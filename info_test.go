package a2s

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	TestHost = "s1.zhenyangli.me"
)

func TestInfo(t *testing.T) {
	c, err := NewClient(TestHost)

	defer c.Close()

	if err != nil {
		t.Error(err)
		return
	}

	i, err := c.QueryInfo()

	if err != nil {
		t.Error(err)
		return
	}

	JSON, _ := json.Marshal(i)

	fmt.Println(string(JSON))
}
