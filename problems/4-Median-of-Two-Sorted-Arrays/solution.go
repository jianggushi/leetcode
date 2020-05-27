package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	i, j := 0, 0
	for i+j < (m+n-1)/2 {
		if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	// 奇数个元素
	if (m+n)%2 == 1 {
		return float64(min(nums1[i], nums2[j]))
	} else {
		if nums1[i] < nums2[j] && i < m-1 {
			return float64(nums1[i]+min(nums1[i+1], nums2[j])) / 2.0
		} else if nums1[i] > nums2[j] && j < n-1 {
			return float64(nums2[j]+min(nums2[j+1], nums1[i])) / 2.0
		} else {
			return float64(nums1[i]+nums2[j]) / 2.0
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 时间复杂度 O(m+n)
// 空间复杂度 O(m+n)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	nums := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
		k++
	}
	for i < m {
		nums[k] = nums1[i]
		i++
		k++
	}
	for j < n {
		nums[k] = nums2[j]
		j++
		k++
	}
	// 奇数个元素
	if (m+n)%2 == 1 {
		return float64(nums[(m+n)/2])
	} else {
		return float64(nums[(m+n-1)/2]+nums[(m+n)/2]) / 2.0
	}
}
