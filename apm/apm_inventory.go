package apm

import (
	"encoding/json"

	"net/http"

	"crypto/tls"
	"os"

	"errors"

	"time"

	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/pkg/runtime"
	"github.com/go-chassis/huawei-apm/common"
	"github.com/go-chassis/huawei-apm/pod"
	"github.com/go-chassis/huawei-apm/utils"
	"github.com/go-mesh/openlogging"
	"github.com/patrickmn/go-cache"
)

// DefaultKPIUrl default url send kpi Message to collector
const DefaultInventoryUrl = "https://elbIp:8923/%s/inventory/istio"

// InventoryCacheKey key for inventory cache key
const InventoryCacheKey string = "inventory_apm_cache_key"

type InventoryApm struct {
	// if send apm failed set data to this
	agentMessage   *cache.Cache
	httpClient     *http.Client
	inventory      common.Inventory
	inventoryCache *cache.Cache
	Url            string
	ProjectID      string
	ServerName     string
	KeyString      string
	KeyInt64       int64
}

func (i *InventoryApm) Set(data interface{}) error {
	in, ok := data.(common.Inventory)
	if !ok {
		return errors.New("set data to inventory failed , because the type of date not common.Inventory")
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
	var datas [][]byte
	inventories, ok := i.Get(InventoryCacheKey)
	if !ok {
		openlogging.GetLogger().Error("not data need to send apm")
		return errors.New("not data need to send apm")
	}
	for _, v := range inventories {
		temp, _ := json.Marshal(v)
		datas = append(datas, temp)
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
	return httpDo(i.httpClient, tAgentMessage, i.Url, i.ProjectID)
}

// Delete delete inventory cache by key , when key is empty will delete all cache of inventory
func (i *InventoryApm) Delete(key string) {
	if key == "" {
		i.inventoryCache.Flush()
		return
	}
	i.inventoryCache.Delete(key)
}

// Get get tKpiMessage from cache
func (k *InventoryApm) Get(key string) ([]common.Inventory, bool) {
	d, ok := k.inventoryCache.Get(key)
	if !ok {
		return []common.Inventory{}, false
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
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      utils.GetCertPool(caPath, ""),
					Certificates: []tls.Certificate{utils.GetCertificate(caPath, "", "")},
				},
			},
		},
	}
}

// NewInventory return new inventory
func NewInventory(hostname, ip, appName, clusterKey, serviceType,
	displayName, instanceName, containerID, appID string, pid int, Props map[string]interface{}) common.Inventory {

	return common.Inventory{
		Hostname:     hostname,
		IP:           ip,
		AppID:        appID,
		AppName:      appName,
		ClusterKey:   clusterKey,
		ServiceType:  serviceType,
		DisplayName:  displayName,
		InstanceName: instanceName,
		ContainerID:  containerID,
		Pid:          pid,
		Props:        Props,
		Created:      utils.GetTimeMillisecond(),
	}
}

func init() {
	t := time.NewTimer(5 * time.Second)
	go func() {
		for range t.C {
			inventoryCache := NewInventory(runtime.HostName, utils.GetLocalIP(), runtime.App, "", config.GlobalDefinition.Cse.Service.Registry.Type,
				"", runtime.InstanceID, pod.GetContainerID(), runtime.App, 0, nil)
			continue
			InventoryApmCache.Set(inventoryCache)
			InventoryApmCache.Send()
		}
	}()
}
