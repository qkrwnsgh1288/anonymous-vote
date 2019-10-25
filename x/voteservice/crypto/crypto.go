package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto/secp256k1"
	"math/big"
)

var (
	curve *secp256k1.BitCurve
)

func init() {
	curve = secp256k1.S256()
}

type Point struct {
	X *big.Int
	Y *big.Int
}
type JacobianPoint struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
	//Z []byte
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
func CreateZKP(senderAddr string, x, v *big.Int, xG Point) (res [4]*big.Int, err error) {
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
	hash := sha256.New()
	sender := common.GetBigInt(senderAddr, 16) // todo: senderAddr check

	hashInput := sender.Bytes()
	hashInput = append(hashInput, curve.Gx.Bytes()...)
	hashInput = append(hashInput, curve.Gy.Bytes()...)
	hashInput = append(hashInput, xG.X.Bytes()...)
	hashInput = append(hashInput, xG.Y.Bytes()...)
	hashInput = append(hashInput, vG.X.Bytes()...)
	hashInput = append(hashInput, vG.Y.Bytes()...)
	//hashInput = append(hashInput, vG.Z...)
	hashInput = append(hashInput, vG.Z.Bytes()...)
	hash.Write(hashInput)

	md := hash.Sum(nil)
	hexStr := hex.EncodeToString(md)
	c := common.GetBigInt(hexStr, 16)

	// Get 'r' the zkp
	xc := mulMod(x, c, curve.N)

	// v - xc
	r := subMod(v, xc, curve.N)

	res[0] = r
	res[1] = vG.X
	res[2] = vG.Y
	//res[3] = new(big.Int).SetBytes(vG.Z)
	res[3] = vG.Z

	return res, nil
}

// Parameters xG, r where r = v - xc, and vG.
// Verify that vG = rG + xcG!
func VerifyZKP(senderAddr string, xG Point, r *big.Int, vG JacobianPoint) bool {
	var G Point
	G.X = curve.Gx
	G.Y = curve.Gy

	// Check both keys are on the curve.
	if !curve.IsOnCurve(xG.X, xG.Y) || !curve.IsOnCurve(vG.X, vG.Y) {
		return false
	}

	// Get c = H(g, g^{x}, g^{v});
	hash := sha256.New()
	sender := common.GetBigInt(senderAddr, 16) // todo: senderAddr check

	hashInput := sender.Bytes()
	hashInput = append(hashInput, curve.Gx.Bytes()...)
	hashInput = append(hashInput, curve.Gy.Bytes()...)
	hashInput = append(hashInput, xG.X.Bytes()...)
	hashInput = append(hashInput, xG.Y.Bytes()...)
	hashInput = append(hashInput, vG.X.Bytes()...)
	hashInput = append(hashInput, vG.Y.Bytes()...)
	//hashInput = append(hashInput, vG.Z...)
	hashInput = append(hashInput, vG.Z.Bytes()...)
	hash.Write(hashInput)

	md := hash.Sum(nil)
	hexStr := hex.EncodeToString(md)
	c := common.GetBigInt(hexStr, 16)

	// Get g^{r}, and g^{xc}
	var rG JacobianPoint
	rG.X, rG.Y = curve.ScalarBaseMult(r.Bytes())

	var xcG JacobianPoint
	xcG.X, xcG.Y = curve.ScalarMult(xG.X, xG.Y, c.Bytes())

	// Add both points together
	var rGxcG JacobianPoint
	rGxcG.X, rGxcG.Y = curve.Add(rG.X, rG.Y, xcG.X, xcG.Y)

	if rGxcG.X.Cmp(vG.X) == 0 && rGxcG.Y.Cmp(vG.Y) == 0 {
		return true
	} else {
		return false
	}
}
