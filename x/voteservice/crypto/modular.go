package crypto

import (
	"math/big"
)

/*** Modular Arithmetic ***/

/* NOTE: Returning a new z each time below is very space inefficient, but the
 * alternate accumulator based design makes the point arithmetic functions look
 * absolutely hideous. I may still change this in the future. */

// addMod computes z = (x + y) % p.
func addMod(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).Add(x, y)
	z.Mod(z, p)
	return z
}

// subMod computes z = (x - y) % p.
func subMod(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).Sub(x, y)
	z.Mod(z, p)
	return z
}

// mulMod computes z = (x * y) % p.
func mulMod(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	n := new(big.Int).Set(x)
	z = big.NewInt(0)

	for i := 0; i < y.BitLen(); i++ {
		if y.Bit(i) == 1 {
			z = addMod(z, n, p)
		}
		n = addMod(n, n, p)
	}

	return z
}

// invMod computes z = (1/x) % p.
func invMod(x *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).ModInverse(x, p)
	return z
}

// expMod computes z = (x^e) % p.
func expMod(x *big.Int, y *big.Int, p *big.Int) (z *big.Int) {
	z = new(big.Int).Exp(x, y, p)
	return z
}

// sqrtMod computes z = sqrt(x) % p.
func sqrtMod(x *big.Int, p *big.Int) (z *big.Int) {
	/* assert that p % 4 == 3 */
	if new(big.Int).Mod(p, big.NewInt(4)).Cmp(big.NewInt(3)) != 0 {
		panic("p is not equal to 3 mod 4!")
	}

	/* z = sqrt(x) % p = x^((p+1)/4) % p */

	/* e = (p+1)/4 */
	e := new(big.Int).Add(p, big.NewInt(1))
	e = e.Rsh(e, 2)

	z = expMod(x, e, p)
	return z
}

func doubleM(P *JacobianPoint) {
	p := Curve.P
	if P.Z.Cmp(big.NewInt(0)) == 0 {
		return
	}
	Px := new(big.Int).SetBytes(P.X.Bytes())
	Py := new(big.Int).SetBytes(P.Y.Bytes())
	Py2 := mulMod(Py, Py, p)
	s := mulMod(big.NewInt(4), mulMod(Px, Py2, p), p)
	m := mulMod(big.NewInt(3), mulMod(Px, Px, p), p)
	subTmp := new(big.Int).Sub(p, addMod(s, s, p))
	PxTemp := addMod(mulMod(m, m, p), subTmp, p)

	P.X = PxTemp
	subTmp2 := new(big.Int).Sub(p, PxTemp)
	subTmp3 := new(big.Int).Sub(p, mulMod(big.NewInt(8), mulMod(Py2, Py2, p), p))
	P.Y = addMod(mulMod(m, addMod(s, subTmp2, p), p), subTmp3, p)
	P.Z = mulMod(big.NewInt(2), mulMod(Py, P.Z, p), p)
}

func AddMixedM(P *JacobianPoint, Q Point) {
	if P.Y.Cmp(big.NewInt(0)) == 0 {
		P.X.SetBytes(Q.X.Bytes())
		P.Y.SetBytes(Q.Y.Bytes())
		P.Z = big.NewInt(1)
		return
	}
	if Q.Y.Cmp(big.NewInt(0)) == 0 {
		return
	}
	p := Curve.P
	zs := MakeDefaultPoint() // Pz^2, Pz^3, Qz^2, Qz^3
	zs.X = mulMod(P.Z, P.Z, p)
	zs.Y = mulMod(P.Z, zs.X, p)
	us := [4]*big.Int{
		P.X,                  // 3875796002492828325
		P.Y,                  // 9905580929764183882
		mulMod(Q.X, zs.X, p), // 9720566412674477352
		mulMod(Q.Y, zs.Y, p), // 1491572980042492800
	} // Pu, Ps, Qu, Qs
	if us[0].Cmp(us[2]) == 0 {
		if us[1].Cmp(us[3]) != 0 {
			P.X = big.NewInt(0)
			P.Y = big.NewInt(0)
			P.Z = big.NewInt(0)
			return
		} else {
			doubleM(P)
			return
		}
	}
	subTmp := new(big.Int).Sub(p, us[0])
	subTmp2 := new(big.Int).Sub(p, us[1])
	h := addMod(us[2], subTmp, p)
	r := addMod(us[3], subTmp2, p)
	h2 := mulMod(h, h, p)
	h3 := mulMod(h2, h, p)

	subTmp3 := new(big.Int).Sub(p, h3)
	subTmp4 := new(big.Int).Sub(p, mulMod(big.NewInt(2), mulMod(us[0], h2, p), p))
	Rx := addMod(mulMod(r, r, p), subTmp3, p)
	Rx = addMod(Rx, subTmp4, p)

	//P.X.SetBytes(Rx.Bytes())
	P.X = Rx
	subTmp5 := new(big.Int).Sub(p, Rx)
	pyTmp := mulMod(r, addMod(mulMod(us[0], h2, p), subTmp5, p), p)
	subTmp6 := new(big.Int).Sub(p, mulMod(us[1], h3, p))
	P.Y = addMod(pyTmp, subTmp6, p)
	P.Z = mulMod(h, P.Z, p)
}

func ToZ1(PJ *JacobianPoint, prime *big.Int) {
	zInv := invMod(PJ.Z, prime)
	zInv2 := mulMod(zInv, zInv, prime)
	PJ.X = mulMod(PJ.X, zInv2, prime)
	PJ.Y = mulMod(PJ.Y, mulMod(zInv, zInv2, prime), prime)
	PJ.Z = big.NewInt(1)
}
