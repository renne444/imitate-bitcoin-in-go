package main

import "chain"

import "transaction"

import "encoding/json"

import "fmt"

func main() {
	bc := chain.NewChain()

	tx1 := transaction.NewTx("Send coin from u1 to u2")
	tx2 := transaction.NewTx("fucking coin from u3 to u8")

	var btx1 []*transaction.Tx
	var btx2 []*transaction.Tx
	btx1 = append(btx1, tx1)
	btx2 = append(btx2, tx2)

	bc.AppendNewBlock(btx1)
	bc.AppendNewBlock(btx2)

	if j, err := json.Marshal(bc); err == nil {
		fmt.Println(string(j))
	}
}
