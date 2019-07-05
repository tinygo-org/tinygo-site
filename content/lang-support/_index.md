---
title: "Go language features"
weight: 6
---

While TinyGo supports a big subset of the Go language, not everything is supported yet.

Here is a list of features that are supported:

* The subset of Go that directly translates to C is well supported. This includes all basic types and all regular control flow (including `switch`).
* Slices are well supported, with the exception of 3-index slicing (see below).
* Interfaces are quite stable and should work well in almost all cases. Type switches and type asserts are also supported, as well as calling methods on interfaces.
* Closures and bound methods are supported, for example inline anonymous (lambda-like) functions.
* The `defer` keyword is supported, with the exception of deferring a call on a function pointer. Immediately applied functions that are deferred are supported, however. In practice, function pointers are little used in deferred calls.

## Concurrency

At the time of writing (2019-07-05), support for goroutines and channels has not been fully realized but simple programs usually work. There may be problems with function pointers or starting more than two goroutines. Also, some things may unexpectedly allocate heap memory like calling a function that blocks. This situation should certainly improve in the future: at the moment, you should treat goroutines as an experimental feature.

## Cgo

While TinyGo embeds the [Clang](https://clang.llvm.org/) compiler to parse `import "C"` blocks, many features of Cgo are still unsupported. For example, `#cgo` statements are currently ignored.

## Reflection

Many packages, especially in the standard library, rely on reflection to work. The `reflect` package in the standards library is closely coupled to the main Go compilers and the runtime, so will have to be replaced. [Work is underway](https://github.com/tinygo-org/tinygo/pull/104) that provides at least initial support for reflection.

## Maps

Support for maps is not yet complete but is usable. You can use any type as a value, but only some types are acceptable as map keys. Also, they have not been optimized for performance and will cause linear lookup times in some cases.

Types supported as map keys include strings, integers, pointers, and structs/arrays that contain only these types. More complex types will depend on better reflection support.

## Standard library

Due to the above missing pieces and because parts of the standard library depend on the particular compiler/runtime in use, many packages do not yet compile.

## Garbage collection

While not directly a language feature (the Go spec doesn't mention it), garbage collection is important for most Go programs to make sure their memory usage stays in reasonable bounds.

Garbage collection is currently only supported on ARM microcontrollers (Cortex-M). For this platform, a simple conservative mark-sweep collector has been implemented. Other platforms will just allocate memory without ever freeing it.

Careful design may avoid memory allocations in main loops. You may want to compile with `-gc=none` and look at link errors to find out where allocations happen: the compiler inserts calls to `runtime.alloc` to allocate memory. For more information, see [heap allocation]({{<ref "heap-allocation.md">}}).

## Little used features

Some features are little used and there hasn't been a real need to implement them yet. These include:

* `recover()`: this can be useful sometimes but in general most programs work just fine with a `panic()` that simply aborts. Supporting `recover()` will also likely increase code size so it has also been left out at the moment for that reason. When `recover()` gets implemented, it will likely be disabled by default and can be enabled with a compiler flag.
* Slice expessions with 3 indexes that sets the capacity as well as the length. This feature was [introduced in Go 1.2](https://tip.golang.org/doc/go1.2#three_index).
