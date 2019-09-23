package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgMakeAgenda
type MsgMakeAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda-proposer"`
	AgendaTopic    string         `json:"agenda-topic"`
	AgendaContent  string         `json:"agenda-content"`
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
func (msg MsgMakeAgenda) ValidateBasic() sdk.Error {
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrInternal("AgendaTopic cannot be empty")
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
