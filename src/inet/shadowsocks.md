# shadowsocks

# Links

# Install

# Configuration

# Server

```
{
        "server":"0.0.0.0",
        "server_port":8388,
        "password":"mypassword",
        "timeout":300,
        "method":"aes-128-gcm",
        "fast_open": true,
}
```

Run:

```
ss-server -uv -c config.json
// run in the background
ss-server -uv -c config.json -d [start|stop]
```


# Client

* [Configuration via Config File · shadowsocks/shadowsocks Wiki · GitHub](https://github.com/shadowsocks/shadowsocks/wiki/Configuration-via-Config-File)

Config File `/etc/shadowsocks-libev/config.json`:

```
{
    "server":"my_server_ip",
    "server_port":8388,
    "local_address": "0.0.0.0",
    "local_port":1080,
    "password":"mypassword",
    "timeout":300,
    "method":"aes-128-gcm",
    "fast_open": true,
    "plugin": "obfs-local",
    "plugin_opts": "obfs=tls;obfs-host=bing.com",
}
```

run:

```
ss-local -uv -c config.json
```

# shadowsocks on raspberry pi

do this is because I need to access the content in network B ( which pi in)
from network A ( which vps in).

and the important thing is , vps in network B is expensive.

1. both need to enable IP_FORWARDING
1. ss on pi
1. wireguard between pi and vps
1. dnat and masquerade on vps.  
	```
	table ip nat {
			chain prerouting {
					type nat hook prerouting priority -100; policy accept;
					iifname "ens4" tcp 12345 tproxy counter log dnat to 10.0.0.1:8388
			}

			chain postrouting {
					type nat hook postrouting priority 100; policy accept;
					ip daddr 10.0.0.1 tcp dport 8388 oif "wg0" counter log masquerade
			}
	}
	```  

[Linux 2.4 NAT HOWTO: Destination NAT Onto the Same Network](https://www.netfilter.org/documentation/HOWTO/NAT-HOWTO-10.html)

**Note:** **MUST** config NAT( snat or masquerade) on vps, cause pi has no
route to the clients. this problem wasted me several hours just now. if you run
`tcpdump`, you can find both vps and pi, there's only packages in, but no
packages out.

[Sun Apr 7 20:21:27 CST 2019]

# PAC

[Shadowrocket-ADBlock-Rules](https://github.com/h2y/Shadowrocket-ADBlock-Rules)

# Troubleshooting

* [[已解决]执行make的时候失败 · Issue #1177 · shadowsocks/shadowsocks-libev · GitHub](https://github.com/shadowsocks/shadowsocks-libev/issues/1177)

Compile on Pi need this:

`./configure --with-sodium-include=/usr/local/include --with-sodium-lib=/usr/local/lib --with-mbedtls-include=/usr/local/include --with-mbedtls-lib=/usr/local/lib`




