---
title: "Adafruit MacroPad RP2040"
weight: 3
---

The [Adafruit MacroPad RP2040](https://www.adafruit.com/product/5100) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `SWITCH`          | `GPIO0`      | `BUTTON`, `UART0_TX_PIN`, `UART_TX_PIN` | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `KEY1`            | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `KEY2`            | `GPIO2`      |                   | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `KEY3`            | `GPIO3`      |                   | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `KEY4`            | `GPIO4`      |                   | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `KEY5`            | `GPIO5`      |                   | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `KEY6`            | `GPIO6`      |                   | `I2C1` (SDA)         | `PWM3` (channel A)   |
| `KEY7`            | `GPIO7`      |                   | `I2C1` (SCL)         | `PWM3` (channel B)   |
| `KEY8`            | `GPIO8`      |                   | `I2C0` (SDA)         | `PWM4` (channel A)   |
| `KEY9`            | `GPIO9`      |                   | `I2C0` (SCL)         | `PWM4` (channel B)   |
| `KEY10`           | `GPIO10`     |                   | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `KEY11`           | `GPIO11`     |                   | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `KEY12`           | `GPIO12`     |                   | `I2C0` (SDA)         | `PWM6` (channel A)   |
| `LED`             | `GPIO13`     |                   | `I2C0` (SCL)         | `PWM6` (channel B)   |
| `SPEAKER_ENABLE`  | `GPIO14`     |                   | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `SPEAKER`         | `GPIO16`     |                   | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `ROT_A`           | `GPIO18`     |                   | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `ROT_B`           | `GPIO17`     |                   | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `OLED_CS`         | `GPIO22`     |                   |                      | `PWM3` (channel A)   |
| `OLED_RST`        | `GPIO23`     |                   |                      | `PWM3` (channel B)   |
| `OLED_DC`         | `GPIO24`     |                   |                      | `PWM4` (channel A)   |
| `NEOPIXEL`        | `GPIO19`     | `WS2812`          | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `I2C0_SDA_PIN`    | `GPIO20`     |                   | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `I2C0_SCL_PIN`    | `GPIO21`     |                   | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `SPI1_SCK_PIN`    | `GPIO26`     | `ADC0`            | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `SPI1_SDO_PIN`    | `GPIO27`     | `ADC1`            | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `SPI1_SDI_PIN`    | `GPIO28`     | `ADC2`            |                      | `PWM6` (channel A)   |
| `ADC3`            | `GPIO29`     |                   |                      | `PWM6` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit MacroPad RP2040](../../machine/macropad-rp2040)

## Flashing

### UF2

The MacroPad RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=macropad-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The MacroPad RP2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the MacroPad RP2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
