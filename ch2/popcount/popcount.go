package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		// 0 - 0 0000 0000
		// 1 - 1 0000 0001
		// 2 - 1 0000 0010
		// 3 - 2 0000 0011
		// 4 - 1 0000 0100
		// 5 - 2 0000 0101
		// 6 - 2 0000 0110
		// 7 - 3 0000 0111
		// 8 - 2 0000 1000
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))]) +
		int(pc[byte(x>>(1*8))]) +
		int(pc[byte(x>>(2*8))]) +
		int(pc[byte(x>>(3*8))]) +
		int(pc[byte(x>>(4*8))]) +
		int(pc[byte(x>>(5*8))]) +
		int(pc[byte(x>>(6*8))]) +
		int(pc[byte(x>>(7*8))])
}
