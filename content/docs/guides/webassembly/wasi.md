---
title: "Using WASI"
weight: 3
description: |
  How to use TinyGo with the WebAssembly System Interface (WASI).
---

TinyGo is very useful for compiling programs for use on servers and other edge devices (WASI).

TinyGo programs can run in [Extism](https://github.com/extism/extism), [Fastly Compute@Edge](https://developer.fastly.com/learning/compute/go/), [Fermyon Spin](https://developer.fermyon.com/spin/go-components), [tau](https://github.com/taubyte/tau), [wazero](https://wazero.io/languages/tinygo/) and many other WebAssembly runtimes.

Both WASI Preview 1 (wasip1) and WASI Preview 2 (wasip2) are currently supported.

Here is a small TinyGo program for use within a WASI host application:

```go
package main

//go:wasmimport yourmodulename add
func add(x, y uint32) uint32 {
	return x + y
}

// main is required for the `wasip1` target, even if it isn't used.
func main() {}
```

To compile the above TinyGo program for use on any WASI runtime:

```shell
GOOS=wasip1 GOARCH=wasm tinygo build -o main.wasm main.go
```
