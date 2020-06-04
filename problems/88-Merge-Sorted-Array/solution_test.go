package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_merge_01(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	expected := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)
	assert.Equal(t, expected, nums1)
}
