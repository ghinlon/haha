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

# run script at boot in initrd

Just put your script in `/etc/rc.local`



