# [Package gob](https://golang.org/pkg/encoding/gob/)

# Links

* [Gobs of data - The Go Blog](https://blog.golang.org/gobs-of-data)

# Overview

[Overview](https://golang.org/pkg/encoding/gob/#pkg-overview)

```go
struct { A, B int }
```

can be sent from or received into any of these Go types:

```go
struct { A, B int }	// the same
*struct { A, B int }	// extra indirection of the struct
struct { *A, **B int }	// extra indirection of the fields
struct { A, B int64 }	// different concrete value type; see below
```

It may also be received into any of these:

```go
struct { A, B int }	// the same
struct { B, A int }	// ordering doesn't matter; matching is by name
struct { A, B, C int }	// extra field (C) ignored
struct { B int }	// missing field (A) ignored; data will be dropped
struct { B, C int }	// missing field (A) ignored; extra field (C) ignored.
```

Attempting to receive into these types will draw a decode error:

```go
struct { A int; B uint }	// change of signedness for B
struct { A int; B float }	// change of type for B
struct { }			// no field names in common
struct { C, D int }		// no field names in common
```

Gob can encode a value of any type implementing the GobEncoder or encoding.BinaryMarshaler interfaces by calling the corresponding method, in that order of preference.

Gob can decode a value of any type implementing the GobDecoder or encoding.BinaryUnmarshaler interfaces by calling the corresponding method, again in that order of preference. 

# 一对Gob編解码接口

```go
type GobEncoder interface {
        // GobEncode returns a byte slice representing the encoding of the
        // receiver for transmission to a GobDecoder, usually of the same
        // concrete type.
        GobEncode() ([]byte, error)
}

type GobDecoder interface {
        // GobDecode overwrites the receiver, which must be a pointer,
        // with the value represented by the byte slice, which was written
        // by GobEncode, usually for the same concrete type.
        GobDecode([]byte) error
}

func Register(value interface{})
func RegisterName(name string, value interface{})
```

# 一对編解码器结构

```go
type Encoder struct {
        // contains filtered or unexported fields
}
func (enc *Encoder) Encode(e interface{}) error {
    return enc.EncodeValue(reflect.ValueOf(e))
}
func (enc *Encoder) EncodeValue(value reflect.Value) error

func NewEncoder(w io.Writer) *Encoder

type Decoder struct {
        // contains filtered or unexported fields
}
func (dec *Decoder) Decode(e interface{}) error
func (dec *Decoder) DecodeValue(v reflect.Value) error

func NewDecoder(r io.Reader) *Decoder
```

[savegob.go][1]

[1]: https://github.com/iofxl/practisego/blob/master/savegob.go




