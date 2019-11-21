package solution

// 使用两个指针，一个放在左侧开始，一个放在右侧末尾
// 计算并更新当前的最大面积
// 因为面积受制于短的一侧，将短的一侧向长的一侧移动，试图寻找更大的面积
func maxArea(height []int) int {
	n := len(height)
	ans := 0
	l, r := 0, n-1
	for l < r {
		ans = max(ans, (r-l)*min(height[l], height[r]))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

// 递归
func maxArea2(height []int) int {
	n := len(height)
	if n == 2 {
		return min(height[0], height[1])
	}
	ans := maxArea2(height[:n-1])
	for i := 0; i < n-1; i++ {
		ans = max(ans, (n-1-i)*min(height[i], height[n-1]))
	}
	return ans
}

// 解法一
// 循环遍历
// 时间复杂度 O(n2)
func maxArea1(height []int) int {
	n := len(height)
	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, (j-i)*min(height[i], height[j]))
		}
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
