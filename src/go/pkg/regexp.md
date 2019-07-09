# [Package regexp](https://golang.org/pkg/regexp/)

# Links

* [Syntax · google/re2 Wiki · GitHub](https://github.com/google/re2/wiki/Syntax)
* [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
* [基础知识 - Golang 中的正则表达式 - GoLove - 博客园](https://www.cnblogs.com/golove/p/3269099.html)
* [Golang学习 - regexp 包 - GoLove - 博客园](https://www.cnblogs.com/golove/p/3270918.html)


For an overview of the syntax, run

`go doc regexp/syntax`


`Find(All)?(String)?(Submatch)?(Index)?`

# type Regexp struct

```go
type Regexp struct {
    // read-only after Compile
    regexpRO
    // cache of machines for running regexp
    mu      sync.Mutex
    machine []*machine
}

func Compile(expr string) (*Regexp, error)
func CompilePOSIX(expr string) (*Regexp, error)
func MustCompile(str string) *Regexp
func MustCompilePOSIX(str string) *Regexp

func (re *Regexp) Find(b []byte) []byte
func (re *Regexp) FindAll(b []byte, n int) [][]byte
func (re *Regexp) FindSubmatch(b []byte) [][]byte
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
func (re *Regexp) FindIndex(b []byte) (loc []int)
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
func (re *Regexp) FindSubmatchIndex(b []byte) []int 
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int 

func (re *Regexp) FindString(s string) string
func (re *Regexp) FindAllString(s string, n int) []string
func (re *Regexp) FindStringSubmatch(s string) []string
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
func (re *Regexp) FindStringIndex(s string) (loc []int)
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
func (re *Regexp) FindStringSubmatchIndex(s string) []int
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int) 
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int 

func (re *Regexp) Match(b []byte) bool
func (re *Regexp) MatchString(s string) bool
func (re *Regexp) MatchReader(r io.RuneReader) bool

func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte

func (re *Regexp) ReplaceAll(src, repl []byte) []byte
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte 
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte 
func (re *Regexp) ReplaceAllString(src, repl string) string 
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

func (re *Regexp) Split(s string, n int) []string 

func (re *Regexp) String() string 
func (re *Regexp) Copy() *Regexp
func (re *Regexp) Longest() 
func (re *Regexp) NumSubexp() int 
func (re *Regexp) SubexpNames() []string
func (re *Regexp) LiteralPrefix() (prefix string, complete bool)

func Match(pattern string, b []byte) (matched bool, err error)
func MatchString(pattern string, s string) (matched bool, err error)
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) 
func QuoteMeta(s string) string
```
