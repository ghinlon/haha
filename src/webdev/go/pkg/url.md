# [Package url](https://golang.org/pkg/net/url/)

# Links

* [How to URL Encode a String in Golang | URLEncoder](https://www.urlencoder.io/golang/)

# Basic

From [How to URL Encode a String in Golang | URLEncoder](https://www.urlencoder.io/golang/):

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// Let's start with a base url
	baseUrl, err := url.Parse("http://www.mywebsite.com")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	// Add a Path Segment (Path segment is automatically escaped)
	baseUrl.Path += "path with?reserved characters"

	// Prepare Query Parameters
	params := url.Values{}
	params.Add("q", "Hello World")
	params.Add("u", "@rajeev")

	// Add Query Parameters to the URL
	baseUrl.RawQuery = params.Encode() // Escape Query Parameters

	fmt.Printf("Encoded URL is %q\n", baseUrl.String())
}
```

# type Values 

```go
type Values map[string][]string
    Values maps a string key to a list of values. It is typically used for query
    parameters and form values. Unlike in the http.Header map, the keys in a
    Values map are case-sensitive.

func ParseQuery(query string) (Values, error)
func (v Values) Add(key, value string)
func (v Values) Del(key string)
func (v Values) Encode() string
func (v Values) Get(key string) string
func (v Values) Set(key, value string)

func (u *URL) Query() Values
    Query parses RawQuery and returns the corresponding values. It silently
    discards malformed value pairs. To check errors use ParseQuery.

// http.Request
type Request struct {
    // URL specifies either the URI being requested (for server requests) or the
    // URL to access (for client requests).
    //
    // For server requests, the URL is parsed from the URI supplied on the
    // Request-Line as stored in RequestURI. For most requests, fields other than
    // Path and RawQuery will be empty. (See RFC 7230, Section 5.3)
    //
    // For client requests, the URL's Host specifies the server to connect to,
    // while the Request's Host field optionally specifies the Host header value to
    // send in the HTTP request.
    URL *url.URL

    // ... other fields elided ...
}
```


