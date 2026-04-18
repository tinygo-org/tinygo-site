---
title: "Seeed Studio XIAO ESP32C3"
weight: 2
image: /images/boards/xiao-esp32c3.svg
---

![Seeed Studio XIAO ESP32C3](/images/boards/xiao-esp32c3.svg)

The [Seeed Studio XIAO ESP32C3](https://wiki.seeedstudio.com/XIAO_ESP32C3_Getting_Started/) is a tiny development board based on the Espressif [ESP32C3](https://www.espressif.com/en/products/socs/esp32-c3) 32-bit RISC-V CPU microcontroller with 400KB of SRAM and 4MB of onboard flash storage. Thanks to the onboard radio, it supports both WiFi and Bluetooth wireless communication.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | YES |
| USBDevice | NO  | NO  |
| WiFi      | YES | YES |
| Bluetooth | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `GPIO2`      | `A0`, `ADC2`      |
| `D1`              | `GPIO3`      | `A1`, `ADC3`      |
| `D2`              | `GPIO4`      | `A2`, `ADC4`      |
| `D3`              | `GPIO5`      | `A3`, `ADC5`      |
| `D4`              | `GPIO6`      | `SDA_PIN`         |
| `D5`              | `GPIO7`      | `SCL_PIN`         |
| `D6`              | `GPIO21`     | `UART_TX_PIN`     |
| `D7`              | `GPIO20`     | `UART_RX_PIN`     |
| `D8`              | `GPIO8`      | `SPI_SCK_PIN`     |
| `D9`              | `GPIO9`      | `SPI_SDI_PIN`     |
| `D10`             | `GPIO10`     | `SPI_SDO_PIN`     |
| `ADC0`            | `GPIO0`      |                   |
| `ADC1`            | `GPIO1`      |                   |

## Machine Package Docs

[Documentation for the machine package for the Seeed Studio XIAO ESP32C3](../../machine/xiao-esp32c3)

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

### WiFi

You can use the onboard wireless chip for WiFi using the `espradio` package.

For more information, see [https://github.com/tinygo-org/espradio](https://github.com/tinygo-org/espradio)
