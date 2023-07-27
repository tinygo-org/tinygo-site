---
title: "Using WASI"
weight: 3
description: |
  How to use TinyGo with the WebAssembly System Interface (WASI).
---

TinyGo is very useful for compiling programs for use on servers and other edge devices (WASI).

TinyGo programs can run in [Fastly Compute@Edge](https://developer.fastly.com/learning/compute/go/), [Fermyon Spin](https://developer.fermyon.com/spin/go-components), [wazero](https://wazero.io/languages/tinygo/) and many other WebAssembly runtimes.

Here is a small TinyGo program for use within a WASI host application:

```go
package main

//go:wasm-module yourmodulename
//export add
func add(x, y uint32) uint32 {
	return x + y
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
```

To compile the above TinyGo program for use on any WASI runtime:

```shell
tinygo build -o main.wasm -target=wasi main.go
```
