package popcount

import (
	"testing"
	"testing/quick"
)

func TestPopCount(t *testing.T) {

	allSame := func(x uint64) bool {
		if PopCount(x) != PopCountLoop(x) {
			return false
		}

		if PopCount(x) != PopCountShift(x) {
			return false
		}

		if PopCount(x) != PopCountClearRight(x) {
			return false
		}

		return true
	}

	if err := quick.Check(allSame, nil); err != nil {
		t.Error(err)
	}

}

func benchmarkPopCount(i uint64, b *testing.B, popc func(uint64) int) {

	for n := 0; n < b.N; n++ {
		popc(i)
	}

}

func BenchmarkPopCount1(b *testing.B)   { benchmarkPopCount(1, b, PopCount) }
func BenchmarkPopCountMax(b *testing.B) { benchmarkPopCount(^uint64(0), b, PopCount) }

func BenchmarkPopCountLoop1(b *testing.B)   { benchmarkPopCount(1, b, PopCountLoop) }
func BenchmarkPopCountLoopMax(b *testing.B) { benchmarkPopCount(^uint64(0), b, PopCountLoop) }

func BenchmarkPopCountShift1(b *testing.B)   { benchmarkPopCount(1, b, PopCountShift) }
func BenchmarkPopCountShiftMax(b *testing.B) { benchmarkPopCount(^uint64(0), b, PopCountShift) }

func BenchmarkPopCountClearRight1(b *testing.B) { benchmarkPopCount(1, b, PopCountClearRight) }
func BenchmarkPopCounClearRighttMax(b *testing.B) {
	benchmarkPopCount(^uint64(0), b, PopCountClearRight)
}
