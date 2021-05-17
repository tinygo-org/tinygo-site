---
title: "Debugging"
weight: 10
description: |
  Debug TinyGo programs using GDB.
---

A debugger can show you things that regular print statements cannot show you, especially on microcontrollers. The debugger used by TinyGo is called [GDB](https://www.gnu.org/software/gdb/). Debugging using GDB can sometimes be intimidating, but it is also extremely powerful. For example:

  * A debugger works even when a serial port is not available. This can happen when there is a bug in an interrupt or in initialization code before the serial port is configured.
  * You can print a stack trace, even in a panic.
  * Depending on the optimization level and other factors, you may be able to inspect function parameters and local variables.
  * You can step through the code and see the results on the device as you do so.

By learning just a few commands such as `continue`, `breakpoint`, and `next`, you will be able to start making good use of the GDB command line. See the [GDB tutorial]({{<ref "gdb.md">}}) for an introduction. Here we will show how to get to the GDB prompt in the first place.

## Requirements

For debugging on your own computer (Linux, MacOS) you only really need to have the `gdb` command installed. No other programs are necessary other than the ones you already need for building programs in TinyGo.

For debugging directly on a chip, you need a few more dependencies. Which dependencies you need vary by chip:

  * You need a GDB variant. It depends on the architecture which one you need, but probably you'll need to install `gdb-multiarch` or `arm-none-eabi-gdb`. For other architectures you will need a different debugger (for example, [`riscv64-unknown-elf-gdb`](https://www.sifive.com/software) for RISC-V).
  * You probably need to install OpenOCD. This program is present in most package managers and it is easiest to install it directly from there.
  * Some boards require something other than OpenOCD, for example some Nordic Semiconductor boards require `JLinkGDBServer`.

### Debian or Ubuntu

You can install the most common dependencies through the package manager:

    sudo apt-get install gdb-multiarch openocd

### MacOS

You can install the most common dependencies through Homebrew:

    brew install open-ocd
    brew tap ARMmbed/homebrew-formulae
    brew install arm-none-eabi-gcc

## Connecting a debug probe

Sometimes you need a separate debug probe. Whether this is necessary depends on the board, some have a debug probe built onto the board already:

  * Evaluation boards from chip vendors (such as STMicroelectronics and Nordic Semiconductors) usually have a debug probe built in.
  * Boards from vendors such as Arduino and Adafruit usually don't have a debug probe. One notable exception is the BBC micro:bit, which does have a debug probe.

If the board you want to debug already has a debug probe, you don't need any extra hardware and can go to the next step.

There is a variety of choice when it comes to debug probes, but most debug probes are able to work with most boards. However, some are better supported than others.

Here are some debug probes that are known to work in TinyGo for at least some boards:

  * [SEGGER J-Link Edu Mini](https://www.segger.com/products/debug-probes/j-link/models/j-link-edu-mini/): this debugger can debug pretty much all ARM Cortex-M chips and is generally very reliable. However, it comes with some possible issues: it's entirely closed source, it is only allowed to be used for non-commercial purposes and it doesn't support as many chips as the full version.
  * [Particle Debugger](https://store.particle.io/products/particle-debugger): a [DAPLink](https://armmbed.github.io/DAPLink/) based debugger that can debug practially all ARM Cortex-M chips (like the SEGGER above) and does not have limitations on how it can be used. It is also open source. While this debugger is designed for Particle hardware it can easily handle chips from other vendors.
  * [ST-Link v2](https://www.st.com/en/development-tools/st-link-v2.html): a debugger often included on boards from STMicroelectronics and also sold separately. It is somewhat less powerful than some other debuggers as it is only intended to be used with ST hardware, even though it works with most microcontrollers that support SWD debugging. Note that many online stores sell counterfeit versions of this debugger that may be unreliable.

Which one you should pick depends on availability and price in your area but if you want to be sure, the SEGGER debugger is a commonly used vendor-neutral debugger that supports almost all chips and is very reliable.

Most ARM Cortex-M chips use a protocol called SWD which uses just two data lines, called SWDIO and SWCLK. Sometimes these are available on separate header pins, but in many cases you'll need to solder some wires to test points on the board. Where they are depends on the board but you can often find them online if you google "&lt;boardname> swd pins" or "&lt;boardname> bootloader programming".

Depending on the debug probe you may need to connect some extra wires:

  * `GND` goes to ground. Not always needed if you're powering the probe from the same USB power source as the board.
  * `Vref` needs to be connected to the power source of the chip, often 3.3V.
  * Optionally, some debuggers (such as the Particle debugger) have a built-in UART that can be useful while debugging.

Check the documentation of your debug probe for more information.

## Getting a GDB shell

Now you have all the hardware ready, it's time to actually start the debugger! If you're lucky, TinyGo already supports debugging the board. For example, the following is all you need to debug the BBC micro:bit:

```
$ tinygo gdb -target=microbit examples/microbit-blink
[...]
Reading symbols from /tmp/tinygo158938990/main...done.
Remote debugging using :3333
0x00000000 in __isr_vector ()
target halted due to debug-request, current mode: Thread 
xPSR: 0x21000000 pc: 0x00003baa msp: 0x200007f0
Loading section .text, size 0xb84 lma 0x0
Loading section .tinygo_stacksizes, size 0x4 lma 0xb84
Loading section .data, size 0x4 lma 0xb88
Start address 0x71c, load size 2956
Transfer rate: 4 KB/sec, 985 bytes/write.
target halted due to debug-request, current mode: Thread 
xPSR: 0xc1000000 pc: 0x0000071c msp: 0x20000800
(gdb) 
```

If you get a result like this without error messages, you have successfully entered the GDB shell! To continue running the program, enter the command `continue`. For more information on using GDB, see the [GDB tutorial]({{<ref "gdb.md">}}).

Otherwise, you may need to specify the programmer:

| Programmer        | Flags                   |
| ----------------- | ----------------------- |
| SEGGER J-Link     | `-programmer=jlink`     |
| Particle Debugger | `-programmer=cmsis-dap` |
| ST-Link V2        | `-programmer=stlink-v2` |

For example:

```
$ tinygo gdb -target=arduino-nano33 -programmer=cmsis-dap examples/blinky1
```
