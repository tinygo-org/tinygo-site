---
title: "STM32 Nucleo WL55JC"
weight: 3
---

The [STM32 Nucleo WL55JC](https://www.st.com/en/evaluation-tools/nucleo-wl55jc.html) is an ARM development board based on the ST Micro [STM32WL55JCI](https://www.st.com/en/microcontrollers-microprocessors/stm32wl55jc.html) SoC.

It has onboard LoRa®, (G)FSK, (G)MSK, and BPSK as well as 3 user LEDs, 3 user buttons and 1 reset push-button

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

[Documentation for the machine package for the STM32 Nucleo WL55JC](../machine/nucleo-wl55jc)

## Flashing

### OpenOCD

Programs are loaded onto the STM32 Nucleo WL55JC with the onboard STLink hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32 Nucleo WL55JC board with your TinyGo code.

- Plug your STM32 Nucleo WL55JC into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nucleo-wl55jc`
