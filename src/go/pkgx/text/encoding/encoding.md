# [package encoding](https://godoc.org/golang.org/x/text/encoding)

# Links

* [simplifiedchinese - GoDoc](https://godoc.org/golang.org/x/text/encoding/simplifiedchinese)
* [charmap - GoDoc](https://godoc.org/golang.org/x/text/encoding/charmap)

# type Decoder struct

```go
type Decoder struct {
    transform.Transformer
    // contains filtered or unexported fields
}

func (d *Decoder) Reader(r io.Reader) io.Reader
func (d *Decoder) Bytes(b []byte) ([]byte, error)
func (d *Decoder) String(s string) (string, error)
```

# type Encoder struct

```go
type Encoder struct {
    transform.Transformer
    // contains filtered or unexported fields
}

func HTMLEscapeUnsupported(e *Encoder) *Encoder
func ReplaceUnsupported(e *Encoder) *Encoder

func (e *Encoder) Writer(w io.Writer) io.Writer
func (e *Encoder) Bytes(b []byte) ([]byte, error)
func (e *Encoder) String(s string) (string, error)
```

# type Encoding interface

```go
type Encoding interface {
    // NewDecoder returns a Decoder.
    NewDecoder() *Decoder

    // NewEncoder returns an Encoder.
    NewEncoder() *Encoder
}

var Nop Encoding = nop{}
var Replacement Encoding = replacement{}
```

# package simplifiedchinese

import "golang.org/x/text/encoding/simplifiedchinese"

* GB 2312 标准共收录 6763 个汉字
* GBK 共收入 21886 个汉字和图形符号, 向下与 GB 2312 完全兼容
* GB 18030 与 GB 2312-1980 和 GBK 兼容，共收录汉字70244个, 编码是一二四字节变长
  编码。

## Variables

```go
var (
    // GB18030 is the GB18030 encoding.
    GB18030 encoding.Encoding = &gbk18030
    // GBK is the GBK encoding. It encodes an extension of the GB2312 character set
    // and is also known as Code Page 936.
    GBK encoding.Encoding = &gbk
)

var All = []encoding.Encoding{GB18030, GBK, HZGB2312}

var HZGB2312 encoding.Encoding = &hzGB2312
```

# package charmap

```go
var (
    // ISO8859_6E is the ISO 8859-6E encoding.
    ISO8859_6E encoding.Encoding = &iso8859_6E

    // ISO8859_6I is the ISO 8859-6I encoding.
    ISO8859_6I encoding.Encoding = &iso8859_6I

    // ISO8859_8E is the ISO 8859-8E encoding.
    ISO8859_8E encoding.Encoding = &iso8859_8E

    // ISO8859_8I is the ISO 8859-8I encoding.
    ISO8859_8I encoding.Encoding = &iso8859_8I
)
```

## type Charmap struct

```go
type Charmap struct {
    // contains filtered or unexported fields
}

// String returns the Charmap's name.
func (m *Charmap) String() string
func (m *Charmap) DecodeByte(b byte) rune
func (m *Charmap) EncodeRune(r rune) (b byte, ok bool)

func (m *Charmap) ID() (mib identifier.MIB, other string)

func (m *Charmap) NewDecoder() *encoding.Decoder
func (m *Charmap) NewEncoder() *encoding.Encoder
```

# Examples

```go
package main

import (
        "fmt"
        "io"
        "io/ioutil"
        "log"
        "os"

        "golang.org/x/text/encoding/simplifiedchinese"
)

func main() {

        if len(os.Args) == 1 {
                fmt.Println("Usage: os.Args[0] <gbk_file>")
                os.Exit(1)
        }

        f, err := os.Open(os.Args[1])
        if err != nil {
                log.Fatal(err)
        }

        defer f.Close()

        r := NewGBK2UTF8Reader(f)

        b, err := ioutil.ReadAll(r)

        if err != nil {
                log.Fatal(err)
        }

        fmt.Print(string(b))

}

func NewGBK2UTF8Reader(r io.Reader) io.Reader {
        dec := simplifiedchinese.GBK.NewDecoder()
        return dec.Reader(r)
}
```


