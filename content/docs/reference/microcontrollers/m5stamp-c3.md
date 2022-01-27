---
title: "M5Stamp C3"
weight: 3
---

The [M5Stamp-C3](https://docs.m5stack.com/en/core/stamp_c3) is a development board based on the Espressif ESP32-C3 a powerful chip that is used on many different board mostly because of the built-in radio that can be used for WiFi or Bluetooth wireless connections.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | Not yet |
| UART      | YES | YES |
| SPI      | YES | Not yet |
| I2C      | YES | Not yet |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |
| WiFi      | YES | Not Yet |
| Bluetooth      | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the M5Stamp-C3](../machine/m5stamp-c3)

## Flashing

### CLI Flashing

In addition, you must install the `esptool` flashing tool:

https://github.com/espressif/esptool#easy-installation

Now you should be able to flash your board as follows:

- Plug your ESP32-C3 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the ESP32-C3 with the serial example:

    ```
    tinygo flash -target=m5stack-core2 examples/serial
    ```

- The ESP32-C3 board should restart and then begin running your program.

### Troubleshooting

Goes here

## Notes

Goes here
