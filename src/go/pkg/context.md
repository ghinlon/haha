# [Package context](https://golang.org/pkg/context/)

# Links

* [Go Concurrency Patterns: Context - The Go Blog](https://blog.golang.org/context)

# type Context interface

Background returns a non-nil, empty Context. It is never canceled, has no
values, and has no deadline. It is typically used by the main function,
initialization, and tests, and as the top-level Context for incoming requests. 

The context package provides functions to derive new Context values from
existing ones. These values form a tree: when a Context is canceled, all
Contexts derived from it are also canceled.

Background is the root of any Context tree; it is never canceled

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
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
```





