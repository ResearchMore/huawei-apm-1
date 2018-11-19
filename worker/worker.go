package worker

import (
	"time"
)

const (
	defaultBatchTime = 60 * time.Second
	defaultProjectID = "default"
	//DefaultExpireTime default expiry time is kept as 0
	DefaultExpireTime = 0
	// CleanupInterval default clean up time is kept as 0
	CleanupInterval = 0
)

// KPIWork record kpi Message
var KPIWork *KpiWorker

// DefaultInventoryUrl default url to send inventory Message for collector like istio
const DefaultInventoryUrl = "https://elbIp:8923/%s/inventory/istio"

// DefaultKPIUrl default url send kpi Message to collector
const DefaultKPIUrl = "https://elbIp:8923/%s/kpi/istio"

type Worker interface {
	send() error
	Set() error
}
