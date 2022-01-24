---
title: "Seeed Studio LoRa-E5 Development Kit"
weight: 3
---

The [LoRa-E5 Development Kit](https://www.seeedstudio.com/LoRa-E5-Dev-Kit-p-4868.html) is an ARM development board based on the ST Micro [STM32WLE5JC](https://www.st.com/en/microcontrollers-microprocessors/stm32wle5jc.html) SoC.

It has onboard LoRa®, (G)FSK, (G)MSK, and BPSK as well as 1 user LED, 1 user buttons and 1 reset push-button

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

[Documentation for the machine package for the LoRa E5](../machine/lorae5)

## Flashing

### OpenOCD

Programs are loaded onto the LoRa-E5 with an STLink hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the LoRa-E5 board with your TinyGo code.

- Connect an ST-Link device to the SWD/SWIM headers.
- Plug your LoRa-E5 board into your computer's USB port.
- Plug your ST-Link device into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=lorae5`
