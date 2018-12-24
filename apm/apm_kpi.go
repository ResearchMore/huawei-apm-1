package apm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	"github.com/patrickmn/go-cache"
)

// DefaultKPIUrl default url send kpi Message to collector
const DefaultKPIUrl = "/svcstg/ats/v1/%s/kpi/istio"

// KpiApm implement APM by forwarding kpi
type KpiApm struct {
	kpiMutex         *sync.Mutex
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
		return KpiNotMatchError
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

// Send implement APM send method , send kpi message to apm , when the kpi message failed to be send apm
// will use cache to storing the kpi message , but the data will only send second time
func (k *KpiApm) Send() error {
	// if has old data sent the old data first and old data only send second time
	amc := k.GetAgentCache()
	if amc != nil {
		k.agentMessage.Flush()
		err := httpDo(k.httpClient, amc, k.Url, k.ProjectID)
		if err != nil {
			openlogging.GetLogger().Errorf("send data again  failed: ,error : [%v]", err)
		}
	}
	// send new data
	k.kpiMutex.Lock()
	items := k.getAllKpiMessageFromCache()
	k.Delete("")
	k.kpiMutex.Unlock()
	if len(items) < 1 {
		//openlogging.GetLogger().Warnf("not kpi message need to send in cache , cache num is : %d", len(items))
		return errors.New(fmt.Sprintf("not kpi message need to send in cache , cache num is : %d", len(items)))
	}

	kpiMessageBytes := [][]byte{}

	for _, v := range items {
		c := v.Object.(common.KPICollectorMessage)
		tKpiMessage := getTKpiMessage(c)
		d, _ := json.Marshal(tKpiMessage)
		kpiMessageBytes = append(kpiMessageBytes, d)
	}

	if len(kpiMessageBytes) < 1 {
		openlogging.GetLogger().Warnf("not kpi message need to send in cache , cache num is : %d", len(items))
		return errors.New(fmt.Sprintf("not kpi message need to send in cache , cache num is : %d", len(items)))
	}

	key := utils.GetAPMKey("istio", k.ProjectID, "default", "cse", k.ServerName)
	int64Key := utils.GetTimeMillisecond() - 60*1000

	tAgentMessage := &common.TAgentMessage{
		AgentContext: utils.UUID16(),
		Messages: map[string]map[int64][][]byte{
			key: {
				int64Key: kpiMessageBytes,
			},
		},
	}

	// send data to apm
	err := httpDo(k.httpClient, tAgentMessage, k.Url, k.ProjectID)
	if err != nil {
		k.setToAgentMessage(tAgentMessage)
	}
	return err
}

// GetAgentCache get agent message , the message is last time send failed
func (k *KpiApm) GetAgentCache() *common.TAgentMessage {
	d, ok := k.agentMessage.Get(common.SecondarySend)
	if !ok {
		return nil
	}
	return d.(*common.TAgentMessage)
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
	var reply common.KPICollectorMessage

	d, ok := k.kpiMessagesCaChe.Get(key)
	if !ok {
		return reply, false
	}
	reply, ok = d.(common.KPICollectorMessage)
	return reply, ok
}

// getAllKpiMessageFromCache method get all cache data form cache
func (k *KpiApm) getAllKpiMessageFromCache() map[string]cache.Item {
	return k.kpiMessagesCaChe.Items()
}

// setToAgentMessage  when send data to apm failed , use this method set
func (k *KpiApm) setToAgentMessage(data *common.TAgentMessage) {
	k.agentMessage.SetDefault(common.SecondarySend, data)
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
		SpanType:          c.SpanType,
	}
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

	tlsConfig, err := utils.GetTLSConfig(caPath, "", "")
	if err != nil {
		openlogging.GetLogger().Errorf("apm kpi: get tls config failed,err[%s]", err)
		return nil
	}
	return &KpiApm{
		ProjectID:        projectID,
		Url:              kpiUrl,
		ServerName:       serverName,
		kpiMessagesCaChe: initCache(),
		agentMessage:     initCache(),
		kpiMutex:         &sync.Mutex{},
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		},
	}
}

// getKpiCacheKey get kpi apmcache key use source name , dest name , TransactionType
func getKpiCacheKey(sourceName, destName, transactionType string) string {
	return utils.GetAPMKey(sourceName, destName, transactionType)
}
