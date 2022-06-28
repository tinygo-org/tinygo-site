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

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `C13`             | `PC13`       | `LED`             |
| `C14`             | `PC14`       |                   |
| `C15`             | `PC15`       |                   |
| `A0`              | `PA0`        | `ADC0`, `BUTTON`  |
| `A1`              | `PA1`        | `ADC1`            |
| `A2`              | `PA2`        | `ADC2`            |
| `A3`              | `PA3`        | `ADC3`            |
| `A4`              | `PA4`        | `ADC4`            |
| `A5`              | `PA5`        | `ADC5`, `SPI0_SCK_PIN` |
| `A6`              | `PA6`        | `ADC6`, `SPI0_SDI_PIN` |
| `A7`              | `PA7`        | `ADC7`, `SPI0_SDO_PIN` |
| `B0`              | `PB0`        | `ADC8`            |
| `B1`              | `PB1`        | `ADC9`            |
| `B10`             | `PB10`       |                   |
| `B11`             | `PB11`       |                   |
| `B12`             | `PB12`       |                   |
| `B13`             | `PB13`       |                   |
| `B14`             | `PB14`       |                   |
| `B15`             | `PB15`       |                   |
| `A8`              | `PA8`        |                   |
| `A9`              | `PA9`        | `UART_TX_PIN`     |
| `A10`             | `PA10`       | `UART_RX_PIN`     |
| `A11`             | `PA11`       |                   |
| `A12`             | `PA12`       |                   |
| `A13`             | `PA13`       |                   |
| `A14`             | `PA14`       |                   |
| `A15`             | `PA15`       |                   |
| `B3`              | `PB3`        |                   |
| `B4`              | `PB4`        |                   |
| `B5`              | `PB5`        |                   |
| `B6`              | `PB6`        | `UART_ALT_TX_PIN`, `I2C0_SCL_PIN` |
| `B7`              | `PB7`        | `UART_ALT_RX_PIN`, `I2C0_SDA_PIN` |
| `B8`              | `PB8`        |                   |
| `B9`              | `PB9`        |                   |

## Machine Package Docs

[Documentation for the machine package for the Bluepill](../machine/bluepill)

## Flashing

### OpenOCD

Programs are loaded onto the Bluepill with a separate hardware programmer such as the [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) board, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Bluepill board with your TinyGo code.

- Plug your STLink v2 programmer into your computer's USB port.
- Plug your Bluepill into the STLink v2 programmer using the Bluepill `SWIO`, `SWCLK`, `3V3` and `GND` pins.
- Build and flash your TinyGo program using `tinygo flash -target=bluepill`
