# Handle each command in ssh.Session.Shell()

Sat Aug 17 18:45:25 CST 2019

# Links

# SendRequest in ssh.Channel interface

I think I can learn something from it.

```go
type Channel interface {
		...
        // SendRequest sends a channel request.  If wantReply is true,
        // it will wait for a reply and return the result as a
        // boolean, otherwise the return value will be false. Channel
        // requests are out-of-band messages so they may be sent even
        // if the data stream is closed or blocked by flow control.
        // If the channel is closed before a reply is returned, io.EOF
        // is returned.
        SendRequest(name string, wantReply bool, payload []byte) (bool, error)
		...
}
    A Channel is an ordered, reliable, flow-controlled, duplex stream that is
    multiplexed over an SSH connection.

func (s *Session) SendRequest(name string, wantReply bool, payload []byte) (bool, error)
    SendRequest sends an out-of-band channel request on the SSH channel
    underlying the session.
```

Exactly, all `ssh.Session.{CombinedOutput, Output, Run, Start, Shell}` finally
is just to call `s.ch.SendRequest(...)`, `ch` is a `ssh.Channel` interface
embed in `ssh.Session`.


