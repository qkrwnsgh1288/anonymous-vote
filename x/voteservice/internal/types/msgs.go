package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
)

const RouterKey = ModuleName // this was defined in your key.go file

var (
	_ sdk.Msg = MsgMakeAgenda{}
	_ sdk.Msg = MsgVoteAgenda{}
)

type StringPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
}

//func (s StringPoint) String() string {
//	return fmt.Sprintf("%s, %s", s.X, s.Y)
//}

// MsgMakeAgenda
type MsgMakeAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`

	AgendaTopic   string   `json:"agenda_topic"`
	AgendaContent string   `json:"agenda_content"`
	SetupList     []string `json:"setuplist"`
	VoteCheckList []string `json:"vote_checklist"`

	State            crypto.State  `json:"state"`
	RegisteredKey    []StringPoint `json:"registered_key"`
	ReconstructedKey []StringPoint `json:"reconstructed_key"`
	Commitment       string        `json:"commitment"`
	Vote             []StringPoint `json:"vote"`
}

func NewMsgMakeAgenda(agendaProposer sdk.AccAddress, agendaTopic string, agendaContent string, whiteList []string) MsgMakeAgenda {
	var voteCheckList []string
	for i := 0; i < len(whiteList); i++ {
		voteCheckList = append(voteCheckList, "empty")
	}
	//a := make([]StringPoint, 0)
	//a = append(a, StringPoint{"x1", "y1"})
	//a = append(a, StringPoint{"x2", "y2"})

	return MsgMakeAgenda{
		AgendaProposer: agendaProposer,

		AgendaTopic:   agendaTopic,
		AgendaContent: agendaContent,
		SetupList:     whiteList,
		VoteCheckList: voteCheckList,

		State: crypto.SETUP,
		//RegisteredKey:  make([]StringPoint, 0),
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
	if len(msg.SetupList) == 0 {
		return sdk.ErrUnknownRequest("SetupList cannot be empty")
	}
	return nil
}
func (msg MsgMakeAgenda) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
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
