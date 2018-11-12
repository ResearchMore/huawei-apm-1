package worker

import "github.com/huaweicse/huawei-apm/common"

const (
	defaultWorker = 10
	defaultPool = 3
)
// WorkerChan is collector , when len of WorkerChan greater or equal to 10 will sent
// it to APM server
var workerChan chan common.TAgentMessage

// SetWorkerChan set apm data into worker chan when it over 10 will sent it to apm server
func SetWorkerChan(f common.TAgentMessage)  {
	workerChan <- f
}
