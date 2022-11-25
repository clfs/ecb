# ecb

<a href="https://pkg.go.dev/github.com/clfs/ecb"><img src="https://pkg.go.dev/badge/github.com/clfs/ecb.svg" alt="Go Reference"></a>

> **Warning**
> Never use this in a production system. ECB mode is not secure.

This package implements the electronic codebook (ECB) block cipher mode, which
is missing from Go's standard library.

It satisfies all the requirements of the [`crypto/cipher.BlockMode`][0]
interface.

[0]: https://pkg.go.dev/crypto/cipher#BlockMode