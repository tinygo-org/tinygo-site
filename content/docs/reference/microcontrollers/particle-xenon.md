---
title: "Particle Xenon"
weight: 3
---

[The Particle Xenon](https://docs.particle.io/datasheets/discontinued/xenon-datasheet/) is a low cost mesh-enabled development board.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |
| Bluetooth      | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `A0`              | `P0_03`      |                   |
| `A1`              | `P0_04`      |                   |
| `A2`              | `P0_28`      |                   |
| `A3`              | `P0_29`      |                   |
| `A4`              | `P0_30`      |                   |
| `A5`              | `P0_31`      |                   |
| `D0`              | `P0_26`      | `SDA_PIN`         |
| `D1`              | `P0_27`      | `SCL_PIN`         |
| `D2`              | `P1_01`      |                   |
| `D3`              | `P1_02`      |                   |
| `D4`              | `P1_08`      |                   |
| `D5`              | `P1_10`      |                   |
| `D6`              | `P1_11`      |                   |
| `D7`              | `P1_12`      | `LED`             |
| `D8`              | `P1_03`      |                   |
| `D9`              | `P0_06`      | `UART_TX_PIN`     |
| `D10`             | `P0_08`      | `UART_RX_PIN`     |
| `D11`             | `P1_14`      | `SPI0_SDI_PIN`    |
| `D12`             | `P1_13`      | `SPI0_SDO_PIN`    |
| `D13`             | `P1_15`      | `SPI0_SCK_PIN`    |
| `LED_GREEN`       | `P0_14`      |                   |
| `LED_RED`         | `P0_13`      |                   |
| `LED_BLUE`        | `P0_15`      |                   |
| `SPI1_SCK_PIN`    | `P0_19`      |                   |
| `SPI1_SDO_PIN`    | `P0_20`      |                   |
| `SPI1_SDI_PIN`    | `P0_21`      |                   |
| `SPI1_CS_PIN`     | `P0_17`      |                   |
| `SPI1_WP_PIN`     | `P0_22`      |                   |
| `SPI1_HOLD_PIN`   | `P0_23`      |                   |
| `MODE_BUTTON_PIN` | `P0_11`      |                   |
| `CHARGE_STATUS_PIN` | `P1_09`      |                   |
| `LIPO_VOLTAGE_PIN` | `P0_05`      |                   |
| `PCB_ANTENNA_PIN` | `P0_24`      |                   |
| `EXTERNAL_UFL_PIN` | `P0_25`      |                   |
| `NFC1_PIN`        | `P0_09`      |                   |
| `NFC2_PIN`        | `P0_10`      |                   |

## Machine Package Docs

[Documentation for the machine package for the Particle Xenon](../machine/particle-xenon)

## Flashing

### OpenOCD

Programs are loaded onto the Particle Xenon board using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Particle Xenon board with your TinyGo code.

- Plug your [Particle Debugger](https://store.particle.io/collections/accessories/products/particle-debugger) to Xenon's debugging connector
- Build and flash your TinyGo program using `tinygo flash -target=particle-xenon`

## Notes

You can use the USB port to the Particle Xenon as a serial port. `UART0` refers to this connection.

Bluetooth support is in development but not yet completed.
