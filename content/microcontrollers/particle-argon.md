---
title: "Particle Argon"
weight: 3
---

[The Particle Argon](https://docs.particle.io/datasheets/wi-fi/argon-datasheet/) is a powerful Wi-Fi enabled development board. It is based on the Nordic [nRF52840](https://www.nordicsemi.com/eng/Products/nRF52840) and ESP32 2.4 GHz Wi-Fi coprocessor.

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

[Documentation for the machine package for the Particle Argon](../machine/particle-argon)

## Flashing

### OpenOCD

Programs are loaded onto the Particle Argon board using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Particle Argon board with your TinyGo code.

- Plug your [Particle Debugger](https://store.particle.io/collections/accessories/products/particle-debugger) to Argon's debugging connector
- Build and flash your TinyGo program using `tinygo flash -target=particle-argon`

## Notes

You can use the USB port to the Particle Argon as a serial port. `UART0` refers to this connection.

Bluetooth support is in development but not yet completed.
