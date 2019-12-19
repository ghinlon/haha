# Proxy Setting

If the program doesn't support socks5，There's two way:

1. Use [provixy](https://www.privoxy.org/) or [polipo](https://www.irif.fr/~jch/software/polipo/) change socks5 to http
    Now I more like to use [zenhack/socks2http](https://github.com/zenhack/socks2http), It's simple and worked.  
	Sun Dec 08 18:24:15 CST 2019  
2. Use [proxychains-ng](https://github.com/rofl0r/proxychains-ng) forward to socks5

# Links

* [proxy - How convert Socks to HTTP in Mac OSX? - Stack Overflow](https://stackoverflow.com/questions/27658665/how-convert-socks-to-http-in-mac-osx)

# socks2http

* [zenhack/socks2http](https://github.com/zenhack/socks2http)

# ssh socks5 proxy

`ssh -gfNTD 1080 <sshd>`

# proxychains-ng

[proxychains-ng](https://github.com/rofl0r/proxychains-ng)

```
sudo apt update
sudo apt install build-essential

./configure --prefix=/usr --sysconfdir=/etc
make
sudo make install

// sample config file:
cp src/proxychains.conf /etc/
```
cfgFile: `/etc/proxychains.conf`

```
[ProxyList]
socsk5 localhost 1080
```

Use:

```
proxychains4 <cmd>
proxychains4 apt upgrade
```

# Privoxy

[Website](http://www.privoxy.org/)

```
// Ubuntu 18.04
apt install privoxy

// Mac
proxychains4 brew install privoxy
```

cfgFile:

* `/usr/local/etc/privoxy/config`		// Mac
* `/etc/privoxy/config`

```
// There's a "dot" at the end.
forward-socks5 / 127.0.0.1:1080 .

// default port: 8118
systemctl enable privoxy
systemctl start privoxy 

// Mac
privoxy /usr/local/etc/privoxy/config
```
## Config Autostart

cfgFile: `Library/LaunchAgents/local.privoxy.plist`

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
        <key>Label</key>
        <string>privoxy</string>
        <key>ProgramArguments</key>
        <array>
          <string>/usr/local/sbin/privoxy</string>
          <string>--no-daemon</string>
          <string>/usr/local/etc/privoxy/config</string>
        </array>
        <key>RunAtLoad</key>
        <true/>
        <key>StandardErrorPath</key>
        <string>~/privoxy.log</string>
        <key>StandardOutPath</key>
        <string>~/privoxy.log</string>
</dict>
</plist>
```

Execute:

`launchctl load local.privoxy.plist`


# polipo

[polipo](https://www.irif.fr/~jch/software/polipo/)

This stupid program can't use single quote in variable assignment, such as:

`proxyAddress = '0.0.0.0'`

**This won't work. MUST use double quote.**

```
// Mac
brew install polipo
```

cfgFile: `.poliporc`

```
socksParentProxy = 127.0.0.1:1080
socksProxyType = socks5

// default port is 8123
proxyAddress = "0.0.0.0"
// proxyPort = 8123
```


Execute: `polipo`

# apt

* Support http_proxy
* Doesn't support socks5

cfgFile: `/etc/apt/apt.conf.d/01proxy`

```
Acquire::http::Proxy "http://localhost:8118";
Acquire::https::Proxy "http://localhost:8118";
```

# docker

using the `HTTP_PROXY`, `HTTPS_PROXY`, and `NO_PROXY` environment variables. 

# git

* use ssh protocol doesn't support http_proxy
* http_proxy here can config with a socks5 proxy

```
// set
git config --global http.proxy $http_proxy
git config --global https.proxy $https_proxy

// unset
git config --global --unset http.proxy
git config --global --unset https.proxy

// check
cat  ~/.gitconfig
```

# go get

* Doesn't support forwarding with proxychains
* support http_proxy

# chrome

`./chrome.exe --proxy-server="socks5://localhost:1080"`

# npm

cfgFile: `.npmrc`

```
proxy=http://localhost:8118/
https-proxy=http://localhost:8118
```

# android

* [wi fi - How to set WIFI proxy via adb shell? - Android Enthusiasts Stack Exchange](https://android.stackexchange.com/questions/98287/how-to-set-wifi-proxy-via-adb-shell)
* [Undo setting proxy via Settings.Global in Android - Stack Overflow](https://stackoverflow.com/questions/31807559/undo-setting-proxy-via-settings-global-in-android)


```
// add
adb shell settings put global http_proxy <address>:<port>
// delete
adb shell settings delete global global_http_proxy_port
```


