package chain

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewChain(t *testing.T) {
	chain := NewChain()
	j, _ := json.Marshal(chain)
	fmt.Println(string(j))
}

func TestAppendNewBlock(t *testing.T) {
	chain := NewChain()

	err1 := chain.AppendNewBlock(nil)
	err2 := chain.AppendNewBlock(nil)
	err3 := chain.AppendNewBlock(nil)
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
		fmt.Println(err3)
		t.Error("连续插入错误")
	}
	j, _ := json.Marshal(chain)
	fmt.Println(string(j))
}
