---
title: "STM32F4 Discovery"
weight: 3
---

The [STM32F4 Discovery](https://www.st.com/en/evaluation-tools/stm32f4discovery.html) is an ARM development board based on the ST Micro [STM32F407](https://www.st.com/resource/en/datasheet/stm32f407vg.pdf) SoC.

It has an onboard LIS302DL or LIS3DSH accelerometer, MP45DT02 MEMS digital microphone, an
CS43L22 audio DAC, 2 user buttons, and 4 user LEDs.

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

[Documentation for the machine package for the STM32F4 Discovery](../machine/stm32f4disco)

## Flashing

### OpenOCD

Programs are loaded onto the STM32F4 Discovery with the onboard [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32F4 Discovery board with your TinyGo code.

- Plug your STM32F4 Discovery into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=stm32f4disco`
