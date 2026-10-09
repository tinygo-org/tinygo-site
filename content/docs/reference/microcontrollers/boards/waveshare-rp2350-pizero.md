---
title: "Waveshare RP2350-PiZero"
weight: 3
---

The [Waveshare RP2350-PiZero](https://www.waveshare.com/wiki/RP2350-PiZero) is a development board based on the Raspberry Pi [RP2350B](https://datasheets.raspberrypi.org/rp2350/rp2350-datasheet.pdf) microcontroller. It has the same form factor as the Raspberry Pi Zero.

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
| `GP0`             | `GPIO0`      | `UART0_TX_PIN`    | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `GP1`             | `GPIO1`      | `UART0_RX_PIN`    | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `GP2`             | `GPIO2`      | `I2C1_SDA_PIN`    | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `GP3`             | `GPIO3`      | `I2C1_SCL_PIN`    | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `GP4`             | `GPIO4`      | `UART1_TX_PIN`, `UART_TX_PIN` | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `GP5`             | `GPIO5`      | `UART1_RX_PIN`, `UART_RX_PIN` | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `GP6`             | `GPIO6`      |                   | `I2C1` (SDA)         | `PWM3` (channel A)   |
| `GP7`             | `GPIO7`      |                   | `I2C1` (SCL)         | `PWM3` (channel B)   |
| `GP8`             | `GPIO8`      |                   | `I2C0` (SDA)         | `PWM4` (channel A)   |
| `GP9`             | `GPIO9`      |                   | `I2C0` (SCL)         | `PWM4` (channel B)   |
| `GP10`            | `GPIO10`     | `SPI1_SCK_PIN`    | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `GP11`            | `GPIO11`     | `SPI1_SDO_PIN`    | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `GP12`            | `GPIO12`     | `SPI1_SDI_PIN`    | `I2C0` (SDA)         | `PWM6` (channel A)   |
| `GP13`            | `GPIO13`     |                   | `I2C0` (SCL)         | `PWM6` (channel B)   |
| `GP14`            | `GPIO14`     |                   | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `GP15`            | `GPIO15`     |                   | `I2C1` (SCL)         | `PWM7` (channel B)   |
| `GP16`            | `GPIO16`     |                   | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `GP17`            | `GPIO17`     |                   | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `GP18`            | `GPIO18`     |                   | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `GP19`            | `GPIO19`     |                   | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `GP20`            | `GPIO20`     |                   | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `GP21`            | `GPIO21`     |                   | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `GP22`            | `GPIO22`     |                   |                      | `PWM3` (channel A)   |
| `GP23`            | `GPIO23`     |                   |                      | `PWM3` (channel B)   |
| `GP24`            | `GPIO24`     |                   |                      | `PWM4` (channel A)   |
| `GP25`            | `GPIO25`     |                   |                      | `PWM4` (channel B)   |
| `GP26`            | `GPIO26`     |                   | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `GP27`            | `GPIO27`     |                   | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `GP28`            | `GPIO28`     |                   |                      | `PWM6` (channel A)   |
| `GP29`            | `GPIO29`     |                   |                      | `PWM6` (channel B)   |
| `GP30`            | `GPIO30`     |                   | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `GP31`            | `GPIO31`     |                   | `I2C1` (SCL)         | `PWM7` (channel B)   |
| `GP32`            | `GPIO32`     |                   | `I2C0` (SDA)         | `PWM8` (channel A)   |
| `GP33`            | `GPIO33`     |                   | `I2C0` (SCL)         | `PWM8` (channel B)   |
| `GP34`            | `GPIO34`     |                   | `I2C1` (SDA)         | `PWM9` (channel A)   |
| `GP35`            | `GPIO35`     |                   | `I2C1` (SCL)         | `PWM9` (channel B)   |
| `GP36`            | `GPIO36`     |                   | `I2C0` (SDA)         | `PWM10` (channel A)  |
| `GP37`            | `GPIO37`     |                   | `I2C0` (SCL)         | `PWM10` (channel B)  |
| `GP38`            | `GPIO38`     |                   | `I2C1` (SDA)         | `PWM11` (channel A)  |
| `GP39`            | `GPIO39`     |                   | `I2C1` (SCL)         | `PWM11` (channel B)  |
| `GP40`            | `GPIO40`     | `ADC0`            | `I2C0` (SDA)         | `PWM8` (channel A)   |
| `GP41`            | `GPIO41`     | `ADC1`            | `I2C0` (SCL)         | `PWM8` (channel B)   |
| `GP42`            | `GPIO42`     | `ADC2`            | `I2C1` (SDA)         | `PWM9` (channel A)   |
| `GP43`            | `GPIO43`     | `ADC3`            | `I2C1` (SCL)         | `PWM9` (channel B)   |
| `GP44`            | `GPIO44`     | `ADC4`            | `I2C0` (SDA)         | `PWM10` (channel A)  |
| `GP45`            | `GPIO45`     | `ADC5`            | `I2C0` (SCL)         | `PWM10` (channel B)  |
| `GP46`            | `GPIO46`     | `ADC6`            | `I2C1` (SDA)         | `PWM11` (channel A)  |
| `GP47`            | `GPIO47`     | `ADC7`            | `I2C1` (SCL)         | `PWM11` (channel B)  |

## Machine Package Docs

[Documentation for the machine package for the Waveshare RP2350-PiZero](../../machine/waveshare-rp2350-pizero)

## Flashing

### UF2

The Waveshare RP2350-PiZero comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=waveshare-rp2350-pizero [PATH TO YOUR PROGRAM]
    ```

- The Waveshare RP2350-PiZero board should restart and then begin running your program.

## Notes

You can use the USB port to the Waveshare RP2350-PiZero as a serial port.

TinyGo has support for the RP2350's on-board Programmable Input/Output (PIO) block.

For more information, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
