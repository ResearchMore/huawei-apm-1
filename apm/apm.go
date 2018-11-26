package apm

import (
	"errors"

	"github.com/go-chassis/huawei-apm/common"
)

var (
	KpiNotMatchError       = errors.New("set data to kpi , but the data to be set type not kpi")
	InventoryNotMatchError = errors.New("set data to Inventory , but the data to be set type not Inventory")
)

// APMI
type APMI interface {
	Set(interface{}) error
	Send() error
	Delete(string)
	GetAgentCache() *common.TAgentMessage
}
