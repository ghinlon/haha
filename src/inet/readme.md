# THIS WORLD IS DANGEROUS

# Reserved IP addresses

* [Reserved IP addresses - Wikipedia](https://en.wikipedia.org/wiki/Reserved_IP_addresses)

```
0.0.0.0/8 
10.0.0.0/8 
100.64.0.0/10 
127.0.0.0/8 
169.254.0.0/16 
172.16.0.0/12 
192.0.0.0/24 
192.0.2.0/24 
192.88.99.0/24 
192.168.0.0/16 
198.18.0.0/15 
198.51.100.0/24 
203.0.113.0/24 
224.0.0.0/4 
240.0.0.0/4 
```

# All TLS Ports

* [List of TCP and UDP port numbers - Wikipedia](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)

```
443
465 		Authenticated SMTP[10] over TLS/SSL (SMTPS)[83]
563 		NNTP over TLS/SSL (NNTPS) 
585 		Legacy use of Internet Message Access Protocol over TLS/SSL (IMAPS), now in use at port 993.
636 		Lightweight Directory Access Protocol over TLS/SSL (LDAPS)[10]
853 		DNS over TLS (RFC 7858) 
989 		FTPS Protocol (data), FTP over TLS/SSL
990 		FTPS Protocol (control), FTP over TLS/SSL
992 		Telnet protocol over TLS/SSL
993 		Internet Message Access Protocol over TLS/SSL (IMAPS)[10]
994 		Internet Relay Chat over TLS/SSL (IRCS). Previously assigned, but not used in common practice.[76]
995 		Post Office Protocol 3 over TLS/SSL (POP3S)[10]
4116 		Smartcard-TLS
4843 		OPC UA TCP Protocol over TLS/SSL for OPC Unified Architecture from OPC Foundation
5061 		Session Initiation Protocol (SIP) over TLS
5349 		TURN over TLS/DTLS, a protocol for NAT traversal[169]
5671 		Advanced Message Queuing Protocol (AMQP)[216] over TLS
6436 		Leap Motion Websocket Server TLS
6513 		NETCONF over TLS
6514 		Syslog over TLS[228]
6619 		odette-ftps, Odette File Transfer Protocol (OFTP) over TLS/SSL
8883 		Secure MQTT (MQTT over TLS)[277][278]
10514 		TLS-enabled Rsyslog (default by convention)
```

# DNS leak

If only run socks5, the DNS leak will happen. this is why I more like vpn than
socks5 proxy.

# all go through socks5

* `proxychains` or tools alike just let one soft go through.
* `redsocks` is used for such case. And `ss-redir` also do such things too.

**Caution: this method doesn't really proxy UDP.**

# UDP over TCP

It seams this trick doesn't fit this case.

[How to forward UDP packets through an SSH tunnel - Quora](https://www.quora.com/How-do-you-forward-UDP-packets-through-an-SSH-tunnel)

```
#!/bin/bash
# $0 UDPPort TCPPort <Client|Server>

if [[ $3 == "client" ]]; then
mkfifo /tmp/udp2tcp
netcat -l -u -p $1 < /tmp/udp2tcp | netcat localhost $2 > /tmp/udp2tcp
fi
if [[ $3 == "server" ]]; then
mkfifo /tmp/tcp2udp
netcat -l -p $2 < /tmp/tcp2udp | netcat -u localhost $1 > /tmp/tcp2udp
fi
```
