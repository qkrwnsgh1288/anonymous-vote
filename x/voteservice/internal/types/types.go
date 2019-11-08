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
	VoteCheckList []string `json:"vote_checklist"` // todo: change to private

	Progress string `json:"progress"`
	//WhiteList map[string]bool  `json:"whitelist"`
	Test map[string]bool `json:"test"`
}

func NewAgenda() Agenda {
	return Agenda{
		Test: make(map[string]bool),
	}
}

func (a Agenda) String() string {
	return strings.TrimSpace(fmt.Sprintf(`AgendaProposer: %s
AgendaTopic: %s
AgendaContent: %s
WhiteList: %v
VoteCheckList: %v
Progress: %s`, a.AgendaProposer, a.AgendaTopic, a.AgendaContent, a.WhiteList, a.VoteCheckList, a.Progress))
}
