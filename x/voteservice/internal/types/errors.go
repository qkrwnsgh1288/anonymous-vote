package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	// basic (100 ~ 200)
	SomethingIsBad          sdk.CodeType = 100
	CodeNameDoesNotExist    sdk.CodeType = 101
	InvalidAnswer           sdk.CodeType = 102
	AgendaTopicAlreadyExist sdk.CodeType = 103
	AgendaTopicDoesNotExist sdk.CodeType = 104
	AgendaTopicIsEmpty      sdk.CodeType = 105
	AgendaContentIsEmpty    sdk.CodeType = 106
	WhiteListIsEmpty        sdk.CodeType = 107

	// about zk (201 ~ 299)
	InvalidPubkeyInCreateZKP  sdk.CodeType = 201
	InvalidVerifyZKP          sdk.CodeType = 202
	DoesNotRegisterAddress    sdk.CodeType = 203
	AlreadyRegisterd          sdk.CodeType = 204
	InvalidTotalRegisteredCnt sdk.CodeType = 205
	DoNotHavePermission       sdk.CodeType = 206
	InvalidVerify1outof2ZKP   sdk.CodeType = 207
	VotingIsNotFinished       sdk.CodeType = 208

	StateIsNotSETUP      sdk.CodeType = 291
	StateIsNotSIGNUP     sdk.CodeType = 292
	StateIsNotCOMMITMENT sdk.CodeType = 293
	StateIsNotVOTE       sdk.CodeType = 294
	StateIsNotFINISHED   sdk.CodeType = 295
)

// basic (100 ~ 200)
func ErrSomethingIsBad(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, SomethingIsBad, "something bad happened")
}
func ErrNameDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNameDoesNotExist, "Name does not exist")
}
func ErrInvalidAnswer(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, InvalidAnswer, "Answer is not valid. It should be yes or no")
}
func ErrAgendaTopicAlreadyExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AgendaTopicAlreadyExist, "AgendaTopic already exist")
}
func ErrAgendaTopicDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AgendaTopicDoesNotExist, "AgendaTopic does not exist")
}
func ErrAgendaTopicIsEmpty(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AgendaTopicIsEmpty, "AgendaTopic cannot be empty")
}
func ErrAgendaContentIsEmpty(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AgendaContentIsEmpty, "AgendaContent cannot be empty")
}
func ErrWhiteListIsEmpty(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, WhiteListIsEmpty, "WhiteList cannot be empty")
}

// about zk (201 ~ 299)
func ErrInvalidPubkeyInCreateZKP(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, InvalidPubkeyInCreateZKP, "Error occured in CreateZKP: XG is not pubKey")
}
func ErrInvalidVerifyZKP(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, InvalidVerifyZKP, "VerifyZKP does not valid")
}
func ErrDoesNotRegisterAddress(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, DoesNotRegisterAddress, "This address does not Registered")
}
func ErrAlreadyRegisterd(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AlreadyRegisterd, "This address already registered")
}
func ErrInvalidTotalRegisteredCnt(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, InvalidTotalRegisteredCnt, "Total registered is smaller than minimum")
}
func ErrDoNotHavePermission(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, DoNotHavePermission, "You don't have permission. It is only possible for the owner")
}
func ErrInvalidVerify1outof2ZKP(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, InvalidVerify1outof2ZKP, "It is Invalid verify 1 outof 2 ZKP")
}
func ErrVotingIsNotFinished(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, VotingIsNotFinished, "Voting is not finished")
}

func ErrStateIsNotSETUP(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, StateIsNotSETUP, "State is not SETUP")
}
func ErrStateIsNotSIGNUP(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, StateIsNotSIGNUP, "State is not SIGNUP")
}
func ErrStateIsNotCOMMITMENT(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, StateIsNotCOMMITMENT, "State is not COMMITMENT")
}
func ErrStateIsNotVOTE(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, StateIsNotVOTE, "State is not VOTE")
}
func ErrStateIsNotFINISHED(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, StateIsNotFINISHED, "State is not StateIsNotFINISHED")
}
