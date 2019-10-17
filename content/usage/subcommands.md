---
title: "Subcommands"
weight: 2
---

TinyGo tries to be similar to the main `go` command in usage. It consists of the following main subcommands:

### build
Compile the given program. The output binary is specified using the ``-o``
parameter. The generated file type depends on the extension:

`.o`
Create a relocatable object file. You can use this option if you don't want to use the TinyGo build system or want to do other custom things.

`.ll`
Create textual LLVM IR, after optimization. This is mainly useful for debugging.

`.bc`
Create LLVM bitcode, after optimization. This may be useful for debugging or for linking into other programs using LTO.

`.hex`
Create an [Intel HEX](https://en.wikipedia.org/wiki/Intel_HEX) file to flash it to a microcontroller.

`.bin`
Similar, but create a binary file.

`.wasm`
Compile and link a WebAssembly file.

(all other)
Compile and link the program into a regular executable. For microcontrollers, it is common to use the .elf file extension to indicate a linked ELF file is generated. For Linux, it is common to build binaries with no extension at all.

### run
Run the program, either directly on the host or in an emulated environment (depending on `-target`).

### flash
Flash the program to a microcontroller.

### gdb
Compile the program, optionally flash it to a microcontroller if it is a remote target, and drop into a GDB shell. From here you can break the current program (Ctrl-C), single-step, show a backtrace, etc. A debugger must be specified for your particular target in the target .json file and the required tools (like GDB for your target) must be installed as well.

### clean
Clean the cache directory, normally stored in `$HOME/.cache/tinygo`. This is not normally needed.

### help
Print a short summary of the available commands, plus a list of command flags.

### version
Print the version of the command and the version of the used `$GOROOT`.

### env
Print a list of environment variables that affect TinyGo (as a shell script). If one or more variable names are given as arguments, env prints the value of each on a new line.
