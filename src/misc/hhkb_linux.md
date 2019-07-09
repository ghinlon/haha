# remap keyboard to hhkb use xmodmap

使用的配置文件[hhkb.Xmodmap](https://github.com/iofxl/hhkb/blob/master/hhkb.Xmodmap)  

使用方法:  
执行这两条命令:  
```
# always recovery the default keymap first
setxkbmap
xmodmap hhkb.Xmodmap
```

**xmodmap配置文件格式**  

```
              Base Shift Alt_Gr Shift+Alt_Gr Compose(?) Shift+Compose(?)
keycode ?? =    ?     ?      ?        ?          ?              ?
```

#### Links
https://deskthority.net/wiki/Xmodmap


#### 关于xmodmap

**常用命令**
```
# 打印modifier
xmodmap -pm 
# 以可恢复的格式打印配置
xmodmap -pke
```

每次执行`xmodmap <config_file>`前都要先执行`setxkbmap`

