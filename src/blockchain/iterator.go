package blockchain

import "github.com/boltdb/bolt"

import "block"

//Iterator 一个向前迭代器，链式查询
type Iterator struct {
	currentHash []byte
	db          *bolt.DB
}

//NewIterator : Blockchain成员函数作用为新建迭代器
func (bc *Blockchain) NewIterator() *Iterator {
	return &Iterator{
		currentHash: bc.tip,
		db:          bc.db,
	}
}

//Next ：迭代器成员函数，查找当前指向块的前一块
func (it *Iterator) Next() *block.Block {
	var b *block.Block

	if string(it.currentHash) == "" {
		return nil
	}

	err := it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		encodeBlock := bucket.Get(it.currentHash)
		b, _ = block.DeserializeBlock(encodeBlock)

		return nil
	})

	if err == nil {
		it.currentHash = []byte(b.PreHash)
		return b
	}
	return nil
}
