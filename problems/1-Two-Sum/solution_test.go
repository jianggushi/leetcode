package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum_01(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	ans := twoSum(nums, target)
	assert.Equal(t, expected, ans)
}

func Test_twoSum_02(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	ans := twoSum(nums, target)
	assert.Equal(t, expected, ans)
}
