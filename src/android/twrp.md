# 命令行方式使用TWRP

# Links
http://www.pocketables.com/2014/10/using-twrps-new-adb-interface.html


# 在Redmi Note 4X上的操作

* 解锁BootLoader
* 刷入TWRP并启动进入TWRP环境  

	```
	fastboot flash recovery mido-twrp.img
	fastboot boot mido-twrp.img
	```
	* 这时候输入`adb devices`可以看到是在recovery里，上传la-mido.zip,并安装   
	```
	adb root 
	adb push la-mido.zip /sdcard
	adb shell
	# adb shell 之后就进入了recovery的shell环境,可以开始安装ROM或者其它应用  
	twrp wipe cache
	twrp wipe dalvik
	twrp wipe data
	twrp install /sdcard/la-mido.zip
	# 用同样的方法刷入su & gapps,刷这两个只需要wipe cache & dalvik
	reboot
	```


 
