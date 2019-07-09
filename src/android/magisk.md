# Magisk and EdXposed(deprecated)


# Links

* [Magisk v19.2 - Root & Universal Systemless Interface](https://forum.xda-developers.com/apps/magisk/official-magisk-v7-universal-systemless-t3473445)
* [How to Install Magisk](https://www.xda-developers.com/how-to-install-magisk/)
* [9.0 pie - What are the alternatives for Xposed framework on Pie? Preferably without triggering Safety Net? - Android Enthusiasts Stack Exchange](https://android.stackexchange.com/questions/207633/what-are-the-alternatives-for-xposed-framework-on-pie-preferably-without-trigge)

# Install

[Releases · topjohnwu/Magisk · GitHub](https://github.com/topjohnwu/Magisk/releases/)

* Flash `Magisk.zip` with twrp
* Install `MagiskManager.apk`  
  This apk is used for install modules.


# EdXposed

## Install riru

Install it in MagiskManager

[GitHub - RikkaApps/Riru: Inject zygote process by replace libmemtrack](https://github.com/RikkaApps/Riru)


* Install `riru-core-v19.1.zip` in Magisk: 
* Install `riru-core-v19.1-magisk-v17.zip` in Magisk 

## Install EdXposed

* [GitHub - ElderDrivers/EdXposed: Elder driver Xposed Framework.](https://github.com/ElderDrivers/EdXposed)
* [Releases · ElderDrivers/EdXposed · GitHub](https://github.com/ElderDrivers/EdXposed/releases)  
  `SandHook` vs `YAHFA` just choose one.

* Install `magisk-EdXposed-SandHook-v0.4.1.2_beta-release.zip` in Magisk
* Install Xposed Installer APK  
  For v0.2.9.8 and later: EdXposed Installer and EdXposed Manager.  
  * [Releases · solohsu/XposedInstaller · GitHub](https://github.com/solohsu/XposedInstaller/releases)  
  * [Releases · ElderDrivers/EdXposedManager · GitHub](https://github.com/ElderDrivers/EdXposedManager/releases)  
  **EdXposedManager doen's work. so don't install it.**  
  ```
  adb install EdXposedInstaller_v2.2.4-release.apk 
  ```
* reboot

## Install Xposed modules

[List of Xposed Modules For Android Pie Working With Ed Xposed F](https://forum.xda-developers.com/xposed/list-xposed-modules-android-pie-ed-t3892768)

* [XPrivacyLua | Xposed Module Repository](https://repo.xposed.info/module/eu.faircode.xlua)

```
// after install active it then reboot
adb install eu.faircode.xlua_v125_25c714.apk 
```



