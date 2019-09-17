# systemd-resolved

# Links

* [systemd-resolved - ArchWiki](https://wiki.archlinux.org/index.php/Systemd-resolved)


# Mannually Config

```
# cat /etc/systemd/resolved.conf 

[Resolve]
DNS=8.8.8.8
#FallbackDNS=
#Domains=
#LLMNR=no
#MulticastDNS=no
#DNSSEC=no
#DNSOverTLS=no
#Cache=yes
#DNSStubListener=yes
#ReadEtcHosts=yes
```

# Check


```
systemctl status systemd-resolved

resolvectl status
```

# DNSSEC

```
[Resolve]
DNSSEC=true
```
