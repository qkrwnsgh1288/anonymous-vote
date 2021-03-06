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

	WhiteList         []string `json:"whitelist"`
	TotalRegistered   int      `json:"total_registered"`
	TotalVoteComplete int      `json:"total_vote_complete"`

	State    crypto.State `json:"state"`
	Voters   []SVoter     `json:"voter"`
	FinalYes int          `json:"final_yes"`
	FinalNo  int          `json:"final_no"`
}

func NewAgenda() Agenda {
	return Agenda{}
}

func (a Agenda) String() string {
	return strings.TrimSpace(fmt.Sprintf(`AgendaProposer: %s
AgendaTopic: %s
AgendaContent: %s
WhiteList: %v`, a.AgendaProposer, a.AgendaTopic, a.AgendaContent, a.WhiteList))
}

type SPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func MakeDefaultSPoint() SPoint {
	return SPoint{X: "", Y: ""}
}

type SVoter struct {
	Addr             string `json:"address"`
	RegisteredKey    SPoint `json:"registered_key"`
	ReconstructedKey SPoint `json:"reconstructed_key"`
	Commitment       string `json:"commitment"`
	Vote             SPoint `json:"vote"`
}

func MakeDefaultSVoter() SVoter {
	return SVoter{
		Addr:             "",
		RegisteredKey:    MakeDefaultSPoint(),
		ReconstructedKey: MakeDefaultSPoint(),
		Commitment:       "",
		Vote:             MakeDefaultSPoint(),
	}
}
