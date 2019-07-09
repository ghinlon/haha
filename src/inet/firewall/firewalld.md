# Firewalld

# Links

* [How to Install, Configure and Use Firewalld in CentOS and Ubuntu](https://www.tecmint.com/install-configure-firewalld-in-centos-ubuntu/)

# Install 

```
// Ubuntu 18.04
sudo apt install firewalld
sudo systemctl enable firewalld
sudo systemctl start firewalld

ufw disable
```

# Basic

* First. We must set `Zone`.
* Then, we add port or service to the zone.

Check

```sh
firewall-cmd --state
firewall-cmd --get-zones
firewall-cmd --get-default-zone
firewall-cmd --get-active-zones
firewall-cmd --get-services
firewall-cmd --list-services
firewall-cmd --list-services --permanent
# find rules associated with the current zone
firewall-cmd --list-all
firewall-cmd --list-ports
firewall-cmd --list-all-zones | less
firewall-cmd --info-zone external
// or 
firewall-cmd --zone=external --list-all

```

Set

```sh
firewall-cmd --reload
```

Zone Config

```sh
firewall-cmd --set-default-zone=external
firewall-cmd --zone=external --add-interface=ens4
firewall-cmd --zone=internal --add-interface=wg0
// --{change,remove}-interface
```

Add

```sh
firewall-cmd --zone=external --add-service={http,https} --permanent
firewall-cmd --zone=external --add-port={465,8080,8443,12345}/tcp --permanent
firewall-cmd --zone=external --add-port={465,8080,8443,12345}/udp --permanent
// --remove-{service,port}
```

# Zones

predefined zones:

* **drop**:		The lowest level of trust. All incoming connections are dropped without reply.
* **block**:	incomming requests are rejected with an `icmp-host-prohibited` or `icmp6-adm-prohibited`
* **public**:	Represents public, untrusted networks. 
* **external**:	when using the firewall as your gateway. It is configured for NAT masquerading.
* **internal**:	The other side of the external zone.
* **dmz**:		DMZ
* **work**:		trust most of the computers in the network.
* **home**:		ditto.
* **trusted**:	Trust all of the machines in the network.

