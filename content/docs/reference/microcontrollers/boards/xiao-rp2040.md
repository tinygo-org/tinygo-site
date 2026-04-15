---
title: "Seeed Studio XIAO RP2040"
weight: 1
---

The [Seeed Studio XIAO RP2040](https://www.seeedstudio.com/XIAO-RP2040-v1-0-p-5026.html) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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

| Pin               | Hardware pin | Alternative names | I2C                  | PWM                  |
| ----------------- | ------------ | ----------------- | -------------------- | -------------------- |
| `D0`              | `GPIO26`     | `A0`, `ADC0`      | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `D1`              | `GPIO27`     | `A1`, `ADC1`      | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `D2`              | `GPIO28`     | `A2`, `I2C0_SDA_PIN`, `ADC2` |                      | `PWM6` (channel A)   |
| `D3`              | `GPIO29`     | `A3`, `I2C0_SCL_PIN`, `ADC3` |                      | `PWM6` (channel B)   |
| `D4`              | `GPIO6`      | `I2C1_SDA_PIN`    | `I2C1` (SDA)         | `PWM3` (channel A)   |
| `D5`              | `GPIO7`      | `I2C1_SCL_PIN`    | `I2C1` (SCL)         | `PWM3` (channel B)   |
| `D6`              | `GPIO0`      | `UART0_TX_PIN`, `UART_TX_PIN` | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `D7`              | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `D8`              | `GPIO2`      | `SPI0_SCK_PIN`    | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `D9`              | `GPIO4`      | `SPI0_SDI_PIN`    | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `D10`             | `GPIO3`      | `SPI0_SDO_PIN`    | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `NEOPIXEL`        | `GPIO12`     | `WS2812`          | `I2C0` (SDA)         | `PWM6` (channel A)   |
| `NEO_PWR`         | `GPIO11`     | `NEOPIXEL_POWER`  | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `LED`             | `GPIO17`     | `LED_RED`         | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `LED_GREEN`       | `GPIO16`     |                   | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `LED_BLUE`        | `GPIO25`     |                   |                      | `PWM4` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Seeed Studio XIAO RP2040](../../machine/xiao-rp2040)

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

For more information, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
