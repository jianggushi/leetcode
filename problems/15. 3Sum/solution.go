package solution

import "sort"

func threeSum(nums []int) [][]int {
	// 先排序，从小到大
	sort.Ints(nums)
	ans := [][]int{}
	n := len(nums)
	// nums[i] > 0 后就不再搜索了
	for i := 0; i < n-2 && nums[i] <= 0; {
		// 两个指针，从两边进行搜索
		l, r := i+1, n-1
		for l < r {
			t := nums[i] + nums[l] + nums[r]
			if t == 0 {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				// 找到一个解后，需要排除相同的解
				nl := nums[l]
				for l++; l < r; l++ {
					if nums[l] != nl {
						break
					}
				}
				nr := nums[r]
				for r--; r > l; r-- {
					if nums[r] != nr {
						break
					}
				}
			} else if t < 0 {
				l++
			} else if t > 0 {
				r--
			}
		}
		// 这里也要避免相同解的产生
		ni := nums[i]
		for i++; i < n-2; i++ {
			if nums[i] != ni {
				break
			}
		}
	}
	return ans
}

// 解法一
// 暴力搜索
func threeSum1(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	mmap := make(map[[3]int]bool)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := [3]int{nums[i], nums[j], nums[k]}
					sort.Ints(tmp[:])
					if _, ok := mmap[tmp]; !ok {
						ans = append(ans, tmp[:])
						mmap[tmp] = true
					}
				}
			}
		}
	}
	return ans
}
