package common

type spanType int64

// TKpiMessage struct for collector
// url https://elbIp:8923/{project_id}/kpi/istio
type TKpiMessage struct {
	// sent endpoint name , if you can know pod name you can set it for you
	// podName,else you can set default name "unknownClient"
	SourceResourceId string `json:"sourceResouceId"`
	// receiver endpoint name,use pod name ,default name use "unknownDestination"
	DestResourceId string `json:"destResourceId"`
	//protocol type use request.path , type of protocol
	// e.g.  request.path|rpc
	TransactionType string `json:"transactionType"`
	// application id , collector use `MD5("istio"|projectid|clusterId|namespace|applicationName)`
	// in this application we use default replace  projectId
	// clusterId is you cluster id. if you did't has cluster,will use "default" to replace it
	// namespace use "cse" for this value
	// applicationName use server name . this name is  what server name set in micro_server.yaml
	// e.g. MD5("istio"|"default"|"default"|"cse"|"trace_plugin")
	AppId string `json:"appId"`
	// you latency error this value is non-essential value
	SelfErrorLatency byte `json:"selfErrorLatency"`
	// request count not need to sent.collector will calculation this value about you value
	Throughput int32 `json:"throughput"`
	// non-essential value
	SelfLatency byte `json:"selfLatency"`
	// non-essential value
	SelfActiveLatency byte `json:"selfActiveLatency"`
	// when return http status is less than 400 value will set response.duration
	// if status greater or equal to 400 , not need set this value
	TotalLatency []byte `json:"totalLatency"`
	// non-essential value
	TotalActiveLatency byte `json:"totalActiveLatency"`
	// enum value , like 1
	SpanType spanType
	// non-essential value
	TotalLatencyList []int32 `json:"totalLatencyList"`
	// non-essential value
	TotalErrorIndicatorList []bool `json:"totalErrorIndicatorList"`
	// when http status greater or equal to 400 set  response.duration for this
	TotalErrorLatency []byte `json:"totalErrorLatency"`
	// non-essential value
	NamespaceName string `json:"namespaceName"`
	// source name , source  version or use " client"
	// source name use call client name,version is client version
	// name didn't set default value, version default "0.0.1"
	SrcTierName string `json:"srcTierName"`
	// destination.service name and version, collector default value "unknownDest"
	// when you know service name andserver version
	DestTierName string `json:"destTierName"`
}

// TAgentMessage struct for collector
type TAgentMessage struct {
	// Uniquely identifies
	AgentContext string `json:"agentContext"`
	// non-essential value
	TenantName string `json:"tenantName"`
	// messages map key : "istio"|projectid|clusterId|namespace|applicationName
	// in this application we use default replace  projectId
	// clusterId is you cluster id. if you did't has cluster,will use "default" to replace it
	// namespace use "cse" for this value
	// applicationName use server name . this name is  what server name set in micro_server.yaml
	// child map key : timestamp , if kpi between 10:00 and 10:01,please use timestamp of 10:00
	// child map value :it is report data byte e.g. TKpiMessage
	Messages map[string]map[int64][][]byte `json:"messages"`
}

// Inventory struct for APM
// url https://elbIp:8923/{project_id}/inventory/istio
type Inventory struct {
	// Hostname setting you host name
	Hostname    string `json:"hostname"`
	IP          string `json:"ip"`
	AgentID     string `json:"agentId"`
	AppName     string `json:"appName"`
	ClusterKey  string `json:"clusterKey"`  // ?
	ServiceType string `json:"serviceType"` // ?
	DisplayName string `json:"displayName"` // ?
	// instance name use server name or find this value for sc
	InstanceName string `json:"instanceName"`
	// container id when you use container run you server set this
	ContainerID string `json:"containerId"`
	Pid         int    `json:"pid"`
	AppID       string `json:"appId"`
	// Props
	Props map[string]interface{} `json:"props"`
	// non-essential value
	Ports interface{} `json:"ports"`
	// non-essential value
	IPs           interface{} `json:"ips"`
	Tier          string      `json:"tier"`
	NamespaceName string      `json:"namespaceName"`
	Created       int64       `json:"created"`
	Updated       int64       `json:"updated"`
	Deleted       int64       `json:"deleted"`
}

// KPICollectorMessage
type KPICollectorMessage struct {
	SourceResourceId   string
	DestResourceId     string
	TransactionType    string
	AppId              string
	SrcTierName        string
	DestTierName       string
	TotalErrorLatencys []int64
	TotalErrorLatency  int64
	TotalLatencys      []int64
	TotalLatency       int64
	SpanType           spanType
}
