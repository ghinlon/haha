# [Package context](https://golang.org/pkg/context/)

# Links

* [Go Concurrency Patterns: Context - The Go Blog](https://blog.golang.org/context)

# Overview


Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it. The Context should be the first parameter, typically named ctx:

```go
func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
```

Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.

Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.

The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines. 

这个东西用于资源释放.大概是当什么情况发生了,就释放资源.

这个东西设计出来,大概是为了解决子routine不能杀掉父routine的问题的.

# type Context interface

>Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests. 

>The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.

>Background is the root of any Context tree; it is never canceled

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}

// An emptyCtx is never canceled, has no values, and has no deadline. It is not
// struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}
func (*emptyCtx) Done() <-chan struct{} {
	return nil
}
func (*emptyCtx) Err() error {
	return nil
}
func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}

var (
    background = new(emptyCtx)
    todo       = new(emptyCtx)
)

func Background() Context {
    return background
}
func TODO() Context {
    return todo
}
func WithValue(parent Context, key, val interface{}) Context 

type CancelFunc func()

func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```





