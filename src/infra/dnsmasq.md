# dnsmasq

# Links

* [dnsmasq - ArchWiki](https://wiki.archlinux.org/index.php/Dnsmasq)
* [How to avoid conflicts between dnsmasq and systemd-resolved? - Unix & Linux Stack Exchange](https://unix.stackexchange.com/questions/304050/how-to-avoid-conflicts-between-dnsmasq-and-systemd-resolved)

#  to disable dnsmasq's DNS server functionality

```
# To disable dnsmasq's DNS server functionality.
port=0
```

# PXE-enabled DHCP

* [matchbox/network-setup.md at master · coreos/matchbox · GitHub](https://github.com/coreos/matchbox/blob/master/Documentation/network-setup.md)
* [cobbler/dnsmasq.template at master · cobbler/cobbler · GitHub](https://github.com/cobbler/cobbler/blob/master/templates/etc/dnsmasq.template)
* [PXE-enabled DHCP](https://github.com/poseidon/matchbox/blob/master/Documentation/network-setup.md#pxe-enabled-dhcp)

Add [ipxe.efi](http://boot.ipxe.org/ipxe.efi) and [unidonly.kpxe](http://boot.ipxe.org/undionly.kpxe) to your tftp-root (e.g. `/var/lib/tftpboot`).

## Tips

```
dhcp-ignore=tag:!known
# Set the DHCP server to authoritative mode. In this mode it will barge in
# and take over the lease for any client which broadcasts on the network,
# whether it has a record of the lease or not. This avoids long timeouts
# when a machine wakes up on a new network. DO NOT enable this if there's
# the slightest chance that you might end up accidentally configuring a DHCP
# server for your campus/company accidentally. The ISC server uses
# the same option, and this URL provides more information:
# http://www.isc.org/files/auth.html
dhcp-authoritative

dhcp-range=set:red,192.168.1.1,192.168.1.254,30m
# router
dhcp-option=tag:red,3,192.168.1.254

dhcp-host=11:22:33:44:55:66,12:34:56:78:90:12,192.168.1.60

enable-tftp
tftp-root=/var/lib/tftpboot

# Legacy PXE
dhcp-match=set:bios,option:client-arch,0
dhcp-boot=tag:bios,undionly.kpxe

# UEFI
# If you want to use a separate TFTP server instead of dnsmasq, specify its IP address after the boot-loader path
# dhcp-boot=tag:efi32,ipxe.efi,192.168.1.101
# and it seems we always should put tftp server address here when on UEFI
dhcp-match=set:efi32,option:client-arch,6
dhcp-boot=tag:efi32,ipxe.efi,192.168.1.100
dhcp-match=set:efibc,option:client-arch,7
dhcp-boot=tag:efibc,ipxe.efi,192.168.1.100
dhcp-match=set:efi64,option:client-arch,9
dhcp-boot=tag:efi64,ipxe.efi,192.168.1.100

# iPXE - chainload to matchbox ipxe boot script
dhcp-userclass=set:ipxe,iPXE
dhcp-boot=tag:ipxe,http://matchbox.example.com:8080/boot.ipxe

# verbose
log-queries
log-dhcp

# static DNS assignements
address=/matchbox.example.com/192.168.1.100

# (optional) disable DNS and specify alternate
# port=0
# dhcp-option=6,192.168.1.100
```

# tftp server address when on UEFI

it seems we always should put tftp server address here when on UEFI

```
dhcp-boot=tag:efi32,ipxe.efi,192.168.1.101
```

# Configurable TFTP

[iPXE - open source boot firmware](https://ipxe.org/start)

If your DHCP server is configured to network boot PXE clients (but not iPXE
clients), add a pxelinux.cfg to serve an iPXE kernel image and append commands.

**Only** when you can't control dhcpd's `dhcp-boot`

Example `/var/lib/tftpboot/pxelinux.cfg/default`:

```
timeout 10
default iPXE
LABEL iPXE
KERNEL ipxe.krn
APPEND dhcp && chain http://matchbox.example.com:8080/boot.ipxe
```

Add [ipxe.krn](http://boot.ipxe.org/ipxe.iso) to `/var/lib/tftpboot` ,  
get it from [http://boot.ipxe.org/ipxe.iso](http://boot.ipxe.org/ipxe.iso)


