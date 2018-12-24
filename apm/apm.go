package apm

import (
	"errors"

	"github.com/go-chassis/huawei-apm/common"
)

var (
	KpiNotMatchError       = errors.New("set data to kpi failed, but the data to be set type not kpi")
	InventoryNotMatchError = errors.New("set data to Inventory failed, but the data to be set type not Inventory")
)

// APM
type APM interface {
	Set(interface{}) error
	Send() error
	Delete(string)
	GetAgentCache() *common.TAgentMessage
}
