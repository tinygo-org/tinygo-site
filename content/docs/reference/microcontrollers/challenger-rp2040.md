---
title: "iLabs Challenger RP2040 LoRa"
weight: 3
---

The [iLabs Challenger RP2040 LoRa](https://ilabs.se/product/challenger-rp2040-lora/) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `LED`             | `GPIO24`     | `I2C0_SDA_PIN`    | `PWM4` (channel A)   |
| `D5`              | `GPIO2`      | `I2C1_SDA_PIN`, `SDA_PIN` | `PWM1` (channel A)   |
| `D6`              | `GPIO3`      | `I2C1_SCL_PIN`, `SCL_PIN` | `PWM1` (channel B)   |
| `D9`              | `GPIO4`      | `UART1_TX_PIN`    | `PWM2` (channel A)   |
| `D10`             | `GPIO5`      | `UART1_RX_PIN`    | `PWM2` (channel B)   |
| `D11`             | `GPIO6`      |                   | `PWM3` (channel A)   |
| `D12`             | `GPIO7`      |                   | `PWM3` (channel B)   |
| `D13`             | `GPIO8`      |                   | `PWM4` (channel A)   |
| `A0`              | `GPIO26`     | `ADC0`            | `PWM5` (channel A)   |
| `A1`              | `GPIO27`     | `ADC1`            | `PWM5` (channel B)   |
| `A2`              | `GPIO28`     | `ADC2`            | `PWM6` (channel A)   |
| `A3`              | `GPIO29`     | `ADC3`            | `PWM6` (channel B)   |
| `I2C0_SCL_PIN`    | `GPIO25`     |                   | `PWM4` (channel B)   |
| `SPI0_SCK_PIN`    | `GPIO22`     |                   | `PWM3` (channel A)   |
| `SPI0_SDO_PIN`    | `GPIO23`     |                   | `PWM3` (channel B)   |
| `SPI0_SDI_PIN`    | `GPIO20`     |                   | `PWM2` (channel A)   |
| `SPI1_SCK_PIN`    | `GPIO10`     | `LORA_SCK`        | `PWM5` (channel A)   |
| `SPI1_SDO_PIN`    | `GPIO11`     | `LORA_SDO`        | `PWM5` (channel B)   |
| `SPI1_SDI_PIN`    | `GPIO12`     | `LORA_SDI`        | `PWM6` (channel A)   |
| `LORA_CS`         | `GPIO9`      |                   | `PWM4` (channel B)   |
| `LORA_RESET`      | `GPIO13`     |                   | `PWM6` (channel B)   |
| `LORA_DIO0`       | `GPIO14`     |                   | `PWM7` (channel A)   |
| `LORA_DIO1`       | `GPIO15`     |                   | `PWM7` (channel B)   |
| `LORA_DIO2`       | `GPIO18`     |                   | `PWM1` (channel A)   |
| `UART0_TX_PIN`    | `GPIO16`     | `UART_TX_PIN`     | `PWM0` (channel A)   |
| `UART0_RX_PIN`    | `GPIO17`     | `UART_RX_PIN`     | `PWM0` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the iLabs Challenger RP2040 LoRa](../machine/challenger-rp2040)

## Flashing

### UF2

The Challenger RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=challenger-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Challenger RP2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Challenger RP2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
