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

var vote1ZK, vote2ZK, vote3ZK, vote4ZK, vote5ZK ZkInfo

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
	vote4ZK = ZkInfo{
		x: common.GetBigInt("45609652874227667464626768270973794726424735209280228041305418571540205918722", 10),
		xG: Point{
			X: common.GetBigInt("53848600353015527901719462274663386164661133925742909896350594460573835076060", 10),
			Y: common.GetBigInt("70164422168210717916177180083129348172943683555132028499398751968528379547725", 10),
		},
		v: common.GetBigInt("94984205120187139432069320990920412063916872908324719047148175125495068276351", 10),
		w: common.GetBigInt("42339855098479556688828468462327389743363366065105188852724290556244035564679", 10),
		r: common.GetBigInt("49622271312245660322508964836660580525574997933152241594905287377874915852505", 10),
		d: common.GetBigInt("114901097096288926209161487110958354481812855654954796268850840847236931849105", 10),
	}
	vote5ZK = ZkInfo{
		x: common.GetBigInt("19881175679920553899753620052540617072335121233390684924327543896539539956978", 10),
		xG: Point{
			X: common.GetBigInt("58111029405235340908680392565409586671747269249747182797126572361176712701953", 10),
			Y: common.GetBigInt("57655573654433694509805387858917377678913378164403565343914597971897728585410", 10),
		},
		v: common.GetBigInt("12261576258828081605166933950686346642241211206833305008622655965813873226117", 10),
		w: common.GetBigInt("3084439821615479104806473970715407580688219722498805605192016321496161973456", 10),
		r: common.GetBigInt("38607954316385349876852970092145248798289765395738923807016981076608959185705", 10),
		d: common.GetBigInt("102328829924740295112199624718737618441718189850331989146008464262116651510633", 10),
	}

}

func TestAdd(t *testing.T) {
	res1, res2 := Curve.Add(vote1ZK.xG.X, vote1ZK.xG.Y, vote2ZK.xG.X, vote2ZK.xG.Y)
	assert.Equal(t, "19726021177552888194148621436129232937104234324513758427865268224158101547130", res1.String())
	assert.Equal(t, "50952383343742199881927221996840986713139267241507858986150651430342248986684", res2.String())

	res1, res2 = Curve.Add(res1, res2, vote2ZK.xG.X, vote2ZK.xG.Y)
	assert.Equal(t, "43373064182507730467407220464395087632331217744190510800352077052212152643252", res1.String())
	assert.Equal(t, "11632633722541884024350322039309093135403445311747067053405721567890952259613", res2.String())
	fmt.Println(Totalregistered)
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
	res1 := Curve.IsOnCurve(vote1ZK.xG.X, vote1ZK.xG.Y)
	assert.True(t, res1)

	vote1ZK.xG.X = common.GetBigInt("30061975807968526978116138222528932566686537412871265156620434532445965483943", 10)
	res2 := Curve.IsOnCurve(vote1ZK.xG.X, vote1ZK.xG.Y)
	assert.False(t, res2)
}

