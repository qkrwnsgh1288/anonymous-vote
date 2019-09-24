package types

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestYesOrNo(t *testing.T) {
	case1 := "Yes"
	case2 := " yEs "
	case3 := "    No "

	assert.Equal(t, "yes", strings.TrimSpace(strings.ToLower(case1)))
	assert.Equal(t, "yes", strings.TrimSpace(strings.ToLower(case2)))
	assert.Equal(t, "no", strings.TrimSpace(strings.ToLower(case3)))
}
