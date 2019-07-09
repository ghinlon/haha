# Networking

`systemctl start|stop|restart|status network`

# Links

* [Chapter 3. Configuring IP Networking - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/networking_guide/ch-configuring_ip_networking)

# Configuring IP Networking with nmcli

`nmcli [OPTIONS] OBJECT { COMMAND | help }`

where OBJECT can be one of the following options: `general`, `networking`,
`radio`, `connection`, `device`, `agent`, and `monitor`. 

```
nmcli connection modify eth0 ipv4.address 192.168.122.88/24

// reload config file
nmcli connection reload
nmcli con load /etc/sysconfig/network-scripts/ifcfg-ifname

// down and up
nmcli dev disconnect interface-name
nmcli con up interface-name
```

