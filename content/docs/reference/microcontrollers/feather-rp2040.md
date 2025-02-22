---
title: "Adafruit Feather RP2040"
weight: 3
---

The [Adafruit Feather RP2040](https://www.adafruit.com/product/4884) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `D4`              | `GPIO6`      |                   | `PWM3` (channel A)   |
| `D5`              | `GPIO7`      |                   | `PWM3` (channel B)   |
| `D6`              | `GPIO8`      | `UART1_TX_PIN`    | `PWM4` (channel A)   |
| `D9`              | `GPIO9`      | `UART1_RX_PIN`    | `PWM4` (channel B)   |
| `D10`             | `GPIO10`     | `SPI1_SCK_PIN`    | `PWM5` (channel A)   |
| `D11`             | `GPIO11`     | `SPI1_SDO_PIN`    | `PWM5` (channel B)   |
| `D12`             | `GPIO12`     | `SPI1_SDI_PIN`    | `PWM6` (channel A)   |
| `D13`             | `GPIO13`     | `LED`             | `PWM6` (channel B)   |
| `D24`             | `GPIO24`     | `I2C0_SDA_PIN`    | `PWM4` (channel A)   |
| `D25`             | `GPIO25`     | `I2C0_SCL_PIN`    | `PWM4` (channel B)   |
| `A0`              | `GPIO26`     | `ADC0`            | `PWM5` (channel A)   |
| `A1`              | `GPIO27`     | `ADC1`            | `PWM5` (channel B)   |
| `A2`              | `GPIO28`     | `ADC2`            | `PWM6` (channel A)   |
| `A3`              | `GPIO29`     | `ADC3`            | `PWM6` (channel B)   |
| `I2C1_SDA_PIN`    | `GPIO2`      | `SDA_PIN`         | `PWM1` (channel A)   |
| `I2C1_SCL_PIN`    | `GPIO3`      | `SCL_PIN`         | `PWM1` (channel B)   |
| `SPI0_SCK_PIN`    | `GPIO18`     |                   | `PWM1` (channel A)   |
| `SPI0_SDO_PIN`    | `GPIO19`     |                   | `PWM1` (channel B)   |
| `SPI0_SDI_PIN`    | `GPIO20`     |                   | `PWM2` (channel A)   |
| `UART0_TX_PIN`    | `GPIO0`      | `UART_TX_PIN`     | `PWM0` (channel A)   |
| `UART0_RX_PIN`    | `GPIO1`      | `UART_RX_PIN`     | `PWM0` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Feather RP2040](../machine/feather-rp2040)

## Flashing

### UF2

The Feather RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Feather RP2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Feather RP2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