func TestVG(t *testing.T) {
	var vG JacobianPoint
	vG.X, vG.Y = Curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	assert.Equal(t, "37002400596499253347436146477359872208984972423528869527866238051389129979940", vG.X.String())
	assert.Equal(t, "46104438919360535329359949165853481514194123783534889415421577162302988165861", vG.Y.String())
	assert.Equal(t, "1", vG.Z.String())

	hashZ := sha256.New()
	hashInputZ := vG.Z.Bytes()
	hashZ.Write(hashInputZ)
	assert.Equal(t, "4bf5122f344554c53bde2ebb8cd2b7e3d1600ad631c385a5d7cce23c7785459a", hex.EncodeToString(hashZ.Sum(nil)))

	hash := sha256.New()
	hashInput := vG.X.Bytes()
	hashInput = append(hashInput, vG.Y.Bytes()...)
	hashInput = append(hashInput, vG.Z.Bytes()...)

	hash.Write(hashInput)

	md := hash.Sum(nil)
	hexStr := hex.EncodeToString(md)
	assert.Equal(t, "79cddfd538609ebc9bda527391369e62f8342e3752f9e532a936e1a415ca854a", hexStr)
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
	vG.X, vG.Y = Curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)
	//fmt.Println(vG)

	hashInput := sender.Bytes()
	hashInput = append(hashInput, Curve.Gx.Bytes()...)
	hashInput = append(hashInput, Curve.Gy.Bytes()...)
	hashInput = append(hashInput, vote1ZK.xG.X.Bytes()...)
	hashInput = append(hashInput, vote1ZK.xG.Y.Bytes()...)
	hashInput = append(hashInput, vG.X.Bytes()...)
	hashInput = append(hashInput, vG.Y.Bytes()...)
	hashInput = append(hashInput, vG.Z.Bytes()...)

	hash.Write(hashInput)

	md := hash.Sum(nil)
	hexStr := hex.EncodeToString(md)
	assert.Equal(t, "66a37f3a320ce9caec790203a7e843166d5873381200dd494f0300f92876ef34", hexStr)

	c := common.GetBigInt(hexStr, 16)
	assert.Equal(t, "46424784717785924143576233396036969011302163563202020539237685875916349566772", c.String())
}

func TestMulMod(t *testing.T) {
	c := common.GetBigInt("46424784717785924143576233396036969011302163563202020539237685875916349566772", 10)
	xc := mulMod(vote1ZK.x, c, Curve.N)

	assert.Equal(t, "13276481680719431021304732231458755388712651776249411785446475489619091431771", xc.String())
}

func TestSubMod(t *testing.T) {
	xc := common.GetBigInt("13276481680719431021304732231458755388712651776249411785446475489619091431771", 10)
	r := subMod(vote1ZK.v, xc, Curve.N)

	assert.Equal(t, "32898198925018782135165360898439062728212081324716852088651049454325506359347", r.String())

}

func TestCreateZKP(t *testing.T) {
	senderAddr := "130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb"
	r, vG, err := CreateZKP(senderAddr, vote1ZK.x, vote1ZK.v, vote1ZK.xG)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "32898198925018782135165360898439062728212081324716852088651049454325506359347", r.String())
	assert.Equal(t, "37002400596499253347436146477359872208984972423528869527866238051389129979940", vG.X.String())
	assert.Equal(t, "46104438919360535329359949165853481514194123783534889415421577162302988165861", vG.Y.String())
	assert.Equal(t, "1", vG.Z.String())
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// 2. VerifyZKP //////////////////////////////////////////////////////////////////////////
func TestIsOnCurveVG(t *testing.T) {
	var vG JacobianPoint
	vG.X, vG.Y = Curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	assert.Equal(t, "37002400596499253347436146477359872208984972423528869527866238051389129979940", vG.X.String())
	assert.Equal(t, "46104438919360535329359949165853481514194123783534889415421577162302988165861", vG.Y.String())
	assert.Equal(t, "1", vG.Z.String())

	res1 := Curve.IsOnCurve(vG.X, vG.Y)
	assert.True(t, res1)
}

