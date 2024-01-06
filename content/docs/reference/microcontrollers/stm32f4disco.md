---
title: 'ST Micro STM32F4 Discovery'
weight: 3
---

The [STM32F4 Discovery](https://www.st.com/en/evaluation-tools/stm32f4discovery.html) is an ARM development board based on the ST Micro [STM32F407](https://www.st.com/resource/en/datasheet/stm32f407vg.pdf) SoC.

## Peripherals and Drivers

- LIS302DL or LIS3DSH accelerometer
- MP45DT02 MEMS digital microphone
- CS43L22 audio DAC
- 2 user buttons
- 4 user LEDs

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `LED1`            | `PD12`       | `LED_GREEN`, `LED`, `LED_BUILTIN` |
| `LED2`            | `PD13`       | `LED_ORANGE`      |
| `LED3`            | `PD14`       | `LED_RED`         |
| `LED4`            | `PD15`       | `LED_BLUE`        |
| `BUTTON`          | `PA0`        | `ADC0`            |
| `ADC1`            | `PA1`        |                   |
| `ADC2`            | `PA2`        | `UART_TX_PIN`     |
| `ADC3`            | `PA3`        | `UART_RX_PIN`     |
| `ADC4`            | `PA4`        |                   |
| `ADC5`            | `PA5`        | `SPI1_SCK_PIN`, `SPI0_SCK_PIN` |
| `ADC6`            | `PA6`        | `SPI1_SDI_PIN`, `SPI0_SDI_PIN` |
| `ADC7`            | `PA7`        | `SPI1_SDO_PIN`, `SPI0_SDO_PIN` |
| `ADC8`            | `PB0`        |                   |
| `ADC9`            | `PB1`        |                   |
| `ADC10`           | `PC0`        |                   |
| `ADC11`           | `PC1`        |                   |
| `ADC12`           | `PC2`        |                   |
| `ADC13`           | `PC3`        |                   |
| `ADC14`           | `PC4`        |                   |
| `ADC15`           | `PC5`        |                   |
| `MEMS_ACCEL_CS`   | `PE3`        |                   |
| `MEMS_ACCEL_INT1` | `PE0`        |                   |
| `MEMS_ACCEL_INT2` | `PE1`        |                   |
| `I2C0_SCL_PIN`    | `PB6`        |                   |
| `I2C0_SDA_PIN`    | `PB9`        |                   |

## Machine Package Docs

[Documentation for the machine package for the STM32F4 Discovery](../machine/stm32f4disco)

## Flashing

### OpenOCD

Programs are loaded onto the STM32F4 Discovery with the onboard [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the STM32F4 Discovery board with your TinyGo code.

- Plug your STM32F4 Discovery into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=stm32f4disco`
