package pow

import (
	"block"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
	"strconv"
)

//Difficulty : 这里的难度指的是签到零的个数
const Difficulty = 24
const maxNonce = 0x7fffffff

//ProofOfWork :工作量证明
type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}

//CalcBlockNonce : public 计算区块的pow具体值，返回挖矿是否成功
func CalcBlockNonce(b *block.Block) bool {
	pow := NewProofOfWork(b)
	nonce, _ := pow.Run()

	if nonce > 0 {
		b.Nonce = nonce
		return true
	}
	return false

}

//NewProofOfWork : 新建工作量证明计算实体
func NewProofOfWork(b *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return &ProofOfWork{
		block:  b,
		target: target,
	}
}

//Run : 计算pow的nonce，也就是挖矿
func (p *ProofOfWork) Run() (int, string) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for {
		data := p.prepareData(nonce)
		hash = sha256.Sum256([]byte(data))
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(p.target) < 0 {
			return nonce, hex.EncodeToString(hashInt.Bytes())
		} else if nonce > maxNonce {
			return -1, "0"
		}

		nonce++
	}
}

//IsValid : 验证区块的nonce值是否合法
func (p *ProofOfWork) IsValid() bool {
	var hashInt big.Int

	data := p.prepareData(p.block.Nonce)
	hash := sha256.Sum256([]byte(data))
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(p.target) < 0
}

func (p *ProofOfWork) prepareData(nonce int) string {
	data := strconv.Itoa(p.block.Index) +
		p.block.PreHash +
		strconv.FormatInt(p.block.Timestamp, 16) +
		p.block.TxHash + strconv.Itoa(nonce)
	return data
}
