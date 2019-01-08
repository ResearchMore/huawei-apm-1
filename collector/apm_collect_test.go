package apm_collector

import (
	"fmt"
	"testing"
)

func TestCreateCollect(t *testing.T) {
	CreateCollect("", "", "")
	fmt.Println(Collector.Apm[Kpi_Collector_Key])
	fmt.Println(Collector.Apm[Kpi_Collector_Key] == nil)
	kpi := Collector.Apm[Kpi_Collector_Key]
	kpi.Delete("")
}
