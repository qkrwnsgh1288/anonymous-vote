package voteservice

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	AgendaRecords []Agenda `json:"agenda_records"`
}

func NewGenesisState(agendaRecords []Agenda) GenesisState {
	return GenesisState{AgendaRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.AgendaRecords {
		if record.AgendaProposer == nil {
			return fmt.Errorf("Invalid AgendaRecord: AgendaProposer: %s. Error: Missing AgendaProposer", record.AgendaProposer)
		}
		if record.AgendaTopic == "" {
			return fmt.Errorf("Invalid AgendaRecord: AgendaTopic: %s. Error: Missing AgendaTopic", record.AgendaTopic)
		}
		if record.AgendaContent == "" {
			return fmt.Errorf("Invalid AgendaRecord: AgendaContent: %s. Error: Missing AgendaContent", record.AgendaContent)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		AgendaRecords: []Agenda{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.AgendaRecords {
		keeper.SetAgenda(ctx, record.AgendaTopic, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Agenda
	iterator := k.GetTopicsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		topic := string(iterator.Key())
		agenda := k.GetAgenda(ctx, topic)
		records = append(records, agenda)
	}
	return GenesisState{AgendaRecords: records}
}
