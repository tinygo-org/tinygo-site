---
title: 'ST Micro STM32F103XX "Bluepill"'
weight: 3
---

The [Bluepill](http://wiki.stm32duino.com/index.php?title=Blue_Pill) is a popular, ultra-cheap and compact ARM development board based on the ST Micro [STM32F103C8](https://www.st.com/en/microcontrollers/stm32f103c8.html) SoC.

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

[Documentation for the machine package for the Bluepill](../machine/bluepill)

## Flashing

### OpenOCD

Programs are loaded onto the Bluepill with a separate hardware programmer such as the [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) board, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Bluepill board with your TinyGo code.

- Plug your STLink v2 programmer into your computer's USB port.
- Plug your Bluepill into the STLink v2 programmer using the Bluepill `SWIO`, `SWCLK`, `3V3` and `GND` pins.
- Build and flash your TinyGo program using `tinygo flash -target=bluepill`
