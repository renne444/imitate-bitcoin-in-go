package blockchain

import (
	"encoding/json"
	"fmt"
	"testing"
)

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
	//coinbaseTx := transaction.NewCoinbaseTx("kitty", "")
	//tx := transaction.Tx{
	//	Vin: []transaction.TxInput{transaction.TxInput{
	//		InputTxID: "0x40e5013d6d26be3719f7e4325c179b60506d2e8f6f79c8e5405fe871afa89cd5",
	//		InputVout: 0,
	//		ScriptSig: "renne"}},
	//	Vout: []transaction.TxOutput{transaction.TxOutput{Value: 27, ScriptPubkey: "renne"}, transaction.TxOutput{Value: 23, ScriptPubkey: "kitty"}},
	//	Hash: "",
	//}

	//tx.Hash = tx.TxHash()

	//	bc := NewBlockchain()
	//	if err := bc.AppendNewBlock([]*transaction.Tx{coinbaseTx, &tx}); err != nil {
	//		t.Error("UTXO 区块插入错误")
	//	}
}

func TestUnspentTransactionsQuery(t *testing.T) {
	bc := NewBlockchain()
	//	bc.AppendNewBlock([]*transaction.Tx{transaction.NewCoinbaseTx("renne", "")})
	txg := bc.FindUnspentTransactions("renne")
	j, _ := json.Marshal(txg)
	fmt.Println(string(j))

	//	txg = bc.FindUnspentTransactions("kitty")
	//	j, _ = json.Marshal(txg)
	//	fmt.Println(string(j))
}
