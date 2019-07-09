# GRUB 2

# Links

* [Chapter 25. Working with GRUB 2 - Red Hat Customer Portal](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/ch-working_with_the_grub_2_boot_loader)

# Introduction

GRUB 2 reads its configuration from the `/boot/grub2/grub.cfg` file on
traditional BIOS-based machines and from the `/boot/efi/EFI/redhat/grub.cfg`
file on UEFI machines. This file contains menu information. 

# grubby

**worked on grub legacy(grub 1) too.**

```
// Listing the Default Kernel
grubby --default-kernel			

// To find out the index number of the default kernel
grubby --default-index			

// This command is very useful 
// you can then know which one to --set-default with this command.
// To list all the kernel menu entries
grubby --info=ALL

// To view the GRUB 2 menu entry for a specific kernel
grubby --info /boot/vmlinuz-3.10.0-229.4.2.el7.x86_64

// Changing the Default Boot Entry persistent
grubby --set-default /boot/vmlinuz-3.10.0-229.4.2.el7.x86_64
```

# GRUB Legacy (grub 1)

The configuration file is located at `/boot/grub/menu.lst`




