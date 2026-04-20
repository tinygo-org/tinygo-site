---
title: "ESP8266 - d1mini"
weight: 3
---

The [Espressif ESP8266](https://www.espressif.com/en/products/socs/esp8266) d1mini is a very small yet powerful SoC that is usually used for WiFi applications thanks to its built-in radio.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | Not yet |
| I2C       | NO (software only) | Not yet |
| ADC       | YES | YES |
| PWM       | YES | Not yet |
| USBDevice | NO  | NO  |
| WiFi      | YES | Not Yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `GPIO16`     |                   |
| `D1`              | `GPIO5`      | `SCL_PIN`         |
| `D2`              | `GPIO4`      | `SDA_PIN`         |
| `D3`              | `GPIO0`      |                   |
| `D4`              | `GPIO2`      | `LED`             |
| `D5`              | `GPIO14`     | `SPI0_SCK_PIN`    |
| `D6`              | `GPIO12`     | `SPI0_SDI_PIN`    |
| `D7`              | `GPIO13`     | `SPI0_SDO_PIN`    |
| `D8`              | `GPIO15`     | `SPI0_CS0_PIN`    |
| `UART_TX_PIN`     | `GPIO1`      |                   |
| `UART_RX_PIN`     | `GPIO3`      |                   |

## Machine Package Docs

[Documentation for the machine package for the ESP8266 d1mini](../../machine/d1mini)

## Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=d1mini [PATH TO YOUR PROGRAM]
    ```

- The d1Mini board should restart and then begin running your program.


### Troubleshooting

Goes here

## Notes

Goes here
