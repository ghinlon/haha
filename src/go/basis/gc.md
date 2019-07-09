# Garbage Collector


# Links

* [Getting to Go: The Journey of Go's Garbage Collector - The Go Blog](https://blog.golang.org/ismmkeynote)
* [proposal/14951-soft-heap-limit.md at master · golang/proposal · GitHub](https://github.com/golang/proposal/blob/master/design/14951-soft-heap-limit.md)
* [runtime: reclaim memory used by huge array that is no longer referenced · Issue #14045 · golang/go · GitHub](https://github.com/golang/go/issues/14045)
* [Print The Current Memory Usage · GolangCode](https://golangcode.com/print-the-current-memory-usage/)
* [runtime - The Go Programming Language](https://golang.org/pkg/runtime/)
* [debug - The Go Programming Language](https://golang.org/pkg/runtime/debug/)
* [Go GC - LINE ENGINEERING](https://engineering.linecorp.com/en/blog/go-gc/)
* [Memory Management in Go](https://dougrichardson.org/2016/01/23/go-memory-allocations.html)


# func FreeOSMemory

FreeOSMemory forces a garbage collection followed by an attempt to return as much memory to the operating system as possible. (Even if this is not called, the runtime gradually returns memory to the operating system in a background task.) 

```go
func debug.FreeOSMemory()
```

# Print The Current Memory Usage 


```go
m := &runtime.MemStats{}

runtime.ReadMemStats(m)

fmt.Printf("Alloc = %v MiB", m.Alloc)
fmt.Printf("\tTotalAlloc = %v MiB", m.TotalAlloc)
fmt.Printf("\tSys = %v MiB", m.Sys)
fmt.Printf("\tNumGC = %v\n", m.NumGC)
```

# Mark Mem As RES

```go
func allocate(n int) []byte {

	b := make([]byte, n)

	for i := 0; i < n; i += os.Getpagesize() {
		b[i] = 1
	}

	return b
}
```
