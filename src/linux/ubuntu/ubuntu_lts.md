# Ubuntu的LTS

LTS: Long term support

# Links

* [Ubuntu release cycle | Ubuntu](https://www.ubuntu.com/about/release-cycle)
* [Releases - Ubuntu Wiki](https://wiki.ubuntu.com/Releases)
* [Mirrors : Ubuntu](https://launchpad.net/ubuntu/+archivemirrors)

一般是以.04结尾的版本是LTS

# 版本升级

`proxychains do-release-upgrade`


# 更新源

cfg:

`/etc/apt/sources.list`

```
# 放在最前头
deb http://linux.yz.yamagata-u.ac.jp/ubuntu/ bionic main restricted universe
deb-src http://linux.yz.yamagata-u.ac.jp/ubuntu/ bionic main restricted universe
```

# 更新

通过proxychains走代理

```
proxychains apt update
proxychains apt upgrade
```

