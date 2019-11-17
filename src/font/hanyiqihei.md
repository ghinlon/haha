# 汉仪旗黑存在的问题
  
## 问题

1. 旗黑只支持zh-cn,如果lang=zh-tw的情况,默认字体就匹配不到该字体.我的办法是在zh-tw后面加上zh-cn,并且弱绑定.  
  ```xml
        <match target="pattern">
                <test qual="any" name="lang"><string>zh-tw</string></test>
                <edit name="lang" mode="assign" binding="weak"><string>zh-tw</string></edit>
                <edit name="lang" mode="assign" binding="weak"><string>zh-cn</string></edit>
        </match>
```

2. 汉仪旗黑所有字重的字体,在用HYQiHei匹配的时候,得分都是一样的.所以就得到了下面这样的结果,因为它第一个匹配.    
   采用的办法是为每个字号配置指定字号的字体.

```
$ fc-match  :lang=zh-cn
HYQiHei-105JF.ttf: "HYQiHei" "105JF"

$ fc-list :lang=zh-cn | grep -i hyqihei[^xy]
/usr/share/fonts/hanyi/HYQiHei-80S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-80S,汉仪旗黑\-80S:style=80S,ExtraBold
/usr/share/fonts/hanyi/HYQiHei-25S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-25S,汉仪旗黑\-25S:style=25S,Hairline
/usr/share/fonts/hanyi/HYQiHei-75S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-75S,汉仪旗黑\-75S:style=75S,Bold
/usr/share/fonts/hanyi/HYQiHei-85S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-85S,汉仪旗黑\-85S:style=85S,Heavy
/usr/share/fonts/hanyi/HYQiHei-105JF.ttf: HYQiHei,汉仪旗黑,HYQiHei\-105JF,汉仪旗黑\-105简繁:style=105JF,105简繁,UltraBlack
/usr/share/fonts/hanyi/HYQiHei-55S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-55S,汉仪旗黑\-55S:style=55S,Book
/usr/share/fonts/hanyi/HYQiHei-70S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-70S,汉仪旗黑\-70S:style=70S,DemiBold
/usr/share/fonts/hanyi/HYQiHei-60S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-60S,汉仪旗黑\-60S:style=60S,Regular
/usr/share/fonts/hanyi/HYQiHei-35S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-35S,汉仪旗黑\-35S:style=35S,Thin
/usr/share/fonts/hanyi/HYQiHei-95S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-95S,汉仪旗黑\-95S:style=95S,ExtraBlack
/usr/share/fonts/hanyi/HYQiHei-30S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-30S,汉仪旗黑\-30S:style=30S,ExtraThin
/usr/share/fonts/hanyi/HYQiHei-65S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-65S,汉仪旗黑\-65S:style=65S,Medium
/usr/share/fonts/hanyi/HYQiHei-50S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-50S,汉仪旗黑\-50S:style=50S,Light
/usr/share/fonts/hanyi/HYQiHei-45S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-45S,汉仪旗黑\-45S:style=45S,ExtraLight
/usr/share/fonts/hanyi/HYQiHei-40S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-40S,汉仪旗黑\-40S:style=40S,UltraLight
```
## 字体的weigh属性

extralight - 40
Thin - 45
light - 50
Demilight - 55
regular - 80
medium - 100
semibold - 180
bold - 200
black - 210


字号的对应关系
https://stackoverflow.com/questions/5912528/font-size-translating-to-actual-point-size

摘录:
```
It may vary by browser slightly but for the most part this should work:

Large is 18 px which is around 13.5 pt

Larger is 19 px which is around 14 pt

Medium is 16 px which is around 12 pt

Small is 13 px which is around 10 pt

Smaller is 13 px which is around 10 pt

X-large is 24 px which is around 18 pt

X-small is 10 px which is around 7.5 pt

XX-large is 32 px which is around 24 pt

XX-small is 9 px which is around 7 pt

This is based off of seeing the computed font-size style in pixels and converting from this chart. This link might also be helpful.
```
http://style.cleverchimp.com/font_size_intervals/altintervals.html#st


