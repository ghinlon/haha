# interfaces

* WHW

#### Links


你知道接口是接口，但你还是很难去理解接口。

https://golang.org/pkg/io/    
https://golang.org/pkg/io/ioutil/    
https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter01/01.1.md 很一般的介绍IO包的文章 例子比较差 但仍然值得学习 中文文档整体质量太差    
https://www.cnblogs.com/golove/p/3276678.html  
https://gobyexample.com/interfaces 这个页面帮助理解什么是接口  
http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go  
https://golangtutorials.blogspot.com/2011/06/interfaces-in-go.html?m=1 解释类型与接口的关系  
https://research.swtch.com/interfaces 解释接口在内存中的本质  
https://blog.golang.org/laws-of-reflection 可以进一步理解接口的本质  
https://golang.org/doc/effective_go.html#interfaces  
http://golangtutorials.blogspot.com/2011/05/table-of-contents.html  
http://www.laktek.com/2012/02/13/learning-go-interfaces-reflections/  


接口是函数的抽象或集合 接口可以作为参数传递给函数

接口也是一种类型 一类这样的东西或者概念 就像情感 函数 文学 这些东西一样

通常的类型是面向属性的类型 接口可以认为是面向行为的类型 面向函数的类型

例子

求长方形 圆形的面积
四种动物的叫 狗 猫 牛 Java程序员

形状是 正方形 长方形 圆形 各种形状的抽象 那它就可以成为接口？抽象的概念就可以是接口？

列车时刻查询也是接口 用英语 法语 手语查询都可以查询到结果

某种接口包含某种函数 某种类型也支持这种函数 那么 这种类型就实现了这个接口 可以调用接口以这个类型的实例为参数

普通类型传值给接口类型可以看成是类型转换
为什么Read这些接口不是这样的用法 因为这些接口类型的函数是带参数的

像ByteReader等接口实现了又有什么意义 只是产生了有这个接口这样的概念 并没有任何价值 因为IO包并没有提供相关的操作函数，类似io.Copy 如果也有ByteCopy这样的函数在 那么所谓的实现了这个接口就有价值了 接口必须得有相应对接才有价值 普通类型也是一样的道理 也就是得有地方去用

但是实现了这个接口，就可以等着被用了。用不用得到另说。

[Effective Go#interfaces](https://golang.org/doc/effective_go.html#interfaces)中的定义：  

>Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.
