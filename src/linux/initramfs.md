# initramfs

# Links

* [How to unpack/uncompress and repack/re-compress an initial ramdisk (initrd/initramfs) boot image file? - Red Hat Customer Portal](https://access.redhat.com/solutions/24029)
* [Using the initial RAM disk (initrd) — The Linux Kernel
  documentation](https://www.kernel.org/doc/html/latest/admin-guide/initrd.html)(Good)


# Resolution

* Identify compression format of the image  
	The most common is a gzip-format image, However, there may also be an  
    XZ/LZMA-format image  
	```
	file /boot/initramfs-($uname -r).img
	```
* Extract / Uncompress  
	```
	zcat /boot/initrd-$(uname -r).img | cpio -idmv
	xz -dc < /boot/initrd-$(uname -r).img | cpio -idmv

	cpio --option:
	   -i, --extract
		             Copy-in. 
   		Operation modifiers valid in copy-in and copy-pass modes
       -d, --make-directories
		             Create leading directories where needed.
       -m, --preserve-modification-time
	                 Retain previous file modification times when creating
					 files.
       -v, --verbose
	                 Verbosely list the files processed.

       -o, --create
	                 Copy-out. 
       -c     Use the old portable (ASCII) archive format. 
	
	xz --option:

	   -d, --decompress, --uncompress
       -c, --stdout, --to-stdout

	For rhel7:
	/usr/lib/dracut/skipcpio /boot/initramfs-$(uname -r).img | zcat | cpio
	-idmv
	```
* Repack / Recompress  
	```
	find . | cpio -o -c | gzip -9 > new.img
	find . 2>/dev/null | cpio -c -o | xz -9 --format=lzma > new.img
	```
