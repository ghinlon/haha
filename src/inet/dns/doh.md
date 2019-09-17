# DNS over HTTPS (DoH)

# Links

* [Running a DNS over HTTPS Client - Cloudflare Resolver](https://developers.cloudflare.com/1.1.1.1/dns-over-https/cloudflared-proxy/)

# Install

* [Downloads - Argo Tunnel](https://developers.cloudflare.com/argo-tunnel/downloads/)

```
brew install cloudflare/cloudflare/cloudflared
```

# Set up cloudflared as a service

```
mkdir -p /usr/local/etc/cloudflared
cat << EOF > /usr/local/etc/cloudflared/config.yml
proxy-dns: true
proxy-dns-upstream:
 - https://1.1.1.1/dns-query
 - https://1.0.0.1/dns-query
EOF
```

install

```
sudo cloudflared service install
```


