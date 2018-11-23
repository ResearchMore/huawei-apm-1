package apm

import (
	"errors"
)

var (
	NotMatchError = errors.New("set data to kpi , but the data to be set type not kpi")
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
}

// Init init
func Init(serverName, url, caPath string) {
	//fmt.Println("====>1", runtime.InstanceID)
	//fmt.Println("====>2", runtime.App)
	//fmt.Println("====>3", runtime.HostName)
	//fmt.Println("====>4", runtime.InstanceStatus)
	//fmt.Println("====>5", runtime.ServiceID)
	//fmt.Println("====>6", runtime.ServiceName)
	//fmt.Println("====>7", runtime.Version)
	KpiApmCache = NewKpiAPM(serverName, url, caPath)
	InventoryApmCache = NewInventoryApm(serverName, url, caPath)
}
