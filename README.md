# About the fork

The fork is created to show that release v0.0.5 (commit hash c595d3a) of the upstream `niclabs/tcrsa` is in deed licensed under MIT.
Its license changing is after the time point.

Notice that while for copyright v0.0.5 is licensed under MIT, since copyright and patent are orthogonal, and MIT does not contain patent authorization, you still need to follow the new license for patent.

The following is the original README:

---

# Golang Threshold Cryptography Library - RSA implementation 
[![Go Report Card](https://goreportcard.com/badge/github.com/niclabs/tcrsa)](https://goreportcard.com/report/github.com/niclabs/tcrsa) [![Build Status](https://travis-ci.org/niclabs/tcrsa.svg?branch=master)](https://travis-ci.org/niclabs/tcrsa) [![GoDoc](https://godoc.org/github.com/niclabs/libtc-rsa?status.svg)](https://godoc.org/github.com/niclabs/tcrsa)

This library implements the cryptographic algorithms of Victor Shoup's paper [Practical Threshold Signatures](http://www.iacr.org/archive/eurocrypt2000/1807/18070209-new.pdf) in the Golang programming language. 

The codebase, commments and optimizations were ported from a previous implementation in C language, called [tchsm-libtc](https://github.com/niclabs/tchsm-libtc). As the previous implementation, the objective of this library is to provide a set of primitives to work with.

### Requirements

Due to Golang extensive standard library, this implementation does not have external requirements (obviously aside of [Golang](https://golang.org), version 1.12 or above).

### Installing

```shell
go get github.com/niclabs/tcrsa
```

To run the tests you just need to use `go test`:

```shell
go test
```
