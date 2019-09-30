package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/qkrwnsgh1288/anonymous-vote/x/voteservice/internal/types"
	"net/http"
	"strings"
)

type makeAgendaReq struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	AgendaProposer string       `json:"agenda_proposer"`
	AgendaTopic    string       `json:"agenda_topic"`
	AgendaContent  string       `json:"agenda_content"`
	WhiteList      []string     `json:"whitelist"`
}

func makeAgendaHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req makeAgendaReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.AgendaProposer)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgMakeAgenda(addr, req.AgendaTopic, req.AgendaContent, req.WhiteList)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type voteAgendaReq struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	AgendaTopic string       `json:"agenda_topic"`
	VoteAddr    string       `json:"vote_addr"`
	YesOrNo     string       `json:"yes_or_no"`
}

func voteAgendaHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req voteAgendaReq

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		addr, err := sdk.AccAddressFromBech32(req.VoteAddr)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		answer := strings.TrimSpace(strings.ToLower(req.YesOrNo))

		msg := types.NewMsgVoteAgenda(addr, req.AgendaTopic, answer)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
