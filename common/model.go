package common

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
}
