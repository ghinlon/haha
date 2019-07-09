# [Package bytes](https://golang.org/pkg/bytes/)

<!-- ToC start -->
# Table of Contents

1. [Links](#links)
1. [Package files](#package-files)
1. [Abstract](#abstract)
1. [Organize](#organize)
   1. [索引(15个函数)](#索引15个函数)
   1. [拆合(8个函数)](#拆合8个函数)
   1. [转换(10个函数)](#转换10个函数)
   1. [清理(9个函数)](#清理9个函数)
   1. [比较(3个函数)](#比较3个函数)
   1. [Buffer类型](#buffer类型)
   1. [读类型](#读类型)
<!-- ToC end -->

# Links

* [Golang学习 - bytes 包](http://www.cnblogs.com/golove/p/3287729.html)  

# Package files

[bytes.go](https://golang.org/src/bytes/bytes.go) [buffer.go](https://golang.org/src/bytes/buffer.go) [reader.go](https://golang.org/src/bytes/reader.go)

# Abstract

```
// index 15
F: IndexByte, Index, IndexRune, IndexAny, IndexFunc
F: LastIndexByte, LastIndex, LastIndexAny, LastIndexFunc
F: Contains, ContainsRune, ContainsAny
F: Count, HasPrefix, HasSuffix

// Split and Join 8
F: Split, SplitN, SplitAfter, SplitAfterN
F: FieldsFunc, Fields
F: Join, Repeat

// Map 10
F: Map, Replace, Runes
F: ToUpper, ToLower, ToTitle, ToUpperSpecial, ToLowerSpecial, ToTitleSpecial, Title

// Trim 9
F: TrimLeftFunc, TrimRightFunc, TrimFunc
F: TrimLeft, TrimRight, Trim
F: TrimSpace, TrimPrefix, TrimSuffix

// 3
F: Equal, EqualFold, Compare

T: Buffer
    b.I: R, W, RF, WT, ByteScanner, ByteWriter, RuneScanner
    b.F: WriteString, WriteRune
    b.F: ReadBytes, ReadString
    b.F: Bytes, String, Next, Len, Cap, Grow, Truncate, Reset
newF: NewBuffer, NewBufferString

T: Reader
    r.I: R, S, WT, RA, ByteScanner, RuneScanner
    r.F: Len, Size, Reset
newF: NewReader
```

# Organize

## 索引(15个函数)

```go
func IndexByte(s []byte, c byte) int
func Index(s, sep []byte) int
func IndexRune(s []byte, r rune) int
func IndexAny(s []byte, chars string) int
func IndexFunc(s []byte, f func(r rune) bool) int

func LastIndexByte(s []byte, c byte) int
func LastIndex(s, sep []byte) int
func LastIndexAny(s []byte, chars string) int
func LastIndexFunc(s []byte, f func(r rune) bool) int

func Contains(b, subslice []byte) bool {
    return Index(b, subslice) != -1
}
func ContainsAny(b []byte, chars string) bool {
    return IndexAny(b, chars) >= 0
}
func ContainsRune(b []byte, r rune) bool {
    return IndexRune(b, r) >= 0
}
func Count(s, sep []byte) int 
func HasPrefix(s, prefix []byte) bool
func HasSuffix(s, suffix []byte) bool
```

## 拆合(8个函数)

```go
func Split(s, sep []byte) [][]byte { return genSplit(s, sep, 0, -1) }
func SplitN(s, sep []byte, n int) [][]byte { return genSplit(s, sep, 0, n) }
func SplitAfter(s, sep []byte) [][]byte { return genSplit(s, sep, len(sep), -1) }
func SplitAfterN(s, sep []byte, n int) [][]byte { return genSplit(s, sep, len(sep), n) }
func FieldsFunc(s []byte, f func(rune) bool) [][]byte
func Fields(s []byte) [][]byte
func Join(s [][]byte, sep []byte) []byte 
func Repeat(b []byte, count int) []byte
```

## 转换(10个函数)

```go
func Map(mapping func(r rune) rune, s []byte) []byte 
func ToUpper(s []byte) []byte
func ToLower(s []byte) []byte 
func ToTitle(s []byte) []byte
func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte 
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte 
func Title(s []byte) []byte
func Replace(s, old, new []byte, n int) []byte
func Runes(s []byte) []rune 
```

## 清理(9个函数)

```
func TrimLeftFunc(s []byte, f func(r rune) bool) []byte
func TrimRightFunc(s []byte, f func(r rune) bool) []byte
func TrimFunc(s []byte, f func(r rune) bool) []byte {
    return TrimRightFunc(TrimLeftFunc(s, f), f)
}

func TrimLeft(s []byte, cutset string) []byte {
    return TrimLeftFunc(s, makeCutsetFunc(cutset))
}
func TrimRight(s []byte, cutset string) []byte {
    return TrimRightFunc(s, makeCutsetFunc(cutset))
}
func Trim(s []byte, cutset string) []byte {
    return TrimFunc(s, makeCutsetFunc(cutset))
}

func TrimSpace(s []byte) []byte {
    return TrimFunc(s, unicode.IsSpace)
}

func TrimPrefix(s, prefix []byte) []byte
func TrimSuffix(s, suffix []byte) []byte 
```

## 比较(3个函数)

```go
func Equal(a, b []byte) bool
func Compare(a, b []byte) int
func EqualFold(s, t []byte) bool
```

## Buffer类型

```
T: Buffer
    b.I: R, W, RF, WT, ByteScanner, ByteWriter, RuneScanner
    b.F: WriteString, WriteRune
    b.F: ReadBytes, ReadString
    b.F: Bytes, String, Next, Len, Cap, Grow, Truncate, Reset
newF: NewBuffer, NewBufferString
```

实现了7个IO接口: 读, 写, 读从, 写到, 字节扫描, 字节写, 符文扫描

实现了2个写函数: 写字符串, 写符文

实现了2个分隔符相关函数: 读字节串, 读符文串

实现了8个缓存相关函数: 字节串, 字符串, 接下来, 长度, 容量, 生长, 截短, 重置

```go
type Buffer struct {
    buf       []byte   // contents are the bytes buf[off : len(buf)]
    off       int      // read at &buf[off], write at &buf[len(buf)]
    bootstrap [64]byte // memory to hold first slice; helps small buffers avoid allocation.
    lastRead  readOp   // last read operation, so that Unread* can work correctly.
    // FIXME: it would be advisable to align Buffer to cachelines to avoid false
    // sharing.
}

func (b *Buffer) Read(p []byte) (n int, err error)
func (b *Buffer) Write(p []byte) (n int, err error) 
func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
func (b *Buffer) WriteTo(w io.Writer) (n int64, err error)
func (b *Buffer) ReadByte() (byte, error) 
func (b *Buffer) UnreadByte() error
func (b *Buffer) WriteByte(c byte) error
func (b *Buffer) ReadRune() (r rune, size int, err error)
func (b *Buffer) UnreadRune() error

func (b *Buffer) WriteString(s string) (n int, err error) 
func (b *Buffer) WriteRune(r rune) (n int, err error) 

func (b *Buffer) ReadBytes(delim byte) (line []byte, err error) 
func (b *Buffer) ReadString(delim byte) (line string, err error)

func (b *Buffer) Bytes() []byte { return b.buf[b.off:] }
func (b *Buffer) String() string
func (b *Buffer) Next(n int) []byte
func (b *Buffer) Len() int { return len(b.buf) - b.off }
func (b *Buffer) Cap() int { return cap(b.buf) }
func (b *Buffer) Grow(n int)
func (b *Buffer) Truncate(n int) 
func (b *Buffer) Reset() 

func NewBuffer(buf []byte) *Buffer { return &Buffer{buf: buf} }
func NewBufferString(s string) *Buffer { return &Buffer{buf: []byte(s)} }
```

## 读类型

实现了6个IO接口: 读, 找, 写到, 读在, 字节扫描, 符文扫描

```
T: Reader
    r.I: R, S, WT, RA, ByteScanner, RuneScanner
    r.F: Len, Size, Reset
newF: NewReader
```

```go
type Reader struct {
    s        []byte
    i        int64 // current reading index
    prevRune int   // index of previous rune; or < 0
}

func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) Seek(offset int64, whence int) (int64, error) 
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error) 
func (r *Reader) ReadByte() (byte, error) 
func (r *Reader) UnreadByte() error
func (r *Reader) ReadRune() (ch rune, size int, err error)
func (r *Reader) UnreadRune() error


func (r *Reader) Len() int
func (r *Reader) Size() int64 { return int64(len(r.s)) }
func (r *Reader) Reset(b []byte) { *r = Reader{b, 0, -1} }

func NewReader(b []byte) *Reader { return &Reader{b, 0, -1} }
```
