# [Package rand](https://golang.org/pkg/crypto/rand/)

```go
var Reader io.Reader
func Int(rand io.Reader, max *big.Int) (n *big.Int, err error)
func Prime(rand io.Reader, bits int) (p *big.Int, err error)
func Read(b []byte) (n int, err error)
```

