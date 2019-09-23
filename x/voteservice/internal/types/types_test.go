package types

import (
	"fmt"
	"testing"
)

func TestNewWhois(t *testing.T) {
	a := NewWhois()
	fmt.Println(a)
}
func TestNewAgenda(t *testing.T) {
	a := NewAgenda()
	a.Voters = append(a.Voters, "A")
	fmt.Println(a.Voters)

	b := make([]string, 1)
	b = append(b, "AA")
	fmt.Println(b)
}
