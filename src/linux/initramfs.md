# initramfs

# Links

* [How to unpack/uncompress and repack/re-compress an initial ramdisk (initrd/initramfs) boot image file? - Red Hat Customer Portal](https://access.redhat.com/solutions/24029)
* [Using the initial RAM disk (initrd) — The Linux Kernel documentation](https://www.kernel.org/doc/html/latest/admin-guide/initrd.html)(Good)
* [Jootamam - Howto create an initramfs image](https://www.jootamam.net/howto-initramfs-image.htm)

# Extract and Repack

* Identify compression format of the image  
	The most common is a gzip-format image, However, there may also be an  
    XZ/LZMA-format image  
	```
	file /boot/initramfs-($uname -r).img
	```

* gzip format

```
// Extract
zcat /boot/initrd-$(uname -r).img | cpio -idmv
// Repack
find . | cpio -o -c | gzip -9 > new.img
// For rhel7
/usr/lib/dracut/skipcpio /boot/initramfs-$(uname -r).img | zcat | cpio -idmv

cpio --option:
   -i, --extract
	             Copy-in. 
	Operation modifiers valid in copy-in and copy-pass modes
   -o, --create
                 Copy-out. 

   -d, --make-directories
	             Create leading directories where needed.
   -m, --preserve-modification-time
                 Retain previous file modification times when creating
				 files.
   -v, --verbose
                 Verbosely list the files processed.

   -c     Use the old portable (ASCII) archive format. 
```

* xz format

```
// Extract
xz -cd < /boot/initrd-$(uname -r).img | cpio -idmv
// Repack
find . 2>/dev/null | cpio -c -o | xz -9 --format=lzma > new.img

xz --option:

   -d, --decompress, --uncompress
   -c, --stdout, --to-stdout
```

# Custom initramfs

* [Custom Initramfs - Gentoo Wiki](https://wiki.gentoo.org/wiki/Custom_Initramfs)

minimalistic /init example

```sh
#!/bin/busybox sh

# Mount the /proc and /sys filesystems.
mount -t proc none /proc
mount -t sysfs none /sys

# Do your stuff here.
echo "This script just mounts and boots the rootfs, nothing else!"

# Mount the root filesystem.
mount -o ro /dev/sda1 /mnt/root

# Clean up.
umount /proc
umount /sys

# Boot the real thing.
exec switch_root /mnt/root /sbin/init
```

# Parse cmdline

* [Custom Initramfs/Examples - Gentoo Wiki](https://wiki.gentoo.org/wiki/Custom_Initramfs/Examples)

```
for x in $(cat /proc/cmdline); do
	case "${x}" in
		crypt_root=*)
			CRYPT_ROOT=${x#*=}
		;;
		net_ipv4=*)
			NET_IPv4=${x#*=}
		;;
		net_gw=*)
			NET_GW=${x#*=}
		;;
	esac
done
```

