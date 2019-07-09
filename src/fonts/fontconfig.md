# fontconfig

## 字体安装
  
全局目录`/usr/share/fonts/`用户目录`~/.local/share/fonts/`  
字体文件的权限至少要对所有用户可读.  
使用命令`fc-cache`更新字体缓存.

## 疑问

1. 如果默认字体不支持中文，字体怎么匹配？  
   就会匹配另一个支持中文的字体，如果匹配不到，就显示乱码了。
2. 如果默认字体支持中文，那么自然也支持英文，怎么使英文用更好看的英文字体匹配？  
   让英文字体的模式在前面。

## 字体调试
  
在浏览器地址栏里输入这个测试:

```
data:text/html,<meta charset="utf8"><p style="font-family: monospace;">☺</p>
```
  
```
<match target="目标">
    <match>必须首先包含一系列<test>组成的列表(可以为空)，然后再包含一系列<edit>组成的列表(可以为空)，
    也就是<test>列表必须位于<edit>列表之前。注意：虽然两个列表都可以为空，但是不可以同时为空。
    如果"目标"满足<test>列表的所有测试条件，那么将被按照<edit>列表中的指令序列进行修改。
    "目标"的默认值是"pattern"，表示此<match>单元针对的是用于匹配的字体模板(第一次修改)。
    如果"目标"的值是"font"，那么就表示此<match>单元针对的是已被选定的字体(第二次修改)。
    如果"目标"的值是"scan"，那么就表示此<match>单元针对的是扫描字体以创建内部配置数据的初始化阶段(FcConfigParse)。

为了帮助诊断字体和应用的各种问题，fontconfig 内置了许多调试功能。这些调试功能可以通过 FC_DEBUG 环境变量进行控制。
这个环境变量是个整数值，它的每个位都对应着一类调试信息。具体如下：

  Name         Value    Meaning
  ---------------------------------------------------------
  MATCH            1    显示字体匹配的简要信息
  MATCHV           2    显示字体匹配的扩展信息
  EDIT             4    监视 match/test/edit 的执行
  FONTSET          8    在启动时跟踪字体信息的加载
  CACHE           16    显示字体缓存变化的简要信息
  CACHEV          32    显示字体缓存变化的扩展信息
  PARSE           64    (已废弃)
  SCAN           128    显示扫描字体文件并创建缓存的过程
  SCANV          256    显示扫描字体文件的详细信息
  MEMORY         512    监视 fontconfig 的内存使用量
  CONFIG        1024    监视加载了哪些配置文件
  LANGSET       2048    显示用来创建 lang 的字符集
  OBJTYPES      4096    显示值类型检查失败的消息

将你期望看到的调试信息类别所对应的 Value 值相加，然后赋给 FC_DEBUG 环境变量，再运行应用程序就可以在 stdout 上看到调试信息了。
[例子]假如你想查看'Courier,mono'的匹配过程，可以在命令行上运行： FC_DEBUG=5 fc-match -s 'Courier,mono'
```

* FC_DEBUG=3   
  可以用来观察每个待选字体的得分情况.这在分析为什么你想用A字体却得到了B字体的时候很有用.  
  分数越接近0,越优.
* FC_DEBUG=4   
  用来观察字体模式的匹配过程.


测试当前默认字体  
```
# fc-match
LiberationMono-Regular.ttf: "Liberation Mono" "Regular"
```
测试当前默认等宽字体
```  
$ fc-match monospace
DejaVuSansMono.ttf: "DejaVu Sans Mono" "Book"
```
随意的一个例子
```
# fc-match 'Noto Sans CJK TC','Noto Sans CJK SC':lang=zh-cn:weight=bold
NotoSansCJK-Bold.ttc: "Noto Sans CJK TC" "Bold"
```
  
## 字体配置

如果默认字体是英文字体，那么中文会用其它字体显示。如果默认字体是中文，那么就会直接用这个字体显示英文，因为字符集包含英文。  

我觉得字体配置的基本要求是：  

* 应用程序要求使用字体A，就要提供字体A，如果没有该字体，就提供一种相近字体。
* 应用程序对字体没要求，就提供默认字体。

那些覆盖应用的字体要求的配置文件，在我看来简直是胡闹。  

因为中文字体会包含英文字符集，为了让英文可以用更好看的英文字体显示，可以在在中文字体前面插入英文字体模式。  

