package apm

import (
	"time"

	"errors"
)

var (
	NotMatchError = errors.New("set data to kpi , but the data to be set type not kpi")
)

// KpiApmCache cache of kpi , cache key use src name,dest name and transaction type
var KpiApmCache *KpiApm

const (
	defaultCAPath    = "/var/paas/srv/kubernetes"
	DefaultBatchTime = 60 * time.Second
	DefaultProjectID = "default"
	//DefaultExpireTime default expiry time is kept as 0
	DefaultExpireTime = 0
	// CleanupInterval default clean up time is kept as 0
	CleanupInterval     = 0
	DefaultClient       = "unknownClient"
	DefaultSDestination = "unknownDestination"
	// need to secondary transmission
	SecondarySend = "SecondarySend"
)

// APMI
type APMI interface {
	Set(interface{}) error
	Send() error
	Delete(string)
}

// Init init
func Init(projectID, serverName, url, caPath string) {
	KpiApmCache = NewKpiAPM(projectID, serverName, url, caPath)
}
