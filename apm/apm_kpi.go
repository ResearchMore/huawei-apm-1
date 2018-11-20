package apm

import (
	"crypto/tls"
	"net"
	"net/http"
	"sync"

	"errors"
	"fmt"

	"encoding/json"

	"time"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	cache "github.com/patrickmn/go-cache"
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
	kd, ok := data.(common.TKpiMessage)
	if !ok {
		return NotMatchError
	}

	var tkpimessages []common.TKpiMessage
	cacheKey := getKpiCacheKey(kd.SrcTierName, kd.DestTierName, kd.TransactionType)

	t, ok := k.Get(cacheKey)
	if ok {
		tkpimessages = t
	}

	tkpimessages = append(tkpimessages, kd)
	k.kpiMessagesCaChe.Set(cacheKey, tkpimessages, 0)
	return nil
}
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

	key := utils.GetAPMKey("istio", k.ProjectID, "dafault", "cse", k.ServerName)
	kpiMessageBytes := [][]byte{}
	dataMap := make(map[int64][][]byte, len(items))

	for _, v := range items {
		d, _ := json.Marshal(v.Object)
		kpiMessageBytes = append(kpiMessageBytes, d)
	}

	int64Key := utils.GetTimeMillisecond() - 60*1000
	dataMap[int64Key] = kpiMessageBytes
	tAgentMessage.Messages[key] = dataMap

	// send data to apm
	err := httpDo(k.httpClient, tAgentMessage, k.Url, k.ProjectID)
	if err != nil {
		k.setTOAgentMessage(tAgentMessage)
	}
	return err
}

func (k *KpiApm) Delete(key string) {
	if key == "" {
		k.kpiMessagesCaChe.DeleteExpired()
		return
	}
	k.kpiMessagesCaChe.Delete(key)
}

func NewKpiAPM(projectID, serverName, url, caPath string) *KpiApm {
	if projectID == "" {
		projectID = DefaultProjectID
	}

	if serverName == "" {
		serverName = DefaultSDestination
	}

	if url == "" {
		url = DefaultKPIUrl
	}

	if caPath == "" {
		caPath = defaultCAPath
	}
	conn, err := NewConnection(DefaultKPIUrl)

	if err != nil {
		openlogging.GetLogger().Errorf("get tcp conn failed err : %v", err)
	}

	return &KpiApm{
		ProjectID:        projectID,
		Url:              url,
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

// Get get tKpiMessage from cache
func (k *KpiApm) Get(key string) ([]common.TKpiMessage, bool) {
	d, ok := k.kpiMessagesCaChe.Get(key)
	if !ok {
		return []common.TKpiMessage{}, false
	}
	message, ok := d.([]common.TKpiMessage)
	return message, ok
}

func (k *KpiApm) getAllKpiMessageFromCache() map[string]cache.Item {
	return k.kpiMessagesCaChe.Items()
}

func initCache() *cache.Cache { return cache.New(DefaultExpireTime, CleanupInterval) }

// getKpiCacheKey get kpi apmcache key use source name , dest name , TransactionType
func getKpiCacheKey(sourceName, destName, transactionType string) string {
	return utils.GetAPMKey(sourceName, destName, transactionType)
}

// setTOAgentMessage  when send data to apm failed , use this method set
func (k *KpiApm) setTOAgentMessage(data *common.TAgentMessage) {
	ms := k.getAgentMessageFormCache()
	ms = append(ms, data)
	k.agentMessage.Set(SecondarySend, ms, 0)
}

// getAgentMessageFormCache
func (k *KpiApm) getAgentMessageFormCache() []*common.TAgentMessage {
	d, ok := k.agentMessage.Get(SecondarySend)
	if !ok {
		return []*common.TAgentMessage{}
	}
	return d.([]*common.TAgentMessage)
}

func init() {
	ticker := time.NewTicker(DefaultBatchTime)
	go func() {
		for range ticker.C {
			if KpiApmCache != nil {
				// if has old data sent the old data first and old data only send data again
				ts := KpiApmCache.getAgentMessageFormCache()
				if len(ts) > 0 {
					KpiApmCache.agentMessage.Delete(SecondarySend)
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
