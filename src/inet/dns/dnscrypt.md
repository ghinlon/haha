# DNSCrypt

# Links

* [DNS over HTTPS - Cloudflare Resolver](https://developers.cloudflare.com/1.1.1.1/dns-over-https/)
* [Installation · jedisct1/dnscrypt-proxy Wiki](https://github.com/jedisct1/dnscrypt-proxy/wiki/installation)


# Install

Config

```
cp example-dnscrypt-proxy.toml dnscrypt-proxy.toml

// modify the content
force_tcp = true
proxy = "socks5://127.0.0.1:1080"
```

Install

```
./dnscrypt-proxy -service install
./dnscrypt-proxy -service start
```

# Basis

```
./dnscrypt-proxy -service stop|restart

./dnscrypt-proxy -resolve example.com
./dnscrypt-proxy -service uninstall
```

## Install on linux

* [Release 2.0.25 · jedisct1/dnscrypt-proxy](https://github.com/jedisct1/dnscrypt-proxy/releases/tag/2.0.25)

It's same as install on macOS.

