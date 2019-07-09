# [telnet - GoDoc](https://godoc.org/github.com/ziutek/telnet)

# Links

* [BatMUD](https://www.bat.org/) an awesome game!

# Overview

E.g.: 

```go
func main() {
	conn, err := telnet.Dial("tcp", "bat.org:23")
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 512)
	for {
		n, err := conn.Read(buf) // Use raw read to find issue #15.
		os.Stdout.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}

	}
}
```