/etc/fonts/conf.d/里面默认的配置，有几个使用了**prepend**，我觉得是很有毒的。其中一个例子，这个匹配结果是由60-latin.conf，65-nonlatin.conf引起.  
所以我把它们删了.基本上默认配置文件中使用了**prepend**的我都給删了.不删不行.  

```
FcConfigSubstitute test pattern any family Equal(ignore blanks) "sans-serif"
Substitute Edit family Prepend "Nachlieli" Comma "Lucida Sans Unicode" Comma "Yudit Unicode" Comma "Kerkis" Comma "ArmNet Helvetica" Comma "Artsounk" Comma "BPG UTF8 M" Comma "Waree" Comma "Loma" Comma "Garuda" Comma "Umpush" Comma "Saysettha Unicode" Comma "JG Lao Old Arial" Comma "GF Zemen Unicode" Comma "Pigiarniq" Comma "B Davat" Comma "B Compset" Comma "Kacst-Qr" Comma "Urdu Nastaliq Unicode" Comma "Raghindi" Comma "Mukti Narrow" Comma "malayalam" Comma "Sampige" Comma "padmaa" Comma "Hapax Berbère" Comma "MS Gothic" Comma "UmePlus P Gothic" Comma "Microsoft YaHei" Comma "Microsoft JhengHei" Comma "WenQuanYi Zen Hei" Comma "WenQuanYi Bitmap Song" Comma "AR PL ShanHeiSun Uni" Comma "AR PL New Sung" Comma "MgOpen Modata" Comma "VL Gothic" Comma "IPAMonaGothic" Comma "IPAGothic" Comma "Sazanami Gothic" Comma "Kochi Gothic" Comma "AR PL KaitiM GB" Comma "AR PL KaitiM Big5" Comma "AR PL ShanHeiSun Uni" Comma "AR PL SungtiL GB" Comma "AR PL Mingti2L Big5" Comma "ＭＳ ゴシック" Comma "ZYSong18030" Comma "TSCu_Paranar" Comma "NanumGothic" Comma "UnDotum" Comma "Baekmuk Dotum" Comma "Baekmuk Gulim" Comma "KacstQura" Comma "Lohit Bengali" Comma "Lohit Gujarati" Comma "Lohit Hindi" Comma "Lohit Marathi" Comma "Lohit Maithili" Comma "Lohit Kashmiri" Comma "Lohit Konkani" Comma "Lohit Nepali" Comma "Lohit Sindhi" Comma "Lohit Punjabi" Comma "Lohit Tamil" Comma "Meera" Comma "Lohit Malayalam" Comma "Lohit Kannada" Comma "Lohit Telugu" Comma "Lohit Oriya" Comma "LKLUG"

Prepend list before  "Bitstream Vera Sans"(w) "DejaVu Sans"(w) "Verdana"(w) "Arial"(w) "Albany AMT"(w) "Luxi Sans"(w) "Nimbus Sans L"(w) "Nimbus Sans"(w) "Helvetica"(w) "Lucida Sans Unicode"(w) "BPG Glaho International"(w) "Tahoma"(w) [marker] **"sans-serif"(s)** "Roya"(w) "Koodak"(w) "Terafik"(w)
Prepend list after  "Bitstream Vera Sans"(w) "DejaVu Sans"(w) "Verdana"(w) "Arial"(w) "Albany AMT"(w) "Luxi Sans"(w) "Nimbus Sans L"(w) "Nimbus Sans"(w) "Helvetica"(w) "Lucida Sans Unicode"(w) "BPG Glaho International"(w) "Tahoma"(w) **"Nachlieli"(w)** "Lucida Sans Unicode"(w) "Yudit Unicode"(w) "Kerkis"(w) "ArmNet Helvetica"(w) "Artsounk"(w) "BPG UTF8 M"(w) "Waree"(w) "Loma"(w) "Garuda"(w) "Umpush"(w) "Saysettha Unicode"(w) "JG Lao Old Arial"(w) "GF Zemen Unicode"(w) "Pigiarniq"(w) "B Davat"(w) "B Compset"(w) "Kacst-Qr"(w) "Urdu Nastaliq Unicode"(w) "Raghindi"(w) "Mukti Narrow"(w) "malayalam"(w) "Sampige"(w) "padmaa"(w) "Hapax Berbère"(w) "MS Gothic"(w) "UmePlus P Gothic"(w) "Microsoft YaHei"(w) "Microsoft JhengHei"(w) "WenQuanYi Zen Hei"(w) "WenQuanYi Bitmap Song"(w) "AR PL ShanHeiSun Uni"(w) "AR PL New Sung"(w) "MgOpen Modata"(w) "VL Gothic"(w) "IPAMonaGothic"(w) "IPAGothic"(w) "Sazanami Gothic"(w) "Kochi Gothic"(w) "AR PL KaitiM GB"(w) "AR PL KaitiM Big5"(w) "AR PL ShanHeiSun Uni"(w) "AR PL SungtiL GB"(w) "AR PL Mingti2L Big5"(w) "ＭＳ ゴシック"(w) "ZYSong18030"(w) "TSCu_Paranar"(w) "NanumGothic"(w) "UnDotum"(w) "Baekmuk Dotum"(w) "Baekmuk Gulim"(w) "KacstQura"(w) "Lohit Bengali"(w) "Lohit Gujarati"(w) "Lohit Hindi"(w) "Lohit Marathi"(w) "Lohit Maithili"(w) "Lohit Kashmiri"(w) "Lohit Konkani"(w) "Lohit Nepali"(w) "Lohit Sindhi"(w) "Lohit Punjabi"(w) "Lohit Tamil"(w) "Meera"(w) "Lohit Malayalam"(w) "Lohit Kannada"(w) "Lohit Telugu"(w) "Lohit Oriya"(w) "LKLUG"(w) **"sans-serif"(s)** "Roya"(w) "Koodak"(w) "Terafik"(w)
```

