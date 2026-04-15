---
title: "Seeed Studio XIAO ESP32C3"
weight: 3
---

The [Seeed Studio XIAO ESP32C3](https://wiki.seeedstudio.com/XIAO_ESP32C3_Getting_Started/) is a tiny development board based on the Espressif [ESP32C3](https://www.espressif.com/en/products/socs/esp32-c3) 32-bit RISC-V CPU microcontroller.

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
| `D0`              | `GPIO2`      | `A0`              |
| `D1`              | `GPIO3`      | `A1`              |
| `D2`              | `GPIO4`      | `A2`              |
| `D3`              | `GPIO5`      | `A3`              |
| `D4`              | `GPIO6`      | `SDA_PIN`         |
| `D5`              | `GPIO7`      | `SCL_PIN`         |
| `D6`              | `GPIO21`     | `UART_TX_PIN`     |
| `D7`              | `GPIO20`     | `UART_RX_PIN`     |
| `D8`              | `GPIO8`      | `SPI_SCK_PIN`     |
| `D9`              | `GPIO9`      | `SPI_SDI_PIN`     |
| `D10`             | `GPIO10`     | `SPI_SDO_PIN`     |

## Machine Package Docs

[Documentation for the machine package for the Seeed Studio XIAO ESP32C3](../machine/xiao-esp32c3)

## Flashing

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=xiao-esp32c3 [PATH TO YOUR PROGRAM]
    ```

- The XIAO ESP32C3 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

Any additional notes go here.
