# [Package http](https://golang.org/pkg/net/http/)

# Links

* [Golang: Making HTTP Requests – Polyglot.Ninja()](https://polyglot.ninja/golang-making-http-requests/)
* [httpbin.org](https://httpbin.org/)
* [Golang Net HTTP Package – All Things Tech](https://mariadesouza.com/2018/11/28/golang-net-http-package/)(**appreciate**)
* [Build You Own Web Framework In Go | Nicolas Merouze](https://www.nicolasmerouze.com/build-web-framework-golang)(**Recommendation**)
* [HTTP | Network Programming with Go (golang)](https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/http/index.html)
* [Encoding and Decoding JSON, with Go’s net/http package | Kevin Burke](https://kev.inburke.com/kevin/golang-json-http/)
* [Writing Web Applications - The Go Programming Language](https://golang.org/doc/articles/wiki/)

# Related Packages

# [Package httputil](https://golang.org/pkg/net/http/httputil/)
# [Package httptest](https://golang.org/pkg/net/http/httptest/)

# Overview

# What is ServeMux?

A ServeMux is a HTTP request multiplexer or router that  matches the incoming
requests with a set of registered patterns and  calls  the associated handler
for that pattern.

simplely saying, a ServeMux is to Handle

# Tricky of HandlerFunc

The HandlerFunc makes it possible for us to pass in any function to make it
a Handler. 

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

# HandeFunc and Handle and HandlerFunc and Handler

HandleFunc needs a pattern and a func

Handle needs a pattern and a Handler

HandlerFunc implements Handler

and Handler is anything that implements ServeHTTP function


```go
// A Handler responds to an HTTP request.
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
```

All we need to make a handler is to implemented a method with the signature
`ServeHTTP(http.ResponseWriter, *http.Request)` on it. 

```go
type home struct {}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is my home page"))
}

mux := http.NewServeMux()
mux.Handle("/", &home{})

// or

func home(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is my home page"))
}

mux := http.NewServeMux()
mux.Handle("/", http.HandlerFunc(home))
// or
mux.HandleFunc("/", home)
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


# type Request struct 

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
```

## func (*Request) ParseForm 

ParseForm populates r.Form and r.PostForm. 

```go
type Request struct {

	...

    // Form contains the parsed form data, including both the URL
    // field's query parameters and the PATCH, POST, or PUT form data.
    // This field is only available after ParseForm is called.
    // The HTTP client ignores Form and uses Body instead.
    Form url.Values

    // PostForm contains the parsed form data from PATCH, POST
    // or PUT body parameters.
    //
    // This field is only available after ParseForm is called.
    // The HTTP client ignores PostForm and uses Body instead.
    PostForm url.Values // Go 1.1

    // MultipartForm is the parsed multipart form, including file uploads.
    // This field is only available after ParseMultipartForm is called.
    // The HTTP client ignores MultipartForm and uses Body instead.
    MultipartForm *multipart.Form

	...
}

func (r *Request) ParseForm() error

```

## Parsing Form Data

1. First, we need to use the `r.ParseForm()` method to parse the request body. 
1. We can then get to the form data contained in `r.PostForm` by using the
   `r.PostForm.Get()` method. 

Strictly speaking, the `r.PostForm.Get()` method that we’ve used above only
returns the first value for a specific form field. This means you can’t use it
with form fields which potentially send multiple values, such as a group of
checkboxes.

The underlying type of the `r.PostForm` map is `url.Values`, which in turn has
the underlying type `map[string][]string`. 

```go
func (v Values) Get(key string) string
    Get gets the first value associated with the given key. If there are no
    values associated with the key, Get returns the empty string. To access
    multiple values, use the map directly.
```

# type Response struct

```go
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

# func Error

```go
func Error(w ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}
```

# type Header struct

```go
type Header map[string][]string

type Header
    func (h Header) Add(key, value string)
    func (h Header) Clone() Header
    func (h Header) Del(key string)
    func (h Header) Get(key string) string
    func (h Header) Set(key, value string)
    func (h Header) Write(w io.Writer) error
    func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
```

for `json`:

```go
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
```

When sending a response Go will automatically set three system-generated
headers for you: `Date` and `Content-Length` and `Content-Type`.

The `Del()` method doesn’t remove system-generated headers. To suppress these,
you need to access the underlying header map directly and set the value to
`nil`.  If you want to suppress the Date header, for example, you need to
write:

```go
w.Header()["Date"] = nil
```

# ReadTimeout and IdleTimeout

One important thing. If you set `ReadTimeout` but don’t set `IdleTimeout`, then
`IdleTimeout` will default to using the same setting as `ReadTimeout`. 


