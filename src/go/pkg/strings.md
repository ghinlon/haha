# [Package strings](https://golang.org/pkg/strings/)

<!-- ToC start -->
# Table of Contents

1. [Links](#links)
1. [Package files](#package-files)
1. [Abstract](#abstract)
1. [Organize](#organize)
   1. [索引(15个函数)](#索引15个函数)
   1. [拆合(8个函数)](#拆合8个函数)
   1. [转换(9个函数)](#转换9个函数)
   1. [清理(9个函数)](#清理9个函数)
   1. [比较(2个函数)](#比较2个函数)
   1. [读类型](#读类型)
   1. [Replacer类型](#Replacer类型)
   1. [Builder类型](#Builder类型)
<!-- ToC end -->

# Links

* [Golang学习 - strings 包](http://www.cnblogs.com/golove/p/3236300.html)  

# Package files

[strings_decl.go](https://golang.org/src/strings/strings_decl.go) [strings_amd64.go](https://golang.org/src/strings/strings_amd64.go) [strings.go](https://golang.org/src/strings/strings.go) [reader.go](https://golang.org/src/strings/reader.go)

# Abstract

```
// index 15
F: IndexByte, Index, IndexRune, IndexAny, IndexFunc
F: LastIndexByte, LastIndex, LastIndexAny, LastIndexFunc
F: Count, Contains, ContainsRune, ContainsAny, HasPrefix, HasSuffix

// Split and Join 8
F: Split, SplitN, SplitAfter, SplitAfterN
F: FieldsFunc, Fields
F: Join, Repeat

// Map 9
F: Map, Replace
F: ToUpper, ToLower, ToTitle, ToUpperSpecial, ToLowerSpecial, ToTitleSpecial, Title

// Trim 9
F: TrimLeftFunc, TrimRightFunc, TrimFunc
F: TrimLeft, TrimRight, Trim
F: TrimSpace, TrimPrefix, TrimSuffix

// 2
F: EqualFold, Compare

T: Reader
    r.I: R, S, WT, RA, ByteScanner, RuneScanner
    r.F: Len, Size, Reset
newF: NewReader

T: Replacer
    r.i: replacer
newF: NewReplacer

T: Builder
    b.I: W, ByteWriter
    b.F: WriteRune, WriteString
    b.F: String, Len, Grow, Reset
```

# Organize

## 索引(15个函数)

空子串""也属于s, Count等于len(s)+1

索引字节成长为索引子串, 索引子串成长为索引符文, 包含函数就是索引函数, 计数也是索引子串的变相

索引任意是另一种方式的索引字节

5个索引函数: 索引字节, 索引, 索引符文, 索引任意, 索引Func

4个末次索引函数: 末次索引字节, 末次索引, 末次索引任意, 末次索引Func

6个相关子串函数:計数, 包含, 包含符文, 包含任意, 有头, 有尾

```
F: IndexByte, Index, IndexRune, IndexAny, IndexFunc
F: LastIndexByte, LastIndex, LastIndexAny, LastIndexFunc
F: Count, Contains, ContainsRune, ContainsAny, HasPrefix, HasSuffix
```

```go
func IndexByte(s string, c byte) int
func Index(s, substr string) int 
func IndexRune(s string, r rune) int
func IndexAny(s, chars string) int 
func IndexFunc(s string, f func(rune) bool) int {
    return indexFunc(s, f, true)
}

func LastIndexByte(s string, c byte) int
func LastIndex(s, substr string) int
func LastIndexAny(s, chars string) int
func LastIndexFunc(s string, f func(rune) bool) int {
    return lastIndexFunc(s, f, true)
}

func Count(s, substr string) int
func Contains(s, substr string) bool {
    return Index(s, substr) >= 0
}
func ContainsRune(s string, r rune) bool
func ContainsAny(s, chars string) bool
func HasPrefix(s, prefix string) bool
func HasSuffix(s, suffix string) bool 
```

## 拆合(8个函数)

也符合常识: 

* 如果 s 不包含 sep, 返回s自己
* 如果 sep 空, 返回s的每个符文
* 如果s sep 都空, 返回空 []
* 如果 s == sep, 返回两个空 [ ], 头一个,尾一个

切, 切N, 切之后, 切之后N, 域Func, 域, 连接, 重复

```
F: Split, SplitN, SplitAfter, SplitAfterN
F: FieldsFunc, Fields
F: Join, Repeat
```

```go
func Split(s, sep string) []string              { return genSplit(s, sep, 0, -1) }
func SplitN(s, sep string, n int) []string      { return genSplit(s, sep, 0, n) }

func SplitAfter(s, sep string) []string         { return genSplit(s, sep, len(sep), -1) }
func SplitAfterN(s, sep string, n int) []string { return genSplit(s, sep, len(sep), n) }

func FieldsFunc(s string, f func(rune) bool) []string
func Fields(s string) []string

func Join(a []string, sep string) string

func Repeat(s string, count int) string
```

## 转换(9个函数)

映射, 替换

到大写, 到小写, 到标题, 到大写特殊, 到小写特殊, 到标题特殊, 标题

其中7个是大小写的

```
F: Map, Replace
F: ToUpper, ToLower, ToTitle, ToUpperSpecial, ToLowerSpecial, ToTitleSpecial, Title
```

```go
func Map(mapping func(rune) rune, s string) string 

func ToUpper(s string) string
func ToLower(s string) string
func ToTitle(s string) string

func ToUpperSpecial(c unicode.SpecialCase, s string) string 
func ToLowerSpecial(c unicode.SpecialCase, s string) string 
func ToTitleSpecial(c unicode.SpecialCase, s string) string 

func Title(s string) string 

func Replace(s, old, new string, n int) string 
```

## 清理(9个函数)

除左Func, 除右Func, 除Func, 除左, 除右, 除, 除空白, 除头, 除尾 

```
F: TrimLeftFunc, TrimRightFunc, TrimFunc
F: TrimLeft, TrimRight, Trim
F: TrimSpace, TrimPrefix, TrimSuffix
```

```go
func TrimLeftFunc(s string, f func(rune) bool) string
func TrimRightFunc(s string, f func(rune) bool) string
func TrimFunc(s string, f func(rune) bool) string {
    return TrimRightFunc(TrimLeftFunc(s, f), f)
}
func TrimLeft(s string, cutset string) string 
func TrimRight(s string, cutset string) string
func Trim(s string, cutset string) string 

func TrimSpace(s string) string {
    return TrimFunc(s, unicode.IsSpace)
}
func TrimPrefix(s, prefix string) string {
    if HasPrefix(s, prefix) {
        return s[len(prefix):]
    }
    return s
}
func TrimSuffix(s, suffix string) string {
    if HasSuffix(s, suffix) {
        return s[:len(s)-len(suffix)]
    }
    return s
}
```

## 比较(2个函数)

等于环绕, 比较

```
F: EqualFold, Compare
```

```go
func EqualFold(s, t string) bool
func Compare(a, b string) int
```

## 读类型

实现了6个IO接口: 读, 找, 写到, 读在, 字节扫描, 符文扫描

还有3个函数: 长度, 大小, 重置

长度返回的是未读部分的长度, 大小返回的是原本字符串的长度

NewReader returns a new Reader reading from s. It is similar to bytes.NewBufferString but more efficient and read-only. 

```
T: Reader
    r.I: R, S, WT, RA, ByteScanner, RuneScanner
    r.F: Len, Size, Reset
newF: NewReader
```

```go
type Reader struct {
    s        string
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
func (r *Reader) Reset(s string) { *r = Reader{s, 0, -1} }

func NewReader(s string) *Reader { return &Reader{s, 0, -1} }

```

## Replacer类型

是对replacer接口的封装, replacer包含2个函数: 替换, 写字符串, 其本身就是一个replacer接口

我想, 封装接口的类型, 其实都可以本身就是该接口, 或者説本就原本就是, 只要同名函数封装一下内部接口的执行就可以了.如下.

那么这么封装一下的价值是什么呢?

因为接口必须依附于类型来实现.这样就有点像类型转换, 将接口转换成了类型.

```
T: Replacer
    r.i: replacer
newF: NewReplacer
```

```go
type Replacer struct {
    r replacer
}

type replacer interface {
    Replace(s string) string
    WriteString(w io.Writer, s string) (n int, err error)
}

func (r *Replacer) Replace(s string) string {
    return r.r.Replace(s)
}

func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error) {
    return r.r.WriteString(w, s)
}

func NewReplacer(oldnew ...string) *Replacer 
```

## Builder类型

实现了2个IO接口: 写, 字节写

实现了2个写的函数: 写符文, 写字符串

实现了4个其它函数: 字符串化, 长度, 生长, 重置

竟然没有New函数, 和bytes.Buffer一样, 可以直接new(strings.Builder)

```go
T: Builder
    b.I: W, ByteWriter
    b.F: WriteRune, WriteString
    b.F: String, Len, Grow, Reset
```

```go
type Builder struct {
    addr *Builder // of receiver, to detect copies by value
    buf  []byte
}

func (b *Builder) Write(p []byte) (int, error)
func (b *Builder) WriteByte(c byte) error 
func (b *Builder) WriteRune(r rune) (int, error)
func (b *Builder) WriteString(s string) (int, error)

func (b *Builder) String() string
func (b *Builder) Len() int
func (b *Builder) Grow(n int) 
func (b *Builder) Reset()
```

# Exercise

[strings.func.go](https://github.com/iofxl/gogogo/blob/master/strings.func.go) [strings.again.go](https://github.com/iofxl/gogogo/blob/master/strings.again.go)



