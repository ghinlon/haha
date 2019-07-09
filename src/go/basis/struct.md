# struct
  

**Everything is a Number.**
  
* 給同样结构的结构体变量取一个名字,就创建了一种*结构体类型*.
  
## slice, map, struct 类型比较  
  
```go  
var i []int //这是int类型的slice类型
var s []string //这是string类型的slice类型
  
var ii map[int]int //这是索引是int类型,值是int类型的map类型
var is map[int]string //这是索引是int类型,值是string类型的map类型
var si map[string]int //这是索引是string类型,值是int类型的map类型
var ss map[string]string //这是索引是string类型,值是string类型的map类型
  
var a struct {    
    name string  
    age int  
    ssn string  
}   //这是一个结构体类型
  
type Person struct {    
    name string  
    age int  
    ssn string  
}   //这是一个结构体类型,取名Person,之后就可以创建Person类型的变量
  
var p Person //这是一个Person类型
var p []Person //这是一个Person类型的slice
  

```
  




