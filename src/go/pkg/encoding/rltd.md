# encoding_rltd

* [Package binary](https://golang.org/pkg/encoding/binary/)
* [Package csv](https://golang.org/pkg/encoding/csv/)


# [Package binary](https://golang.org/pkg/encoding/binary/)

Package binary implements simple translation between numbers and byte sequences and encoding and decoding of varints. 

```go
func PutUvarint(buf []byte, x uint64) int
func Uvarint(buf []byte) (uint64, int) 

func PutVarint(buf []byte, x int64) int
func Varint(buf []byte) (int64, int) 

func ReadUvarint(r io.ByteReader) (uint64, error)
func ReadVarint(r io.ByteReader) (int64, error) 
```


## 1个字节序接口

### 2个实现了该接口的变量

```go
type ByteOrder interface {
    Uint16([]byte) uint16
    Uint32([]byte) uint32
    Uint64([]byte) uint64
    PutUint16([]byte, uint16)
    PutUint32([]byte, uint32)
    PutUint64([]byte, uint64)
    String() string
}

var LittleEndian littleEndian
var BigEndian bigEndian
```

## 2个从读写接口編解码的函数

Data must be a fixed-size value or a slice of fixed-size values, or a pointer to such data.

```go
func Read(r io.Reader, order ByteOrder, data interface{}) error
func Write(w io.Writer, order ByteOrder, data interface{}) error
```

## Size函数

返回编码v需要的多少字节

```go
func Size(v interface{}) int 
```

# [Package csv](https://golang.org/pkg/encoding/csv/)

空格也算数据

这里的读写类型不是读写接口

```
T: Reader
    r.F: Read, ReadAll
newF: NewReader
T: Writer
    w.F: Write, WriteAll, Flush, Error
newF: NewWriter
T: ParseError
    e.F: Error
```

## 读类型

```go
type Reader struct {
    Comma rune
    Comment rune
    FieldsPerRecord int
    LazyQuotes bool
    TrimLeadingSpace bool
    ReuseRecord bool
    TrailingComma bool // Deprecated: No longer used.
    // contains filtered or unexported fields
}

func NewReader(r io.Reader) *Reader
func (r *Reader) Read() (record []string, err error)
func (r *Reader) ReadAll() (records [][]string, err error)

```

## 写类型

```go
type Writer struct {
    Comma   rune // Field delimiter (set to ',' by NewWriter)
    UseCRLF bool // True to use \r\n as the line terminator
    // contains filtered or unexported fields
}

func NewWriter(w io.Writer) *Writer
func (w *Writer) Write(record []string) error
func (w *Writer) WriteAll(records [][]string) error
func (w *Writer) Flush()
func (w *Writer) Error() error
```

## 解析错误类型

```go
type ParseError struct {
        StartLine int   // Line where the record starts
        Line      int   // Line where the error occurred
        Column    int   // Column (rune index) where the error occurred
        Err       error // The actual error
}

func (e *ParseError) Error() string
```

