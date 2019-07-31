# Pipe io.ReadWriteCloser

# Links

* [chisel/pipe.go at master · jpillora/chisel · GitHub](https://github.com/jpillora/chisel/blob/master/share/pipe.go#L8)

# Code

```go
func Pipe(src io.ReadWriteCloser, dst io.ReadWriteCloser) (int64, int64) {
	var sent, received int64
	var wg sync.WaitGroup
	var once sync.Once
	close := func() {
		src.Close()
		dst.Close()
	}
	wg.Add(2)
	go func() {
		received, _ = io.Copy(src, dst)
		once.Do(close)
		wg.Done()
	}()
	go func() {
		sent, _ = io.Copy(dst, src)
		once.Do(close)
		wg.Done()
	}()
	wg.Wait()
	return sent, received
}
```



