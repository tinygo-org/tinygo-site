---
title: "Low Power"
linkTitle: "Low Power"
type: "docs"
weight: 35
description: >
  How to consume less power on microcontrollers
---

Microcontrollers are usually rather non-demanding when it comes to power consumption.
Couple hudreds mW is the most you may expect a hobby project would use.

Sometimes, even that is not low enough.
Power your project from a battery, and it'd benefit greatly from consuming less.

Common scenario for such projects is sleeping most of the time, and only rarely briefly wake up to query sensors and execute actions.
The main caveat here is if not configured properly, microcontroller and peripherals continue to consume power while your program "sleeps".

The standard way to pause execution in Go is to use `time.Sleep()` function that is available in TinyGo too. Depending on the chip and effort put into its support, the efficiency of this function may vary. The only thing certain: this is _not worse_ than busy looping until a certain time has passed.  
Still, with such high-level function as `time.Sleep()` it is not always possible to disable everything that consumes power as the decisions may depend on the application, presense of peripherals and the board used.

Most microcontrollers have one or more "low power" modes, also known as "deep sleep" and, sometimes, "dormant".  
In such modes microcontroller may consume as low as milli- or even micro-Watts.  
TinyGo gives you low level access to hardware to try and make most of it.

## Implementations

Manufacturers implement different modes, depending of capabilites of their products and name these modes differently too.
Hence no naming consistency across microcontrollers, unfortunately.

Always consult "Power consumption" part of a respective datasheet to understand capabilities of your hardware.  
When reading datasheets, keep in mind peripherals that draw power too. These include \[power\] LEDs, various sensors or co-processors usually present on development boards.

Capabilities of different chips:
- Some very good chips (like the nRF series) can go to a very low power state (just a few µA) with RAM retention and waking up from a low-power RTC.
- Some limited chips (like the esp32) can't really "sleep" with low power -- they can disable the core and memory and reset after a predefined time losing state stored in memory.
- Other chips need to use a special feature to go into deep sleep mode.

Especially the distinction between waking up with or without RAM is important, because it greatly affects how you write software.

The common way to wake up from deep sleep is to use [interrupts](compiler-internals/interrupts), such as timer, real-time-clock or an interrupt on a pin change.

### RP2040

Minimal achievable power consumption is 0.18mA in "Dormant" and 0.39mA in "Sleep" modes.  
See [RP2040 datasheet](https://datasheets.raspberrypi.com/rp2040/rp2040-datasheet.pdf#page=621), section 5.7 Power Consumption.

### nRF Family

[nRF series chips](https://www.nordicsemi.com/Products) are exceptionally efficient and is recommended choice for low-power projects.  
Even with [BLE broadcasting enabled](https://devzone.nordicsemi.com/power/w/opp/2/online-power-profiler-for-bluetooth-le) it can use less power than RP2040 in its lowest power state!

Minimal achievable power consumption is in the area of µA.

TinyGo implements standard `time.Sleep()` very efficiently for nRF chips so you can just use it and have great results.  
The only thing you may want to tweak is to disable the UART with help of `-serial=none` flag to go even lower than default.

## The Way to Low Power

> Attached [debugger](../guides/debugging) usually blocks the low power state in chip.  
> The best way to measure when using a debugger is to flash the code, do a power cycle, and then measure.

> The changes discussed in this paragraph shall be made in TinyGo and not in your application code.  
> These are `runtime` and/or `machine` packages, in files related to your target.

Following steps shall help you achieve low power consumption:

- Put a `for { arm.Asm("wfi") }` loop before anything else is initialized (at the beginning of [Reset_Handler](https://github.com/tinygo-org/tinygo/search?q=%22export+Reset_Handler%22)).  
  _`wfi` pauses the processor, puts it in "wait for interrupt" state. And when inserted before any initialisation, this ensures no internal subsystem is running consuming power._
- Measure the current, this should be the lowest consumption as advertised for the chip.
- If it is higher than that, investigate why (it could be a power LED, regulator, some clock that's enabled by the bootloader, maybe other things).  
  _Usually this ranges somewhere between 1µA and 100µA, depending on the low-power capabilities of the chip. Don't continue until you really do reach the lowest possible power consumption._
- Slowly move the loop to later in the code, until you move it into the main function.
- There are probably things initialized by default that consume current, like UART, PLLs, different clocks. Disable them before continuing, or investigate why they consume power.
