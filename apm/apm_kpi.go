package apm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"crypto/tls"
	"os"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	"github.com/patrickmn/go-cache"
)

// DefaultKPIUrl default url send kpi Message to collector
const DefaultKPIUrl = "https://elbIp:8923/%s/kpi/istio"

// KpiApm implement APMI by forwarding kpi
type KpiApm struct {
	kpiMutex         *sync.Mutex
	tcpConn          *net.TCPConn
	httpClient       *http.Client
	kpiMessagesCaChe *cache.Cache
	// if send apm failed set data to this
	agentMessage *cache.Cache
	Url          string
	ProjectID    string
	ServerName   string
}

// Set method set data to api cache
func (k *KpiApm) Set(data interface{}) error {
	// if data not kpi type,return error
	collector, ok := data.(common.KPICollectorMessage)
	if !ok {
		return NotMatchError
	}

	var kpiCollector = collector
	cacheKey := getKpiCacheKey(collector.SrcTierName, collector.DestTierName, collector.TransactionType)

	temp, ok := k.Get(cacheKey)
	if ok {
		kpiCollector = temp
	}
	if collector.TotalErrorLatency != 0 {
		kpiCollector.TotalErrorLatencys = append(kpiCollector.TotalErrorLatencys, collector.TotalErrorLatency)
	}
	if collector.TotalLatency != 0 {
		kpiCollector.TotalLatencys = append(kpiCollector.TotalLatencys, collector.TotalLatency)
	}

	k.kpiMessagesCaChe.Set(cacheKey, kpiCollector, 0)
	return nil
}

// Send implement APMI send method , send kpi message to apm , when the kpi message failed to be send apm
// will use cache to storing the kpi message , but the data will only send second time
func (k *KpiApm) Send() error {
	k.kpiMutex.Lock()
	items := k.getAllKpiMessageFromCache()
	k.Delete("")
	k.kpiMutex.Unlock()
	if len(items) < 1 {
		openlogging.GetLogger().Errorf("not kpi message need to send in cache , cache num is : %d", len(items))
		return errors.New(fmt.Sprintf("not kpi message need to send in cache , cache num is : %d", len(items)))
	}

	tAgentMessage := &common.TAgentMessage{
		AgentContext: utils.UUID16(),
		Messages:     make(map[string]map[int64][][]byte),
	}

	key := utils.GetAPMKey("istio", k.ProjectID, "default", "cse", k.ServerName)
	kpiMessageBytes := [][]byte{}
	dataMap := make(map[int64][][]byte, len(items))

	for _, v := range items {

		c := v.Object.(common.KPICollectorMessage)
		tKpiMessage := getTKpiMessage(c)
		d, _ := json.Marshal(tKpiMessage)
		kpiMessageBytes = append(kpiMessageBytes, d)
	}

	int64Key := utils.GetTimeMillisecond() - 60*1000
	dataMap[int64Key] = kpiMessageBytes
	tAgentMessage.Messages[key] = dataMap

	// send data to apm
	err := httpDo(k.httpClient, tAgentMessage, k.Url, k.ProjectID)
	if err != nil {
		k.setToAgentMessage(tAgentMessage)
	}
	return err
}

// getTKpiMessage method get tKpiMessage with kpi collector
func getTKpiMessage(c common.KPICollectorMessage) common.TKpiMessage {
	var totalLatency []byte
	var totalErrorLatency []byte
	if len(c.TotalLatencys) > 0 {
		totalLatency, _ = json.Marshal(c.TotalLatencys)
	}
	if len(c.TotalErrorLatencys) > 0 {
		totalErrorLatency, _ = json.Marshal(c.TotalErrorLatencys)
	}
	return common.TKpiMessage{
		SourceResourceId:  c.SourceResourceId,
		DestResourceId:    c.DestResourceId,
		TransactionType:   c.TransactionType,
		AppId:             c.AppId,
		SrcTierName:       c.SrcTierName,
		DestTierName:      c.DestTierName,
		TotalErrorLatency: totalErrorLatency,
		TotalLatency:      totalLatency,
	}
}

