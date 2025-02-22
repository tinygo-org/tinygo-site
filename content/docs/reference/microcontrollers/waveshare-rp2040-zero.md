---
title: "Waveshare RP2040-Zero"
weight: 3
---

The [Waveshare RP2040-Zero](https://www.waveshare.com/wiki/RP2040-Zero) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `D0`              | `GPIO0`      | `I2C0_SDA_PIN`, `UART0_TX_PIN`, `UART_TX_PIN` | `PWM0` (channel A)   |
| `D1`              | `GPIO1`      | `I2C0_SCL_PIN`, `UART0_RX_PIN`, `UART_RX_PIN` | `PWM0` (channel B)   |
| `D2`              | `GPIO2`      | `I2C1_SDA_PIN`    | `PWM1` (channel A)   |
| `D3`              | `GPIO3`      | `I2C1_SCL_PIN`, `SPI0_SDO_PIN` | `PWM1` (channel B)   |
| `D4`              | `GPIO4`      | `SPI0_SDI_PIN`    | `PWM2` (channel A)   |
| `D5`              | `GPIO5`      |                   | `PWM2` (channel B)   |
| `D6`              | `GPIO6`      | `SPI0_SCK_PIN`    | `PWM3` (channel A)   |
| `D7`              | `GPIO7`      |                   | `PWM3` (channel B)   |
| `D8`              | `GPIO8`      | `UART1_TX_PIN`    | `PWM4` (channel A)   |
| `D9`              | `GPIO9`      | `UART1_RX_PIN`    | `PWM4` (channel B)   |
| `D10`             | `GPIO10`     | `SPI1_SCK_PIN`    | `PWM5` (channel A)   |
| `D11`             | `GPIO11`     | `SPI1_SDO_PIN`    | `PWM5` (channel B)   |
| `D12`             | `GPIO12`     | `SPI1_SDI_PIN`    | `PWM6` (channel A)   |
| `D13`             | `GPIO13`     |                   | `PWM6` (channel B)   |
| `D14`             | `GPIO14`     |                   | `PWM7` (channel A)   |
| `D15`             | `GPIO15`     |                   | `PWM7` (channel B)   |
| `D16`             | `GPIO16`     | `NEOPIXEL`, `WS2812` | `PWM0` (channel A)   |
| `D17`             | `GPIO17`     |                   | `PWM0` (channel B)   |
| `D18`             | `GPIO18`     |                   | `PWM1` (channel A)   |
| `D19`             | `GPIO19`     |                   | `PWM1` (channel B)   |
| `D20`             | `GPIO20`     |                   | `PWM2` (channel A)   |
| `D21`             | `GPIO21`     |                   | `PWM2` (channel B)   |
| `D22`             | `GPIO22`     |                   | `PWM3` (channel A)   |
| `D23`             | `GPIO23`     |                   | `PWM3` (channel B)   |
| `D24`             | `GPIO24`     |                   | `PWM4` (channel A)   |
| `D25`             | `GPIO25`     |                   | `PWM4` (channel B)   |
| `D26`             | `GPIO26`     | `A0`, `ADC0`      | `PWM5` (channel A)   |
| `D27`             | `GPIO27`     | `A1`, `ADC1`      | `PWM5` (channel B)   |
| `D28`             | `GPIO28`     | `A2`, `ADC2`      | `PWM6` (channel A)   |
| `D29`             | `GPIO29`     | `A3`, `ADC3`      | `PWM6` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Waveshare RP2040-Zero](../machine/waveshare-rp2040-zero)

## Flashing

### UF2

The Waveshare RP2040-Zero comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=waveshare-rp2040-zero [PATH TO YOUR PROGRAM]
    ```

- The Waveshare RP2040-Zero board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Waveshare RP2040-Zero as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
