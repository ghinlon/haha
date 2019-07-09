# iPXE

iPXE is the leading open source network boot firmware. 

# Links

* [iPXE - open source boot firmware](https://ipxe.org/howto/chainloading)

# Chainloading iPXE

Download [undionly.kpxe](http://boot.ipxe.org/undionly.kpxe) and save it to
your TFTP server directory.  Configure your DHCP server to hand out
`undionly.kpxe` as the boot file.

Example `/var/lib/tftpboot/pxelinux.cfg/default`:

```
timeout 10
default iPXE
LABEL iPXE
// May I change ipxe.lkrn to undionly.kpxe, so this is the right way ?
KERNEL ipxe.lkrn
APPEND dhcp && chain http://matchbox.example.com:8080/boot.ipxe
```

## Breaking the infinite loop

PXE will load iPXE which will load iPXE which will load iPXE which will load
iPXE… 

## UEFI

If you have machines which attempt to perform a UEFI network boot, then
download [ipxe.ef](http://boot.ipxe.org/ipxe.efi) and save it to your TFTP
server directory. 

**Note** that UEFI network booting tends to be substantially slower than BIOS
network booting, due to fundamental architectural limitations in UEFI. Most
UEFI systems provide the ability to perform a network boot in a BIOS
compatibility mode. You may wish to upgrade your system to use BIOS mode for
network booting. 



