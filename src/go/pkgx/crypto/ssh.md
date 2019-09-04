# [ssh - GoDoc](https://godoc.org/golang.org/x/crypto/ssh) 

# Links

* [Go and SSH for network devices – Egor Krivosheev – Medium](https://medium.com/@Vasya4k/go-and-ssh-for-network-devices-128937852ccb)

# Handle each command in ssh.Session.Shell()

# SendRequest in ssh.Channel interface

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

How Client.NewSession ?

First, `Client.OpenChannel()`, then use `newSession()` to turn it into a session.

`OpenChannel()` belongs to `Conn` interface. 

`mux` struct implements `Conn` interface.

`channel` struct implements `Channel` interface.


How `s.ch.SendRequest` work ?

channel.sendMessage(channelRequestMsg)

`sendMessage()` is `p = Marshal(msg)`, `ch.writePacket(p)`

`writePacket()` belongs to `packetConn` interface.

type `transport` implements `packetConn` interface




