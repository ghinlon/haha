# Linux interfaces for virtual networking

# Links

* [Introduction to Linux interfaces for virtual networking - Red Hat Developer Blog](https://developers.redhat.com/blog/2018/10/22/introduction-to-linux-interfaces-for-virtual-networking/)

# Bridge

A Linux bridge behaves like a network switch. 


Here’s how to create a bridge:

```
# ip link add br0 type bridge
# ip link set eth0 master br0
# ip link set tap1 master br0
# ip link set tap2 master br0
# ip link set veth1 master br0
```

This creates a bridge device named br0 and sets two TAP devices (tap1, tap2),
a VETH device (veth1), and a physical device (eth0) as its slaves, as shown in
the diagram above.

![bridge](img/bridge.png)


# Bonded interface

[Linux Ethernet Bonding Driver HOWTO](https://www.kernel.org/doc/Documentation/networking/bonding.txt)

The Linux bonding driver provides a method for aggregating multiple network
interfaces into a single logical “bonded” interface. 

Here’s how to create a bonded interface:

```
ip link add bond1 type bond miimon 100 mode active-backup
ip link set eth0 master bond1
ip link set eth1 master bond1
```

![bond](img/bond.png)

# Team device

[Bonding vs. Team features · jpirko/libteam Wiki · GitHub](https://github.com/jpirko/libteam/wiki/Bonding-vs.-Team-features)

Similar a bonded interface

But there are also some functional differences between a bonded interface and
a team. For example, a team supports LACP load-balancing, NS/NA (IPV6) link
monitoring, D-Bus interface, etc., which are absent in bonding. 


Here’s how to create a team:

```
# teamd -o -n -U -d -t team0 -c '{"runner": {"name": "activebackup"},"link_watch": {"name": "ethtool"}}'
# ip link set eth0 down
# ip link set eth1 down
# teamdctl team0 port add eth0
# teamdctl team0 port add eth1
```

# VLAN


Use a VLAN when you want to separate subnet in VMs, namespaces, or hosts.

Here’s how to create a VLAN:

```
# ip link add link eth0 name eth0.2 type vlan id 2
# ip link add link eth0 name eth0.3 type vlan id 3
```

The VLAN header looks like:

![vlan](img/vlan_01.png)

# VXLAN

VXLAN (Virtual eXtensible Local Area Network) is a tunneling protocol designed
to solve the problem of limited VLAN IDs (4,096) in IEEE 802.1q.

With a 24-bit segment ID, aka VXLAN Network Identifier (VNI), VXLAN allows up
to 2^24 (16,777,216) virtual LANs, which is 4,096 times the VLAN capacity.

VXLAN encapsulates Layer 2 frames with a VXLAN header into a UDP-IP packet,
which looks like this:

![vxlan](img/vxlan_01.png)

Here’s how to use VXLAN:

```
# ip link add vx0 type vxlan id 100 local 1.1.1.1 remote 2.2.2.2 dev eth0 dstport 4789
```

# MACVLAN

With VLAN, you can create multiple interfaces on top of a single one and filter
packages based on a VLAN tag. With MACVLAN, you can create multiple interfaces
with different Layer 2 (that is, Ethernet MAC) addresses on top of a single
one.

Use a MACVLAN when you want to connect directly to a physical network from
containers.

# IPVLAN

IPVLAN is similar to MACVLAN with the difference being that the endpoints have
the same MAC address.

IPVLAN supports L2 and L3 mode.

![ipvlan](img/ipvlan.png)

# MACVTAP/IPVTAP

![macvtap](img/macvtap.png)


Typically, MACVLAN/IPVLAN is used to make both the guest and the host show up
directly on the switch to which the host is connected. The difference between
MACVTAP and IPVTAP is same as with MACVLAN/IPVLAN.

Here’s how to create a MACVTAP instance:

```
# ip link add link eth0 name macvtap0 type macvtap
```

# MACsec

The main use case for MACsec is to secure all messages on a standard LAN
including ARP, NS, and DHCP messages.

# VETH

The VETH (virtual Ethernet) device is a local Ethernet tunnel. Devices are
created in pairs.

Packets transmitted on one device in the pair are immediately received on the
other device. When either device is down, the link state of the pair is down.

![veth](img/veth.png)

Use a VETH configuration when namespaces need to communicate to the main host
namespace or between each other.


Here’s how to set up a VETH configuration:

```
# ip netns add net1
# ip netns add net2
# ip link add veth1 netns net1 type veth peer name veth2 netns net2
```
# VCAN

CAN (Controller Area Network) 


Here’s how to create a VCAN:

```
# ip link add dev vcan1 type vcan
```

# VXCAN

Similar to the VETH driver, a VXCAN (Virtual CAN tunnel) implements a local CAN
traffic tunnel between two VCAN network devices.


Here’s how to set up a VXCAN instance:

```
# ip netns add net1
# ip netns add net2
# ip link add vxcan1 netns net1 type vxcan peer name vxcan2 netns net2
```

# IPOIB

An IPOIB device supports the IP-over-InfiniBand protocol. 

Use an IPOIB device when you have an IB device and want to communicate with
a remote host via IP.

Here’s how to create an IPOIB device:

```
# ip link add ipoib0 type ipoib mode connected
```

# NLMON

NLMON is a Netlink monitor device.


Here’s how to create an NLMON device:

```
# ip link add nlmon0 type nlmon
# ip link set nlmon0 up
# tcpdump -i nlmon0 -w nlmsg.pcap
```

# Dummy interface

Here’s how to create a dummy interface:

```
# ip link add dummy1 type dummy
# ip addr add 1.1.1.1/24 dev dummy1
# ip link set dummy1 up
```

# IFB

# netdevsim interface


