package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto/secp256k1"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var vote1ZK, vote2ZK, vote3ZK ZkInfo

func init() {
	vote1ZK = ZkInfo{
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
	vote2ZK = ZkInfo{
		x: common.GetBigInt("73044129382900516458626751513450444633224877614886040552580274724707341882358", 10),
		xG: Point{
			X: common.GetBigInt("106453131882900883561540729696424913020938673149822726580895600813441888567406", 10),
			Y: common.GetBigInt("51103279871057056523744718969849587301335546334788824374456705394361157035715", 10),
		},
		v: common.GetBigInt("5209700577050836730122816381945828534280019944306124503689657137675013206313", 10),
		w: common.GetBigInt("3281651291674397017871631723438190271143959716071949645966259657936788081884", 10),
		r: common.GetBigInt("41403887247771194901357327673253115844662353380189037573280322289018911955215", 10),
		d: common.GetBigInt("54881002424480711715545502563057680527631329702287064744481231671320736961772", 10),
	}
	vote3ZK = ZkInfo{
		x: common.GetBigInt("109643633626514401630001551396577794344562341547838637839149212543909734236096", 10),
		xG: Point{
			X: common.GetBigInt("107956135215754977339644472077254825401575884648279129012018898429310504004233", 10),
			Y: common.GetBigInt("113679158974756670989576148654313567926994200253163665614193081831818003969237", 10),
		},
		v: common.GetBigInt("90296205232189910611570761372692972689976252523802034275699368039112551113416", 10),
		w: common.GetBigInt("67234808599408419035045387287500787848801498653651572042289356038771497986569", 10),
		r: common.GetBigInt("63027092517635873569811565836959959067401728725379842671725115301137381260942", 10),
		d: common.GetBigInt("92351701889870080263340561154237384290634945237799585777221631564623413795918", 10),
	}
}

func TestCalc(t *testing.T) {
	a := 1
	b := 2
	var c, d int

	a, b, c, d = b, a+3, a, b
	assert.Equal(t, 2, a)
	assert.Equal(t, 4, b)
	assert.Equal(t, 1, c)
	assert.Equal(t, 2, d)
}

func TestIsOnCurve(t *testing.T) {
	curve := secp256k1.S256()
	res1 := curve.IsOnCurve(vote1ZK.xG.X, vote1ZK.xG.Y)
	assert.True(t, res1)

	vote1ZK.xG.X = common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483943", 10)
	res2 := curve.IsOnCurve(vote1ZK.xG.X, vote1ZK.xG.Y)
	assert.False(t, res2)
}

func TestVG(t *testing.T) {
	curve := secp256k1.S256()

	var vG JacobianPoint
	vG.X, vG.Y = curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	assert.Equal(t, "37002400596499253347436146477359872208984972423528869527866238051389129979940", vG.X.String())
	assert.Equal(t, "46104438919360535329359949165853481514194123783534889415421577162302988165861", vG.Y.String())
	assert.Equal(t, byte(0x01), vG.Z[31])

	hashZ := sha256.New()
	hashInputZ := vG.Z
	hashZ.Write(hashInputZ)
	assert.Equal(t, "ec4916dd28fc4c10d78e287ca5d9cc51ee1ae73cbfde08c6b37324cbfaac8bc5", hex.EncodeToString(hashZ.Sum(nil)))

	hash := sha256.New()
	hashInput := vG.X.Bytes()
	hashInput = append(hashInput, vG.Y.Bytes()...)
	hashInput = append(hashInput, vG.Z...)

	hash.Write(hashInput)

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	assert.Equal(t, "3671c70fe36d399d158f71abd58f782cba4ca924d073d3c27630ac1eb050fa7a", mdStr)
}

func TestLittleEndian(t *testing.T) {
	aaa := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	fmt.Println(aaa)
	bbb := new(big.Int).SetBytes(aaa)
	bbbBytes := bbb.Bytes()
	for i := 0; i < len(bbbBytes)/2; i++ {
		bbbBytes[i], bbbBytes[len(bbbBytes)-i-1] = bbbBytes[len(bbbBytes)-i-1], bbbBytes[i]
	}
	fmt.Println(bbbBytes)

	hash := sha256.New()
	hash.Write(bbbBytes)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	assert.Equal(t, "ec4916dd28fc4c10d78e287ca5d9cc51ee1ae73cbfde08c6b37324cbfaac8bc5", mdStr)
}
func TestByteHash(t *testing.T) {
	for i := 0; i < 256; i++ {
		hash := sha256.New()
		hashInput := []byte{byte(i)}

		hash.Write(hashInput)
		md := hash.Sum(nil)
		mdStr := hex.EncodeToString(md)
		fmt.Println(hashInput, mdStr)
	}
}

func TestSha256(t *testing.T) {
	curve := secp256k1.S256()
	hash := sha256.New()

	sender := common.GetBigInt("130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb", 16)

	var vG JacobianPoint
	vG.X, vG.Y = curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)
	//fmt.Println(vG)

	hashInput := sender.Bytes()
	hashInput = append(hashInput, curve.Gx.Bytes()...)
	hashInput = append(hashInput, curve.Gy.Bytes()...)
	hashInput = append(hashInput, vote1ZK.xG.X.Bytes()...)
	hashInput = append(hashInput, vote1ZK.xG.Y.Bytes()...)
	hashInput = append(hashInput, vG.X.Bytes()...)
	hashInput = append(hashInput, vG.Y.Bytes()...)
	hashInput = append(hashInput, vG.Z...)

	hash.Write(hashInput)

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	assert.Equal(t, "056167e4948e5800f8fa96822d0a6c545535a29d76f6fec0ea93ed7d653d19a5", mdStr)
}

func TestCreateZKP(t *testing.T) {
	res, err := CreateZKP(vote1ZK.x, vote1ZK.v, vote1ZK.xG)
	fmt.Println(res, err)
}
