---
title: "Adafruit Feather M0"
weight: 3
---

The [Adafruit Feather M0](https://www.adafruit.com/product/3403) is a tiny ARM development board based on the Atmel [ATSAMD21G18](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of SoC.

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

The Feather M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

- Plug your Feather M0 into your computer's USB port.
- Press the "RESET" button on the board two times to get the Feather M0 board ready to receive code.
- The Feather M0 board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Build your TinyGo program to the board in `.uf2` format using the `tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.uf2 -target=feather-m0 [PATH TO YOUR PROGRAM]` command.
- The Feather M0 board should restart and then begin running your program.

### CLI Flashing

Once you have updated your Feather M0 board the first time, after that you should be able to flash it entirely from the command line using the `stty` command like this:

```
stty -F /dev/ttyACM0 1200 hupcl; tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.uf2 -target=feather-m0 [PATH TO YOUR PROGRAM]
```

Replace `/dev/ttyACM0` in the command above with the correct USB port name for your board.

The Feather M0 board should restart and then begin running your program.
