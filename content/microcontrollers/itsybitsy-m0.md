---
title: "ItsyBitsy M0"
weight: 3
---

The [Adafruit ItsyBitsy M0](https://www.adafruit.com/product/3727) is very compact ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of SoC.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | Not yet |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | YES |

## Flashing

### UF2

The ItsyBitsy M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

- Plug your ItsyBitsy M0 into your computer's USB port.
- Build your TinyGo program to `.bin` format using the `tinygo build -target=itsybitsy-m0` command.
- Press the "RESET" button on the board two times.
- Convert/flash the resulting `.bin` using the [`uf2conv.py`](https://github.com/Microsoft/uf2/blob/master/utils/uf2conv.py) program.

### BOSSA

The ItsyBitsy M0 UF2 bootloader also supports the BOSSA format. You must install the latest version of the `bossac` command line utility version 1.9+ from [https://github.com/shumatech/BOSSA](https://github.com/shumatech/BOSSA).

It is slower to flash the board using `bossac` as the dual bootloader takes longer to become active than just using the UF2 format.

- Plug your ItsyBitsy M0 into your computer's USB port.
- Press the "RESET" button on the board two times.
- Build and flash your TinyGo program using `tinygo flash -target=itsybitsy-m0`
