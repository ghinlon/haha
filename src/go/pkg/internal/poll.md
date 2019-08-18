# [Package poll](https://golang.org/pkg/internal/poll/)

# Links

# type FD struct

FD is a file descriptor. The net and os packages use this type as a field of
a larger type representing a network connection or OS file. 

```go
type FD struct {

        // System file descriptor. Immutable until Close.
        Sysfd int

        // Whether this is a streaming descriptor, as opposed to a
        // packet-based descriptor like a UDP socket. Immutable.
        IsStream bool

        // Whether a zero byte read indicates EOF. This is false for a
        // message based socket connection.
        ZeroReadIsEOF bool
        // contains filtered or unexported fields
}

func (fd *FD) Read(p []byte) (int, error) {
        ...
		n, err := syscall.Read(fd.Sysfd, p)
		...
}
```



