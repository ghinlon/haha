# FTP

# Links

* [File Transfer Protocol - Wikipedia](https://en.wikipedia.org/wiki/File_Transfer_Protocol)
* [List of FTP commands - Wikipedia](https://en.wikipedia.org/wiki/List_of_FTP_commands)
* [List of FTP server return codes - Wikipedia](https://en.wikipedia.org/wiki/List_of_FTP_server_return_codes)
* [ietf.org/rfc/rfc959.txt Page 36](https://www.ietf.org/rfc/rfc959.txt)


ftp是一个Request-Response的基于文本的协议。

# Q

我有一个疑问，送LIST, PWD...Cmd, 结果是通过哪种方式返回的，是在textproto.Conn里返回的，还是在net.Conn里返回的？

**A:** 

PWD是在msg里返回的。而LIST的结果是建立新的连接返回的，甚至不是在原来的net.Conn里。

那么不有法直接在原来的net.Conn里传输数据呢？

```go
// CurrentDir issues a PWD FTP command, which Returns the path of the current
// directory.
func (c *ServerConn) CurrentDir() (string, error) {
	_, msg, err := c.cmd(StatusPathCreated, "PWD")
	if err != nil {
		return "", err
	}

	start := strings.Index(msg, "\"")
	end := strings.LastIndex(msg, "\"")

	if start == -1 || end == -1 {
		return "", errors.New("Unsuported PWD response format")
	}

	return msg[start+1 : end], nil
}

// cmdDataConnFrom executes a command which require a FTP data connection.
// Issues a REST FTP command to specify the number of bytes to skip for the transfer.
func (c *ServerConn) cmdDataConnFrom(offset uint64, format string, args ...interface{}) (net.Conn, error) {
	conn, err := c.openDataConn()
	if err != nil {
		return nil, err
	}

	if offset != 0 {
		_, _, err := c.cmd(StatusRequestFilePending, "REST %d", offset)
		if err != nil {
			conn.Close()
			return nil, err
		}
	}

	_, err = c.conn.Cmd(format, args...)
	if err != nil {
		conn.Close()
		return nil, err
	}

	code, msg, err := c.conn.ReadResponse(-1)
	if err != nil {
		conn.Close()
		return nil, err
	}
	if code != StatusAlreadyOpen && code != StatusAboutToSend {
		conn.Close()
		return nil, &textproto.Error{Code: code, Msg: msg}
	}

	return conn, nil
}
```

# FTP REPLIES

GO的textproto包里的 [func (*Reader) ReadResponse](https://golang.org/pkg/net/textproto/#Reader.ReadResponse) 遵循这个协议的

`func (r *Reader) ReadResponse(expectCode int) (code int, message string, err error)`


	# 4.2. FTP REPLIES

      An FTP reply consists of a three digit number (transmitted as
            three alphanumeric characters) followed by some text.

      A reply is defined to contain the 3-digit code, followed by Space
      <SP>, followed by one line of text (where some maximum line length
      has been specified), and terminated by the Telnet end-of-line
      code.  There will be cases however, where the text is longer than
      a single line. 

         Thus the format for multi-line replies is that the first line
         will begin with the exact required reply code, followed
         immediately by a Hyphen, "-" (also known as Minus), followed by
         text.  The last line will begin with the same code, followed
         immediately by Space <SP>, optionally some text, and the Telnet
         end-of-line code.

            For example:
```
                                123-First line
                                Second line
                                  234 A line beginning with numbers
                                123 The last line
```

         The user-process then simply needs to search for the second
         occurrence of the same reply code, followed by <SP> (Space), at
         the beginning of a line, and ignore all intermediary lines.  If
         an intermediary line begins with a 3-digit number, the Server
         must pad the front  to avoid confusion.

         There are five values for the first digit of the reply code:

	    1yz   Positive Preliminary reply
	    2yz   Positive Completion reply
	    3yz   Positive Intermediate reply
	    4yz   Transient Negative Completion reply
	    5yz   Permanent Negative Completion reply



