package kpi

import (
	"github.com/go-chassis/huawei-apm/apm"
	"github.com/go-chassis/huawei-apm/common"
)

func Collect(data common.KPICollectorMessage) error {
	return apm.KpiApmCache.Set(data)
}
