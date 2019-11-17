# [Package json](https://golang.org/pkg/encoding/json/)

# Links

* [JSON | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/dataserialisation/json.html)
* [JSON and Go - The Go Blog](https://blog.golang.org/json-and-go)
* [JSON-to-Go: Convert JSON to Go instantly](https://mholt.github.io/json-to-go/)


# Encoder and Decoder


```go
type Encoder struct {
        // Has unexported fields.
}
    An Encoder writes JSON values to an output stream.

func NewEncoder(w io.Writer) *Encoder
func (enc *Encoder) Encode(v interface{}) error
func (enc *Encoder) SetEscapeHTML(on bool)
func (enc *Encoder) SetIndent(prefix, indent string)

type Decoder struct {
        // Has unexported fields.
}
    A Decoder reads and decodes JSON values from an input stream.

func NewDecoder(r io.Reader) *Decoder
func (dec *Decoder) Buffered() io.Reader
func (dec *Decoder) Decode(v interface{}) error
func (dec *Decoder) DisallowUnknownFields()
func (dec *Decoder) More() bool
func (dec *Decoder) Token() (Token, error)
func (dec *Decoder) UseNumber()

func (dec *Decoder) More() bool
    More reports whether there is another element in the current array or object
    being parsed.
```

