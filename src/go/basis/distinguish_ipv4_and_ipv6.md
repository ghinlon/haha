# [go - golang distinguish IPv4 IPv6 - Stack Overflow](https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6)

# Links

[go - golang distinguish IPv4 IPv6 - Stack Overflow](https://stackoverflow.com/questions/22751035/golang-distinguish-ipv4-ipv6/22752227#22752227)

```
 simply check ip.To4() != nil. Since the documentation says "if ip is not an IPv4 address, To4 returns nil" this condition should return true if and only if the address is an IPv4 address.
```
