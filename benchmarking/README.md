# benchmarking

## Run Benchmarks Once
```
# Run all benchmarks
$ go test -bench=.
BenchmarkConcatString/concat_string-12    114427  10303 ns/op  53488 B/op  99 allocs/op
BenchmarkConcatString/concat_buffer-12    774298   1499 ns/op  3440 B/op    6 allocs/op
BenchmarkConcatString/concat_builder-12  1305826	862 ns/op  2032 B/op    7 allocs/op
```

## Run Benchmarks Multiple Times
```
# Run benchmarks for BenchmarkConcatString 8 times and compute mean and variance
$ go test -bench=BenchmarkConcatString -count=8 >> old.txt
$ benchstat old.txt
name                            time/op
ConcatString/concat_string-12   10.1µs ± 1%
ConcatString/concat_buffer-12   1.43µs ± 4%
ConcatString/concat_builder-12   881ns ± 2%

name                            alloc/op
ConcatString/concat_string-12   53.5kB ± 0%
ConcatString/concat_buffer-12   3.44kB ± 0%
ConcatString/concat_builder-12  2.03kB ± 0%

name                            allocs/op
ConcatString/concat_string-12     99.0 ± 0%
ConcatString/concat_buffer-12     6.00 ± 0%
ConcatString/concat_builder-12    7.00 ± 0%
```
