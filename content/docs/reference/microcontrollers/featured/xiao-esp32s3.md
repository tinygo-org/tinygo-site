---
title: "Seeed Studio XIAO ESP32S3"
weight: 2
---

The [Seeed Studio XIAO ESP32S3](https://wiki.seeedstudio.com/xiao_esp32s3_getting_started/) is a tiny development board based on the Espressif [Xtensa LX7 dual-core, 32-bit processor ](https://www.espressif.com/en/products/socs/esp32-s3/) microcontroller with 512KB of SRAM and 8MB of onboard flash storage. Thanks to the onboard radio, it supports both WiFi and Bluetooth wireless communication.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | YES |
| USBDevice | YES | Not yet |
| WiFi      | YES | YES |
| Bluetooth | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `GPIO1`      | `A0`, `ADC0`      |
| `D1`              | `GPIO2`      | `A1`, `ADC2`      |
| `D2`              | `GPIO3`      | `A2`, `ADC3`      |
| `D3`              | `GPIO4`      | `A3`, `ADC4`      |
| `D4`              | `GPIO5`      | `SDA_PIN`, `ADC5` |
| `D5`              | `GPIO6`      | `SCL_PIN`, `ADC6` |
| `D6`              | `GPIO43`     | `UART_TX_PIN`     |
| `D7`              | `GPIO44`     | `UART_RX_PIN`     |
| `D8`              | `GPIO7`      | `SPI1_SCK_PIN`, `ADC7` |
| `D9`              | `GPIO8`      | `SPI1_MISO_PIN`, `ADC8` |
| `D10`             | `GPIO9`      | `SPI1_MOSI_PIN`, `ADC9` |
| `LED`             | `GPIO21`     |                   |
| `ADC10`           | `GPIO10`     |                   |
| `ADC11`           | `GPIO11`     |                   |
| `ADC12`           | `GPIO12`     |                   |
| `ADC13`           | `GPIO13`     |                   |
| `ADC14`           | `GPIO14`     |                   |
| `ADC15`           | `GPIO15`     |                   |
| `ADC16`           | `GPIO16`     |                   |
| `ADC17`           | `GPIO17`     |                   |
| `ADC18`           | `GPIO18`     |                   |
| `ADC19`           | `GPIO19`     |                   |
| `ADC20`           | `GPIO20`     |                   |

## Machine Package Docs

[Documentation for the machine package for the Seeed Studio XIAO ESP32S3](../../machine/xiao-esp32s3)

## Flashing

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=xiao-esp32s3 [PATH TO YOUR PROGRAM]
    ```

- The XIAO ESP32S3 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

### WiFi

You can use the onboard wireless chip for WiFi using the `espradio` package.

For more information, see [https://github.com/tinygo-org/espradio](https://github.com/tinygo-org/espradio)
