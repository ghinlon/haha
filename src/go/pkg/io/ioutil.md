# [Package ioutil](https://golang.org/pkg/io/ioutil/)

# Links

* [Golang学习 - io/ioutil 包](https://www.cnblogs.com/golove/p/3278444.html)


# Discard

`var Discard io.Writer = devNull(0)`

# Func

```go
func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)

func TempFile(dir, pattern string) (f *os.File, err error)
func ReadFile(filename string) ([]byte, error)
func WriteFile(filename string, data []byte, perm os.FileMode) error

func TempDir(dir, prefix string) (name string, err error)
func ReadDir(dirname string) ([]os.FileInfo, error)
```

