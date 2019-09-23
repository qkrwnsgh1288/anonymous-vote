package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	restName = "vote"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	// query
	r.HandleFunc(fmt.Sprintf("/%s/agenda/{%s}", storeName, restName), agendaHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/agenda/topics", storeName), topicsHandler(cliCtx, storeName)).Methods("GET")

	// tx
	r.HandleFunc(fmt.Sprintf("/%s/agenda", storeName), makeAgendaHandler(cliCtx)).Methods("POST")
}
