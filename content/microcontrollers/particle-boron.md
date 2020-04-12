---
title: "Particle Boron"
weight: 3
---

[The Particle Boron](https://docs.particle.io/datasheets/cellular/boron-datasheet/) The Boron is a powerful LTE CAT-M1/2G/3G enabled development kit that supports cellular networks and Bluetooth LE (BLE). It is based on the Nordic [nRF52840](https://www.nordicsemi.com/eng/Products/nRF52840) and u-blox SARA coprocessor.

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

[Documentation for the machine package for the Particle Boron](../machine/particle-boron)

## Flashing

### OpenOCD

Programs are loaded onto the Particle Boron board using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Particle Boron board with your TinyGo code.

- Plug your [Particle Debugger](https://store.particle.io/collections/accessories/products/particle-debugger) to Boron's debugging connector
- Build and flash your TinyGo program using `tinygo flash -target=particle-boron`

## Notes

You can use the USB port to the Particle Boron as a serial port. `UART0` refers to this connection.

Bluetooth support is in development but not yet completed.
