# [Package utf16](https://golang.org/pkg/unicode/utf16/)

```go
func Encode(s []rune) []uint16
func EncodeRune(r rune) (r1, r2 rune)
func Decode(s []uint16) []rune
func DecodeRune(r1, r2 rune) rune
func IsSurrogate(r rune) bool
```
