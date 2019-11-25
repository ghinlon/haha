# Client and Server


1. Need a struct called Client or Server hold Options
1. Need a `func NewClient(conn net.Conn) ( *Client, error)` and a `func Dial(address string) (*Client, error)`
1. Need a `func Close() error`
1. cmd func
1. Server struct need `func Serve(l net.Listener) error` and `func ListenAndServe() error`
1. Server need type Session struct


```go
type Server struct {
	Addr string
	Hostname string
}

func (s *Server) Serve(l net.Listener) error {

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}

		ss := s.newSession(c)
		go ss.Serve()

	}
}


func (s *Server) ListenAndServe() error {
	addr := s.Addr
	if addr == "" {
		addr = ":1080"
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.Serve(l)

}

func ListenAndServe(addr string) error {
	s := &Server{Addr: addr}
	return s.ListenAndServe()
}

type Session struct {
	s *Server
	conn net.Conn
}

func (s *Server) newSession(conn net.Conn) *Session {
	return &Session{
		s:    s,
		conn: conn,
	}
}

func( ss *Session) Serve()  error {
	initSession()
	for {
		switch c {
		    case x: handleCmd()
		    default:
		}
	}
}
```
