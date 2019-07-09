# [Package testing](://golang.org/pkg/testing/)

# Links

* [An Introduction to Testing in Go | TutorialEdge.net](https://tutorialedge.net/golang/intro-testing-in-go/)

# Overview

[Overview](https://golang.org/pkg/testing/#pkg-overview)

Package testing provides support for automated testing of Go packages. It is intended to be used in concert with the “go test” command, which automates execution of any function of the form

`func TestXxx(*testing.T)`

where Xxx does not start with a lowercase letter. The function name serves to identify the test routine.

Within these functions, use the Error, Fail or related methods to signal failure.

To write a new test suite, create a file whose name ends _test.go that contains the TestXxx functions as described here. Put the file in the same package as the one being tested. The file will be excluded from regular package builds but will be included when the “go test” command is run. 

简单来说，就是在你的包里建一个以`xxx_test.go`结尾的文件，里面有`TestXxx`为名的functions，在里向去判断`Xxx` functions 的输出跟你的期望么一样的，如果否一样，使`Error`, `Fail` 等methods去报错

示例：

`main.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func Calc(x int) (result int) {
	return x + 2
}
```

`main_test.go`:

```go
package main

import (
	"testing"
)

func TestCalc(t *testing.T) {

	foo := map[int]int{
		0:     2,
		-1:    1,
		1:     3,
		99999: 100001,
	}

	for k, v := range foo {

		if output := Calc(k); output != v {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", k, v, output)
		}

	}
}
```

再执行`go test`就瞅得到测试情况罢

```
$ go test
PASS
ok  	_/xx/practisego/01testing	0.004s
```
