package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stretchr/testify/assert"
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
	Name  string
	Value []string
	//Value map[string]bool
}

func TestMarshal(t *testing.T) {
	a := TestStruct{
		Name:  "name",
		Value: []string{"aaa", "bbb"},
		//Value: make(map[string]bool),
	}
	//a.Value["aaa"]=true
	//a.Value["bbb"]=false
	fmt.Println("a=", a)

	testCdc := codec.New()
	testCdc.RegisterConcrete(TestStruct{}, "test", nil)

	encodedData := testCdc.MustMarshalBinaryBare(a)

	var b TestStruct
	testCdc.MustUnmarshalBinaryBare(encodedData, &b)
	fmt.Println("b=", b)

}
