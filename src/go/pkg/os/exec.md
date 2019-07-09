# [Package exec](https://golang.org/pkg/os/exec/)


# Overview

Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout, connect I/O with pipes, and do other adjustments.

Unlike the "system" library call from C and other languages, the os/exec package intentionally does not invoke the system shell and does not expand any glob patterns or handle other expansions, pipelines, or redirections typically done by shells. The package behaves more like C's "exec" family of functions. To expand glob patterns, either call the shell directly, taking care to escape any dangerous input, or use the path/filepath package's Glob function. To expand environment variables, use package os's ExpandEnv.

```go
func LookPath(file string) (string, error)
```

# type Cmd struct

```go
func Command(name string, arg ...string) *Cmd
func CommandContext(ctx context.Context, name string, arg ...string) *Cmd

func (c *Cmd) CombinedOutput() ([]byte, error)
func (c *Cmd) Output() ([]byte, error)
func (c *Cmd) Run() error
func (c *Cmd) Start() error
func (c *Cmd) Wait() error
func (c *Cmd) StderrPipe() (io.ReadCloser, error)
func (c *Cmd) StdinPipe() (io.WriteCloser, error)
func (c *Cmd) StdoutPipe() (io.ReadCloser, error)
```

