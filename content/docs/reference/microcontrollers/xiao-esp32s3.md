---
title: "Seeed Studio XIAO ESP32S3"
weight: 3
---

The [Seeed Studio XIAO ESP32S3](https://wiki.seeedstudio.com/xiao_esp32s3_getting_started/) is a tiny development board based on the Espressif [Xtensa LX7 dual-core, 32-bit processor ](https://www.espressif.com/en/products/socs/esp32-s3/) microcontroller.

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
| WiFi      | YES | Not Yet |
| Bluetooth | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `GPIO1`      | `A0`              |
| `D1`              | `GPIO2`      | `A1`              |
| `D2`              | `GPIO3`      | `A2`              |
| `D3`              | `GPIO4`      | `A3`              |
| `D4`              | `GPIO5`      | `SDA_PIN`         |
| `D5`              | `GPIO6`      | `SCL_PIN`         |
| `D6`              | `GPIO43`     | `UART_TX_PIN`     |
| `D7`              | `GPIO44`     | `UART_RX_PIN`     |
| `D8`              | `GPIO7`      | `SPI1_SCK_PIN`    |
| `D9`              | `GPIO8`      | `SPI1_MISO_PIN`   |
| `D10`             | `GPIO9`      | `SPI1_MOSI_PIN`   |
| `LED`             | `GPIO21`     |                   |

## Machine Package Docs

[Documentation for the machine package for the Seeed Studio XIAO ESP32S3](../machine/xiao-esp32s3)

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

Any additional notes go here.
