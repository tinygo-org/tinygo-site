---
title: "Makerdiary nRF52840-MDK"
weight: 3
---

The [nRF52840-MDK](https://wiki.makerdiary.com/nrf52840-mdk/) (not to be confused with its sibling, the [nRF52840-MDK-USB-Dongle](https://wiki.makerdiary.com/nrf52840-mdk-usb-dongle/)) is an open-source, micro development kit for IoT applications based on the Nordic Semiconductor [nRF52840](https://www.nordicsemi.com/eng/Products/nRF52840) SoC chip.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |
| Bluetooth      | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `LED_GREEN`       | `P0_22`      | `LED`             |
| `LED_RED`         | `P0_23`      |                   |
| `LED_BLUE`        | `P0_24`      |                   |
| `UART_TX_PIN`     | `P0_20`      |                   |
| `UART_RX_PIN`     | `P0_19`      |                   |

## Machine Package Docs

[Documentation for the machine package for the nRF52840-MDK](../machine/nrf52840-mdk)

## Flashing

There are two options to flash the nRF52840-MDK board.

### OpenOCD (Recommended)

You must install [OpenOCD](http://openocd.org/) before you will be able to flash the nRF52840-MDK board with your TinyGo code.
You should check [OpenOCD Documentation](http://openocd.org/Documentation) for installation and configuration instructions.

Once you have installed it correctly, you will be able to flash the nRF52840-MDK board with your TinyGo code.

- Plug your nRF52840-MDK into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nrf52840-mdk`

### nrfjprog/J-Link

Programs can be loaded onto the nRF52840-MDK board using the `nrfjprog` command line utility program.

First install the J-Link Software and Documentation Pack from Segger: [https://www.segger.com/downloads/jlink/#J-LinkSoftwareAndDocumentationPack](https://www.segger.com/downloads/jlink/#J-LinkSoftwareAndDocumentationPack)

Then install the nRF5x Command-Line Tools: [https://docs.zephyrproject.org/latest/guides/tools/nordic_segger.html#nrf5x-command-line-tools-installation](https://docs.zephyrproject.org/latest/guides/tools/nordic_segger.html#nrf5x-command-line-tools-installation)

Once you have installed both of these correctly, you will be able to flash the nRF52840-MDK board with your TinyGo code.

- Plug your nRF52840-MDK into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nrf52840-mdk -programmer command`

## Notes

Bluetooth support is now available for nRF52840 boards. See https://github.com/tinygo-org/bluetooth for more information.
