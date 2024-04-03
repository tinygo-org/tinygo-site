---
title: "Seeed XIAO ESP32 C3"
weight: 3
---

The [XIAO ESP32 C3](https://www.seeedstudio.com/Seeed-XIAO-ESP32C3-p-5431.html) is a development board based on the [Espressif ESP32-C3](https://www.espressif.com/en/products/socs/esp32-c3).

## Peripherals and Drivers

A powerful chip that is used on many different boards mostly because of the built-in radio that can be used for WiFi or Bluetooth wireless connections.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |
| WiFi      | YES | Not yet |
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

[Documentation for the machine package for the XIAO ESP32 C3](../machine/xiao-esp32c3)

## Flashing

### CLI Flashing

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32-C3 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32-C3 with the serial example:

    ```shell
    tinygo flash -target=xiao-esp32c3 examples/serial
    ```

- The ESP32-C3 board should restart and then begin running your program.
