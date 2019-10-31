package solution

// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray/

// Given an integer array nums, find the contiguous subarray (containing at
// least one number) which has the largest sum and return its sum.

// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

func maxSubArray1(nums []int) int {
	n := len(nums)
	result := nums[0]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			tmp := sum(nums[i:j])
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}

func sum(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	return total
}

func maxSubArray(nums []int) int {
	n := len(nums)
	result := nums[0]
	for i := 0; i < n; i++ {
		tmp := 0
		for j := i; j < n; j++ {
			tmp += nums[j]
			if tmp > result {
				result = tmp
			}
		}
	}
	return result
}
