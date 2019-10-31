package solution

// Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
// You may assume that the array is non-empty and the majority element always exist in the array.

// Example 1:
// Input: [3,2,3]
// Output: 3

// Example 2:
// Input: [2,2,1,1,1,2,2]
// Output: 2

// 使用 map 统计数组中元素出现次数
func majorityElement(nums []int) int {
	n := len(nums)
	mapping := make(map[int]int)

	for i := 0; i < n; i++ {
		if _, ok := mapping[nums[i]]; ok {
			mapping[nums[i]]++
		} else {
			mapping[nums[i]] = 1
		}
	}
	thr := n >> 1
	for num, value := range mapping {
		if value > thr {
			return num
		}
	}
	return 0
}
