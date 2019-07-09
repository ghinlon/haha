# [Package http](https://golang.org/pkg/net/http/)

# Links

* [Golang: Making HTTP Requests – Polyglot.Ninja()](https://polyglot.ninja/golang-making-http-requests/)
* [httpbin.org](https://httpbin.org/)
* [Golang Net HTTP Package – All Things Tech](https://mariadesouza.com/2018/11/28/golang-net-http-package/)(**appreciate**)
* [Build You Own Web Framework In Go | Nicolas Merouze](https://www.nicolasmerouze.com/build-web-framework-golang)(**Recommendation**)
* [HTTP | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/http/index.html)
* [Encoding and Decoding JSON, with Go’s net/http package | Kevin Burke](https://kev.inburke.com/kevin/golang-json-http/)
* [Writing Web Applications - The Go Programming Language](https://golang.org/doc/articles/wiki/)

# Overview

**What is ServeMux?**

```go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux
var defaultServeMux ServeMux

type ServeMux
    // NewServeMux allocates and returns a new ServeMux.
    func NewServeMux() *ServeMux { return new(ServeMux) }
	func (mux *ServeMux) Handle(pattern string, handler Handler)
	func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)
	func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

**What is a Handler?**

```go
// A Handler responds to an HTTP request.
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}

func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// HandleFunc registers the handler function for the given pattern.
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	if handler == nil {
		panic("http: nil handler")
	}
	mux.Handle(pattern, HandlerFunc(handler))
}
```

**It's a little bit tricky**

```go
// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

# type Client struct

**Always Use a Timeout**

To make a request with custom headers, use NewRequest and Client.Do. 

**So that’s it – create a client, create a request and then let the client Do the request.**

```go
type Client struct {
        Transport RoundTripper
        CheckRedirect func(req *Request, via []*Request) error
        Jar CookieJar
        Timeout time.Duration
}
 

func (c *Client) Do(req *Request) (*Response, error)
func (c *Client) Get(url string) (resp *Response, err error)
func (c *Client) Head(url string) (resp *Response, err error)
func (c *Client) Post(url, contentType string, body io.Reader) (resp *Response, err error)
func (c *Client) PostForm(url string, data url.Values) (resp *Response, err error)
```

# CookieJar

After the first time client.Do(req), The next time client.Do(req), it will add
cookies automatically.


# type Request struct and Response

NewRequest returns a new Request given a method, URL, and optional body.

NewRequest returns a Request suitable for use with Client.Do or Transport.RoundTrip.

```go
type Request struct {
	Method string
	URL *url.URL
        Proto      string // "HTTP/1.0"
        ProtoMajor int    // 1
        ProtoMinor int    // 0
	Header Header
	Body io.ReadCloser
	GetBody func() (io.ReadCloser, error) // Go 1.8
	ContentLength int64
	...
}


func NewRequest(method, url string, body io.Reader) (*Request, error)

func (t *Transport) RoundTrip(req *Request) (*Response, error)

type Response struct {
        Status     string // e.g. "200 OK"
        StatusCode int    // e.g. 200
        Proto      string // e.g. "HTTP/1.0"
        ProtoMajor int    // e.g. 1
        ProtoMinor int    // e.g. 0
	Header Header
	Body io.ReadCloser
	...
}
```

# What is HTTP referer

Referer logging is used to allow websites and web servers to identify where
people are visiting them from, for promotional or statistical purposes.

**Caution: The default behaviour of referer leaking puts websites at risk of
privacy and security breaches.**



