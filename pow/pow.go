package pow

import (
	"math/big"

	"github.com/sipovalovdmitro/building-blockchain-in-go/blockchain"
)

const targetBits = 24

type ProofOfWork struct {
	block  *blockchain.Block
	target *big.Int
}
