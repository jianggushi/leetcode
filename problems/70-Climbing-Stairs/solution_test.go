package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs_01(t *testing.T) {
	n := 1
	expected := 1
	ans := climbStairs(n)
	assert.Equal(t, expected, ans)
}

func Test_climbStairs_02(t *testing.T) {
	n := 2
	expected := 2
	ans := climbStairs(n)
	assert.Equal(t, expected, ans)
}

func Test_climbStairs_03(t *testing.T) {
	n := 3
	expected := 3
	ans := climbStairs(n)
	assert.Equal(t, expected, ans)
}

func Test_climbStairs_04(t *testing.T) {
	n := 44
	expected := 1134903170
	ans := climbStairs(n)
	assert.Equal(t, expected, ans)
}
