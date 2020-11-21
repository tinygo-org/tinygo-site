---
title: "ESP8266 - NodeMCU"
weight: 3
---

The [Espressif ESP8266](https://www.espressif.com/en/products/socs/esp8266) NodeMCU is a small yet powerful SoC that is usually used for WiFi applications thanks to its built-in radio.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | Not Yet |
| I2C      | YES | Not Yet |
| ADC      | YES | Not Yet |
| PWM      | YES | Not Yet |
| WiFi      | YES | Not Yet |

## Machine Package Docs

[Documentation for the machine package for the ESP8266 NodeMCU](../machine/nodemcu)

## Flashing

### CLI Flashing on Linux

You need to install the same toolchain for the ESP8266 as is used for the ESP32 to use the ESP8266 with TinyGo: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/linux-setup.html#standard-setup-of-toolchain-for-linux

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP8266 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP8266 with the blinky1 example:

    ```
    tinygo flash -target=nodemcu -port=/dev/ttyUSB examples/blinky1
    ```

- The ESP8266 board should restart and then begin running your program.

### CLI Flashing on macOS

You need to install the same toolchain for the ESP8266 as is used for the ESP32 to use the ESP8266 with TinyGo:

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/macos-setup.html

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP8266 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP8266 with the blinky1 example:

    ```
    tinygo flash -target=nodemcu examples/blinky1
    ```

- The ESP826 board should restart and then begin running your program.

### CLI Flashing on Windows

You need to install the same toolchain for the ESP8266 as is used for the ESP32 to use the ESP8266 with TinyGo:

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/windows-setup.html

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP826 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP826 with the blinky1 example:

    ```
    tinygo flash -target=nodemcu examples/blinky1
    ```

- The ESP826 board should restart and then begin running your program.

### Troubleshooting

Goes here

## Notes

Goes here
