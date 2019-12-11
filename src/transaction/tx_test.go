package transaction

import "testing"

import "encoding/json"

import "fmt"

//func TestTxHash(t *testing.T) {
//	var tx Tx
//	tx.Input = "fuckyou"
//	var hash = tx.TxHash()
//	fmt.Println(hash)
//}
//
//func TestNewTx(t *testing.T) {
//	var tx *Tx = NewTx("fuckyou")
//	j, _ := json.Marshal(tx)
//	fmt.Println(string(j))
//}
//

func TestNewCoinbase(t *testing.T) {
	tx := NewCoinbaseTx("renne", "")
	j, _ := json.Marshal(tx)
	fmt.Println(string(j))
}
