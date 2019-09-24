package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

var (
	_ sdk.Msg = MsgMakeAgenda{}
	_ sdk.Msg = MsgVoteAgenda{}
)

// MsgMakeAgenda
type MsgMakeAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`
	AgendaTopic    string         `json:"agenda_topic"`
	AgendaContent  string         `json:"agenda_content"`
}

func NewMsgMakeAgenda(agendaProposer sdk.AccAddress, agendaTopic string, agendaContent string) MsgMakeAgenda {
	return MsgMakeAgenda{
		AgendaProposer: agendaProposer,
		AgendaTopic:    agendaTopic,
		AgendaContent:  agendaContent,
	}
}
func (msg MsgMakeAgenda) Route() string { return RouterKey }
func (msg MsgMakeAgenda) Type() string  { return "make_agenda" }

// todo: more
func (msg MsgMakeAgenda) ValidateBasic() sdk.Error {
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrUnknownRequest("AgendaTopic cannot be empty")
	}
	if len(msg.AgendaContent) == 0 {
		return sdk.ErrUnknownRequest("AgendaContent cannot be empty")
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

// todo: more
func (msg MsgVoteAgenda) ValidateBasic() sdk.Error {
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
