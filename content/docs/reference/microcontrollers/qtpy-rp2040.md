---
title: "Adafruit QT Py RP2040"
weight: 3
---

The [Adafruit QT Py RP2040](https://www.adafruit.com/product/4900) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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

## Pins

| Pin               | Hardware pin | Alternative names | PWM                  |
| ----------------- | ------------ | ----------------- | -------------------- |
| `SDA`             | `GPIO24`     | `I2C0_SDA_PIN`, `SDA_PIN`, `SPI1_SDI_PIN` | `PWM4` (channel A)   |
| `SCL`             | `GPIO25`     | `I2C0_SCL_PIN`, `SCL_PIN`, `SPI1_CS` | `PWM4` (channel B)   |
| `TX`              | `GPIO20`     | `UART1_TX_PIN`    | `PWM2` (channel A)   |
| `MO`              | `GPIO3`      | `MOSI`, `SPI0_SDO_PIN` | `PWM1` (channel B)   |
| `MI`              | `GPIO4`      | `MISO`, `SPI0_SDI_PIN` | `PWM2` (channel A)   |
| `SCK`             | `GPIO6`      | `SPI0_SCK_PIN`    | `PWM3` (channel A)   |
| `RX`              | `GPIO5`      | `SPI0_CS`, `UART1_RX_PIN` | `PWM2` (channel B)   |
| `QT_SCL1`         | `GPIO23`     | `I2C1_QT_SCL_PIN` | `PWM3` (channel B)   |
| `QT_SDA1`         | `GPIO22`     | `I2C1_QT_SDA_PIN` | `PWM3` (channel A)   |
| `A0`              | `GPIO29`     | `UART0_RX_PIN`, `UART_RX_PIN`, `ADC3` | `PWM6` (channel B)   |
| `A1`              | `GPIO28`     | `UART0_TX_PIN`, `UART_TX_PIN`, `ADC2` | `PWM6` (channel A)   |
| `A2`              | `GPIO27`     | `I2C1_SCL_PIN`, `SPI1_SDO_PIN`, `ADC1` | `PWM5` (channel B)   |
| `A3`              | `GPIO26`     | `I2C1_SDA_PIN`, `SPI1_SCK_PIN`, `ADC0` | `PWM5` (channel A)   |
| `NEOPIXEL`        | `GPIO12`     | `WS2812`          | `PWM6` (channel A)   |
| `NEOPIXEL_POWER`  | `GPIO11`     |                   | `PWM5` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit QT Py RP2040](../machine/qtpy-rp2040)

## Flashing

### UF2

The QT Py RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=qtpy-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The QT Py RP2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the QT Py RP2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
