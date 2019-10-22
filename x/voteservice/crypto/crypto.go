package crypto

import (
	"errors"
	"math/big"
)

/// @dev Modular inverse of a (mod p) using euclid.
/// "a" and "p" must be co-prime.
/// @param a The number.
/// @param p The mmodulus.
/// @return x such that ax = 1 (mod p)
// does not use
func Invmod(a, p uint) (uint, error) {
	if a == 0 || a == p || p == 0 {
		return 0, errors.New("error occured in Invmod, (a==0 or p==0 or a==p)")
	}
	if a > p {
		a = a % p
	}
	var t1 int
	var t2 int = 1

	var r1 uint = p
	var r2 uint = a

	var q uint
	for r2 != 0 {
		q = r1 / r2
		t1, t2, r1, r2 = t2, t1-int(q)*t2, r2, r1-q*r2
	}
	if t1 < 0 {
		return p - uint(-t1), nil
	}
	return uint(t1), nil
}

/// @dev Modular exponentiation, b^e % m
/// Basically the same as can be found here:
/// https://github.com/mgenware/go-modular-exponentiation
/// @param b The base.
/// @param e The exponent.
/// @param m The modulus.
/// @return x such that x = b^e (mod m)
// does not use
func Expmod(b, e, m uint) (uint, error) {
	if b == 0 {
		return 0, nil
	}
	if e == 0 {
		return 1, nil
	}
	if m == 1 {
		return 0, nil
	} else if m == 0 {
		return 0, errors.New("error occured in Expmod. (m == 0)")
	}

	result, err := Expmod(b, e/2, m)
	if err != nil {
		return 0, err
	}
	result = (result * result) % m
	if e&1 != 0 {
		return ((b % m) * result) % m, nil
	}
	return result % m, nil
}

/// @dev Converts a point (Px, Py, Pz) expressed in Jacobian coordinates to (Px", Py", 1).
/// Mutates P.
/// @param P The point.
/// @param zInv The modular inverse of "Pz".
/// @param z2Inv The square of zInv
/// @param prime The prime modulus.
/// @return (Px", Py", 1)
func ToZ1(P [3]*big.Int, zInv, z2Inv, prime *big.Int) {
	P[0] = mulMod(P[0], z2Inv, prime)
	P[1] = mulMod(P[1], mulMod(zInv, z2Inv, prime), prime)
	P[2] = big.NewInt(1)
}

/// @dev See _toZ1(uint[3], uint, uint).
/// Warning: Computes a modular inverse.
/// @param PJ The point.
/// @param prime The prime modulus.
/// @return (Px", Py", 1)
func ToZ1_2(PJ [3]*big.Int, prime *big.Int) {
	zInv := invMod(PJ[2], prime)
	zInv2 := mulMod(zInv, zInv, prime)
	PJ[0] = mulMod(PJ[0], zInv2, prime)
	PJ[1] = mulMod(PJ[1], mulMod(zInv, zInv2, prime), prime)
	PJ[2] = big.NewInt(1)
}
