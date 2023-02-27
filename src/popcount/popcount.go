package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var a [8]int

	var result int

	for i := range a {
		result += int(pc[byte(x>>(i*8))])
	}

	return result
}
