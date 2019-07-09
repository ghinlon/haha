# [Package rsa](https://golang.org/pkg/crypto/rsa/)

# Links

* [Prime Factorization](https://www.mathsisfun.com/prime-factorization.html)
* [Fundamental Theorem of Arithmetic](https://www.mathsisfun.com/numbers/fundamental-theorem-arithmetic.html)
* [Euler's totient function - Wikipedia](https://en.wikipedia.org/wiki/Euler's_totient_function)
* [RSA (cryptosystem) - Wikipedia](https://en.wikipedia.org/wiki/RSA_(cryptosystem))
* [What is RSA OAEP & RSA PSS in simple terms - Information Security Stack Exchange](https://security.stackexchange.com/questions/183179/what-is-rsa-oaep-rsa-pss-in-simple-terms)

# type PrivateKey struct

```go
type PrivateKey struct {
        PublicKey            // public part.
        D         *big.Int   // private exponent
        Primes    []*big.Int // prime factors of N, has >= 2 elements.

        // Precomputed contains precomputed values that speed up private
        // operations, if available.
        Precomputed PrecomputedValues
}

func GenerateMultiPrimeKey(random io.Reader, nprimes int, bits int) (*PrivateKey, error)
func GenerateKey(random io.Reader, bits int) (*PrivateKey, error) {
	return GenerateMultiPrimeKey(random, 2, bits)
}

func (priv *PrivateKey) Validate() error
func (priv *PrivateKey) Precompute()
func (priv *PrivateKey) Public() crypto.PublicKey
func (priv *PrivateKey) Decrypt(rand io.Reader, ciphertext []byte, opts crypto.DecrypterOpts) (plaintext []byte, err error)
//  implements crypto.Signer interface
func (priv *PrivateKey) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error)
```

# type PublicKey struct

```go
type PublicKey struct {
        N *big.Int // modulus
        E int      // public exponent
}

func (pub *PublicKey) Size() int
```

# Func

```go
func EncryptOAEP(hash hash.Hash, random io.Reader, pub *PublicKey, msg []byte, label []byte) ([]byte, error)
func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)

func DecryptOAEP(hash hash.Hash, random io.Reader, priv *PrivateKey, ciphertext []byte, label []byte) ([]byte, error)
func DecryptPKCS1v15(rand io.Reader, priv *PrivateKey, ciphertext []byte) ([]byte, error)
func DecryptPKCS1v15SessionKey(rand io.Reader, priv *PrivateKey, ciphertext []byte, key []byte) error

func SignPSS(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte, opts *PSSOptions) ([]byte, error)
func SignPKCS1v15(rand io.Reader, priv *PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error)

func VerifyPSS(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte, opts *PSSOptions) error
func VerifyPKCS1v15(pub *PublicKey, hash crypto.Hash, hashed []byte, sig []byte) error
```
