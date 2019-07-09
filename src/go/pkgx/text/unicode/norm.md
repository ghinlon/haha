# [norm - GoDoc](https://godoc.org/golang.org/x/text/unicode/norm)

# Links

* [Unicode等價性 - 维基百科，自由的百科全书](https://zh.wikipedia.org/zh/Unicode%E7%AD%89%E5%83%B9%E6%80%A7)
* [UAX #15: Unicode Normalization Forms](https://unicode.org/reports/tr15/)
* [Canonical Equivalence in Applications](https://unicode.org/notes/tn5/)


# type Form int

```go
type Form int

const (
    NFC Form = iota
    NFD
    NFKC
    NFKD
)


func (f Form) Bytes(b []byte) []byte
func (f Form) String(s string) string
func (f Form) IsNormal(b []byte) bool
func (f Form) IsNormalString(s string) bool

func (f Form) Append(out []byte, src ...byte) []byte
func (f Form) AppendString(out []byte, src string) []byte
func (f Form) FirstBoundary(b []byte) int
func (f Form) FirstBoundaryInString(s string) int
func (f Form) NextBoundary(b []byte, atEOF bool) int
func (f Form) NextBoundaryInString(s string, atEOF bool) int
func (f Form) LastBoundary(b []byte) int

func (f Form) Properties(s []byte) Properties
func (f Form) PropertiesString(s string) Properties

// Transform and Reset implements the Transform method of the transform.Transformer interface.
func (f Form) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)
func (Form) Reset()

// Span implements transform.SpanningTransformer.
func (f Form) Span(b []byte, atEOF bool) (n int, err error)
func (f Form) SpanString(s string, atEOF bool) (n int, err error)
func (f Form) QuickSpan(b []byte) int
func (f Form) QuickSpanString(s string) int

func (f Form) Reader(r io.Reader) io.Reader
func (f Form) Writer(w io.Writer) io.WriteCloser
```

# type Properties struct

```go
type Properties struct {
    // contains filtered or unexported fields
}

func (p Properties) BoundaryAfter() bool
func (p Properties) BoundaryBefore() bool

func (p Properties) Size() int
func (p Properties) Decomposition() []byte

func (p Properties) CCC() uint8
func (p Properties) LeadCCC() uint8
func (p Properties) TrailCCC() uint8
```

# type Iter struct

```go
type Iter struct {
    // contains filtered or unexported fields
}

func (i *Iter) Init(f Form, src []byte)
func (i *Iter) InitString(f Form, src string)
func (i *Iter) Next() []byte
func (i *Iter) Done() bool
func (i *Iter) Pos() int
func (i *Iter) Seek(offset int64, whence int) (int64, error)
```
