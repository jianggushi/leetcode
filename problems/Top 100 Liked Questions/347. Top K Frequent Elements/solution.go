package solution

import "sort"

func topKFrequent(nums []int, k int) []int {
	// 计数
	m := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}
	// 排序
	a := make([][2]int, 0)
	for k, v := range m {
		a = append(a, [2]int{k, v})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] > a[j][1]
	})
	// 结果
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, a[i][0])
	}
	return ans
}
