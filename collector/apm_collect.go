package apm_collector

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
	// make to goroutine to send kpi and inventory data
	for k := range Collector.Apm {
		go func() {
			for range t.C {
				if Collector.Apm[k] != nil {
					Collector.Apm[k].Send()
				}
			}
		}()
	}
}
