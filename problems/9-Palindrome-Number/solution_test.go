package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome_01(t *testing.T) {
	x := 121
	expected := true
	ans := isPalindrome(x)
	assert.Equal(t, expected, ans)
}

func Test_isPalindrome_02(t *testing.T) {
	x := -121
	expected := false
	ans := isPalindrome(x)
	assert.Equal(t, expected, ans)
}

func Test_isPalindrome_03(t *testing.T) {
	x := 10
	expected := false
	ans := isPalindrome(x)
	assert.Equal(t, expected, ans)
}
