---
title: 'ST Micro "Nucleo" L552ZE'
weight: 3
---

The [STM32 Nucleo L552ZE](https://www.st.com/en/evaluation-tools/nucleo-l552ze-q.html) is an ARM development board based on the ST Micro [STM32L552ZE](https://www.st.com/en/microcontrollers-microprocessors/stm32l552ze.html) SoC.

It has onboard ethernet, 2 user buttons, and 3 user LEDs.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the STM32 Nucleo l552ze](../machine/stm32nucleol552ze)

## Flashing

### OpenOCD

Programs are loaded onto the STM32 Nucleo L552ZE with the onboard [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32 Nucleo L552ZE board with your TinyGo code.

- Plug your STM32 Nucleo L552ZE into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nucleo-l552ze`
