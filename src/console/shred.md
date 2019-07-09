# Shred

# Links

* [shred (Unix) - Wikipedia](https://en.wikipedia.org/wiki/Shred_(Unix))
* [shred(1) - Linux man page](https://linux.die.net/man/1/shred)
* [macos - Using `shred` from the command line - Super User](https://superuser.com/questions/617515/using-shred-from-the-command-line)

```
brew install coreutils
```

# HowTo

```
shred <foofile>
```

因为现在各种文件系统都会自动备份，所以这个方法不一定起作用。

可以写入之后不要删除，这样就会被同步到备份里去，过段时间再删除

正确的做法是**MUST**文件系统加密



