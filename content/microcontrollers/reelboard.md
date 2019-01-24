---
title: "reel board"
weight: 3
---

The [reel board](https://www.phytec.eu/product-eu/internet-of-things/reelboard/) is an evaluation board based on the Nordic Semiconductor [nRF52840](https://www.nordicsemi.com/eng/Products/nRF52840) SoC.

It is equipped with an Electrophoretic (electronic ink) Display (EPD), along with temperature, humidity, light, and accelerometer sensors, and Bluetooth connectivity.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |

## Flashing

### OpenOCD

Programs are loaded onto the reelboard using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the reelboard with your TinyGo code.

- Plug your reelboard into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=reelboard`
