# [Command go](https://golang.org/cmd/go/)

# Links

# go get -u update all

You can also use `go get -u all` to update all packages in your `GOPATH`

# Command objdump 

[objdump - The Go Programming Language](https://golang.org/cmd/objdump/)


```
go tool objdump [-s symregexp] binary
go tool objdump binary start end
```

# Build

build for raspberry

`env GOOS=linux GOARCH=arm GOARM=7 go build -o guochan.raspberry`

