# [Package log](https://golang.org/pkg/log/)


A Logger represents an active logging object that generates lines of output to an io.Writer. Each logging operation makes a single call to the Writer's Write method. A Logger can be used simultaneously from multiple goroutines; it guarantees to serialize access to the Writer. 

# Constants

```go
const (
        Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
        Ltime                         // the time in the local time zone: 01:23:23
        Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
        Llongfile                     // full file name and line number: /a/b/c/d.go:23
        Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
        LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
        LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```

```go
var std = New(os.Stderr, "", LstdFlags)

func Print(v ...interface{})
func Printf(format string, v ...interface{})
func Println(v ...interface{})

func Fatal(v ...interface{})
func Fatalf(format string, v ...interface{})
func Fatalln(v ...interface{})

func Panic(v ...interface{})
func Panicf(format string, v ...interface{})
func Panicln(v ...interface{})

func SetFlags(flag int)
func Flags() int

func SetPrefix(prefix string)
func Prefix() string

func SetOutput(w io.Writer)
func Output(calldepth int, s string) error
```

# type Loger struct

```go
func New(out io.Writer, prefix string, flag int) *Logger
```
