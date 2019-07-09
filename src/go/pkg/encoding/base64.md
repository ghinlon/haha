# [Package base64](https://golang.org/pkg/encoding/base64/)

# Links

* [Base64 - 维基百科，自由的百科全书](https://zh.wikipedia.org/zh-hans/Base64)

# Overview

同asn1, json, gob这些库不大一样，别的库统都，要么编解组函数，要么套在读写接口上生成编解码器，再去编解码

这个是有一个编码结构，由这个结构的东西去编解码，或者是套在读写接口上，转换读写接口成读接口，写关接口，直接写数据到生成的编解码器

# Principle

摘自[Base64 - 维基百科，自由的百科全书](https://zh.wikipedia.org/zh-hans/Base64):

```
Base64是一种基于64个可打印字符来表示二进制数据的表示方法。由于 2^6 = 64 ，所以每6个位元为一个单元，对应某个可打印字符。3个字节有24个位元，对应于4个Base64单元，即3个字节可由4个可打印字符来表示。它可用来作为电子邮件的传输编码。在Base64中的可打印字符包括字母A-Z、a-z、数字0-9，这样共有62个字符，此外两个可打印符号在不同的系统中而不同。一些如uuencode的其他编码方法，和之后BinHex的版本使用不同的64字符集来代表6个二进制数字，但是不被称为Base64。 

Base64常用于在通常处理文本数据的场合，表示、传输、存储一些二进制数据，包括MIME的电子邮件及XML的一些复杂数据。
```

# Encoding结构

```go
type Encoding struct {
    encode    [64]byte
    decodeMap [256]byte
    padChar   rune
    strict    bool
}

func (enc *Encoding) Encode(dst, src []byte)
func (enc *Encoding) EncodeToString(src []byte) string
func (enc *Encoding) EncodedLen(n int) int

func (enc *Encoding) Decode(dst, src []byte) (n int, err error)
func (enc *Encoding) DecodeString(s string) ([]byte, error)
func (enc *Encoding) DecodedLen(n int) int

func (enc Encoding) Strict() *Encoding
func (enc Encoding) WithPadding(padding rune) *Encoding

func NewEncoding(encoder string) *Encoding

var StdEncoding = NewEncoding(encodeStd)
var URLEncoding = NewEncoding(encodeURL)
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)
```

# 编解码器本身即IO接口

写接口转换后变为写关闭接口,写操作之后一定要执行**关闭**

```go
func NewEncoder(enc *Encoding, w io.Writer) io.WriteCloser {
	return &encoder{enc: enc, w: w}
}

func NewDecoder(enc *Encoding, r io.Reader) io.Reader {
	return &decoder{enc: enc, r: &newlineFilteringReader{r}}
}


type encoder struct {
	err  error
	enc  *Encoding
	w    io.Writer
	buf  [3]byte    // buffered data waiting to be encoded
	nbuf int        // number of bytes in buf
	out  [1024]byte // output buffer
}

type decoder struct {
	err     error
	readErr error // error from r.Read
	enc     *Encoding
	r       io.Reader
	buf     [1024]byte // leftover input
	nbuf    int
	out     []byte // leftover decoded output
	outbuf  [1024 / 4 * 3]byte
}
```

[base64 Client and Server][1]

[1]: https://github.com/iofxl/practisego/blob/master/base64c.e.s.go

