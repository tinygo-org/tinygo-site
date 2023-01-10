---
title: "Arduino Mega 2560"
weight: 3
---

The [Arduino Mega 2560](https://store.arduino.cc/arduino-mega-2560-rev3) is based on the AVR [ATmega2560](https://www.microchip.com/wwwproducts/en/ATmega2560) microcontroller.

Note: the AVR backend of LLVM is still experimental so you may encounter bugs.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | Not yet |
| USBDevice | NO  | NO  |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `A0`              | `PF0`        | `ADC0`            |
| `A1`              | `PF1`        | `ADC1`            |
| `A2`              | `PF2`        | `ADC2`            |
| `A3`              | `PF3`        | `ADC3`            |
| `A4`              | `PF4`        | `ADC4`            |
| `A5`              | `PF5`        | `ADC5`            |
| `A6`              | `PF6`        | `ADC6`            |
| `A7`              | `PF7`        | `ADC7`            |
| `A8`              | `PK0`        | `ADC8`            |
| `A9`              | `PK1`        | `ADC9`            |
| `A10`             | `PK2`        | `ADC10`           |
| `A11`             | `PK3`        | `ADC11`           |
| `A12`             | `PK4`        | `ADC12`           |
| `A13`             | `PK5`        | `ADC13`           |
| `A14`             | `PK6`        | `ADC14`           |
| `A15`             | `PK7`        | `ADC15`           |
| `D0`              | `PE0`        | `UART_RX_PIN`, `UART0_RX_PIN` |
| `D1`              | `PE1`        | `UART_TX_PIN`, `UART0_TX_PIN` |
| `D2`              | `PE4`        |                   |
| `D3`              | `PE5`        |                   |
| `D4`              | `PG5`        |                   |
| `D5`              | `PE3`        |                   |
| `D6`              | `PH3`        |                   |
| `D7`              | `PH4`        |                   |
| `D8`              | `PH5`        |                   |
| `D9`              | `PH6`        |                   |
| `D10`             | `PB4`        |                   |
| `D11`             | `PB5`        |                   |
| `D12`             | `PB6`        |                   |
| `D13`             | `PB7`        | `LED`             |
| `D14`             | `PJ1`        | `UART3_TX_PIN`    |
| `D15`             | `PJ0`        | `UART3_RX_PIN`    |
| `D16`             | `PH1`        | `UART2_TX_PIN`    |
| `D17`             | `PH0`        | `UART2_RX_PIN`    |
| `D18`             | `PD3`        | `UART1_TX_PIN`    |
| `D19`             | `PD2`        | `UART1_RX_PIN`    |
| `D20`             | `PD1`        |                   |
| `D21`             | `PD0`        |                   |
| `D22`             | `PA0`        |                   |
| `D23`             | `PA1`        |                   |
| `D24`             | `PA2`        |                   |
| `D25`             | `PA3`        |                   |
| `D26`             | `PA4`        |                   |
| `D27`             | `PA5`        |                   |
| `D28`             | `PA6`        |                   |
| `D29`             | `PA7`        |                   |
| `D30`             | `PC7`        |                   |
| `D31`             | `PC6`        |                   |
| `D32`             | `PC5`        |                   |
| `D33`             | `PC4`        |                   |
| `D34`             | `PC3`        |                   |
| `D35`             | `PC2`        |                   |
| `D36`             | `PC1`        |                   |
| `D37`             | `PC0`        |                   |
| `D38`             | `PD7`        |                   |
| `D39`             | `PG2`        |                   |
| `D40`             | `PG1`        |                   |
| `D41`             | `PG0`        |                   |
| `D42`             | `PL7`        |                   |
| `D43`             | `PL6`        |                   |
| `D44`             | `PL5`        |                   |
| `D45`             | `PL4`        |                   |
| `D46`             | `PL3`        |                   |
| `D47`             | `PL2`        |                   |
| `D48`             | `PL1`        |                   |
| `D49`             | `PL0`        |                   |
| `D50`             | `PB3`        |                   |
| `D51`             | `PB2`        |                   |
| `D52`             | `PB1`        |                   |
| `D53`             | `PB0`        |                   |

## Machine Package Docs

[Documentation for the machine package for the Arduino Mega 2560](../machine/arduino-mega2560)

## Installing dependencies

The Arduino Mega 2560 needs a few extra dependencies to work, for example, if you get an error like this:

```text
/bin/sh: 1: avrdude: not found
```

To fix this, see the installation guide for [Linux](../../../../getting-started/install/linux/#avr-eg-arduino-uno-2) and for [macOS](../../../../getting-started/install/macos/#avr-eg-arduino-uno-2).

## Flashing

### AVRDude

Programs are loaded onto the Arduino Mega 2560 using the `avrdude` command line utility program. You must install this program before you will be able to flash the Arduino Uno board with your TinyGo code.

- Plug your Arduino Uno into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target arduino-mega2560 /path/to/code`
