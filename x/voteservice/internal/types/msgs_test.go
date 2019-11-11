package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strings"
	"testing"
)

func TestYesOrNo(t *testing.T) {
	case1 := "Yes"
	case2 := " yEs "
	case3 := "    No "

	assert.Equal(t, "yes", strings.TrimSpace(strings.ToLower(case1)))
	assert.Equal(t, "yes", strings.TrimSpace(strings.ToLower(case2)))
	assert.Equal(t, "no", strings.TrimSpace(strings.ToLower(case3)))
}

type TestStruct struct {
	Name     string
	Value    []string
	VoteList []string
	//Value map[string]bool
}

func TestMarshal(t *testing.T) {
	valList := []string{"AAA", "BBB", "CCC"}
	var voteList []string
	//for _, val := range valList {
	//	fmt.Println("valid check", val)
	//	voteList = append(voteList, false)
	//}
	for i := 0; i < len(valList); i++ {
		voteList = append(voteList, "notyet")
	}
	a := TestStruct{
		Name:     "name",
		Value:    valList,
		VoteList: voteList,
		//Value: make(map[string]bool),
	}
	//a.Value["aaa"]=true
	//a.Value["bbb"]=false
	a.VoteList[1] = "yes"

	fmt.Println("a=", a)

	testCdc := codec.New()
	testCdc.RegisterConcrete(TestStruct{}, "test", nil)

	encodedData := testCdc.MustMarshalBinaryBare(a)

	var b TestStruct
	testCdc.MustUnmarshalBinaryBare(encodedData, &b)
	fmt.Println("b=", b)
}

type testst struct {
	Name          *string
	Value         *int
	Value2        *big.Int
	Value3        float64 `amino:"unsafe"`
	State         crypto.State
	RegisteredKey []SPoint `json:"registered_key"`
}

func TestMarshal2(t *testing.T) {
	str := "AAA"
	value := 111
	//value2 := big.NewInt(222)
	test := testst{
		Name:          &str,
		Value:         &value,
		Value2:        big.NewInt(222),
		Value3:        12.12,
		State:         crypto.SIGNUP,
		RegisteredKey: make([]SPoint, 0),
	}
	fmt.Println(*test.Name, *test.Value, test.Value2, test.Value3, test.State, test.RegisteredKey)

	cdc := codec.New()
	encode := cdc.MustMarshalBinaryBare(test)

	var decode testst
	cdc.MustUnmarshalBinaryBare(encode, &decode)
	fmt.Println(*decode.Name, *test.Value, test.Value2, test.Value3, test.State, test.RegisteredKey)
}
