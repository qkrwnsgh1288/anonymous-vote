package crypto

import (
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	a := 1
	b := 2
	var c, d int

	a, b, c, d = b, a+3, a, b
	fmt.Println(a, b, c, d)
}

func TestInvmod(t *testing.T) {
	a, err := Invmod(3, 17)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	a, err = Invmod(17, 3)
	fmt.Println(a)

	a, err = Invmod(3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func TestExpmod(t *testing.T) {
	res, err := Expmod(81792, 73363, 233)
	if err != nil {
		fmt.Println(err)
	}
	var exp uint = 161
	assert.Equal(t, exp, res)

	res, err = Expmod(1000, 1000, 19)
	if err != nil {
		fmt.Println(err)
	}
	exp = 7
	assert.Equal(t, exp, res)

	res, err = Expmod(12, 9, 1)
	if err != nil {
		fmt.Println(err)
	}
	exp = 0
	assert.Equal(t, exp, res)

	res, err = Expmod(111, 123, 53)
	if err != nil {
		fmt.Println(err)
	}
	exp = 35
	assert.Equal(t, exp, res)
}

func TestCreateZKPValidcheck1(t *testing.T) {
	vote1ZK := ZkInfo{
		x: common.GetBigInt("9913299858144681957286823974289411938574605225645739615654527694124463202819", 10),
		xG: Point{
			X: common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483942", 10),
			Y: common.GetBigInt("98141067444202828032016841245494455215374046124323329249557735915756843740538", 10),
		},
		v: common.GetBigInt("46174680605738213156470093129897818116924733100966263874097524943944597791118", 10),
		w: common.GetBigInt("38363520556597256917446345152267010746310705659165182691192706661717283647109", 10),
		r: common.GetBigInt("50335626772706697871408471165498611599437960211637817517623137749599062304789", 10),
		d: common.GetBigInt("63472464783107388493770567796117006062886472127577491241883308220111272611979", 10),
	}
	_, err := CreateZKP(vote1ZK.x, vote1ZK.v, vote1ZK.xG)
	assert.Nil(t, err)

	vote1ZK.xG.X = common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483943", 10)
	_, err = CreateZKP(vote1ZK.x, vote1ZK.v, vote1ZK.xG)
	assert.Equal(t, "error occured in CreateZKP: xG is not pubKey", err.Error())

}
