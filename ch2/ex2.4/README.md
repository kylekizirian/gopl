# Exercise 2.3, 2.4, and 2.5

This exercise has us benchmark a few implementations of `popcount`, the
results on my machine are below. `popcount` takes a `uint64` and each
implementation is run with the value `1` and `2^64 - 1`.

```
goos: darwin
goarch: amd64
pkg: popcount
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkPopCount1-8              	1000000000	         0.2642 ns/op
BenchmarkPopCountMax-8            	1000000000	         0.2684 ns/op
BenchmarkPopCountLoop1-8          	284938018	         4.207 ns/op
BenchmarkPopCountLoopMax-8        	199043202	         5.882 ns/op
BenchmarkPopCountShift1-8         	62907314	        17.80 ns/op
BenchmarkPopCountShiftMax-8       	63476053	        17.80 ns/op
BenchmarkPopCountClearRight1-8    	1000000000	         0.7868 ns/op
BenchmarkPopCounClearRightMax-8   	40695819	        28.27 ns/op
PASS
ok  	popcount	8.492s
```

