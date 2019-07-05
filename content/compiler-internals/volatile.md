---
title: "The volatile keyword"
weight: 3
---

Go does not have the `volatile` keyword like C/C++. Volatile is used to read and write memory mapped registers that can change or have side effects that the compiler does not know about.

To implement volatile operations, the runtime/volatile package has been added with functions that are treated specially by the compiler. For example, you can use `volatile.LoadUint32` to load an `*uint32`, a type that is often used in memory mapped peripherals on 32-bit microcontrollers.

In practice, you do not normally need to use these functions. Memory mapped registers are available with special `Get` and `Set` functions that call these volatile operations behind the scenes. One exception is when you want to communicate between normal code and interrupt handlers, in which case you need to use volatile operations for all variables that are accessed by both the normal code and the interrupt handler.
