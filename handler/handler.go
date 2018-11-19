package handler

import (
	"time"

	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/go-chassis/go-chassis/pkg/runtime"
	"github.com/go-chassis/huawei-apm/common"
)

// ApmName  name of huawei collector handler
const ApmName = "collector-handler"

//  APMHandler collector struct
type APMHandler struct{}

func init() {
	handler.RegisterHandler(ApmName, New)
}
func New() handler.Handler {
	return &APMHandler{}
}

// Name returns collector-handler string
func (a *APMHandler) Name() string {
	return ApmName
}

// Handle is to handle the collector tracing related things
func (a *APMHandler) Handle(chain *handler.Chain, inv *invocation.Invocation, cb invocation.ResponseCallBack) {
	// start time
	start := time.Now()
	chain.Next(inv, func(response *invocation.Response) (err error) {
		err = cb(response)
		var totalErrorLatency, totalLatency byte
		if response.Status < 200 || response.Status > 300 {
			totalErrorLatency += byte(time.Since(start).Nanoseconds() / 1e6)
		}
		getTKpiMesssage(inv.SourceServiceID, inv.SchemaID, inv.Protocol, runtime.App, inv.SourceMicroService,
			inv.MicroServiceName, totalLatency, totalErrorLatency)
		return
	})
}

// getTKpiMesssage get kpi message
func getTKpiMesssage(sourceResourceID, destResourceID, transactionType,
	appID, srcTierName, destTierName string,
	totalLatency, totalErrorLatency byte) common.TKpiMessage {
	return common.TKpiMessage{
		SourceResourceId: sourceResourceID,
		DestResourceId:   destResourceID,
		TransactionType:  transactionType,
		AppId:            appID,
		SrcTierName:      srcTierName,
		DestTierName:     destTierName,
	}
}
