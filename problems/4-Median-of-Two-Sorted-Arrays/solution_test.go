package solution

import "testing"

func Test_findMedianSortedArrays_01(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	expected := 2.0
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_findMedianSortedArrays_02(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	expected := 2.5
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_findMedianSortedArrays_03(t *testing.T) {
	nums1 := []int{1, 2, 2, 2}
	nums2 := []int{3, 4}
	expected := 2.0
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_findMedianSortedArrays_04(t *testing.T) {
	nums1 := []int{1, 2, 4, 5}
	nums2 := []int{3, 3}
	expected := 3.0
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_findMedianSortedArrays_05(t *testing.T) {
	nums1 := []int{4, 5}
	nums2 := []int{3, 3}
	expected := 3.5
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_findMedianSortedArrays_06(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}
	expected := 1.0
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}

func Test_findMedianSortedArrays_07(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1, 2, 3}
	expected := 2.0
	ans := findMedianSortedArrays(nums1, nums2)
	if ans != expected {
		t.Errorf("expected: %v, ans: %v", expected, ans)
	}
}
