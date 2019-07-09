# Character Sets and Encodings

# Links

* [Definitions | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/encoding/definitions.html)
* [基础知识 - 字符编码简介 - GoLove - 博客园](https://www.cnblogs.com/golove/p/3222096.html)


# Definitions

**Character**

A character is a “unit of information that roughly corresponds to a grapheme
(written symbol) of a natural language, such as a letter, numeral, or
punctuation mark” (Wikipedia). A character is “the smallest component of
written language that has a semantic value” (Unicode). 


A character is some sort of abstraction of any actual symbol: the character “a”
is to any written “a” as a Platonic circle is to any actual circle. The concept
of character also includes control characters, which do not correspond to
natural language symbols but to other bits of information used to process texts
of the language.


A character does not have any particular appearance, although we use the
appearance to help recognize the character. 

**Character Repertoire/Character Set**

A character repertoire is a set of distinct characters, such as the Latin
alphabet. 

**Character Code**

A character code is a mapping from characters to integers. The mapping for a
character set is also called a coded character set or code set. The value of
each character in this mapping is often called a code point. ASCII is a code
set. The code point for “a” is 97 and for “A” is 65 (decimal).

The character code is still an abstraction. It isn’t yet what we will see in
text files, or in TCP packets. However, it is getting close, as it supplies the
mapping from human-oriented concepts into numerical ones.

**Character Encoding**

To communicate or store a character, you need to encode it in some way. To
transmit a string, you need to encode all characters in the string. 

**Transport Encoding**

An encoding can be based on space- and hence bandwidth-saving techniques such
as zipping the text. Or it could be reduced to a 7-bit format to allow a parity
checking bit, such as base64.

**Unicode**


# UTF-8, Go, and Runes

Go uses UTF-8 encoded characters in its strings. Each character is of type
rune. This is an alias for int32. A Unicode character can be up to 4 bytes in
UTF-8 encoding so that 4 bytes are needed to represent all characters. In terms
of characters, a string is an array of runes using 1, 2, or 4 bytes per rune.

# UCS and Unicode

UCS（Universal Character Set，通用字符集）使用 4 个字节来表示一个字符，其中最高
字节的最高二进制位始终为 0，这样能够表示的字符数量就达到了约 20 亿个，足以容纳
全世界所有的字符。在 UCS 字符集中，0～255 这些数值的定义与 Latin1 一样，只不过
是用 4 个字节表示 1 个字符，除最低位字节外，其它 3 个字节全部用 0 填充。这样就
保证了与 Latin1 字符集的兼容。

UCS 将最高字节定义为 2^7=128 个组（group），每个组再根据次高字节分为 256 个平面
（plane），每个平面再根据次低字节分为 256 行（row），每行再根据最低字节分为 256
个码位（cell）。UCS 中第 0 组的第 0 平面被称为“基本多文种平面”，里面存放了包含
世界各国的常用字符。除了第 0 平面，UCS 还定义了 16 个辅助平面，用来存放更多的字
符。

Unicode 字符集是 UCS 字符集的子集，只包含 0x0～0x10FFFF 范围的字符。

UTF: Unicode/UCS Transformation Format

utf32, utf16:

linux, win:	Little-Endian
Mac:		Big-Endian

# Unicode Gotchas

* [norm - GoDoc](https://godoc.org/golang.org/x/text/unicode/norm)


