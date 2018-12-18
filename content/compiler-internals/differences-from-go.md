---
title: "Differences from Go"
weight: 2
---

* A whole program is compiled in a single step, without intermediate linking. This makes incremental development much slower for large programs but enables far more optimization opportunities. In the future, an option should be added for incremental compilation during edit-compile-test cycles.
* Interfaces are always represented as a `{typecode, value}` pair. Unlike Go [https://research.swtch.com/interfaces](https://research.swtch.com/interfaces), TinyGo will not precompute a list of function pointers for fast interface method calls. Instead, all interface method calls are looked up where they are used. This may sound expensive, but it avoids memory allocation at interface creation.
* Global variables are computed during compilation whenever possible (unlike Go, which does not have the equivalent of a `.data` section). This is an important optimization for several reasons:

 - Startup time is reduced. This is nice, but not the main reason.
 - Initializing globals by copying the initial data from flash to RAM costs much less flash space as only the actual data needs to be stored, instead of all instructions to initialize these globals.
 - Data can often be statically allocated instead of dynamically allocated at startup.
 - Dead globals are trivially optimized away by LLVM.
 - Constant globals are trivially recognized by LLVM and marked `constant`. This makes sure they can be stored in flash instead of RAM.
 - Global constants are useful for constant propagation and thus for dead code elimination (like an `if` that depends on a global variable).
