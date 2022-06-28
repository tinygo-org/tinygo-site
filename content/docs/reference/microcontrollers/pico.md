---
title: "Raspberry Pi Pico"
weight: 3
---

The [Raspberry Pi Pico](https://www.raspberrypi.org/products/raspberry-pi-pico/) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `GP2`             | `GPIO2`      | `I2C1_SDA_PIN`    |
| `GP3`             | `GPIO3`      | `I2C1_SCL_PIN`    |
| `GP4`             | `GPIO4`      | `I2C0_SDA_PIN`    |
| `GP5`             | `GPIO5`      | `I2C0_SCL_PIN`    |
| `GP6`             | `GPIO6`      |                   |
| `GP7`             | `GPIO7`      |                   |
| `GP8`             | `GPIO8`      | `UART1_TX_PIN`    |
| `GP9`             | `GPIO9`      | `UART1_RX_PIN`    |
| `GP10`            | `GPIO10`     | `SPI1_SCK_PIN`    |
| `GP11`            | `GPIO11`     | `SPI1_SDO_PIN`    |
| `GP12`            | `GPIO12`     | `SPI1_SDI_PIN`    |
| `GP13`            | `GPIO13`     |                   |
| `GP14`            | `GPIO14`     |                   |
| `GP15`            | `GPIO15`     |                   |
| `GP16`            | `GPIO16`     | `SPI0_SDI_PIN`    |
| `GP17`            | `GPIO17`     |                   |
| `GP18`            | `GPIO18`     | `SPI0_SCK_PIN`    |
| `GP19`            | `GPIO19`     | `SPI0_SDO_PIN`    |
| `GP20`            | `GPIO20`     |                   |
| `GP21`            | `GPIO21`     |                   |
| `GP22`            | `GPIO22`     |                   |
| `GP26`            | `GPIO26`     | `ADC0`            |
| `GP27`            | `GPIO27`     | `ADC1`            |
| `GP28`            | `GPIO28`     | `ADC2`            |
| `LED`             | `GPIO25`     |                   |
| `ADC3`            | `GPIO29`     |                   |

## Machine Package Docs

[Documentation for the machine package for the Pico](../machine/pico)

## Flashing

### UF2

The Pico comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Pico into your computer's USB port while holding down the RESET button on the board.
- One plugged in, release the RESET button.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pico [PATH TO YOUR PROGRAM]
    ```

- The Pico board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Pico as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.

You can refer to [getting started with Raspberry Pi Pico](https://datasheets.raspberrypi.org/pico/getting-started-with-pico.pdf) documentation on how to connect two Picos together (see Appendix A: Using Picoprobe) to debug and convert `UART0` output on target pico to USB output on picoprobe. You will need the [Picoprobe UF2](https://www.raspberrypi.org/documentation/rp2040/getting-started/#board-specifications), available on the Pico's website under "About" tab.
