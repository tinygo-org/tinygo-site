---
title: "GDB"
weight: 10
description: |
  Short introduction into debugging with GDB.
---

The debugger used by TinyGo is called [GDB](https://www.gnu.org/software/gdb/). This tutorial assumes you already have GDB installed and know how to enter the GDB shell, but don't know how to use it yet to actually debug something. For information on how to install and access GDB, see [the debugging guide]({{<ref "debugging.md">}}).

For this tutorial we will assume the following:

  * You have a BBC micro:bit.
  * You have installed the required software: `openocd` and GDB (either `gdb-multiarch` or `arm-none-eabi-gdb`).

## Enter GDB

Now you can enter the GDB shell using the following command:

```
$ tinygo gdb -target=microbit examples/microbit-blink
[...]
target halted due to debug-request, current mode: Thread 
xPSR: 0xc1000000 pc: 0x0000071c msp: 0x20000800
(gdb) 
```

The following line is important:

```
target halted due to debug-request, current mode: Thread 
```

It means the target (in this case, the nrf51822 chip on the BBC micro:bit) is paused right at the start of the program.

## Pause and continue

You can start to run it by running the `c` or `continue` command:

```
(gdb) continue
Continuing.
```

Now you should see that the LED on the BBC micro:bit starts to blink, which indicates that it is running.

Pausing the program is done using Ctrl-C. It doesn't stop like you might expect, it only pauses:

```
(gdb) continue
Continuing.
^C
Program received signal SIGINT, Interrupt.
0x0000085c in runtime.waitForEvents () at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/runtime_nrf_bare.go:8
8		arm.Asm("wfe")
(gdb) 
```

You can see here that the program pauses at the `wfe` instruction, the "wait for event" instruction that is executed to suspend the microcontroller. This is because the blinking LED example spends the vast majority of its time sleeping, and to reduce power consumption it enters a low power state using this instruction.

You can also see that the LED stopped blinking: the processor was paused by the debug probe.

## Backtrace

One very useful command is `backtrace` or `bt`:

```
(gdb) bt
#0  0x0000085c in runtime.waitForEvents () at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/runtime_nrf_bare.go:8
#1  runtime.rtc_sleep (ticks=<optimized out>) at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/runtime_nrf.go:135
#2  runtime.sleepTicks (d=16383) at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/runtime_nrf.go:73
#3  runtime.scheduler () at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/scheduler.go:158
#4  runtime.run () at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/scheduler_any.go:24
#5  Reset_Handler () at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/runtime_nrf.go:27
(gdb) 
```

You can see that the program was somewhere in the scheduler and not in a goroutine.

## Breakpoints

One of the main features of GDB is setting breakpoints. For example, let's say we want to break on printing a string. For that, we can break on the `runtime.printstring` function using the `b` or `break` command:

```
~$ tinygo gdb -target=microbit examples/serial
[...]
target halted due to debug-request, current mode: Thread 
xPSR: 0xc1000000 pc: 0x0000071c msp: 0x20000800
(gdb) b runtime.printstring
Breakpoint 1 at 0x19c: file /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/print.go, line 14.
(gdb) c
Continuing.
Note: automatically using hardware breakpoints for read-only addresses.

Breakpoint 1, runtime.printstring (s=...) at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/print.go:14
14			putchar(s[i])
(gdb) bt
#0  runtime.printstring (s=...) at /home/ayke/src/github.com/tinygo-org/tinygo/src/runtime/print.go:14
#1  0x0000096e in main.main () at /home/ayke/src/github.com/tinygo-org/tinygo/src/examples/serial/serial.go:7
(gdb) list
9	}
10	
11	//go:nobounds
12	func printstring(s string) {
13		for i := 0; i < len(s); i++ {
14			putchar(s[i])
15		}
16	}
17	
18	func printuint8(n uint8) {
(gdb) 
```

Here, you can see the following commands:

 1. `b runtime.printstring` to set a breakpoint on the `printstring` function in the runtime package.
 2. `c` to start running the program (it was paused).
 3. The program hits the breakpoint almost immediately.
 4. `bt` shows the backtrace, which stops at the `main.main` function because that's the root function of the goroutine.
 5. `list` shows the code around the breakpoint.
