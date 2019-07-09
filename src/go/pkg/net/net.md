# [Package net](https://golang.org/pkg/net/)

<!-- ToC start -->

# Table of Contents

1. [Links](#links)
1. [Package files](#package-files)
1. [Overview](#overview)
1. [5个基础网络接口: Conn, Listener, PacketConn, Addr, Error](#5个基础网络接口:-conn-listener-packetconn-addr-error)
   1. [Conn接口](#conn接口)
   1. [Listener接口](#listener接口)
   1. [PacketConn接口](#packetconn接口)
   1. [Addr接口](#addr接口)
   1. [关于网络类型](#关于网络类型)
   1. [Error接口](#error接口)
1. [4个基础网络结构: Interface, HardwareAddr, IP, IPMask](#4个基础网络结构:-interface-hardwareaddr-ip-ipmask)
   1. [Interface结构](#interface结构)
   1. [HardwareAddr结构](#hardwareaddr结构)
   1. [IP结构](#ip结构)
   1. [IPMask结构](#ipmask结构)
1. [TCP协议里3个结构：TCPAddr, TCPConn, TCPListener](#tcp协议里3个结构：tcpaddr-tcpconn-tcplistener)
   1. [TCPAddr结构](#tcpaddr结构)
   1. [TCPConn结构](#tcpconn结构)
   1. [TCPListener结构](#tcplistener结构)
1. [UDP协议里2个结构：UDPAddr, UDPConn](#udp协议里2个结构：udpaddr-udpconn)
   1. [UDPAddr结构](#udpaddr结构)
   1. [UDPConn结构](#udpconn结构)
1. [引出Conn, PacketConn, Listener接口的小总结](#引出conn-packetconn-listener接口的小总结)
   1. [重写前头的例子](#重写前头的例子)
1. [IP协议里3个结构：IPNet, IPAddr, IPConn](#ip协议里3个结构：ipnet-ipaddr-ipconn)
   1. [IPNet结构](#ipnet结构)
   1. [IPAddr结构](#ipaddr结构)
   1. [IPConn结构](#ipconn结构)
1. [Unix协议里3个结构：UnixAddr, UnixConn, UnixListener](#unix协议里3个结构：unixaddr-unixconn-unixlistener)
   1. [UnixAddr结构](#unixaddr结构)
   1. [UnixConn结构](#unixconn结构)
   1. [UnixListener结构](#unixlistener结构)
1. [Dialer结构](#dialer结构)
1. [ListenConfig结构](#listenconfig结构)
1. [DNS的3个记录结构：SRV, MX, NS](#dns的3个记录结构：srv-mx-ns)
1. [Resolver结构](#resolver结构)
1. [DNS的9个查找函数](#dns的9个查找函数)
1. [6种网络错误结构](#6种网络错误结构)
1. [Buffers结构](#buffers结构)
1. [2个拆合主机端口的函数](#2个拆合主机端口的函数)
<!-- ToC end -->

# Links

* [Socket-level Programming | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/index.html)
* [Implementing UDP vs TCP in Golang - Mina Andrawos](http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/)
* [networking - When is it appropriate to use UDP instead of TCP? - Stack Overflow](https://stackoverflow.com/questions/1099672/when-is-it-appropriate-to-use-udp-instead-of-tcp/1099734#1099734)
* [Graham King » Raw sockets in Go: IP layer](https://www.darkcoding.net/uncategorized/raw-sockets-in-go-ip-layer/)
* [TCP header ](https://github.com/grahamking/latency/blob/master/tcp.go)

# Package files

* 定义接口和錯誤类型

    [net.go](https://golang.org/src/net/net.go) 

* 基础类型

    [interface.go](https://golang.org/src/net/interface.go) [mac.go](https://golang.org/src/net/mac.go) [ip.go](https://golang.org/src/net/ip.go)  [dial.go](https://golang.org/src/net/dial.go) [lookup.go](https://golang.org/src/net/lookup.go) 
* 实现Conn, PacketConn, Listener 

    [iprawsock.go](https://golang.org/src/net/iprawsock.go) [tcpsock.go](https://golang.org/src/net/tcpsock.go)  [udpsock.go](https://golang.org/src/net/udpsock.go) [unixsock.go ](https://golang.org/src/net/unixsock.go) 

# Overview

[Overview](https://golang.org/pkg/net/#pkg-overview)

# 5个基础网络接口: Conn, Listener, PacketConn, Addr, Error

## Conn接口

Conn接口是这个包的核心

Conn接口包含8个函数，实现了Conn接口也就实现了"读写关"接口

这8个函数返回Conn接口: Dial, DialTimeout, FileConn, Pipe, *Dialer.Dial, *Dialer.DialContext, *TCPListener.Accept, *UnixListener.Accept

包内有一个conn类型实现了Conn接口, 还实现了2个设置缓存函数和1个File函数

```go
type Conn interface {
        Read(b []byte) (n int, err error)
        Write(b []byte) (n int, err error)
        Close() error

        LocalAddr() Addr
        RemoteAddr() Addr

        // A deadline is an absolute time after which I/O operations
        // fail with a timeout (see type Error) instead of
        // blocking.
        SetDeadline(t time.Time) error
        SetReadDeadline(t time.Time) error
        SetWriteDeadline(t time.Time) error
}

func Dial(network, address string) (Conn, error)
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
func FileConn(f *os.File) (c Conn, err error)

func Pipe() (Conn, Conn)

func (d *Dialer) Dial(network, address string) (Conn, error)
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)

func (l *TCPListener) Accept() (Conn, error)
func (l *UnixListener) Accept() (Conn, error)

type conn struct {
    fd *netFD
}

func (c *conn) SetReadBuffer(bytes int) error
func (c *conn) SetWriteBuffer(bytes int) error
// On Unix systems this will cause the SetDeadline methods to stop working.
func (c *conn) File() (f *os.File, err error)
```

## Listener接口

```go
type Listener interface {
    // Accept waits for and returns the next connection to the listener.
    Accept() (Conn, error)
    // Close closes the listener.
    // Any blocked Accept operations will be unblocked and return errors.
    Close() error
    // Addr returns the listener's network address.
    Addr() Addr
}

func Listen(network, address string) (Listener, error)
func FileListener(f *os.File) (ln Listener, err error)

func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
```

## PacketConn接口

包含7个函数,区别于Conn接口, 其中的读写接口换成了另一种读写函数,没有了RemoteAddr函数

```go
type PacketConn interface {
    // ReadFrom reads a packet from the connection,
    // copying the payload into b. It returns the number of
    // bytes copied into b and the return address that
    // was on the packet.
    ReadFrom(b []byte) (n int, addr Addr, err error)
    // WriteTo writes a packet with payload b to addr.
    WriteTo(b []byte, addr Addr) (n int, err error)
    Close() error

    LocalAddr() Addr

    SetDeadline(t time.Time) error
    SetReadDeadline(t time.Time) error
    SetWriteDeadline(t time.Time) error
}

func ListenPacket(network, address string) (PacketConn, error)
func FilePacketConn(f *os.File) (c PacketConn, err error)

func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
```

## Addr接口

```go
// Addr represents a network end point address.
//
// The two methods Network and String conventionally return strings
// that can be passed as the arguments to Dial, but the exact form
// and meaning of the strings is up to the implementation.
type Addr interface {
    Network() string // name of the network (for example, "tcp", "udp")
    String() string  // string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80")
}

func (ifi *Interface) Addrs() ([]Addr, error)
func (ifi *Interface) MulticastAddrs() ([]Addr, error)
func InterfaceAddrs() ([]Addr, error)

func (n *IPNet) Network() string
func (n *IPNet) String() string

func (a *IPAddr) Network() string
func (a *IPAddr) String() string

func (c *IPConn) LocalAddr() Addr
func (c *IPConn) RemoteAddr() Addr
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)

func (a *TCPAddr) Network() string
func (a *TCPAddr) String() string

func (c *TCPConn) LocalAddr() Addr
func (c *TCPConn) RemoteAddr() Addr

func (l *TCPListener) Addr() Addr

func (a *UDPAddr) Network() string
func (a *UDPAddr) String() string

func (c *UDPConn) LocalAddr() Addr
func (c *UDPConn) RemoteAddr() Addr
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)

func (a *UnixAddr) Network() string
func (a *UnixAddr) String() string

func (c *UnixConn) LocalAddr() Addr
func (c *UnixConn) RemoteAddr() Addr
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error)

func (l *UnixListener) Addr() Addr
```

## 关于网络类型

```go
func Dial(network, address string) (Conn, error)
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
func (d *Dialer) Dial(network, address string) (Conn, error)
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)


func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)

func Listen(network, address string) (Listener, error)
func ListenPacket(network, address string) (PacketConn, error)

func ResolveIPAddr(network, address string) (*IPAddr, error)
func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
func ListenIP(network string, laddr *IPAddr) (*IPConn, error)

func ResolveTCPAddr(network, address string) (*TCPAddr, error)
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)

func ResolveUDPAddr(network, address string) (*UDPAddr, error)
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)

func ResolveUnixAddr(network, address string) (*UnixAddr, error)
func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error)
func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)

func LookupPort(network, service string) (port int, err error)
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)
```

Known networks are:

* "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only) 
* "udp", "udp4" (IPv4-only), "udp6" (IPv6-only)
* "ip", "ip4" (IPv4-only), "ip6" (IPv6-only)
* "unix", "unixgram" and "unixpacket". 

对于IP协议,参数network是要带上协议名,或协议号的.例:

```
Dial("ip4:1", "192.0.2.1")
Dial("ip6:ipv6-icmp", "2001:db8::1")
Dial("ip6:58", "fe80::1%lo0")
```

Resovle用正确的网络名生成各协议地址结构

IPNet, IPAddr, TCPAddr, UDPAddr, UnixAddr 5种地址结构都实现了Addr接口,可以通过Addr接口得到网络类型

## Error接口

```go
// An Error represents a network error.
type Error interface {
    error
    Timeout() bool   // Is the error a timeout?
    Temporary() bool // Is the error temporary?
}
```

# 4个基础网络结构: Interface, HardwareAddr, IP, IPMask

HardwareAddr, IP, IPMask all are []byte

## Interface结构

```go
type Interface struct {
    Index        int          // positive integer that starts at one, zero is never used
    MTU          int          // maximum transmission unit
    Name         string       // e.g., "en0", "lo0", "eth0.100"
    HardwareAddr HardwareAddr // IEEE MAC-48, EUI-48 and EUI-64 form
    Flags        Flags        // e.g., FlagUp, FlagLoopback, FlagMulticast
}

func (ifi *Interface) Addrs() ([]Addr, error) 
func (ifi *Interface) MulticastAddrs() ([]Addr, error)

func InterfaceByIndex(index int) (*Interface, error) 
func InterfaceByName(name string) (*Interface, error) 
func Interfaces() ([]Interface, error)
func InterfaceAddrs() ([]Addr, error)

type Flags uint
func (f Flags) String() string 
```

## HardwareAddr结构

```go
type HardwareAddr []byte

func (a HardwareAddr) String() string
func ParseMAC(s string) (hw HardwareAddr, err error)
```

## IP结构

这些函数返回IP结构:

ParseIP, ParseCIDR, IPv4, LookupIP, IP.To4, IP.To16, ip.Mask

```go
type IP []byte

func ParseIP(s string) IP 
func ParseCIDR(s string) (IP, *IPNet, error) 
func IPv4(a, b, c, d byte) IP
func LookupIP(host string) ([]IP, error)

func (ip IP) Equal(x IP) bool
func (ip IP) IsUnspecified() bool {
    return ip.Equal(IPv4zero) || ip.Equal(IPv6unspecified)
}
func (ip IP) IsLoopback() bool 

func (ip IP) IsMulticast() bool
func (ip IP) IsInterfaceLocalMulticast() bool
func (ip IP) IsLinkLocalMulticast() bool 
func (ip IP) IsLinkLocalUnicast() bool 
func (ip IP) IsGlobalUnicast() bool 

func (ip IP) To4() IP 
func (ip IP) To16() IP 

func (ip IP) DefaultMask() IPMask
func (ip IP) Mask(mask IPMask) IP

func (ip IP) String() string 
func (ip IP) MarshalText() ([]byte, error)
func (ip *IP) UnmarshalText(text []byte) error 
```

## IPMask结构

```go
// An IP mask is an IP address.
type IPMask []byte

func IPv4Mask(a, b, c, d byte) IPMask
func CIDRMask(ones, bits int) IPMask 

func (m IPMask) Size() (ones, bits int) 
func (m IPMask) String() string

func (ip IP) DefaultMask() IPMask
```


# TCP协议里3个结构：TCPAddr, TCPConn, TCPListener

## TCPAddr结构

TCPAddr结构实现了Addr接口

```go
type TCPAddr struct {
    IP   IP
    Port int
    Zone string // IPv6 scoped addressing zone
}

func (a *TCPAddr) Network() string { return "tcp" }
func (a *TCPAddr) String() string

func ResolveTCPAddr(network, address string) (*TCPAddr, error)
```

## TCPConn结构

TCPConn is an implementation of the Conn interface for TCP network connections.

TCPConn内部就是一个conn结构，所以自然它实现了Conn接口

```go
type TCPConn struct {
    conn
}

func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)

// SyscallConn returns a raw network connection.
// This implements the syscall.Conn interface.
func (c *TCPConn) SyscallConn() (syscall.RawConn, error)

func (c *TCPConn) CloseRead() error
func (c *TCPConn) CloseWrite() error 

// 也是conn结构实现的
func (c *TCPConn) File() (f *os.File, err error)

func (c *TCPConn) SetKeepAlive(keepalive bool) error 
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
func (c *TCPConn) SetLinger(sec int) error
func (c *TCPConn) SetNoDelay(noDelay bool) error

// DialTCP acts like Dial for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.// DialTCP acts like Dial for TCP networks.
//
// The network must be a TCP network name; see func Dial for details.
//
// If laddr is nil, a local address is automatically chosen.
// If the IP field of raddr is nil or an unspecified IP address, the
// local system is assumed.
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error) 
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
```

客户端的基本流程就是Dial建立连接，再么写，再么读。

[tcpc.go][1]

## TCPListener结构

实现Listener接口

```go
type TCPListener struct {
    fd *netFD
}

// Accept implements the Accept method in the Listener interface; it
// waits for the next call and returns a generic Conn.
func (l *TCPListener) Accept() (Conn, error)
func (l *TCPListener) Close() error
func (l *TCPListener) Addr() Addr { return l.fd.laddr }

// AcceptTCP accepts the next incoming call and returns the new
// connection.
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
func (l *TCPListener) File() (f *os.File, err error)
func (l *TCPListener) SetDeadline(t time.Time) error
func (l *TCPListener) SyscallConn() (syscall.RawConn, error) 

func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
```

[Daytime Server][2]

[Echo Server][3]

困扰我的是，怎样知道把客户端发过来的东西读光罢呢？: 不需要知道读光否读光的，使一个固定长的字节串 只要读的过程否碰的错误 一直读一直写转就可以罢

# UDP协议里2个结构：UDPAddr, UDPConn

## UDPAddr结构

UDPAddr结构实现了地址接口

```go
type UDPAddr struct {
    IP   IP
    Port int
    Zone string // IPv6 scoped addressing zone
}

func (a *UDPAddr) Network() string { return "udp" }
func (a *UDPAddr) String() string

func ResolveUDPAddr(network, address string) (*UDPAddr, error)
```

## UDPConn结构

IPConn, TCPConn, UDPConn, UnixConn 本质上都一样, 结构内部都是一模一样的一个conn结构.只是实现的差别

除了conn结构实现的,UDPConn结构自己实现了包连接接口,系统调用.连接接口,还有另外4个读写函数

UDPConn实现的4个读写函数: 读从UDP, 写到UDP, 读消息UDP, 写消息UDP. 和IPConn结构如出一辙

```go
type UDPConn struct {
    conn
}

func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error) 
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)

func (c *UDPConn) SyscallConn() (syscall.RawConn, error) 

func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error) 
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)

func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error) 
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error) 

func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)

```

UDP协议里的核心的函数是这几个：

```go
func ResolveUDPAddr(network, address string) (*UDPAddr, error)
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
```


[Daytime Server UDP][4]

[DialUDP][5]

# 引出Conn, PacketConn, Listener接口的小总结

对比以瞅，Dial函数就是简化原来各个协议里的:

* 先通过`network,address`去`ResolveTCPAddr`,再通过`DialTCP`返回`*TCPConn`
* 先通过`network,address`去`ResolveUDPAddr`,再通过`DialUDP`返回`*UDPConn`

而直接通过`network, address`去`Dial`, 返回`Conn`接口

`Listen`之于`ListenTCP`, `ListenPacket`之于`ListenUDP`, 统都一回事，`*TCPListener`结构实现`Listener`接口，`*UDPConn`结构实现`PacketConn`接口

一种是抽象的，一种是具体的。

**这种学习方式是由具体到抽象的步骤，我认为就是应该这样学，而不应该倒过来。**

```go
func ResolveTCPAddr(network, address string) (*TCPAddr, error)
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)

func ResolveUDPAddr(network, address string) (*UDPAddr, error)
func DialUDP(network string, laddr, raddr *UDPAddr) (*UDPConn, error)

func Dial(network, address string) (Conn, error)

func Listen(network, address string) (Listener, error)
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)

func ListenPacket(network, address string) (PacketConn, error)
func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
```

## 重写前头的例子

[TCP Client with Dial][6]

[Daytime Server Support both TCP and UDP][7]

[Echo Server Support both TCP and UDP][8]

[Ping][9]

[9]: https://github.com/iofxl/practisego/blob/master/ping.go

# IP协议里3个结构：IPNet, IPAddr, IPConn

IPNet, IPAddr 都实现了Addr 地址接口

## IPNet结构

```go
type IPNet struct {
    IP   IP     // network number
    Mask IPMask // network mask
}

func (n *IPNet) Network() string { return "ip+net" }
func (n *IPNet) String() string 

func (n *IPNet) Contains(ip IP) bool 
```

## IPAddr结构

```go
type IPAddr struct {
    IP   IP
    Zone string // IPv6 scoped addressing zone
}

func (a *IPAddr) Network() string { return "ip" }
func (a *IPAddr) String() string 

func ResolveIPAddr(network, address string) (*IPAddr, error) 
```

## IPConn结构

IPConn 类型实现了3个接口: Conn, PacketConn, syscall.Conn

```go
// IPConn is the implementation of the Conn and PacketConn interfaces
// for IP network connections.
type IPConn struct {
    conn
}

type conn struct {
    fd *netFD
}


// ReadFromIP acts like ReadFrom but returns an IPAddr.
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error) 
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)

func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error) 

func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error) 
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error) 

func (c *IPConn) SyscallConn() (syscall.RawConn, error) 

//net.go内的conn类型实现了Conn接口, 还实现了另外2个缓存设置函数及1个File函数
func (c *conn) SetReadBuffer(bytes int) error
func (c *conn) SetWriteBuffer(bytes int) error
// On Unix systems this will cause the SetDeadline methods to stop working.
func (c *conn) File() (f *os.File, err error)

func DialIP(network string, laddr, raddr *IPAddr) (*IPConn, error)
func ListenIP(network string, laddr *IPAddr) (*IPConn, error) 
```

# Unix协议里3个结构：UnixAddr, UnixConn, UnixListener

## UnixAddr结构

UnixAddr实现了Addr接口

```go
type UnixAddr struct {
    Name string
    Net  string
}

func (a *UnixAddr) Network() string 
func (a *UnixAddr) String() string 

func ResolveUnixAddr(network, address string) (*UnixAddr, error) 
```

## UnixConn结构

IPConn, TCPConn, UDPConn, UnixConn 本质上都一样, 结构内部都是一模一样的一个conn类型.只是实现的差别

UnixConn自己实现了包连接接口, 系统调用.连接接口,还有6个读写关闭相关函数

UnixConn类型实现的6个函数: 关闭读, 关闭写, 读从Unix, 写到Unix, 读消息Unix, 写消息Unix

两个返回UnixConn类型的函数: 拔号Unix, 监听Unixgram(一定要是Unixgram)

```go
type UnixConn struct {
    conn
}


func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error) 
func (c *UnixConn) WriteTo(b []byte, addr Addr) (int, error) 

func (c *UnixConn) SyscallConn() (syscall.RawConn, error) 

func (c *UnixConn) CloseRead() error
func (c *UnixConn) CloseWrite() error

func (c *UnixConn) ReadFromUnix(b []byte) (int, *UnixAddr, error)
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (int, error)

func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)

func DialUnix(network string, laddr, raddr *UnixAddr) (*UnixConn, error)
func ListenUnixgram(network string, laddr *UnixAddr) (*UnixConn, error) 
```

## UnixListener结构

```go
type UnixListener struct {
    fd         *netFD
    path       string
    unlink     bool
    unlinkOnce sync.Once
}


func (l *UnixListener) Accept() (Conn, error) 
func (l *UnixListener) Close() error
func (l *UnixListener) Addr() Addr { return l.fd.laddr }

func (l *UnixListener) SyscallConn() (syscall.RawConn, error) 

func (l *UnixListener) AcceptUnix() (*UnixConn, error)
func (l *UnixListener) File() (f *os.File, err error) 
func (l *UnixListener) SetDeadline(t time.Time) error
func (l *UnixListener) SetUnlinkOnClose(unlink bool)

func ListenUnix(network string, laddr *UnixAddr) (*UnixListener, error)
```

# Dialer结构

A Dialer contains options for connecting to an address.

The zero value for each field is equivalent to dialing without that option. Dialing with the zero value of Dialer is therefore equivalent to just calling the Dial function. 

```go
type Dialer struct {
    // The default is no timeout.
    // 会影响超时的东西, 操作系统设置的超时
    // Timeout, Deadline誰短誰有效吧
    Timeout time.Duration
    Deadline time.Time
    LocalAddr Addr
    // 没説这个默认值是什么
    DualStack bool
    FallbackDelay time.Duration
    // 默认值?
    KeepAlive time.Duration
    Resolver *Resolver
}

func Dial(network, address string) (Conn, error) {
    var d Dialer
    return d.Dial(network, address)
}
func DialTimeout(network, address string, timeout time.Duration) (Conn, error) {
    d := Dialer{Timeout: timeout}
    return d.Dial(network, address)
}

func (d *Dialer) Dial(network, address string) (Conn, error) {
    return d.DialContext(context.Background(), network, address)
}
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)
```

# ListenConfig结构

ListenConfig contains options for listening to an address. 

```go
type ListenConfig struct {
// If Control is not nil, it is called after creating the network
// connection but before binding it to the operating system.
//
// Network and address parameters passed to Control method are not
// necessarily the ones passed to Listen. For example, passing "tcp" to
// Listen will cause the Control function to be called with "tcp4" or "tcp6".
Control func(network, address string, c syscall.RawConn) error
}

func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)
```


# DNS的3个记录结构：SRV, MX, NS

```go
// An SRV represents a single DNS SRV record.
type SRV struct {
    Target   string
    Port     uint16
    Priority uint16
    Weight   uint16
}

type MX struct {
    Host string
    Pref uint16
}

type NS struct {
    Host string
}
```

# Resolver结构

是对 Dial() 的封装

# DNS的9个查找函数

LookupAddr(r.LookupAddr)是反向查询, LookupIP(r.LookupIPAddr)不是

**LookupIP(r.LookupIPAddr)** 和 LookupHost(r.LookupHost)是相同功能的.只不过返回类型不一样.前者返回相应的类型,后者返回字符串的切片.

```go
var DefaultResolver = &Resolver{}

type Resolver struct {
    PreferGo bool
    StrictErrors bool
    Dial func(ctx context.Context, network, address string) (Conn, error)
}

func LookupHost(host string) (addrs []string, err error) {
    return DefaultResolver.LookupHost(context.Background(), host)
}
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)

func LookupIP(host string) ([]IP, error) {
    addrs, err := DefaultResolver.LookupIPAddr(context.Background(), host)
    ...
}
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error) 

func LookupPort(network, service string) (port int, err error)
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error) 

func LookupCNAME(host string) (cname string, err error) 
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (cname string, err error)

func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error) 
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (cname string, addrs []*SRV, err error) 

func LookupMX(name string) ([]*MX, error) 
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error) 

func LookupNS(name string) ([]*NS, error)
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error) 

func LookupTXT(name string) ([]string, error) 
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)

func LookupAddr(addr string) (names []string, err error)
func (r *Resolver) LookupAddr(ctx context.Context, addr string) (names []string, err error)
```

# 6种网络错误结构

```go
type ParseError struct {
    Type string
    Text string
}

type OpError struct {
    Op string
    Net string
    Source Addr
    Addr Addr
    Err error
}

type AddrError struct {
    Err  string
    Addr string
}

type UnknownNetworkError string

type InvalidAddrError string

type DNSError struct {
    Err         string // description of the error
    Name        string // name looked for
    Server      string // server used
    IsTimeout   bool   // if true, timed out; not all timeouts set this
    IsTemporary bool   // if true, error is temporary; not all errors set this
}
```

# Buffers结构

```go
type Buffers [][]byte
func (v *Buffers) Read(p []byte) (n int, err error)
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error)
```

# 2个拆合主机端口的函数

```go
func SplitHostPort(hostport string) (host, port string, err error)
func JoinHostPort(host, port string) string 
```

[1]: https://github.com/iofxl/gogogo/blob/master/tcpc.go
[2]: https://github.com/iofxl/practisego/blob/master/daytimes.go
[3]: https://github.com/iofxl/practisego/blob/master/echos.go
[4]: https://github.com/iofxl/practisego/blob/master/daytimesudp.go
[5]: https://github.com/iofxl/practisego/blob/master/udpc.go
[6]: https://github.com/iofxl/practisego/blob/master/tcpcInterface.go
[7]: https://github.com/iofxl/practisego/blob/master/daytimes.with.listen.go
[8]: https://github.com/iofxl/practisego/blob/master/echos.use.Listen.go
