package apm_collector

import (
	"time"

	"github.com/go-chassis/huawei-apm/apm"
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

// CreateDefaultCollect use default value to create default apm collector
func CreateDefaultCollect() {
	CreateCollect("", "", "")
}

// CreateCollect create collector
func CreateCollect(serverName, kpiUrl, caPath string) {
	Collector.Apm[Kpi_Collector_Key] = apm.NewKpiAPM(serverName, kpiUrl, caPath)
	Collector.Apm[Inventory_Collector_Key] = apm.NewInventoryApm(serverName, kpiUrl, caPath)
}

// StartCollector when you init collect will start collector
func StartCollector() {
	// make  goroutine to send kpi and inventory data
	for k := range Collector.Apm {
		t := time.NewTicker(10 * time.Second)
		//t := time.NewTicker(common.DefaultBatchTime)
		go func(k string) {
			for range t.C {
				// 启动判断
				if Collector.Apm[k] != nil {
					Collector.Apm[k].Send()
				}
			}
		}(k)
	}
}
func init() {
	Collector = Collect{
		Apm: make(map[string]apm.APMI, 2),
	}

}