func TestVerifyZKP(t *testing.T) {
	senderAddr := "130e42fFa25b341b81aC1eb9E53Bc9FF0b16BBeb"

	r := common.GetBigInt("32898198925018782135165360898439062728212081324716852088651049454325506359347", 10)

	var vG JacobianPoint
	vG.X, vG.Y = Curve.ScalarBaseMult(vote1ZK.v.Bytes())
	vG.Z = secp256k1.ZForAffine(vG.X, vG.Y)

	res := VerifyZKP(senderAddr, vote1ZK.xG, r, vG)
	assert.True(t, res)

	rFalse := common.GetBigInt("26559763677273002448312780653724801258659183715468606358605122125501101448492", 10)
	res2 := VerifyZKP(senderAddr, vote1ZK.xG, rFalse, vG)
	assert.False(t, res2)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////// 3. register //////////////////////////////////////////////////////////////////////////
func TestRegister(t *testing.T) {
	addr1 := "2a931b52132d9ca43d5355fecc234fb7b1d02674"
	addr2 := "bb2b24b84d6eee7895b20fdeed2a5b0735046706"
	addr3 := "411609f475dfa35e931bf0f5b59ae933380756a6"
	addr4 := "4851d64bc9cd1817561e113212a1890bd3d84eab"
	addr5 := "c7dbe023afb9c66eb9465e46f3d4a8e3f3da8b30"

	v1_r, v1_vG, _ := CreateZKP(addr1, vote1ZK.x, vote1ZK.v, vote1ZK.xG)
	v2_r, v2_vG, _ := CreateZKP(addr2, vote2ZK.x, vote2ZK.v, vote2ZK.xG)
	v3_r, v3_vG, _ := CreateZKP(addr3, vote3ZK.x, vote3ZK.v, vote3ZK.xG)
	v4_r, v4_vG, _ := CreateZKP(addr4, vote4ZK.x, vote4ZK.v, vote4ZK.xG)
	v5_r, v5_vG, _ := CreateZKP(addr5, vote5ZK.x, vote5ZK.v, vote5ZK.xG)

	assert.True(t, VerifyZKP(addr1, vote1ZK.xG, v1_r, v1_vG))
	assert.True(t, VerifyZKP(addr2, vote2ZK.xG, v2_r, v2_vG))
	assert.True(t, VerifyZKP(addr3, vote3ZK.xG, v3_r, v3_vG))
	assert.True(t, VerifyZKP(addr4, vote4ZK.xG, v4_r, v4_vG))
	assert.True(t, VerifyZKP(addr5, vote5ZK.xG, v5_r, v5_vG))

	Register(addr1, vote1ZK.xG, v1_vG, v1_r)
	Register(addr2, vote2ZK.xG, v2_vG, v2_r)
	Register(addr3, vote3ZK.xG, v3_vG, v3_r)
	Register(addr4, vote4ZK.xG, v4_vG, v4_r)
	Register(addr5, vote5ZK.xG, v5_vG, v5_r)

	assert.Equal(t, 5, Totalregistered)

	err := finishRegistrationPhase()
	if err != nil {
		fmt.Println(err)
	}
	//for i:=0; i<Totalregistered; i++{
	//	fmt.Println(Voters[i].ReconstructedKey.X, Voters[i].ReconstructedKey.Y)
	//}

	assert.Equal(t, "13640588435166186727072570872841920017273057013114604476956539355021275854144", Voters[0].ReconstructedKey.X.String())
	assert.Equal(t, "90715709871810868701227023413915222907311739236101232519930156567199700809709", Voters[0].ReconstructedKey.Y.String())
	assert.Equal(t, "29649779067344416281603749355821590019952822407947728238700922695212875405379", Voters[1].ReconstructedKey.X.String())
	assert.Equal(t, "48119921710368519564274046835194847709739768059431168092090665719136860379594", Voters[1].ReconstructedKey.Y.String())
	assert.Equal(t, "87261116829116053902355813311953430020061743503720332457245262174159519523247", Voters[2].ReconstructedKey.X.String())
	assert.Equal(t, "50725870913308555836107121124447058283822308921418136225940978370414232207742", Voters[2].ReconstructedKey.Y.String())
	assert.Equal(t, "19661706804235466481235433058866766500314639389480257333831084820319603447544", Voters[3].ReconstructedKey.X.String())
	assert.Equal(t, "28751956283219203239390587914594206013138850149638812878587734004990969454344", Voters[3].ReconstructedKey.Y.String())
	assert.Equal(t, "113801752441250897077142944944718213653143011080785242322121729713585417265186", Voters[4].ReconstructedKey.X.String())
	assert.Equal(t, "16070335387338018856501682106237699376075221278867300502691659950109073760271", Voters[4].ReconstructedKey.Y.String())
}
