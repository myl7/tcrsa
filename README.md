# About the fork

The fork is created to show that release v0.0.5 (commit hash c595d3a) of the upstream `niclabs/tcrsa` is in deed licensed under MIT.
Its license changing is after the time point.

Notice that while for copyright v0.0.5 is licensed under MIT, since copyright and patent are orthogonal, and MIT does not contain patent authorization, you still need to follow the new license for patent.

The upstream overview follows, with installation and maintenance instructions for this fork:

---

# Golang Threshold Cryptography Library - RSA implementation 
[![Go tests](https://github.com/myl7/tcrsa/actions/workflows/ci.yaml/badge.svg)](https://github.com/myl7/tcrsa/actions/workflows/ci.yaml) [![Go Reference](https://pkg.go.dev/badge/github.com/myl7/tcrsa.svg)](https://pkg.go.dev/github.com/myl7/tcrsa)

This library implements the cryptographic algorithms of Victor Shoup's paper [Practical Threshold Signatures](http://www.iacr.org/archive/eurocrypt2000/1807/18070209-new.pdf) in the Golang programming language. 

The codebase, commments and optimizations were ported from a previous implementation in C language, called [tchsm-libtc](https://github.com/niclabs/tchsm-libtc). As the previous implementation, the objective of this library is to provide a set of primitives to work with.

### Requirements

Due to Golang extensive standard library, this implementation does not have external requirements (obviously aside of [Golang](https://golang.org), version 1.12 or above).

### Installing

```shell
go get github.com/myl7/tcrsa@v0.0.6
```

To run the tests you just need to use `go test`:

```shell
go test
```

## Go maintenance checks

CI builds, vets, and runs the existing tests on Go 1.27.2 and Go 1.26.9.
Go 1.27.2 is recommended for reproduction; the existing module requirements
and Go directive are preserved:

```sh
go build ./...
go vet ./...
go test -race -timeout=10m ./...
```

These checks do not measure paper benchmark performance. Reproducing published
measurements requires the original toolchain and experimental environment.

The maintained fork is published as `github.com/myl7/tcrsa` at v0.0.6.

Its historical `go 1.12` directive preserves the module's language and
compatibility settings; it does not require installing Go 1.12. CI uses the
specified toolchain directly, without automatic switching or a workspace.
