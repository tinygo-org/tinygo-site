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
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `GP0`             | `GPIO0`      | `UART0_TX_PIN`, `UART_TX_PIN` |
| `GP1`             | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` |
| `GP2`             | `GPIO2`      | `SPI0_SCK_PIN`    |
| `GP3`             | `GPIO3`      | `SPI0_SDO_PIN`    |
| `GP4`             | `GPIO4`      | `SPI0_SDI_PIN`    |
| `GP6`             | `GPIO6`      | `I2C0_SCL_PIN`, `I2C1_SDA_PIN`, `SDA_PIN` |
| `GP7`             | `GPIO7`      | `I2C0_SDA_PIN`, `I2C1_SCL_PIN`, `SCL_PIN` |
| `GP8`             | `GPIO8`      |                   |
| `GP9`             | `GPIO9`      |                   |
| `GP10`            | `GPIO10`     |                   |
| `GP11`            | `GPIO11`     |                   |
| `GP12`            | `GPIO12`     | `SPI1_SDI_PIN`    |
| `GP14`            | `GPIO14`     | `SPI1_SCK_PIN`    |
| `GP15`            | `GPIO15`     | `SPI1_SDO_PIN`    |
| `GP16`            | `GPIO16`     |                   |
| `GP17`            | `GPIO17`     |                   |
| `GP18`            | `GPIO18`     |                   |
| `GP19`            | `GPIO19`     |                   |
| `GP20`            | `GPIO20`     |                   |
| `GP21`            | `GPIO21`     |                   |
| `GP22`            | `GPIO22`     |                   |
| `GP23`            | `GPIO23`     |                   |
| `GP25`            | `GPIO25`     | `LED`             |
| `GP26`            | `GPIO26`     | `A0`, `ADC0`      |
| `GP27`            | `GPIO27`     | `A1`, `ADC1`      |
| `GP28`            | `GPIO28`     | `A2`, `ADC2`      |
| `GP29`            | `GPIO29`     | `A3`, `ADC3`      |

## Machine Package Docs

[Documentation for the machine package for the Sparkfun Thing Plus RP2040](../machine/thingplus-rp2040)

## Flashing

### UF2

The Sparkfun Thing Plus RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Thing Plus RP2040 into your computer's USB port while pressing the BOOT button
- Release BOOT Button
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=thingplus-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Thing Plus RP2040 board should restart and run your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Thing Plus RP2040 as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.

The Neopixel LED and the SD Card are supported by [tinygo drivers](https://github.com/tinygo-org/drivers)