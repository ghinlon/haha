# qemu

# Links

* [QEMU - ArchWiki](https://wiki.archlinux.org/index.php/QEMU)

# Install

```
// archlinux
pacman -S qemu
```

# By specifying kernel and initrd manually

```
qemu-system-x86_64  \
	-kernel /boot/vmlinuz-linux \
	-initrd /boot/initramfs-linux.img  \
	-append 'console=ttyS0 root=/dev/sda1' \
	--nographic \
	-serial mon:stdio \
	/dev/sda3 

// In the above example, the physical partition being used for the guest's root
// file system is /dev/sda3 on the host, but it shows up as /dev/sda on the guest.

```

