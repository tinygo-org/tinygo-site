---
title: "Sparkfun Thing Plus RP2040"
weight: 3
---

The [Sparkfun Thing Plus RP2040](https://www.sparkfun.com/products/17745) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller. 

Peripherals: 
- ws2812 Neopixel
- sdcard

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
| `GP0`             | `GPIO0`      | `UART0_TX_PIN`, `UART_TX_PIN` | `PWM0` (channel A)   |
| `GP1`             | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` | `PWM0` (channel B)   |
| `GP2`             | `GPIO2`      | `SPI0_SCK_PIN`    | `PWM1` (channel A)   |
| `GP3`             | `GPIO3`      | `SPI0_SDO_PIN`    | `PWM1` (channel B)   |
| `GP4`             | `GPIO4`      | `SPI0_SDI_PIN`    | `PWM2` (channel A)   |
| `GP6`             | `GPIO6`      | `I2C0_SCL_PIN`, `I2C1_SDA_PIN`, `SDA_PIN` | `PWM3` (channel A)   |
| `GP7`             | `GPIO7`      | `I2C0_SDA_PIN`, `I2C1_SCL_PIN`, `SCL_PIN` | `PWM3` (channel B)   |
| `GP8`             | `GPIO8`      | `WS2812`          | `PWM4` (channel A)   |
| `GP9`             | `GPIO9`      |                   | `PWM4` (channel B)   |
| `GP10`            | `GPIO10`     |                   | `PWM5` (channel A)   |
| `GP11`            | `GPIO11`     |                   | `PWM5` (channel B)   |
| `GP12`            | `GPIO12`     | `SPI1_SDI_PIN`    | `PWM6` (channel A)   |
| `GP14`            | `GPIO14`     | `SPI1_SCK_PIN`    | `PWM7` (channel A)   |
| `GP15`            | `GPIO15`     | `SPI1_SDO_PIN`    | `PWM7` (channel B)   |
| `GP16`            | `GPIO16`     |                   | `PWM0` (channel A)   |
| `GP17`            | `GPIO17`     |                   | `PWM0` (channel B)   |
| `GP18`            | `GPIO18`     |                   | `PWM1` (channel A)   |
| `GP19`            | `GPIO19`     |                   | `PWM1` (channel B)   |
| `GP20`            | `GPIO20`     |                   | `PWM2` (channel A)   |
| `GP21`            | `GPIO21`     |                   | `PWM2` (channel B)   |
| `GP22`            | `GPIO22`     |                   | `PWM3` (channel A)   |
| `GP23`            | `GPIO23`     |                   | `PWM3` (channel B)   |
| `GP25`            | `GPIO25`     | `LED`             | `PWM4` (channel B)   |
| `GP26`            | `GPIO26`     | `A0`, `ADC0`      | `PWM5` (channel A)   |
| `GP27`            | `GPIO27`     | `A1`, `ADC1`      | `PWM5` (channel B)   |
| `GP28`            | `GPIO28`     | `A2`, `ADC2`      | `PWM6` (channel A)   |
| `GP29`            | `GPIO29`     | `A3`, `ADC3`      | `PWM6` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Sparkfun Thing Plus RP2040](../machine/thingplus-rp2040)

## Flashing

### UF2

The Sparkfun Thing Plus RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=thingplus-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Thing Plus RP2040 board should restart and run your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Thing Plus RP2040 as a serial port.

The Neopixel LED and the SD Card are supported by [tinygo drivers](https://github.com/tinygo-org/drivers)

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
