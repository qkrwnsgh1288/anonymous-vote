package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/crypto"
)

const RouterKey = ModuleName // this was defined in your key.go file

var (
	_ sdk.Msg = MsgMakeAgenda{}
	_ sdk.Msg = MsgRegisterByVoter{}
	_ sdk.Msg = MsgRegisterByProposer{}
	_ sdk.Msg = MsgVoteAgenda{}
)

type SPoint struct {
	X string `json:"x"`
	Y string `json:"y"`
}

func MakeDefaultStringPoint() SPoint {
	return SPoint{X: "", Y: ""}
}

type SVoter struct {
	Addr             string `json:"address"`
	RegisteredKey    SPoint `json:"registered_key"`
	ReconstructedKey SPoint `json:"reconstructed_key"`
	Commitment       string `json:"commitment"`
	Vote             SPoint `json:"vote"`
}

func MakeDefaultStringVoter() SVoter {
	return SVoter{
		Addr:             "",
		RegisteredKey:    MakeDefaultStringPoint(),
		ReconstructedKey: MakeDefaultStringPoint(),
		Commitment:       "",
		Vote:             MakeDefaultStringPoint(),
	}
}

//func (s SPoint) String() string {
//	return fmt.Sprintf("%s, %s", s.X, s.Y)
//}

// 1. MsgMakeAgenda
type MsgMakeAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`
	AgendaTopic    string         `json:"agenda_topic"`
	AgendaContent  string         `json:"agenda_content"`

	SetupList     []string `json:"setuplist"`
	VoteCheckList []string `json:"vote_checklist"`

	State crypto.State `json:"state"`
	Voter []SVoter     `json:"voter"`
}

func NewMsgMakeAgenda(agendaProposer sdk.AccAddress, agendaTopic string, agendaContent string, whiteList []string) MsgMakeAgenda {
	var voteCheckList []string
	var voterList []SVoter
	for i := 0; i < len(whiteList); i++ {
		voteCheckList = append(voteCheckList, "empty")
		voterList = append(voterList, MakeDefaultStringVoter())
	}

	return MsgMakeAgenda{
		AgendaProposer: agendaProposer,
		AgendaTopic:    agendaTopic,
		AgendaContent:  agendaContent,

		SetupList:     whiteList,
		VoteCheckList: voteCheckList,

		State: crypto.SETUP,
		Voter: voterList,
	}
}
func (msg MsgMakeAgenda) Route() string { return RouterKey }
func (msg MsgMakeAgenda) Type() string  { return "make_agenda" }
func (msg MsgMakeAgenda) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrUnknownRequest("AgendaTopic cannot be empty")
	}
	if len(msg.AgendaContent) == 0 {
		return sdk.ErrUnknownRequest("AgendaContent cannot be empty")
	}
	if len(msg.SetupList) == 0 {
		return sdk.ErrUnknownRequest("SetupList cannot be empty")
	}
	return nil
}
func (msg MsgMakeAgenda) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgMakeAgenda) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.AgendaProposer}
}

// 2. MsgRegisterByVoter
type MsgRegisterByVoter struct {
	AgendaTopic string         `json:"agenda_topic"`
	VoteAddr    sdk.AccAddress `json:"vote_addr"`
	ZkInfo      []string       `json:"zk_info"`
}

func NewMsgRegisterByVoter(voteAddr sdk.AccAddress, topic string, zkInfos []string) MsgRegisterByVoter {
	return MsgRegisterByVoter{
		AgendaTopic: topic,
		VoteAddr:    voteAddr,
		ZkInfo:      zkInfos,
	}
}
func (msg MsgRegisterByVoter) Route() string { return RouterKey }
func (msg MsgRegisterByVoter) Type() string  { return "register_by_voter" }
func (msg MsgRegisterByVoter) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrUnknownRequest("AgendaTopic cannot be empty")
	}

	return nil
}
func (msg MsgRegisterByVoter) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgRegisterByVoter) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.VoteAddr}
}

// 3. MsgRegisterByProposer
type MsgRegisterByProposer struct {
	AgendaTopic  string         `json:"agenda_topic"`
	ProposerAddr sdk.AccAddress `json:"proposer_addr"`
}

func NewMsgRegisterByProposer(proposerAddr sdk.AccAddress, topic string) MsgRegisterByProposer {
	return MsgRegisterByProposer{
		AgendaTopic:  topic,
		ProposerAddr: proposerAddr,
	}
}
func (msg MsgRegisterByProposer) Route() string { return RouterKey }
func (msg MsgRegisterByProposer) Type() string  { return "register_by_proposer" }
func (msg MsgRegisterByProposer) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return sdk.ErrUnknownRequest("AgendaTopic cannot be empty")
	}

	return nil
}
func (msg MsgRegisterByProposer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgRegisterByProposer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ProposerAddr}
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
func (msg MsgVoteAgenda) ValidateBasic() sdk.Error {
	// todo: more
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