配置文件中,字体名的-不需要用\转义.
```xml
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>90</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>60</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-105JF</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>60</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>40</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-95S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>40</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>30</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-85S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>30</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>24</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-80S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>24</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>18</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-75S</string></edit>
</match>



<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>18</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>16</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-70S</string></edit>
</match>

<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>16</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>14</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-65S</string></edit>
</match>


<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>14</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>13</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-60S</string></edit>
</match>

<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>13</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>12</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-55S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>12</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>11</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-50S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>11</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>10</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-45S</string></edit>
</match>

<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>10</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>9</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-40S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>9</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>8</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-35S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>8</double></test>
        <test qual="any" name="pixelsize" compare="more"><double>7</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-30S</string></edit>
</match>
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <test qual="any" name="pixelsize" compare="less_eq"><double>7</double></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-25S</string></edit>
</match>

<!-- HYQiHei fallback-->
<match target="pattern">
        <test qual="any" name="family"><string>HYQiHei</string></test>
        <edit name="family" mode="assign" binding="same"><string>HYQiHei-65S</string></edit>
</match>
```

很遗憾,chromium传出来的模式,不含pixelsize.所以永远只能fallback到HYQiHei-65S.还不知道怎么办?

在Telegram上了解到,chromium使用skia作为底层渲染引擎,它和fontconfig目前没有很好对接.像語言,字重可能都没有传給fontconfig的.

所以,就这样吧!

唉,遗憾.

测试了firefox,也是存在问题.

终于找到原因了,发现汉仪好几个字体文件字重没有设置正确
```
$ fc-list HYQiHei:regular
/usr/share/fonts/hanyi/HYQiHei-80S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-80S,汉仪旗黑\-80S:style=80S,ExtraBold
/usr/share/fonts/hanyi/HYQiHei-25S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-25S,汉仪旗黑\-25S:style=25S,Hairline
/usr/share/fonts/hanyi/HYQiHei-105JF.ttf: HYQiHei,汉仪旗黑,HYQiHei\-105JF,汉仪旗黑\-105简繁:style=105JF,105简繁,UltraBlack
/usr/share/fonts/hanyi/HYQiHei-55S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-55S,汉仪旗黑\-55S:style=55S,Book
/usr/share/fonts/hanyi/HYQiHei-60S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-60S,汉仪旗黑\-60S:style=60S,Regular
/usr/share/fonts/hanyi/HYQiHei-95S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-95S,汉仪旗黑\-95S:style=95S,ExtraBlack
/usr/share/fonts/hanyi/HYQiHei-30S.ttf: HYQiHei,汉仪旗黑,HYQiHei\-30S,汉仪旗黑\-30S:style=30S,ExtraThin
```

这些都被当然regular字重去匹配了.使用fontforge修改为相应字重后,正常.

```
HYQiHei-25S.ttf
	style: "25S"(s) "Hairline"(s) 10
HYQiHei-30S.ttf
	style: "30S"(s) "ExtraThin"(s) 75
HYQiHei-35S.ttf
	style: "35S"(s) "Thin"(s) 100 不用改
HYQiHei-40S.ttf
	style: "40S"(s) "UltraLight"(s) 200 不用改
HYQiHei-45S.ttf
	style: "45S"(s) "ExtraLight"(s) 200 不用改
HYQiHei-50S.ttf
	style: "50S"(s) "Light"(s)  300 不用改
HYQiHei-55S.ttf
	style: "55S"(s) "Book"(s)  375 
HYQiHei-60S.ttf
	style: "60S"(s) "Regular"(s) 400 不用改
HYQiHei-65S.ttf
	style: "65S"(s) "Medium"(s) 500 不用改
HYQiHei-70S.ttf
	style: "70S"(s) "DemiBold"(s) 600 不用改
HYQiHei-75S.ttf
	style: "75S"(s) "Bold"(s) 700 不用改
HYQiHei-80S.ttf
	style: "80S"(s) "ExtraBold"(s) 800
HYQiHei-85S.ttf
	style: "85S"(s) "Heavy"(s) 900 不用改
HYQiHei-95S.ttf
	style: "95S"(s) "ExtraBlack"(s) 1000
HYQiHei-105JF.ttf
	style: "105JF"(s) "105简繁"(s) "UltraBlack"(s)  to:1100

```

## 参考资料

https://stackoverflow.com/questions/17508/how-to-modify-the-style-property-of-a-font-on-windows  


