package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}
func (msg MsgSetName) Route() string { return RouterKey }
func (msg MsgSetName) Type() string  { return "set_name" }
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}
func (msg MsgBuyName) Route() string { return RouterKey }
func (msg MsgBuyName) Type() string  { return "buy_name" }
func (msg MsgBuyName) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

// MsgDeleteName defines a DeleteName message
type MsgDeleteName struct {
	Name  string         `json:"name"`
	Owner sdk.AccAddress `json:"owner"`
}

func NewMsgDeleteName(name string, owner sdk.AccAddress) MsgDeleteName {
	return MsgDeleteName{
		Name:  name,
		Owner: owner,
	}
}
func (msg MsgDeleteName) Route() string { return RouterKey }
func (msg MsgDeleteName) Type() string  { return "delete_name" }
func (msg MsgDeleteName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	return nil
}
func (msg MsgDeleteName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgDeleteName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// 1. MakeAgenda
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
