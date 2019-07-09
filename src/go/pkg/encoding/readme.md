# [Package encoding](https://golang.org/pkg/encoding/)

<!-- ToC start -->

# Table of Contents

1. [Links](#links)
1. [Package files](#package-files)
1. [Overview](#overview)
1. [4个编解组接口](#4个编解组接口)
<!-- ToC end -->

# Links

# Package files

[encoding.go](https://golang.org/src/encoding/encoding.go)

# Overview

Package encoding defines interfaces shared by other packages that convert data to and from byte-level and textual representations. Packages that check for these interfaces include encoding/gob, encoding/json, and encoding/xml. As a result, implementing an interface once can make a type useful in multiple encodings. Standard types that implement these interfaces include time.Time and net.IP. The interfaces come in pairs that produce and consume encoded data. 

# 4个编解组接口

二进制編解组接口, 文本編解组接口

```go
type BinaryMarshaler interface {
    MarshalBinary() (data []byte, err error)
}

type BinaryUnmarshaler interface {
    UnmarshalBinary(data []byte) error
}

type TextMarshaler interface {
    MarshalText() (text []byte, err error)
}

type TextUnmarshaler interface {
    UnmarshalText(text []byte) error
}
```

