---
title: "nRF52840-MDK"
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
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |
| Bluetooth      | YES | YES |

## Machine Package Docs

[Documentation for the machine package for the nRF52840-MDK](../machine/nrf52840-mdk)

## Flashing

### OpenOCD

Programs are loaded onto the nRF52840-MDK board using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the nRF52840-MDK board with your TinyGo code.

- Plug your nRF52840-MDK into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nrf52840-mdk`

## Notes

Bluetooth support is now available for nRF52840 boards. See https://github.com/tinygo-org/bluetooth for more information.
