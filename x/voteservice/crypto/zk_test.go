package crypto

import (
	"fmt"
	"math/big"
	"testing"
)

type Voter struct {
	Addr             string
	RegisteredKey    [2]*big.Int
	ReconstructedKey [2]*big.Int
	Commitment       []byte
	Vote             [2]*big.Int
}

func TestBasic(t *testing.T) {

}

func TestVote(t *testing.T) {

	fmt.Println(vote1ZK.x.String(), vote2ZK, vote3ZK)
}
