# Install

# Links

* [Raspberry Pi Documentation](https://www.raspberrypi.org/documentation/)
* [Remote Access - Raspberry Pi Documentation](https://www.raspberrypi.org/documentation/remote-access/)
* [Setting WiFi up via the command line - Raspberry Pi Documentation](https://www.raspberrypi.org/documentation/configuration/wireless/wireless-cli.md)


```
// Mac
diskutil list

    /dev/disk2 (external, physical):

diskutil unmountDisk /dev/disk2s1

sudo dd bs=1m if=image.img of=/dev/rdisk<disk# from diskutil> conv=sync
or
sudo dd bs=1m if=image.img of=/dev/disk2 conv=sync // I choose this 

sudo diskutil eject /dev/rdisk<disk# from diskutil>
```

Enable SSH on a headless Raspberry Pi (add file to SD card on another machine)

For headless setup, SSH can be enabled by placing a file named ssh, without any extension, onto the boot partition of the SD card from another computer. When the Pi boots, it looks for the ssh file. If it is found, SSH is enabled and the file is deleted. The content of the file does not matter; it could contain text, or nothing at all.

`cd /Volumes/boot && echo ssh > ssh`

```  
unzip -p 2017-11-29-raspbian-stretch.zip | sudo dd of=/dev/sdX bs=4M conv=fsync
sync
/*  
       -p     extract  files  to pipe (stdout).  Nothing but the file data is sent to stdout, and the files are always extracted in binary format, just as they are stored (no
              conversions).
*/
```

# Ip

```
brew install nmap

nmap -sn 192.168.1.0/24	    // ping scan
```

# Default user

`pi/raspberry`

Add user

```
useradd -ms /bin/bash <user>
passwd <user>
```

# Access from Internet

```
curl -s 'https://pgp.mit.edu/pks/lookup?op=get&search=0x1657198823E52A61' | gpg --import && \
if z=$(curl -s 'https://install.zerotier.com/' | gpg); then echo "$z" | sudo bash; fi

sudo systemctl enable zerotier-one

sudo zerotier-cli status

sudo zerotier-cli join [Network ID]

Authenticate your device by going to https://my.zerotier.com/network/[Network ID]
```

# AS VPN Server

```
sudo vim /etc/sysctl.conf
net.ipv4.ip_forward = 1

sudo sysctl -p /etc/sysctl.conf
```

# update

```
apt update ; apt upgrade
apt install git vim curl neovim tmux
```

# WIFI

[Setting WiFi up via the command line - Raspberry Pi Documentation](https://www.raspberrypi.org/documentation/configuration/wireless/wireless-cli.md)

You will need to define a `wpa_supplicant.conf` file for your particular wireless
network. Put this file in the boot folder, and when the Pi first boots, it will
copy that file into the correct location in the Linux root file system and use
those settings to start up wireless networking.
```
echo "country=GB" >> /etc/wpa_supplicant/wpa_supplicant.conf

iwlist wlan0 scan | grep SSID
wpa_passphrase "<SSID>" | sudo tee -a /etc/wpa_supplicant/wpa_supplicant.conf

wpa_cli -i wlan0 reconfigure
```

**priority**

```
network={
    ssid="HomeOneSSID"
    psk="passwordOne"
    priority=1
    id_str="homeOne"
}

network={
    ssid="HomeTwoSSID"
    psk="passwordTwo"
    priority=2
    id_str="homeTwo"
    }
```
