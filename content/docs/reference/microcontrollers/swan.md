---
title: "Blues Wireless Swan"
weight: 3
---

The [Swan](https://blues.io/products/swan/) is an ARM development board based on the ST Micro [stm32l4r5zi](https://www.st.com/en/microcontrollers-microprocessors/stm32l4r5zi.html) SoC.

The Swan has a user button and an LED, LiPo charging and USB.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `LED`             | `PE2`        |                   |
| `UART_TX_PIN`     | `PA9`        |                   |
| `UART_RX_PIN`     | `PA10`       |                   |
| `I2C0_SCL_PIN`    | `PB6`        |                   |
| `I2C0_SDA_PIN`    | `PB7`        |                   |
| `SPI1_SCK_PIN`    | `PD1`        | `SPI0_SCK_PIN`    |
| `SPI1_SDI_PIN`    | `PB14`       | `SPI0_SDI_PIN`    |
| `SPI1_SDO_PIN`    | `PB15`       | `SPI0_SDO_PIN`    |

## Machine Package Docs

[Documentation for the machine package for the Swan](../machine/swan)

## Flashing

### DFU Util

Programs are loaded onto the Swan using DFU mode by default. You must install [DFU Util](http://dfu-util.sourceforge.net/) before you will be able to flash the STM32F4 Discovery board with your TinyGo code.

- Plug the Swan into your computer's USB port.
- Put the Swan into DFU mode (press and hold BOOT while toggling RESET).
- Build and flash your TinyGo program using `tinygo flash -target=swan`

### ST-Link

The Swan includes an SWD header that can be used with an ST Link such as the [STLINK-V3MINI](https://www.st.com/en/development-tools/stlink-v3mini.html). You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Swan with your TinyGo code.

- Connect the ST-Link debugger to the Swan using the ribbon cable.
- Plug the Swan and the ST-Link into USB ports on your computer.
- Build and flash your TinyGo program using `tinygo flash -target=swan -programmer=openocd`
