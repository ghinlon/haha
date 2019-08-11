# jazigo

# Links

* [GitHub - udhos/jazigo](https://github.com/udhos/jazigo)

# Overview

  代码质量相当差。

  接口的使法，log的处理，函数的设计，代码的风格，完全是乱七八糟，一塌糊涂。

  这个项目名字听起来就像夹边沟。一个危险的名字。

# Details

funny interface name

```go
type hasPrintf interface {
	Printf(fmt string, v ...interface{})
}
```

我不喜欢这个结构，我觉得它至少应该包含一个`Body`才算是一个`Request`.

OK, 所以这个名字才叫`FetchRequest`, 而并不是`Request`. kind of signal.

```go
// FetchRequest is a request for fetching a device configuration.
type FetchRequest struct {
	ID        string           // fetch this device
	ReplyChan chan FetchResult // reply on this channel
}
```

`func Scan` just write a FetchRequest with deviceid and response channel into
the reqChan. and this reqChan is like a global valuable. I don't like it.  It's
hard to debug.

I don't like a function named `XXXer`. `func xxxer` just make nonsense.

所有函数都是这么一长串的参数，真的是瞅得就头痛。

烦躁。Sigh...

```go
func Spawner(tab DeviceUpdater, logger hasPrintf, reqChan chan FetchRequest, repository, logPathPrefix string, options *conf.Options, ft *FilterTable) 
go d.Fetch(tab, logger, replyChan, 0, repository, logPathPrefix, opt, ft) // spawn per-request goroutine
func Scan(tab DeviceUpdater, devices []*Device, logger hasPrintf, opt *conf.AppConfig, reqChan chan FetchRequest) (int, int, int)
```

看不下去了。



