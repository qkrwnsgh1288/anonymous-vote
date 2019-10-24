package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto/secp256k1"
	"github.com/stretchr/testify/assert"
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
	res1 := curve.IsOnCurve(vote1ZK.xG.X, vote1ZK.xG.Y)
	assert.True(t, res1)

	vote1ZK.xG.X = common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483943", 10)
	res2 := curve.IsOnCurve(vote1ZK.xG.X, vote1ZK.xG.Y)
	assert.False(t, res2)
}

func TestVG(t *testing.T) {
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
	hexStr := hex.EncodeToString(md)
	assert.Equal(t, "3671c70fe36d399d158f71abd58f782cba4ca924d073d3c27630ac1eb050fa7a", hexStr)
}

/*func TestLittleEndian(t *testing.T) {
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
	hexStr := hex.EncodeToString(md)
	assert.Equal(t, "ec4916dd28fc4c10d78e287ca5d9cc51ee1ae73cbfde08c6b37324cbfaac8bc5", hexStr)
}
func TestByteHash(t *testing.T) {
	for i := 0; i < 256; i++ {
		hash := sha256.New()
		hashInput := []byte{byte(i)}

		hash.Write(hashInput)
		md := hash.Sum(nil)
		hexStr := hex.EncodeToString(md)
		fmt.Println(hashInput, hexStr)
	}
}*/

////////////////////////////////////////////////////////// 1. CreateZKP //////////////////////////////////////////////////////////////////////////
func TestSha256(t *testing.T) {
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
	hexStr := hex.EncodeToString(md)
	assert.Equal(t, "056167e4948e5800f8fa96822d0a6c545535a29d76f6fec0ea93ed7d653d19a5", hexStr)

	c := common.GetBigInt(hexStr, 16)
	assert.Equal(t, "2433665450586170755636384720285970258250106297820018300061622432952343140773", c.String())
}

func TestMulMod(t *testing.T) {
	c := common.GetBigInt("056167e4948e5800f8fa96822d0a6c545535a29d76f6fec0ea93ed7d653d19a5", 16)
	xc := mulMod(vote1ZK.x, c, curve.N)

	assert.Equal(t, "19614916928465210708157312476173016858265549385497657515492402818443496342626", xc.String())
}

func TestSubMod(t *testing.T) {
	xc := common.GetBigInt("19614916928465210708157312476173016858265549385497657515492402818443496342626", 10)
	r := subMod(vote1ZK.v, xc, curve.N)

	assert.Equal(t, "26559763677273002448312780653724801258659183715468606358605122125501101448492", r.String())

}

func TestCreateZKP(t *testing.T) {
	senderAddr := "130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb"
	res, err := CreateZKP(senderAddr, vote1ZK.x, vote1ZK.v, vote1ZK.xG)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 4, len(res))
	assert.Equal(t, "26559763677273002448312780653724801258659183715468606358605122125501101448492", res[0].String())
	assert.Equal(t, "37002400596499253347436146477359872208984972423528869527866238051389129979940", res[1].String())
	assert.Equal(t, "46104438919360535329359949165853481514194123783534889415421577162302988165861", res[2].String())
	assert.Equal(t, "1", res[3].String())
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// 2. VerifyZKP //////////////////////////////////////////////////////////////////////////
func TestIsOnCurveVG(t *testing.T) {
	var vG JacobianPoint
	vG.X, vG.Y = curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	assert.Equal(t, "37002400596499253347436146477359872208984972423528869527866238051389129979940", vG.X.String())
	assert.Equal(t, "46104438919360535329359949165853481514194123783534889415421577162302988165861", vG.Y.String())

	res1 := curve.IsOnCurve(vG.X, vG.Y)
	assert.True(t, res1)
}

func TestVerifyZKP(t *testing.T) {
	senderAddr := "130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb"
	r := common.GetBigInt("26559763677273002448312780653724801258659183715468606358605122125501101448492", 10)

	var vG JacobianPoint
	vG.X, vG.Y = curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	res := VerifyZKP(senderAddr, vote1ZK.xG, r, vG)
	assert.True(t, res)

	senderAddr2 := "130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBec"
	res2 := VerifyZKP(senderAddr2, vote1ZK.xG, r, vG)
	assert.False(t, res2)
}
