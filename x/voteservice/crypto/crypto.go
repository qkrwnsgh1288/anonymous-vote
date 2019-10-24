package crypto

import (
	"errors"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto/secp256k1"
	"math/big"
)

type Point struct {
	X *big.Int
	Y *big.Int
}
type JacobianPoint struct {
	X *big.Int
	Y *big.Int
	Z []byte
}

type ZkInfo struct {
	x  *big.Int // private key
	xG Point    // public key
	v  *big.Int // random nonce for zkp
	w  *big.Int // random nonce for 1outof2 zkp
	r  *big.Int // 1 or 2, random nonce for 1outof2 zkp
	d  *big.Int // 1 or 2, random nonce for 1outof2 zkp
}

// vG (blinding value), xG (public key), x (what we are proving)
// c = H(g, g^{v}, g^{x});
// r = v - xz (mod p);
// return(r,vG)
func CreateZKP(x, v *big.Int, xG Point) (res [4]*big.Int, err error) {
	curve := secp256k1.S256()

	var G Point
	G.X = curve.Gx
	G.Y = curve.Gy

	if !curve.IsOnCurve(xG.X, xG.Y) {
		return res, errors.New("error occured in CreateZKP: xG is not pubKey")
	}

	// Get g^{v}
	// Convert to Affine Co-ordinates
	var vG JacobianPoint
	vG.X, vG.Y = curve.ScalarBaseMult(v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	// Get c = H(g, g^{x}, g^{v});
	//hash := sha256.New()
	//tmpSender := common.GetBigInt("130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb", 16) // todo: have to change

	return res, nil
}
