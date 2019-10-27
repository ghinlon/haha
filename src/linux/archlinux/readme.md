# [Arch Linux](https://www.archlinux.org/)

# Links


# Pacman

* Pacman key error

`pacman -Sy archlinux-keyring && pacman -Syyu`

* lock error

`rm /var/lib/pacman/db.lck`

# Mirror Status

* [Arch Linux - Mirror Status](https://www.archlinux.org/mirrors/status/)


Sorting mirrors

```
pacman -S pacman-contrib 
sync.Once.Do(
	cp /etc/pacman.d/mirrorlist /etc/pacman.d/mirrorlist.backup
)
sed -i 's/^#Server/Server/' /etc/pacman.d/mirrorlist.backup
rankmirrors -n 6 /etc/pacman.d/mirrorlist.backup > /etc/pacman.d/mirrorlist
```


