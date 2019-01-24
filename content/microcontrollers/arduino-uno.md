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
| SPI      | YES | Not yet |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Flashing

### AVRDude

Programs are loaded onto the Arduino Uno using the `avrdude` command line utility program. You must install this program before you will be able to flash the Arduino Uno board with your TinyGo code.

- Plug your Arduino Uno into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=arduino`
