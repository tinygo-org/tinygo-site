---
title: "Particle Xenon"
weight: 3
---

[The Particle Xenon](https://docs.particle.io/datasheets/discontinued/xenon-datasheet/) is a low cost mesh-enabled development board.

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

[Documentation for the machine package for the Particle Xenon](../machine/particle-xenon)

## Flashing

### OpenOCD

Programs are loaded onto the Particle Xenon board using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Particle Xenon board with your TinyGo code.

- Plug your [Particle Debugger](https://store.particle.io/collections/accessories/products/particle-debugger) to Xenon's debugging connector
- Build and flash your TinyGo program using `tinygo flash -target=particle-xenon`

## Notes

You can use the USB port to the Particle Xenon as a serial port. `UART0` refers to this connection.

Bluetooth support is in development but not yet completed.
