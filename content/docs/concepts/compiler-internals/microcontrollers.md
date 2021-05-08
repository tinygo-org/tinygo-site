---
title: "Microcontroller Targets"
weight: 1
---

TinyGo was designed to run on microcontrollers, but the Go language wasn't. This means there are a few challenges to writing Go code for microcontrollers.

Microcontrollers have very little RAM and execute code directly from flash. Also, constant globals are generally put in flash whenever possible. The Go language itself heavily relies on garbage collection so care must be taken to avoid dynamic memory allocation.

## ARM

ARM Cortex-M processors are well supported. There is support for multiple chips and the backend appears to be stable. In fact, it uses the same underlying technology (LLVM) as the proprietary ARM compiler for code generation.

Please note that while Cortex-M is well supported, every chip and chip family needs to have some special support added for things like the memory map, clock configuration, timekeeping (`time.Now` and friends), and some more things. This extra code is required for TinyGo to be usable on a given chip. If you want to add support for a new board, take a look at [this wiki page](https://github.com/tinygo-org/tinygo/wiki/Adding-a-new-board).

## AVR

The AVR backend of LLVM is still experimental so you may encounter bugs. Larger programs often fail to work due to this.

## RISC-V

There is some support for RISC-V, in particular the [HiFive1 rev B](https://github.com/tinygo-org/tinygo/wiki/Adding-a-new-board) board.

## Xtensa

Support to run TinyGo directly on the ESP8266/ESP32 chips now exists, although is still in an early stage. See the [TinyGo FAQ](../../faq/what-about-esp8266-esp32) for details.
