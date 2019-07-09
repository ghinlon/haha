# 安装archlinux

在U盘上安装

#### 参考链接

https://wiki.archlinux.org/index.php/Installing_Arch_Linux_on_a_USB_key
https://wiki.archlinux.org/index.php/installation_guide
https://wiki.archlinux.org/index.php/Improving_performance#Reduce_disk_reads.2Fwrites
http://www.wyae.de/docs/boot-usb3/
https://wiki.archlinux.org/index.php/General_recommendations
https://wiki.archlinux.org/index.php/List_of_applications
https://wiki.archlinux.org/index.php/Wayland
https://wiki.archlinux.org/index.php/GRUB


#### 连网

* 有线一般通过 dhcpcd eth0 连接网络
* 配置代理

  ```bash
  export http_proxy="https://username:password@server:port"
  ```

#### 创建文件系统

##### 分区及文件系统

为了兼容性，所以采用msdos的分区表。  
为了仍然可以有部分容量作为U盘使用，所以系统安装在第二个分区,因为windows只认U盘的第一个分区。
给第一个分区创建exfat文件系统，这样可以放超过4G大小的文件.  

  ```bash
parted -s -- /dev/sda mkpart 1m 32g
parted -s -- /dev/sda mkpart 32g -1
parted -s -- /dev/sda set 2 boot on
mkfs.exfat /dev/sda1
mkfs.ext4 /dev/sda2
mount /dev/sda2 /mnt
```

#### pacstrap

`pacstrap /mnt`

#### 配置系统

##### 生成fstab

  `# genfstab -U /mnt >> /mnt/etc/fstab`
##### chroot

  `# arch-chroot /mnt`

##### 安装基础软件包

  目标是重启后可以连网  
```bash
pacman -S privoxy shadowsocks-libev simple-obfs proxychains-ng kcptun iw wpa_supplicant dhcpcd
```
##### 时区

  ```bash
# ln -sf /usr/share/zoneinfo/Region/City /etc/localtime
# hwclock --systohc
```

##### locale

  编辑/etc/locale.gen，然后执行 `# locale-gen`:  

  ```bash
/etc/locale.conf
LANG=en_US.UTF-8
```

##### 主机名

  ```bash
# cat /etc/hostname
myhostname
```

##### Initramfs (重要，和常规有变化)

  Before creating the initial RAM disk # mkinitcpio -p linux, in /etc/mkinitcpio.conf add the block hook to the hooks array right after udev. This is necessary for appropriate module loading in early userspace.

  要hooks里udev后面马上跟block

  `# mkinitcpio -p linux`

##### 配置引导程序

  ```bash
pacman -S grub
# grub-install --target=i386-pc /dev/sda --removable --recheck
# grub-mkconfig -o /boot/grub/grub.cfg
```

##### 配置root密码，添加用户，配置用户密码

  ```bash
passwd
# useradd -m -G wheel -s /bin/bash <username>
passwd <username>
```

##### 重启

  ```bash
# umount /mnt
reboot
```

#### 第一次进入系统

##### [配置网络]()

##### [配置代理]()

##### 安装基本软件包

  ```bash
pacman -S base-devel dosfstools ntfs-3g exfat-utils openssh screen gvim p7zip zip unzip git gnupg  lsof  ntp gnu-netcat nmap cronie bc dstat bash-completion
```

##### 安装图形

  ```bash
pacman -S  xf86-input-synaptics xf86-video-vesa  xf86-video-ati xf86-video-intel  xf86-video-nouveau weston xorg-server-xwayland xorg-xrdb gtk3  sway rxvt-unicode dmenu alsa-utils alsa-firmware libsamplerate pulseaudio pulseaudio-alsa pavucontrol fcitx noto-fonts noto-fonts-cjk
```

##### 安装多媒体工具

  ```bash
pacman -S mpv imv
```

##### 为了访问Android

  ```bash
pacman -s android-tools android-udev
# gpasswd -a <username> adbusers
```

##### 文档处理

  ```bash
pacman -S mupdf
```
