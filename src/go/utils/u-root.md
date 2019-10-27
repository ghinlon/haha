# u-root

# Links

* [u-root/u-root](https://github.com/u-root/u-root)

# Howto

```
go get github.com/u-root/u-root

// arch
sudo pacman -S qemu

qemu-system-x86_64 \
	-kernel path/to/kernel \
	-initrd /tmp/initramfs.linux_amd64.cpio \
	-append 'console=ttyS0' \
	--nographic  \
	-serial mon:stdio 
```


# Issue with loading modules 

* [modprobe fails inside u-root shell · Issue #1214 · u-root/u-root · GitHub](https://github.com/u-root/u-root/issues/1214)

```
cp -RLp /lib/modules/`uname -r` /tmp/modules
unxz `find -name '*.ko.xz' /tmp/modules`
./u-root -format=cpio -build=source -files `which depmod` -files `which modprobe` -files /lib/modules/`uname -r` -files /tmp/modules:mytmp -files /usr/bin/vim -files /bin/nano -files /usr/share/terminfo/l/linux -o initramfs-sluinit-files.cpio ./cmds/* ./examples/sluinit
```

# How to add init scripts or extend init process?

* [How to add init scripts or extend init process? · Issue #569 · u-root/u-root · GitHub](https://github.com/u-root/u-root/issues/569#issuecomment-359951972)

Use [hashicorp/go-plugin: Golang plugin system over RPC.](https://github.com/hashicorp/go-plugin)

or

Include any Go command named `uinit`.

And this:

[U-root Init Proposal](https://github.com/u-root/u-root/issues/60#issuecomment-364807661)

## type Servicer interface

```go
type Servicer interface {
	Start() error
	Stop() error
	Reload() error
	Restart() error
        Status() state.Value
        Unit() Unit
}
```

Currently state:

Sat Oct 19 22:46:29 CST 2019

* [bb mode: move init out of bb-specific code · Issue #615 · u-root/u-root · GitHub](https://github.com/u-root/u-root/issues/615)
* [De-couple bb mode from init and rush. Add standalone makebb tool. by hugelgupf · Pull Request #708 · u-root/u-root · GitHub](https://github.com/u-root/u-root/pull/708)


# uinit

* [systemboot](https://github.com/u-root/u-root#systemboot)


