package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type Agenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`
	AgendaTopic    string         `json:"agenda_topic"`
	AgendaContent  string         `json:"agenda_content"`

	Voters   []string `json:"voters"`
	ProCount uint32   `json:"pro_count"`
	NegCount uint32   `json:"neg_count"`
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
