---
title: "BBC:Microbit"
weight: 3
---

The [BBC micro:bit](https://microbit.org) is a tiny programmable computer designed for learning. It is based on the Nordic Semiconductor [nRF51822](https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF51822) ARM Cortex MO chip.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | Software support | Not yet |

## Flashing

### OpenOCD

Programs are loaded onto the BBC:Microbit using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the BBC:Microbit board with your TinyGo code.

- Plug your Microbit into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=microbit`
