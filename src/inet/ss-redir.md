# ss-redir

# Links

* [使用 nftables 与 ss-redir 实现 shadowsocks transparent proxy](https://huntxu.github.io/2015-01-20-shadowsocks-transparent-proxy-with-nft.html)
* [ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest](https://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest)

# postrouting 


why the masquerade rule in nat chain hook at postrouting lead to no connection when redirect:

* when packages in, they are being redirected to `:12345`, the `saddr` left no
  change.
* the next step packages will go through `input`, and then processed by `ss-redir`,
  after this, the packages daadr will be encapsulated with the server's ip. it
  will go through `output` directly due to the rule `ip daddr <server> return`.
* I have found why:  
  cause the package back from server which need to send to client or send to
  another middle man, these packages saddr is the target server's. so the rule
  should be limited to your local saddr like:  
  `ip saddr 192.168.0.0/24 oif eth0 masquerade`
* BUT I'M WRONG, the rule above totaly doesn't work.


```
table ip nat {
        chain prerouting {
                type nat hook prerouting priority -1; policy accept;
                ip daddr { 0.0.0.0/8, 10.0.0.0/8, 127.0.0.0/8, 169.254.0.0/16, 172.16.0.0/12, 192.168.0.0/16, 224.0.0.0/4, 240.0.0.0/4, 172.16.39.0/24} return
                ip daddr <server> return
                # By using redirect, packets will be forwarded to local machine. 
                # Is a special case of DNAT where the destination is the current machine. 
                # tcp sport {32768-61000} redirect to 12345     # /proc/sys/net/ipv4/ip_local_port_range
                ip protocol tcp redirect to 12345  
        }

        chain input {
                type nat hook input priority 0; policy accept;
        }


        # this is for local packages generated from "lo"
        chain output {
                type nat hook output priority 0; policy accept;
                ip daddr { 0.0.0.0/8, 10.0.0.0/8, 127.0.0.0/8, 169.254.0.0/16, 172.16.0.0/12, 192.168.0.0/16, 224.0.0.0/4, 240.0.0.0/4, 172.16.39.0/24} return 
                ip daddr <server> return 
                ip protocol tcp redirect to 12345  
        }

        chain postrouting {
                type nat hook postrouting priority 0; policy accept;
        }

}
```









