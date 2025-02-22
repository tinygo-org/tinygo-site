---
title: "Arduino Nano RP2040 Connect"
weight: 3
---

The [Nano RP2040 Connect](https://store.arduino.cc/nano-rp2040-connect) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller. 

Peripherals: 
- NINA-W102 chip with [wifinina](https://github.com/tinygo-org/drivers/tree/release/wifinina) firmware (wifi and bluetooth)
- [lsm6dox](https://github.com/tinygo-org/drivers/tree/release/lsm6dox) IMU chip (acceleration, rotation and temperature)
- microphone

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
| `D2`              | `GPIO25`     |                   | `PWM4` (channel B)   |
| `D3`              | `GPIO15`     |                   | `PWM7` (channel B)   |
| `D4`              | `GPIO16`     |                   | `PWM0` (channel A)   |
| `D5`              | `GPIO17`     |                   | `PWM0` (channel B)   |
| `D6`              | `GPIO18`     | `I2C1_SDA_PIN`    | `PWM1` (channel A)   |
| `D7`              | `GPIO19`     | `I2C1_SCL_PIN`    | `PWM1` (channel B)   |
| `D8`              | `GPIO20`     |                   | `PWM2` (channel A)   |
| `D9`              | `GPIO21`     |                   | `PWM2` (channel B)   |
| `D10`             | `GPIO5`      |                   | `PWM2` (channel B)   |
| `D11`             | `GPIO7`      | `SPI0_SDO_PIN`    | `PWM3` (channel B)   |
| `D12`             | `GPIO4`      | `SPI0_SDI_PIN`    | `PWM2` (channel A)   |
| `D13`             | `GPIO6`      | `LED`, `SPI0_SCK_PIN` | `PWM3` (channel A)   |
| `D14`             | `GPIO26`     | `A0`, `ADC0`      | `PWM5` (channel A)   |
| `D15`             | `GPIO27`     | `A1`, `ADC1`      | `PWM5` (channel B)   |
| `D16`             | `GPIO28`     | `A2`, `ADC2`      | `PWM6` (channel A)   |
| `D17`             | `GPIO29`     | `A3`, `ADC3`      | `PWM6` (channel B)   |
| `D18`             | `GPIO12`     | `I2C0_SDA_PIN`    | `PWM6` (channel A)   |
| `D19`             | `GPIO13`     | `I2C0_SCL_PIN`    | `PWM6` (channel B)   |
| `SPI1_SCK_PIN`    | `GPIO22`     | `SPI1_SDO_PIN`, `SPI1_SDI_PIN` | `PWM3` (channel A)   |
| `NINA_SCK`        | `GPIO14`     |                   | `PWM7` (channel A)   |
| `NINA_SDO`        | `GPIO11`     | `NINA_RTS`        | `PWM5` (channel B)   |
| `NINA_SDI`        | `GPIO8`      | `NINA_TX`         | `PWM4` (channel A)   |
| `NINA_CS`         | `GPIO9`      | `NINA_RX`         | `PWM4` (channel B)   |
| `NINA_ACK`        | `GPIO10`     | `NINA_CTS`        | `PWM5` (channel A)   |
| `NINA_GPIO0`      | `GPIO2`      |                   | `PWM1` (channel A)   |
| `NINA_RESETN`     | `GPIO3`      |                   | `PWM1` (channel B)   |
| `UART0_TX_PIN`    | `GPIO0`      | `UART_TX_PIN`     | `PWM0` (channel A)   |
| `UART0_RX_PIN`    | `GPIO1`      | `UART_RX_PIN`     | `PWM0` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Nano RP2040](../machine/nano-rp2040)

## Flashing

### UF2

The Nano RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=nano-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Nano RP2040 board should restart and run your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Nano RP2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
