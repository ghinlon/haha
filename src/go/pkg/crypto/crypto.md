# [Package crypto](https://golang.org/pkg/crypto/)

# Overview

Package crypto collects common cryptographic constants. 

这个包的主要功能瞅着就是注册Hash方法使的

# 注册Hash和Hash类型

```go
type Hash uint

func (h Hash) HashFunc() Hash {
    return h
}
func (h Hash) Size() int
func (h Hash) Available() bool
func (h Hash) New() hash.Hash

var hashes = make([]func() hash.Hash, maxHash)

// RegisterHash registers a function that returns a new instance of the given
// hash function. This is intended to be called from the init function in
// packages that implement hash functions.
func RegisterHash(h Hash, f func() hash.Hash) {
    if h >= maxHash {
        panic("crypto: RegisterHash of unknown hash function")
    }
    hashes[h] = f
}
```

# Hash Management

* `hash.Hash` is `type Hash interface`
* `crypto.Hash` is `type Hash uint`
* `func RegisterHash(h Hash, f func() hash.Hash)` bind them.

[sha256.go](https://golang.org/src/crypto/sha256/sha256.go)里开头就有这样的代码:

```go
func init() {
    crypto.RegisterHash(crypto.SHA224, New224)
    crypto.RegisterHash(crypto.SHA256, New)
}
```

接口就是函数的集合.接口是"属性是函数"的类型.某种类型实现了这个函数的集合,就説这个类型实现了这个接口,这个类型就是这个接口.

RegisterHash函数的第一个参数是一个crypto.Hash无意义类型,第二参数是一个可以返回hash.Hash接口的函数.各个hash函数包里的New函数的本质是返回一个未导出的类型,它实现了hash.Hash接口.

为什么要注册呢,注册了有什么用? 当然目前我发现,注册了至少可以用crypto.Hash类型实现的这几个函数了.

其中,crypto.Hash类型实现的New函数,可以直接直接返回hash.Hash接口.其实就是调用了注册时提供的New函数.可以学习一下这种管理思路.


# 公私钥，签名，解密相关接口

```go
type PublicKey interface{}

type PrivateKey interface{}

type Signer interface {
    Public() PublicKey
    Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}

type SignerOpts interface {
    HashFunc() Hash
}

type Decrypter interface {
    Public() PublicKey
    Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

type DecrypterOpts interface{}
```
