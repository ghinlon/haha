# [package proto](https://godoc.org/github.com/golang/protobuf/proto)

import "github.com/golang/protobuf/proto"

# Links

* [Getting Started with Protocol Buffers in Go - Tutorial - YouTube](https://www.youtube.com/watch?v=NoDRq6Twkts&t=14s)
* [GitHub - golang/protobuf: Go support for Google's protocol buffers](https://github.com/golang/protobuf)
* [protobuf/README.md at master · protocolbuffers/protobuf · GitHub](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)
* [Protocol Buffer Basics: Go  |  Protocol Buffers  |  Google Developers](https://developers.google.com/protocol-buffers/docs/gotutorial)
* [protobuf/examples at master · protocolbuffers/protobuf](https://github.com/protocolbuffers/protobuf/tree/master/examples)
* [Language Guide (proto3)  |  Protocol Buffers  |  Google Developers](https://developers.google.com/protocol-buffers/docs/proto3)
* [Go Generated Code  |  Protocol Buffers  |  Google Developers](https://developers.google.com/protocol-buffers/docs/reference/go-generated)
* [Encoding  |  Protocol Buffers  |  Google Developers](https://developers.google.com/protocol-buffers/docs/encoding)

# Installation

直接在[releases](https://github.com/protocolbuffers/protobuf/releases)下载编译好的程序

然后安装`protoc-gen-go`插件, `proto`库

需要配置`GOPATH`, 以及将`${GOPATH}/bin`加进`PATH`

```
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/proto
```

# Getting Started

定义

```proto
$ cat person.proto 
syntax = "proto3";
package person;

message Person {
	
	string name = 1;
	int32 age =2;
}
```

生成代码

`protoc --go_out=. person.proto`

执行这条命令后，会生成一个文件：`person.pb.go`

使用示例

```go
package main

import (
	"fmt"
	"log"

	// 这里是创建一个目录，把生成个person.pb.go放里面
	person "./protobuf"
	proto "github.com/golang/protobuf/proto"
)

func main() {

	// 这里不加上Name: Age: 就要报错
	p := &person.Person{Name: "zhangsan", Age: 99}

	data, err := proto.Marshal(p)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data, string(data))
	// [10 8 122 104 97 110 103 115 97 110 16 99] 
	// zhangsanc

	newp := new(person.Person)

	err = proto.Unmarshal(data, newp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(newp.GetName(), newp.GetAge())
	// zhangsan 99

}
```

# 一对编解组函数

Marshal takes a protocol buffer message and encodes it into the wire format, returning the data. This is the main entry point. 

```go
func Marshal(pb Message) ([]byte, error)
func Unmarshal(buf []byte, pb Message) error
```
# Message接口

```go
type Message interface {
    Reset()
    String() string
    ProtoMessage()
}
```

