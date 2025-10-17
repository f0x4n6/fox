# Benchmark

Core data structure benchmark:
```
goos: darwin
goarch: arm64
pkg: github.com/cuhsat/fox/internal/pkg/types/smap
cpu: Apple M4
BenchmarkMap-10       	     770	   1579411 ns/op	10259438 B/op	   31131 allocs/op
BenchmarkRender-10    	     344	   3512507 ns/op	 6404429 B/op	      53 allocs/op
BenchmarkFormat-10    	     242	   4922445 ns/op	 5809618 B/op	   79307 allocs/op
BenchmarkWrap-10      	      31	  37800359 ns/op	12121736 B/op	      55 allocs/op
BenchmarkGrep-10      	      92	  11491231 ns/op	 6578986 B/op	      72 allocs/op
BenchmarkPick-10      	   20880	     56936 ns/op	 1001022 B/op	      33 allocs/op
PASS
```
