package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	oproto "github.com/golang/protobuf/proto"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types/proto"
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

func TestProto(t *testing.T) {
	msg := MsgMakeAgenda{
		AgendaTopic:   "A",
		AgendaContent: "B",
		Test:          make(map[string]bool),
	}
	msg.Test["AAA"] = true

	msgProto := proto.MsgMakeAgenda{
		AgendaTopic:   msg.AgendaTopic,
		AgendaContent: msg.AgendaContent,
		Test:          msg.Test,
	}
	fmt.Println(msgProto)

	b, err := oproto.Marshal(&msgProto)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	var de proto.MsgMakeAgenda
	err = oproto.Unmarshal(b, &de)
	fmt.Println(de)

}
