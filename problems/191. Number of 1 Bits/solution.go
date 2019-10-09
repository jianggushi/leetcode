package solution

// 位操作小技巧
func hammingWeight(num uint32) int {
	bits := 0
	for num != 0 {
		bits++
		num &= (num - 1)
	}
	return bits
}

// 移位
func hammingWeight1(num uint32) int {
	bits := 0
	for j := num; j != 0; j >>= 1 {
		if j&0x01 == 1 {
			bits++
		}
	}
	return bits
}
