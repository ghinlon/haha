# [Package utf8](https://golang.org/pkg/unicode/utf8/)  

# Links

* [Golang学习 - unicode/utf8 包](https://www.cnblogs.com/golove/p/3271597.html)
* [标准库 - unicode/utf8/utf8.go 解读](http://www.cnblogs.com/golove/p/5889790.html)
* [基础知识 - 字符编码简介 - GoLove - 博客园](https://www.cnblogs.com/golove/p/3222096.html)
* [encoding - GoDoc](https://godoc.org/golang.org/x/text/encoding)


# UTF8 Encoding

1. 对于单字节的符号，字节的第一位设为0，后面7位为这个符号的unicode码。因此对于
   英语字母，UTF-8编码和ASCII码是相同的。
2. 对于n字节的符号（n>1），第一个字节的前n位都设为1，第n+1位设为0，后面字节的前
   两位一律设为10。剩下的没有提及的二进制位，全部为这个符号的unicode码。

如表：

**utf8 encoded char's first byte MUST NOT begin with 10xxxxxx(bit)**

```
1字节 0xxxxxxx
2字节 110xxxxx 10xxxxxx
3字节 1110xxxx 10xxxxxx 10xxxxxx
4字节 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
5字节 111110xx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
6字节 1111110x 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
```

因此UTF-8中可以用来表示字符编码的实际位数最多有31位，即上表中x所表示的位。除去
那些控制位（每字节开头的10等），这些x表示的位与UNICODE编码是一一对应的，位高低
顺序也相同。


# Constants

```go
const (
        RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
        RuneSelf  = 0x80         // characters below Runeself are represented as themselves in a single byte.
        MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
        UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
)

const (                                                                                                                                                                                
        t1 = 0x00 // 0000 0000                                                                                                                                                         
        tx = 0x80 // 1000 0000                                                                                                                                                         
        t2 = 0xC0 // 1100 0000                                                                                                                                                         
        t3 = 0xE0 // 1110 0000                                                                                                                                                         
        t4 = 0xF0 // 1111 0000                                                                                                                                                         
        t5 = 0xF8 // 1111 1000                                                                                                                                                         
                                                                                                                                                                                       
        maskx = 0x3F // 0011 1111                                                                                                                                                      
        mask2 = 0x1F // 0001 1111                                                                                                                                                      
        mask3 = 0x0F // 0000 1111                                                                                                                                                      
        mask4 = 0x07 // 0000 0111                                                                                                                                                      
                                                                                                                                                                                       
        rune1Max = 1<<7 - 1                                                                                                                                                            
        rune2Max = 1<<11 - 1                                                                                                                                                           
        rune3Max = 1<<16 - 1                                                                                                                                                           
                                                                                                                                                                                       
        // The default lowest and highest continuation byte.                                                                                                                           
        locb = 0x80 // 1000 0000                                                                                                                                                       
        hicb = 0xBF // 1011 1111                                                                                                                                                       
                                                                                                                                                                                       
        // These names of these constants are chosen to give nice alignment in the                                                                                                     
        // table below. The first nibble is an index into acceptRanges or F for                                                                                                        
        // special one-byte cases. The second nibble is the Rune length or the                                                                                                         
        // Status for the special one-byte case.                                                                                                                                       
        xx = 0xF1 // invalid: size 1                                                                                                                                                   
        as = 0xF0 // ASCII: size 1                                                                                                                                                     
        s1 = 0x02 // accept 0, size 2                                                                                                                                                  
        s2 = 0x13 // accept 1, size 3                                                                                                                                                  
        s3 = 0x03 // accept 0, size 3                                                                                                                                                  
        s4 = 0x23 // accept 2, size 3                                                                                                                                                  
        s5 = 0x34 // accept 3, size 4                                                                                                                                                  
        s6 = 0x04 // accept 0, size 4                                                                                                                                                  
        s7 = 0x44 // accept 4, size 4                                                                                                                                                  
) 
```

# Var

```go
// first is information about the first byte in a UTF-8 sequence.                                                                                                                      
var first = [256]uint8{                                                                                                                                                                
        //   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F                                                                                                                 
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x00-0x0F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x10-0x1F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x20-0x2F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x30-0x3F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x40-0x4F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x50-0x5F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x60-0x6F                                                                                                   
        as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, as, // 0x70-0x7F                                                                                                   
        //   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F                                                                                                                 
        xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0x80-0x8F                                                                                                   
        xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0x90-0x9F                                                                                                   
        xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xA0-0xAF                                                                                                   
        xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xB0-0xBF                                                                                                   
        xx, xx, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, // 0xC0-0xCF                                                                                                   
        s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, s1, // 0xD0-0xDF                                                                                                   
        s2, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s3, s4, s3, s3, // 0xE0-0xEF                                                                                                   
        s5, s6, s6, s6, s7, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, xx, // 0xF0-0xFF                                                                                                   
}                                                                                                                                                                                      
                                                                                                                                                                                       
// acceptRange gives the range of valid values for the second byte in a UTF-8                                                                                                          
// sequence.                                                                                                                                                                           
type acceptRange struct {                                                                                                                                                              
        lo uint8 // lowest value for second byte.                                                                                                                                      
        hi uint8 // highest value for second byte.                                                                                                                                     
}                                                                                                                                                                                      
                                                                                                                                                                                       
var acceptRanges = [...]acceptRange{                                                                                                                                                   
        0: {locb, hicb},                                                                                                                                                               
        1: {0xA0, hicb},                                                                                                                                                               
        2: {locb, 0x9F},                                                                                                                                                               
        3: {0x90, hicb},                                                                                                                                                               
        4: {locb, 0x8F},                                                                                                                                                               
}
```

16进制的1位相当于2进制的4位{0xF=1111(2)=15(10)},16进制1位可以表达16个值,2进制要
4位来表达 16个值.

2进制是每8位一个字节,所以16进制是每2位一字节.2进制的8位可表达的值,在16进制下用2
位就可以表达完 8位的值:0-255,最大值是255

假如第一个字节的值是127,对应二进制 0111 1111, 对应16进制7F(把2进制换算到16进制,
感觉上还是像换到10进制一样), 对应 first表里值是最后一个as,就是最大的那个ascii码

假如0xfb = 128 = 1000 0000 = 0x80 --> first[128] == xx,表明不是合法的utf8編码该
有的值.

根据上面的算法,一直到1011 1111 = 0xbf = 191 (下一个1100 0000 就合法了)都是不合
法的,但first表里一直到0xc0,0xc1也都是不合法的. 1100 0000, 11000001这两个为什么
不合法呢?

然后, 0xfb = 1110 0000 = 0xe0(8421 - > 14 -9 =5 ->e (11100000,e0,224), 作为
first 的索引得到值 s2 = 0x13,作为acceptRanges的索引,得到第二字节的值范围
{0xA0,hicb}, 开始3字节的編码

# Func

```go
func DecodeRune(p []byte) (r rune, size int)
func DecodeRuneInString(s string) (r rune, size int)
func DecodeLastRune(p []byte) (r rune, size int)
func DecodeLastRuneInString(s string) (r rune, size int)

// EncodeRune writes into p (which must be large enough) the UTF-8 encoding of the rune.
// It returns the number of bytes written.
func EncodeRune(p []byte, r rune) int
func RuneLen(r rune) int
func RuneCount(p []byte) int
func RuneCountInString(s string) (n int)

func FullRune(p []byte) bool
func FullRuneInString(s string) bool
func RuneStart(b byte) bool
func Valid(p []byte) bool
func ValidRune(r rune) bool
func ValidString(s string) bool
```
```go
func FullRune(p []byte) bool {
	n := len(p)
	if n == 0 {
		return false
	}
	x := first[p[0]]
	// x&7 means x&0111(2) get x's little-ending 4 bits.
	if n >= int(x&7) {
		return true // ASCII, invalid or valid.
	}
	// Must be short or invalid.
	accept := acceptRanges[x>>4]
	// invalid is true
	if n > 1 && (p[1] < accept.lo || accept.hi < p[1]) {
		return true
	// invalid is true
	} else if n > 2 && (p[2] < locb || hicb < p[2]) {
		return true
	}
	// now must be short
	return false
}
```
