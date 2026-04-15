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

Programs will need to be flashed onto a GBA reproduction cartridge.
You can purchase reproduction cartridges from a site like [inside gadgets](https://shop.insidegadgets.com/?post_type=product&filter_flash-carts=gba), though there are other retailers that carry them as well.

There are a couple pieces of hardware that can be used to flash your games onto a reproduction cartridge.
* [GBxCart RW](https://www.gbxcart.com/)
* [Joey Jr](https://bennvenn.myshopify.com/collections/game-cart-to-pc-interface/products/usb-gb-c-cart-dumper-the-joey-jr)

Both of these tools use [FlashGBX](https://github.com/lesserkuma/FlashGBX/tree/4.2) for flashing cartridges.
Once you have installed FlashGBX, insert a reproduction cart into your hardware flasher of choice and plug it into your computer.
FlashGBX has a GUI you can then use to write your program onto the cart.
