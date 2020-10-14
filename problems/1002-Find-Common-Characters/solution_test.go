package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_commonChars_01(t *testing.T) {
	A := []string{"bella", "label", "roller"}
	want := []string{"e", "l", "l"}
	got := commonChars(A)
	assert.Equal(t, want, got)
}

func Test_commonChars_02(t *testing.T) {
	A := []string{"cool", "lock", "cook"}
	want := []string{"c", "o"}
	got := commonChars(A)
	assert.Equal(t, want, got)
}

func Test_commonChars_03(t *testing.T) {
	A := []string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}
	want := []string{}
	got := commonChars(A)
	assert.Equal(t, want, got)
}
