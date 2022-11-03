// Source : https://leetcode.cn/problems/find-pivot-index/
// Author : jianggushi
// Date   : 2022-10-26

package solution

// 设 total 为数组元素总和，lsum[i]为位置i左侧元素和
// 那么 rsum[i] = total-lsum[i]-nums[i]
func pivotIndex(nums []int) int {
	n := len(nums)
	total, lsum := 0, 0
	for i := 0; i < n; i++ {
		total += nums[i]
	}
	for i := 0; i < n; i++ {
		if total-lsum-nums[i] == lsum {
			return i
		}
		lsum += nums[i]
	}
	return -1
}

// 其中特殊处理了lsum[0]的情况，但也可以换个位置，去掉if语句
func pivotIndex3(nums []int) int {
	n := len(nums)
	lsum, rsum := 0, 0
	for i := 0; i < n; i++ {
		rsum += nums[i]
	}
	for i := 0; i < n; i++ {
		rsum -= nums[i]
		if lsum == rsum {
			return i
		}
		lsum += nums[i] // 提前计算下次迭代的lsum
	}
	return -1
}

// 解法二：前缀和
//
// 在暴力搜索中每次都要计算位置i左侧和右侧元素的和，导致时间复杂度比较高
// 在遍历过程中左侧元素逐渐增加，右侧元素逐渐减少
// 在位置i时
// lsum[i] = nums[0] + ... + nums[i-1]
// rsum[i] = nums[i+1] + ... + nums[n-1]
// 在位置i+1时
// lsum[i+1] = nums[0] + ... + nums[i]
// rsum[i+1] = nums[i+2] + ... + nums[n-1]
// 可以得到
// lsum[i+1] = lsum[i] + nums[i]
// rsum[i+1] = rsum[i] - nums[i+1]
//
// 特殊位置 0 套用上面的公式
// lsum[0] = lsum[-1] + nums[-1]
// rsum[0] = rsum[-1] - nums[0]
//
// 其中需要位置 -1，但这个位置并不真实存在，为了减少特殊情况我们可以假设
// lsum[-1] = 0
// rsum[-1] = nums[0] + ... + nums[n-1] 为全部元素的和
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func pivotIndex2(nums []int) int {
	n := len(nums)
	lsum, rsum := 0, 0
	for i := 0; i < n; i++ {
		rsum += nums[i]
	}
	for i := 0; i < n; i++ {
		if i > 0 {
			lsum += nums[i-1]
		}
		rsum -= nums[i]
		if lsum == rsum {
			return i
		}
	}
	return -1
}

// 解法一：暴力搜索
//
// 遍历数组中的每个位置i，以位置i为中心点
// 对位置i左边和右边的元素分别进行求和
// 设lsum[i] = nums[0] + ... + nums [i-1] 为位置i左边元素和
// 设rsum[i] = nums[i+1] + ... + nums[n-1] 为位置I右边元素和
// 判断 lsum[i] 和 rsum[i] 是否相等
// 如果 lsum[i] == rsum[i]，则位置i就是所求数组的中心下标
// 如果 lsum[i] != rsum[i]，尝试下一个位置 i+1，直到数组最后一个位置
// 如果都不满足，那么返回 -1
//
// 需要注意首尾两个特殊的位置：
// 位置0，左边没有元素，lsum[0] = 0
// 位置n-1，右边没有元素，rsum[n-1] = 0
//
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
func pivotIndex1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if sum(nums[0:i]) == sum(nums[i+1:n]) {
			return i
		}
	}
	return -1
}

func sum(nums []int) int {
	ans := 0
	for i := range nums {
		ans += nums[i]
	}
	return ans
}
