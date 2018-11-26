package inventory

import (
	"github.com/go-chassis/huawei-apm/collector"
	"github.com/go-chassis/huawei-apm/common"
)

func CollectInventory(inventory common.Inventory) error {
	return apm_collector.Collector.Apm[apm_collector.Inventory_Collector_Key].Set(inventory)
}
