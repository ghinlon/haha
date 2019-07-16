# IP FORWARDING and NAT

# Links

* [iptables - ArchWiki](https://wiki.archlinux.org/index.php/Iptables)
* [Simple stateful firewall - ArchWiki](https://wiki.archlinux.org/index.php/Simple_stateful_firewall)
* [tables_traverse.jpg (JPEG Image, 647 × 1100 pixels) - Scaled (78%)](https://www.frozentux.net/iptables-tutorial/images/tables_traverse.jpg)
* [Getting Started with Software-Defined Networking and Creating a VPN with ZeroTier One | DigitalOcean](https://www.digitalocean.com/community/tutorials/getting-started-software-defined-networking-creating-vpn-zerotier-one?utm_medium=social&utm_source=twitter&utm_campaign=zerotier_tut&utm_content=no_image)
* [【转】iptables配置——NAT地址转换 - veryhappy - 博客园](https://www.cnblogs.com/shineshqw/articles/2351028.html)

# system config

**net.ipv4.ip_forward=1**

cmd: 

```
// print
sysctl net.ipv4.ip_forward
net.ipv4.ip_forward = 1

// set
sysctl net.ipv4.ip_forward=1
net.ipv4.ip_forward = 1
```

cfg file: `/etc/sysctl.conf`

```
net.ipv4.ip_forward = 1
```

Save and close the file, then run `sysctl -p` to trigger the kernel's adoption
of the new configuration

**net.ipv4.conf.all.rp_filter=2**

This configuration change is required to alter the kernel's view of what an
acceptable return path for your client traffic is. Due to the way that the
ZeroTier VPN is configured, the traffic coming back from your server to
your client can sometimes appear to come from a different network address
than the one it was sent it to. By default, the Linux kernel views these as
invalid and drops them, making it necessary to override that behavior.


# waht is iptables

![img:tables_traverse.jpg](imgs/tables_traverse.jpg "tables_traverse")

```
PREROUTING(nat/DNAT) 	- FORWARD(filter) 					- POSTROUTING(nat/SNAT)
PREROUTING 				- INPUT - LOCAL PROCESS - OUTPUT	- POSTROUTING 
```

# NAT 原理

同filter表一样，nat表也有三条缺省的"链"(chains)：

* PREROUTING: when DNAT

	把从外来的访问重定向到其他的机子上，比如内部SERVER，或者DMZ。

	因为路由时只检查数据包的目的ip地址，所以必须在路由之前就进行目的PREROUTING DNAT;

	**系统先PREROUTING DNAT翻译——>再过滤（FORWARD）——>最后路由。**

	**路由和过滤（FORWARD)中match 的目的地址，都是针对被PREROUTING DNAT之后的。**

* POSTROUTING: when SNAT

	在路由以后在执行该链中的规则。

	**系统先路由——>再过滤（FORWARD)——>最后才进行POSTROUTING SNAT地址翻译**

	**其match 源地址是翻译前的。**

* OUTPUT:定义对本地产生的数据包的目的NAT规则

-i ，-o 参数在NAT中的用途:

`-i` only used in `PREROUTING`, when `-o` only used in `POSTROUTING and OUTPUT`

# Internal to access Internet with -j SNAT

```
-j SNAT --to-source ipaddr[-ipaddr][:port-port]
网络地址转换，SNAT就是重写包的源IP地址
only valid in POSTROUTING

iptables -t nat -A POSTROUTING -s 192.168.0.0/24 -o eth0 -j SNAT --to 你的eth0地址

-j MASQUERADE
用于外网口public地址是DHCP动态获取的（如ADSL）

iptables -t nat -A POSTROUTING –o eth1 –s 192.168.1.0/24 –j MASQUERADE
```

真正实现还需要FORWARD链中两条rule的配合,并且这两条规则是先执行的：

```
iptables -A FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
iptables -A FORWARD -i zt0 -o eth0 -j ACCEPT
```

## Internet to access Internal with –j DNAT

```
-j DNAT --to-destination ipaddr[-ipaddr][:port-port]
目的网络地址转换，重写包的目的IP地址
only valid in PREROUTING

DNAT Example:

外部接口ip：	210.83.2.206
内部接口ip：	192.168.1.1
ftp Server:		192.168.1.3
web Server:		192.168.1.4

iptables -t nat -A PREROUTING -d 210.83.2.206 -p tcp --dport 21 -j DNAT --to 192.168.1.3
iptables -t nat -A PREROUTING -d 210.83.2.206 -p tcp --dport 80 -j DNAT --to 192.168.1.4
```

需要在FORWARD链中允许DNAT后的目的地址：

```
-A FORWARD -d 192.168.1.3/32 -p tcp --dport 21 -j ACCEPT  
-A FORWARD -d 192.168.1.4/32 -p tcp --dport 80 -j ACCEPT  
```

而且用于转发的机器是**不需要监听**这两个端口的。以前我一直觉得凡能接受连接必有监听...

And no matter SNAT or DNAT, because the packets always will go though the
filter table, so must config the FORWARD chain to accept them.

# DNAT for Port Mapping

```
iptables -t nat -A PREROUTING -p tcp -d 216.94.87.37 --dport 2121 -j DNAT --to-destination 192.168.100.125:21
```

# DNAT as load-balance

```
iptables –t nat –A PREROUTING –d 219.142.217.161 –j DNAT --to-destination 192.168.1.24-192.168.1.25
```

# Used for Jumping

```
iptables -t nat -A PREROUTING -i zt0 -p tcp -d 216.94.87.37 --dport 8443 -j DNAT --to-destination 192.168.100.125:443

iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
// 如果eth0的IP是固定的，也可以用 -j SNAT --to-source ipaddr[-ipaddr][:port-port]

iptables -A FORWARD -d 192.168.100.125 -p tcp --dport 443 -j ACCEPT
```

# iptables config

* enable 

```
# iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
// 这条命令大小写敏感,何必呢
// 这条命令不指定 -o 也是可以的,就所有网卡上都可以NAT了
```
* Permit traffic and track active connections:

```
iptables -A FORWARD -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
这里为什么要conntrack, 为什么要--ctstate呢?
conntrack是一个模块,提供追踪连接状态的功能.
后面--ctstate是它的选项.ctstate就是conntrack state的简写,就是我要追踪后面这两个状态的报文.
```

* 指定从zt0进的报文转发到eth0

```
iptables -A FORWARD -i zt0 -o eth0 -j ACCEPT
-i -o 能不能是同一个接口呢?
测试了命令是可以这么配置成功,能不能有什么作用就不知道了.
一般的场景,NAT就是用于这个网口转发到那个网口.如果同一个网口有不同的网段,应该也是允许的.猜的.
```

# Persistent config

ubuntu

install `iptables-persistent`, then `netfilter-persistent save`

others use `iptables-save`

archlinux 

systemd automatically load `/etc/iptables/iptables.rules`

rhel6 and rhel7

automatically load `/etc/sysconfig/iptables`

# iptables basis

```
// save and restore
iptables-save > /etc/iptables/iptables.rules
iptables-restore <  /etc/iptables/iptables.rules

// show current config
iptables -nvL --line-numbers
```

# Simple stateful firewall

* reset iptables

```
iptables -F
iptables -X
iptables -t nat -F
iptables -t nat -X
iptables -t mangle -F
iptables -t mangle -X
iptables -t raw -F
iptables -t raw -X
iptables -t security -F
iptables -t security -X
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT
```
* begin

```
# 这条命令避免已经建立的ssh被中断
iptables -A INPUT -m conntrack --ctstate RELATED,ESTABLISHED -j ACCEPT
# accept all from lo
iptables -A INPUT -i lo -j ACCEPT
# accept another interface
iptables -A INPUT -i wg1 -j ACCEPT
# drop all INVALID
iptables -A INPUT -m conntrack --ctstate INVALID -j DROP
# if you want pings
# iptables -A INPUT -p icmp --icmp-type 8 -m conntrack --ctstate NEW -j ACCEPT
```
* create two chains

```
iptables -N TCP
iptables -N UDP
iptables -A INPUT -p udp -m conntrack --ctstate NEW -j UDP
iptables -A INPUT -p tcp --syn -m conntrack --ctstate NEW -j TCP
iptables -A INPUT -p udp -j REJECT --reject-with icmp-port-unreachable
iptables -A INPUT -p tcp -j REJECT --reject-with tcp-reset
iptables -A INPUT -j REJECT --reject-with icmp-proto-unreachable
```
* accept basic services

```
iptables -A TCP -p tcp --dport 22 -j ACCEPT
iptables -A TCP -p tcp --dport 26462 -j ACCEPT
iptables -A TCP -p tcp --dport 53 -j ACCEPT
iptables -A UDP -p udp --dport 53 -j ACCEPT
iptables -A TCP -p tcp -m multiport --dports 80,443 -j ACCEPT
iptables -A TCP -p tcp -m multiport --dports 8080,65000:65009 -j ACCEPT
iptables -A UDP -p udp -m multiport --dports 8080,48570:48579 -j ACCEPT
```
* default policy

```
iptables -P INPUT DROP
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT
```

# LOG

* [Linux iptables LOG everything - Jesin's Blog](https://websistent.com/linux-iptables-log-everything/)

```
iptables -I INPUT 1 -j LOG
iptables -I FORWARD 1 -j LOG
iptables -I OUTPUT 1 -j LOG

iptables -t nat -I PREROUTING 1 -j LOG
iptables -t nat -I POSTROUTING 1 -j LOG
iptables -t nat -I OUTPUT 1 -j LOG
```
