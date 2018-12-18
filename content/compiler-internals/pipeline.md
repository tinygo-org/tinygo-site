---
title: "Pipeline"
weight: 5
---

Like most compilers, TinyGo is a compiler built as a pipeline of transformations, that each translate an input to a simpler output version (also called lowering). However, most of these part are not in TinyGo itself. The frontend is mostly implemented by external Go libraries, and most optimizations and code generation is implemented by LLVM.

This is roughly the pipeline for TinyGo:

* Lexing, parsing, typechecking and [AST](https://en.wikipedia.org/wiki/Abstract_syntax_tree) building is done by
packages in the [standard library](https://godoc.org/go) and in the [golang.org/x/tools/go library] (https://godoc.org/golang.org/x/tools/go).
* [SSA](https://en.wikipedia.org/wiki/Static_single_assignment_form) construction (a very important step) is done by the [golang.org/x/tools/go/ssa](https://godoc.org/golang.org/x/tools/go/ssa) package.
* This SSA form is then analyzed by the [ir package](https://godoc.org/github.com/aykevl/tinygo/ir) to learn all kinds of things about the code that help the optimizer.
* The Go SSA is then transformed into LLVM IR by the [compiler package](https://godoc.org/github.com/aykevl/tinygo/compiler). Both forms are SSA, but because Go SSA is higher level and contains Go-specific constructs (like interfaces and goroutines) this is non-trivial. However, the vast majority of the work is simply lowering the available Go SSA into LLVM IR, possibly calling some runtime library intrinsics in the process (for example, operations on maps).
* This LLVM IR is then optimized by the LLVM optimizer, which has a large array of standard [optimization passes]
(https://llvm.org/docs/Passes.html). Currently, the standard optimization pipeline is used as is also be used by Clang, but a pipeline better tuned for TinyGo might be used in the future.
* After all optimizations have run, a few fixups are needed for AVR for globals. This is implemented by the compiler package.
* Finally, the resulting machine code is emitted by LLVM to an object file, to be linked by an architecture-specific linker in a later step.

After this whole list of compiler phases, the Go source has been transformed into object code. It can then be emitted directly to a file (for linking in a different build system), or it can be linked directly or even be flashed to a target by TinyGo (using external tools under the hood).
