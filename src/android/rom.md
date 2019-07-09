# ROM

# 一个ROM包里有什么

**What does new Android zips (full roms, but also otas) contain?**

New Android flashable zips are made this way:

boot.img (kernel)
file_contexts (selinux related)
META-INF (folder containing scripts)
system.new.dat (compressed /system partition)
system.patch.dat (for OTAs)
system.transfer.list (see explanation below)

# 怎么提取system.new.dat里的东西?

因为我不小心在`/system/fonts/`目录里执行了`rm -f *`,所以我需要抽取出里面的字体文件.

使用[xpirt/sdat2img](git@github.com:xpirt/sdat2img.git):

# Example

This is a simple example on a Linux system:

```bash
~$ ./sdat2img.py system.transfer.list system.new.dat system.img
~$ sudo mount -o loop -t ext4 system.img /mnt/
```























#### 参考链接
https://forum.xda-developers.com/android/software-hacking/how-to-conver-lollipop-dat-files-to-t2978952
