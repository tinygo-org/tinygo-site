---
title: "ESP32"
weight: 3
---

The [Espressif ESP32](https://www.espressif.com/en/products/socs/esp32) is a powerful chip that is used on many different development boards. It includes a built-in radio that can be used for WiFi or Bluetooth wireless connections.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | Not Yet |
| ADC      | YES | Not Yet |
| PWM      | YES | Not Yet |
| WiFi      | YES | Not Yet |
| Bluetooth      | YES | Not Yet |

## Machine Package Docs

[Documentation for the machine package for the ESP32](../machine/esp32)

## Flashing

### CLI Flashing on Linux

You need to install the Espressif toolchain for Linux to use TinyGo with the ESP32:

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/linux-setup.html#standard-setup-of-toolchain-for-linux

In addition, you must install the `esptool` flashing tool and clone the tinygo repo where you'll find the examples folder:

https://github.com/espressif/esptool#easy-installation

https://github.com/tinygo-org/tinygo

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the blinky1 example:

    ```
    tinygo flash -target=esp32-wroom-32 -port=/dev/ttyUSB0 examples/blinky1
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
    tinygo flash -target=esp32-wroom-32 examples/blinky1
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
    tinygo flash -target=esp32-wroom-32 examples/blinky1
    ```

- The ESP32 board should restart and then begin running your program.

### Troubleshooting

Goes here

## Notes

Goes here
