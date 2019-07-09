# KVM

# Links

* [KVM_Virtualization_in_RHEL_7_Made_Easy](https://linux.dell.com/files/whitepapers/KVM_Virtualization_in_RHEL_7_Made_Easy.pdf) **Excenlent**
* [Virtualization Getting Started Guide - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_getting_started_guide/index)
* [Virtualization Deployment and Administration Guide - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/virtualization_deployment_and_administration_guide/index)

# kvm vs libvirt vs qemu

kvm virtualize cpu and mem

qemu virtualize others except cpu and mem

libvirt is an interface, an api.

libvirt supports qemu-kvm, vmware, virtualbox and xen.

# Begin

* Verify your server support VT(Virtualization Technology)

```
grep -E 'svm|vmx' /proc/cpuinfo

	- vmx is for Intel processors
	- svm is for AMD processors
```

* Install KVM software

```
// on rhel7
yum install qemu-kvm libvirt libvirt-python libguestfs-tools virt-install virt-top
systemctl enable libvirtd 
systemctl start libvirtd

// verify kvm modules are loaded
# lsmod | grep kvm
kvm_intel              53484  3 
kvm                   316506  1 kvm_intel

// enable ip_forward
# grep ip_forward /etc/sysctl.conf
net.ipv4.ip_forward = 1

// load config 
# sysctl -p 

// print the full list of variant
// when creating vms, pick a suitable one.
virt-install --os-variant list
```

# Creating VMs

If no network option is specified, the guest virtual machine is configured with a default network with NAT.

```sh
virt-install -v \
--os-type=linux \
--os-variant=rhel7 \
--name=centos7.6 \
--vcpus=1 \
--memory=1024 \
--nographics --extra-args='console=ttyS0' \
--location=CentOS-7-x86_64-DVD-1810.iso
--disk=centos7.6.1810.img,size=10 \
```

## virt-viewer 

```
virt-viewer [OPTION...] DOMAIN-NAME|ID|UUID - Virtual machine graphical console
```

Getting X11 forwarding through ssh working after running su:

```
ssh -X user@hostname
su -
xauth merge /home/user/.Xauthority
virt-view <id>
```
There's one time I have trouble with keyboard when use `virt-view <id>`, I simply change to `vnc-view :5900`, then everything's ok.

# virsh - management user interface

```
// exit with ctrl-]
# console vm1

# list --all
# destroy vm1 
# undefine vm1

/*
      shutdown domain [--mode acpi|agent]
           Gracefully shuts down a domain.
       destroy domain [--graceful]
           Immediately terminate the domain domain.  This doesn’t give the domain OS any chance to react, and it’s the equivalent of ripping the power cord out
           on a physical machine.  In most cases you will want to use the shutdown command instead.  However, this does not delete any storage volumes used by
           the guest, and if the domain is persistent, it can be restarted later.
       undefine domain [--managed-save] [--snapshots-metadata] [ {--storage volumes | --remove-all-storage} --wipe-storage]
           Undefine a domain. If the domain is running, this converts it to a transient domain, without stopping it. If the domain is inactive, the domain
           configuration is removed.

simplely saying, destroy is same to "poweroff",  undefine is same to "delete".

*/


# dominfo vm1
# shutdown vm1
# start vm1
# autostart [--diable] vm1


# domifaddr vm1	    // obtain ip
```

and

```
# virt-top
# virt-df vm1
```

# Cloning VMs

clone will have a new mac, but same ip, so we need to modify ip manually.

```
# virsh suspend vm1
# virt-clone --original vm1 -n vm1-clone -f vm/vm1-clone.img
# virsh resume vm1
```

# Attaching storage device to a VM

```
// vdb is the device name in vm1
# virsh attach-disk vm1 /dev/sdb vdb --driver qemu --mode shareable
# virsh detach-disk vm1 vdb
```

# Memory

You can dynamically change the memory in a VM up to what its **maximum memory setting** is . 

* reduce

```
# virsh
# dominfo vm1-clone 
...
Max memory:     1048576 KiB
Used memory:    1048576 KiB
...

# setmem vm1-clone 512000
```

* increase  

  **increase the amount above maximum, must shutdown first.**

```
# virsh shutdown vm1-clone
# virsh edit vm1-clone
// chage the value to
    <memory unit='KiB'>32000000</memory>

# virsh create /etc/libvirt/qemu/vm1-clone.xml 
Domain vm1-clone created from /etc/libvirt/qemu/vm1-clone.xml

# virsh setmem <id|name> --size 32G
// make persistent
# virsh setmem 8 --size 32G --config
```

# vCPUs

```
# virsh shutdown vm1-clone
// change the value
 <vcpu placement='static'>2</vcpu>

# virsh create /etc/libvirt/qemu/vm1-clone.xml 
Domain vm1-clone created from /etc/libvirt/qemu/vm1-clone.xml 
```

# Disk capacity

1. `fallocate -l 10G vm1-clone-add.img`
2. `virsh shutdown vm1-clone`
3. `virsh edit vm1-clone` 

    ```
        <disk type='file' device='disk'>
          <driver name='qemu' type='raw' cache='none'/>
          <source file='/mnt/vm/vm1-clone.img'/>
          <target dev='vda' bus='virtio'/>
          <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
        </disk>
    ```
    copy to below, and modify 3 place(file=, dev=, slot=):  

    ```
        <disk type='file' device='disk'>
          <driver name='qemu' type='raw' cache='none'/>
          <source file='/mnt/vm/vm1-clone-add.img'/>
          <target dev='vdb' bus='virtio'/>
          <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
        </disk>
    ```

    * **make sure that the name of the device (i.e. vdb) follows the first one in sequential order.**  
    * **in the address tag, use a unique slot address ( check the address tag of ALL devices , not just storage devices)**  

    as mine ,slot must change to `0x07`, but not `0x06`:

    ```
    error: XML error: Attempted double use of PCI Address '0:0:6.0'
    Failed. Try again? [y,n,f,?]:
    Domain vm1-clone XML configuration edited.
    ```

4. `virsh create /etc/libvirt/qemu/vm1-clone.xml`


# Deleting VMs

```sh
virsh shutdown vm1-clone
virsh destroy vm1-clone
virsh undefine vm1-clone
rm /mnt/vm/vm1-clone.img
rm /mnt/vm/vm1-clone-add.img
```
