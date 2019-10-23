package secp256k1

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"math/big"
	"testing"
)

type Point struct {
	X *big.Int
	Y *big.Int
}

//func TestBitCurve_IsOnCurve(t *testing.T) {
//	x := common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483942", 10)
//	y := common.GetBigInt("98141067444202828032016841245494455215374046124323329249557735915756843740538", 10)
//
//
//}

func TestDouble(t *testing.T) {
	inputA := common.GetBigInt("1", 10)
	inputB := common.GetBigInt("2", 10)
	inputC := common.GetBigInt("3", 10)
	a, b, c := theCurve.doubleJacobian(inputA, inputB, inputC)
	fmt.Println(a, b, c)
}

func TestBitCurve_ScalarBaseMult(t *testing.T) {
	v := common.GetBigInt("46174680605738213156470093129897818116924733100966263874097524943944597791118", 10)
	a, b := theCurve.ScalarBaseMult(v.Bytes())
	fmt.Println(a, b)

	//addr := "0x130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb"
	//sha256.New()
}
func TestHash(t *testing.T) {
	data := "aaa"
	Gx := common.GetBigInt("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	Gy := common.GetBigInt("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
	xG := Point{
		X: common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483942", 10),
		Y: common.GetBigInt("98141067444202828032016841245494455215374046124323329249557735915756843740538", 10),
	}
	hash := sha256.New()

	dd := []byte(data)
	dd = append(dd, Gx.Bytes()...)
	dd = append(dd, Gy.Bytes()...)
	dd = append(dd, xG.X.Bytes()...)
	dd = append(dd, xG.Y.Bytes()...)

	hash.Write(dd)

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)
}
func TestHash2(t *testing.T) {
	addr := common.GetBigInt("130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb", 16)

	hash := sha256.New()
	dd := addr.Bytes()
	hash.Write(dd)

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)
}
