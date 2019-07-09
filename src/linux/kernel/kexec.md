# kexec - directly boot into a new kernel 

# Links

* [kexec(8): directly boot into new kernel - Linux man page](https://linux.die.net/man/8/kexec)

# Usage

 Using **kexec** consists of:

1. loading the kernel to be rebooted to into memory, and
1. actually rebooting to the pre-loaded kernel.

```
kexec -l kernel-image --append=command-line-options --initrd=initrd-image 

// After this kernel is loaded, it can be booted to at any time using the command:
kexec -e

		-e (--exec)
			Run the currently loaded kernel. 
		-l (--load) kernel
			Load the specified kernel into the current kernel. 

		--append=string
				Append string to the kernel command line.

		--command-line=string
				Set the kernel command line to string.

		--reuse-cmdline
				Use the command line from the running system. When a panic kernel is loaded, it strips the crashkernel parameter automatically. The BOOT_IMAGE parame‐
				ter is also stripped.
```

# Example

With this, I can pxe boot into a raid-tool kernel, after raided, then `kexec`
into the real installation kernel. doesn't need to boot twice.



