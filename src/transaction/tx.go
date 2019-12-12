package transaction

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

//TxInput 交易输入部分
type TxInput struct {
	InputTxID []byte `json:"ID"`
	InputVout int    `json:"index"`
	ScriptSig string `json:"sig"`
}

//CanUnlockOutputWith : 某个输入，能否解锁拥有某个公钥
func (txi *TxInput) CanUnlockOutputWith(scriptPub string) bool {
	return scriptPub == txi.ScriptSig
}

//TxOutput : 交易输出部分，其实就是UTXO统计的部分
type TxOutput struct {
	Value        int    `json:"value"`
	ScriptPubkey string `json:"pubkey"`
}

//CanBeUnlockedWith : 某笔输出能否被某个签名解锁
func (txo *TxOutput) CanBeUnlockedWith(scriptSig string) bool {
	return txo.ScriptPubkey == scriptSig
}

//Tx 交易的格式
type Tx struct {
	Vin  []TxInput  `json:"vin"`
	Vout []TxOutput `json:"vout"`
	ID   string     `json:"id"`
	Hash string     `json:"hash"`
}

const (
	coinbaseReward = 50
)

//NewCoinbaseTx : 新建创块交易
func NewCoinbaseTx(to, data string) *Tx {
	if data == "" {
		data = fmt.Sprintf("Reward to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{coinbaseReward, to}
	tx := Tx{[]TxInput{txin}, []TxOutput{txout}, "id_sample", ""}

	tx.Hash = tx.TxHash()
	return &tx
}

//NewTx 新建裸交易
//func NewTx(input string) *Tx {
//	tx := &Tx{
//		Input: input,
//	}
//	tx.Hash = tx.TxHash()
//	return tx
//}

//TxHash hash of tx
func (tx *Tx) TxHash() string {
	jsonInput, _ := json.Marshal(tx.Vin)
	jsonOutput, _ := json.Marshal(tx.Vout)

	data := bytes.Join([][]byte{jsonInput, jsonOutput, []byte(tx.ID)}, []byte{})

	hashInByte := sha256.Sum256(data)
	hashValue := "0x" + hex.EncodeToString(hashInByte[:])

	return hashValue
}
