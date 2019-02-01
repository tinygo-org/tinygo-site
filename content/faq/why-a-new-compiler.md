---
title: "Why a new compiler?"
weight: 3
---


Why not modify the existing compiler to produce binaries for microcontrollers?

There are several reasons for this:

* The standard Go compiler (`gc`) does not support instruction sets as used on microcontrollers:

 * The Thumb instruction set is unsupported, but it should be possible to add support for it as it already has an ARM backend.
 * The AVR instruction set (as used in the Arduino Uno) is unsupported and unlikely to be ever supported.

Of course, it is possible to use ``gccgo``, but that has different problems (see below).

* The runtime is really big. A standard 'hello world' on a desktop PC produces a binary of about 1MB, even when using the builtin ``println`` function and nothing else. All this overhead is due to the runtime. Of course, it may be possible to use a different runtime with the same compiler but that will be kind of painful as the exact ABI as used by the compiler has to be matched, limiting optimization opportunities (see below).

* The compiler is optimized for speed, not for code size or memory consumption (which are usually far more important on MCUs). This results in design choices like allocating memory on every value → interface conversion while TinyGo sacrifices some performance for reduced GC pressure.

* With the existing Go libraries for parsing Go code and the pretty awesome LLVM optimizer/backend it is relatively easy to get simple Go programs working with a very small binary size. Extra features can be added where needed in a pay-as-you-go manner similar to C++ avoiding their cost when unused. Most programs on microcontrollers are relatively small so a not-complete compiler is still useful.

* The standard Go compilers do not allocate global variables as static data, but as zero-initialized data that is initialized during program startup. This is not a big deal on desktop computers but prevents allocating these values in flash on microcontrollers. Part of this is due to how the [language specification defines package initialization](https://golang.org/ref/spec#Package_initialization), but this can be worked around to a large extent.

* The standard Go compilers do a few special things for CGo calls. This is necessary because only Go code can use the (small) Go stack while C code will need a much bigger stack. A new compiler can avoid this limitation if it ensures stacks are big enough for C, greatly reducing the C ↔ Go calling overhead.

[At one point](https://github.com/tinygo-org/tinygo-gccgo), a real Go compiler had been used to produce binaries for various platforms, and the result was painful enough to start writing a new compiler:

* The ABI was fixed, so could not be optimized for speed. Also, the ABI didn't seem to be documented anywhere.

* Working arount limitations in the `go` toolchain was rather burdensome and quite a big hack.

* The binaries produced were quite bloated, for various reasons:

 * The Go calling convention places all arguments on the stack. Due to this, stack usage was really bad and code size was bigger than it needed to be.

 * Global initialization was very inefficient, see above.

 * There seemed to be no way to optimize across packages.
