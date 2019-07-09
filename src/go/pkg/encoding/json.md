# [Package json](https://golang.org/pkg/encoding/json/)

# Links

* [JSON | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/dataserialisation/json.html)
* [JSON and Go - The Go Blog](https://blog.golang.org/json-and-go)
* [JSON-to-Go: Convert JSON to Go instantly](https://mholt.github.io/json-to-go/)

# 一对编解组函数 

这对函数会优先使json.编解码器接口，否则使encoding.文本编解组接口，再不行使一般json处理规则

```go
func Marshal(v interface{}) ([]byte, error)
func Unmarshal(data []byte, v interface{}) error

func Indent(dst *bytes.Buffer, src []byte, prefix, indent string) error
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	b, err := Marshal(v)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Valid(data []byte) bool
```

[json1.go][1]

[1]: https://github.com/iofxl/practisego/blob/master/json1.go

# 一对编解码器

基本的使用方法就是套在读写接口里生成编解码器，再使这个编解码器去编解码东西

解码要先dec.More()判断么有东西类

```go
type Encoder struct {
	w          io.Writer
	err        error
	escapeHTML bool

	indentBuf    *bytes.Buffer
	indentPrefix string
	indentValue  string
}

func (enc *Encoder) Encode(v interface{}) error
func (enc *Encoder) SetEscapeHTML(on bool)
func (enc *Encoder) SetIndent(prefix, indent string)

func NewEncoder(w io.Writer) *Encoder

type Decoder struct {
	r       io.Reader
	buf     []byte
	d       decodeState
	scanp   int   // start of unread data in buf
	scanned int64 // amount of data already scanned
	scan    scanner
	err     error
	tokenState int
	tokenStack []int
}

func (dec *Decoder) Buffered() io.Reader
func (dec *Decoder) Decode(v interface{}) error
func (dec *Decoder) DisallowUnknownFields()
func (dec *Decoder) More() bool
func (dec *Decoder) Token() (Token, error)
func (dec *Decoder) UseNumber()

func NewDecoder(r io.Reader) *Decoder
```

[saveJSON.go][2] [loadJSON.go][3] [json client and server][4]

[2]: https://github.com/iofxl/practisego/blob/master/savejson.go
[3]: https://github.com/iofxl/practisego/blob/master/loadjson.go
[4]: https://github.com/iofxl/practisego/blob/master/json.c.and.s.go






