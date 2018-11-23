package common

import "time"

// const for zipkin
const (
	APMKeySeparator = "|"
)
const (
	DefaultCAPath    = "/var/paas/srv/kubernetes"
	DefaultBatchTime = 60 * time.Second
	DefaultProjectID = "default"
	//DefaultExpireTime default expiry time is kept as 0
	DefaultExpireTime = 0
	// CleanupInterval default clean up time is kept as 0
	CleanupInterval = 0

	DefaultClient       = "unknownClient"
	DefaultSDestination = "unknownDestination"

	DefaultServerName = "default"
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
