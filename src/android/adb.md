# adb 

android debug bridge

# Links

* [Android ADB命令大全(通过ADB命令查看wifi密码、MAC地址、设备信息、操作文件、查看文件、日志信息、卸载、启动和安装APK等) | 张明云的博客](http://zmywly8866.github.io/2015/01/24/all-adb-command.html)

# INSTALL

```sh
brew cask install android-platform-tools
```

Run `adb root && adb shell` to get root permision

# set proxy

* [wi fi - How to set WIFI proxy via adb shell? - Android Enthusiasts Stack Exchange](https://android.stackexchange.com/questions/98287/how-to-set-wifi-proxy-via-adb-shell)
* [Undo setting proxy via Settings.Global in Android - Stack Overflow](https://stackoverflow.com/questions/31807559/undo-setting-proxy-via-settings-global-in-android)

```
adb shell settings put global http_proxy <address>:<port>

adb shell settings delete global http_proxy
adb shell settings delete global global_http_proxy_host
adb shell settings delete global global_http_proxy_port
```

# adb-sync

[GitHub - google/adb-sync](https://github.com/google/adb-sync)

works a little like `rsync`

```
adb-sync --help

adb-sync ~/Music /sdcard
adb-sync --delete ~/Music /sdcard

adb-sync --reverse /sdcard/Download/ ~/Downloads
```

# adb over net

[networking - How can I connect to Android with ADB over TCP? - Stack Overflow](https://stackoverflow.com/questions/2604727/how-can-i-connect-to-android-with-adb-over-tcp/3623727#3623727)

```
adb connect 192.168.0.101:5555
adb root
adb shell
// To tell the ADB daemon return to listening over USB
adb usb
```

# Command

<pre>
获取序列号：
adb get-serialno

查看连接计算机的设备：
adb devices

重启机器：
adb reboot

重启到bootloader，即刷机模式：
adb reboot bootloader

重启到recovery，即恢复模式：
 adb reboot recovery

查看log：
adb logcat

终止adb服务进程：
adb kill-server

重启adb服务进程：
adb start-server

获取机器MAC地址：
adb shell  cat /sys/class/net/wlan0/address

获取CPU序列号：
adb shell cat /proc/cpuinfo

安装APK：
adb install <apkfile> //比如：adb install ss.apk

保留数据和缓存文件，重新安装apk：
adb install -r <apkfile> //比如：adb install -r ss.apk

安装apk到sd卡：
adb install -s <apkfile> // 比如：adb install -s ss.apk

卸载APK：
adb uninstall <package> //比如：adb uninstall com.ss.search

卸载app但保留数据和缓存文件：
adb uninstall -k <package> //比如：adb uninstall -k com.ss.search

启动应用：
adb shell am start -n <package_name>/.<activity_class_name>

查看设备cpu和内存占用情况：
adb shell top

查看占用内存前6的app：
adb shell top -m 6

刷新一次内存信息，然后返回：
adb shell top -n 1

查询各进程内存使用情况：
adb shell procrank

杀死一个进程：
adb shell kill [pid]

查看进程列表：
adb shell ps

查看指定进程状态：
adb shell ps -x [PID]

查看后台services信息：
adb shell service list

查看当前内存占用：
adb shell cat /proc/meminfo

查看IO内存分区：
adb shell cat /proc/iomem

将system分区重新挂载为可读写分区：
adb remount

从本地复制文件到设备  
adb push <local> <remote>

从设备复制文件到本地：
adb pull <remote>  <local>

列出目录下的文件和文件夹，等同于dos中的dir命令：
adb shell ls

进入文件夹，等同于dos中的cd 命令：
adb shell cd <folder>

重命名文件：
adb shell rename path/oldfilename path/newfilename

删除system/avi.apk：
adb shell rm /system/avi.apk

删除文件夹及其下面所有文件：
adb shell rm -r <folder>

移动文件：
adb shell mv path/file newpath/file

设置文件权限：
adb shell chmod 777 /system/fonts/DroidSansFallback.ttf

新建文件夹：
adb shell mkdir path/foldelname

查看文件内容：
adb shell cat <file>

查看wifi密码：
adb shell cat /data/misc/wifi/*.conf

清除log缓存：
adb logcat -c

查看bug报告：
adb bugreport

获取设备名称：
adb shell cat /system/build.prop

查看ADB帮助：
adb help

跑monkey：
adb shell monkey -v -p your.package.name 500
</pre>


