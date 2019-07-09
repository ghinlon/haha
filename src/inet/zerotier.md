# ZeroTier

# Links

* [ZeroTier | Home](https://www.zerotier.com/)
* [ZeroTier | Manual](https://www.zerotier.com/manual.shtml)
* [Getting Started with Software-Defined Networking and Creating a VPN with ZeroTier One | DigitalOcean](https://www.digitalocean.com/community/tutorials/getting-started-software-defined-networking-creating-vpn-zerotier-one?utm_medium=social&utm_source=twitter&utm_campaign=zerotier_tut&utm_content=no_image#prerequisites)

# Uninstall

[ZeroTierOne/uninstall.sh at master · zerotier/ZeroTierOne · GitHub](https://github.com/zerotier/ZeroTierOne/blob/master/ext/installfiles/mac/uninstall.sh)

```
apt remove zerotier-one
```

# Quick Start

```
curl -s 'https://pgp.mit.edu/pks/lookup?op=get&search=0x1657198823E52A61' | gpg --import && \
if z=$(curl -s 'https://install.zerotier.com/' | gpg); then echo "$z" | sudo bash; fi

sudo systemctl enable zerotier-one

sudo zerotier-cli status

sudo zerotier-cli join [Network ID]

Authenticate your device by going to https://my.zerotier.com/network/[Network ID]
```

# Managed Routes   

Go to the top-right of your ZeroTier Networks page and add a new route with the following parameters. You can find the ZeroTier IP for your server in the Members section of your ZeroTier Network configuration page. In the network/bits field, enter in 0.0.0.0/0, in the (LAN) field, enter your ZeroTier server's IP address.

* Enable allowDefault and net.ipv4.conf.all.rp_filter=2 on client

    ```
    sudo zerotier-cli set NetworkID allowDefault=1

    Each time the ZeroTier service on the client is restarted, the allowDefault=1 value gets reset to 0, so remember to re-execute it in order to activate the VPN functionality.

# Firewall

allow traffic to/from port 9993.
