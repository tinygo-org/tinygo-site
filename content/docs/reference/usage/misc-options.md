---
title: "Misc. Build Options"
weight: 4
---

- `-no-debug`
Disable outputting debug symbols. This can be useful for WebAssembly, as there is no debugger for .wasm files yet and .wasm files are generally served directly. Avoiding debug symbols can have a big impact on generated binary size, reducing them by more than half.
This is not necessary on microcontrollers because debugging symbols are not flashed to the microcontroller. Additionally, you will need it when you use `tinygo gdb`. In general, it is recommended to include debug symbols unless you have a good reason not to.
Note: while there is some support for debug symbols, only line numbers have been implemented so far. That means single-stepping and stacktraces work just fine, but no variables can be inspected.

- `-size`
    Print size (``none``, ``short``, or ``full``) of the output (linked) binary. Note that the calculated size includes RAM reserved for the stack.

    - `none` (default)
        Print nothing.
    
    - `short`
        Print size statistics, roughly like what the ``size`` binutils program would print but with useful flash and RAM columns:
        ```
        code    data     bss |   flash     ram
        5780     144    2132 |    5924    2276
        ```
    
    - `full`
        Try to determine per package how much space is used. Note that these calculations are merely guesses and can somethimes be way off due to various reasons like inlining:
        ```
        code  rodata    data     bss |   flash     ram | package
        876        0       4       0 |     880       4 | (bootstrap)
        38         0       0       0 |      38       0 | device/arm
        0          0       0      66 |       0      66 | machine
        2994     440     124       0 |    3558     124 | main
        948      127       4       1 |    1079       5 | runtime
        4856     567     132      67 |    5555     199 | (sum)
        5780       -     144    2132 |    5924    2276 | (all)
        ```

- `-print-allocs=<regex>`
    Print heap allocations in the program and why they happen. You can use `-print-allocs=.` if you want to see all heap allocations in the program. You can also match the function name, for example you can use `-print-allocs=strconv` to print heap allocations that contain the `strconv` substring (which includes all functions from the strconv standard library package).
