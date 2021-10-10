---
title: "Package organization"
weight: 1
description: >
  How code is organized in TinyGo.
---

There are many packages in the TinyGo repository. Here is an overview, and a brief description what they're for.

Packages can be split in two kinds: compiler packages and GOROOT packages. Compiler packages are the ones that are part of the TinyGo compiler itself: they will generate the object files and call a linker to link them together. The GOROOT packages are the packages you can import from a program, such as the TinyGo `"machine"` package.

## Compiler packages

  - root: contains the command line interface for the `tinygo` command and all its subcommands
  - `builder`: orchestrates the build
  - `loader`: loads and typechecks the code, and produces an AST
  - `compiler`: the compiler itself, makes little attempt at optimizing code
  - `interp`: tries to run package initializers at compile time as far as possible
  - `transform`: implements various optimizations necessary to produce working and efficient code

There are a few more, but you don't usually need to look at them. The above packages are the most important packages. To see how these packages work together, you can also take a look at [Pipeline]({{< ref "pipeline.md" >}}).

## `GOROOT` packages

The `GOROOT` directory in TinyGo is actually a constructed GOROOT, cached and reused over different compiler invocations. It is a combination of Go standard library packages and special TinyGo modified packages. For example, the `fmt` package is the exact same package that's used in regular Go, while the `runtime` package is a custom implementation for TinyGo and replaced altogether.

These replaced (or added) packages live in the src subdirectory:

  - `src/runtime`: the replacement runtime package that contains the things you would expect of a runtime package. It implements a heap with a garbage collector, a scheduler for goroutines, channels, and perhaps less obviously: it contains various things that are not implemented in the compiler like maps, the slice `append` built-in, and various other language features. It also implements a few device specific things, namely: runtime initialization and timers (used by the time packkage). These last two are device specific and need to be implemented per device.
  - `src/device`: contains hardware register access (memory-mapped I/O) without abstractions. Most of the subpackages are automatically generated. These packages are only used for baremetal programming.
  - `src/machine`: contains a hardware abstraction layer for chips. You could think of this as the equivalent of the `os` package: it provides portable APIs for common hardware peripherals. For example, every chip vendor implements [I2C]({{<ref "i2c.md">}}) differently but the machine package wraps this in a single `*machine.I2C` type that has the same interface on any supported chip. For more information, see the [machine package doucumentation]({{<ref "../../reference/machine.md">}})

There are a few more packages but these are the ones that are the most important for contributing to TinyGo.
