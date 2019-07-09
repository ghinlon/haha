# [Package os](https://golang.org/pkg/os/)

# Constants

Flags

```go
const (
        // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
        O_RDONLY int = syscall.O_RDONLY // open the file read-only.
        O_WRONLY int = syscall.O_WRONLY // open the file write-only.
        O_RDWR   int = syscall.O_RDWR   // open the file read-write.
        // The remaining values may be or'ed in to control behavior.
        O_APPEND int = syscall.O_APPEND // append data to the file when writing.
        O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
        O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
        O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
        O_TRUNC  int = syscall.O_TRUNC  // if possible, truncate file when opened.
)
```
DevNull

`const DevNull = "/dev/null"`


# Dir and File

```go
func Getwd() (dir string, err error)
func Chdir(dir string) error
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error
func Remove(name string) error
func RemoveAll(path string) error

// Rename renames (moves) oldpath to newpath. If newpath already exists and is not a directory, Rename replaces it.
func Rename(oldpath, newpath string) error
func Remove(name string) error
func RemoveAll(path string) error
func Truncate(name string, size int64) error
func IsPathSeparator(c uint8) bool

func TempDir() string
func UserCacheDir() (string, error)

func (f *File) Chdir() error
func (f *File) Readdir(n int) ([]FileInfo, error)
func (f *File) Readdirnames(n int) (names []string, err error)

func ioutil.ReadFile(filename string) ([]byte, error)
// WriteFile writes data to a file named by filename. 
// If the file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing. 
func ioutil.WriteFile(filename string, data []byte, perm os.FileMode) error
func ioutil.ReadDir(dirname string) ([]os.FileInfo, error)
func ioutil.TempDir(dir, prefix string) (name string, err error)
func ioutil.TempFile(dir, pattern string) (f *os.File, err error)
```

# type File struct

```go
type File struct {
        // contains filtered or unexported fields
}

// Create creates the named file with mode 0666 (before umask), truncating it if it already exists. 
func Create(name string) (*File, error)
func NewFile(fd uintptr, name string) *File
func Open(name string) (*File, error)
func OpenFile(name string, flag int, perm FileMode) (*File, error)
func Pipe() (r *File, w *File, err error)

func (f *File) Name() string

func ioutil.TempFile(dir, pattern string) (f *os.File, err error)
```

# FileInfo

```go
type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}
func Stat(name string) (FileInfo, error)
func Lstat(name string) (FileInfo, error)
func SameFile(fi1, fi2 FileInfo) bool

func (f *File) Stat() (FileInfo, error)
func (f *File) Readdir(n int) ([]FileInfo, error)

func ioutil.ReadDir(dirname string) ([]os.FileInfo, error)
```

# FileMode

A FileMode represents a file's mode and permission bits.

```go
type FileMode uint32

const (
        // The single letters are the abbreviations
        // used by the String method's formatting.
        ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
        ModeAppend                                     // a: append-only
        ModeExclusive                                  // l: exclusive use
        ModeTemporary                                  // T: temporary file; Plan 9 only
        ModeSymlink                                    // L: symbolic link
        ModeDevice                                     // D: device file
        ModeNamedPipe                                  // p: named pipe (FIFO)
        ModeSocket                                     // S: Unix domain socket
        ModeSetuid                                     // u: setuid
        ModeSetgid                                     // g: setgid
        ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
        ModeSticky                                     // t: sticky
        ModeIrregular                                  // ?: non-regular file; nothing else is known about this file

        // Mask for the type bits. For regular files, none will be set.
        ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeIrregular

        ModePerm FileMode = 0777 // Unix permission bits
)

func FileInfo.Mode() FileMode 

func Chmod(name string, mode FileMode) error
func OpenFile(name string, flag int, perm FileMode) (*File, error)
func Mkdir(name string, perm FileMode) error
func MkdirAll(path string, perm FileMode) error

func (m FileMode) IsDir() bool
func (m FileMode) IsRegular() bool
// Perm returns the Unix permission bits in m.
func (m FileMode) Perm() FileMode
func (m FileMode) String() string

func (f *File) Chmod(mode FileMode) error
```
# uid gid

```go
func Chown(name string, uid, gid int) error
func Lchown(name string, uid, gid int) error
func Getuid() int
func Getgid() int
func Geteuid() int
func Getegid() int
func Getgroups() ([]int, error)

func (f *File) Chown(uid, gid int) error
```

# times

```go
func Chtimes(name string, atime time.Time, mtime time.Time) error
```

# Link

```go
func Link(oldname, newname string) error
func Symlink(oldname, newname string) error
func Readlink(name string) (string, error)
```


# Env

```go
func Expand(s string, mapping func(string) string) string
func ExpandEnv(s string) string
func Setenv(key, value string) error
func Unsetenv(key string) error
func Getenv(key string) string
func LookupEnv(key string) (string, bool)
func Environ() []string
func Clearenv()
```

# os

```go
func Hostname() (name string, err error)
func Getpagesize() int
func Exit(code int)
```

# Is(error)

```go
func IsExist(err error) bool
func IsNotExist(err error) bool
func IsPermission(err error) bool
func IsTimeout(err error) bool
```

# Process

```go
func Executable() (string, error)
func Getpid() int
func Getppid() int
```
