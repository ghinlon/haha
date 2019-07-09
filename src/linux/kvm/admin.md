# Administration

# Links

* [Red Hat Enterprise Linux 7 Virtualization Deployment and Administration Guide - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/)

# Add a second NIC for a libvirt guest


1. `virsh domiflist vm1`
2. 
	```
	virsh attach-interface --domain vm1 --type network \
			--source openstackvms --model virtio \
			--mac 52:54:00:4b:73:5f --config --live
	```  

* [linux - What is the difference between virbr# and vnet#? - Unix & Linux Stack Exchange](https://unix.stackexchange.com/questions/52855/what-is-the-difference-between-virbr-and-vnet)
* [kashyapc.fedorapeople.org/virt/add-network-card-in-guest.txt](https://kashyapc.fedorapeople.org/virt/add-network-card-in-guest.txt)

# CREATING A VIRTUAL MACHINE

#### Storage Pools

```
// <dp> is the name of the pool.
virsh pool-define-as dp --type dir --target /mnt/vm/dirpool/

//this command will create the <dirpool> dir for us. we don't need create the dir by ourself.
virsh pool-build dp

virsh pool-start dp
virsh pool-autostart dp
virsh pool-info dp
virsh pool-list --all
virsh pool-destroy dp

// will delete the dir
virsh pool-delete dp

virsh pool-undefine dp
```

#### Storage Volumes 


```
# virsh vol-create-as dp v1 10G --format qcow2
Vol v1 created

# virsh vol-create-as dp v2 10G 
Vol v2 created

# virsh vol-info dp v1
error: failed to get pool 'v1'
error: failed to get vol 'dp', specifying --pool might help
error: Storage volume not found: no storage vol with matching path

# virsh vol-info v1 dp
Name:           v1
Type:           file
Capacity:       10.00 GiB
Allocation:     196.00 KiB

# virsh vol-info v2 dp
Name:           v2
Type:           file
Capacity:       10.00 GiB
Allocation:     10.00 GiB

# ls -l dirpool/*
-rw------- 1 root root      197120 Mar 27 21:37 dirpool/v1
-rw------- 1 root root 10737418240 Mar 27 21:38 dirpool/v2
# du -sh dirpool/*
196K    dirpool/v1
11G     dirpool/v2

[root@host dirpool]# qemu-img info v1
image: v1
file format: qcow2
virtual size: 10G (10737418240 bytes)
disk size: 196K
cluster_size: 65536
[root@host dirpool]# qemu-img info v2
image: v2
file format: raw
virtual size: 10G (10737418240 bytes)
disk size: 10G
```

#### Convert vm image

```
# qemu-img info vdisk.img 
image: vdisk.img
file format: raw
virtual size: 10G (10737418240 bytes)
disk size: 5.0G
# qemu-img convert -f raw -O qcow2 vdisk.img vdisk.qcow2
```

#### Re-sizing the Disk Image

` # qemu-img resize filename size`

You can also resize relative to the current size of the disk image.   

`# qemu-img resize filename [+|-]size[K|M|G|T]``  

**This command doesn't affect `virsh volume` direct.**  

     Before using this command to shrink a disk image, you must use file system and partitioning tools inside the VM itself to reduce allocated file systems and partition sizes accordingly. Failure to do so will result in data loss.
    After using this command to grow a disk image, you must use file system and partitioning tools inside the VM to actually begin using the new space on the device. 




# Links

https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/pdf/virtualization_deployment_and_administration_guide/Red_Hat_Enterprise_Linux-7-Virtualization_Deployment_and_Administration_Guide-en-US.pdf

