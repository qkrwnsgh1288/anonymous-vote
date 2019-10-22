package common

import "math/big"

func GetBigInt(str string, base int) *big.Int {
	res, _ := new(big.Int).SetString(str, base)
	return res
}
