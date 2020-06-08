package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestConsecutive_01(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	expected := 4
	ans := longestConsecutive(nums)
	assert.Equal(t, expected, ans)
}

func Test_longestConsecutive_02(t *testing.T) {
	nums := []int{}
	expected := 0
	ans := longestConsecutive(nums)
	assert.Equal(t, expected, ans)
}

func Test_longestConsecutive_03(t *testing.T) {
	nums := []int{1}
	expected := 1
	ans := longestConsecutive(nums)
	assert.Equal(t, expected, ans)
}

func Test_longestConsecutive_04(t *testing.T) {
	nums := []int{4, 3, 2, 1, 5}
	expected := 5
	ans := longestConsecutive(nums)
	assert.Equal(t, expected, ans)
}

func Test_longestConsecutive_05(t *testing.T) {
	nums := []int{1, 2, 0, 1}
	expected := 3
	ans := longestConsecutive(nums)
	assert.Equal(t, expected, ans)
}
