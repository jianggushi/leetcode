package solution

// 方法一：
// 一个重要的条件 1 <= a[i] <= n，n = len(a)
// 构造 hash table，key = [1, n], value = key 出现的次数
// 遍历数组 a，计算每个 key 出现的次数
// 最后遍历 hash table，key 出现次数为 0 的就是缺失的 number

// 方法二：
// 调整 a[i] 到它应该在的下标位置

// 方法三：
// 一个萝卜一个坑，可以用数组作为 hash table
// 数组下标是 key，数组的值是出现的次数

// 方法四：
// 把每个数组元素对应的下标位置加上 n
// 那么，那些重复出现的元素，对应的下标位置会加上多次 n
// 那些缺失的元素，对应的下标位置不会加上 n
// 最后遍历数组，找到数组中 <= n 的对应的下标就是缺失的元素

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		// 元素可能会重复出现，nums[i]可能已经加了 n，需要取余数
		nums[(nums[i]-1)%n] += n
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			result = append(result, i+1)
		}
	}
	return result
}

func findDisappearedNumbers4(nums []int) []int {
	n := len(nums)
	numMap := make([]int, n)
	for i := 0; i < n; i++ {
		numMap[nums[i]-1] = 1
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if numMap[i] == 0 {
			result = append(result, i+1)
		}
	}
	return result
}

// 调整位置
func findDisappearedNumbers3(nums []int) []int {
	n := len(nums)

	for i := 0; i < n; {
		if nums[i] != i+1 && nums[i] != 0 {
			if nums[nums[i]-1] != nums[i] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			} else {
				nums[i] = 0
			}
		} else {
			i++
		}
	}
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			result = append(result, i+1)
		}
	}
	return result
}

// hash table
// 构造 hash table 和计算每个 key 出现的次数，两个循环合并
func findDisappearedNumbers2(nums []int) []int {
	n := len(nums)
	numMap := make(map[int]int, 0)

	for i := 0; i < n; i++ {
		if _, ok := numMap[i+1]; !ok {
			numMap[i+1] = 0
		}
		if _, ok := numMap[nums[i]]; ok {
			numMap[nums[i]]++
		} else {
			numMap[nums[i]] = 1
		}
	}

	result := make([]int, 0)
	for num, count := range numMap {
		if count == 0 {
			result = append(result, num)
		}
	}
	return result
}

// hash table
func findDisappearedNumbers1(nums []int) []int {
	n := len(nums)
	numMap := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		numMap[i+1] = 0
	}

	for i := 0; i < n; i++ {
		numMap[nums[i]]++
	}

	result := make([]int, 0)
	for num, count := range numMap {
		if count == 0 {
			result = append(result, num)
		}
	}
	return result
}
