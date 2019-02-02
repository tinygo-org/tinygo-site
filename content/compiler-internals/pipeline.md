---
title: "Pipeline"
weight: 5
---

Like most compilers, TinyGo is a compiler built as a pipeline of transformations
that each translate an input to a simpler output version (also called lowering).
However, most of these part are not in TinyGo itself. The frontend is mostly
implemented by external Go libraries, and most optimizations and code generation
is implemented by LLVM.

This is roughly the pipeline for TinyGo:

  * Lexing, parsing, and typechecking is done by packages in the
    [standard library](https://godoc.org/go) and in the
    [golang.org/x/tools/go library](https://godoc.org/golang.org/x/tools/go).
  * [SSA](https://en.wikipedia.org/wiki/Static_single_assignment_form)
    construction (a very important step) is done by the
    [golang.org/x/tools/go/ssa](https://godoc.org/golang.org/x/tools/go/ssa)
    package.
  * The Go SSA is then transformed into LLVM IR by the
    [compiler package](https://godoc.org/github.com/tinygo-org/tinygo/compiler).
    Both forms are SSA, but because Go SSA is higher level and contains
    Go-specific constructs (like interfaces and goroutines) this is non-trivial.
    However, the vast majority of the work is simply lowering the available Go
    SSA into LLVM IR, possibly calling some runtime library intrinsics in the
    process (for example, operations on maps).
  * Go does a lot of initialization at runtime, which is really bad for
    code size. This includes all global variables: they are all initialized at
    runtime, not at compile time like C. So TinyGo
    [interprets these functions at compile time](https://github.com/tinygo-org/tinygo/tree/master/interp)
    as far as it is able to.
  * The resulting IR is then first optimized by a mixture of handpicked LLVM
    optimization passes, TinyGo-specific
    optimizations (escape analysis, `string`-to-`[]byte` optimizations, etc.)
    and custom lowering. For example, this is the time when interfaces are
    lowered to their final form to benefit from the optimizations already done
    until that point.
  * This LLVM IR is then optimized by the LLVM optimizer, which has a large
    array of standard [optimization passes](https://llvm.org/docs/Passes.html).
    This is the standard optimization pipeline as is also used by Clang.
  * After all optimizations have run, a few fixups are needed for AVR for
    globals, because AVR has separate address spaces for flash and RAM. This is
    implemented by the compiler package.
  * Finally, the resulting machine code is emitted by LLVM to an object file.

This is just the compiler. TinyGo can also automatically link the file and flash
it to a device, if needed.
