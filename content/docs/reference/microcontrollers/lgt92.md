---
title: "Dragino LoRaWAN GPS Tracker LGT-92"
weight: 3
---

The [LoRaWAN GPS Tracker LGT-92](https://www.dragino.com/products/lora-lorawan-end-node/item/142-lgt-92.html) is based on the ST Micro [STM32L072CZT6](https://www.st.com/en/microcontrollers-microprocessors/stm32l072cz.html) SoC.

## Peripherals and Drivers

- low power GPS module L76-L
- 9-axis accelerometer for motion and attitude detection

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `LED1`            | `PA12`       | `LED_RED`, `LED`  |
| `LED2`            | `PA8`        | `LED_BLUE`        |
| `LED3`            | `PA11`       | `LED_GREEN`       |
| `BUTTON`          | `PB14`       |                   |
| `GPS_STANDBY_PIN` | `PB3`        |                   |
| `GPS_RESET_PIN`   | `PB4`        |                   |
| `GPS_POWER_PIN`   | `PB5`        |                   |
| `MEMS_ACCEL_CS`   | `PE3`        |                   |
| `MEMS_ACCEL_INT1` | `PE0`        |                   |
| `MEMS_ACCEL_INT2` | `PE1`        |                   |
| `SPI1_SCK_PIN`    | `PA5`        | `SPI0_SCK_PIN`    |
| `SPI1_SDI_PIN`    | `PA6`        | `SPI0_SDI_PIN`    |
| `SPI1_SDO_PIN`    | `PA7`        | `SPI0_SDO_PIN`    |
| `RFM95_DIO0_PIN`  | `PC13`       |                   |
| `UART_RX_PIN`     | `PA13`       |                   |
| `UART_TX_PIN`     | `PA14`       |                   |
| `UART1_RX_PIN`    | `PB6`        |                   |
| `UART1_TX_PIN`    | `PB7`        |                   |
| `I2C0_SCL_PIN`    | `PA9`        |                   |
| `I2C0_SDA_PIN`    | `PA10`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Dragino LGT-92](../machine/lgt-92)

## Flashing

### OpenOCD

Programs are loaded onto the Dragino LGT-92 with a separate hardware programmer such as the [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) board, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Dragino LGT-92 board with your TinyGo code.

- Plug your STLink v2 programmer into your computer's USB port.
- Plug your Dragino LGT-92 into the STLink v2 programmer using the Dragino LGT-92 `SWIO`, `SWCLK`, `3V3` and `GND` pins.
- Build and flash your TinyGo program using `tinygo flash -target=lgt92`
