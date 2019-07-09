# [transform - GoDoc](https://godoc.org/golang.org/x/text/transform)

# Variables

```go
var (
    // Discard is a Transformer for which all Transform calls succeed
    // by consuming all bytes and writing nothing.
    Discard Transformer = discard{}

    // Nop is a SpanningTransformer that copies src to dst.
    Nop SpanningTransformer = nop{}
)
```
# type Transformer and SpanningTransformer interface

```go
type Transformer interface {
    // Transform writes to dst the transformed bytes read from src, and
    // returns the number of dst bytes written and src bytes read. The
    // atEOF argument tells whether src represents the last bytes of the
    // input.
    //
    // Callers should always process the nDst bytes produced and account
    // for the nSrc bytes consumed before considering the error err.
    //
    // A nil error means that all of the transformed bytes (whether freshly
    // transformed from src or left over from previous Transform calls)
    // were written to dst. A nil error can be returned regardless of
    // whether atEOF is true. If err is nil then nSrc must equal len(src);
    // the converse is not necessarily true.
    //
    // ErrShortDst means that dst was too short to receive all of the
    // transformed bytes. ErrShortSrc means that src had insufficient data
    // to complete the transformation. If both conditions apply, then
    // either error may be returned. Other than the error conditions listed
    // here, implementations are free to report other errors that arise.
    Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error)

    // Reset resets the state and allows a Transformer to be reused.
    Reset()
}

func Chain(t ...Transformer) Transformer
func RemoveFunc(f func(r rune) bool) Transformer

func Append(t Transformer, dst, src []byte) (result []byte, n int, err error)
func Bytes(t Transformer, b []byte) (result []byte, n int, err error)
func String(t Transformer, s string) (result string, n int, err error)

type NopResetter struct{}
	func (NopResetter) Reset()
```

# type SpanningTransformer interface

```go
type SpanningTransformer interface {
    Transformer

    // Span returns a position in src such that transforming src[:n] results in
    // identical output src[:n] for these bytes. It does not necessarily return
    // the largest such n. The atEOF argument tells whether src represents the
    // last bytes of the input.
    //
    // Callers should always account for the n bytes consumed before
    // considering the error err.
    //
    // A nil error means that all input bytes are known to be identical to the
    // output produced by the Transformer. A nil error can be returned
    // regardless of whether atEOF is true. If err is nil, then n must
    // equal len(src); the converse is not necessarily true.
    //
    // ErrEndOfSpan means that the Transformer output may differ from the
    // input after n bytes. Note that n may be len(src), meaning that the output
    // would contain additional bytes after otherwise identical output.
    // ErrShortSrc means that src had insufficient data to determine whether the
    // remaining bytes would change. Other than the error conditions listed
    // here, implementations are free to report other errors that arise.
    //
    // Calling Span can modify the Transformer state as a side effect. In
    // effect, it does the transformation just as calling Transform would, only
    // without copying to a destination buffer and only up to a point it can
    // determine the input and output bytes are the same. This is obviously more
    // limited than calling Transform, but can be more efficient in terms of
    // copying and allocating buffers. Calls to Span and Transform may be
    // interleaved.
    Span(src []byte, atEOF bool) (n int, err error)
}
```

# type Reader struct

```go
type Reader struct {
    // contains filtered or unexported fields
}

func NewReader(r io.Reader, t Transformer) *Reader

func (r *Reader) Read(p []byte) (int, error)
```

# type Writer struct

Writer wraps another io.Writer by transforming the bytes read. The user needs
to call Close to flush unwritten bytes that may be buffered. 

```go
type Writer struct {
    // contains filtered or unexported fields
}
func NewWriter(w io.Writer, t Transformer) *Writer
func (w *Writer) Close() error
func (w *Writer) Write(data []byte) (n int, err error)
```
