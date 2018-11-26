package apm

import (
	"testing"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/stretchr/testify/assert"
)

var kpiApm *KpiApm

var collectorMessages []common.KPICollectorMessage

var agentMessageKpi common.TAgentMessage

func init() {
	kpiApm = NewKpiAPM("", "", "")
	collectorMessages = []common.KPICollectorMessage{
		{
			SourceResourceId:   "apm_source_id01",
			DestResourceId:     "apm_dest_id01",
			TransactionType:    "apm_transaction_type01",
			AppId:              "apm_app_id01",
			SrcTierName:        "apm_src_tier_name01",
			DestTierName:       "apm_dest_tier_name01",
			TotalErrorLatencys: []int64{1, 2, 3, 4, 5},
			TotalErrorLatency:  6,
			TotalLatencys:      []int64{8, 9, 10, 11, 12},
			TotalLatency:       7,
		},
		{
			SourceResourceId:   "apm_source_id02",
			DestResourceId:     "apm_dest_id02",
			TransactionType:    "apm_transaction_type02",
			AppId:              "apm_app_id02",
			SrcTierName:        "apm_src_tier_name02",
			DestTierName:       "apm_dest_tier_name02",
			TotalErrorLatencys: []int64{1, 2, 3, 4, 5},
			TotalErrorLatency:  6,
			TotalLatencys:      []int64{8, 9, 10, 11, 12},
			TotalLatency:       7,
		},
	}
	agentMessageKpi = common.TAgentMessage{
		AgentContext: "apm_anent_context",
		TenantName:   "apm_tenant_name",
		Messages: map[string]map[int64][][]byte{
			"apm_map_01": {123456789: [][]byte{}},
		},
	}
}
func TestKpiApm_Set(t *testing.T) {
	var keys []string
	for _, v := range collectorMessages {
		// test set method
		err := kpiApm.Set(v)
		assert.NoError(t, err)
		keys = append(keys, utils.GetAPMKey(v.SrcTierName, v.DestTierName, v.TransactionType))
	}

	for _, v := range keys {
		// test get method
		mCache, ok := kpiApm.Get(v)
		assert.Equal(t, ok, true)
		assert.NotNil(t, mCache)
	}

	// get all cache data
	ms := kpiApm.getAllKpiMessageFromCache()
	assert.NotNil(t, ms)

	// delete
	//kpiApm.Delete(key)
	kpiApm.Delete("")
	ms = kpiApm.getAllKpiMessageFromCache()
	assert.Empty(t, ms)

	for _, v := range keys {
		// test get method
		mCache, ok := kpiApm.Get(v)
		assert.Equal(t, ok, false)
		assert.Empty(t, mCache)

	}

}
func TestAgent(t *testing.T) {
	kpiApm.setToAgentMessage(&agentMessageKpi)
	agentMessages := kpiApm.GetAgentCache()
	assert.NotNil(t, agentMessages)
}
