---
title: "Adafruit Feather STM32F405 Express"
weight: 3
---

The [Adafruit Feather STM32F405](https://www.adafruit.com/product/4382) is a tiny ARM development board based on the ST Micro [STM32F405](https://www.st.com/resource/en/datasheet/dm00037051.pdf) family of microcontrollers.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Feather STM32F405](../machine/feather-stm32f405)

## Flashing

Flashing this board using its DFU bootloader can be a bit cumbersome (see Adafruit docs at https://learn.adafruit.com/adafruit-stm32f405-feather-express/dfu-bootloader-details), but it is possible - without requiring an external programmer - by pulling the B0 pin high (you can use the board's 3.3V output pin) at bootup when connected to host PC via USB. This puts the device in bootloader mode.

Once in bootloader mode, the device can be programmed using the open-source tool `dfu-util`.

### CLI Flashing on Linux

You must first install the `dfu-util` program in order to flash the Adafruit Feather STM32F405 board.

    sudo apt update 
    sudo apt install dfu-util

### CLI Flashing on macOS

You must first install the `dfu-util` program in order to flash the Adafruit Feather STM32F405 board. You can do this using Homebrew on macOS:

    brew install dfu-util

### CLI Flashing on Windows

You must first install the `dfu-util` program in order to flash the Adafruit Feather STM32F405 board.

- Download dfu-util from the website here: http://dfu-util.sourceforge.net/releases/dfu-util-0.9-win64.zip
- Decompress the files to a directory such as `C:\dfu-util`
- Add `C:\dfu-util` to your `PATH`.

### Troubleshooting

If you run into trouble getting dfu-util installed and working on Windows, see the blog post at https://www.hanselman.com/blog/HowToFixDfuutilSTMWinUSBZadigBootloadersAndOtherFirmwareFlashingIssuesOnWindows.aspx

## Notes

Goes here
