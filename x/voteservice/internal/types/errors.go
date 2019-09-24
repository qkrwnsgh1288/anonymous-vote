package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeNameDoesNotExist    sdk.CodeType = 101
	AgendaTopicAlreadyExist sdk.CodeType = 102
	AgendaTopicDoesNotExist sdk.CodeType = 103
	InvalidAnswer           sdk.CodeType = 104
)

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
