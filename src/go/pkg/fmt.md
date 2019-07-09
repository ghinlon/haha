# [Package fmt](https://golang.org/pkg/fmt/)

<!-- ToC start -->
# Table of Contents

1. [Links](#links)
1. [Package files](#package-files)
1. [Abstract](#abstract)
1. [Organize](#organize)
   1. [9种Print函数](#9种print函数)
   1. [1个转换成error类型的函数](#1个转换成error类型的函数)
   1. [9种scan函数](#9种scan函数)
   1. [4个print相关接口](#4个print相关接口)
   1. [2个scan相关接口](#2个scan相关接口)
1. [Exercise](#exercise)
<!-- ToC end -->

# Links

[基础知识 - Golang 中的格式化输入输出](https://www.cnblogs.com/golove/p/3284304.html)  
[Golang学习 - fmt 包](https://www.cnblogs.com/golove/p/3286303.html)  
[标准库 - fmt/format.go 解读](https://www.cnblogs.com/golove/p/5861971.html)  
[标准库 - fmt/print.go 解读](https://www.cnblogs.com/golove/p/5857338.html)  
[标准库 - fmt/scan.go 解读](https://www.cnblogs.com/golove/p/5888441.html)  

# Package files

[doc.go](https://golang.org/src/fmt/doc.go) [format.go](https://golang.org/src/fmt/format.go) [print.go](https://golang.org/src/fmt/print.go) [scan.go](https://golang.org/src/fmt/scan.go)

# Abstract

```
F: Fprintf, Fprint, Fprintln, Printf, Print, Println, Sprintf, Sprint, Sprintln  
F: Errorf
F: Fscanf, Fscan, Fscanln, Scanf, Scan, Scanln, Sscanf, Sscan, Sscanln   
I: State, Formatter, String, GoString  
I: ScanState, Scanner  
```

# Organize

## 9种Print函数

`F: Fprintf, Fprint, Fprintln, Printf, Print, Println, Sprintf, Sprint, Sprintln`  

* `Fprintf/Fprint/Fprintln`的写入目标是写接口`w`  
* `Printf/Print/Println`就是`Fprint{f,,ln}`的特例,指定写給`os.Stdout`  
* `Sprintf/Sprint/Sprintln`则以字符串类型返回  

`printf, print, println` 这三种的区别:  

* `Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.`   
* `Fprint formats using the default formats for its operands and writes to w. Spaces are added between operands when neither is a string. It returns the number of bytes written and any write error encountered.`   
* `Fprintln formats using the default formats for its operands and writes to w. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.`   

**简单的説:** `print` 非字符串参数之间添加空格; `println` 所有参数之间添加空格，结尾添加换行符

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintf(format, a)
    n, err = w.Write(p.buf)
    p.free()
    return
}

func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}

func Fprint(w io.Writer, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrint(a)
    n, err = w.Write(p.buf)
    p.free()
    return
}

func Print(a ...interface{}) (n int, err error) {
    return Fprint(os.Stdout, a...)
}

func Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
    p := newPrinter()
    p.doPrintln(a)
    n, err = w.Write(p.buf)
    p.free()
    return
}

func Println(a ...interface{}) (n int, err error) {
    return Fprintln(os.Stdout, a...)
}

func Sprintf(format string, a ...interface{}) string {
    p := newPrinter()
    p.doPrintf(format, a)
    s := string(p.buf)
    p.free()
    return s
}


func Sprint(a ...interface{}) string {
    p := newPrinter()
    p.doPrint(a)
    s := string(p.buf)
    p.free()
    return s
}

func Sprintln(a ...interface{}) string {
    p := newPrinter()
    p.doPrintln(a)
    s := string(p.buf)
    p.free()
    return s
}

```

## 1个转换成error类型的函数

`F: Errorf` 

`Errorf`是对`Sprintf`返回结果的再次转换,转换成`error`类型返回

```go
func Errorf(format string, a ...interface{}) error {
    return errors.New(Sprintf(format, a...))
}
```

## 9种scan函数

`F: Fscanf, Fscan, Fscanln, Scanf, Scan, Scanln, Sscanf, Sscan, Sscanln`   

* `Fscanf/Fscan/Fscanln`从读接口`r`中扫描  
* `Scanf/Scan/Scanln`就是`Fscan{f,,ln}`的特例,指定从`os.Stdin`中扫描  
* `Sscanf/Sscan/Sscanln`也是对应`Fscan{f,,ln}`的特例,指定从字符串`str`中扫描

`scanf, scan, scanln` 这三种的区别:  

**简单的説:** `scan` 换行视为空白, `scanln` 换行结束解析

```go
func Scan(a ...interface{}) (n int, err error) {
    return Fscan(os.Stdin, a...)
}

func Scanln(a ...interface{}) (n int, err error) {
    return Fscanln(os.Stdin, a...)
}

func Scanf(format string, a ...interface{}) (n int, err error) {
    return Fscanf(os.Stdin, format, a...)
}

func Sscan(str string, a ...interface{}) (n int, err error) {
    return Fscan((*stringReader)(&str), a...)
}

func Sscanln(str string, a ...interface{}) (n int, err error) {
    return Fscanln((*stringReader)(&str), a...)
}

func Sscanf(str string, format string, a ...interface{}) (n int, err error) {
    return Fscanf((*stringReader)(&str), format, a...)
}

func Fscan(r io.Reader, a ...interface{}) (n int, err error) {
    s, old := newScanState(r, true, false)
    n, err = s.doScan(a)
    s.free(old)
    return
}

func Fscanln(r io.Reader, a ...interface{}) (n int, err error) {
    s, old := newScanState(r, false, true)
    n, err = s.doScan(a)
    s.free(old)
    return
}

func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error) {
    s, old := newScanState(r, false, false)
    n, err = s.doScanf(format, a)
    s.free(old)
    return
}
```

## 4个print相关接口

`I: State, Formatter, String, GoString`  

State接口是一个写接口的扩展  
Formatter接口是操作State接口的接口,实现该接口可以自定义格式,这是一个格式接口  
某种类型只要实现了Stringer, GoStringer接口,就可以被输出该类型的字符串格式  


```go
type State interface {
    Write(b []byte) (n int, err error)
    Width() (wid int, ok bool)
    Precision() (prec int, ok bool)
    Flag(c int) bool
}

type Formatter interface {
    Format(f State, c rune)
}

type Stringer interface {
    String() string
}

type GoStringer interface {
    GoString() string
}
```

## 2个scan相关接口

`I: ScanState, Scanner`

Scanner 用于让自定义类型实现自己的扫描过程。  
Scan 方法会从输入端读取数据并将处理结果存入接收端，接收端必须是有效的指针。  
Scan 方法会被扫描器调用，只要对应的 arg 实现了该方法。  

```go
type ScanState interface {
        // ReadRune reads the next rune (Unicode code point) from the input.
        // If invoked during Scanln, Fscanln, or Sscanln, ReadRune() will
        // return EOF after returning the first '\n' or when reading beyond
        // the specified width.
        ReadRune() (r rune, size int, err error)
        // UnreadRune causes the next call to ReadRune to return the same rune.
        UnreadRune() error
        // SkipSpace skips space in the input. Newlines are treated appropriately
        // for the operation being performed; see the package documentation
        // for more information.
        SkipSpace()
        // Token skips space in the input if skipSpace is true, then returns the
        // run of Unicode code points c satisfying f(c).  If f is nil,
        // !unicode.IsSpace(c) is used; that is, the token will hold non-space
        // characters. Newlines are treated appropriately for the operation being
        // performed; see the package documentation for more information.
        // The returned slice points to shared data that may be overwritten
        // by the next call to Token, a call to a Scan function using the ScanState
        // as input, or when the calling Scan method returns.
        Token(skipSpace bool, f func(rune) bool) (token []byte, err error)
        // Width returns the value of the width option and whether it has been set.
        // The unit is Unicode code points.
        Width() (wid int, ok bool)
        // Because ReadRune is implemented by the interface, Read should never be
        // called by the scanning routines and a valid implementation of
        // ScanState may choose always to return an error from Read.
        Read(buf []byte) (n int, err error)
}

type Scanner interface {
        Scan(state ScanState, verb rune) error
}
```

# Exercise


