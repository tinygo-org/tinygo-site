---
title: 'ST Micro "Nucleo" L432KC'
weight: 3
---

The [STM32 Nucleo L432KC](https://www.st.com/en/evaluation-tools/nucleo-l432kc.html) is an ARM development board based on the ST Micro [STM32L432KC](https://www.st.com/en/microcontrollers-microprocessors/stm32l432kc.html) SoC.

It has 2 user buttons, and 3 user LEDs.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |
| USBDevice | NO  | NO  |

## Machine Package Docs

[Documentation for the machine package for the STM32 Nucleo L432KC](../machine/nucleo-l432kc)

## Flashing

### OpenOCD

Programs are loaded onto the STM32 Nucleo L432KC with the onboard [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32 Nucleo L432KC board with your TinyGo code.

- Plug your STM32 Nucleo L432KC into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nucleo-l432kc`
