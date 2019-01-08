package apm

import (
	"testing"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	"github.com/stretchr/testify/assert"
)

var inventoryApm *InventoryApm

var inventories []common.TDiscoveryInfo

func init() {
	inventoryApm = NewInventoryApm("", "", "")
	inventories = []common.TDiscoveryInfo{
		{
			Hostname:     "apm_hostname_01",
			IP:           "127.0.0.1",
			AppId:        "apm_appID_01",
			AppName:      "apm_appName_01",
			ServiceType:  "apm_serviceType_01",
			DisplayName:  "apm_displayName_01",
			InstanceName: "apm_instanceName_01",
			ContainerId:  "apm_containerID_01",
			Pid:          0,
			Props:        nil,
			Created:      utils.GetTimeMillisecond(),
		},
		{
			Hostname:     "apm_hostname_02",
			IP:           "127.0.0.1",
			AppId:        "apm_appID_02",
			AppName:      "apm_appName_0",
			ServiceType:  "apm_serviceType_0",
			DisplayName:  "apm_displayName_02",
			InstanceName: "apm_instanceName_02",
			ContainerId:  "apm_containerID_02",
			Pid:          1,
			Props:        nil,
			Created:      utils.GetTimeMillisecond(),
		},
	}

}

func TestInventoryApm(t *testing.T) {
	openlogging.GetLogger().Info("test about inventory method start")

	var err error
	for _, v := range inventories {
		err = inventoryApm.Set(v)
		assert.NoError(t, err)
	}

	is, ok := inventoryApm.Get(InventoryCacheKey)

	assert.Equal(t, ok, true)
	assert.Equal(t, is, inventories)

	// delete
	inventoryApm.Delete(InventoryCacheKey)

	is, ok = inventoryApm.Get(InventoryCacheKey)

	assert.Equal(t, ok, false)
	assert.Nil(t, is)

}
