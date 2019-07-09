# linux

# Mirrors List

* [CentOS](https://centos.org/download/mirrors/)

# pxeboot

* [CentOS](http://mirror.centos.org/centos/7/os/x86_64/images/pxeboot/)

# Distribution

* [Downloads | Ubuntu](https://www.ubuntu.com/download/alternative-downloads)
* [stable.release.core-os.net/amd64-usr/current/](https://stable.release.core-os.net/amd64-usr/current/)

# Date and Time

[Chapter 3. Configuring the Date and Time - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/chap-configuring_the_date_and_time)

```
timedatectl 
timedatectl list-timezones | grep Rome
timedatectl set-timezone “Asia/Kolkata”
timedatectl set-timezone UTC
timedatectl set-time 15:58:30
timedatectl set-time 20151120
timedatectl set-time '2015-11-20 16:14:50'

// Hardware Clock
timedatectl set-local-rtc 0		// set to utc

// must have NTP installed on the system
timedatectl set-ntp true

systemctl restart systemd-timedated.service


date
date --utc
date +"format"
date --set HH:MM:SS
date --set HH:MM:SS --utc
date --set YYYY-MM-DD
date --set "2017-06-02 23:26:00"

hwclock
hwclock --systohc
hwclock --systohc --utc
hwclock --hctosys
```

# Users and Groups

[Chapter 4. Managing Users and Groups - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/ch-managing_users_and_groups)

# Networking

* [Chapter 3. Configuring IP Networking - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/networking_guide/ch-configuring_ip_networking)

## Configuring IP Networking with nmcli


```
systemctl start|stop|restart|status network 

nmcli connection modify eth0 ipv4.address 192.168.122.88/24

// reload config file
nmcli connection reload
nmcli con load /etc/sysconfig/network-scripts/ifcfg-ifname

// down and up
nmcli dev disconnect interface-name
nmcli con up interface-name
```

# speedtest

```
curl -s https://raw.githubusercontent.com/sivel/speedtest-cli/master/speedtest.py | python -
```

# X11 Forwarding

rhel7

```sh
sudo yum install xauth

// cfg file: /etc/ssh/sshd_config

X11Forwarding yes
X11DisplayOffset 10
X11UseLocalhost no

// reload
systemctl restart sshd
```


# Network Scan

```
// -sn: Ping Scan - disable port scan
nmap -sn 192.168.1.0/24
```

# LS COLOR

[LSCOLORS Generator](https://geoff.greer.fm/lscolors/)

# PERMIT APP BIND SERVICE

```
setcap cap_net_bind_service+ep /path/to/bin
```

# mv: Directory not empty

[mv: Directory not empty](https://askubuntu.com/questions/269775/mv-directory-not-empty#269818)

Though its man page doesn't document it, mv will refuse to rename a directory
to another directory if the target directory contains files. This is a good
thing in your case because you turn out to want to merge the content of the
source into the target, which mv will not do.  Use:

`rsync -a backup/ backupArchives/`

instead. After that run:

`rm -rf backup/*`

Instead of using rsync, you also can do the classical

`(cd backup && tar c .) | (cd backupArchives && tar xf -)`

which earns you more geek points.


