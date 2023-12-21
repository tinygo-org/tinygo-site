---
title: "Game Boy Advance"
weight: 3
---

The [Game Boy Advance](https://en.wikipedia.org/wiki/Game_Boy_Advance) is a handheld videogame platform based on the [ARM7TDMI](http://infocenter.arm.com/help/topic/com.arm.doc.ddi0210c/DDI0210B.pdf) microcontroller.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | ?   | ?   |
| UART      | ?   | ?   |
| SPI       | ?   | ?   |
| I2C       | ?   | ?   |
| ADC       | ?   | ?   |
| PWM       | ?   | ?   |
| USBDevice | ?   | ?   |

## Machine Package Docs

[Documentation for the machine package for the Gameboy Advance](../machine/gameboy-advance)

## Installing dependencies

You can use a Game Boy Advance software emulator such as MGBA (https://mgba.io) to test your programs.


## Building code

Build your Game Boy Advance programs using `-target gameboy-advance` like this:

```shell
tinygo build -o main.gba -target gameboy-advance examples/gba-display
```

You can now use the GBA file with your emulator or flash it onto your physical hardware.

## Examples

* [gba-display](https://github.com/tinygo-org/tinygo/blob/release/src/examples/gba-display/gba-display.go)
* [Other GBA examples (Buttons, Gopher, Snake...)](https://github.com/tinygo-org/tinygba/tree/main/examples)

## Flashing

Information needed here...
