package solution

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// Note:
// Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

// Example 1:
// Input: [2,2,1]
// Output: 1

// Example 2:
// Input: [4,1,2,1,2]
// Output: 4
func singleNumber(nums []int) int {
	mapping := make(map[int]bool)
	n := len(nums)

	for i := 0; i < n; i++ {
		if _, ok := mapping[nums[i]]; ok {
			delete(mapping, nums[i])
		} else {
			mapping[nums[i]] = true
		}
	}

	result := 0
	for num := range mapping {
		result = num
		break
	}
	return result
}
