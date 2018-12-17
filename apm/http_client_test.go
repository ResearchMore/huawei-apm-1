package apm

import (
	"fmt"
	"testing"
)

func TestInventoryApm_Get(t *testing.T) {
	u := getURL("", "")
	fmt.Println(u)
}
