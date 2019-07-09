# init functions in Go

# Links

* [init functions in Go – golangspec – Medium](https://medium.com/golangspec/init-functions-in-go-eac191b3860a)
* [Initialization dependencies in Go – golangspec – Medium](https://medium.com/golangspec/initialization-dependencies-in-go-51ae7b53f24c)
* [The init function ¶](https://golang.org/doc/effective_go.html#init)
* [go - When is the init() function run? - Stack Overflow](https://stackoverflow.com/questions/24790175/when-is-the-init-function-run/49831018#49831018)

![init.png](img/init.png)


init functions are defined in package block and are used for:

* variables initialization if cannot be done in initialization expression,
* checking / fixing program’s state,
* registering,
* running one-time computations,
* and many more.


init function doesn’t take arguments neither returns any value. In contrast to main, identifier init is not declared so cannot be referenced:

```go
package main

import "fmt"

func init() {
    fmt.Println("init")
}

func main() {
    init()
}
```

and it gives “undefined: init” error while compilation.

	Formally speaking init identifier doesn’t introduce binding. In the same way works blank identifier represented by underscore character.


===

The most common use case of init function is to assign a value which cannot be calculated as a part of initialization expression:

```go
var precomputed = [20]float64{}

func init() {
    var current float64 = 1
    precomputed[0] = current
    for i := 1; i < len(precomputed); i++ {
        precomputed[i] = precomputed[i-1] * 1.2
    }
}
```

It’s not possible to use for loop as an expression (it’s a statement in Go) so putting it into init function solves the problem.

这个的意思是for loop只有法在函数里面，而这时候又需要precomputed，所以放在init()里
