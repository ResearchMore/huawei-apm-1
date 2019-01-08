package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/core/handler"
	"github.com/go-chassis/go-chassis/core/invocation"
	"github.com/go-chassis/go-chassis/pkg/runtime"
	"github.com/go-chassis/huawei-apm/collector/inventory"
	"github.com/go-chassis/huawei-apm/collector/kpi"
	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/pod"
	"github.com/go-chassis/huawei-apm/utils"
)

// ApmName  name of huawei collector handler
const ApmName = "apm-handler"

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
	transactionType := inv.Protocol
	var callbackFunc func(*invocation.Response, time.Time) error
	finish := make(chan *invocation.Response, 1)

	chain.Next(inv, func(response *invocation.Response) error {
		err := response.Err
		select {
		case finish <- response:
		default:
		}
		return err
		//return cb(response)
	})

	callbackFunc = func(response *invocation.Response, t time.Time) error {

		var totalErrorLatency, totalLatency int64
		if response.Err != nil {
			totalErrorLatency = time.Since(t).Nanoseconds() / 1e6
		} else {
			if transactionType == "rpc" {
				totalLatency = time.Since(t).Nanoseconds() / 1e6
			} else {
				transactionType = fmt.Sprintf("%s%s", inv.Endpoint, inv.URLPathFormat)
				resp, _ := response.Result.(*http.Response)
				if resp.StatusCode < 200 || resp.StatusCode > 300 {
					totalErrorLatency = time.Since(t).Nanoseconds() / 1e6
				} else {
					totalLatency = time.Since(t).Nanoseconds() / 1e6
				}
			}
		}

		kpiCollectorMessage := getKPIMessage(pod.GetPodName(), getDestResourceId(),
			transactionType, runtime.ServiceName,
			inv.MicroServiceName, totalLatency, totalErrorLatency)
		kpi.CollectKpi(kpiCollectorMessage)

		in := getInventoryMessage(runtime.HostName, utils.GetLocalIP(), runtime.App,
			"servicemesh_displayName", runtime.InstanceID, pod.GetContainerID(),
			inv.MicroServiceName, config.MicroserviceDefinition.ServiceDescription.Name, 0)

		inventory.CollectInventory(in)
		return cb(response)
	}

	callbackFunc(<-finish, start)
}

// getKPIMessage get kpi message
func getKPIMessage(sourceResourceID, destResourceID, transactionType,
	srcTierName, destTierName string,
	totalLatency, totalErrorLatency int64) common.KPICollectorMessage {
	clusterKey := utils.GetClusterID()
	return common.KPICollectorMessage{
		SourceResourceId:  sourceResourceID,
		DestResourceId:    destResourceID,
		TransactionType:   transactionType,
		AppId:             clusterKey,
		SrcTierName:       srcTierName,
		DestTierName:      destTierName,
		TotalLatency:      totalLatency,
		TotalErrorLatency: totalErrorLatency,
	}
}

// getInventoryMessage return new inventory
func getInventoryMessage(hostname, ip, appName,
	displayName, instanceName, containerID, tier, namespaceName string, pid int32) common.TDiscoveryInfo {
	clusterKey := utils.GetClusterID()
	projectID := utils.GetProjectID()
	return common.TDiscoveryInfo{
		CollectorId:   "servicemesh",
		ProjectId:     projectID,
		Hostname:      hostname,
		IP:            ip,
		AgentId:       "servicemesh",
		AppName:       appName,
		ClusterKey:    clusterKey,
		ServiceType:   "servicemesh",
		DisplayName:   displayName,
		InstanceName:  instanceName,
		ContainerId:   containerID,
		Pid:           pid,
		AppId:         clusterKey,
		Tier:          tier,
		NamespaceName: namespaceName,
		Created:       utils.GetTimeMillisecond(),
		Updated:       utils.GetTimeMillisecond(),
	}
}

// getDestResourceId get pod name of endpoint , default name  "unknownDestination"
func getDestResourceId() string {
	return common.DefaultSDestination
}
