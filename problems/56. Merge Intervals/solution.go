package solution

import "sort"

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	n := len(intervals)
	if n == 0 {
		return ans
	}
	// sort by intervals[i][0]
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for i, j := 1, 0; i < n; i++ {
		// overlap
		if intervals[i][0] >= ans[j][0] && intervals[i][0] <= ans[j][1] {
			if intervals[i][1] > ans[j][1] {
				ans[j][1] = intervals[i][1]
			}
		} else {
			ans = append(ans, intervals[i])
			j++
		}
	}
	return ans
}
