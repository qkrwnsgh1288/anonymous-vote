package types

import "strings"

// Query Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

func (r QueryResResolve) String() string {
	return r.Value
}

// Query Result Payload for a names query
type QueryResNames []string

func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}

type QueryResAgendas []string

func (a QueryResAgendas) String() string {
	return strings.Join(a[:], "\n")
}
