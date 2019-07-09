# Unbound
  
# Links

* [Wireguard VPN: Typical Setup - The poetry of (in)security](https://www.ckn.io/blog/2017/11/14/wireguard-vpn-typical-setup/)

# Install

```  
pacman -S unbound

apt install unbound unbound-host
```

# Config

```
curl -o /var/lib/unbound/root.hints https://www.internic.net/domain/named.cache
chown unbound:unbound /var/lib/unbound/root.hints

// solve root.key anchor error
sudo -u unbound unbound-anchor -v -a /var/lib/unbound/root.key
```

Config File: [repo]

Check Config: `unbound-checkconf`

run:

```
systemctl stop systemd-resolved
systemctl disable systemd-resolved
systemctl start unbound


// view log
journalctl -f -u unbound
```

# Troubleshooting

[Unbound DNS server behind a VIP - solving reply from unexpected source](https://www.claudiokuenzler.com/blog/695/unbound-behind-a-virtual-ip-vip-reply-from-unexpected-source)


# resolv.conf

There has trouble with set `nameserver` to `127.0.0.1`, so I set it to the ip
of the external interface, then there's no trouble.

some days later, I find this is because the masquerade rule in the firewall.
just remove it, then everything's ok.

