package worker

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	"github.com/patrickmn/go-cache"
)

// KpiWorker implement Worker by forwarding worker to kpi worker
type KpiWorker struct {
	kpiMutex         *sync.Mutex
	tcpConn          *net.TCPConn
	httpClient       *http.Client
	kpiMessages      chan map[string]common.TKpiMessage
	kpiMessagesCaChe chan *cache.Cache
	Message          common.TKpiMessage
	ParentID         *int64
	Max              int
	Url              string
	ProjectID        string
	ServerName       string
}

// Set add new kpi message to kpi worker
func (k *KpiWorker) Set() error {

	var km = make(map[string]common.TKpiMessage)
	key := getKpiCacheKey(k.Message.SrcTierName, k.Message.DestTierName, k.Message.TransactionType)

	k.kpiMutex.Lock()
	if len(k.kpiMessages) > 0 {
		km = <-k.kpiMessages
	}
	if v, ok := km[key]; ok {
		v.TotalLatency = append(v.TotalLatency, k.Message.TotalLatency...)
		v.TotalErrorLatency = append(v.TotalErrorLatency, k.Message.TotalErrorLatency...)
		km[key] = v
	} else {
		km[key] = k.Message
	}

	k.kpiMessages <- km
	k.kpiMutex.Unlock()
	return nil
}

// send sent data to collector , sent once every minute.
func (k *KpiWorker) send() error {
	// if has data in  kpiMessages , will lock the data . unlock  when data has been taken
	k.kpiMutex.Lock()
	if len(k.kpiMessages) < 1 {
		openlogging.GetLogger().Warnf("apmcache no message need to sent")
		return errors.New("apmcache no message need to sent")
	}
	msgs := <-k.kpiMessages
	k.kpiMutex.Unlock()

	if len(msgs) < 1 {
		openlogging.GetLogger().Warnf("apmcache no message need to sent")
		return errors.New("apmcache no message need to sent")
	}

	tAgentMessage := common.TAgentMessage{
		AgentContext: utils.UUID16(),
		Messages:     make(map[string]map[int64][][]byte),
	}

	key := utils.EncryptionMD5(utils.GetAPMKey("istio", k.ProjectID, "default", "cse", k.ServerName))
	int64Key := utils.GetTimeMillisecond() - 60*1000
	datas := [][]byte{}
	dataMap := make(map[int64][][]byte)

	for _, v := range msgs {
		data, _ := json.Marshal(v)
		datas = append(datas, data)
	}

	dataMap[int64Key] = datas
	tAgentMessage.Messages[key] = dataMap

	// send message to collector
	return httpDo(k.httpClient, &tAgentMessage, k.Url, k.ProjectID)
}

// NewKpiWorker the method return new kpiWorker
func NewKpiWorker(projectID, serverName string) (*KpiWorker, error) {

	if projectID == "" {
		projectID = defaultProjectID
	}
	conn, err := NewConnection(DefaultKPIUrl)

	if err != nil {
		openlogging.GetLogger().Errorf("get tcp conn failed err : %v", err)
		return nil, err
	}
	return &KpiWorker{
		Url:              DefaultKPIUrl,
		ProjectID:        projectID,
		ServerName:       serverName,
		kpiMessages:      make(chan map[string]common.TKpiMessage, 1),
		kpiMessagesCaChe: make(chan *cache.Cache, 1),
		tcpConn:          conn,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      getCAs(""),
					Certificates: []tls.Certificate{},
				},
			},
		},
		kpiMutex: &sync.Mutex{},
	}, nil
}

func init() {

	// definite time to sent message to collector
	go func() {
		ticker := time.NewTicker(defaultBatchTime)
		for range ticker.C {
			KPIWork.send()
		}
	}()
}

// getKpiCacheKey get kpi apmcache key use source name , dest name , TransactionType
func getKpiCacheKey(sourceName, destName, transactionType string) string {
	return utils.GetAPMKey(sourceName, destName, transactionType)
}
