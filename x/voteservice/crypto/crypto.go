package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto/secp256k1"
	"math/big"
)

const (
	Minimum_voter_count = 3
)

var (
	Curve           *secp256k1.BitCurve
	Voters          []Voter
	Totalregistered int
)

func init() {
	Curve = secp256k1.S256()
	Voters = make([]Voter, 0)
}

type Point struct {
	X *big.Int
	Y *big.Int
}

func MakeDefaultPoint() Point {
	return Point{
		X: new(big.Int),
		Y: new(big.Int),
	}
}

type JacobianPoint struct {
	X *big.Int
	Y *big.Int
	Z *big.Int
}

func MakeDefaultJacobianPoint() JacobianPoint {
	return JacobianPoint{
		X: new(big.Int),
		Y: new(big.Int),
		Z: new(big.Int),
	}
}

type ZkInfo struct {
	x  *big.Int // private key
	xG Point    // public key
	v  *big.Int // random nonce for zkp
	w  *big.Int // random nonce for 1outof2 zkp
	r  *big.Int // 1 or 2, random nonce for 1outof2 zkp
	d  *big.Int // 1 or 2, random nonce for 1outof2 zkp
}

type Voter struct {
	Addr             string
	RegisteredKey    Point // xG
	ReconstructedKey Point // yG
	Commitment       []byte
	Vote             [2]*big.Int
}

func MakeVoter(addr string, registerKey Point) Voter {
	return Voter{
		Addr:             addr,
		RegisteredKey:    registerKey,
		ReconstructedKey: MakeDefaultPoint(),
	}
}

// vG (blinding value), xG (public key), x (what we are proving)
// c = H(g, g^{v}, g^{x});
// r = v - xz (mod p);
// return(r,vG)
func CreateZKP(senderAddr string, x, v *big.Int, xG Point) (r *big.Int, vG JacobianPoint, err error) {
	var G Point
	G.X = Curve.Gx
	G.Y = Curve.Gy

	if !Curve.IsOnCurve(xG.X, xG.Y) {
		return r, vG, errors.New("error occured in CreateZKP: xG is not pubKey")
	}

	// Get g^{v}
	// Convert to Affine Co-ordinates
	vG.X, vG.Y = Curve.ScalarBaseMult(v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	// Get c = H(g, g^{x}, g^{v});
	hash := sha256.New()
	sender := common.GetBigInt(senderAddr, 16) // todo: senderAddr check

	hashInput := sender.Bytes()
	hashInput = append(hashInput, Curve.Gx.Bytes()...)
	hashInput = append(hashInput, Curve.Gy.Bytes()...)
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
	xc := mulMod(x, c, Curve.N)

	// v - xc
	r = subMod(v, xc, Curve.N)

	//res[0] = r
	//res[1] = vG.X
	//res[2] = vG.Y
	////res[3] = new(big.Int).SetBytes(vG.Z)
	//res[3] = vG.Z

	return r, vG, nil
}

// Parameters xG, r where r = v - xc, and vG.
// Verify that vG = rG + xcG!
func VerifyZKP(senderAddr string, xG Point, r *big.Int, vG JacobianPoint) bool {
	var G Point
	G.X = Curve.Gx
	G.Y = Curve.Gy

	// Check both keys are on the Curve.
	if !Curve.IsOnCurve(xG.X, xG.Y) || !Curve.IsOnCurve(vG.X, vG.Y) {
		return false
	}

	// Get c = H(g, g^{x}, g^{v});
	hash := sha256.New()
	sender := common.GetBigInt(senderAddr, 16) // todo: senderAddr check

	hashInput := sender.Bytes()
	hashInput = append(hashInput, Curve.Gx.Bytes()...)
	hashInput = append(hashInput, Curve.Gy.Bytes()...)
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
	rG.X, rG.Y = Curve.ScalarBaseMult(r.Bytes())

	var xcG JacobianPoint
	xcG.X, xcG.Y = Curve.ScalarMult(xG.X, xG.Y, c.Bytes())

	// Add both points together
	var rGxcG JacobianPoint
	rGxcG.X, rGxcG.Y = Curve.Add(rG.X, rG.Y, xcG.X, xcG.Y)

	if rGxcG.X.Cmp(vG.X) == 0 && rGxcG.Y.Cmp(vG.Y) == 0 {
		return true
	} else {
		return false
	}
}

// Called by participants to register their voting public key
// Participant mut be eligible, and can only register the first key sent key.
func Register(senderAddr string, xG Point, vG JacobianPoint, r *big.Int) {
	// todo:  dead line check
	// todo: white list check
	Voters = append(Voters, MakeVoter(senderAddr, xG))
	//Voters[Totalregistered] = MakeVoter(senderAddr, xG)
	Totalregistered += 1
}

// Calculate the reconstructed keys
func finishRegistrationPhase() error {
	if Totalregistered < 3 {
		return errors.New("total registered is smaller than minimum(3)")
	}
	temp := MakeDefaultPoint()
	yG := MakeDefaultPoint()
	beforei := MakeDefaultPoint()
	afteri := MakeDefaultPoint()

	// Step 1 is to compute the index 1 reconstructed key
	afteri.X.SetBytes(Voters[1].RegisteredKey.X.Bytes())
	afteri.Y.SetBytes(Voters[1].RegisteredKey.Y.Bytes())

	for i := 2; i < Totalregistered; i++ {
		afteri.X, afteri.Y = Curve.Add(afteri.X, afteri.Y, Voters[i].RegisteredKey.X, Voters[i].RegisteredKey.Y)
	}

	Voters[0].ReconstructedKey.X.SetBytes(afteri.X.Bytes())
	Voters[0].ReconstructedKey.Y.Sub(Curve.P, afteri.Y)

	// Step 2 is to add to beforei, and subtract from afteri.
	for i := 1; i < Totalregistered; i++ {
		if i == 1 {
			beforei.X.SetBytes(Voters[0].RegisteredKey.X.Bytes())
			beforei.Y.SetBytes(Voters[0].RegisteredKey.Y.Bytes())
		} else {
			beforei.X, beforei.Y = Curve.Add(beforei.X, beforei.Y, Voters[i-1].RegisteredKey.X, Voters[i-1].RegisteredKey.Y)
		}

		// If we have reached the end... just store beforei
		// Otherwise, we need to compute a key.
		// Counting from 0 to n-1...
		if i == (Totalregistered - 1) {
			Voters[i].ReconstructedKey.X.SetBytes(beforei.X.Bytes())
			Voters[i].ReconstructedKey.Y.SetBytes(beforei.Y.Bytes())
		} else {
			// Subtract 'i' from afteri
			temp.X.SetBytes(Voters[i].RegisteredKey.X.Bytes())
			temp.Y.Sub(Curve.P, Voters[i].RegisteredKey.Y)

			// Grab negation of afteri (did not seem to work with Jacob co-ordinates)
			afteri.X, afteri.Y = Curve.Add(afteri.X, afteri.Y, temp.X, temp.Y)

			temp.X.SetBytes(afteri.X.Bytes())
			temp.Y.Sub(Curve.P, afteri.Y)

			// Now we do beforei - afteri...
			yG.X, yG.Y = Curve.Add(beforei.X, beforei.Y, temp.X, temp.Y)

			Voters[i].ReconstructedKey.X = yG.X
			Voters[i].ReconstructedKey.Y = yG.Y
		}
	}

	return nil
}
