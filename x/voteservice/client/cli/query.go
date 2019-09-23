package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/qkrwnsgh1288/vote-dapp/x/voteservice/internal/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the voteservice module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(client.GetCommands(
		GetCmdAgenda(storeKey, cdc),
		GetCmdTopics(storeKey, cdc),
	)...)
	return nameserviceQueryCmd
}

func GetCmdAgenda(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "agenda [topic]",
		Short: "Show agenda details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			topic := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/agenda/%s", queryRoute, topic), nil)
			if err != nil {
				fmt.Printf("does not have topic in agendas - %s \n", topic)
				return nil
			}

			var out types.Agenda
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
func GetCmdTopics(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "topics",
		Short: "topics",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/topics", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query topics\n")
				return nil
			}

			var out types.QueryResTopics
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
