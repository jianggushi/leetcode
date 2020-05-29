package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect_01(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2, 2}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_02(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{4, 9}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_03(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{}
	expected := []int{}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_04(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{}
	expected := []int{}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_05(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}
	expected := []int{}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_06(t *testing.T) {
	nums1 := []int{2}
	nums2 := []int{1}
	expected := []int{}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_07(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{1}
	expected := []int{1}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}

func Test_intersect_08(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2, 3, 2}
	expected := []int{2, 2}
	ans := intersect(nums1, nums2)
	assert.ElementsMatch(t, expected, ans)
}
