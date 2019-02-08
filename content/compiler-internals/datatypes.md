---
title: "Datatypes"
weight: 3
---

TinyGo uses a different representation for some data types than standard Go.

### string
A string is encoded as a `{ptr, len}` tuple. The type is actually defined in the runtime as `runtime._string`, in [src/runtime/string.go](https://github.com/tinygo-org/tinygo/blob/master/src/runtime/string.go). That file also contains some compiler intrinsics for dealing with strings and UTF-8.

### slice
A slice is encoded as a `{ptr, len, cap}` tuple. There is no runtime definition of it as slices are a generic type and the pointer type is different for each slice. That said, the bit layout is exactly the same for every slice and generic `copy` and `append` functions are implemented in [src/runtime/slice.go](https://github.com/tinygo-org/tinygo/blob/master/src/runtime/slice.go).

### array
Arrays are simple: they are simply lowered to a LLVM array type.

### complex
Complex numbers are implemented the same way as Clang implements them: as a struct with two `float32` or `float64` elements.

### map
The map type is a very complex type and is implemented as an (incomplete) hashmap. It is defined as `runtime.hashmap` in [src/runtime/hashmap.go](https://github.com/tinygo-org/tinygo/blob/master/src/runtime/hashmap.go). As maps are reference types, they are lowered to a pointer to the aforementioned struct. See for example `runtime.hashmapMake` that is the compiler intrinsic to create a new hashmap.

### interface
An interface is a `{typecode, value}` tuple and is defined as `runtime._interface` in [src/runtime/interface.go](https://github.com/tinygo-org/tinygo/blob/master/src/runtime/interface.go). The typecode is a small integer unique to the type of the value. See interface.go for a detailed description of how typeasserts and interface calls are implemented.

### function value
A function value is a fat function pointer in the form of `{context, function
pointer}` where context is a pointer which may have any value. See [calling
convention]({{<ref "calling-convention.md">}}) for details.

### goroutine
A goroutine is a linked list of [LLVM
coroutines](https://llvm.org/docs/Coroutines.html). Every blocking call will
create a new coroutine and pass itself to the coroutine as a parameter (see
[calling convention]({{<ref "calling-convention.md">}})). The callee then
re-activates the caller once it would otherwise return to the parent.
Non-blocking calls are normal calls, unaware of the fact that they're running on
a particular goroutine. For details, see
[src/runtime/scheduler.go](https://github.com/tinygo-org/tinygo/blob/master/src/runtime/scheduler.go).

This is rather expensive and should be optimized in the future. But the way it works now, a single stack can be used for all goroutines lowering memory consumption.
