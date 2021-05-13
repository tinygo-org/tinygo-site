---
title: "Arduino Mega 1280"
weight: 3
---

The [Arduino Mega 1280](https://www.arduino.cc/en/Main/arduinoBoardMega/) is based on the AVR [ATmega1280](https://www.microchip.com/wwwproducts/en/ATMEGA1280) microcontroller.

Note: the AVR backend of LLVM is still experimental so you may encounter bugs.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Machine Package Docs

[Documentation for the machine package for the Arduino Mega 1280](../machine/arduino-mega1280)

## Installing dependencies

The Arduino Mega 1280 needs a few extra dependencies to work, for example, if you get an error like this:

```text
/usr/lib/gcc/avr/5.4.0/../../../avr/bin/ld: cannot find -lm
/usr/lib/gcc/avr/5.4.0/../../../avr/bin/ld: cannot find -lc
collect2: error: ld returned 1 exit status
```

Or like this:

```text
/bin/sh: 1: avrdude: not found
```

To fix this, see the installation guide for [Linux](../../getting-started/linux/#avr-arduino) and for [macOS](../../getting-started/macos/#avr-arduino).

## Flashing

### AVRDude

Programs are loaded onto the Arduino Mega 1280 using the `avrdude` command line utility program. You must install this program before you will be able to flash the Arduino Uno board with your TinyGo code.

- Plug your Arduino Uno into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target arduino-mega1280 /path/to/code`
