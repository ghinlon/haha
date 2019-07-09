# Markdown对照表

<!-- ToC start -->

# Table of Contents

1. [Links](#links)
1. [标题](#标题)
1. [重点](#重点)
1. [列表](#列表)
1. [链接](#链接)
1. [图片](#图片)
1. [代码和语法高亮](#代码和语法高亮)
1. [表格](#表格)
1. [块引用](#块引用)
1. [嵌入HTML](#嵌入html)
1. [水平线规则](#水平线规则)
1. [换行](#换行)
1. [油管视频](#油管视频)
<!-- ToC end -->

# Links

* [为什么换行不是换行而要敲两下](https://meta.stackexchange.com/a/27111)
* [Markdown对照表](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet)


该文意为一份速查表和示例。更多完整信息，看[John Gruber的原始规范](http://daringfireball.net/projects/markdown/)和[Github味儿的Markdown信息页](http://github.github.com/github-flavored-markdown/)。
 
注意这还有一份[针对Markdown Here的对照表](https://github.com/adam-p/markdown-here/wiki/Markdown-Here-Cheatsheet),如果你需要的话。你也可以看看[更多Markdown工具](https://github.com/adam-p/markdown-here/wiki/Other-Markdown-Tools)。

# 标题

```markdown
# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

另外，对于一级标题和二级标题,还有下划线风格： 

另一种一级标题
======

另一种二级标题
------
```

# 一级标题
## 二级标题
### 三级标题
#### 四级标题
##### 五级标题
###### 六级标题

另外，对于一级标题和二级标题,还有下划线风格：
 
另一种一级标题
======

另一种二级标题
------

# 重点

```markdown
重点,也就是斜体, 用*星号*或_下划线_。 
很重点,也就是粗体,用**双星号**或__双下划线__。 
重点里的重点，用**双星号和_单下划线_**一起。
删除线用两个波浪线。~~划这里~~.
```

重点,也就是斜体, 用*星号*或_下划线_。

很重点,也就是粗体,用**双星号**或__双下划线__。

重点里的重点，用**双星号和_单下划线_**一起。

删除线用两个波浪线。~~划这里~~.

# 列表

(这个例子里, 开头和结尾的空格用点表示：.)  
```markdown
1. 第一个带序号的项
2. 另一个项
⋅⋅* 不带序号的子列表
1. 实际上数字是几不重要，只要它是数字。
⋅⋅1. 带序号的子列表
4. 另一项

⋅⋅⋅你应该已经和列表项一起段落缩进了. 注意上面的空行和开头的空格 (至少一个, 但我们在这里也用了三个是为了对齐原始Markdown).

⋅⋅⋅为了只换行不分段, 你需要在行尾使用两个空格.⋅⋅
⋅⋅⋅注意这行是新行,但在同一个段落里.⋅⋅
⋅⋅⋅(这和典型的GFM换行行为相反,GFM不需要行尾空格.)..


* 无序号的列表可以用星号
- 或者减号
+ 或者加号
```
1. 第一个带序号的项
2. 另一个项
  * 不带序号的子列表
1. 实际上数字是几不重要，只要它是数字。
  1. 带序号的子列表
4. 另一项

   你应该已经和列表项一起段落缩进了. 注意上面的空行和开头的空格 (至少一个, 但我们在这里也用了三个是为了对齐原始Markdown).

   为了只换行不分段, 你需要在行尾使用两个空格.  
   注意这行是新行,但在同一个段落里.  
   (这和典型的GFM换行行为相反,GFM不需要行尾空格.)  


* 无序号的列表可以用星号
- 或者减号
+ 或者加号

# 链接

有两种方式创建链接。  

```markdown
[我是一个内嵌链接](https://www.google.com)

[我是带标题的内嵌链接](https://www.google.com "谷歌的主页")

[我是一个参考链接][任意的不区分大小写的参考文本]

[我是一个到库文件的相对引用](README.md)
[你可以用数字定义参考链接][1] 
或者留空使用[链接文本自身].

URLs 和尖括号中的URLs会自动转换成链接. 
http://www.example.com 或 <http://www.example.com> 和有时 
example.com (比如，github上就不会).

一些可以之后作为参考链接用的文本。

[任意的不区分大小写的参考文本]: https://www.mozilla.org
[1]: http://slashdot.org
[链接文本自身]: http://www.reddit.com
```
[我是一个内嵌链接](https://www.google.com)

[我是带标题的内嵌链接](https://www.google.com "谷歌的主页")

[我是一个参考链接][任意的不区分大小写的参考文本]

[我是一个到库文件的相对引用](README.md)

[你可以用数字定义参考链接][1]

或者留空使用[链接文本自身].

URLs 和尖括号中的URLs会自动转换成链接. 
http://www.example.com 或 <http://www.example.com> 和有时 
example.com (比如，github上就不会).

一些可以之后作为参考链接用的文本。

[任意的不区分大小写的参考文本]: https://www.mozilla.org
[1]: http://slashdot.org
[链接文本自身]: http://www.reddit.com

# 图片
```markdown
这是我们的标志（悬停看标题文本）：

内嵌风格：
![alt text](images/icon48.png "标志标题文本一")

引用风格：
![alt text][标志]

[标志]: images/icon48.png "标志标题文本二"
```
这是我们的标志（悬停看标题文本）：

内嵌风格：
![alt text](images/icon48.png "标志标题文本一")

引用风格：
![alt text][标志]

[标志]: images/icon48.png "标志标题文本二"

# 代码和语法高亮

代码块是Markdown规范的一部分，但语法高亮不是。但是，许多渲染器--像Github的和Markdown Here -- 都支持语法高亮。支持哪些语言和那些语言的名字要怎么写渲染器和渲染器之间就会不同。Markdown Here支持数十种语言的语法高亮（还支持不是真正的语言，像diffs和HTTP头）;查看完整列表和怎么写语言名字，看[highlight.js演示页面](http://softwaremaniacs.org/media/soft/highlight/test.html)。  
```markdown
行内 `代码` 用 `反斜杠包围` 它.
``` 

行内 `代码` 用 `反斜杠包围` 它.  

代码块用三个反斜杠```的行隔开，或用4个空格缩进。我建议只用隔开的代码块--他们更简单和只有他们支持语法高亮。  

<pre>
```javascript
var s = "JavaScript 语法高亮";
alert(s);
```
 
```python
s = "Python 语法高亮"
print s
```
 
```
未指定语言, 所以没有语法高亮。  
但是我们扔一个&ltb&gt标记&lt/b&gt.
```
</pre>

```javascript
var s = "JavaScript 语法高亮";
alert(s);
```
 
```python
s = "Python 语法高亮"
print s
```
 
```
未指定语言, 所以没有语法高亮。  
但是我们扔一个<b>标记</b>.
```

# 表格

表格不是核心Markdown规范的一部分,但它们是GFM的一部分和Markdown Here 支持它们. 它们是一种給你的邮件添加表格的简单的方式.--否则完成这项任务就必须从其它应用程序复制粘贴了.

<pre>
冒号可以用来对齐列.

| 表格        | 很           | 牛逼  |
| ------------|:------------:| -----:|
| 第三列是    | 右对齐       | $1600 |
| 第二列是    | 居中         |   $12 |
| 斑马条纹    | 是整洁的     |    $1 |

每个标题单元格必须至少有3个划线.
外围的管道(|)是可选的, 而且你不用一定让原始Markdown排列的整齐.
你也可以使用内嵌的Markdown.

Markdown | 不 | 整齐的
--- | --- | ---
*仍然* | `渲染的` | **很漂亮**
1 | 2 | 3
</pre>

冒号可以用来对齐列.

| 表格        | 很           | 牛逼  |
| ------------|:------------:| -----:|
| 第三列是    | 右对齐       | $1600 |
| 第二列是    | 居中         |   $12 |
| 斑马条纹    | 是整洁的     |    $1 |

每个标题单元格必须至少有3个划线.
外围的管道(|)是可选的, 而且你不用一定让原始Markdown排列的整齐.
你也可以使用内嵌的Markdown.

Markdown | 不 | 整齐的
--- | --- | ---
*仍然* | `渲染的` | **很漂亮**
1 | 2 | 3

# 块引用

<pre>
> 块引用非常方便在邮件中模拟回复文本.
> 这一行仍然是同一个引用的部分.

引用打断.

> 这是一行非常长的行就算超过屏幕了也能被正常引用.  噢男娃子我们接着写确保它足够长让每个人都真的可以超过屏幕. 噢, 你可以*放* **Markdown**在引用里. 
</pre>

> 块引用非常方便在邮件中模拟回复文本.  
> 这一行仍然是同一个引用的部分.

引用打断.

> 这是一行非常长的行就算超过屏幕了也能被正常引用.  噢男娃子我们接着写确保它足够长让每个人都真的可以超过屏幕. 噢, 你可以*放* **Markdown**在引用里. 

# 嵌入HTML

你也可以用纯HTML在你的Markdown里, 而且它大多数情况会工作的非常好.
<pre>
&ltdl>
  &ltdt>定义列表&lt/dt>
  &ltdd>是人们有时用的一些东西&lt/dd>

  &ltdt>Markdown在HTML里&lt/dt>
  &ltdd>其 *未* 工作的 **很** 好. 用 HTML &ltem>标志&lt/em>.&lt/dd>
&lt/dl>
</pre>

<dl>
  <dt>定义列表</dt>
  <dd>是人们有时用的一些东西</dd>

  <dt>Markdown在HTML里</dt>
  <dd>其 *未* 工作的 **很** 好. 用 HTML <em>标志</em>.</dd>
</dl>

# 水平线规则

<pre>
三个或更多...

---

连字符

***

星号

___

下划线

</pre>

---

连字符

***

星号

___

下划线

# 换行

对于换行是怎么工作的我的建议是试验和发现 -- 按 <Enter> 一次 (即, 插入一个换行符), 然后按两次(即, 插入两个换行符), 看发生了什么. 你会很快学到得到你想要的. "Markdown Toggle" 是你的朋友.

这里有一些事情要尝试:

<pre>
我们用这行开始.

这行与上面那行用两个换行符隔开, 所以它会是 *单独的段落*.

这行也是一个单独的段落, 但...
这行只用一个换行符隔开, 所以它是单独的一行在 *同一个段落*.
</pre>

我们用这行开始.

这行与上面那行用两个换行符隔开, 所以它会是 *单独的段落*.  

这行也是一个单独的段落, 但...  
这行只用一个换行符隔开, 所以它是单独的一行在 *同一个段落*.  

# 油管视频


他们不能被直接添加但你可以像这样添加一张带有到视频链接的图片:

<pre>
&lta href="http://www.youtube.com/watch?feature=player_embedded&v=YOUTUBE_VIDEO_ID_HERE
" target="_blank">&ltimg src="http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg" 
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>
</pre>
或者, 在純Markdown里, 但失去图片尺寸和边框:

<pre>
[![IMAGE ALT TEXT HERE](http://img.youtube.com/vi/YOUTUBE_VIDEO_ID_HERE/0.jpg)](http://www.youtube.com/watch?v=YOUTUBE_VIDEO_ID_HERE)
</pre> 

在git commit中引用#bugID引起的错误将其链接到传票。 例如＃1。


许可証:[CC-BY](http://creativecommons.org/licenses/by/3.0/)

