---
title: "Why Go instead of Rust?"
weight: 4
---

Rust is another "new" and safer language that is now made ready for embedded processors. There is [a fairly active community around it](https://rust-embedded.github.io/blog/).

However, apart from personal language preference, Go has a few advantages:

* Subjective, but in general Go is [easier to learn](https://matthias-endler.de/2017/go-vs-rust/). Rust is in general far more complicated than Go, with difficult-to-grasp ownership rules, traits, generics, etc. Go prides itself on being a simple and slightly dumb language, sacrificing some expressiveness for readability.

* Built-in support for concurrency with goroutines and channels that do not rely on a particular implementation threads. This avoids the need for a custom [RTOS-like framework](https://blog.japaric.io/rtfm-v2/) or a [full-blown RTOS](https://github.com/rust-embedded/wg/issues/45) with the associated API one has to learn. In Go, everything is handled by goroutines which are built into the language itself.

* A batteries-included standard library that consists of loosely-coupled packages. Rust uses a monolithic standard library that is currently unusable on bare-metal, while the Go standard library is much more loosely coupled so is more likely to be (partially) supported. Also, non-standard packages in Go do not have to be marked with something like `#![no_std]` to be usable on bare metal. Note: most standard library packages cannot yet be compiled, but this situation will hopefully improve in the future.

At the same time, Rust has other advantages:

* Unlike Go, Rust does not have a garbage collector by default and carefully written Rust code can avoid most or all uses of the heap. Go relies heavily on garbage collection and often implicitly allocates memory on the heap.

* Rust has stronger memory-safety guarantees.

* In general, Rust is more low-level and easier to support on a microcontroller. Of course, this doesn't mean one shouldn't try to run Go on a microcontroller, just that it is more difficult. When even dynamic languages like [Python](https://micropython.org/), [Lua](https://nodemcu.readthedocs.io/en/master) and [JavaScript](https://www.espruino.com/) can run on a microcontroller, then certainly Go can.
