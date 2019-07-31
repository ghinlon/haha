# meshbird

If you want to understand something, you always need to get the **structure**
of the thing.

# Links

* [meshbird/meshbird: Distributed private networking](https://github.com/meshbird/meshbird)

# Overview

# Review

Structure:  

```go
a.server = transport.NewServer(a.config.BindAddrs, a, a.config.Key)
a.bootstrap()

go a.server.Start()
return a.runIface()
```

# func bootstrap

What does bootstrap do is `func NewPeer` and do `peer.Start()`

```go
peer := NewPeer([]string{seedAddr}, a.config, a.getRoutes)
peer.Start()

func (p *Peer) Start() {
	p.client.Start()
	go p.process()
}
```

## type Peer struct

Why peer do `SendPing()` every second ?

## func (c *Client) Start() 

```go
for _, remoteAddr := range c.remoteAddrs {
			c.wg.Add(1)
			conn := NewClientConn(remoteAddr, c.key, connIndex, &c.wg)
			c.conns[connIndex] = conn
			go c.conns[connIndex].run()
		}
```

# type ClientConn struct

* The method `func (cc *ClientConn) run()` really do is run `func (cc
  *ClientConn) process() error`
* What `func ( cc *ClientConn) process()` do is get data from it's embeded
  `chan []byte`, then write to it's `*net.TCPConn` field, and I think the
  implementation of this `write()` method is not beautiful.

The protocol is:


	|--------|--------|------|
	| secure | length | data |
	|--------|--------|------|
	| 1      | 2      | N    |
	|--------|--------|------|

# go a.server.Start()

What it does is just `s.listen(tcpAddr)` on `config.BindAddrs`.

what `s.listen` do is `go serverConn.run()`

```go
func (sc *ServerConn) run() {
	...
	sc.handler.OnData(data)
	...
}
```

# func (a *App) OnData(buf []byte)


# func (a *App) runIface() error 



