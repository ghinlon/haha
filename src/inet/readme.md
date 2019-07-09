# THIS WORLD IS DANGEROUS

# DNS leak

If only run socks5, the DNS leak will happen. this is why I more like vpn than
socks5 proxy.

# VPN in VPN

you can run VPN on host, run another VPN in virtual machine.


# all go through socks5

* `proxychains` or tools alike can just let one soft go through, I want all.
* `redsocks` is used for such case. but this tool **can't really proxy UDP**.
* so the best way is to use `wireguard`, then run a `udp2tcp`, then work
  together with `redsocks`:  
  ```
  user_data >
    wg_client >
	  udp2tcp 	>
	    			fw > redsocks > ss_local  > 
									ss_server >
	  tcp2udp 	>
	wg_server >
  www
  ```

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
