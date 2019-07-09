# shadowsocks

# Links

* [Shadowsocks - Protocol](https://shadowsocks.org/en/spec/Protocol.html)
* [Shadowsocks - AEAD Ciphers](https://shadowsocks.org/en/spec/AEAD-Ciphers.html)
* [GitHub - shadowsocks/go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)
* [GitHub - UESTC-BBS/socket-tcp-proxy](https://github.com/UESTC-BBS/socket-tcp-proxy)

# Protocol Format

```
[1-byte type][variable-length host][2-byte port]
```

协议的全部内容就是socks5中的地址的部分

## 和socks5的区别

socks5的做法是，解析得到addr后，去连接addr，回复成功或失败，再开始转发

ss的是做法是，C端将表达addr的字节串直接发得S端，S端解析得到addr后，就跟原socks5一样罢

# Implementation

[6ss](https://github.com/iofxl/practisego/tree/master/06ss)

# 接口的设计

[go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)的接口设计很漂亮，很GO。

最重要的一个接口就是一个以net.Conn接口为参数，并返回net.Conn接口的接口，经过这个接口处理之后，新的net.Conn接口的流为加密流

```go
type Guochaner interface {                                                                                                                                                             
        Guochan( conn net.Conn ) net.Conn                                                                                                                                              
} 
```

要注意这种接口跟读写接口那种接口是有区别的：这种接口要求返回一个东西，读写接口那种要求实现一种功能。

```go
func ( m Method ) Guochan( c net.Conn ) net.Conn {                                                                                                                                     
        return NewConn(m,c)                                                                                                                                                            
} 
```

接下来这个工作也就是一种设计，我感觉这是非常漂亮的设计，不需要等接口真正实现出来就有法做这些事体：

```go
type listener struct {                                                                                                                                                                 
        net.Listener                                                                                                                                                                   
        Guochaner                                                                                                                                                                      
}                                                                                                                                                                                      
                                                                                                                                                                                       
func (l *listener ) Accept() (net.Conn, error) {                                                                                                                                       
        c, err := l.Listener.Accept()                                                                                                                                                  
        return l.Guochan(c), err                                                                                                                                                       
}                                                                                                                                                                                      
                                                                                                                                                                                       
func Listen( g Guochaner, network, address string ) (net.Listener, error ) {                                                                                                           
        l, err := net.Listen(network, address)                                                                                                                                         
        return &listener{l, g}, err                                                                                                                                                    
}                                                                                                                                                                                      
                                                                                                                                                                                       
func Dial( g Guochaner, network, address string ) (net.Conn, error ) {                                                                                                                 
        c, err := net.Dial(network, address)                                                                                                                                           
        return g.Guochan(c), err                                                                                                                                                       
} 
```

这种接口设计的好处，就是不管么要使加密，使哪种加密方法，只要实现这个接口，代码统都不需要修改的。

到这时，所有的工作就是要去实现这个接口。

# 加密

要解决的一个问题，AEAD接口跟BlockMode接口都不是Stream接口这样的，有法直接读写流。要自己实现一下。

使Stream接口去实现简单

# 加密方法的管理

可以学习`crypto.Hash`结构管理`func() hash.Hash`的方法:

每一种方法使常量编号， 使切片存相关固定长度的东西等相关参数，还有生成接口的函数，通过New函数找得要使的函数并直接执行相应的函数，通过切片的key，将相关的东西跟编号联系起。

各个编号都实现这几个方法： New, Size, Available, HashFunc

这种东西可以喊作方法选择器

搞么很多时才没搞明白这么简单的道理： **函数返回接口，返回的其实是实现了这个接口的东西**

在这里，我需要选择的一个加密方法，就要对应的需要的函数并执行这个函数
