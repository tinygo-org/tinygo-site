---
title: "Microcontroller Targets"
weight: 5
---

TinyGo was designed to run on microcontrollers, but the Go language wasn't. This means there are a few challenges to writing Go code for microcontrollers.

Microcontrollers have very little RAM and execute code directly from flash. Also, constant globals are generally put in flash whenever possible. The Go language itself heavily relies on garbage collection so care must be taken to avoid dynamic memory allocation.

## ARM

ARM Cortex-M processors are well supported. There is support for multiple chips and the backend appears to be stable. In fact, it uses the same underlying technology (LLVM) as the proprietary ARM compiler for code generation.

## AVR

The AVR backend of LLVM is still experimental so you may encounter bugs.

## Xtensa

Support for the ESP8266/ESP32 chips will take a lot of work if they ever get support. See the [TinyGo FAQ](../../faq/what-about-esp8266-esp32) for details.
