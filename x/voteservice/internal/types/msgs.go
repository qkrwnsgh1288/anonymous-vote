package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	oproto "github.com/golang/protobuf/proto"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types/proto"
)

const RouterKey = ModuleName // this was defined in your key.go file

var (
	_ sdk.Msg = MsgMakeAgenda{}
	_ sdk.Msg = MsgVoteAgenda{}
)

// MsgMakeAgenda
type MsgMakeAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`

	AgendaTopic   string   `json:"agenda_topic"`
	AgendaContent string   `json:"agenda_content"`
	WhiteList     []string `json:"whitelist"`
	VoteCheckList []string `json:"vote_checklist"`
	//WhiteList     map[string]bool `json:"whitelist"`
	Test map[string]bool `json:"test"`
}

func NewMsgMakeAgenda(agendaProposer sdk.AccAddress, agendaTopic string, agendaContent string, whiteList []string) MsgMakeAgenda {
	/*whiteList := make(map[string]bool)
	for _, addr := range whiteListSlice {
		whiteList[addr] = false
	}*/
	var voteCheckList []string
	for i := 0; i < len(whiteList); i++ {
		voteCheckList = append(voteCheckList, "empty")
	}
	test := make(map[string]bool)
	test["A"] = true

	return MsgMakeAgenda{
		AgendaProposer: agendaProposer,
		AgendaTopic:    agendaTopic,
		AgendaContent:  agendaContent,
		WhiteList:      whiteList,
		VoteCheckList:  voteCheckList,
		Test:           test,
	}
}
func (msg MsgMakeAgenda) GetPb() proto.MsgMakeAgenda {
	return proto.MsgMakeAgenda{
		AgendaProposer: msg.AgendaProposer,
		AgendaTopic:    msg.AgendaTopic,
		AgendaContent:  msg.AgendaContent,
		WhiteList:      msg.WhiteList,
		VoteCheckList:  msg.VoteCheckList,
		Test:           msg.Test,
	}
}
func (msg MsgMakeAgenda) Route() string { return RouterKey }
func (msg MsgMakeAgenda) Type() string  { return "make_agenda" }
func (msg MsgMakeAgenda) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrUnknownRequest("AgendaTopic cannot be empty")
	}
	if len(msg.AgendaContent) == 0 {
		return sdk.ErrUnknownRequest("AgendaContent cannot be empty")
	}
	if len(msg.WhiteList) == 0 {
		return sdk.ErrUnknownRequest("WhiteList cannot be empty")
	}
	return nil
}
func (msg MsgMakeAgenda) GetSignBytes() []byte {
	//return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
	pb := msg.GetPb()
	b, err := oproto.Marshal(&pb)
	if err != nil {
		panic(err)
	}
	return b
}
func (msg MsgMakeAgenda) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.AgendaProposer}
}

// MsgVoteAgenda
type MsgVoteAgenda struct {
	AgendaTopic string         `json:"agenda_topic"`
	VoteAddr    sdk.AccAddress `json:"vote_addr"`
	YesOrNo     string         `json:"yes_or_no"`
}

func NewMsgVoteAgenda(voteAddr sdk.AccAddress, topic string, yesOrNo string) MsgVoteAgenda {
	return MsgVoteAgenda{
		AgendaTopic: topic,
		VoteAddr:    voteAddr,
		YesOrNo:     yesOrNo,
	}
}
func (msg MsgVoteAgenda) Route() string { return RouterKey }
func (msg MsgVoteAgenda) Type() string  { return "vote_agenda" }
func (msg MsgVoteAgenda) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrUnknownRequest("AgendaTopic cannot be empty")
	}

	if !(msg.YesOrNo == "yes" || msg.YesOrNo == "no") {
		return sdk.ErrUnknownRequest("Answer should be yes or no")
	}

	return nil
}
func (msg MsgVoteAgenda) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgVoteAgenda) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.VoteAddr}
}
