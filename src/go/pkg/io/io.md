# [Package io](https://golang.org/pkg/io/)

# Links

* [Golang学习 - io 包](http://www.cnblogs.com/golove/p/3276678.html)


# 4个最基本的接口

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

# 7个基本接口的组合接口

```go
type ReadWriter interface {
    Reader
    Writer
}

type ReadCloser interface {
    Reader
    Closer
}

type WriteCloser interface {
    Writer
    Closer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

type ReadSeeker interface {
    Reader
    Seeker
}

type WriteSeeker interface {
    Writer
    Seeker
}

type ReadWriteSeeker interface {
    Reader
    Writer
    Seeker
}
```

# 2个操作读写接口的接口

以读写接口作为参数的2个接口

读从接口,写到接口

既然是接口就一定要某种类型实现了才有效,然后由那种类型的东西来调用.而不能由io包来调用.

而下面的6个操作读写接口的函数则由io包来调用.

也就是，一个包里的接口，需要东西去实现了才能用 函数不是 函数可以用包名直接调用 

```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}

type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

# 2个偏移读写接口

这两个接口和前面的读写接口没有关系，是另一种接口，和基本的读写接口比，多了一个偏移参数 

```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}

type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

# 5个字节,符文相关的读写接口

```go
type ByteReader interface {
    ReadByte() (byte, error)
}

type ByteScanner interface {
    ByteReader
    UnreadByte() error
}

type ByteWriter interface {
    WriteByte(c byte) error
}

type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}

type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```

# 6个操作读写接口的函数

ReadAtLeast reads from r into buf until it has read at least min bytes. It returns the number of bytes copied and an error if fewer bytes were read. The error is EOF only if no bytes were read. If an EOF happens after reading fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF. If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer. On return, n >= min if and only if err == nil.

```go
func WriteString(w Writer, s string) (n int, err error) 

func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) 
func ReadFull(r Reader, buf []byte) (n int, err error) {
    return ReadAtLeast(r, buf, len(buf))
}

func Copy(dst Writer, src Reader) (written int64, err error)  
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
func CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
    written, err = Copy(dst, LimitReader(src, n))
    if written == n {
        return n, nil
    }
    if written < n && err == nil {
        // src stopped early; must have been EOF.
        err = EOF
    }
    return
}
```

# 4个包内实现了基本接口的类型

限制读类型、部分读类型、管道读类型、管道写类型

限制读类型是对读接口包装,部分读类型是对偏移读接口的包装.

也就是説,只要某种类型实现了读接口就可以转换为限制读类型,实现了偏移读接口就可以转换为部分读类型.

SectionReader implements Read, Seek, and ReadAt on a section of an underlying ReaderAt.

Pipe creates a synchronous in-memory pipe. It can be used to connect code expecting an io.Reader with code expecting an io.Writer.

Reads and Writes on the pipe are matched one to one except when multiple Reads are needed to consume a single Write. That is, each Write to the PipeWriter blocks until it has satisfied one or more Reads from the PipeReader that fully consume the written data. The data is copied directly from the Write to the corresponding Read (or Reads); there is no internal buffering.

It is safe to call Read and Write in parallel with each other or with Close. Parallel calls to Read and parallel calls to Write are also safe: the individual calls will be gated sequentially.

```go
type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}

func LimitReader(r Reader, n int64) Reader

func (l *LimitedReader) Read(p []byte) (n int, err error)


type SectionReader struct {
    r     ReaderAt
    base  int64
    off   int64
    limit int64
}

func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader

func (s *SectionReader) Read(p []byte) (n int, err error)
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
func (s *SectionReader) Size() int64


type PipeReader struct {
    p *pipe
}

type PipeWriter struct {
    p *pipe
}

func Pipe() (*PipeReader, *PipeWriter)

func (r *PipeReader) Read(data []byte) (n int, err error)
func (r *PipeReader) Close() error
func (r *PipeReader) CloseWithError(err error) error

func (w *PipeWriter) Write(data []byte) (n int, err error)
func (w *PipeWriter) Close() error
func (w *PipeWriter) CloseWithError(err error) error
```

# 4个生成特殊功能读写接口的函数

LimitReader returns a Reader that reads from r but stops with EOF after n bytes. The underlying implementation is a *LimitedReader.

TeeReader returns a Reader that writes to w what it reads from r. All reads from r performed through it are matched with corresponding writes to w. There is no internal buffering - the write must complete before the read completes. Any error encountered while writing is reported as a read error.

MultiReader returns a Reader that's the logical concatenation of the provided input readers. They're read sequentially. Once all inputs have returned EOF, Read will return EOF. If any of the readers return a non-nil, non-EOF error, Read will return that error.

MultiWriter creates a writer that duplicates its writes to all the provided writers, similar to the Unix tee(1) command.

Each write is written to each listed writer, one at a time. If a listed writer returns an error, that overall write operation stops and returns the error; it does not continue down the list.

```go
func LimitReader(r Reader, n int64) Reader
func TeeReader(r Reader, w Writer) Reader
func MultiReader(readers ...Reader) Reader
func MultiWriter(writers ...Writer) Writer
```

[Package bytes](https://golang.org/pkg/bytes/#pkg-index)中的Buffer类型实现IO接口
[Package os](https://golang.org/pkg/os/#pkg-index)中的File类型实现IO接口




