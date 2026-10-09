---
title: "ESP32-H2-DevKitM-1"
weight: 3
---

The [ESP32-H2-DevKitM-1](https://docs.espressif.com/projects/esp-dev-kits/en/latest/esp32h2/esp32-h2-devkitm-1/user_guide.html) is a development board based on the Espressif [ESP32-H2](https://www.espressif.com/en/products/socs/esp32-h2) 32-bit RISC-V microcontroller. It has an onboard radio for Bluetooth Low Energy and IEEE 802.15.4.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | Not yet |
| I2C       | YES | Not yet |
| ADC       | YES | YES |
| PWM       | YES | Not yet |
| Bluetooth | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `WS2812`          | `GPIO8`      |                   |
| `BUTTON`          | `GPIO9`      |                   |
| `UART_TX_PIN`     | `GPIO24`     |                   |
| `UART_RX_PIN`     | `GPIO23`     |                   |

## Machine Package Docs

[Documentation for the machine package for the ESP32-H2-DevKitM-1](../../machine/esp32h2-devkitm-1)

## Flashing

### CLI Flashing

- Plug your ESP32-H2-DevKitM-1 board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the board with the serial example:

    ```shell
    tinygo flash -target=esp32h2-devkitm-1 examples/serial
    ```

- The ESP32-H2-DevKitM-1 board should restart and then begin running your program.

## Notes

The board has an addressable RGB LED on `GPIO8` named `WS2812`. It has no plain LED, so `examples/blinky1` does not work on this board.
