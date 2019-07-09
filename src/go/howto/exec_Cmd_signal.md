# How to signal to a exec.Cmd

We can:

```
cmd := exec.Command(name, args...)
err := cmd.Start()

// balabala

err = cmd.Process(os.Interrupt)
```

# Material

From `go doc exec.Cmd`:

```
type Cmd struct {

		...

        // Process is the underlying process, once started.
        Process *os.Process

        // ProcessState contains information about an exited process,
        // available after a call to Wait or Run.
        ProcessState *os.ProcessState

        // Has unexported fields.

		...
}
```

From `go doc os.Process`:

```
type Process struct {
        Pid int

        // Has unexported fields.
}
    Process stores the information about a process created by StartProcess.

func (p *Process) Signal(sig Signal) error
```

From `go doc os.Signal`:

```
type Signal interface {
        String() string
        Signal() // to distinguish from other Stringers
}
    A Signal represents an operating system signal. The usual underlying
    implementation is operating system-dependent: on Unix it is syscall.Signal.


var Interrupt Signal = syscall.SIGINT ...
```


