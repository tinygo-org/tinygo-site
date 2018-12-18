---
title: "What about the ESP8266/ESP32?"
weight: 5
---

These chips use the rather obscure Xtensa instruction set. While a port of GCC exists and Espressif provides precompiled GNU toolchains, there is no support yet in LLVM (although there have been [multiple attempts](http://lists.llvm.org/pipermail/llvm-dev/2018-July/124789.html).

There are two ways these chips might be supported in the future, and both will take a considerable amount of work:

* The compiled LLVM IR can be converted into (ugly) C and then be compiled with a supported C compiler (like GCC for Xtensa). This has been [done before](https://github.com/JuliaComputing/llvm-cbe) so should be doable.

* One of the work-in-progress LLVM backends can be worked on to get it in a usable state. If this is finished, a true TinyGo port is possible.
