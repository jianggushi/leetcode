package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_myAtoi_01(t *testing.T) {
	s := "42"
	expected := 42
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_02(t *testing.T) {
	s := "   -42"
	expected := -42
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_03(t *testing.T) {
	s := "4193 with words"
	expected := 4193
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_04(t *testing.T) {
	s := "words and 987"
	expected := 0
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_05(t *testing.T) {
	s := "-91283472332"
	expected := -2147483648
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_06(t *testing.T) {
	s := "-2147483647"
	expected := -2147483647
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_07(t *testing.T) {
	s := "-2147483648"
	expected := -2147483648
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_08(t *testing.T) {
	s := "-2147483649"
	expected := -2147483648
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}
func Test_myAtoi_09(t *testing.T) {
	s := "2147483647"
	expected := 2147483647
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_10(t *testing.T) {
	s := "2147483648"
	expected := 2147483647
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_11(t *testing.T) {
	s := "2147483649"
	expected := 2147483647
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_12(t *testing.T) {
	s := "91283472332"
	expected := 2147483647
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_13(t *testing.T) {
	s := ""
	expected := 0
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_14(t *testing.T) {
	s := "-"
	expected := 0
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_15(t *testing.T) {
	s := "+1"
	expected := 1
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_16(t *testing.T) {
	s := "-+1"
	expected := 0
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}

func Test_myAtoi_17(t *testing.T) {
	s := "+-1"
	expected := 0
	ans := myAtoi(s)
	assert.Equal(t, expected, ans)
}
