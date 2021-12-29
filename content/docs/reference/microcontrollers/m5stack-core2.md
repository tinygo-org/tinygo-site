---
title: "M5Stack Core2"
weight: 3
---

The [m5stack-core2](https://shop.m5stack.com/products/m5stack-core2-esp32-iot-development-kit) is a development board based on the Espressif ESP32 a powerful chip that is used on many different board mostly because of the built-in radio that can be used for WiFi or Bluetooth wireless connections.

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

## Machine Package Docs

[Documentation for the machine package for the M5Stack Core2](../machine/m5stack-core2)

## Flashing

### CLI Flashing on Linux

You need to install the Espressif toolchain for Linux to use TinyGo with the ESP32: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/linux-setup.html#standard-setup-of-toolchain-for-linux

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the serial example:

    ```
    tinygo flash -target=m5stack-core2 -port=/dev/ttyUSB0 examples/serial
    ```

- The ESP32 board should restart and then begin running your program.

### CLI Flashing on macOS

You need to install the Espressif toolchain for macOS to use TinyGo with the ESP32: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/macos-setup.html

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the serial example:

    ```
    tinygo flash -target=m5stack-core2 examples/serial
    ```

- The ESP32 board should restart and then begin running your program.

### CLI Flashing on Windows

You need to install the Espressif toolchain for Windows to use TinyGo with the ESP32: 

https://docs.espressif.com/projects/esp-idf/en/release-v3.0/get-started/windows-setup.html

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32 with the serial example:

    ```
    tinygo flash -target=m5stack-core2 examples/serial
    ```

- The ESP32 board should restart and then begin running your program.

### Troubleshooting

Goes here

## Notes

Goes here
