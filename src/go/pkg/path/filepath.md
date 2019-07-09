# [Package filepath](https://golang.org/pkg/path/filepath/)

# Links

* [Golang学习 - path/filepath 包](https://www.cnblogs.com/golove/p/5903579.html)

# Constants

```go
const (
        Separator     = os.PathSeparator
        ListSeparator = os.PathListSeparator
)
```

```go
var SkipDir = errors.New("skip this directory")
```

```go
func Clean(path string) string 
func Join(elem ...string) string
func Split(path string) (dir, file string)
func Dir(path string) string 
func Base(path string) string 
func Ext(path string) string
func IsAbs(path string) bool 
func Match(pattern, name string) (matched bool, err error)

func Abs(path string) (string, error)
// Rel returns a relative path that is lexically equivalent to targpath when joined to basepath with an intervening separator. 
// That is, Join(basepath, Rel(basepath, targpath)) is equivalent to targpath itself.
func Rel(basepath, targpath string) (string, error)

func ToSlash(path string) string
func FromSlash(path string) string

func Glob(pattern string) (matches []string, err error)
func SplitList(path string) []string 
func VolumeName(path string) string
func EvalSymlinks(path string) (string, error)

func Walk(root string, walkFn WalkFunc) error
type WalkFunc func(path string, info os.FileInfo, err error) error
```

