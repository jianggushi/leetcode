package solution

// 思路二
// 4  = 0 1 0 0
// 14 = 1 1 1 0
// 2  = 0 0 1 0
// 从后往前看,
// 第一个 bit 位都是 0, 汉明距离 0
// 第二个 bit 位 1 个 bit0, 2 个 bit1, 汉明距离 2
// 第三个 bit 位 1 个 bit0, 2 个 bit1, 汉明距离 2
// 第四个 bit 位 2 个 bit0, 1 个 bit1, 汉明距离 2
// 总结, 同样的 bit 位, 汉明距离 = count(bit0) * count(bit1)
// 时间复杂度 O(n)
func totalHammingDistance(nums []int) int {
	total := 0
	n := len(nums)
	for i := 0; i < 32; i++ {
		bit1 := 0
		for j := 0; j < n; j++ {
			if nums[j]&1 == 1 {
				bit1++
			}
			nums[j] >>= 1
		}
		total += bit1 * (n - bit1)
	}
	return total
}

// 思路一
// 两层循环遍历所有的两两组合，分别计算两个数的 Hamming Distance，并求和
// 时间复杂度 O(n2)
func totalHammingDistance1(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			total += hammingDistance(nums[i], nums[j])
		}
	}
	return total
}

func hammingDistance(x int, y int) int {
	n := x ^ y
	bits := 0
	for n != 0 {
		bits++
		n &= (n - 1)
	}
	return bits
}
