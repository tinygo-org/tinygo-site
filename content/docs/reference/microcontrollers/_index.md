---
title: "Microcontrollers"
chapter: true
weight: 3
description: |
  Documentation for each microcontroller board supported by TinyGo.
---

TinyGo lets you run Go directly on microcontrollers.

TinyGo has support for 85 different boards and devices such as the Arduino Nano33 IoT, Adafruit Circuit Playground Express, BBC micro:bit and more. Click on a board name below to see the what features are supported for the given hardware.

Support for some boards and processor types are more complete than others.
As of early 2023, boards using the following microcontrollers are
well-supported:

* [SAMD21](https://www.microchip.com/en-us/product/ATSAMD21G18) based on the
  ARM Cortex-M0+ processor
* [SAMD51](https://www.microchip.com/en-us/product/ATSAMD51N19A) based on the
  ARM Cortex-M4 processor
* [nRF52840](https://infocenter.nordicsemi.com/topic/ps_nrf52840/keyfeatures_html5.html)
  based on the Arm Cortex-M4F processor
* [RP2040](https://en.wikipedia.org/wiki/RP2040) with dual ARM Cortex-M0+
  processors

The introductory Arduino boards based on the 8-bit AVR processors work
relatively well under TinyGo. But they have limited amounts of flash and static
memory so they support only small applications:

* [ATmega328P](https://www.microchip.com/en-us/product/ATmega328P), used by
  Arduino Nano, Arduino UNO, etc.

Boards using the Espressif microcontrollers have become popular in IoT
applications because of their support for WiFi. Unfortunately TinyGo does not
support WiFi nor Bluetooth on these boards:

* [ESP8266](https://en.wikipedia.org/wiki/ESP8266) based on the Xtensa LX106
  processor
* [ESP32](https://en.wikipedia.org/wiki/ESP32) based on the Xtensa LX6
  processor
* [ESP32-C3](https://www.espressif.com/en/products/socs/esp32-c3) based on the
  RISC-V processor

We also give you the ability to add new boards. If your target isn't listed here, please raise an issue in the [issue tracker](https://github.com/tinygo-org/tinygo/issues).

Want to know the details about how it is possible to compile Go for microcontrollers? Check out the [microcontrollers](../../concepts/compiler-internals/microcontrollers/) page in our "Compiler Internals" section.
