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

}
