---
title: "Arduino Uno"
weight: 3
---

The [Arduino Uno](https://store.arduino.cc/arduino-uno-rev3) is based on the AVR [ATmega328p](https://www.microchip.com/wwwproducts/en/ATmega328p) microcontroller.

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
| `D0`              | `PD0`        | `UART_RX_PIN`     |                      |
| `D1`              | `PD1`        | `UART_TX_PIN`     |                      |
| `D2`              | `PD2`        |                   |                      |
| `D3`              | `PD3`        |                   | `Timer2` (channel B) |
| `D4`              | `PD4`        |                   |                      |
| `D5`              | `PD5`        |                   | `Timer0` (channel B) |
| `D6`              | `PD6`        |                   | `Timer0` (channel A) |
| `D7`              | `PD7`        |                   |                      |
| `D8`              | `PB0`        |                   |                      |
| `D9`              | `PB1`        |                   | `Timer1` (channel A) |
| `D10`             | `PB2`        |                   | `Timer1` (channel B) |
| `D11`             | `PB3`        |                   | `Timer2` (channel A) |
| `D12`             | `PB4`        |                   |                      |
| `D13`             | `PB5`        | `LED`             |                      |
| `ADC0`            | `PC0`        |                   |                      |
| `ADC1`            | `PC1`        |                   |                      |
| `ADC2`            | `PC2`        |                   |                      |
| `ADC3`            | `PC3`        |                   |                      |
| `ADC4`            | `PC4`        |                   |                      |
| `ADC5`            | `PC5`        |                   |                      |

## Machine Package Docs

[Documentation for the machine package for the Arduino Uno](../machine/arduino)

## Installing dependencies

The Arduino Uno needs a few extra dependencies to work, for example, if you get an error like this:

```text
/bin/sh: 1: avrdude: not found
```

To fix this, see the installation guide for [Linux](../../../../getting-started/install/linux/#avr-eg-arduino-uno-2) and for [macOS](../../../../getting-started/install/macos/#avr-eg-arduino-uno-2).

## Flashing

### AVRDude

Programs are loaded onto the Arduino Uno using the `avrdude` command line utility program. You must install this program before you will be able to flash the Arduino Uno board with your TinyGo code.

- Plug your Arduino Uno into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target arduino /path/to/code`
