---
title: "Digispark"
weight: 3
---

The [Digispark](http://digistump.com/products/1) is an [ATtiny85](https://www.microchip.com/wwwproducts/en/ATtiny85) based microcontroller development board similar to the Arduino line, only cheaper, smaller, and a bit less powerful. 

There is very limited support at the moment for this board.

Note: the AVR backend of LLVM is still experimental so you may encounter bugs.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | Not yet |
| SPI      | Requires software | Not yet |
| I2C      | Requires software | Not yet |
| ADC      | YES | YES |
| PWM      | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `P0`              | `PB0`        |                   |
| `P1`              | `PB1`        | `LED`             |
| `P2`              | `PB2`        |                   |
| `P3`              | `PB3`        |                   |
| `P4`              | `PB4`        |                   |
| `P5`              | `PB5`        |                   |

## Machine Package Docs

[Documentation for the machine package for the Digispark](../machine/digispark)

## Flashing

### Micronucleus

Programs are loaded onto the Digispark using the `micronucleus` command line utility program. You must install [micronucleus](https://littlewire.github.io/) before you will be able to flash the Digispark board with your TinyGo code.

- Build and flash your TinyGo program using `tinygo flash -target=digispark`
- Plug your Digispark into your computer's USB port after you start the flashing process. This is because the bootloader is only active for 5 seconds after power up. Micronucleus will show a message when to insert the device.
