# [WireGuard: fast, modern, secure VPN tunnel](https://www.wireguard.com/)

When use vpn, the server knows who you are.

# Links

* [Installation - WireGuard](https://www.wireguard.com/install/)
* [Quick Start - WireGuard](https://www.wireguard.com/quickstart/)
* [Wireguard - Debian Wiki](https://wiki.debian.org/Wireguard)
* [WireGuard - ArchWiki](https://wiki.archlinux.org/index.php/WireGuard)
* [Cheatsheet for setting up a WireGuard client on a Mac](https://medium.com/@headquartershq/setting-up-wireguard-on-a-mac-8a121bfe9d86)
* [WireGuard VPN Road Warrior Setup – emanuelduss.ch](https://emanuelduss.ch/2018/09/wireguard-vpn-road-warrior-setup/)
* [Wireguard VPN: Typical Setup - The poetry of (in)security](https://www.ckn.io/blog/2017/11/14/wireguard-vpn-typical-setup/)
* [Wireguard VPN: Portable Raspberry Pi Setup - The poetry of (in)security](https://www.ckn.io/blog/2017/12/28/wireguard-vpn-portable-raspberry-pi-setup/)

# HowTo

[How to setup a VPN server using WireGuard (with NAT and IPv6)](https://angristan.xyz/how-to-setup-vpn-server-wireguard-nat-ipv6/)

1. Make the client's WireGuard interface its gateway (default route)
1. Enable IP Forwarding on the server
1. Enable NAT between the WireGuard interface and public interface on the
   server

# Install

```
// Ubuntu
$ sudo add-apt-repository ppa:wireguard/wireguard
$ sudo apt-get update
$ sudo apt-get install wireguard

// macOS
brew install wireguard-tools

// Raspberry Pi
// building kernel first. 
git clone https://git.zx2c4.com/WireGuard
cd WireGuard/src
make
sudo make install
sudo modprobe wireguard

//  Red Hat Enterprise Linux / CentOS
sudo curl -Lo /etc/yum.repos.d/wireguard.repo https://copr.fedorainfracloud.org/coprs/jdoss/wireguard/repo/epel-7/jdoss-wireguard-epel-7.repo
sudo yum install epel-release
sudo yum install wireguard-dkms wireguard-tools
```

# IP Forwarding

`net.ipv4.ip_forward = 1` 

meaning allowing the kernel to forward packets from one network interface to
another. 

cmd:

```
cat /proc/sys/net/ipv4/ip_forward
echo 1 > /proc/sys/net/ipv4/ip_forward
// or
sysctl -w net.ipv4.ip_forward=1
```

cfg:

Config File `/etc/sysctl.conf`:

```
net.ipv4.ip_forward=1
```

apply:

`sysctl -p /etc/sysctl.conf`

# NAT

vpn server needs to config nat masquerade.

```
nft add table nat
nft add chain nat prerouting { type nat hook prerouting priority 0 \; }
nft add chain nat postrouting { type nat hook postrouting priority 100 \; }

nft add rule nat postrouting oif ens4 masquerade
```

# Import by reading a QR code

```
apt install qrencode

qrencode -t ansiutf8 < client.conf
```

# Conceptual Overview

Wireguard works like ssh. 

**Cryptokey Routing**

works by associating public keys with a list of tunnel IP addresses that are
allowed inside the tunnel. 

Each network interface has a private key and a list of peers. Each peer has
a public key. Public keys are short and simple, and are used by peers to
authenticate each other. They can be passed around for use in configuration
files by any out-of-band method, similar to how one might send their SSH public
key to a friend for access to a shell server.

In other words, when sending packets, the list of allowed IPs behaves as a sort
of routing table, and when receiving packets, the list of allowed IPs behaves
as a sort of access control list.

This is what we call a Cryptokey Routing Table: the simple association of
public keys and allowed IPs.

# Quick Start

* Generate key and exchange public key with server

	Every wireguard interface needs a pair of key:

	```
	wg genkey | tee -a privatekey | wg pubkey >> publickey
	chmod 600 *key
	```

* Create config file `/etc/wireguard/wg0.conf`: 

	```
	[Interface]
	PrivateKey = <server privatekey>
	// client doesn't need config ListenPort
	ListenPort = <public listen port>
    Address = <client ip>
	// client append below
	DNS = x.x.x.x

	[Peer]
	PublicKey = <client publickey>
	AllowedIPs = <client ip>
	// client append below
	AllowedIPs = 0.0.0.0/0
    Endpoint = <server public ip:port>
    PersistentKeepalive = 25

	```

* wg-quick 

	```
	wg-quick [up|down] wg0
	// or
	wg-quick [up|down] ./wg0.conf
	```

* auto start

 `systemctl enable wg-quick@wg0` 

* show

```
# ip a
# wg 
```

# Manual setup

[WireGuard - Fast and secure kernelspace VPN](https://git.zx2c4.com/WireGuard/about/src/tools/man/wg-quick.8)

```
ip link add dev wg0 type wireguard
ip address add dev wg0 192.168.2.1/24
ip link set dev wg0 up
wg setconf wg0 myconfig.conf
// or
wg setconf wgpi <(wg-quick strip ./wgpi.conf)
```

# Config with systemd

Config File `/etc/systemd/network/wg0.netdev`:

```
[NetDev]
Name=wg0
Kind=wireguard
Description=Wireguard

[WireGuard]
PrivateKey=<paste the private key of the local host here>
DNS = <server ip>

[WireGuardPeer]
PublicKey=<paste the public key of the remote host here>
AllowedIPs=0.0.0.0/0
AllowedIPs=::/0
Endpoint=<remote IP or hostname>:<remote port>
```

change file permissions:

```sh
chown root.systemd-network /etc/systemd/network/wg0.netdev
chmod 0640 /etc/systemd/network/wg0.netdev
```
Config File `/etc/systemd/network/wg0.network`:

```
[Match]
Name=wg0

[Network]
Address=10.88.88.1/24
Address=2001:db8:1234:5678::1
```

Start

```
systemctl daemon-reload
systemctl start systemd-networkd
```

# Rule-based Routing


# Chain Routing

* [Routing & Network Namespaces - WireGuard](https://www.wireguard.com/netns/)
* [Wireguard VPN: Chained Setup - The poetry of (in)security](https://www.ckn.io/blog/2017/12/28/wireguard-vpn-chained-setup/)

# Delete utunN on macOS

`rm -f /var/run/wireguard/utun1.sock`

# can't self ping wg0's ip 

[can't ping remote side IP range from WG instance](https://lists.zx2c4.com/pipermail/wireguard/2018-March/002559.html)

the method is to remove masquerade rule in firewall. but, then how the vpn will
work ?

