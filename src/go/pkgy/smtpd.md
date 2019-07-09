# [smtpd - GoDoc](https://godoc.org/github.com/bradfitz/go-smtpd/smtpd)

import "github.com/bradfitz/go-smtpd/smtpd"


# type Server struct

```go
type Server struct {
    Addr         string        // TCP address to listen on, ":25" if empty
    Hostname     string        // optional Hostname to announce; "" to use system hostname
    ReadTimeout  time.Duration // optional read timeout
    WriteTimeout time.Duration // optional write timeout

    PlainAuth bool // advertise plain auth (assumes you're on SSL)

    // OnNewConnection, if non-nil, is called on new connections.
    // If it returns non-nil, the connection is closed.
    OnNewConnection func(c Connection) error

    // OnNewMail must be defined and is called when a new message beings.
    // (when a MAIL FROM line arrives)
    OnNewMail func(c Connection, from MailAddress) (Envelope, error)
}

func (srv *Server) ListenAndServe() error
func (srv *Server) Serve(ln net.Listener) error

func (srv *Server) newSession(rwc net.Conn) (s *session, err error) 
```


