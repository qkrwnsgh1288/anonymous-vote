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

	WhiteList     []string `json:"whitelist"`
	VoteCheckList []bool   `json:"vote_checklist"` // todo: change to private
	ProCount      uint32   `json:"pro_count"`      // todo: change to private
	NegCount      uint32   `json:"neg_count"`      // todo: change to private
	//WhiteList map[string]bool  `json:"whitelist"`
}

func NewAgenda() Agenda {
	return Agenda{}
}

func (a Agenda) String() string {
	return strings.TrimSpace(fmt.Sprintf(`AgendaProposer: %s
AgendaTopic: %s
AgendaContent: %s
WhiteList: %v
VoteCheckList: %v
ProCount: %d
NegCount: %d`, a.AgendaProposer, a.AgendaTopic, a.AgendaContent, a.WhiteList, a.VoteCheckList, a.ProCount, a.NegCount))
}
