---
title: "Go language features"
weight: 6
description: |
  Which Go language features are supported by TinyGo and which are still a work in progress.
---

The TinyGo compiler implements all major Go language features, although some details are missing. Below you will find a description of some of the missing features, as of june 2023.

If you're wondering "Does TinyGo support feature X", we often cannot give a good answer! The rest of this page gives a good indication, but other than that you will just have to try for yourself to see whether a particular piece of software works with TinyGo. There are just way too many edge cases that are not entirely supported or work slightly differently, many of which we don't even know about.

## Cgo

While TinyGo embeds the [Clang](https://clang.llvm.org/) compiler to parse `import "C"` blocks, some features of Cgo are still unsupported or may work slightly differently. For example, `#cgo` statements are only partially supported.

## Reflection

Many packages, especially in the standard library, rely on reflection to work. The `reflect` package has been re-implemented in TinyGo and most of it works, but some parts are not yet fully supported.

## Maps

Maps generally work fine, but may be slower than you expect them to be. There are a few reasons for this, one of which is that some types (like structs) may internally be compared using reflection instead of using a dedicated hash/compare function.

## Standard library

Due to the above missing pieces and because parts of the standard library depend on the particular compiler/runtime in use, many packages do not yet compile. See the [list of compiling packages here]({{<ref "stdlib.md">}}) (but note that "compiling" does not imply that works entirely).

## Garbage collection

Garbage collection generally works fine, but may work not as well on very small chips (AVR) and on WebAssembly. It is also a lot slower than the usual Go garbage collector.

Careful design may avoid memory allocations in main loops where they can reduce performance a lot. You may want to compile with `-print-allocs=.` to find out where allocations happen and why they happen. For more information, see [heap allocation]({{<ref "heap-allocation.md">}}).

## `recover` builtin

The `recover` builtin is supported on most architectures, with the notable exception of WebAssembly. For WebAssembly, we need the [exception handling proposal](https://webassembly.org/roadmap/) which is implemented in browsers but is not implemented in many WASI runtimes.

On architectures where `recover` is not implemented, a panic will always exit the program without running any deferred functions.

Some notes on `recover` support in TinyGo:

  * We don't follow the Go language specification to the letter, in particular `recover()` also returns a value in functions that aren't directly called by `defer` (meaning, it returns a value inside a function that is called by a deferred function). In practice, this happens very rarely. This inconsistency should eventually be fixed.
  * Runtime panics can currently not be recovered from. This includes things like divide-by-zero and nil pointer dereferences, which are used in some standard library tests.
