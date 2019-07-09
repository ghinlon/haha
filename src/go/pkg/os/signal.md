# [Package signal](https://golang.org/pkg/os/signal/)


```go
func Ignore(sig ...os.Signal)
func Ignored(sig os.Signal) bool
func Notify(c chan<- os.Signal, sig ...os.Signal)
func Reset(sig ...os.Signal)
func Stop(c chan<- os.Signal)
```

# func Notify

Notify causes package signal to relay incoming signals to c. If no signals are provided, all incoming signals will be relayed to c. Otherwise, just the provided signals will.

Package signal will not block sending to c: the caller must ensure that c has sufficient buffer space to keep up with the expected signal rate. For a channel used for notification of just one signal value, a buffer of size 1 is sufficient. 

Ex:

```go
c := make(chan os.Signal, 1)
signal.Notify(c, os.Interrupt)

// Block until a signal is received.
s := <-c
fmt.Println("Got signal:", s)
```



