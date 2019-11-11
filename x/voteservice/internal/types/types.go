package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
	"strings"
)

type Agenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`
	AgendaTopic    string         `json:"agenda_topic"`
	AgendaContent  string         `json:"agenda_content"`

	SetupList       []string `json:"setuplist"`
	VoteCheckList   []string `json:"vote_checklist"` // todo: change to private
	Progress        string   `json:"progress"`
	TotalRegistered int      `json:"total_registered"`

	State crypto.State `json:"state"`
	Voter []SVoter     `json:"voter"`
}

func NewAgenda() Agenda {
	return Agenda{}
}

func (a Agenda) String() string {
	return strings.TrimSpace(fmt.Sprintf(`AgendaProposer: %s
AgendaTopic: %s
AgendaContent: %s
SetupList: %v
VoteCheckList: %v
Progress: %s`, a.AgendaProposer, a.AgendaTopic, a.AgendaContent, a.SetupList, a.VoteCheckList, a.Progress))
}
