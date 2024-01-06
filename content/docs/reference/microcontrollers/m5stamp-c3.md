---
title: "M5Stack M5Stamp-C3"
weight: 3
---

The [M5Stamp-C3](https://docs.m5stack.com/en/core/stamp_c3) is a development board based on the [Espressif ESP32-C3](https://www.espressif.com/en/products/socs/esp32-c3).

## Peripherals and Drivers

A powerful chip that is used on many different boards mostly because of the built-in radio that can be used for WiFi or Bluetooth wireless connections.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | Not yet |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |
| WiFi      | YES | Not Yet |
| Bluetooth | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `IO0`             | `GPIO0`      | `XTAL_32K_P`      |
| `IO1`             | `GPIO1`      | `XTAL_32K_N`      |
| `IO2`             | `GPIO2`      | `WS2812`          |
| `IO3`             | `GPIO3`      |                   |
| `IO4`             | `GPIO4`      | `MTMS`            |
| `IO5`             | `GPIO5`      | `MTDI`            |
| `IO6`             | `GPIO6`      | `MTCK`            |
| `IO7`             | `GPIO7`      | `MTDO`            |
| `IO8`             | `GPIO8`      |                   |
| `IO9`             | `GPIO9`      |                   |
| `IO10`            | `GPIO10`     |                   |
| `IO11`            | `GPIO11`     | `VDD_SPI`         |
| `IO12`            | `GPIO12`     | `SPIHD`           |
| `IO13`            | `GPIO13`     | `SPISP`           |
| `IO14`            | `GPIO14`     | `SPICS0`          |
| `IO15`            | `GPIO15`     | `SPICLK`          |
| `IO16`            | `GPIO16`     | `SPID`            |
| `IO17`            | `GPIO17`     | `SPIQ`            |
| `IO18`            | `GPIO18`     |                   |
| `IO19`            | `GPIO19`     |                   |
| `IO20`            | `GPIO20`     | `U0RXD`, `UART_RX_PIN` |
| `IO21`            | `GPIO21`     | `U0TXD`, `UART_TX_PIN` |

## Machine Package Docs

[Documentation for the machine package for the M5Stamp-C3](../machine/m5stamp-c3)

## Flashing

### CLI Flashing

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32-C3 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32-C3 with the serial example:

    ```shell
    tinygo flash -target=m5stack-core2 examples/serial
    ```

- The ESP32-C3 board should restart and then begin running your program.
