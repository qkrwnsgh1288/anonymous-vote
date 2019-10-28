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
	X  *big.Int // private key
	xG Point    // public key
	V  *big.Int // random nonce for zkp
	W  *big.Int // random nonce for 1outof2 zkp
	R  *big.Int // 1 or 2, random nonce for 1outof2 zkp
	D  *big.Int // 1 or 2, random nonce for 1outof2 zkp
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
func Register(senderAddr string, xG Point, vG JacobianPoint, r *big.Int) error {
	// todo:  dead line check
	// todo: white list check
	Voters = append(Voters, MakeVoter(senderAddr, xG))
	//Voters[Totalregistered] = MakeVoter(senderAddr, xG)
	Totalregistered += 1

	return nil
}

// Calculate the reconstructed keys
func FinishRegistrationPhase() error {
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

// random 'W', 'r1', 'd1'
func Create1outof2ZKPYesVote(sender string, xG, yG Point, w, r1, d1, x *big.Int) (y, a1, b1, a2, b2 Point, res [4]*big.Int, err error) {
	// 1. y = h^{X} * g
	y.X, y.Y = Curve.ScalarMult(yG.X, yG.Y, x.Bytes())
	y.X, y.Y = Curve.Add(y.X, y.Y, Curve.Gx, Curve.Gy)

	// 2. a1 = g^{r1} * x^{d1}
	a1.X, a1.Y = Curve.ScalarBaseMult(r1.Bytes())
	tmp1 := MakeDefaultPoint()
	tmp1.X, tmp1.Y = Curve.ScalarMult(xG.X, xG.Y, d1.Bytes())
	a1.X, a1.Y = Curve.Add(a1.X, a1.Y, tmp1.X, tmp1.Y)

	// 3. b1 = h^{r1} * y^{d1} (temp = affine 'y')
	tmp1.X, tmp1.Y = Curve.ScalarMult(yG.X, yG.Y, r1.Bytes())
	// Setting temp to 'y'
	temp := MakeDefaultPoint()
	temp.X.SetBytes(y.X.Bytes())
	temp.Y.SetBytes(y.Y.Bytes())

	b1_tmpX, b1_tmpY := Curve.ScalarMult(temp.X, temp.Y, d1.Bytes())
	b1.X, b1.Y = Curve.Add(tmp1.X, tmp1.Y, b1_tmpX, b1_tmpY)

	// 4. a2 = g^{w}
	a2.X, a2.Y = Curve.ScalarBaseMult(w.Bytes())

	// 5. b2 = h^{w} (where h = g^{y})
	b2.X, b2.Y = Curve.ScalarMult(yG.X, yG.Y, w.Bytes())

	// Get c = H(id, xG, Y, a1, b1, a2, b2)
	// id is H(round, voter_index, voter_address)...
	hash := sha256.New()
	hInput := common.GetBigInt(sender, 16).Bytes()
	hInput = append(hInput, xG.X.Bytes()...)
	hInput = append(hInput, xG.Y.Bytes()...)
	hInput = append(hInput, y.X.Bytes()...)
	hInput = append(hInput, y.Y.Bytes()...)
	hInput = append(hInput, a1.X.Bytes()...)
	hInput = append(hInput, a1.Y.Bytes()...)
	hInput = append(hInput, b1.X.Bytes()...)
	hInput = append(hInput, b1.Y.Bytes()...)
	hInput = append(hInput, a2.X.Bytes()...)
	hInput = append(hInput, a2.Y.Bytes()...)
	hInput = append(hInput, b2.X.Bytes()...)
	hInput = append(hInput, b2.Y.Bytes()...)
	hash.Write(hInput)

	md := hash.Sum(nil)
	hexStr := hex.EncodeToString(md)
	c := common.GetBigInt(hexStr, 16)

	// d2 = c - d1
	d2 := subMod(c, d1, Curve.N)

	// r2 = w - (x * d2)
	r2 := subMod(w, mulMod(x, d2, Curve.N), Curve.N)

	res[0] = d1
	res[1] = d2
	res[2] = r1
	res[3] = r2

	return y, a1, b1, a2, b2, res, err
}

// random 'W', 'r2', 'd2'
func Create1outof2ZKPNoVote(sender string, xG, yG Point, w, r2, d2, x *big.Int) (y, a1, b1, a2, b2 Point, res [4]*big.Int, err error) {
	temp_affine1 := MakeDefaultPoint()
	temp_affine2 := MakeDefaultPoint()
	temp1 := MakeDefaultPoint()

	// 1. y = h^{x} * g
	y.X, y.Y = Curve.ScalarMult(yG.X, yG.Y, x.Bytes())

	// 2. a1 = g^{w}
	a1.X, a1.Y = Curve.ScalarBaseMult(w.Bytes())

	// 3. b1 = h^{w} (where h = g^{y})
	b1.X, b1.Y = Curve.ScalarMult(yG.X, yG.Y, w.Bytes())

	// 4. a2 = g^{r2} * x^{d2}
	a2.X, a2.Y = Curve.ScalarBaseMult(r2.Bytes())
	temp1.X, temp1.Y = Curve.ScalarMult(xG.X, xG.Y, d2.Bytes())
	a2.X, a2.Y = Curve.Add(a2.X, a2.Y, temp1.X, temp1.Y)

	// 5. b2
	// Negate the 'y' co-ordinate of G
	temp_affine1.X = Curve.Gx
	temp_affine1.Y.Sub(Curve.P, Curve.Gy)

	// We need the public key y in affine co-ordinates
	temp_affine2.X = y.X
	temp_affine2.Y = y.Y

	// We should end up with y^{d2} + g^{d2} .... (but we have the negation of g.. so y-g).
	tmpMul := MakeDefaultPoint()
	tmpMul1 := MakeDefaultPoint()
	tmpMul2 := MakeDefaultPoint()
	tmpMul2.X, tmpMul2.Y = Curve.ScalarMult(temp_affine2.X, temp_affine2.Y, d2.Bytes())
	tmpMul1.X, tmpMul1.Y = Curve.ScalarMult(temp_affine1.X, temp_affine1.Y, d2.Bytes())
	temp1.X, temp1.Y = Curve.Add(tmpMul2.X, tmpMul2.Y, tmpMul1.X, tmpMul1.Y)

	// Now... it is h^{r2} + temp2..
	tmpMul.X, tmpMul.Y = Curve.ScalarMult(yG.X, yG.Y, r2.Bytes())
	b2.X, b2.Y = Curve.Add(tmpMul.X, tmpMul.Y, temp1.X, temp1.Y)

	// Get c = H(i, xG, Y, a1, b1, a2, b2)
	hash := sha256.New()
	hInput := common.GetBigInt(sender, 16).Bytes()
	hInput = append(hInput, xG.X.Bytes()...)
	hInput = append(hInput, xG.Y.Bytes()...)
	hInput = append(hInput, y.X.Bytes()...)
	hInput = append(hInput, y.Y.Bytes()...)
	hInput = append(hInput, a1.X.Bytes()...)
	hInput = append(hInput, a1.Y.Bytes()...)
	hInput = append(hInput, b1.X.Bytes()...)
	hInput = append(hInput, b1.Y.Bytes()...)
	hInput = append(hInput, a2.X.Bytes()...)
	hInput = append(hInput, a2.Y.Bytes()...)
	hInput = append(hInput, b2.X.Bytes()...)
	hInput = append(hInput, b2.Y.Bytes()...)
	hash.Write(hInput)

	md := hash.Sum(nil)
	b_c := hex.EncodeToString(md)
	c := common.GetBigInt(b_c, 16)

	// d1 = c - d2
	d1 := subMod(c, d2, Curve.N)

	// r1 = w - (x * d1)
	r1 := subMod(w, mulMod(x, d1, Curve.N), Curve.N)

	res[0] = d1
	res[1] = d2
	res[2] = r1
	res[3] = r2

	return y, a1, b1, a2, b2, res, err
}
