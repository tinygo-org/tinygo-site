---
title: "What is TinyGo exactly?"
weight: 2
---

A new compiler and a new runtime implementation.

Specifically:

* A new compiler using (mostly) the standard library to parse Go programs and using LLVM to optimize the code and generate machine code for the target architecture.

* A new runtime library that implements some compiler intrinsics, like a memory allocator, a scheduler, and operations on strings. Also, some packages that are strongly connected to the runtime like the `sync` package and the `reflect` package have been or will be re-implemented for use with this new compiler.
