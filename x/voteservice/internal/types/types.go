package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type Agenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda-proposer"`
	AgendaTopic    string         `json:"agenda-topic"`
	AgendaContent  string         `json:"agenda-content"`

	Voters   []string `json:"voters"`
	ProCount uint32   `json:"pro-count"`
	NegCount uint32   `json:"neg-count"`
}

func NewAgenda() Agenda {
	return Agenda{}
}

func (a Agenda) String() string {
	return strings.TrimSpace(fmt.Sprintf(`AgendaProposer: %s
AgendaTopic: %s
AgendaContent: %s
Voters: %v
ProCount: %d
NegCount: %d`, a.AgendaProposer, a.AgendaTopic, a.AgendaContent, a.Voters, a.ProCount, a.NegCount))
}

// Whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// Returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}