// Delete delete tAgentMessage for tAgentMessageCache , when the key is empty will delete all cache data,
// key not empty will delete cache data by key
func (k *KpiApm) Delete(key string) {
	if key == "" {
		k.kpiMessagesCaChe.Flush()

		return
	}
	k.kpiMessagesCaChe.Delete(key)
}

// Get get tKpiMessage from cache
func (k *KpiApm) Get(key string) (common.KPICollectorMessage, bool) {
	d, ok := k.kpiMessagesCaChe.Get(key)
	if !ok {
		return common.KPICollectorMessage{}, false
	}
	message, ok := d.(common.KPICollectorMessage)
	return message, ok
}

// getAllKpiMessageFromCache method get all cache data form cache
func (k *KpiApm) getAllKpiMessageFromCache() map[string]cache.Item {
	return k.kpiMessagesCaChe.Items()
}

// setToAgentMessage  when send data to apm failed , use this method set
func (k *KpiApm) setToAgentMessage(data *common.TAgentMessage) {
	ms := k.getAgentMessageFormCache()
	ms = append(ms, data)
	k.agentMessage.Set(common.SecondarySend, ms, 0)
}

// getAgentMessageFormCache
func (k *KpiApm) getAgentMessageFormCache() []*common.TAgentMessage {
	d, ok := k.agentMessage.Get(common.SecondarySend)
	if !ok {
		return []*common.TAgentMessage{}
	}
	return d.([]*common.TAgentMessage)
}

// initCache
func initCache() *cache.Cache { return cache.New(common.DefaultExpireTime, common.CleanupInterval) }

// NewKpiAPM return new KpiAPM with projectID , server name , KpiUrl , caPath
// when input params is empty will use default value.
// e.g. projectID is empty , use default value is "default" , more default please see system const
func NewKpiAPM(serverName, kpiUrl, caPath string) *KpiApm {

	projectID, isExist := os.LookupEnv(common.EnvProjectID)
	if !isExist {
		projectID = common.DefaultProjectID
	}

	serverName = utils.GetStringWithDefaultName(serverName, common.DefaultServerName)
	kpiUrl = utils.GetStringWithDefaultName(kpiUrl, DefaultKPIUrl)
	caPath = utils.GetStringWithDefaultName(caPath, common.DefaultCAPath)

	conn, err := NewConnection(kpiUrl, projectID)

	if err != nil {
		openlogging.GetLogger().Errorf("get tcp conn failed err : %v", err)
	}

	return &KpiApm{
		ProjectID:        projectID,
		Url:              kpiUrl,
		ServerName:       serverName,
		kpiMessagesCaChe: initCache(),
		agentMessage:     initCache(),
		kpiMutex:         &sync.Mutex{},
		tcpConn:          conn,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      utils.GetCertPool(caPath, ""),
					Certificates: []tls.Certificate{utils.GetCertificate(caPath, "", "")},
				},
			},
		},
	}
}

// getKpiCacheKey get kpi apmcache key use source name , dest name , TransactionType
func getKpiCacheKey(sourceName, destName, transactionType string) string {
	return utils.GetAPMKey(sourceName, destName, transactionType)
}
func init() {
	ticker := time.NewTicker(common.DefaultBatchTime)
	go func() {
		for range ticker.C {
			if KpiApmCache != nil {
				// if has old data sent the old data first and old data only send second time
				ts := KpiApmCache.getAgentMessageFormCache()
				if len(ts) > 0 {
					KpiApmCache.agentMessage.Delete(common.SecondarySend)
					for _, v := range ts {
						err := httpDo(KpiApmCache.httpClient, v, KpiApmCache.Url, KpiApmCache.ProjectID)
						openlogging.GetLogger().Errorf("send data again  failed: [%v],error : [%v]", v, err)
					}
				}
				// send new data
				KpiApmCache.Send()
			} else {

			}
		}
	}()
}
