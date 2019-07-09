# Termux

# Links

* [Termux-setup-storage - Termux Wiki](https://wiki.termux.com/wiki/Termux-setup-storage)
* [Remote Access - Termux Wiki](https://wiki.termux.com/wiki/Remote_Access#SSH)


It is necessary to grant storage permission for Termux on Android 6 and higher. Use 'Settings>Apps>Termux>Permissions>Storage' and set to true.

Execute `termux-setup-storage` (run apt update && apt upgrade to make sure that this tool is available) 

# ssh server

Default Port: `8022`


# Accessing Termux from the Internet : Tor

```
pkg install tor torsocks

// modify $PREFIX/etc/tor/torrc

HiddenServiceDir /data/data/com.termux/files/home/tor/hidden_service
HiddenServicePort 22 127.0.0.1:8022

cat ~/.tor/hidden_ssh/hostname 
```

service test:

`torsocks ssh <hostname>`

# DNS error

This can be resolved with `termux-chroot`

