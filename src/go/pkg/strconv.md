# [Package strconv](https://golang.org/pkg/strconv/)

<!-- ToC start -->
# Table of Contents

1. [Links](#links)
1. [Package files](#package-files)
1. [Abstract](#abstract)
1. [Overview](#overview)
1. [Organize](#organize)
   1. [Int (8个函数)](#int-8个函数)
   1. [Float (3个函数)](#float-3个函数)
   1. [Bool (3个函数)](#bool-3个函数)
   1. [Quote (17个函数)](#quote-17个函数)
   1. [NumError类型](#numerror类型)
1. [Exercise](#exercise)
<!-- ToC end -->

# Links

* [Golang学习 - strconv 包 - GoLove - 博客园](https://www.cnblogs.com/golove/p/3262925.html)

# Package files

[atoi.go](https://golang.org/src/strconv/atoi.go) [itoa.go](https://golang.org/src/strconv/itoa.go)

# Abstract

```
// 8
F: ParseUint, ParseInt, Atoi
F: FormatUint, FormatInt, Itoa, AppendUint, AppendInt

// 3
F: ParseFloat, FormatFloat, AppendFloat

// 3
F: ParseBool, FormatBool, AppendBool

// 17
F: Quote, QuoteToASCII, QuoteToGraphic, AppendQuote, AppendQuoteToASCII, AppendQuoteToGraphic
F: QuoteRune, QuoteRuneToASCII, QuoteRuneToGraphic, AppendQuoteRune, AppendQuoteRuneToASCII, AppendQuoteRuneToGraphic
F: Unquote, UnquoteChar
F: CanBackquote, IsPrint, IsGraphic

T: NumError
    e.I: error
```

# Overview

Package strconv implements conversions to and from string representations of basic data types. 

# Organize

## Int (8个函数)

```
F: ParseUint, ParseInt, Atoi
F: FormatUint, FormatInt, Itoa, AppendUint, AppendInt
```

```go
func ParseUint(s string, base int, bitSize int) (uint64, error) 
func ParseInt(s string, base int, bitSize int) (i int64, err error)
// Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
func Atoi(s string) (int, error) 

func FormatUint(i uint64, base int) string
func FormatInt(i int64, base int) string
// Itoa is shorthand for FormatInt(int64(i), 10).
func Itoa(i int) string 

func AppendUint(dst []byte, i uint64, base int) []byte
func AppendInt(dst []byte, i int64, base int) []byte
```

## Float (3个函数)

```
F: ParseFloat, FormatFloat, AppendFloat
```

```go
func ParseFloat(s string, bitSize int) (float64, error)
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
```

## Bool (3个函数)

```
F: ParseBool, FormatBool, AppendBool
```

```go
func ParseBool(str string) (bool, error) 
func FormatBool(b bool) string 
func AppendBool(dst []byte, b bool) []byte 
```

## Quote (17个函数)

```
F: Quote, QuoteToASCII, QuoteToGraphic, AppendQuote, AppendQuoteToASCII, AppendQuoteToGraphic
F: QuoteRune, QuoteRuneToASCII, QuoteRuneToGraphic, AppendQuoteRune, AppendQuoteRuneToASCII, AppendQuoteRuneToGraphic
F: Unquote, UnquoteChar
F: CanBackquote, IsPrint, IsGraphic
```

```go
func Quote(s string) string
func QuoteToASCII(s string) string
func QuoteToGraphic(s string) string

func AppendQuote(dst []byte, s string) []byte
func AppendQuoteToASCII(dst []byte, s string) []byte 
func AppendQuoteToGraphic(dst []byte, s string) []byte 

func QuoteRune(r rune) string 
func QuoteRuneToASCII(r rune) string
func QuoteRuneToGraphic(r rune) string

func AppendQuoteRune(dst []byte, r rune) []byte
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte 
func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte 

func UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)
func Unquote(s string) (string, error) 

func CanBackquote(s string) bool 
func IsPrint(r rune) bool
func IsGraphic(r rune) bool
```

## NumError类型

```
T: NumError
    e.I: error
```

```go
type NumError struct {
    Func string // the failing function (ParseBool, ParseInt, ParseUint, ParseFloat)
    Num  string // the input
    Err  error  // the reason the conversion failed (e.g. ErrRange, ErrSyntax, etc.)
}

func (e *NumError) Error() string 
```

# Exercise
