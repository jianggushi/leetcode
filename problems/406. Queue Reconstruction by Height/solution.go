package solution

// 先排序：按照身高 h 降序 k 升序的规则
// 排序后，前面的人的身高肯定比后面的人高，而且 k 是增加的
// 需要再遍历一遍排序后的结果，把对应的人，按照 k 的值插入到对应的位置
func reconstructQueue(people [][]int) [][]int {
	// 插入排序
	n := len(people)
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if (people[j][0] < people[j+1][0]) ||
				(people[j][0] == people[j+1][0] && people[j][1] > people[j+1][1]) {
				people[j], people[j+1] = people[j+1], people[j]
			} else {
				// 提前退出
				break
			}
		}
	}
	// 根据 k 的值插入到对应的位置
	for i := 0; i < n; i++ {
		if people[i][1] < i {
			for j := i - 1; j >= people[j+1][1]; j-- {
				people[j], people[j+1] = people[j+1], people[j]
			}
		}
	}
	return people
}
