package collector

import (
	"time"

	"github.com/go-chassis/huawei-apm/apm"
	"github.com/go-chassis/huawei-apm/common"
)

const (
	Kpi_Collector_Key       = "Kpi_Collector"
	Inventory_Collector_Key = "Inventory_Collector"
)

// Collector
var Collector Collect

type Collect struct {
	Apm map[string]apm.APMI
}

// CreateCollect create collector
func CreateCollect(serverName, kpiUrl, caPath string) {
	Collector.Apm[Kpi_Collector_Key] = apm.NewKpiAPM(serverName, kpiUrl, caPath)
	Collector.Apm[Inventory_Collector_Key] = apm.NewInventoryApm(serverName, kpiUrl, caPath)

}
func init() {
	Collector = Collect{
		Apm: make(map[string]apm.APMI, 2),
	}
	t := time.NewTicker(common.DefaultBatchTime)
	go func() {
		for range t.C {
			Collector.Apm[Kpi_Collector_Key].Send()
			Collector.Apm[Inventory_Collector_Key].Send()
		}
	}()
}
