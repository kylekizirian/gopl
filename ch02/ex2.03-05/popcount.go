package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3 - rewrite PopCount to use a loop and compare the performance
func PopCountLoop(x uint64) int {

	count := 0

	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}

	return count
}

// Exercise 2.4 - Write a version of PopCount that shifts over 64-bits,
// counting the rightmost bit each time
func PopCountShift(x uint64) int {

	const one uint64 = 1
	count := 0

	for i := 0; i < 64; i++ {
		count += int(x & one)
		x = x >> 1
	}

	return count
}

// Exercise 2.5 - the expression x&(x-1) clears the rightmost non-zero bit
// of x. Write a version of PopCount that counts bits using this fact.
func PopCountClearRight(x uint64) int {

	count := 0

	for x > 0 {
		x = x & (x - 1)
		count++
	}

	return count
}
