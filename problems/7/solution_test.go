package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse_01(t *testing.T) {
	x := 123
	expected := 321
	ans := reverse(x)
	assert.Equal(t, expected, ans)
}

func Test_reverse_02(t *testing.T) {
	x := -123
	expected := -321
	ans := reverse(x)
	assert.Equal(t, expected, ans)
}

func Test_reverse_03(t *testing.T) {
	x := 120
	expected := 21
	ans := reverse(x)
	assert.Equal(t, expected, ans)
}

func Test_reverse_04(t *testing.T) {
	x := 2147483648
	expected := 0
	ans := reverse(x)
	assert.Equal(t, expected, ans)
}

func Test_reverse_05(t *testing.T) {
	x := 1534236469
	expected := 0
	ans := reverse(x)
	assert.Equal(t, expected, ans)
}
