package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countAndSay_01(t *testing.T) {
	n := 1
	expected := "1"
	ans := countAndSay(n)
	assert.Equal(t, expected, ans)
}

func Test_countAndSay_02(t *testing.T) {
	n := 2
	expected := "11"
	ans := countAndSay(n)
	assert.Equal(t, expected, ans)
}

func Test_countAndSay_03(t *testing.T) {
	n := 3
	expected := "21"
	ans := countAndSay(n)
	assert.Equal(t, expected, ans)
}

func Test_countAndSay_04(t *testing.T) {
	n := 4
	expected := "1211"
	ans := countAndSay(n)
	assert.Equal(t, expected, ans)
}

func Test_countAndSay_05(t *testing.T) {
	n := 5
	expected := "111221"
	ans := countAndSay(n)
	assert.Equal(t, expected, ans)
}
