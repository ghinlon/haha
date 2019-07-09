# LineageOS

# Links

* [TWRP OnePlus 3/3T](https://twrp.me/oneplus/oneplusthree.html)
* [Install LineageOS on oneplus3 | LineageOS Wiki](https://wiki.lineageos.org/devices/oneplus3/install)
* [LineageOS Downloads | oneplus3](https://download.lineageos.org/oneplus3)
* [Google apps | LineageOS Wiki](https://wiki.lineageos.org/gapps.html)  
  Google apps should be installed via recovery immediately after installing LineageOS.(before reboot)  
* [LineageOS Downloads | su](https://download.lineageos.org/extras)
* [TWRP Commandline Guide](https://twrp.me/faq/openrecoveryscript.html)
* [TWRP FAQ](https://twrp.me/FAQ/)

# Sun Jul 07 20:54:57 CST 2019

* [/GUIDE//MODDED FIRMWARE//9.0.4//OP3/T/The le… | OnePlus 3T](https://forum.xda-developers.com/oneplus-3t/how-to/guide-cope-9-0-3-5-0-8-firmware-barrier-t3941164)  
  Thank you very much.  

# Backup

```bash
cd ${BACKUPDIR}
adb pull /storage/emulated/0/DCIM/Camera .
adb pull /storage/emulated/0/Documents .
adb pull /storage/emulated/0/Pictures/Twitter .
adb pull /storage/emulated/0/pool .
```


# Install TWRP

**troubleshooting:**

[[RECOVERY][unified] Official TWRP touch recovery for OnePlus 3/3T - Post #373](https://forum.xda-developers.com/showpost.php?p=78332159&postcount=373)

```
adb reboot bootloader
fastboot oem unlock		if already unlocked, doesn't need do this.

fastboot flash recovery twrp.img
// then boot into recovery with power button

# or
# not every phone work this way, op3 is the one.
fastboot boot twrp-3.2.3-1-oneplus3.img
```

# Install ROM

## Packages

* [Flashable Firmware+Modem ZIPs | OnePlus 3](https://forum.xda-developers.com/oneplus-3/how-to/op3-flashable-firmware-modem-zips-t3816066)
* [The Open GApps Project](https://opengapps.org/)  

```
adb push {firmware,rom,gapps,addonsu}.zip /sdcard

adb shell
// now in twrp
// twrp wipe {system,data,cache,dalvik}
twrp wipe {cache,dalvik}
twrp install sdcard/{firmware,rom,gapps,addonsu}.zip

reboot
// Pray...
```

# Post Install

If Failed to reboot, Maybe it's because the Encryption doesn't work. Must
format `data` in recovery, then reflash.

	Hey man. First of all you need to update your twrp to the latest version for op3. That is 3.2.3.0. 
	and second of all. Google made changes in the vold with android pie,
	so encryption isn't working at all and is not supported at the moment. So you have to decrypt first. 
	Go to twrp recovery>advanced>format data and type 'yes'.
	Remember this will erase your internal storage so backup to a pc is must.
	After this,you can install the rom without any problems

# Checking info

if stock at `Checking info`, boot into recovery, then factory reset. That is,
the default wipe option.

# Basic APK

* [F-Droid](https://f-droid.org/en/)
* [Shelter](https://f-droid.org/en/packages/net.typeblog.shelter/)
* [WireGuard](https://f-droid.org/en/packages/com.wireguard.android/)
* [Releases · shadowsocks/shadowsocks-android · GitHub](https://github.com/shadowsocks/shadowsocks-android/releases)

```
// Doesn't need root shell
adb install FDroid.apk
adb install shelter.apk
adb install ss.apk
```

# TWRP COMMAND

```
wipe cache (eg: “twrp wipe cache”)
wipe dalvik
wipe data
backup (switches, name)
restore (source, switches)
mount (partition)
unmount (partition)
set variable = something
install <folder/filename>
and many more
```


