# WEIXIN_WEB

# Links

* [微信网页版接口详解 - 闪客sun - 博客园](https://www.cnblogs.com/flashsun/p/8493306.html)
* [GitHub - songtianyi/wechat-go: go version wechat web api and message framework for building wechat robot](https://github.com/songtianyi/wechat-go)
* [企业微信API](https://work.weixin.qq.com/api/doc#90002/90151/90854)
* [workwx - GoDoc](https://godoc.org/github.com/xen0n/go-workwx#example-Workwx)

# Notes

* plugin should in a sepatate package.and I think it's more elegible that
  a session to register a plugin, than the vice visa.  
  I more like this way to register:  
  `func RegisterHash(h Hash, f func() hash.Hash)`  
* there's a lot of keyword `else`, `else` is the enemy of clearnity.
* It seems fetchJokes will be memory leak.

# How it work

```go
func CreateSession(common *Common, handlerRegister *HandlerRegister, qrmode int) (*Session, error) {

VVV

// a session to LoginAndServe ? 
func (s *Session) LoginAndServe(useCache bool) error

VVV

func (s *Session) serve() error {
	msg := make(chan []byte, 1000)
	// syncheck
	errChan := make(chan error)
	go s.producer(msg, errChan)
	for {
		select {
		case m := <-msg:
			go s.consumer(m)
		case err := <-errChan:
			// TODO maybe not all consumed messages ended
			return err
		}
	}
}
```


