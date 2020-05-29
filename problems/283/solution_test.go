package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes_01(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

func Test_moveZeroes_02(t *testing.T) {
	nums := []int{}
	expected := []int{}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

func Test_moveZeroes_03(t *testing.T) {
	nums := []int{1}
	expected := []int{1}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

func Test_moveZeroes_04(t *testing.T) {
	nums := []int{0}
	expected := []int{0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}

func Test_moveZeroes_05(t *testing.T) {
	nums := []int{1, 0}
	expected := []int{1, 0}
	moveZeroes(nums)
	assert.Equal(t, expected, nums)
}
