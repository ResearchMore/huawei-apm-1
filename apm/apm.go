package apm

import (
	"errors"

	"github.com/go-chassis/huawei-apm/common"
)

var (
	KpiNotMatchError       = errors.New("set data to kpi , but the data to be set type not kpi")
	InventoryNotMatchError = errors.New("set data to Inventory , but the data to be set type not Inventory")
)

// KpiApmCache cache of kpi , cache key use src name,dest name and transaction type
var KpiApmCache *KpiApm

// InventoryApmCache cache of inventory apm ,and the cache save the inventory when sent inventory to apm failed
var InventoryApmCache *InventoryApm

// APMI
type APMI interface {
	Set(interface{}) error
	Send() error
	Delete(string)
	GetAgentCache() *common.TAgentMessage
}

// Init init
func Init(serverName, url, caPath string) {
	KpiApmCache = NewKpiAPM(serverName, url, caPath)
	InventoryApmCache = NewInventoryApm(serverName, url, caPath)
}