**注：**结果可以看出来，prepend的作用，不是前置在整个模式最前面的，而只前置在参与匹配值的前面。  



## 字体配置的逻辑

```
 Files begining with:   Contain:
 
 00 through 09          Font directories
 10 through 19          system rendering defaults (AA, etc)
 20 through 29          font rendering options
 30 through 39          family substitution
 40 through 49          generic identification, map family->generic
 50 through 59          alternate config file loading
 60 through 69          generic aliases, map generic->family
 70 through 79          select font (adjust which fonts are available)
 80 through 89          match target="scan" (modify scanned patterns)
 90 through 99          font synthesis
```

正确的玩法就是按照编号的这个思路来.

* 30-39:对同类型的字体做一个字体替换.  
  我们希望这些字体中的每一个字体都可以回滚到它们中的任意一个,
  但是相似的字体优先.我们用三步来实现:

  1. 别名每个特殊到它的一般字族    
     例. Liberation Sans 至 Arial
```xml
        <alias binding="same">
          <family>Liberation Sans</family>
          <default>
          <family>Arial</family>
          </default>
        </alias>
```
  1. 弱别名每个一般字族到其它一般字族  
     例. Arial to Helvetica
```xml
        <alias>
          <family>Arial</family>
          <default>
          <family>Helvetica</family>
          </default>
        </alias>
```

  1. 别名每个一般到它的特殊  
     例. Arial to Liberation Sans, Arimo, Albany, and Albany AMT
```xml
        <alias binding="same">
          <family>Arial</family>
          <accept>
            <family>Arimo</family>
            <family>Liberation Sans</family>
            <family>Albany</family>
            <family>Albany AMT</family>
          </accept>
        </alias>
```

  我自己准备添加一个:32-metric-aliases.conf(純想法,懒得弄)
  1. 对中文字体做相同的配置.列出一张字体表
  1. 对于不想用的字体,我会用alias binding="same" prefer 想用的字体覆盖.
   
* 40-49:映射字体类型,就是定义字体的类别:monospace,sans-serif,serif.
```xml
        <alias>
                <family>ＭＳ 明朝</family>
                <default><family>serif</family></default>
        </alias>
```

## 配置默认字体

[我的配置文件](https://raw.githubusercontent.com/iofxl/config/master/etc/fonts/local.conf)

## 参考资料

https://wiki.archlinux.org/index.php/fonts  
https://wiki.archlinux.org/index.php/Font_configuration  
http://www.jinbuguo.com/gui/fonts.conf.html  
https://eev.ee/blog/2015/05/20/i-stared-into-the-fontconfig-and-the-fontconfig-stared-back-at-me/  
https://stackoverflow.com/questions/17508/how-to-modify-the-style-property-of-a-font-on-windows  









