package apm

import (
	"encoding/json"

	"net/http"

	"crypto/tls"
	"os"

	"errors"

	"sync"

	"time"

	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	"github.com/patrickmn/go-cache"
)

// DefaultKPIUrl default url send kpi Message to collector
const DefaultInventoryUrl = "/svcstg/ats/v1/%s/inventory/istio"

// InventoryCacheKey key for inventory cache key
const InventoryCacheKey string = "inventory_apm_cache_key"

type InventoryApm struct {
	// if send apm failed set data to this
	agentMessage   *cache.Cache
	inventoryCache *cache.Cache
	mutexInventory *sync.Mutex
	httpClient     *http.Client
	Url            string
	ProjectID      string
	ServerName     string
	KeyString      string
	KeyInt64       int64
}

func (i *InventoryApm) Set(data interface{}) error {
	in, ok := data.(common.Inventory)
	if !ok {
		return InventoryNotMatchError
	}
	var inventories []common.Inventory
	iCache, ok := i.Get(InventoryCacheKey)
	if ok {
		inventories = iCache
	}
	inventories = append(inventories, in)
	i.inventoryCache.SetDefault(InventoryCacheKey, inventories)

	return nil
}
func (i *InventoryApm) Send() error {
	// if has old data in agent cache will sent it again
	agent := i.GetAgentCache()
	if agent != nil {
		i.agentMessage.Flush()
		err := httpDo(i.httpClient, agent, i.Url, i.ProjectID)
		if err != nil {
			openlogging.GetLogger().Errorf("send data again  failed: ,error : [%v]", err)
		}
	}

	var datas [][]byte

	i.mutexInventory.Lock()
	inventories, ok := i.Get(InventoryCacheKey)
	i.Delete(InventoryCacheKey)
	i.mutexInventory.Unlock()

	if !ok {
		//openlogging.GetLogger().Warn("not data need to send apm")
		return errors.New("not data need to send apm")
	}
	for _, v := range inventories {
		temp, _ := json.Marshal(v)
		datas = append(datas, temp)
	}
	if len(datas) == 0 {
		openlogging.GetLogger().Warn("not data need to send apm")
		return errors.New("not data need to send apm")
	}
	i.KeyString = utils.GetAPMKey("istio", i.ProjectID, "default", "cse", i.ServerName)
	i.KeyInt64 = utils.GetTimeMillisecond() - 60*1000
	tAgentMessage := &common.TAgentMessage{
		AgentContext: utils.UUID16(),
		Messages: map[string]map[int64][][]byte{
			i.KeyString: {
				i.KeyInt64: datas,
			},
		},
	}
	err := httpDo(i.httpClient, tAgentMessage, i.Url, i.ProjectID)
	if err != nil {
		i.agentMessage.SetDefault(common.SecondarySend, tAgentMessage)
	}
	return err
}

// Delete delete inventory cache by key , when key is empty will delete all cache of inventory
func (i *InventoryApm) Delete(key string) {
	if key == "" {
		i.inventoryCache.Flush()
		return
	}
	i.inventoryCache.Delete(key)
}

// GetAgentCache get agent message , the message is last time send failed
func (i *InventoryApm) GetAgentCache() *common.TAgentMessage {
	d, ok := i.agentMessage.Get(common.SecondarySend)
	if !ok {
		return nil
	}
	return d.(*common.TAgentMessage)
}

// Get get tKpiMessage from cache
func (k *InventoryApm) Get(key string) ([]common.Inventory, bool) {
	d, ok := k.inventoryCache.Get(key)
	if !ok {
		return nil, false
	}
	message, ok := d.([]common.Inventory)
	return message, ok
}

// NewInventoryApm return new InventoryApm
func NewInventoryApm(serverName, inventoryUrl, caPath string) *InventoryApm {
	projectID, isExist := os.LookupEnv(common.EnvProjectID)
	if !isExist {
		projectID = common.DefaultProjectID
	}

	inventoryUrl = utils.GetStringWithDefaultName(inventoryUrl, DefaultInventoryUrl)
	caPath = utils.GetStringWithDefaultName(caPath, common.DefaultCAPath)
	serverName = utils.GetStringWithDefaultName(serverName, common.DefaultServerName)

	return &InventoryApm{
		ServerName:     serverName,
		Url:            inventoryUrl,
		ProjectID:      projectID,
		agentMessage:   initCache(),
		inventoryCache: initCache(),
		mutexInventory: &sync.Mutex{},
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					ClientAuth:   tls.RequireAndVerifyClientCert,
					RootCAs:      utils.GetX509CACertPool(caPath, ""),
					Certificates: []tls.Certificate{utils.GetCertificate(caPath, "", "")},
				},
			},
		},
	}
}
