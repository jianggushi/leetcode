package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber_01(t *testing.T) {
	nums := []int{2, 2, 1}
	expected := 1
	ans := singleNumber(nums)
	assert.Equal(t, expected, ans)
}

func Test_singleNumber_02(t *testing.T) {
	nums := []int{4, 1, 2, 1, 2}
	expected := 4
	ans := singleNumber(nums)
	assert.Equal(t, expected, ans)
}
