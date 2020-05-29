package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne_01(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	ans := plusOne(nums)
	assert.Equal(t, expected, ans)
}

func Test_plusOne_02(t *testing.T) {
	nums := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}
	ans := plusOne(nums)
	assert.Equal(t, expected, ans)
}

func Test_plusOne_03(t *testing.T) {
	nums := []int{1, 9}
	expected := []int{2, 0}
	ans := plusOne(nums)
	assert.Equal(t, expected, ans)
}

func Test_plusOne_04(t *testing.T) {
	nums := []int{1, 9, 9}
	expected := []int{2, 0, 0}
	ans := plusOne(nums)
	assert.Equal(t, expected, ans)
}

func Test_plusOne_05(t *testing.T) {
	nums := []int{9, 9, 9}
	expected := []int{1, 0, 0, 0}
	ans := plusOne(nums)
	assert.Equal(t, expected, ans)
}
