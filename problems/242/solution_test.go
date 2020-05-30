package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram_01(t *testing.T) {
	ss := "anagram"
	tt := "nagaram"
	expected := true
	ans := isAnagram(ss, tt)
	assert.Equal(t, expected, ans)
}

func Test_isAnagram_02(t *testing.T) {
	ss := "rat"
	tt := "cat"
	expected := false
	ans := isAnagram(ss, tt)
	assert.Equal(t, expected, ans)
}

func Test_isAnagram_03(t *testing.T) {
	ss := "ra"
	tt := "cat"
	expected := false
	ans := isAnagram(ss, tt)
	assert.Equal(t, expected, ans)
}

func Test_isAnagram_04(t *testing.T) {
	ss := "aabb"
	tt := "bba"
	expected := false
	ans := isAnagram(ss, tt)
	assert.Equal(t, expected, ans)
}
