package kpi

import (
	"context"

	"time"

	"github.com/go-chassis/huawei-apm/worker"
)

const (
	defaultUrl = "https://elbIp:8923/%s/kpi/istio"
)

// defaultBatchInterval in 60 seconds
const defaultHTTPBatchInterval = 60 * time.Second

// Init initialization WorkKPI
func Init(projectID, serverName string) {
	worker.KPIWork, _ = worker.NewKpiWorker(projectID, serverName)
}

// KpiHandle install this to chassis handle
func KpiHandle(ctx context.Context) {
	//kpidata := GetTKpiMesssage("", "",
	//	"", "", "", "", byte(10), byte(10))

	//worker.WorkKPI.Set(tAgentMessage)
}
