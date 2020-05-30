package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr_01(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	expected := 2
	ans := strStr(haystack, needle)
	assert.Equal(t, expected, ans)
}

func Test_strStr_02(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	expected := -1
	ans := strStr(haystack, needle)
	assert.Equal(t, expected, ans)
}

func Test_strStr_03(t *testing.T) {
	haystack := "aaaaa"
	needle := ""
	expected := 0
	ans := strStr(haystack, needle)
	assert.Equal(t, expected, ans)
}

func Test_strStr_04(t *testing.T) {
	haystack := "aaaaa"
	needle := "aaa"
	expected := 0
	ans := strStr(haystack, needle)
	assert.Equal(t, expected, ans)
}

func Test_strStr_05(t *testing.T) {
	haystack := "aaaab"
	needle := "b"
	expected := 4
	ans := strStr(haystack, needle)
	assert.Equal(t, expected, ans)
}

func Test_strStr_06(t *testing.T) {
	haystack := ""
	needle := ""
	expected := 0
	ans := strStr(haystack, needle)
	assert.Equal(t, expected, ans)
}
