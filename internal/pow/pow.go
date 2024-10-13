package pow

import (
	"math/big"

	"github.com/rchelbii/tock/internal/block"
)

const TARGET_BITS = 24

type ProofOfWork struct {
	block  *block.Block
	target *big.Int
}

func NewProofOfWork(b *block.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-TARGET_BITS))
	pow := &ProofOfWork{
		block:  b,
		target: target,
	}
	return pow
}
