package blockchain

import (
	"block"
	"log"
	"pow"
	"strconv"
	"transaction"

	"github.com/boltdb/bolt"
)

const (
	dbFile      = "my.db"
	blockBucket = "block"
)

//Blockchain : 可持久化的区块链类
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

//NewBlockchain : 新建区块链，读数据库文件/新建一条链
func NewBlockchain() *Blockchain {
	db, err := bolt.Open(dbFile, 0600, nil)
	var tip []byte

	if err != nil {
		log.Fatal("blockchain file open error")
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			//新建创币交易
			genesisBlock := block.NewGenesisBlock(transaction.NewCoinbaseTx("renne", ""))
			b, _ := tx.CreateBucket([]byte(blockBucket))
			genesisBlockSerialize, _ := genesisBlock.Serialize()
			err = b.Put([]byte(genesisBlock.Hash), genesisBlockSerialize)
			err = b.Put([]byte("l"), []byte(genesisBlock.Hash))
			err = b.Put([]byte("h"), []byte(strconv.Itoa(1)))
			tip = []byte(genesisBlock.Hash)
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	return &Blockchain{
		tip: tip,
		db:  db,
	}
}

//AppendNewBlock : 增加新区块
func (bc *Blockchain) AppendNewBlock(Txs []*transaction.Tx) error {
	var lastHash []byte
	var height int

	err := bc.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		lastHash = bucket.Get([]byte("l"))
		height, _ = strconv.Atoi(string(bucket.Get([]byte("h"))))

		return nil
	})

	if err != nil {
		log.Println("错误点1")
		return err
	}

	b := block.NewBlock(Txs, string(lastHash), height)
	pow.CalcBlockNonce(b)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		blockSerialize, _ := b.Serialize()

		bucket.Put([]byte(b.Hash), blockSerialize)
		bucket.Put([]byte("l"), []byte(b.Hash))
		bucket.Put([]byte("h"), []byte(strconv.Itoa(height+1)))

		bc.tip = []byte(b.Hash)
		return nil
	})

	if err != nil {
		log.Println("错误点2")
	}
	return err
}

//FindUnspentTransactions : 遍历整条链，将与某交易相关的所有未花费交易都找出来
func (bc *Blockchain) FindUnspentTransactions(address string) []*transaction.Tx {
	bcit := bc.NewIterator()

	unspentTransaction := []*transaction.Tx{}
	spentTransactionOutput := make(map[string][]int)

	for block := bcit.Next(); block != nil; block = bcit.Next() {
		for _, tx := range block.Tx {

			hasOutputUnspent := false
			for outputIndex, txOutput := range tx.Vout {
				isOutputSpent := false
				if txOutput.CanBeUnlockedWith(address) != true {
					continue
				}
				for _, spentIndex := range spentTransactionOutput[tx.Hash] {
					if outputIndex == spentIndex {
						isOutputSpent = true
						break
					}
				}
				if isOutputSpent == false {
					hasOutputUnspent = true
				}
			}
			if hasOutputUnspent == true {
				unspentTransaction = append(unspentTransaction, tx)
			}

			for _, txInput := range tx.Vin {
				if txInput.CanUnlockOutputWith(address) {
					spentTransactionOutput[txInput.InputTxID] = append(spentTransactionOutput[txInput.InputTxID], txInput.InputVout)
				}
			}

		}
	}
	return unspentTransaction
}
