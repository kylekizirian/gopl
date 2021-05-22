// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join.

// Run this with a few arguments to test the difference. On my machine with
// 5 arguments, the inefficient method takes about 440ns and the efficient
// method takes about 140ns.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const NumIterations = 10_000_000

func main() {

	var inefficientTotalTime int64 = 0
	var efficientTotalTime int64 = 0

	for i := 0; i < NumIterations; i++ {
		// time the string building operation by looping over args
		start := time.Now()
		s, sep := "", ""

		for _, arg := range os.Args {
			s += sep + arg
			sep = " "
		}

		inefficientTotalTime += time.Since(start).Nanoseconds()
	}

	for i := 0; i < NumIterations; i++ {
		// time the string building operation using strings.Join
		start := time.Now()
		s := strings.Join(os.Args, " ")
		_ = s
		efficientTotalTime += time.Since(start).Nanoseconds()
	}

	avgInefficientTime := float64(inefficientTotalTime) / NumIterations
	avgEfficientTime := float64(efficientTotalTime) / NumIterations

	fmt.Printf("Average inefficient time: %f ns\n", avgInefficientTime)
	fmt.Printf("Average efficient time: %f ns\n", avgEfficientTime)

}
