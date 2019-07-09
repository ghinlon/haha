# Network Optimization

# BBR

* show  
   `sysctl net.ipv4.tcp_congestion_control`
* check if available  
    `sysctl net.ipv4.tcp_available_congestion_control`  
    if output has no bbr, run `modprobe tcp_bbr`
* TCP algorithm
    高延迟使用hybla  
    低延迟使用htcp,或者bbr.   

# sysctl.conf

Config File `/etc/sysctl.conf`:

After edited, run `sysctl -p` to apply.

```
# max open files
fs.file-max = 1024000
# max read buffer
net.core.rmem_max = 67108864
# max write buffer
net.core.wmem_max = 67108864
# default read buffer
net.core.rmem_default = 65536
# default write buffer
net.core.wmem_default = 65536
# max processor input queue
net.core.netdev_max_backlog = 4096
# max backlog
net.core.somaxconn = 4096

# resist SYN flood attacks
net.ipv4.tcp_syncookies = 1
# reuse timewait sockets when safe
net.ipv4.tcp_tw_reuse = 1
# turn off fast timewait sockets recycling
net.ipv4.tcp_tw_recycle = 0
# short FIN timeout
net.ipv4.tcp_fin_timeout = 30
# short keepalive time
net.ipv4.tcp_keepalive_time = 1200
# outbound port range
net.ipv4.ip_local_port_range = 10000 65000
# max SYN backlog
net.ipv4.tcp_max_syn_backlog = 4096
# max timewait sockets held by system simultaneously
net.ipv4.tcp_max_tw_buckets = 5000
# TCP receive buffer
net.ipv4.tcp_rmem = 4096 87380 67108864
# TCP write buffer
net.ipv4.tcp_wmem = 4096 65536 67108864
# turn on path MTU discovery
net.ipv4.tcp_mtu_probing = 1

# for high-latency network
net.ipv4.tcp_congestion_control = bbr
# forward ipv4
net.ipv4.ip_forward = 1

# turn on TCP Fast Open on both client and server side
net.ipv4.tcp_fastopen = 3
```

# limits.conf

Config File `/etc/security/limits.conf`:

```
*               soft    nofile          512000
*               hard    nofile          1024000
*               hard    noproc          64000
*               soft    noproc          64000
```
relogin into shell to apply, doesn't need to reboot server.

To show:

`ulimit -n`


