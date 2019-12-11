package pow

import (
	"block"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {

	b := block.NewBlock(nil, "0x0", 1)
	pof := NewProofOfWork(b)
	fmt.Printf("    target = %s\n", hex.EncodeToString(pof.target.Bytes()))

	nonce, hashCalc := pof.Run()

	fmt.Printf("nonce = %d ,hash_calc = %s\n", nonce, hashCalc)
}

func TestMine(t *testing.T) {
	b := block.NewBlock(nil, "0x0", 1)
	if CalcBlockNonce(b) == false {
		t.Error("mining error")
	} else {
		fmt.Printf("挖矿成功，对应nonce = %d\n", b.Nonce)
	}
	pow := NewProofOfWork(b)
	if pow.IsValid() == false {
		t.Error("判断1: 验证正确性错误，对->错")
	}
}

func TestVaildFalse(t *testing.T) {
	b := block.NewBlock(nil, "0x0", 1)
	b.Nonce = 10
	pow := NewProofOfWork(b)

	if pow.IsValid() == true {
		t.Error("判断2: 验证正确性错误，错->对")
	}
}
