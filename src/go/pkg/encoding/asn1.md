# [Package asn1](https://golang.org/pkg/encoding/asn1/)

<!-- ToC start -->

# Table of Contents

1. [Links](#links)
1. [Overview](#overview)
1. [Package files](#package-files)
1. [一对编解组函数](#两个基本的编解组函数)
1. [一末概念](#一末概念)
<!-- ToC end -->

# Links

* [ASN.1 - Abstract Syntax Notation](https://www.obj-sys.com/asn1tutorial/node4.html)
* [Introduction to ASN.1](https://www.itu.int/en/ITU-T/asn1/Pages/introduction.aspx)
* [ASN.1 | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/dataserialisation/asn1.html)
* [asn浅析--协议的设计-博客-云栖社区-阿里云](https://yq.aliyun.com/articles/494985?spm=a2c4e.11153940.blogcont494986.9.3bc82a74DRPHjg)
* [asn浅析--简单介绍-博客-云栖社区-阿里云](https://yq.aliyun.com/articles/494986)

# Overview

我本来以为这个包会实现encoding里的4个接口，否想到否实现

# Package files

＃ [Index](https://golang.org/pkg/encoding/asn1/#pkg-index)

# 一对编解组函数

```go
func Marshal(val interface{}) ([]byte, error)
func Unmarshal(val interface{}, b []byte) (rest []byte, err error)
```

编组东西,返回字节串跟错误

解组字节串进东西，返回剩余跟错误

# 一末概念

语法三元组：实际语法、抽象语法和传输语法

其实在早期的一些标准如ASCII，它们既定义了抽象语法（比如字母A），又定义了传输语法（0x41）。ASN.1分离了这两种概念，以便可以选择一种适合要求的编解码方法。

```
asn1以字节为基础，定义1个字节为“类型”和“长度”的元长度，首先考虑“类型”，从高位到低位的规定：8位和7位两位定义class，一共四种，第6位表示后面的“值”是基本类型还是复合类型，再后面的5位表示一个tag，如果我们把class想象成大的类型，那么tag就是该大类中的小类，但是如果5个字节不够了怎么办，实际上如果你第5位为1，那么就说明接下来的字节还是表示tag，直到最高位不为1为止结束tag；接下来就是“长度”了，一个字节仅能表示最多一个字节的长度，如果长度超过256怎么办？解决办法就是只要“长度”的最高位为1，那么此字节表示的就是“长度”的长度，这样就可以定义更长的长度了；接下来的“值”要么是一个确定的数字要么是上面三元组的重复。
```

[asn1.go][1]

[1]: https://github.com/iofxl/practisego/blob/master/asn1.go

[ASN.1 Daytime Server][2]  [ASN.1 Daytime Client][3]

[2]: https://github.com/iofxl/practisego/blob/master/asndaytimes.go
[3]: https://github.com/iofxl/practisego/blob/master/asndaytimec.go


