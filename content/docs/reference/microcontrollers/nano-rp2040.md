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

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D2`              | `GPIO25`     |                   |
| `D3`              | `GPIO15`     |                   |
| `D4`              | `GPIO16`     |                   |
| `D5`              | `GPIO17`     |                   |
| `D6`              | `GPIO18`     | `I2C1_SDA_PIN`    |
| `D7`              | `GPIO19`     | `I2C1_SCL_PIN`    |
| `D8`              | `GPIO20`     |                   |
| `D9`              | `GPIO21`     |                   |
| `D10`             | `GPIO5`      |                   |
| `D11`             | `GPIO7`      | `SPI0_SDO_PIN`    |
| `D12`             | `GPIO4`      | `SPI0_SDI_PIN`    |
| `D13`             | `GPIO6`      | `LED`, `SPI0_SCK_PIN` |
| `D14`             | `GPIO26`     | `A0`, `ADC0`      |
| `D15`             | `GPIO27`     | `A1`, `ADC1`      |
| `D16`             | `GPIO28`     | `A2`, `ADC2`      |
| `D17`             | `GPIO29`     | `A3`, `ADC3`      |
| `D18`             | `GPIO12`     | `I2C0_SDA_PIN`    |
| `D19`             | `GPIO13`     | `I2C0_SCL_PIN`    |
| `SPI1_SCK_PIN`    | `GPIO22`     | `SPI1_SDO_PIN`, `SPI1_SDI_PIN` |
| `NINA_SCK`        | `GPIO14`     |                   |
| `NINA_SDO`        | `GPIO11`     |                   |
| `NINA_SDI`        | `GPIO8`      | `NINA_RX`         |
| `NINA_CS`         | `GPIO9`      | `NINA_TX`         |
| `NINA_ACK`        | `GPIO10`     |                   |
| `NINA_GPIO0`      | `GPIO2`      |                   |
| `NINA_RESETN`     | `GPIO3`      |                   |
| `UART0_TX_PIN`    | `GPIO0`      | `UART_TX_PIN`     |
| `UART0_RX_PIN`    | `GPIO1`      | `UART_RX_PIN`     |

## Machine Package Docs

[Documentation for the machine package for the Nano RP2040](../machine/nano-rp2040)

## Flashing

### UF2

The Nano RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Nano RP2040 into your computer's USB port while shorting the pins REC and GND with a jumper wire.
- Once plugged in, remove the jumper pin.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=nano-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Nano RP2040 board should restart and run your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Nano RP2040 as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.
