---
title: "Overview"
linkTitle: "Overview"
type: "docs"
weight: 20
description: >
  Short introduction to TinyGo
---

TinyGo is a new compiler for an existing programming language, the [Go programming language](https://golang.org/). TinyGo focuses on compiling code written in Go, but for smaller kinds of systems:

  - The Go compiler and tools (from [golang.org](https://golang.org/)) are the reference implementation of the Go programming language. They are primarily intended for server side programming but also used for command line tools and other purposes.
  - The TinyGo project implements the [exact same programming language](https://golang.org/ref/spec). However, TinyGo uses a different compiler and tools to make it suited for embedded systems and WebAssembly. It does this primarily by creating much smaller binaries and targeting a much wider variety of systems.

For example, the following piece of code does exactly the same thing in both compilers:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
```

You can compile it with similar commands:

```shell
go build -o hello ./hello.go
```

```shell
tinygo build -o hello ./hello.go
```

The programming language is the same. The standard library (including the fmt package) is the same. The only difference here, is which compiler is used, and the associated runtime.

What is significantly different is the output binary size. Here are the executables produced by each compiler:

```shell
1937340 may 12 12:42 helloworld-with-fmt.go1.16.3
 251376 may 12 12:42 helloworld-with-fmt.tinygo0.18
```

And then again after having the `strip` command run on each to remove all symbols and debugging info:

```shell
 837624 may 12 20:03 helloworld-with-fmt.go1.16.3-stripped
  10392 may 12 20:05 helloworld-with-fmt.tinygo0.18-stripped
```

In this case the Go compiled binary is 837k in size (or 1.9MB before stripping) while TinyGo produces a binary of only 10k in size (251k before stripping)!

This is 1% of the size of the original binary, which allows such binaries to be used on much smaller systems that were previously unsupported, simply because of the binary size.

By using TinyGo you can actually compile and run a binary on a variety of bare metal hardware platforms. For example, you can run this program directly on a [BBC micro:bit](https://microbit.org/):

    tinygo flash -target=microbit ./hello.go

Many different boards from vendors such as [Adafruit](https://www.adafruit.com/) and [Arduino](https://www.arduino.cc/) and many other companies are supported. For a complete list, see [this list](../../../docs/reference/microcontrollers).

Interested? Play around a bit with it on our [playground](https://play.tinygo.org/) or continue with [installing TinyGo](../../install).
