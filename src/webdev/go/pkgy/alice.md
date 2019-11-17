# [alice - GoDoc](https://godoc.org/github.com/justinas/alice)

# Links

* [justinas/alice: Painless middleware chaining for Go](https://github.com/justinas/alice) 

# Overview

In short, it transforms

`Middleware1(Middleware2(Middleware3(App)))`

to

`alice.New(Middleware1, Middleware2, Middleware3).Then(App)`



