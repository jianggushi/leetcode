package solution

func isHappy(n int) bool {
	mmap := make(map[int]bool)
	for {
		if n == 1 {
			return true
		}
		if _, ok := mmap[n]; ok {
			return false
		}
		mmap[n] = true
		m := 0
		for n != 0 {
			m += (n % 10) * (n % 10)
			n /= 10
		}
		n = m
	}
}
