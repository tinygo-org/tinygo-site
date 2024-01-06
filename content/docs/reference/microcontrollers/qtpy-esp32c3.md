---
title: "Adafruit QT Py ESP32-C3"
weight: 3
---

The [QT Py ESP32-C3](https://www.adafruit.com/product/5405) is a development board based on the [Espressif ESP32-C3](https://www.espressif.com/en/products/socs/esp32-c3).

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
| `D0`              | `GPIO4`      | `A0`              |
| `D1`              | `GPIO3`      | `A1`              |
| `D2`              | `GPIO1`      | `A2`              |
| `D3`              | `GPIO0`      | `A3`              |
| `RX_PIN`          | `GPIO20`     | `UART_RX_PIN`     |
| `TX_PIN`          | `GPIO21`     | `UART_TX_PIN`     |
| `SDA_PIN`         | `GPIO5`      | `I2C0_SDA_PIN`    |
| `SCL_PIN`         | `GPIO6`      | `I2C0_SCL_PIN`    |
| `SCK_PIN`         | `GPIO10`     | `SPI_SCK_PIN`     |
| `MI_PIN`          | `GPIO8`      | `SPI_SDI_PIN`     |
| `MO_PIN`          | `GPIO7`      | `SPI_SDO_PIN`     |
| `NEOPIXEL`        | `GPIO2`      | `WS2812`          |
| `BUTTON`          | `GPIO9`      |                   |

## Machine Package Docs

[Documentation for the machine package for the QT Py ESP32-C3](../machine/qtpy-esp32c3)

## Flashing

### CLI Flashing

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32-C3 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32-C3 with the serial example:

    ```shell
    tinygo flash -target=qtpy-esp32c3 examples/serial
    ```

- The ESP32-C3 board should restart and then begin running your program.
