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
* The `defer` keyword is almost entirely supported, with the exception of deferring some builtin functions.

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

Due to the above missing pieces and because parts of the standard library depend on the particular compiler/runtime in use, many packages do not yet compile. See the [list of compiling packages here]({{<ref "stdlib.md">}}) (but note that "compiling" does not imply that works entirely).

## Garbage collection

While not directly a language feature (the Go spec doesn't mention it), garbage collection is important for most Go programs to make sure their memory usage stays in reasonable bounds.

Garbage collection is currently supported on all platforms, although it works best on 32-bit chips. A simple conservative mark-sweep collector is used that will trigger a collection cycle when the heap runs out (that is fixed at compile time) or when requested manually using `runtime.GC()`. Some other collector designs are used for other targets, TinyGo will automatically pick a good GC for a given target.

Careful design may avoid memory allocations in main loops. You may want to compile with `-gc=none` and look at link errors to find out where allocations happen: the compiler inserts calls to `runtime.alloc` to allocate memory. For more information, see [heap allocation]({{<ref "heap-allocation.md">}}).

## A note on the `recover` builtin

The `recover` builtin is not yet supported. Instead, a `panic` will always terminate a program and `recover` simply returns nil.

This is a deviation from the Go spec but so far works well in practice. While there are no immediate plans to implement `recover`, if it can be shown to be necessary for compatibility it will be implemented. Please note that this comes at a cost: it means that every `defer` call will need some extra memory (both code and stack), so this feature is not free. It might also be architecture dependent. If it gets implemented, it will likely be opt-in to not increase code size for existing projects.
