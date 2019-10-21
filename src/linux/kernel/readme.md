# kernel

Links

* [The Linux Kernel Archives](https://www.kernel.org/)

# Install

* [Linux Kernel 5.3 released and here is how to install it - nixCraft](https://www.cyberciti.biz/linux-news/linux-kernel-5-3-released-and-here-is-how-to-install-it/)
* [Linux kernel release 5.x <http://kernel.org/> — The Linux Kernel documentation](https://www.kernel.org/doc/html/latest/admin-guide/README.html#configuring-the-kernel)

```
// rhel
sudo yum groupinstall "Development Tools"
sudo yum install ncurses-devel bison flex elfutils-libelf-devel openssl-devel bc


xz -cd linux-5.x.tar.xz | tar xvf -
cd linux-5.3
make defconfig
make
sudo make modules_install install
```

