---
title: "Seeed XIAO BLE"
weight: 3
---

The [Seeed XIAO BLE](https://www.seeedstudio.com/Seeed-XIAO-BLE-nRF52840-p-5201.html) is a tiny ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | YES |
| USBDevice | YES | YES |
| Bluetooth | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `P0_02`      | `A0`              |
| `D1`              | `P0_03`      | `A1`              |
| `D2`              | `P0_28`      | `A2`              |
| `D3`              | `P0_29`      | `A3`              |
| `D4`              | `P0_04`      | `A4`, `SDA0_PIN`  |
| `D5`              | `P0_05`      | `A5`, `SCL0_PIN`  |
| `D6`              | `P1_11`      | `UART_TX_PIN`     |
| `D7`              | `P1_12`      | `UART_RX_PIN`     |
| `D8`              | `P1_13`      | `SPI0_SCK_PIN`    |
| `D9`              | `P1_14`      | `SPI0_SDO_PIN`    |
| `D10`             | `P1_15`      | `SPI0_SDI_PIN`    |
| `LED`             | `P0_17`      | `LED_CHG`         |
| `LED1`            | `P0_26`      | `LED_RED`         |
| `LED2`            | `P0_30`      | `LED_GREEN`       |
| `LED3`            | `P0_06`      | `LED_BLUE`        |
| `SDA_PIN`         | `P0_07`      | `SDA1_PIN`        |
| `SCL_PIN`         | `P0_27`      | `SCL1_PIN`        |
| `LSM_PWR`         | `P1_08`      |                   |
| `LSM_INT`         | `P0_11`      |                   |
| `MIC_PWR`         | `P1_10`      |                   |
| `MIC_CLK`         | `P1_00`      |                   |
| `MIC_DIN`         | `P0_16`      |                   |

## Machine Package Docs

[Documentation for the machine package for the Seeed XIAO BLE](../machine/xiao-ble)

## Flashing

### UF2

The XIAO BLE comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your XIAO BLE into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=xiao-ble [PATH TO YOUR PROGRAM]
    ```

- The XIAO BLE board should restart and then begin running your program.

### Troubleshooting

Add troubleshooting tips here.

## Notes

You can use the USB port to the XIAO BLE as a serial port. `UART0` refers to this connection.

