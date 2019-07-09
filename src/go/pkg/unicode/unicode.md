# [Package unicode](https://golang.org/pkg/unicode/)  

# Links

* [Golang学习 - unicode 包](http://www.cnblogs.com/golove/p/3273585.html)

# Unicode in Regexp

```go
\pN            Unicode character class (one-letter name)
\p{Greek}      Unicode character class
\PN            negated Unicode character class (one-letter name)
\P{Greek}      negated Unicode character class
```



# Constants

```go 
const (
        MaxRune         = '\U0010FFFF' // Maximum valid Unicode code point.
        ReplacementChar = '\uFFFD'     // Represents invalid code points.
        MaxASCII        = '\u007F'     // maximum ASCII value.
        MaxLatin1       = '\u00FF'     // maximum Latin-1 value.
)

const (
        UpperCase = iota
        LowerCase
        TitleCase
        MaxCase
)
```

# Variables

```go
var Categories = map[string]*RangeTable{
        "C":  C,
        "Cc": Cc,
        "Cf": Cf,
        "Co": Co,
        "Cs": Cs,
	...
}

// Letter, Mark, Number, punctuation, Symbol, Space separator

var GraphicRanges = []*RangeTable{
        L, M, N, P, S, Zs,
}

var PrintRanges = []*RangeTable{
        L, M, N, P, S,
}

var Scripts = map[string]*RangeTable{
        "Adlam":                  Adlam,
        "Ahom":                   Ahom,
	...
}
```
# Func

```go
func Is(rangeTab *RangeTable, r rune) bool

func In(r rune, ranges ...*RangeTable) bool
func IsOneOf(ranges []*RangeTable, r rune) bool

func IsControl(r rune) bool
func IsLetter(r rune) bool
func IsMark(r rune) bool
func IsNumber(r rune) bool
func IsPunct(r rune) bool
func IsSymbol(r rune) bool
func IsSpace(r rune) bool

func IsGraphic(r rune) bool
func IsPrint(r rune) bool
func IsDigit(r rune) bool

func IsTitle(r rune) bool
func IsUpper(r rune) bool
func IsLower(r rune) bool

func SimpleFold(r rune) rune
func To(_case int, r rune) rune
func ToLower(r rune) rune
func ToTitle(r rune) rune
func ToUpper(r rune) rune
```

E.G:

```go
fmt.Println(unicode.Is(unicode.L, '的'))    // true
fmt.Println(unicode.Is(unicode.M, '的'))    // false
fmt.Println(unicode.Is(unicode.N, '的'))    // false
fmt.Println(unicode.Is(unicode.P, '的'))    // false
fmt.Println(unicode.Is(unicode.S, '的'))    // false
fmt.Println(unicode.Is(unicode.Zs, '的'))   // false
fmt.Println(unicode.Is(unicode.Han, '的'))  // true
```



