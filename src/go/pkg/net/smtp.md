# [Package smtp](https://golang.org/pkg/net/smtp/)

# Links

* [mailhog/MailHog: Web and API based SMTP testing](https://github.com/mailhog/MailHog)
* [smtpd - GoDoc](https://godoc.org/github.com/bradfitz/go-smtpd/smtpd)

# Run MailHog

```
docker pull mailhog/mailhog
docker run -p 1025:1025 -p 8025:8025 mailhog/mailhog
```

# type Client

```go
type Client struct {
        // Text is the textproto.Conn used by the Client. It is exported to allow for
        // clients to add extensions.
        Text *textproto.Conn
        // contains filtered or unexported fields
}

func Dial(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return NewClient(conn, host)
}
func NewClient(conn net.Conn, host string) (*Client, error)
// Calling this method is only necessary if the client needs control over the host name used. 
func (c *Client) Hello(localName string) error
func (c *Client) Auth(a Auth) error
func (c *Client) Mail(from string) error
func (c *Client) Rcpt(to string) error
func (c *Client) Data() (io.WriteCloser, error)
// 发完数据关闭连接
func (c *Client) Close() error
func (c *Client) Quit() error

func (c *Client) Extension(ext string) (bool, string)
func (c *Client) Noop() error
func (c *Client) Reset() error
func (c *Client) StartTLS(config *tls.Config) error
func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool)
func (c *Client) Verify(addr string) error

```

# func SendMail

```go
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
```

Example:

```go
func main() {
        addr := "localhost:1025"
        auth := smtp.CRAMMD5Auth("username", "password")
        from := "x@y.org"
        to := []string{"z1@z.org", "z2@z.org"}
        msg := []byte("From: from@y.org\r\n" +
                "To: z1@z.org, z2@z.org\r\n" +
                "Subject: discount Gophers!\r\n" +
                "\r\n" +
                "This is the email body.\r\n")

        err := smtp.SendMail(addr, auth, from, to, msg)
        if err != nil {
                log.Fatal(err)
        }
}
```
