# Mac

# Index

* [New to MAC](newtomac.md)
* [homebrew](homebrew.md)
* [MacPorts](macports.md)
* [More Downloads for Apple Developers](developers.md)
* [Remapping Caps Lock key to something more natural on Mac OS X](Remapping_Caps_Lock_key_to_something_more_natural_on_Mac_OS_X.md)

# Must 

## Stupid REPLACE my dir but not MERGE

  **SB**  
  Sun Jul 07 21:35:57 CST 2019

## Disable FaceTime

1. Open FaceTime
1. Click on the FaceTime menu in the top menu bar
1. Select the third option, Turn FaceTime Off

# Capture a Screen Shot

[Keyboard Shortcuts to Capture a Screen Shot with Mac OS X Mavericks and Yosemite | Information Technology Group](https://www.itg.ias.edu/content/keyboard-shortcuts-capture-screen-shot-mac-os-x)

* To capture the entire screen, press `Command-Shift-3`.
* To copy the entire screen, press `Command-Control-Shift-3`. 
* To capture a portion of the screen, press `Command-Shift-4`.
* To copy a portion of the screen to the clipboard, press `Command-Control-Shift-4`.


# How to Write to NTFS Drives on a Mac

* install ntfs-3g

`brew install ntfs-3g`

* mount

```
sudo mkdir /Volumes/ntfs1
diskutil list
```

The NTFS partition was probably automatically mounted by your Mac, so you’ll
need to unmount it first. Run the following command, replacing /dev/disk2s1
with the device name of your NTFS partition.

`sudo umount /dev/disk2s1`

* mount

`/usr/local/bin/ntfs-3g /dev/disk2s1 /Volumes/ntfs1 -olocal -oallow_other`

## Paragon

[Microsoft NTFS for Mac | Paragon Software](https://www.paragon-software.com/home/ntfs-mac/)

* Reset trial

[Manually remove Paragon NTFS v15 leftovers MacOS](https://gist.github.com/guycalledseven/7b3fbeb521f74c682799932d64856f03)

```
sudo su
rm -rf "/Library/Application Support/Paragon Software/"
rm /Library/LaunchDaemons/com.paragon-software.installer.plist
rm /Library/LaunchDaemons/com.paragon-software.ntfs.loader.plist
rm /Library/LaunchDaemons/com.paragon-software.ntfsd.plist
rm /Library/LaunchAgents/com.paragon-software.ntfs.notification-agent.plist
rm -rf /Library/PrivilegedHelperTools/com.paragon-software.installer
rm -rf /Library/PreferencePanes/ParagonNTFS.prefPane
```

# Finder is not responding

```
pgrep -fl Finder
kill <pid>
```

# If the sound is weird

```
sudo killall coreaudiod
```

# Config Gateway

 [macos - How to change the default gateway of a Mac OSX machine - Ask Different](https://apple.stackexchange.com/questions/33097/how-to-change-the-default-gateway-of-a-mac-osx-machine)

```
route delete default
route add default 192.168.0.1<Paste>
// or
$ route change default -interface $INTF
$ route change 192.168.0.0/16 -interface $INTF
```

# chattr use chflags

```
sudo chflags uimmutable /etc/resolv.conf
ls -lO /etc/resolv.conf
```
