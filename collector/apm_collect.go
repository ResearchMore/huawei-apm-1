package apm_collector

import (
	"time"

	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/huawei-apm/apm"
	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-mesh/openlogging"
)

const (
	Kpi_Collector_Key       = "Kpi_Collector"
	Inventory_Collector_Key = "Inventory_Collector"
)

// Collector
var Collector Collect

type Collect struct {
	Apm map[string]apm.APM
}

// CreateDefaultCollect use default value to create default apm collector
func CreateDefaultCollect() error {
	return CreateCollect("", "", "")
}

// CreateCollect create collector
func CreateCollect(serverName, url, caPath string) error {
	kpiApm, err := apm.NewKpiAPM(serverName, url, "")
	if err != nil {
		return err
	}
	Collector.Apm[Inventory_Collector_Key] = kpiApm
	inven, err := apm.NewKpiAPM(serverName, url, "")
	if err != nil {
		return err
	}
	Collector.Apm[Kpi_Collector_Key] = inven
	return err
}

// StartCollector when you init collect will start collector
func StartCollector(serverName, url, caPath string) {
	if !config.GlobalDefinition.Cse.APM.Enable {
		openlogging.GetLogger().Warn("apm collect not enable")
		return
	}
	err := CreateCollect(serverName, url, caPath)
	if err != nil {
		openlogging.GetLogger().Errorf("create collect for apm failed err :%+v\n", err)
		return
	}
	// make  goroutine to send kpi and inventory data
	for k := range Collector.Apm {
		t := time.NewTicker(common.DefaultBatchTime)
		go func(k string) {
			for range t.C {
				if !config.GlobalDefinition.Cse.APM.Enable {
					openlogging.GetLogger().Warn("apm collect not enable")
					return
				}
				Collector.Apm[k].Send()
			}
		}(k)
	}
}
func init() {
	Collector = Collect{
		Apm: make(map[string]apm.APM, 2),
	}
}
