// Source : https://leetcode.cn/problems/binary-search/
// Author : jianggushi
// Date   : 2022-11-22

package solution

// 解法一：二分查找算法
//
// nums 元素唯一，升序排列
// 初始化两个指针l在首，r在尾
// 取中间节点 mid=l+(r-l)/2=(l+r)/2
// 如果中间节点nums[mid]==target，找到目标位置，返回下标
// 如果中间节点nums[mid]<target（中间节点小了），要向右边搜索，让l=mid+1
// 如果中间节点nums[mid]>target（中间节点大了），要向左边搜索，让r=mid-1
// 什么时候结束呢？只要l<=r都可以进行，l>r就结束
//
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l       // 取中间元素
		if nums[mid] == target { // 找到目标元素
			return mid
		} else if nums[mid] < target { // 向右搜索
			l = mid + 1
		} else { // 向左搜索
			r = mid - 1
		}
	}
	return -1
}
