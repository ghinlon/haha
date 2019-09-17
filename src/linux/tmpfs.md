# memory based file system (RAM disk)

# Links

* [Create a RAM disk in Linux – JamesCoyle.net Limited](https://www.jamescoyle.net/how-to/943-create-a-ram-disk-in-linux)
* [The Difference Between a tmpfs and ramfs RAM Disk – JamesCoyle.net Limited](https://www.jamescoyle.net/knowledge/951-the-difference-between-a-tmpfs-and-ramfs-ram-disk)


# tmpfs vs ramfs

* ramfs is the older file system type and is largely replaced in most scenarios
  by tmpfs.

* ramfs file systems cannot be limited in size like a disk base file system
  which is limited by it’s capacity. ramfs will continue using memory storage
  until the system runs out of RAM and likely crashes or becomes unresponsive.
  This is a problem if the application writing to the file system cannot be
  limited in total size. Another issue is you cannot see the size of the file
  system in df and it can only be estimated by looking at the cached entry in
  free.

* You can specify a size limit in tmpfs which will give a ‘disk full’ error
  when the limit is reached. This behaviour is exactly the same as a partition
  of a physical disk.

# Create

```
mkdir /mnt/ramdisk
mount -t tmpfs -o size=512m tmpfs /mnt/ramdisk
```

make persist, `vim /etc/fstab:`

```
tmpfs       /mnt/ramdisk tmpfs   nodev,nosuid,noexec,nodiratime,size=1024M   0 0
```


