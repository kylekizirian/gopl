package main

import "fmt"

// Count set bits by checking lower bit and shifting each time
func PopCount(x uint64) int {

	const one uint64 = 1
	count := 0

	for i := 0; i < 64; i++ {
		count += int(x & one)
		x = x >> 1
	}

	return count
}

func main() {
	fmt.Println(PopCount(7))
}
