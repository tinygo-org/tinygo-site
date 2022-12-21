---
title: "Arduino Leonardo"
weight: 3
---

The [Arduino Leonardo](https://store.arduino.cc/collections/boards/products/arduino-leonardo-with-headers) is based on the AVR [ATmega32u4](https://www.microchip.com/en-us/product/ATmega32u4) microcontroller.

Note: the AVR backend of LLVM is still experimental so you may encounter bugs.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | NO  |
| SPI       | YES | NO  |
| I2C       | YES | NO  |
| ADC       | YES | YES |
| PWM       | YES | NO  |
| USBDevice | YES | NO  |

## Pins

| Pin               | Hardware pin | Alternative names | PWM                  |
| ----------------- | ------------ | ----------------- | -------------------- |
| `D0`              | `PD2`        |                   |                      |
| `D1`              | `PD3`        |                   |                      |
| `D2`              | `PD1`        |                   |                      |
| `D3`              | `PD0`        |                   |                      |
| `D4`              | `PD4`        |                   |                      |
| `D5`              | `PC6`        |                   |                      |
| `D6`              | `PD7`        |                   |                      |
| `D7`              | `PE6`        |                   |                      |
| `D8`              | `PB4`        |                   |                      |
| `D9`              | `PB5`        |                   |                      |
| `D10`             | `PB6`        |                   |                      |
| `D11`             | `PB7`        |                   |                      |
| `D12`             | `PD6`        |                   |                      |
| `D13`             | `PC7`        | `LED`             |                      |
| `ADC0`            | `PF7`        |                   |                      |
| `ADC1`            | `PF6`        |                   |                      |
| `ADC2`            | `PF5`        |                   |                      |
| `ADC3`            | `PF4`        |                   |                      |
| `ADC4`            | `PF1`        |                   |                      |
| `ADC5`            | `PF0`        |                   |                      |

## Machine Package Docs

[Documentation for the machine package for the Arduino Leonardo](../machine/arduino-leonardo)

## Installing dependencies

The Arduino Leonardo needs a few extra dependencies to work, for example, if you get an error like this:

```text
/usr/lib/gcc/avr/5.4.0/../../../avr/bin/ld: cannot find -lm
/usr/lib/gcc/avr/5.4.0/../../../avr/bin/ld: cannot find -lc
collect2: error: ld returned 1 exit status
```

Or like this:

```text
/bin/sh: 1: avrdude: not found
```

To fix this, see the installation guide for [Linux](../../../../getting-started/install/linux/#avr-eg-arduino-uno-2) and for [macOS](../../../../getting-started/install/macos/#avr-eg-arduino-uno-2).

## Flashing

### AVRDude

Programs are loaded onto the Arduino Leonardo using the `avrdude` command line utility program. You must install this program before you will be able to flash the Arduino Leonardo board with your TinyGo code.

- Plug your Arduino Leonardo into your computer's USB port.
- Reset your Arduino Leonardo to get into the bootloader mode.
- Build and flash your TinyGo program using `tinygo flash -target arduino-leonardo /path/to/code`
