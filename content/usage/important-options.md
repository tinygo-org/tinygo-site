---
title: "Important Build Options"
weight: 3
---

There are a few flags to control how binaries are built:

`-o`
Output filename, see the ``build`` command.

`-target`
Select the target you want to use. Leave it empty to compile for the host.
Example targets:

- wasm
WebAssembly target. Creates .wasm files that can be run in a web browser.
- arduino
Compile using the experimental AVR backend to run Go programs on an Arduino Uno.
- microbit
Compile programs for the `BBC micro:bit <https://microbit.org/>`_.
- qemu
Run on a Stellaris LM3S as emulated by QEMU.

This switch also configures the emulator, flash tool and debugger to use so
you don't have to fiddle with those options.

Read [supported targets]() for a list of supported targets.

`-port`
Specify the serial port used for flashing. This is used for the Arduino Uno, which is flashed over a serial port. It defaults to ``/dev/ttyACM0`` as that is the default port on Linux.

`-opt`
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
-`-gc`
Use the specified memory manager.
    `-gc=none`
    Do not use a memory manager at all. This will cause a link error at every place in the program that tries to allocate memory. The primary use case for this is finding such locations.

    `-gc=dumb`
    Only allocate memory, never free it. This is the simplest allocator possible and uses very few resources while being very portable. Also, allocation is very fast. Larger programs will likely need a real garbage collector.

    `-gc=marksweep`
    Simple conservative mark/sweep garbage collector. This collector does not yet work on all platforms. Also, the performance of the collector is highly unpredictable as any allocation may trigger a garbage collection cycle.
