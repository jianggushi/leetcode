package solution

/*
 * 1. Two Sum
 *
 * https://leetcode.com/problems/two-sum/
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 */

// 暴力搜索
// 遍历数组中的每一个元素 x，并查找数组中是否存在一个元素 y，满足 x + y = target
// 时间复杂度 O(n2)
// 空间复杂度 O(1)
func twoSum1(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 两遍哈希表
// 在查找数组中的元素 y 时，暴力搜索是采用顺序查找，时间复杂度高
// hash map 查找的时间复杂度 O(1)，可以把数组元素存到 hash map 中，数组元素的值作为 key，数组元素的下标作为 value
// 遍历数组的中每一个元素 x，并在 hash map 中查找是否存在元素 y = target - x
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func twoSum2(nums []int, target int) []int {
	n := len(nums)
	records := make(map[int]int, n)
	for i := 0; i < n; i++ {
		records[nums[i]] = i
	}

	for i := 0; i < n; i++ {
		// NOTE: 不能使用同一个数组元素
		if j, ok := records[target-nums[i]]; ok && j != i {
			return []int{i, j}
		}
	}
	return nil
}

// 一遍哈希表
func twoSum(nums []int, target int) []int {
	n := len(nums)
	records := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if j, ok := records[target-nums[i]]; ok {
			return []int{j, i}
		}
		records[nums[i]] = i
	}
	return nil
}
