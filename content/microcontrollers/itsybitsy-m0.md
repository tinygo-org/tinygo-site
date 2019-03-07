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
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Flashing

### UF2

The ItsyBitsy M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

- Plug your ItsyBitsy M0 into your computer's USB port.
- Press the "RESET" button on the board two times to get the ItsyBitsy M0 board ready to receive code.
- The ItsyBitsy M0 board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Build your TinyGo program to the board in `.uf2` format using the `tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.uf2 -target=itsybitsy-m0 [PATH TO YOUR PROGRAM]` command.
- The ItsyBitsy M0 board should restart and then begin running your program.
