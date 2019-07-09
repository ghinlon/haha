# [Package mail](https://golang.org/pkg/net/mail/)

# Links

* [Email address - Wikipedia](https://en.wikipedia.org/wiki/Email_address)
* [International email - Wikipedia](https://en.wikipedia.org/wiki/International_email)

# type Message struct

```go
type Message struct {
        Header Header
	Body   io.Reader
}

func ReadMessage(r io.Reader) (msg *Message, err error)
```

# type Header

```go
type Header map[string][]string

func (h Header) Get(key string) string {
	return textproto.MIMEHeader(h).Get(key)
}
func (h Header) AddressList(key string) ([]*Address, error)
// Date parses the Date header field.
func (h Header) Date() (time.Time, error) {
	hdr := h.Get("Date")
	if hdr == "" {
		return time.Time{}, ErrHeaderNotPresent
	}
	return ParseDate(hdr)
}

func ParseDate(date string) (time.Time, error)
```

#  type Address struct

```go
type Address struct {
        Name    string // Proper name; may be empty.
	        Address string // user@domain
}

func ParseAddress(address string) (*Address, error)
func ParseAddressList(list string) ([]*Address, error)

func (a *Address) String() string
```

# type AddressParser struct

```go
type AddressParser struct {
        // WordDecoder optionally specifies a decoder for RFC 2047 encoded-words.
	WordDecoder *mime.WordDecoder
}

func (p *AddressParser) Parse(address string) (*Address, error)
func (p *AddressParser) ParseList(list string) ([]*Address, error)
```
