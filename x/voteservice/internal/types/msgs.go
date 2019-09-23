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

// 1. Agenda
type MsgAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda-proposer"`
	AgendaKey      uint64         `json:"agenda-key"`
	AgendaContent  string         `json:"agenda-content"`

	Voters   []string `json:"voters"`
	ProCount uint32   `json:"pro-count"`
	NegCount uint32   `json:"neg-count"`
}

func NewMsgAgenda(agendaProposer sdk.AccAddress, agendaKey uint64, agendaContent string) MsgAgenda {
	return MsgAgenda{
		AgendaProposer: agendaProposer,
		AgendaKey:      agendaKey,
		AgendaContent:  agendaContent,
	}
}
func (msg MsgAgenda) Route() string { return RouterKey }
func (msg MsgAgenda) Type() string  { return "agenda" }
func (msg MsgAgenda) ValidateBasic() sdk.Error {
	if msg.AgendaKey == 0 {
		return sdk.ErrInternal("AgendaKey cannot be 0")
	}
	if len(msg.AgendaContent) == 0 {
		return sdk.ErrUnknownRequest("AgendaContent cannot be empty")
	}
	return nil
}
func (msg MsgAgenda) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgAgenda) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.AgendaProposer}
}
