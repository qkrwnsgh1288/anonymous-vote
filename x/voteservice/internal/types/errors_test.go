package types

import (
	"fmt"
	"testing"
)

func TestErrNameDoesNotExist(t *testing.T) {
	a := ErrNameDoesNotExist("test")
	fmt.Println(a)
}
