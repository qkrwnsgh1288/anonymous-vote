package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	a := 1
	b := 2
	var c, d int

	a, b, c, d = b, a+3, a, b
	fmt.Println(a, b, c, d)
}

func TestInvmod(t *testing.T) {
	a, err := Invmod(3, 17)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)

	a, err = Invmod(17, 3)
	fmt.Println(a)

	a, err = Invmod(3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}

func TestExpmod(t *testing.T) {
	res, err := Expmod(81792, 73363, 233)
	if err != nil {
		fmt.Println(err)
	}
	var exp uint = 161
	assert.Equal(t, exp, res)

	res, err = Expmod(1000, 1000, 19)
	if err != nil {
		fmt.Println(err)
	}
	exp = 7
	assert.Equal(t, exp, res)

	res, err = Expmod(12, 9, 1)
	if err != nil {
		fmt.Println(err)
	}
	exp = 0
	assert.Equal(t, exp, res)

	res, err = Expmod(111, 123, 53)
	if err != nil {
		fmt.Println(err)
	}
	exp = 35
	assert.Equal(t, exp, res)
}
