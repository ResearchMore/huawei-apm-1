package pod

import (
	"os"

	"github.com/go-chassis/huawei-apm/common"
)

const (
	DefaultClusterID   = "default"
	DefaultContainerID = "default"
)

// GetClusterId
func GetClusterId() string {
	return DefaultClusterID
}

// GetPodName get pod name for this value,default value is "unknownClient"
// pod name use as src id
func GetPodName() string {
	PodName := os.Getenv("POD_NAME")
	if PodName != "" {
		return PodName
	}
	return common.DefaultClient
}

// GetContainerID
func GetContainerID() string {
	return DefaultContainerID
}
