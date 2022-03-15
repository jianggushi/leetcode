package solution

import "math"

// 02 使用 map 空间换时间
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func findRestaurant(list1 []string, list2 []string) []string {
	res := []string{}
	minIndexSum := math.MaxInt

	map1 := make(map[string]int, len(list1))
	for i, item1 := range list1 {
		map1[item1] = i
	}

	for j, item2 := range list2 {
		if i, ok := map1[item2]; ok {
			if i+j < minIndexSum {
				minIndexSum = i + j
				res = []string{item2}
			} else if i+j == minIndexSum {
				res = append(res, item2)
			}
		}
	}
	return res
}

// 01 两次循环遍历
// 时间复杂度 O(n2)
// 空间复杂度 O(n)
func findRestaurant01(list1 []string, list2 []string) []string {
	res := []string{}
	minIndexSum := math.MaxInt

	for i, item1 := range list1 {
		for j, item2 := range list2 {
			if item2 == item1 {
				if i+j < minIndexSum {
					minIndexSum = i + j
					res = []string{item2}
				} else if i+j == minIndexSum {
					res = append(res, item2)
				}
			}
		}
	}
	return res
}
