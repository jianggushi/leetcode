package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords_01(t *testing.T) {
	s := "the sky is blue"
	expected := "blue is sky the"
	ans := reverseWords(s)
	assert.Equal(t, expected, ans)
}

func Test_reverseWords_02(t *testing.T) {
	s := "  hello world!  "
	expected := "world! hello"
	ans := reverseWords(s)
	assert.Equal(t, expected, ans)
}

func Test_reverseWords_03(t *testing.T) {
	s := "a good   example"
	expected := "example good a"
	ans := reverseWords(s)
	assert.Equal(t, expected, ans)
}
