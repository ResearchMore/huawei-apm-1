package inventory

import (
	"github.com/go-chassis/huawei-apm/collector"
	"github.com/go-chassis/huawei-apm/common"
)

func CollectInventory(inventory common.Inventory) error {
	return collector.Collector.Apm[collector.Inventory_Collector_Key].Set(inventory)
}
