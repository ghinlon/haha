# [Package textproto](https://golang.org/pkg/net/textproto/)

这个包是http, smtp等包的基础，所以非常有必要搞清楚。

配合smtp包一时学这个包。

借着这个包，还可以同时回顾一下bufio包

textproto.Conn这个东西客户端需要，实现服务端的时候，有textproto.Reader, textproto.Writer就会得罢

# Overview

Package textproto implements generic support for text-based request/response protocols in the style of HTTP, NNTP, and SMTP.

The package provides:

Error, which represents a numeric error response from a server.

Pipeline, to manage pipelined requests and responses in a client.

Reader, to read numeric response code lines, key: value headers, lines wrapped with leading spaces on continuation lines, and whole text blocks ending with a dot on a line by itself.

Writer, to write dot-encoded text blocks.

Conn, a convenient packaging of Reader, Writer, and Pipeline for use with a single network connection. 

# type Conn struct

a textproto.Conn is just a conn with bufio.ReadWriter

```go
type Conn struct {
        Reader
        Writer
        Pipeline
        // contains filtered or unexported fields
}

func Dial(network, addr string) (*Conn, error) {
	c, err := net.Dial(network, addr)
	if err != nil {
		return nil, err
	}
	return NewConn(c), nil
}
func NewConn(conn io.ReadWriteCloser) *Conn {
	return &Conn{
		Reader: Reader{R: bufio.NewReader(conn)},
		Writer: Writer{W: bufio.NewWriter(conn)},
		conn:   conn,
	}
}

// c.Cmd内部使c.PrintfLine, c.W.PrintfLine内部使fmt.Fprintf,加上一个w.W.Write(crnl)
func (c *Conn) Cmd(format string, args ...interface{}) (id uint, err error) {
	id = c.Next()
	c.StartRequest(id)
	err = c.PrintfLine(format, args...)
	c.EndRequest(id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (c *Conn) Close() error
```

For example, a client might run a HELP command that returns a dot-body by using:

```go
id, err := c.Cmd("HELP")
if err != nil {
	return nil, err
}

c.StartResponse(id)
defer c.EndResponse(id)

if _, _, err = c.ReadCodeLine(110); err != nil {
	return nil, err
}
text, err := c.ReadDotBytes()
if err != nil {
	return nil, err
}
return c.ReadCodeLine(250)
```

# type Pipeline struct

A Pipeline manages a pipelined in-order request/response sequence.

To use a Pipeline p to manage multiple clients on a connection, each client should run:

```go
id := p.Next()	// take a number

p.StartRequest(id)	// wait for turn to send request
«send request»
p.EndRequest(id)	// notify Pipeline that request is sent

p.StartResponse(id)	// wait for turn to read response
«read response»
p.EndResponse(id)	// notify Pipeline that response is read
```

Pipeline这个东西，别个场地都有法使的。

```go
type Pipeline struct {
mu       sync.Mutex
id       uint
request  sequencer
response sequencer
}

func (p *Pipeline) Next() uint {
	p.mu.Lock()
	id := p.id
	p.id++
	p.mu.Unlock()
	return id
}
func (p *Pipeline) StartRequest(id uint)
func (p *Pipeline) EndRequest(id uint)
func (p *Pipeline) StartResponse(id uint)
func (p *Pipeline) EndResponse(id uint)
```


# type Reader struct

```go
type Reader struct {
        R *bufio.Reader
	        // contains filtered or unexported fields
}

func NewReader(r *bufio.Reader) *Reader

// the msg after process by ReadRespnse is in this form: "msg1\nmsg2\nmsg3\nmsg4", the end is no "\n"
// so can be taken to used by strings.Split(msg, "\n")
func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error) {
	code, continued, message, err := r.readCodeLine(expectCode)
	...
	for continued {
		line, err := r.ReadLine()
		...
		code2, continued, moreMessage, err = parseCodeLine(line, 0)
		if err != nil || code2 != code {
			message += "\n" + strings.TrimRight(line, "\r\n")
			continued = true
			continue
		}
		message += "\n" + moreMessage
	}
	...	
	return
}

func (r *Reader) ReadCodeLine(expectCode int) (code int, message string, err error)
func (r *Reader) ReadLine() (string, error)
func (r *Reader) ReadLineBytes() ([]byte, error)
func (r *Reader) ReadContinuedLine() (string, error)
func (r *Reader) ReadContinuedLineBytes() ([]byte, error)

func (r *Reader) DotReader() io.Reader
func (r *Reader) ReadDotLines() ([]string, error)
func (r *Reader) ReadDotBytes() ([]byte, error)

func (r *Reader) ReadMIMEHeader() (MIMEHeader, error)

// it can't handle the second form. but r.ReadResponse will handle form 2 correct.
func parseCodeLine(line string, expectCode int) (code int, continued bool, message string, err error) {
	...
	continued = line[3] == '-'
	code, err = strconv.Atoi(line[0:3])
	...
	message = line[4:]
	...
	return
}
```

ReadResponse reads a multi-line response of the form:

```
code-message line 1
code-message line 2
...
code message line n
```
another form of response accepted:

```
code-message line 1
message line 2
...
code message line n
```
If the prefix of the status does not match the digits in expectCode, ReadResponse returns with err set to &Error{code, message}. For example, if expectCode is 31, an error will be returned if the status is not in the range [310,319].

An expectCode <= 0 disables the check of the status code.

# type Writer struct

```go
type Writer struct {
        W *bufio.Writer
        // contains filtered or unexported fields
}

func NewWriter(w *bufio.Writer) *Writer

func (w *Writer) PrintfLine(format string, args ...interface{}) error {
	w.closeDot()
	fmt.Fprintf(w.W, format, args...)
	w.W.Write(crnl)
	return w.W.Flush()
}
func (w *Writer) DotWriter() io.WriteCloser
```


