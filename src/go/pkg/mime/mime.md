# [Package mime](https://golang.org/pkg/mime/)

# Links

* [MIME 类型 | MDN](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/MIME_types)

# type WordEncoder 

```go
type WordEncoder byte

func (e WordEncoder) Encode(charset, s string) string

const (
	// BEncoding represents Base64 encoding scheme as defined by RFC 2045.
	BEncoding = WordEncoder('b')
	// QEncoding represents the Q-encoding scheme as defined by RFC 2047.
	QEncoding = WordEncoder('q')
)
```

# type WordDecoder struct 

```go
type WordDecoder struct {
	// CharsetReader, if non-nil, defines a function to generate
	// charset-conversion readers, converting from the provided
	// charset into UTF-8.
	// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
	// are handled by default.
	// One of the CharsetReader's result values must be non-nil.
	CharsetReader func(charset string, input io.Reader) (io.Reader, error)
}

func (d *WordDecoder) Decode(word string) (string, error)
func (d *WordDecoder) DecodeHeader(header string) (string, error)
```

# ParseMediaType and FormatMediaType

```go
func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
func FormatMediaType(t string, param map[string]string) string
```

