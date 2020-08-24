package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotIndex_01(t *testing.T) {
	nums := []int{1, 7, 3, 6, 5, 6}
	expect := 3
	ans := pivotIndex(nums)
	assert.Equal(t, expect, ans)
}

func Test_pivotIndex_02(t *testing.T) {
	nums := []int{1, 2, 3}
	expect := -1
	ans := pivotIndex(nums)
	assert.Equal(t, expect, ans)
}

func Test_pivotIndex_03(t *testing.T) {
	nums := []int{}
	expect := -1
	ans := pivotIndex(nums)
	assert.Equal(t, expect, ans)
}

func Test_pivotIndex_04(t *testing.T) {
	nums := []int{1}
	expect := 0
	ans := pivotIndex(nums)
	assert.Equal(t, expect, ans)
}

func Test_pivotIndex_05(t *testing.T) {
	nums := []int{1, 1}
	expect := -1
	ans := pivotIndex(nums)
	assert.Equal(t, expect, ans)
}

func Test_pivotIndex_06(t *testing.T) {
	nums := []int{0, 0}
	expect := 0
	ans := pivotIndex(nums)
	assert.Equal(t, expect, ans)
}
