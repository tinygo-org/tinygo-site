---
title: "Seeed XIAO RP2040"
weight: 3
---

The [Seeed XIAO RP2040](https://www.seeedstudio.com/XIAO-RP2040-v1-0-p-5026.html) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `D0`              | `GPIO26`     | `A0`, `ADC0`      | `PWM5` (channel A)   |
| `D1`              | `GPIO27`     | `A1`, `ADC1`      | `PWM5` (channel B)   |
| `D2`              | `GPIO28`     | `A2`, `I2C0_SDA_PIN`, `ADC2` | `PWM6` (channel A)   |
| `D3`              | `GPIO29`     | `A3`, `I2C0_SCL_PIN`, `ADC3` | `PWM6` (channel B)   |
| `D4`              | `GPIO6`      | `I2C1_SDA_PIN`    | `PWM3` (channel A)   |
| `D5`              | `GPIO7`      | `I2C1_SCL_PIN`    | `PWM3` (channel B)   |
| `D6`              | `GPIO0`      | `UART0_TX_PIN`, `UART_TX_PIN` | `PWM0` (channel A)   |
| `D7`              | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` | `PWM0` (channel B)   |
| `D8`              | `GPIO2`      | `SPI0_SCK_PIN`    | `PWM1` (channel A)   |
| `D9`              | `GPIO4`      | `SPI0_SDI_PIN`    | `PWM2` (channel A)   |
| `D10`             | `GPIO3`      | `SPI0_SDO_PIN`    | `PWM1` (channel B)   |
| `NEOPIXEL`        | `GPIO12`     | `WS2812`          | `PWM6` (channel A)   |
| `NEO_PWR`         | `GPIO11`     | `NEOPIXEL_POWER`  | `PWM5` (channel B)   |
| `LED`             | `GPIO17`     | `LED_RED`         | `PWM0` (channel B)   |
| `LED_GREEN`       | `GPIO16`     |                   | `PWM0` (channel A)   |
| `LED_BLUE`        | `GPIO25`     |                   | `PWM4` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Seeed XIAO RP2040](../machine/xiao-rp2040)

## Flashing

### UF2

The XIAO RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=xiao-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The XIAO RP2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the XIAO RP2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
