---
title: "Requirements"
weight: 1
---

These are the base requirements and enough for most (desktop) use.

- Go 1.11+
- LLVM 7 (for example, from [apt.llvm.org](http://apt.llvm.org/)

Linking a binary needs an installed C compiler (`cc`). At the moment it
expects GCC or a recent Clang.

## ARM Cortex-M

The Cortex-M family of microcontrollers is well supported, as it uses the stable
ARM LLVM backend (which is even used in the propietary C compiler from ARM).
Compiling to object code should be supported out of the box, but compiling the
final binary and flashing it needs some extra tools.

- binutils (`arm-none-eabi-ld`, `arm-none-eabi-objcopy`) for linking and for producing .hex files for flashing.
- Clang 7 (`clang-7`) for building assembly files and the compiler runtime library [https://compiler-rt.llvm.org/](https://compiler-rt.llvm.org/).
- The flashing tool for the particular chip, like `openocd` or `nrfjprog`.

## AVR (Arduino)

The AVR backend has similar requirements as the `ARM Cortex-M` backend. It
needs the following tools:

- binutils (`avr-objcopy`) for flashing.
- GCC (`avr-gcc`) for linking object files.
- libc (`avr-libc`), which is not installed on Debian as a dependency of `avr-gcc`.
- `avrdude` for flashing to an Arduino.

## WebAssembly

The WebAssembly backend only needs a special linker from the LLVM project:

- LLVM linker (`ld.lld-7`) for linking WebAssembly files together.
