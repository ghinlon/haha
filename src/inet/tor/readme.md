# [Tor: Onion Service Protocol](https://www.torproject.org/docs/onion-services.html.en)


# Links

* [Tor: Onion Service Protocol](https://www.torproject.org/docs/onion-services.html.en)
* [Tor Project: Onion Service Configuration Instructions](https://www.torproject.org/docs/tor-onion-service.html.en)
* [Tor: Linux Install Instructions](https://www.torproject.org/docs/tor-doc-unix.html.en)
* [Tor Project: Mac OS X Install Instructions](https://www.torproject.org/docs/tor-doc-osx.html.en)
* [关于 Tor 的常见问题解答](https://program-think.blogspot.com/2013/11/tor-faq.html)
* [socat](http://www.dest-unreach.org/socat/)
* [doc/TorFAQ – Tor Bug Tracker & Wiki](https://trac.torproject.org/projects/tor/wiki/doc/TorFAQ#SOCKSAndDNS)
* [Tor Project: manual](https://www.torproject.org/docs/tor-manual.html.en)

Tor allows clients and relays to offer onion services. That is, you can offer a web server, SSH server, etc., without revealing your IP address to its users. In fact, because you don't use any public address, you can run an onion service from behind your firewall. 

# Really Important Thing

* **[Warning](https://www.torproject.org/download/download.html.en#warning)**
* 什么是“陷阱节点”/“蜜罐节点”？

  如果你使用的线路中，“出口节点”正好是蜜罐，那么该蜜罐就会窥探到你的上网行为——前面已经说了，出口节点肯定会知道你的访问了哪个网站。并且，假如你访问的目标网站没有 HTTPS 加密，蜜罐就会知道你浏览的页面内容。
* 如何避免“陷阱节点”/“蜜罐节点”？

  比较简单的做法，就是通过修改 Tor 的配置文件，规避这些不安全国家的节点。

  in the end of `torrc`, add these config, 表示排除这些国家/地区的节点，StrictNodes 表示强制执行）:

  ```
  ExcludeNodes {cn},{hk},{mo},{kp},{ir},{sy},{pk},{cu},{vn}
  StrictNodes  1
  Socks5Proxy 127.0.0.1:1080
  // 北朝鲜  {kp} 伊朗  {ir} 叙利亚  {sy} 巴基斯坦  {pk} 古巴  {cu} 越南  {vn}
  ```

# Install

* [Download Tor](https://www.torproject.org/download/download.html.en#warning)

```
// Mac 
brew install tor

You will find a sample Tor configuration file at /usr/local/etc/tor/torrc.sample. Remove the .sample extension to make it effective.
```

# NeedToUseAProxy

* [Tor Project: FAQ](https://www.torproject.org/docs/faq.html.en#NeedToUseAProxy)

```
// config in `torrc`
Socks5Proxy host[:port] 
// Tor will make all OR connections through the SOCKS 5 proxy at host:port (or host:1080 if port is not specified). 
```

Raspberry:

```
apt install tor
```

Termux

```
pkg install tor torsocks
```

Launch it: `tor`

# TBB

The recommended way to use Tor is to simply download the Tor Browser and you are done. 

[Tor Browser](https://www.torproject.org/projects/torbrowser.html.en)

# Onion Service

```
// config in torrc

HiddenServiceDir $HOME/tor/hidden_service/
HiddenServicePort 80 127.0.0.1:8080
HiddenServicePort 22 127.0.0.1:22
```

When Tor starts, it will automatically create the HiddenServiceDir that you specified (if necessary), and it will create two files there.

`private_key` and `hostname`




