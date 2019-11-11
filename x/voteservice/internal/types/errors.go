package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	// basic (100 ~ 200)
	CodeNameDoesNotExist    sdk.CodeType = 101
	AgendaTopicAlreadyExist sdk.CodeType = 102
	AgendaTopicDoesNotExist sdk.CodeType = 103
	InvalidAnswer           sdk.CodeType = 104

	// about zk (201 ~ 299)
	InvalidPubkeyInCreateZKP sdk.CodeType = 201
	InvalidVerifyZKP         sdk.CodeType = 202
	DoesNotRegisterAddress   sdk.CodeType = 203
)

// basic (100 ~ 200)
func ErrNameDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNameDoesNotExist, "Name does not exist")
}
func ErrAgendaTopicAlreadyExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AgendaTopicAlreadyExist, "AgendaTopic already exist")
}
func ErrAgendaTopicDoesNotExist(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, AgendaTopicDoesNotExist, "AgendaTopic does not exist")
}
func ErrInvalidAnswer(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, InvalidAnswer, "Answer is not valid. It should be yes or no")
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
