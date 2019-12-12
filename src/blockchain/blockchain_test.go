package blockchain

import "testing"

import "fmt"

func TestAppend(t *testing.T) {
	bc := NewBlockchain()

	fmt.Println(string(bc.tip))
	//	bc.AppendNewBlock(nil)
	//	fmt.Println(string(bc.tip))
	//	bc.AppendNewBlock(nil)
	//	fmt.Println(string(bc.tip))
	//	bc.AppendNewBlock(nil)
	//	fmt.Println(string(bc.tip))

}

func TestAppendUTXO(t *testing.T) {

}
