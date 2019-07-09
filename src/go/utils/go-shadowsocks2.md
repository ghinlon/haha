# [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)

这是一个设计很美，代码实现很烂的项目。

# Overview

# Design

## type Ciper interface

这是这个包最美的东西，也是最核心的东西。有了这个东西，这个包就要一个好的骨架。不管代码实现的有多烂，整个项目都不会烂到哪里去的。

**M**: 不过以我目前的知识，我认为对udp的处理是有问题的。我认为udp不是加密一下就会得的，而是要转化为tcp去传的，

有这个接口的设计，余下来的工作就是去使东西实现这个接口就会得罢。

```go
type Cipher interface {
    StreamConnCipher
    PacketConnCipher
}

type PacketConnCipher interface {
    PacketConn(net.PacketConn) net.PacketConn
}

type StreamConnCipher interface {
    StreamConn(net.Conn) net.Conn
}

func Dial(network, address string, ciph StreamConnCipher) (net.Conn, error)
func Listen(network, address string, ciph StreamConnCipher) (net.Listener, error)
func ListenPacket(network, address string, ciph PacketConnCipher) (net.PacketConn, error)
```

# Implementation

不得不曰，为了实现上面的接口，这个项目的实现方式真滴是非常难瞅。

先瞅瞅这里是哪样实现的，俺再整理整理思路，瞅瞅哪样实现更是好。

# 问题

