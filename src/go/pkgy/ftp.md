# [package ftp](https://godoc.org/github.com/jlaffaye/ftp)

# Links

* [fclairamb/ftpserver: golang ftp server library with a sample implementation](https://github.com/fclairamb/ftpserver#quick-test-with-docker)

# Run ftpserver

```
# Creating a storage dir
mkdir data

# Starting the sample FTP server
docker run --rm -d -p 2121-2200:2121-2200 -v $(pwd)/data:/data fclairamb/ftpserver

# Connecting to it and uploading a file
ftp ftp://test:test@localhost:2121
!wget -c -O ftpserver-v0.3 https://github.com/fclairamb/ftpserver/releases/download/v0.3/ftpserver
put ftpserver-v0.3 ftpserver-v0.3
quit
ls -lh data/ftpserver-v0.3
```

# type Entry struct

```go
type Entry struct {
    Name string
    Type EntryType
    Size uint64
    Time time.Time
}
```

# type ServerConn struct

```go
type ServerConn struct {
    // Do not use EPSV mode
    DisableEPSV bool

    // Timezone that the server is in
    Location *time.Location
    // contains filtered or unexported fields
}

func Dial(addr string) (*ServerConn, error)
func DialTimeout(addr string, timeout time.Duration) (*ServerConn, error)

func (c *ServerConn) Login(user, password string) error
func (c *ServerConn) Logout() error
func (c *ServerConn) NoOp() error
func (c *ServerConn) Quit() error

func (c *ServerConn) CurrentDir() (string, error)
func (c *ServerConn) ChangeDir(path string) error
func (c *ServerConn) ChangeDirToParent() error
func (c *ServerConn) MakeDir(path string) error
func (c *ServerConn) RemoveDir(path string) error
func (c *ServerConn) RemoveDirRecur(path string) error

func (c *ServerConn) Rename(from, to string) error
func (c *ServerConn) Delete(path string) error
func (c *ServerConn) FileSize(path string) (int64, error)

func (c *ServerConn) List(path string) (entries []*Entry, err error)
func (c *ServerConn) NameList(path string) (entries []string, err error)

// The returned ReadCloser must be closed to cleanup the FTP data connection.
func (c *ServerConn) Retr(path string) (*Response, error)
func (c *ServerConn) RetrFrom(path string, offset uint64) (*Response, error)

func (c *ServerConn) Stor(path string, r io.Reader) error
func (c *ServerConn) StorFrom(path string, r io.Reader, offset uint64) error
```
