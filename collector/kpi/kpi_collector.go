package kpi

import (
	"github.com/go-chassis/huawei-apm/collector"
	"github.com/go-chassis/huawei-apm/common"
)

// CollectKpi
func CollectKpi(data common.KPICollectorMessage) error {

	return collector.Collector.Apm[collector.Kpi_Collector_Key].Set(data)
}
