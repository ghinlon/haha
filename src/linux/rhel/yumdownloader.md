# How do I use yum to download a package without installing it?

# Links

* [How to use yum to download a package without installing it - Red Hat Customer Portal](https://access.redhat.com/solutions/10154)

# Environment

* Red Hat Enterprise Linux (RHEL) 7
* Red Hat Enterprise Linux 6
* Red Hat Enterprise Linux 5

# Issue

* How do I use yum to download a package without installing it?

# Resolution

There are two ways to download a package without installing it.

One is using the "downloadonly" plugin for yum, the other is using "yumdownloader" utility.

```
yum install yum-plugin-downloadonly
yum install --downloadonly --downloaddir=<directory> <package>


Note:

    Before using the plugin, check /etc/yum/pluginconf.d/downloadonly.conf to confirm that this plugin is "enabled=1"
    This is applicable for "yum install/yum update" and not for "yum groupinstall". Use "yum groupinfo" to identify packages within a specific group.
    If only the package name is specified, the latest available package is downloaded (such as sshd). Otherwise, you can specify the full package name and version (such as httpd-2.2.3-22.el5).
    If you do not use the --downloaddir option, files are saved by default in /var/cache/yum/ in rhel-{arch}-channel/packages
    If desired, you can download multiple packages on the same command.
    You still need to re-download the repodata if the repodata expires before you re-use the cache. By default it takes two hours to expire.
```

# Yumdownloader

If downloading a installed package, "yumdownloader" is useful.

```
yum install yum-utils
yumdownloader <package>
```

Note:

* The package is saved in the current working directly by default; use the
  `--destdir` option to specify an alternate location.
* Be sure to add `--resolve` if you need to download dependencies.

