---
title: "M5Stack"
weight: 3
---

The [m5stack](https://docs.m5stack.com/en/core/basic) is a development board based on the Espressif ESP32 a powerful chip that is used on many different board mostly because of the built-in radio that can be used for WiFi or Bluetooth wireless connections.

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
| `IO0`             | `GPIO0`      |                   |
| `IO1`             | `GPIO1`      | `UART0_TX_PIN`, `UART_TX_PIN` |
| `IO2`             | `GPIO2`      |                   |
| `IO3`             | `GPIO3`      | `UART0_RX_PIN`, `UART_RX_PIN` |
| `IO4`             | `GPIO4`      | `SDCARD_SS_PIN`   |
| `IO5`             | `GPIO5`      |                   |
| `IO6`             | `GPIO6`      |                   |
| `IO7`             | `GPIO7`      |                   |
| `IO8`             | `GPIO8`      |                   |
| `IO9`             | `GPIO9`      |                   |
| `IO10`            | `GPIO10`     |                   |
| `IO11`            | `GPIO11`     |                   |
| `IO12`            | `GPIO12`     |                   |
| `IO13`            | `GPIO13`     |                   |
| `IO14`            | `GPIO14`     | `SPI0_CS0_PIN`, `LCD_SS_PIN` |
| `IO15`            | `GPIO15`     |                   |
| `IO16`            | `GPIO16`     | `UART1_RX_PIN`    |
| `IO17`            | `GPIO17`     | `UART1_TX_PIN`    |
| `IO18`            | `GPIO18`     | `SPI0_SCK_PIN`, `LCD_SCK_PIN`, `SDCARD_SCK_PIN` |
| `IO19`            | `GPIO19`     | `SPI0_SDI_PIN`, `LCD_SDI_PIN`, `SDCARD_SDI_PIN` |
| `IO21`            | `GPIO21`     | `SDA0_PIN`, `SDA_PIN` |
| `IO22`            | `GPIO22`     | `SCL0_PIN`, `SCL_PIN` |
| `IO23`            | `GPIO23`     | `SPI0_SDO_PIN`, `LCD_SDO_PIN`, `SDCARD_SDO_PIN` |
| `IO25`            | `GPIO25`     | `SPEAKER_PIN`, `DAC1` |
| `IO26`            | `GPIO26`     | `DAC2`            |
| `IO27`            | `GPIO27`     | `LCD_DC_PIN`      |
| `IO32`            | `GPIO32`     | `LCD_BL_PIN`      |
| `IO33`            | `GPIO33`     | `LCD_RST_PIN`     |
| `IO34`            | `GPIO34`     |                   |
| `IO35`            | `GPIO35`     | `ADC1`            |
| `IO36`            | `GPIO36`     | `ADC2`            |
| `IO37`            | `GPIO37`     | `BUTTON_C`        |
| `IO38`            | `GPIO38`     | `BUTTON_B`        |
| `IO39`            | `GPIO39`     | `BUTTON_A`, `BUTTON` |

## Machine Package Docs

[Documentation for the machine package for the M5Stack Core2](../machine/m5stack)

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
    tinygo flash -target=m5stack -port=/dev/ttyUSB0 examples/serial
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
    tinygo flash -target=m5stack examples/serial
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
    tinygo flash -target=m5stack examples/serial
    ```

- The ESP32 board should restart and then begin running your program.

### Troubleshooting

Goes here

## Notes

Goes here
