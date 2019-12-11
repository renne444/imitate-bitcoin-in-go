package blockchain

import "testing"

import "fmt"

import "encoding/json"

import "block"

func TestIteratorNext(t *testing.T) {
	bc := NewBlockchain()
	it := bc.NewIterator()
	var b *block.Block = it.Next()
	for ; b != nil; b = it.Next() {
		j, _ := json.Marshal(b)
		fmt.Println(string(j))
	}
}
