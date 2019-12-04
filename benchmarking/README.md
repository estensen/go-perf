# benchmarking

## Run Benchmarks Once
```
# Run all benchmarks
$ go test -benchmarks=.
```

## Run Benchmarks Multiple Times
```
$ go test -bench=BenchmarkConcatString -count=8 >> old.txt
$ benchstat old.txt
```

