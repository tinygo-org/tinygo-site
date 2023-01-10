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
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | YES |
| USBDevice | NO  | NO  |

## Pins

| Pin               | Hardware pin | Alternative names | PWM                  |
| ----------------- | ------------ | ----------------- | -------------------- |
| `A0`              | `PF0`        | `ADC0`            |                      |
| `A1`              | `PF1`        | `ADC1`            |                      |
| `A2`              | `PF2`        | `ADC2`            |                      |
| `A3`              | `PF3`        | `ADC3`            |                      |
| `A4`              | `PF4`        | `ADC4`            |                      |
| `A5`              | `PF5`        | `ADC5`            |                      |
| `A6`              | `PF6`        | `ADC6`            |                      |
| `A7`              | `PF7`        | `ADC7`            |                      |
| `A8`              | `PK0`        | `ADC8`            |                      |
| `A9`              | `PK1`        | `ADC9`            |                      |
| `A10`             | `PK2`        | `ADC10`           |                      |
| `A11`             | `PK3`        | `ADC11`           |                      |
| `A12`             | `PK4`        | `ADC12`           |                      |
| `A13`             | `PK5`        | `ADC13`           |                      |
| `A14`             | `PK6`        | `ADC14`           |                      |
| `A15`             | `PK7`        | `ADC15`           |                      |
| `D0`              | `PE0`        |                   |                      |
| `D1`              | `PE1`        |                   |                      |
| `D2`              | `PE4`        |                   | `Timer3` (channel B) |
| `D3`              | `PE5`        |                   | `Timer3` (channel C) |
| `D4`              | `PG5`        |                   | `Timer0` (channel B) |
| `D5`              | `PE3`        |                   | `Timer3` (channel A) |
| `D6`              | `PH3`        |                   | `Timer4` (channel A) |
| `D7`              | `PH4`        |                   | `Timer4` (channel B) |
| `D8`              | `PH5`        |                   | `Timer4` (channel C) |
| `D9`              | `PH6`        |                   | `Timer0` (channel B) |
| `D10`             | `PB4`        |                   | `Timer2` (channel A) |
| `D11`             | `PB5`        |                   | `Timer1` (channel A) |
| `D12`             | `PB6`        |                   | `Timer1` (channel B) |
| `D13`             | `PB7`        | `LED`             | `Timer0` (channel A) |
| `D14`             | `PJ1`        |                   |                      |
| `D15`             | `PJ0`        |                   |                      |
| `D16`             | `PH1`        |                   |                      |
| `D17`             | `PH0`        |                   |                      |
| `D18`             | `PD3`        |                   |                      |
| `D19`             | `PD2`        |                   |                      |
| `D20`             | `PD1`        |                   |                      |
| `D21`             | `PD0`        |                   |                      |
| `D22`             | `PA0`        |                   |                      |
| `D23`             | `PA1`        |                   |                      |
| `D24`             | `PA2`        |                   |                      |
| `D25`             | `PA3`        |                   |                      |
| `D26`             | `PA4`        |                   |                      |
| `D27`             | `PA5`        |                   |                      |
| `D28`             | `PA6`        |                   |                      |
| `D29`             | `PA7`        |                   |                      |
| `D30`             | `PC7`        |                   |                      |
| `D31`             | `PC6`        |                   |                      |
| `D32`             | `PC5`        |                   |                      |
| `D33`             | `PC4`        |                   |                      |
| `D34`             | `PC3`        |                   |                      |
| `D35`             | `PC2`        |                   |                      |
| `D36`             | `PC1`        |                   |                      |
| `D37`             | `PC0`        |                   |                      |
| `D38`             | `PD7`        |                   |                      |
| `D39`             | `PG2`        |                   |                      |
| `D40`             | `PG1`        |                   |                      |
| `D41`             | `PG0`        |                   |                      |
| `D42`             | `PL7`        |                   |                      |
| `D43`             | `PL6`        |                   |                      |
| `D44`             | `PL5`        |                   | `Timer5` (channel C) |
| `D45`             | `PL4`        |                   | `Timer5` (channel B) |
| `D46`             | `PL3`        |                   | `Timer5` (channel A) |
| `D47`             | `PL2`        |                   |                      |
| `D48`             | `PL1`        |                   |                      |
| `D49`             | `PL0`        |                   |                      |
| `D50`             | `PB3`        |                   |                      |
| `D51`             | `PB2`        |                   |                      |
| `D52`             | `PB1`        |                   |                      |
| `D53`             | `PB0`        |                   |                      |

## Machine Package Docs

[Documentation for the machine package for the Arduino Mega 1280](../machine/arduino-mega1280)

## Installing dependencies

The Arduino Mega 1280 needs a few extra dependencies to work, for example, if you get an error like this:

```text
/bin/sh: 1: avrdude: not found
```

To fix this, see the installation guide for [Linux](../../../../getting-started/install/linux/#avr-eg-arduino-uno-2) and for [macOS](../../../../getting-started/install/macos/#avr-eg-arduino-uno-2).

## Flashing

### AVRDude

Programs are loaded onto the Arduino Mega 1280 using the `avrdude` command line utility program. You must install this program before you will be able to flash the Arduino Uno board with your TinyGo code.

- Plug your Arduino Uno into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target arduino-mega1280 /path/to/code`
