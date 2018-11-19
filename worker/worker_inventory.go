package worker

import (
	"net/http"
	"sync"

	"github.com/go-chassis/huawei-apm/common"
)

// InventoryWorker implement Worker by forwarding worker to Inventory worker
type InventoryWorker struct {
	inventoryRWMutex *sync.RWMutex
	httpClient       *http.Client
	Message          common.Inventory
	Url              string
	ProjectID        string
	ServerName       string
}
