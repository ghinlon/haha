# modules

# Links

* [Linux: How to load a kernel module automatically at boot time - nixCraft](https://www.cyberciti.biz/faq/linux-how-to-load-a-kernel-module-automatically-at-boot-time/)

# modprobe at startup

```
# cat /etc/modules-load.d/modules.conf 

# /etc/modules: kernel modules to load at boot time.
#
# This file contains the names of kernel modules that should be loaded
# at boot time, one per line. Lines beginning with "#" are ignored.
```

# Basic

```
modprobe {module_name}
lsmod
modinfo {module_name} 				
rmmod {module_name} 
// or 
modprobe -r {module_name}
```



