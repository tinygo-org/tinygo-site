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

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `LED_BLUE`        | `PB15`       |                   |
| `LED_GREEN`       | `PB9`        |                   |
| `LED_RED`         | `PB11`       | `LED`             |
| `BTN1`            | `PA0`        | `BUTTON`          |
| `BTN2`            | `PA1`        |                   |
| `BTN3`            | `PC6`        |                   |
| `SPI0_NSS_PIN`    | `PA4`        |                   |
| `SPI0_SCK_PIN`    | `PA5`        |                   |
| `SPI0_SDO_PIN`    | `PA6`        |                   |
| `SPI0_SDI_PIN`    | `PA7`        |                   |
| `UART1_TX_PIN`    | `PB6`        |                   |
| `UART1_RX_PIN`    | `PB7`        |                   |
| `UART2_RX_PIN`    | `PA3`        | `UART_RX_PIN`     |
| `UART2_TX_PIN`    | `PA2`        | `UART_TX_PIN`     |
| `I2C1_SCL_PIN`    | `PA9`        | `I2C0_SCL_PIN`    |
| `I2C1_SDA_PIN`    | `PA10`       | `I2C0_SDA_PIN`    |
| `I2C2_SCL_PIN`    | `PA12`       |                   |
| `I2C2_SDA_PIN`    | `PA11`       |                   |

## Machine Package Docs

[Documentation for the machine package for the STM32 Nucleo WL55JC](../machine/nucleo-wl55jc)

## Flashing

### OpenOCD

Programs are loaded onto the STM32 Nucleo WL55JC with the onboard STLink hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32 Nucleo WL55JC board with your TinyGo code.

- Plug your STM32 Nucleo WL55JC into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nucleo-wl55jc`
