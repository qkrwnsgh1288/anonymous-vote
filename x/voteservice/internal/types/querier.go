package types

import "strings"

type QueryResTopics []string

func (a QueryResTopics) String() string {
	return strings.Join(a[:], "\n")
}
