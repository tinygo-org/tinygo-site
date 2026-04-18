---
title: "Raspberry Pi Pico"
weight: 2
image: /images/boards/pico.svg
---

![Raspberry Pi Pico](/images/boards/pico.svg)

The [Raspberry Pi Pico](https://www.raspberrypi.org/products/raspberry-pi-pico/) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller dual-core ARM Cortex-M0+ processor with 264KB of SRAM and 2MB of onboard flash storage.

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
| `GP0`             | `GPIO0`      | `UART0_TX_PIN`, `UART_TX_PIN` | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `GP1`             | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `GP2`             | `GPIO2`      | `I2C1_SDA_PIN`    | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `GP3`             | `GPIO3`      | `I2C1_SCL_PIN`    | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `GP4`             | `GPIO4`      | `I2C0_SDA_PIN`    | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `GP5`             | `GPIO5`      | `I2C0_SCL_PIN`    | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `GP6`             | `GPIO6`      |                   | `I2C1` (SDA)         | `PWM3` (channel A)   |
| `GP7`             | `GPIO7`      |                   | `I2C1` (SCL)         | `PWM3` (channel B)   |
| `GP8`             | `GPIO8`      | `UART1_TX_PIN`    | `I2C0` (SDA)         | `PWM4` (channel A)   |
| `GP9`             | `GPIO9`      | `UART1_RX_PIN`    | `I2C0` (SCL)         | `PWM4` (channel B)   |
| `GP10`            | `GPIO10`     | `SPI1_SCK_PIN`    | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `GP11`            | `GPIO11`     | `SPI1_SDO_PIN`    | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `GP12`            | `GPIO12`     | `SPI1_SDI_PIN`    | `I2C0` (SDA)         | `PWM6` (channel A)   |
| `GP13`            | `GPIO13`     |                   | `I2C0` (SCL)         | `PWM6` (channel B)   |
| `GP14`            | `GPIO14`     |                   | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `GP15`            | `GPIO15`     |                   | `I2C1` (SCL)         | `PWM7` (channel B)   |
| `GP16`            | `GPIO16`     | `SPI0_SDI_PIN`    | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `GP17`            | `GPIO17`     |                   | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `GP18`            | `GPIO18`     | `SPI0_SCK_PIN`    | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `GP19`            | `GPIO19`     | `SPI0_SDO_PIN`    | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `GP20`            | `GPIO20`     |                   | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `GP21`            | `GPIO21`     |                   | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `GP22`            | `GPIO22`     |                   |                      | `PWM3` (channel A)   |
| `GP26`            | `GPIO26`     | `ADC0`            | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `GP27`            | `GPIO27`     | `ADC1`            | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `GP28`            | `GPIO28`     | `ADC2`            |                      | `PWM6` (channel A)   |
| `LED`             | `GPIO25`     |                   |                      | `PWM4` (channel B)   |
| `ADC3`            | `GPIO29`     |                   |                      | `PWM6` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Pico](../../machine/pico)

## Flashing

### UF2

The Pico comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pico [PATH TO YOUR PROGRAM]
    ```

- The Pico board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Pico as a serial port.

### Programmable Input/Output (PIO)

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)

### Debugging

You can refer to [getting started with Raspberry Pi Pico](https://datasheets.raspberrypi.org/pico/getting-started-with-pico.pdf) documentation on how to connect two Picos together (see Appendix A: Using Picoprobe) to debug and convert `UART0` output on target pico to USB output on picoprobe. You will need the [Picoprobe UF2](https://www.raspberrypi.org/documentation/rp2040/getting-started/#board-specifications), available on the Pico's website under "About" tab.
