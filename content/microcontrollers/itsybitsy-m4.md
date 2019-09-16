---
title: "Adafruit ItsyBitsy M4"
weight: 3
---

The [Adafruit ItsyBitsy M4](https://www.adafruit.com/product/3727) is very compact ARM development board based on the Atmel [SAMD51](https://www.microchip.com/wwwproducts/en/ATSAMD51G19A) family of SoC.

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

The ItsyBitsy M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

- Plug your ItsyBitsy M4 into your computer's USB port.
- Press the "RESET" button on the board two times to get the ItsyBitsy M4 board ready to receive code.
- The ItsyBitsy M4 board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Build your TinyGo program to the board in `.uf2` format using the `tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.uf2 -target=itsybitsy-m4 [PATH TO YOUR PROGRAM]` command.
- The ItsyBitsy M4 board should restart and then begin running your program.

### CLI Flashing

Once you have updated your ItsyBitsy M4 board the first time, after that you should be able to flash it entirely from the command line using the `stty` command like this:

```
stty -F /dev/ttyACM0 1200 hupcl; tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.uf2 -target=itsybitsy-m4 [PATH TO YOUR PROGRAM]
```

Replace `/dev/ttyACM0` in the command above with the correct USB port name for your board.

The ItsyBitsy M4 board should restart and then begin running your program.
