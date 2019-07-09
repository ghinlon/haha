# [Package bufio](https://golang.org/pkg/bufio/)

# Links

* [Golang学习 - bufio 包](https://www.cnblogs.com/golove/p/3282667.html)   
* [“Introduction to bufio package in Golang” @mlowicki](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762)  
* [“In-depth introduction to bufio.Scanner in Golang” @mlowicki](https://medium.com/golangspec/in-depth-introduction-to-bufio-scanner-in-golang-55483bb689b4)  

# Overview

Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O. 

# type Reader struct

```go
const (
    defaultBufSize = 4096
)

type Reader struct {
    buf          []byte
    rd           io.Reader // reader provided by the client
    r, w         int       // buf read and write positions
    err          error
    lastByte     int
    lastRuneSize int
}

func NewReaderSize(rd io.Reader, size int) *Reader
func NewReader(rd io.Reader) *Reader {
    return NewReaderSize(rd, defaultBufSize)
}
//只要缓存有数据就只从缓存读
//要读的比缓存还大就不用缓存直接读底层
func (b *Reader) Read(p []byte) (n int, err error)
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)
func (b *Reader) ReadByte() (byte, error)
func (b *Reader) UnreadByte() error
func (b *Reader) ReadRune() (r rune, size int, err error) 
func (b *Reader) UnreadRune() error

//缓存满了也没有找到delim就报错io.ErrBufferFull，也就是说这个函数每次调用最多处理到缓存满
//遇到EOF了没有找到delim就报错io.EOF
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
//如果一行的长度超过缓存的长度会截断返回
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
//ReadBytes会多次调用ReadSlice积累数据直到找到delim,所以这个函数不受缓存大小限制
func (b *Reader) ReadBytes(delim byte) ([]byte, error) 
func (b *Reader) ReadString(delim byte) (string, error) {
    bytes, err := b.ReadBytes(delim)
    return string(bytes), err
}

func (r *Reader) Size() int { return len(r.buf) }
func (b *Reader) Reset(r io.Reader) { b.reset(b.buf, r) }
//要看的没看够就拉东西进缓存
//最多看缓存大小那么多，如果想看的比缓存大小还多就报错ErrBufferFull
//到最后还要看的没看够会报错
func (b *Reader) Peek(n int) ([]byte, error)
//实际discarded小于n就报错
//如果缓存里数据不够就会去读底层
func (b *Reader) Discard(n int) (discarded int, err error) 
func (b *Reader) Buffered() int { return b.w - b.r }
```

# type Writer struct

```go
type Writer struct {
    err error
    buf []byte
    n   int
    wr  io.Writer
}

func NewWriterSize(w io.Writer, size int) *Writer 
func NewWriter(w io.Writer) *Writer {
    return NewWriterSize(w, defaultBufSize)
}
//只要报错了就不再写了
//要写的比缓存还大就不用缓存直接读写底层
func (b *Writer) Write(p []byte) (nn int, err error)
func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
func (b *Writer) WriteByte(c byte) error
func (b *Writer) WriteRune(r rune) (size int, err error)
func (b *Writer) WriteString(s string) (int, error)

func (b *Writer) Size() int { return len(b.buf) }
//重置可以换底层，就不用重新分配内存
func (b *Writer) Reset(w io.Writer) 
func (b *Writer) Available() int { return len(b.buf) - b.n }
func (b *Writer) Flush() error 
func (b *Writer) Buffered() int { return b.n }
```

# type ReadWriter struct

```go
type ReadWriter struct {
    *Reader
    *Writer
}

func NewReadWriter(r *Reader, w *Writer) *ReadWriter {
    return &ReadWriter{r, w}
}
```

# type Scanner struct

```go
type Scanner struct {
    r            io.Reader // The reader provided by the client.
    split        SplitFunc // The function to split the tokens.
    maxTokenSize int       // Maximum size of a token; modified by tests.
    token        []byte    // Last token returned by split.
    buf          []byte    // Buffer used as argument to split.
    start        int       // First non-processed byte in buf.
    end          int       // End of data in buf.
    err          error     // Sticky error.
    empties      int       // Count of successive empty tokens.
    scanCalled   bool      // Scan has been called; buffer is in use.
    done         bool      // Scan has finished.
}

func NewScanner(r io.Reader) *Scanner {
    return &Scanner{
        r:            r,
        split:        ScanLines,
        maxTokenSize: MaxScanTokenSize,
    }
}

func (s *Scanner) Buffer(buf []byte, max int)
func (s *Scanner) Split(split SplitFunc)
//不允许Token长度超过缓存，会报错ErrTooLong
//缓存不空或者EOF了就会调用SplitFunc
func (s *Scanner) Scan() bool
func (s *Scanner) Bytes() []byte {
    return s.token
}
func (s *Scanner) Text() string {
    return string(s.token)
}
// Err returns the first non-EOF error that was encountered by the Scanner.
func (s *Scanner) Err() error {
    if s.err == io.EOF {
        return nil
    }
    return s.err
}
```

# SplitFunc

**SplitFunc's three behavior **  

* return `0，nil, nil` means want more data
* Token Found
* Once Error Scanner stop
    * 有一个特别的报错是ErrFinalToken，报这个错本次Scan返回正确，下次调用Scan直接停止。io.EOF和ErrFinalToken这两个报错都不被当作真正的错误处理的，也就是Err处理会返回nil.  

```go
type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
```


# ReadBytes('\n') ReadLine Scanner 比较

* `ReadBytes`不自动处理`\r\n`，并且和`\n`一起返回
* 行长度限制  
    * `ReadLine`如果一行长度超过缓存长度需要再次调用读取剩下的
    * `Scanner`如果行长度超过缓存长度就不会做任何处理直接报错`token too long`
    * `ReadBytes`没有任何限制  
* `Scanner`最简单 `ReadBytes`最牛逼 `ReadLine`有点假

# 使用bufio包的标准库

* archive/zip  
* compress/*  
* encoding/*  
* image/*  
* net/http  

