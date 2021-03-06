package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

var (
	_ sdk.Msg = MsgMakeAgenda{}
	_ sdk.Msg = MsgRegisterByVoter{}
	_ sdk.Msg = MsgRegisterByProposer{}
	_ sdk.Msg = MsgVoteAgenda{}
	_ sdk.Msg = MsgTally{}
)

// 1. MsgMakeAgenda
type MsgMakeAgenda struct {
	AgendaProposer sdk.AccAddress `json:"agenda_proposer"`
	AgendaTopic    string         `json:"agenda_topic"`
	AgendaContent  string         `json:"agenda_content"`

	WhiteList []string `json:"whitelist"`
	Voters    []SVoter `json:"voter"`
}

func NewMsgMakeAgenda(agendaProposer sdk.AccAddress, agendaTopic string, agendaContent string, whiteList []string) MsgMakeAgenda {
	var voterList []SVoter
	for i := 0; i < len(whiteList); i++ {
		voterList = append(voterList, MakeDefaultSVoter())
	}

	return MsgMakeAgenda{
		AgendaProposer: agendaProposer,
		AgendaTopic:    agendaTopic,
		AgendaContent:  agendaContent,

		WhiteList: whiteList,
		//Voters:    voterList,
	}
}
func (msg MsgMakeAgenda) Route() string { return RouterKey }
func (msg MsgMakeAgenda) Type() string  { return "make_agenda" }
func (msg MsgMakeAgenda) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return ErrAgendaTopicIsEmpty(DefaultCodespace)
	}
	if len(msg.AgendaContent) == 0 {
		return ErrAgendaContentIsEmpty(DefaultCodespace)
	}
	if len(msg.WhiteList) == 0 {
		return ErrWhiteListIsEmpty(DefaultCodespace)
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
		return ErrAgendaTopicIsEmpty(DefaultCodespace)
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
		return ErrAgendaTopicIsEmpty(DefaultCodespace)
	}

	return nil
}
func (msg MsgRegisterByProposer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgRegisterByProposer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ProposerAddr}
}

// 4. MsgVoteAgenda
type MsgVoteAgenda struct {
	AgendaTopic string         `json:"agenda_topic"`
	VoteAddr    sdk.AccAddress `json:"vote_addr"`
	ZkInfo      []string       `json:"zk_info"`
	YesOrNo     string         `json:"yes_or_no"`
}

func NewMsgVoteAgenda(voteAddr sdk.AccAddress, topic string, yesOrNo string, zkInfos []string) MsgVoteAgenda {
	return MsgVoteAgenda{
		AgendaTopic: topic,
		VoteAddr:    voteAddr,
		ZkInfo:      zkInfos,
		YesOrNo:     yesOrNo,
	}
}
func (msg MsgVoteAgenda) Route() string { return RouterKey }
func (msg MsgVoteAgenda) Type() string  { return "vote_agenda" }
func (msg MsgVoteAgenda) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return ErrAgendaTopicIsEmpty(DefaultCodespace)
	}

	if !(msg.YesOrNo == "yes" || msg.YesOrNo == "no") {
		return ErrInvalidAnswer(DefaultCodespace)
	}

	return nil
}
func (msg MsgVoteAgenda) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgVoteAgenda) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.VoteAddr}
}

// 5. MsgTally
type MsgTally struct {
	AgendaTopic  string         `json:"agenda_topic"`
	ProposerAddr sdk.AccAddress `json:"proposer_addr"`
}

func NewMsgTally(proposerAddr sdk.AccAddress, topic string) MsgTally {
	return MsgTally{
		AgendaTopic:  topic,
		ProposerAddr: proposerAddr,
	}
}
func (msg MsgTally) Route() string { return RouterKey }
func (msg MsgTally) Type() string  { return "tally" }
func (msg MsgTally) ValidateBasic() sdk.Error {
	// todo: more
	if len(msg.AgendaTopic) == 0 {
		return ErrAgendaTopicIsEmpty(DefaultCodespace)
	}

	return nil
}
func (msg MsgTally) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgTally) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ProposerAddr}
}
