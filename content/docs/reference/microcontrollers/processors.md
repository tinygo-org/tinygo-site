---
title: "Processors"
chapter: true
weight: 3
description: |
  Different processors supported by TinyGo.
---

As of early 2026, boards using the following microcontrollers are
well-supported:

* [SAMD21](https://www.microchip.com/en-us/product/ATSAMD21G18) based on the
  ARM Cortex-M0+ processor
    * Some companies (Adafruit) call these boards the "M0".
* [SAMD51](https://www.microchip.com/en-us/product/ATSAMD51N19A) based on the
  ARM Cortex-M4 processor
    * Some companies call these boards the "M4".
* [nRF52840](https://www.nordicsemi.com/Products/nRF52840)
  based on the Arm Cortex-M4F processor
    * Other nRF microcontrollers (e.g.
    [nRF52832](https://www.nordicsemi.com/Products/nRF52832),
    [nRF52833](https://www.nordicsemi.com/Products/nRF52833),
    [nRF51822](https://www.nordicsemi.com/Products/nRF51822)) are less common
    but should work well with TinyGo.
* [RP2040](https://en.wikipedia.org/wiki/RP2040) with dual ARM Cortex-M0+
  processors
    * The Raspberry Pi Pico is a famous example using this, but there are many
      other boards using this microcontroller now.
* [RP2350](https://en.wikipedia.org/wiki/RP2350) with dual ARM Cortex-M33
  processors
    * The Raspberry Pi Pico2 is another well-known board, but there are many
      others.

Boards using the Espressif microcontrollers have become popular in IoT
applications because of their support for WiFi. TinyGo now
supports WiFi on some of these processors with Bluetooth coming soon:

* [ESP32-C3](https://www.espressif.com/en/products/socs/esp32-c3) based on the
  RISC-V processor. WiFi supported.
* [ESP32-S3](https://www.espressif.com/en/products/socs/esp32-s3) based on the
  dual-core XTensa LX7 processor. WiFi supported.
* [ESP8266](https://en.wikipedia.org/wiki/ESP8266) based on the Xtensa LX106
  processor. WiFi not yet supported.
* [ESP32](https://en.wikipedia.org/wiki/ESP32) based on the Xtensa LX6
  processor.  WiFi not yet supported.

The introductory Arduino boards based on the 8-bit AVR processors work
relatively well under TinyGo. But they have limited amounts of flash and static
memory so they support only small applications (e.g. the `fmt` package may
consume too much flash memory, and goroutines may consume too much static
memory):

* [ATmega328P](https://www.microchip.com/en-us/product/ATmega328P), used by
  Arduino Nano, Arduino UNO, etc.
* [ATtiny85](https://www.microchip.com/en-us/product/ATtiny85), used by
  Digispark.

Want to know the details about how it is possible to compile Go for microcontrollers? Check out the [microcontrollers](../../concepts/compiler-internals/microcontrollers/) page in our "Compiler Internals" section.
