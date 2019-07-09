# [Package pem](https://golang.org/pkg/encoding/pem/)


Package pem implements the PEM data encoding, which originated in Privacy
Enhanced Mail. The most common use of PEM encoding today is in TLS keys and
certificates. See RFC 1421.


```go
func Encode(out io.Writer, b *Block) error
func EncodeToMemory(b *Block) []byte
func Decode(data []byte) (p *Block, rest []byte)

// A Block represents a PEM encoded structure.
//
// The encoded form is:
//    -----BEGIN Type-----
//    Headers
//    base64-encoded Bytes
//    -----END Type-----
// where Headers is a possibly empty sequence of Key: Value lines.
type Block struct {
        Type    string            // The type, taken from the preamble (i.e. "RSA PRIVATE KEY").
        Headers map[string]string // Optional headers.
        Bytes   []byte            // The decoded bytes of the contents. Typically a DER encoded ASN.1 structure.
}
```
