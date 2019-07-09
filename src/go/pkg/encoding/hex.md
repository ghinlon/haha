# [Package hex](https://golang.org/pkg/encoding/hex/)

# Links

# Principle

hex也称为base16，意思是使用16个可见字符来表示一个二进制数组，编码后数据大小将翻倍,因为1个字符需要用2个可见字符来表示。

# 编解码函数

```go
func EncodeToString(src []byte) string
func DecodeString(s string) ([]byte, error)

func Encode(dst, src []byte) int
func Decode(dst, src []byte) (int, error)

func EncodedLen(n int) int
func DecodedLen(x int) int

func Dump(data []byte) string 
```

# 编解码器本身即IO接口

这个跟base64包里是一样的

```go
func NewEncoder(w io.Writer) io.Writer
func NewDecoder(r io.Reader) io.Reader

func Dumper(w io.Writer) io.WriteCloser 
```

