# [Package aes](https://golang.org/pkg/crypto/aes/)

# 生成块接口的函数

```go
func NewCipher(key []byte) (cipher.Block, error)
```

NewCipher creates and returns a new cipher.Block. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256. 

key的长度是有讲究的，生成固定长度的办法：

1. hash
2. [hkdf - GoDoc](https://godoc.org/golang.org/x/crypto/hkdf)



