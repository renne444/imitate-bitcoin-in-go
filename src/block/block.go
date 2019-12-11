package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"strconv"
	"transaction"
	"util"
)

//Block 区块格式
type Block struct {
	PreHash   string            `json:"pre_hash"`
	Tx        []*transaction.Tx `json:"transactions"`
	Index     int               `json:"index"`
	Hash      string            `json:"hash"`
	Timestamp int64             `json:"timestamp"`
	TxHash    string            `json:"tx_hash"`
	Nonce     int               `json:"nonce"`
}

//NewGenesisBlock 根据配置生成新区块
func NewGenesisBlock() *Block {
	return NewBlock(nil, "0x0", 0)
}

//NewBlock 新增普通区块
func NewBlock(inputTx []*transaction.Tx, preHash string, index int) *Block {
	var block = &Block{
		PreHash:   preHash,
		Tx:        inputTx,
		Index:     index,
		Timestamp: util.GetTimestampNow(),
	}
	block.TxHash = block.calculateTxHash()
	block.Hash = block.calculateBlockHash()

	return block
}

//GetHash 获取区块哈希值
func (b *Block) GetHash() string {
	return b.Hash
}

//Serialize : 序列化区块对象
func (b *Block) Serialize() ([]byte, error) {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	if err := encoder.Encode(b); err != nil {
		return []byte{}, err
	}
	return result.Bytes(), nil
}

//DeserializeBlock : 反序列化
func DeserializeBlock(d []byte) (*Block, error) {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(&block); err != nil {
		return &Block{}, err
	}
	return &block, nil
}

func (b *Block) calculateTxHash() string {
	record := ""
	for _, tx := range b.Tx {
		record += tx.Hash
	}
	h := sha256.New()
	h.Write([]byte(record))
	return "0x" + hex.EncodeToString(h.Sum(nil))
}

func (b *Block) calculateBlockHash() string {
	record := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10) + b.PreHash
	for _, tx := range b.Tx {
		record += tx.Hash
	}
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return "0x" + hex.EncodeToString(hashed)
}
