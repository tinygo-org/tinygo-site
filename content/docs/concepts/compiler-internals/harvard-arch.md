---
title: "Harvard architectures (AVR)"
weight: 10
---

The AVR architecture is a modified Harvard architecture, which means that flash and RAM live in different address spaces. In practice, this means that any given pointer may either point to RAM or flash, but this is not visible from the pointer itself.

To get TinyGo to work on the Arduino, which uses the AVR architecutre, all global variables (which include string constants!) are marked non-constant and thus are stored in RAM and all pointer dereferences assume that pointers point to RAM. At some point this should be optimized so that obviously constant data is kept in read-only memory but this optimization has not yet been implemented.
