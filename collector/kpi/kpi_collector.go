package kpi

import "github.com/go-chassis/huawei-apm/common"

// KPICollector implements Collector by forwarding kpi to a http server
type KPICollector struct {
	batch []common.TKpiMessage
	kpi   chan *common.TKpiMessage
	quit  chan struct{}
}
