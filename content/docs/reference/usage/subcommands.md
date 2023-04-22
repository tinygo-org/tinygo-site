---
title: "Subcommands"
weight: 2
description: |
  Subcommands like `tinygo flash` and `tinygo version`.
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
Flash the program to a microcontroller. Some common flags are described below.

`-target={name}`: Specifies the type of microcontroller that is used. The `name`
of the microcontroller is given on the individual pages for each board type
listed under [Microcontrollers]({{<ref "../microcontrollers/">}}). For example,
`arduino-nano`, `d1mini`, `xiao`.

`-monitor`: Start the serial monitor (see below) immediately after flashing.
However, some microcontrollers need a split second or two to configure the
serial port after flashing, and using the `-monitor` flag can fail because the
serial monitor starts too quickly. In that case, use the `tinygo monitor`
command explicitly.

### monitor
Start the serial monitor on the serial port that is connected to the
microcontroller. If there is only a single board attached to the host computer,
the default values for various options should be sufficient. In other
situations, particularly if you have multiple microcontrollers attached, some
parameters may need to be overridden using the following flags:

`-port={port}`: If there are multiple microcontroller attached, an error message
will display a list of potential serial ports. The appropriate port can be
specified by this flag. On Linux, the port will be something like `/dev/ttyUSB0`
or `/dev/ttyACM1`. On MacOS, the port will look like `/dev/cu.usbserial-1420`.
On Windows, the port will be something like `COM1` or `COM31`.

`-baudrate={rate}`: The default baud rate is 115200. Boards using the AVR
processor (e.g. [Arduino Nano]({{<ref "../microcontrollers/arduino-nano.md">}}),
[Arduino Mega 2560]({{<ref "../microcontrollers/arduino-mega2560">}})) use 9600
instead.

`-target={name}`: If you have more than one microcontrollers attached, you can
sometimes just specify the target name and let `tinygo monitor` figure out the
port. Sometimes, this does not work and you have to explicitly use the `-port`
flag.

The serial monitor intercepts several control characters for its own use instead
of sending them to the microcontroller:

* Control-C: terminates the `tinygo monitor`
* Control-Z: suspends the `tinygo monitor` and drops back into shell
* Control-\\: terminates the `tinygo monitor` with a stack trace
* Control-S: flow control, suspends output to the console
* Control-Q: flow control, resumes output to the console
* Control-@: thrown away by `tinygo monitor`

**Note**: If you are using `os.Stdin` on the microcontroller, you may find that
a CR character on the host computer (also known as Enter, `^M`, or `\r`) is
transmitted to the microcontroller without conversion, so `os.Stdin` returns a
`\r` character instead of the expected `\n` (also known as `^J`, NL, or LF) to
indicate end-of-line. You may be able to get around this problem by hitting
`Control-J` in `tinygo monitor` to transmit the `\n` end-of-line character.

### gdb
Compile the program, optionally flash it to a microcontroller if it is a remote target, and drop into a GDB shell. From there you can set breakpoints, start the program with `run` or `continue` (`run` for a local program, `continue` for on-chip debugging), single-step, show a backtrace, break and resume the program with Ctrl-C/`continue`, etc. You may need to install extra tools (like `openocd` and `arm-none-eabi-gdb`) to be able to do this. Also, you may need a dedicated debugger to be able to debug certain boards if no debugger is integrated. Some boards (like the BBC micro:bit and most professional evaluation boards) have an integrated debugger.

### clean
Clean the cache directory, normally stored in `$HOME/.cache/tinygo`. This is not normally needed.

### help
Print a short summary of the available commands, plus a list of command flags.

### version
Print the version of the command and the version of the used `$GOROOT`.

### env
Print a list of environment variables that affect TinyGo (as a shell script). If one or more variable names are given as arguments, env prints the value of each on a new line.
