package chain

import (
	"block"
	"transaction"
)

// Chain :链
type Chain struct {
	Blocks []*block.Block `json:"chain"`
	Height int            `json:"height"`
}

// NewChain :新建链
func NewChain() *Chain {
	ch := new(Chain)
	genenisBlock := block.NewGenesisBlock(transaction.NewCoinbaseTx("renne", ""))
	ch.Blocks = append(ch.Blocks, genenisBlock)
	ch.Height = 1
	return ch
}

// GetHeight :获取区块高度
func (ch *Chain) GetHeight() int {
	return ch.Height
}

//GetTailBlock 获取最新的区块
func (ch *Chain) GetTailBlock() *block.Block {
	return ch.Blocks[len(ch.Blocks)-1]
}

//AppendNewBlock :添加区块，会补全信息
func (ch *Chain) AppendNewBlock(Txs []*transaction.Tx) error {
	block := block.NewBlock(Txs, ch.GetTailBlock().Hash, ch.GetHeight())

	ch.Blocks = append(ch.Blocks, block)
	ch.Height++
	return nil
}
