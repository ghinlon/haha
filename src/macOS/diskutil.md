# diskutil

# Links

# Basics

```
diskutil list
diskutil listfilesystems

diskutil partitionDisk disk2 GPT jhfs+ Foo 100%
```

# Delete partition

Since diskutil has no ablity to delete partition, Here is a hack:

```
sudo dd if=/dev/zero of=/dev/disk2 bs=512 count=1
```

