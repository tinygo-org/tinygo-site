---
title: "Overview"
weight: 1
---

TinyGo is a new compiler for an existing programming language, the [Go programming language](https://golang.org/). It is the same programming language but the compiler focuses on different kinds of systems:

  - The Go toolchain (from [golang.org](https://golang.org/)) is the reference implementation of the Go programming language. It is primarily intended for server side programming but is also used for command line tools and other purposes.
  - The TinyGo project is the exact same language but it is a different compiler and an ecosystem around this compiler to make it suited for embedded and WebAssembly. It does this primarily by creating much smaller binaries and targeting a much wider variety of systems.

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

The language is the same. The standard library (including the fmt package) is the same. The only difference here is the compiler and associated runtime.

What is different is the output binary size. In this case the Go compiled binary is 1.4MB in size (or 2.0MB before stripping) while TinyGo produces a binary of only 50kB in size (225kB before stripping)! This is 3-4% of the size of the original binary, which allows such binaries to be used on much smaller systems that were previously unsupported simply because of the binary size.

What is more, is that you can actually run this binary on a variety of bare metal hardware platforms. For example, you can run this directly on a [BBC micro:bit](https://microbit.org/):


    tinygo flash -target=microbit ./hello.go

Many different boards from vendors such as [Adafruit](https://www.adafruit.com/) and [Arduino](https://www.arduino.cc/) and many other companies are supported. For a complete list, see [this list](../../microcontrollers).

Interested? Play around a bit with it on our [playground](https://play.tinygo.org/) or continue with [installing TinyGo](install).
