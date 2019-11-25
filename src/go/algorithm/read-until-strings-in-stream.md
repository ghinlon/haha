# read until strings in stream

# Links

# Overview

# Code

* [telnet/conn.go#L376 · ziutek/telnet](https://github.com/ziutek/telnet/blob/master/conn.go#L376)

```go
func (c *Conn) readUntil(read bool, delims ...string) ([]byte, int, error) {
	if len(delims) == 0 {
		return nil, 0, nil
	}
	p := make([]string, len(delims))
	for i, s := range delims {
		if len(s) == 0 {
			return nil, 0, nil
		}
		p[i] = s
	}
	var line []byte
	for {
		b, err := c.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		if read {
			line = append(line, b)
		}
		for i, s := range p {
			if s[0] == b {
				if len(s) == 1 {
					return line, i, nil
				}
				p[i] = s[1:]
			} else {
				p[i] = delims[i]
			}
		}
	}
	panic(nil)
}
```

What makes me confusion is this code: 

How it works ?

```go
for i, s := range p {
	if s[0] == b {
		if len(s) == 1 {
			return line, i, nil
		}
		p[i] = s[1:]
	} else {
		p[i] = delims[i]
	}
}
```

# Find out

It looks like a state machine. Even though I don't really understand what is
state machine.


