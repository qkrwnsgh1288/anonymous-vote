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
	Z *big.Int
}

type ZkInfo struct {
	x  *big.Int // private key
	xG Point    // public key
	v  *big.Int // random nonce for zkp
	w  *big.Int // random nonce for 1outof2 zkp
	r  *big.Int // 1 or 2, random nonce for 1outof2 zkp
	d  *big.Int // 1 or 2, random nonce for 1outof2 zkp
}

//var curve EllipticCurve
//func init() {
//	/* See SEC2 pg.9 http://www.secg.org/collateral/sec2_final.pdf */
//	/* secp256k1 elliptic curve parameters */
//	curve.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
//	curve.A, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000000", 16)
//	curve.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
//	curve.G.X, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
//	curve.G.Y, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
//	curve.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
//	curve.H, _ = new(big.Int).SetString("01", 16)
//}

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
