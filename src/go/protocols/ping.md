# ICMP:ping

# Links

* [Raw sockets and the type IPConn | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/raw_sockets_and_the_type_ipconn.html)
* [ping (networking utility) - Wikipedia](https://en.wikipedia.org/wiki/Ping_(networking_utility))
* [Internet Control Message Protocol - Wikipedia](https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol)
* [Mike Muuss - Wikipedia](https://en.wikipedia.org/wiki/Mike_Muuss)
* [Checksum算法 - 简单的过客 - CSDN博客](https://blog.csdn.net/zjli321/article/details/74908451)

# 协议格式

# Header

The ICMP header starts after the [IPv4 header](https://en.wikipedia.org/wiki/IPv4#Header) and is identified by IP protocol number. All ICMP packets have an 8-byte header and variable-sized data section. The first 4 bytes of the header have fixed format, while the last 4 bytes depend on the type/code of that ICMP packet.

# ICMP Header Format 

```go
msg[0] = Type
msg[1] = Code
msg[2:4] = Checksum
msg[4:8] = Rest of Header
```

# ping的包头

ping 使Type = 8 and Type = 0 这两个。分别就是 Echo Request and Echo Reply, 两种都Code = 0

```
The format of the ICMP packet payload is as follows:

• The first byte is 8, standing for the echo message.
• The second byte is zero.
• The third and fourth bytes are a checksum on the entire message.
• The fifth and sixth bytes are an arbitrary identifier.
• The seventh and eight bytes are an arbitrary sequence number.
• The rest of the packet is user data.
```
# Checksum算法

里面的checksum怎么算,是只算包头部分还是要跟payload一时算?

是要跟payload一时算的。

依据就是: `^(x + ^x) == 0`

数据按位取反，再跟原来的数相加，再取反，就是0；^x就是校验码，接收方收到数据，做上面那个公式就可以罢

具体算法就是把数据切分成每16bit一组，全部相加，如果数据是奇数个字节的，最末尾补一字节0

```
checksum的计算方法

1、 先将需要计算checksum数据中的checksum设为0；
2、 计算checksum的数据按2byte划分开来，每2byte组成一个16bit的值，如果最后有单个byte的数据，补一个byte的0组成2byte；
3、 将所有的16bit值累加到一个32bit的值中；
4、 将32bit值的高16bit与低16bit相加到一个新的32bit值中，若新的32bit值大于0Xffff, 再将新值的高16bit与低16bit相加；
5、 将上一步计算所得的16bit值按位取反，即得到checksum值，存入数据的checksum字段即可
```

