package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLengthOfLCIS_01(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7}
	expected := 3
	ans := findLengthOfLCIS(nums)
	assert.Equal(t, expected, ans)
}

func Test_findLengthOfLCIS_02(t *testing.T) {
	nums := []int{2, 2, 2, 2, 2}
	expected := 1
	ans := findLengthOfLCIS(nums)
	assert.Equal(t, expected, ans)
}

func Test_findLengthOfLCIS_03(t *testing.T) {
	nums := []int{2}
	expected := 1
	ans := findLengthOfLCIS(nums)
	assert.Equal(t, expected, ans)
}

func Test_findLengthOfLCIS_04(t *testing.T) {
	nums := []int{2, 1}
	expected := 1
	ans := findLengthOfLCIS(nums)
	assert.Equal(t, expected, ans)
}

func Test_findLengthOfLCIS_05(t *testing.T) {
	nums := []int{1, 3, 5, 4, 7, 8, 9, 2, 4, 5}
	expected := 4
	ans := findLengthOfLCIS(nums)
	assert.Equal(t, expected, ans)
}
