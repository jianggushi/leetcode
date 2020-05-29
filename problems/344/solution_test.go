package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString_01(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(s)
	assert.Equal(t, expected, s)
}

func Test_reverseString_02(t *testing.T) {
	s := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expected := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(s)
	assert.Equal(t, expected, s)
}

func Test_reverseString_03(t *testing.T) {
	s := []byte{'H'}
	expected := []byte{'H'}
	reverseString(s)
	assert.Equal(t, expected, s)
}

func Test_reverseString_04(t *testing.T) {
	s := []byte{}
	expected := []byte{}
	reverseString(s)
	assert.Equal(t, expected, s)
}
