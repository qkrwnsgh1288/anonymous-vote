package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/common"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types"
	"github.com/spf13/cobra"
	"strings"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Voteservice transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	nameserviceTxCmd.AddCommand(client.PostCommands(
		GetCmdMakeAgenda(cdc),
		GetCmdRegisterByVoter(cdc),
		GetCmdVoteAgenda(cdc),
	)...)

	return nameserviceTxCmd
}

func GetCmdMakeAgenda(cdc *codec.Codec) *cobra.Command {
	var whiteList []string
	c := &cobra.Command{
		Use:   "make-agenda [topic] [content]",
		Short: "make agenda",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgMakeAgenda(cliCtx.GetFromAddress(), args[0], args[1], whiteList)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	c.Flags().StringSliceVarP(&whiteList, "whitelist", "w", []string{}, "")

	return c
}
func GetCmdRegisterByVoter(cdc *codec.Codec) *cobra.Command {
	c := &cobra.Command{
		Use:   "register-by-voter [topic] [file_path]",
		Short: "register my zk-info",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			zkSlice, err := common.ReadZkInfoFromFile(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterByVoter(cliCtx.GetFromAddress(), args[0], zkSlice)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return c
}
func GetCmdVoteAgenda(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vote-agenda [topic] [yes or no]",
		Short: "vote agenda about topic",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			answer := strings.TrimSpace(strings.ToLower(args[1]))

			msg := types.NewMsgVoteAgenda(cliCtx.GetFromAddress(), args[0], answer)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
