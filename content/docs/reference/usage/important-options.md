---
title: "Important Build Options"
weight: 3
description: |
  Command line flags you will commonly use with TinyGo that often control how the binary is built or flashed to the device.
---

There are a few flags to control how binaries are built:

- `-o`
Output filename, see the ``build`` command.

- `-target`
    Select the target you want to use. Leave it empty to compile for the host. This switch also configures the emulator, flash tool and debugger to use so you don't have to fiddle with those options. Read [supported targets](../../microcontrollers/) for a list of supported targets. Example targets:

  - `-target=wasm`  
    WebAssembly target. Creates .wasm files that can be run in a web browser.

  - `-target=arduino`  
    Compile using the experimental AVR backend to run Go programs on an Arduino Uno.

  - `-target=microbit`  
    Compile programs for the [BBC micro:bit](https://microbit.org/).

  - `-target=cortex-m-qemu`  
    Run on a Stellaris LM3S as emulated by QEMU. This is mostly useful for testing.

- `-port`
Specify the serial port used for flashing. This is used for the Arduino Uno, which is flashed over a serial port. It defaults to ``/dev/ttyACM0`` as that is the default port on Linux.

- `-opt`
Which optimization level to use. Optimization levels roughly follow standard `-O` level options like ``-O2``, ``-Os``, etc. Available optimization levels:

  - `-opt=0`
Disable as much optimizations as possible. There are still a few optimization passes that run to make sure the program executes correctly, but all LLVM passes that can be disabled are disabled.

  - `-opt=1`
Enable only a few optimization passes. In particular, this keeps the inliner disabled. It can be useful when you want to look at the generated IR but want to avoid the noise that is common in non-optimized code.

  - `-opt=2`
A good optimization level for use cases not strongly limited by code size. Provided here for completeness. It enables most optimizations and will likely result in the fastest code.

  - `-opt=s`
Like `-opt=2`, but while being more careful about code size. It provides a balance between performance and code size.

  - `-opt=z` (default)
Like ``-opt=s``, but more aggressive about code size. This pass also reduces the inliner threshold by a large margin. Use this pass if you care a lot about code size.

- `-ocd-output`
Print output of the on-chip debugger tool (like OpenOCD) while in a `tinygo gdb` session. This can be useful to diagnose connection problems.

- `-gc`
Use the specified memory manager. The default is usually the best option, so leave this value at the default unless you know what you're doing. The possible values are:

  - `-gc=none`  
    Do not use a memory manager at all. This will cause a link error at every place in the program that tries to allocate memory. The primary use case for this is finding such locations.

  - `-gc=leaking`  
    Only allocate memory, never free it. This is the simplest allocator possible and uses very few resources while being very portable. Also, allocation is very fast. Larger programs will likely need a real garbage collector.

  - `-gc=conservative`  
    Simple conservative mark/sweep garbage collector. This collector works on all platforms. Also, the performance of the collector is highly unpredictable as any allocation may trigger a garbage collection cycle.

- `-panic`
Use the specified panic strategy. That is, what the compiled program should do when a panic occurs.

  - `-panic=abort`
    Print the panic message and abort the program. This is the default. On a desktop system this results in a call to [`abort`](https://manpages.debian.org/stretch/manpages-dev/abort.3.en.html). On WebAssembly this is implemented as the [`unreachable`](https://webassembly.github.io/spec/core/syntax/instructions.html#syntax-instr-control) instruction. On microcontrollers, it results in a hang (endless loop).

  - `-panic=trap`
    Do not print the panic message but instead of printing anything, it directly hits a trap instruction. This instruction varies by platform but it will result in the immediate termination of the program. It could either exit with `SIGILL` or cause a call to the `HardFault_Handler`. It can be used to reduce the size of the compiled program while keeping standard Go safety rules intact at the cost of debuggability.

- `-serial` Use the specified serial output as the default stdout. This applies to various runtime outputs such as `println` and `panic`, but also to `os.Stdout` (and therefore `fmt.Println`).

  The default varies by board, but it's usually `-serial=usb` on many newer boards (those with a rp2040, samd21, samd51, or nrf52 chip). Older boards like the Arduino Uno typically use `-serial=uart`, and very small chips like the attiny1616 will use `-serial=none` to disable serial output entirely.

  To see the output, you can use the `-monitor` flag to `tinygo flash`, or use `tinygo monitor`.

  - `-serial=uart` Use the default UART as set in `machine.DefaultUART`. Which UART this is varies by board, but it's typically the first UART indicated on the silkscreen (with TX/RX markings). You may need an external USB-UART connector to receive output this way.

  - `-serial=usb` Use USB-CDC as the serial transport. This will usually not need any extra cables: chips with USB-CDC support are usually also programmed this way.

  - `-serial=rtt` Use [Segger RTT](https://wiki.segger.com/index.php?title=RTT). This only works if you are connected to the chip using a debug probe (most debug probes will work, not just the one from Segger). There are a few boards with a built-in debugger where it might work out of the box, but in most cases you will need an [external debugger]({{< ref "/docs/guides/debugging" >}}).

  - `-serial=none` Disable serial output entirely. This can be useful if you want to use the default serial output for something else. In some cases, disabling serial output will reduce current consumption and/or reduce binary size.

- `-scheduler`
Use the specified scheduler. The default scheduler varies by platform. For example, `AVR` currently defaults to `none` because it has such limited memory while `asyncify` and `tasks` are used for other platforms. Normally you do not need to override the default except on AVR where you can optionally select the tasks scheduler if you want concurrency.

  - `scheduler=tasks` The tasks scheduler is a scheduler much like an RTOS available for non-WASM platforms. This is usually the preferred scheduler.
  - `scheduler=asyncify` The asyncify scheduler is a scheduler for WASM based off of [Binaryen's Asyncify Pass](https://github.com/WebAssembly/binaryen/blob/main/src/passes/Asyncify.cpp).
  - `scheduler=none` The none scheduler disables scheduler support, which means that goroutines and channels are not available. It can be used to reduce firmware size and RAM consumption if goroutines and channels are not needed.

- `-ldflags`
The TinyGo `ldflags` flag has the same behavior as the Go link tool `ldflags` flag. It passes a value along to the TinyGo linker. This is commonly used to set the value for a variable at the time that the program is compiled for example `-ldflags="-X 'package_path.variable_name=new_value'"`
