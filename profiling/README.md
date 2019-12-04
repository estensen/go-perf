# tmp-go-meetup

## Run
```
$ go build
$ ./profiling
```

## Pad one string
```
$ curl -X GET "localhost:8080/leftpad?str=hello&padding=10"
"          hello"
```

## Pad many strings (probably not realistic usecase of the service)
```
$ ab -k -c 8 -n 1000000 "localhost:8080/leftpad?str=hello&length=10000"
```

## Profile CPU
```
$ go tool pprof -http=:8000 http://localhost:8080/debug/pprof/profile\?seconds\=10
```

## Profile Memory Allocations
```
$ go tool pprof -http=:8000 http://localhost:8080/debug/pprof/heap
```

## Get Tracing Date
```
# Profiles can also be saved to files
$ curl http://localhost:8080/debug/pprof/trace\?seconds\=5 > trace.out
$ go tool trace trace.out
```

