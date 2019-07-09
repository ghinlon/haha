# [zap - GoDoc](https://godoc.org/go.uber.org/zap)

import "go.uber.org/zap"

# Links

* [uber-go/zap: Blazing fast, structured, leveled logging in Go.](https://github.com/uber-go/zap)
* [Using Zap - Simple use cases · sandipb.net](https://blog.sandipb.net/2018/05/02/using-zap-simple-use-cases/)


# Quick Start

```go
// sugar := zap.NewExample().Sugar()
// defer sugar.Sync()

logger, _ := zap.NewProduction()
defer logger.Sync() // flushes buffer, if any
sugar := logger.Sugar()
sugar.Infow("failed to fetch URL",
  // Structured context as loosely typed key-value pairs.
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)
sugar.Infof("Failed to fetch URL: %s", url)
```

When performance and type safety are critical, use the Logger. It's even faster than the SugaredLogger and allocates far less, but it only supports structured logging.

# type Logger struct

```go
func L() *Logger
func New(core zapcore.Core, options ...Option) *Logger
func NewExample(options ...Option) *Logger
func NewDevelopment(options ...Option) (*Logger, error)
func NewProduction(options ...Option) (*Logger, error)
func NewNop() *Logger

func (log *Logger) Sugar() *SugaredLogger
func (log *Logger) With(fields ...Field) *Logger
func (log *Logger) Sync() error
```

# type SugaredLogger struct

```go
func S() *SugaredLogger

func (s *SugaredLogger) Desugar() *Logger
func (s *SugaredLogger) With(args ...interface{}) *SugaredLogger
func (s *SugaredLogger) Sync() error
```



