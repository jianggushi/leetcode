package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	expected := true
	ans := isPalindrome(s)
	assert.Equal(t, expected, ans)
}
