package block

import (
	"encoding/json"
	"fmt"
	"testing"
	"transaction"
)

func TestNewGenesisBlock(t *testing.T) {
	fmt.Println("Test TestNewGenesisBlock START")
	block := NewGenesisBlock()
	byteBlock, _ := json.Marshal(block)
	strBlock := string(byteBlock)
	fmt.Println(strBlock)
	fmt.Println("Test TestNewGenesisBlock END")
}

func TestNewBlock(t *testing.T) {
	fmt.Println("Test TestNewBlock ")
	var txgp []*transaction.Tx

	txgp = append(txgp, transaction.NewTx("fuckyou"))
	block := NewBlock(txgp, "0x0", 1)

	j, _ := json.Marshal(block)
	fmt.Println(string(j))
	fmt.Println("Test TestNewBlock END")
}

func serializeSimulate(t *testing.T) []byte {
	var txgp []*transaction.Tx
	txgp = append(txgp, transaction.NewTx("fuckyou"))
	block := NewBlock(txgp, "0x0", 1)

	by, err := block.Serialize()
	if err != nil {
		t.Error("error on serialize " + err.Error())
	}
	return by
}

func TestSerialize(t *testing.T) {
	serializeSimulate(t)
}

func TestDeserialize(t *testing.T) {

	decodeByte := serializeSimulate(t)

	b, err := DeserializeBlock(decodeByte)
	if err != nil {
		t.Error("error on deserializing " + err.Error())
	}
	j, _ := json.Marshal(b)
	fmt.Println(string(j))
}
