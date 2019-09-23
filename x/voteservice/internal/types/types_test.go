package types

import (
	"fmt"
	"testing"
)

func TestNewAgenda(t *testing.T) {
	a := NewAgenda()
	a.Voters = append(a.Voters, "A", "B")
	fmt.Println(a)
}
