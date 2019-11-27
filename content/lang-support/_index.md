---
title: "Go language features"
weight: 6
---

While TinyGo supports a big subset of the Go language, not everything is supported yet.

Here is a list of features that are supported:

* The subset of Go that directly translates to C is well supported. This includes all basic types and all regular control flow (including `switch`).
* Slices are well supported.
* Interfaces are quite stable and should work well in almost all cases. Type switches and type asserts are also supported, as well as calling methods on interfaces. The only exception is comparing two interface values (but comparing against `nil` works).
* Closures and bound methods are supported, for example inline anonymous (lambda-like) functions.
* The `defer` keyword is supported, with the exception of a deferred call on an interface. This happens very infrequently in practice.

## Concurrency

At the time of writing (2019-11-27), support for goroutines and channels works for the most part. Support for concurrency on ARM microcontrollers is complete but may have some edge cases that don't work. Support for other platforms (such as WebAssembly) is a bit more limited: calling a blocking function may for example allocate heap memory.

## Cgo

While TinyGo embeds the [Clang](https://clang.llvm.org/) compiler to parse `import "C"` blocks, many features of Cgo are still unsupported. For example, `#cgo` statements are only partially supported.

## Reflection

Many packages, especially in the standard library, rely on reflection to work. The `reflect` package has been re-implemented in TinyGo and most common types like numbers, strings, and structs are supported now.

## Maps

Support for maps is not yet complete but is usable. You can use any type as a value, but only some types are acceptable as map keys. Also, they have not been optimized for performance and will cause linear lookup times in some cases.

Types supported as map keys include strings, integers, pointers, and structs/arrays that contain only these types.

## Standard library

Due to the above missing pieces and because parts of the standard library depend on the particular compiler/runtime in use, many packages do not yet compile. See the [list of compiling packages here]({{<ref "stdlib.md">}}).

## Garbage collection

While not directly a language feature (the Go spec doesn't mention it), garbage collection is important for most Go programs to make sure their memory usage stays in reasonable bounds.

Garbage collection is currently supported on all platforms except AVR. A simple conservative mark-sweep collector is used that will trigger a collection cycle when the heap runs out (that is fixed at compile time) or when requested manually using `runtime.GC()`.

Careful design may avoid memory allocations in main loops. You may want to compile with `-gc=none` and look at link errors to find out where allocations happen: the compiler inserts calls to `runtime.alloc` to allocate memory. For more information, see [heap allocation]({{<ref "heap-allocation.md">}}).

## Little used features

Some features are little used and there hasn't been a real need to implement them yet. These include:

* `recover()`: this can be useful sometimes but in general most programs work just fine with a `panic()` that simply aborts. Supporting `recover()` will also likely increase code size so it has also been left out at the moment for that reason. When `recover()` gets implemented, it will likely be disabled by default and can be enabled with a compiler flag.
