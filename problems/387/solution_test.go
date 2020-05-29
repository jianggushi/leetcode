package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstUniqChar_01(t *testing.T) {
	s := "leetcode"
	expected := 0
	ans := firstUniqChar(s)
	assert.Equal(t, expected, ans)
}

func Test_firstUniqChar_02(t *testing.T) {
	s := "loveleetcode"
	expected := 2
	ans := firstUniqChar(s)
	assert.Equal(t, expected, ans)
}

func Test_firstUniqChar_03(t *testing.T) {
	s := "lovelove"
	expected := -1
	ans := firstUniqChar(s)
	assert.Equal(t, expected, ans)
}
