package solution

// 思路：
// 求解括号的组合方式，分治策略
// 当 n = 1 时，只有一种组合方式 [()]
// 当 n = 2 时，有两种组合方式 [()(), (())]，这个结果可以从 n = 1 时扩展过来，新增的括号在原来有三个可以插入的位置
// 插入后组成 [()(), (()), ()()]，去除重复的就是 [()(), (())]
// 当 n = 3 时，比上面也多一个括号，新增的括号也可以插入原来的结果中组成
// ()() -> [()()(), (())(), ()()(), ()(()), ()()()]
// (()) -> [()(()), (()()), ((())), (()()), (())()]
// 去除重复的就是 [()()(), (())(), ()(()), (()()), ((()))]

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	// a, b 两个 map 交替使用
	a := make(map[string]bool)
	b := make(map[string]bool)
	a["()"] = true
	for i := 2; i <= n; i++ {
		for p := range a {
			for j := 0; j <= len(p); j++ {
				tmp := p[0:j] + "()" + p[j:len(p)]
				b[tmp] = true
			}
		}
		// 遍历后，b -> a，同时清空 b
		a = b
		b = make(map[string]bool)
	}
	// 生成结果
	ans := make([]string, 0)
	for p := range a {
		ans = append(ans, p)
	}
	return ans
}
