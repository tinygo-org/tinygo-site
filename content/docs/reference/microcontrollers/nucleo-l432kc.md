---
title: 'ST Micro STM32 Nucleo-32 L432KC'
weight: 3
---

The [STM32 Nucleo-32 L432KC](https://www.st.com/en/evaluation-tools/nucleo-l432kc.html) is an ARM development board based on the ST Micro [STM32L432KC](https://www.st.com/en/microcontrollers-microprocessors/stm32l432kc.html) SoC.

## Peripherals and Drivers

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

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `A0`              | `PA0`        | `BUTTON`          |
| `A1`              | `PA1`        |                   |
| `A2`              | `PA3`        |                   |
| `A3`              | `PA4`        |                   |
| `A4`              | `PA5`        |                   |
| `A5`              | `PA6`        |                   |
| `A6`              | `PA7`        |                   |
| `A7`              | `PA2`        | `UART_TX_PIN`     |
| `D0`              | `PA10`       |                   |
| `D1`              | `PA9`        |                   |
| `D2`              | `PA12`       |                   |
| `D3`              | `PB0`        |                   |
| `D4`              | `PB7`        | `I2C0_SDA_PIN`    |
| `D5`              | `PB6`        | `I2C0_SCL_PIN`    |
| `D6`              | `PB1`        |                   |
| `D7`              | `PC14`       |                   |
| `D8`              | `PC15`       |                   |
| `D9`              | `PA8`        |                   |
| `D10`             | `PA11`       |                   |
| `D11`             | `PB5`        | `SPI1_SDI_PIN`, `SPI0_SDI_PIN` |
| `D12`             | `PB4`        | `SPI1_SDO_PIN`, `SPI0_SDO_PIN` |
| `D13`             | `PB3`        | `LED`, `LED_BUILTIN`, `LED_GREEN`, `SPI1_SCK_PIN`, `SPI0_SCK_PIN` |
| `UART_RX_PIN`     | `PA15`       |                   |

## Machine Package Docs

[Documentation for the machine package for the STM32 Nucleo L432KC](../machine/nucleo-l432kc)

## Flashing

### OpenOCD

Programs are loaded onto the STM32 Nucleo L432KC with the onboard [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32 Nucleo L432KC board with your TinyGo code.

- Plug your STM32 Nucleo L432KC into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nucleo-l432kc`
