# [package yaml](https://godoc.org/gopkg.in/yaml.v2)

import "gopkg.in/yaml.v2"

# Links

* [GitHub - go-yaml/yaml: YAML support for the Go language.](https://github.com/go-yaml/yaml)

# 一对编解组函数

```go
func Marshal(in interface{}) (out []byte, err error)
func Unmarshal(in []byte, out interface{}) (err error)
func UnmarshalStrict(in []byte, out interface{}) (err error)
```

# 一对编解码器

```go
type Encoder struct {
    // contains filtered or unexported fields
}
func NewEncoder(w io.Writer) *Encoder

func (e *Encoder) Encode(v interface{}) (err error)
func (e *Encoder) Close() (err error)

type Decoder struct {
    // contains filtered or unexported fields
}
func NewDecoder(r io.Reader) *Decoder

func (dec *Decoder) Decode(v interface{}) (err error)
func (dec *Decoder) SetStrict(strict bool)
```

# 一对编解组接口

```go
type Marshaler interface {
    MarshalYAML() (interface{}, error)
}

type Unmarshaler interface {
    UnmarshalYAML(unmarshal func(interface{}) error) error
}
