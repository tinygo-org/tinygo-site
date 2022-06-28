---
title: "ESP32 - mini32"
weight: 3
---

The mini32 is a small development board based on the popular [Espressif ESP32](https://www.espressif.com/en/products/socs/esp32). The ESP32 includes a built-in radio that can be used for WiFi or Bluetooth wireless connections.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | Not yet |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |
| WiFi      | YES | Not Yet |
| Bluetooth      | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `CLK`             | `GPIO6`      |                   |
| `CMD`             | `GPIO11`     |                   |
| `IO0`             | `GPIO0`      | `PWM1_PIN`        |
| `IO1`             | `GPIO1`      | `TXD`, `UART_TX_PIN` |
| `IO2`             | `GPIO2`      | `LED`, `PWM0_PIN` |
| `IO3`             | `GPIO3`      | `RXD`, `UART_RX_PIN` |
| `IO4`             | `GPIO4`      | `PWM2_PIN`        |
| `IO5`             | `GPIO5`      | `SPI0_CS0_PIN`    |
| `IO9`             | `GPIO9`      | `SD2`, `UART1_TX_PIN` |
| `IO10`            | `GPIO10`     | `SD3`, `UART1_RX_PIN` |
| `IO16`            | `GPIO16`     |                   |
| `IO17`            | `GPIO17`     |                   |
| `IO18`            | `GPIO18`     | `SPI0_SCK_PIN`    |
| `IO19`            | `GPIO19`     | `SPI0_SDI_PIN`    |
| `IO21`            | `GPIO21`     | `SDA_PIN`         |
| `IO22`            | `GPIO22`     | `SCL_PIN`         |
| `IO23`            | `GPIO23`     | `SPI0_SDO_PIN`    |
| `IO25`            | `GPIO25`     |                   |
| `IO26`            | `GPIO26`     |                   |
| `IO27`            | `GPIO27`     |                   |
| `IO32`            | `GPIO32`     |                   |
| `IO33`            | `GPIO33`     |                   |
| `IO34`            | `GPIO34`     | `ADC0`            |
| `IO35`            | `GPIO35`     | `ADC1`            |
| `IO36`            | `GPIO36`     | `SVP`, `ADC2`     |
| `IO39`            | `GPIO39`     | `SVN`, `ADC3`     |
| `SD0`             | `GPIO7`      |                   |
| `SD1`             | `GPIO8`      |                   |
| `TCK`             | `GPIO13`     |                   |
| `TD0`             | `GPIO15`     |                   |
| `TDI`             | `GPIO12`     |                   |
| `TMS`             | `GPIO14`     |                   |

## Machine Package Docs

[Documentation for the machine package for the ESP32-mini32](../machine/esp32-mini32)

## Flashing

### CLI Flashing on Linux

You need to install the Espressif toolchain for Linux to use TinyGo with the ESP32: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/linux-setup.html#standard-setup-of-toolchain-for-linux

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the blinky1 example:

    ```
    tinygo flash -target=esp32-mini32 -port=/dev/ttyUSB0 examples/blinky1
    ```

- The ESP32 board should restart and then begin running your program.

### CLI Flashing on macOS

You need to install the Espressif toolchain for macOS to use TinyGo with the ESP32: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/macos-setup.html

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the blinky1 example:

    ```
    tinygo flash -target=esp32-mini32 examples/blinky1
    ```

- The ESP32 board should restart and then begin running your program.

### CLI Flashing on Windows

You need to install the Espressif toolchain for Windows to use TinyGo with the ESP32: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/windows-setup.html

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the blinky1 example:

    ```
    tinygo flash -target=esp32-mini32 examples/blinky1
    ```

- The ESP32 board should restart and then begin running your program.

### Troubleshooting

Goes here

## Notes

Goes here
