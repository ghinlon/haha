# [expect - GoDoc](https://godoc.org/github.com/google/goexpect)

# Links

* [GitHub - google/goexpect: Expect for Go](https://github.com/google/goexpect)


# Review

```go
		select {
		case <-timer.C:
			// Expect timeout.
			nr, err := io.Copy(&tbuf, e)
			if err != nil {
				return tbuf.String(), nil, -1, fmt.Errorf("io.Copy failed: %v", err)
			}
			// If we got no new data we return otherwise give it another chance to match.
			if nr == 0 {
				return tbuf.String(), nil, -1, TimeoutError(timeout)
			}
			timer = time.NewTimer(timeout)
		case <-chTicker.C:
			// Periodical timer to make sure data is handled in case the <-e.rcv channel
			// was missed.
			if _, err := io.Copy(&tbuf, e); err != nil {
				return tbuf.String(), nil, -1, fmt.Errorf("io.Copy failed: %v", err)
			}
		case <-e.rcv:
			// Data to fetch.
			if _, err := io.Copy(&tbuf, e); err != nil {
				return tbuf.String(), nil, -1, fmt.Errorf("io.Copy failed: %v", err)
			}
		}
```

I think this select is useless, it's the same as just one line `io.Copy(&tbuf, e)`,
cause no matter in which case, it just read out from `e` anyway.

am I wrong ?

I think there should be have a reliable way to know:

1. e.Send(in) has done.
2. the sended in command has response back.



