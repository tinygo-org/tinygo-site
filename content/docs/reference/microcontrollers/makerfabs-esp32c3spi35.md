---
title: "Makerfabs ESP32-C3 SPI"
weight: 3
---

The [ESP32-C3 SPI](https://wiki.makerfabs.com/ESP32_C3_SPI_3.5_TFT_with_Touch.html) is a development board based on the [Espressif ESP32-C3](https://www.espressif.com/en/products/socs/esp32-c3).

## Peripherals and Drivers

- WiFi
- Bluetooth
- ili9488: 3.5inch TFT LCD, 480*320, RGB, via SPI
- FT6236: Capacitive Touch Panel
- MicroSD

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

## Machine Package Docs

[Documentation for the machine package for the ESP32-C3 SPI](../machine/makerfabs-esp32c3spi35)

## Flashing

### CLI Flashing

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32-C3 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32-C3 with the serial example:

    ```shell
    tinygo flash -target=makerfabs-esp32c3spi35 examples/serial
    ```

- The ESP32-C3 board should restart and then begin running your program.
