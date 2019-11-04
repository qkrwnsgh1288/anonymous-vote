package secp256k1

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

type Point struct {
	X *big.Int
	Y *big.Int
}

func TestAdd(t *testing.T) {
	x1 := common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483942", 10)
	x2 := common.GetBigInt("98141067444202828032016841245494455215374046124323329249557735915756843740538", 10)
	x3 := common.GetBigInt("1", 10)
	y1 := common.GetBigInt("106453131882900883561540729696424913020938673149822726580895600813441888567406", 10)
	y2 := common.GetBigInt("51103279871057056523744718969849587301335546334788824374456705394361157035715", 10)
	y3 := common.GetBigInt("1", 10)
	a, b, c := theCurve.addJacobian(x1, x2, x3, y1, y2, y3)
	fmt.Println(a, b, c)

	res1, res2 := theCurve.AffineFromJacobian(a, b, c)

	assert.Equal(t, "19726021177552888194148621436129232937104234324513758427865268224158101547130", res1.String())
	assert.Equal(t, "50952383343742199881927221996840986713139267241507858986150651430342248986684", res2.String())
}
func TestAdd2(t *testing.T) {
	x := big.NewInt(3)
	y := big.NewInt(5)
	z := big.NewInt(1)
	x1 := big.NewInt(3)
	y1 := big.NewInt(5)
	z1 := big.NewInt(1)

	a, b, c := theCurve.addJacobian(x, y, z, x1, y1, z1)
	fmt.Println(a, b, c)
}

func TestDouble(t *testing.T) {
	inputA := big.NewInt(1) //common.GetBigInt("1", 10)
	inputB := big.NewInt(2) //common.GetBigInt("2", 10)
	inputC := big.NewInt(3) //common.GetBigInt("3", 10)
	a, b, c := theCurve.doubleJacobian(inputA, inputB, inputC)
	fmt.Println(a, b, c)
}

func TestBitCurve_ScalarBaseMult(t *testing.T) {
	v := common.GetBigInt("46174680605738213156470093129897818116924733100966263874097524943944597791118", 10)
	a, b := theCurve.ScalarBaseMult(v.Bytes())
	fmt.Println(a, b)
}
func TestBitCurve_ScalarMult(t *testing.T) {
	x := common.GetBigInt("9913299858144681957286823974289411938574605225645739615654527694124463202819", 10)
	yG := Point{
		X: common.GetBigInt("13640588435166186727072570872841920017273057013114604476956539355021275854144", 10),
		Y: common.GetBigInt("90715709871810868701227023413915222907311739236101232519930156567199700809709", 10),
	}
	a, b := theCurve.ScalarMult(yG.X, yG.Y, x.Bytes())
	fmt.Println(a, b)
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
