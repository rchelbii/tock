package pow

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"

	"github.com/rchelbii/tock/internal/block"
)

const (
	TARGET_BITS = 24

	MAX_NONCE = math.MaxInt64
)

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

func (pow *ProofOfWork) DataToHash(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.Data,
			pow.block.PrevBlockHash,
			[]byte(intToHex(int(pow.block.Timestamp))),
			[]byte(intToHex(int64(TARGET_BITS))),
			[]byte(intToHex(int64(nonce))),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Proof() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte

	nonce := 0

	for nonce < MAX_NONCE {
		data := pow.DataToHash(nonce)
		hash := sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.DataToHash(0)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
