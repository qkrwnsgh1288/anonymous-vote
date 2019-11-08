package crypto

import (
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestDoubleM(t *testing.T) {
	a := MakeDefaultJacobianPoint()
	a.X = big.NewInt(1)
	a.Y = big.NewInt(2)
	a.Z = big.NewInt(3)
	doubleM(&a)

	assert.Equal(t, "115792089237316195423570985008687907853269984665640564039457584007908834671640", a.X.String())
	assert.Equal(t, "115792089237316195423570985008687907853269984665640564039457584007908834671652", a.Y.String())
	assert.Equal(t, "12", a.Z.String())
}

func TestAddMixedM(t *testing.T) {
	a := MakeDefaultJacobianPoint()
	b := MakeDefaultPoint()

	a.X = big.NewInt(3)
	a.Y = big.NewInt(5)
	a.Z = big.NewInt(1)
	b.X = big.NewInt(3)
	b.Y = big.NewInt(5)

	AddMixedM(&a, b)

	assert.Equal(t, "129", a.X.String())
	assert.Equal(t, "115792089237316195423570985008687907853269984665640564039457584007908834671280", a.Y.String())
	assert.Equal(t, "10", a.Z.String())

	res1, res2 := Curve.AffineFromJacobian(a.X, a.Y, a.Z)
	assert.Equal(t, "19684655170343753222007067451476944335055897393158895886707789281344501894184", res1.String())
	assert.Equal(t, "51064311353656442181794804388831367363292063237547488741400794547487796090203", res2.String())

	ToZ1(&a, Curve.P)
	assert.Equal(t, "19684655170343753222007067451476944335055897393158895886707789281344501894184", a.X.String())
	assert.Equal(t, "51064311353656442181794804388831367363292063237547488741400794547487796090203", a.Y.String())
	assert.Equal(t, "1", a.Z.String())
}

func TestAddmixedM2(t *testing.T) {
	temp := MakeDefaultJacobianPoint()
	vote := MakeDefaultPoint()

	temp.X = common.GetBigInt("74467521297191565178513432739089809393952090909697815815307594610304762694309", 10)
	temp.Y = common.GetBigInt("95455015536110063042853633233820093016164655237802526178747111496863700373322", 10)
	temp.Z = big.NewInt(1)
	vote.X = common.GetBigInt("92515557875610636761004533075988517481553681564945778683886252164442967911720", 10)
	vote.Y = common.GetBigInt("29569670867748894430471170272272888267757597049512796102421410804787799190400", 10)

	AddMixedM(&temp, vote)

	assert.Equal(t, "57573336721267290568522211640320386453566183723721414949685575050653965439170", temp.X.String())
	assert.Equal(t, "92113145688917356339994382103993184066300731337667578833771166601659128048974", temp.Y.String())
	assert.Equal(t, "18048036578419071582491100336898708087601590655247962868578657554138205217411", temp.Z.String())
}