---
title: "Go language features"
weight: 7
---

While TinyGo supports a big subset of the Go language, not everything is supported yet.

Here is a list of features that are supported:

* The subset of Go that directly translates to C is well supported. This includes all basic types (except for `complex64` and `complex128`) and all regular control flow (including `switch`).
* Slices are well supported, with the exception of 3-index slicing (see below).
* Interfaces, while tested less than other features, are quite stable and should work well in almost all cases. Type switches and type asserts are also supported, as well as calling methods on interfaces.
* Closures and bound methods are supported, for example inline anonymous (lambda-like) functions.
* The `defer` keyword is supported, with the exception of deferring a call on a function pointer. Immediately applied functions that are deferred are supported, however. In practice, function pointers are little used in deferred calls.

## Concurrency

At the time of writing (2019-01-29), support for goroutines and channels is weak. There is some support, but you will often encounter compiler errors when trying to use concurrency in more complicated ways (for example with function pointers). Also, some things may unexpectedly allocate memory like calling a function that blocks. This situation should certainly improve in the future, but at the moment you shouldn't rely on concurrency features to work well.

## Cgo

While TinyGo embeds the [Clang](https://clang.llvm.org/) compiler to parse `import "C"` blocks, many features of Cgo are still unsupported. For example, `#cgo` statements are currently ignored.

## Reflection

Many packages, especially in the standard library, rely on reflection to work. The `reflect` package in the standards library is closely coupled to the main Go compilers and the runtime, so will have to be replaced. [Work is underway](https://github.com/aykevl/tinygo/pull/104) that provides at least initial support for reflection.

## Maps

Support for maps is still very limited and in general you shouldn't rely on their availability. The main reason they are supported is so the `unicode` package compiles.

To use maps, you will currently have to deal with the following limitations:

* A map can store at most 8 entries. The exception is when it is initialized as a global variable, in which case it can store more entries.
* Map keys must be either strings, integers, or (nested) structs that contain only integers. Supporting more will likely depend on reflection for code size reasons.
* Map operations are likely to be very slow.
* The current implementation is little tested so may contain bugs.

## Standard library

Due to the above missing pieces and because parts of the standard library depend on the particular compiler/runtime in use, many packages do not yet compile.

## Garbage collection

While not directly a language feature (the Go spec doesn't mention it), garbage collection is important for most Go programs to make sure their memory usage stays in reasonable bounds.

Garbage collection is currently only supported on ARM microcontrollers (Cortex-M). For this platform, a simple conservative mark-sweep collector has been implemented. Other platforms will just allocate memory without ever freeing it.

Careful design may avoid memory allocations in main loops. You may want to compile with `-gc=none` and look at link errors to find out where allocations happen: the compiler inserts calls to `runtime.alloc` to allocate memory. For more information, see [heap allocation]({{<ref "heap-allocation.md">}}).

## Little used features

Some features are little used and there hasn't been a real need to implement them yet. These include:

* `recover()`: this can be useful sometimes but in general most programs work just fine with a `panic()` that simply aborts. Supporting `recover()` will also likely increase code size so it has also been left out at the moment for that reason. When `recover()` gets implemented, it will likely be disabled by default and can be enabled with a compiler flag.
* Arithmetic on complex numbers is not yet supported (add, sub, mul, div). However, the `complex`, `real` and `imag` builtins, as well as complex number constants, are all supported. Apart from the fact that complex number operations are little used, they are also [hard to implement correctly](https://github.com/golang/go/issues/29846).
* Slice expessions with 3 indexes that sets the capacity as well as the length. This feature was [introduced in Go 1.2](https://tip.golang.org/doc/go1.2#three_index).
