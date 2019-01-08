package common

import "time"

// const for zipkin
const (
	APMKeySeparator = "|"
)
const (
	DefaultK8sCrtFileName = "kubecfg.crt"
	DefaultK8sKeyFileName = "kubecfg_crypto.key"
	DefaultCAPath         = "/var/paas/srv/kubernetes"
	DefaultClusterKey     = "CLUSTER_KEY_MESHER_APM"
	DefaultElbIP          = "ELBURL_MESHER_APM"
)
const (
	DefaultBatchTime = 5 * time.Second
	DefaultProjectID = "default"
	//DefaultExpireTime default expiry time is kept as 0
	DefaultExpireTime = 0
	// CleanupInterval default clean up time is kept as 0
	CleanupInterval = 0

	DefaultClient       = "unknownClient"
	DefaultSDestination = "unknownDestination"
	DefaultCluster      = "default"
	DefaultServerName   = "default"
	// need to secondary transmission
	SecondarySend = "SecondarySend"
	EnvProjectID  = "CSE_PROJECT_ID"
)

const (
	INTERMEDIATE = iota + 1
	FIRST_FOR_BACKEND
	FIRST_FOR_CLIENT
	FIRST_FOR_UNKNOWN
	FIRST_FOR_ENDPOINT
	ENDPOINT
	EXTERNAL
	FIRST_EXTERNAL
	ENDPOINT_EXTERNAL
	ISTIO
	ISTIO_EXTERNAL
)
