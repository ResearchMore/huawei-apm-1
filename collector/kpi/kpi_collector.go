package kpi

import (
	"github.com/go-chassis/huawei-apm/collector"
	"github.com/go-chassis/huawei-apm/common"
)

// CollectKpi
func CollectKpi(data common.KPICollectorMessage) error {
	return apm_collector.Collector.Apm[apm_collector.Kpi_Collector_Key].Set(data)
}
