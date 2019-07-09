# [Package tar](https://golang.org/pkg/archive/tar/)

# Index


# type Header struct

```go
type Header struct
  func FileInfoHeader(fi os.FileInfo, link string) (*Header, error)
  func (h *Header) FileInfo() os.FileInfo
```

# type Reader struct

```go
type Reader
  func NewReader(r io.Reader) *Reader
  func (tr *Reader) Next() (*Header, error)
  func (tr *Reader) Read(b []byte) (int, error)
```

# type Writer struct

```go
type Writer
  func NewWriter(w io.Writer) *Writer
  func (tw *Writer) Close() error
  func (tw *Writer) Flush() error
  func (tw *Writer) Write(b []byte) (int, error)
  func (tw *Writer) WriteHeader(hdr *Header) error
```
