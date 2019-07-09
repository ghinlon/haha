# [Package path](https://golang.org/pkg/path/)


```go
func Clean(path string) string
func Join(elem ...string) string
func Split(path string) (dir, file string)
func Dir(path string) string
func Base(path string) string
func Ext(path string) string
func Match(pattern, name string) (matched bool, err error)
func IsAbs(path string) bool
```


