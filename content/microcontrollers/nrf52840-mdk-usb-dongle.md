---
title: "nRF52840-MDK-USB-Dongle"
weight: 3
---

The [nRF52840 MDK USB Dongle](https://wiki.makerdiary.com/nrf52840-mdk-usb-dongle/) (not to be confused with its sibling, the [nRF52840-MDK](https://wiki.makerdiary.com/nrf52840-mdk/)) is an open-source, micro development kit dongle for IoT applications based on the Nordic Semiconductor [nRF52840](https://www.nordicsemi.com/eng/Products/nRF52840) SoC chip.

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

[Documentation for the machine package for the nRF52840-MDK-USB-Dongle](../machine/nrf52840-mdk-usb-dongle)

## Flashing

### UF2

The nRF52840 MDK USB Dongle currently ships with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed. However, previously (pre mid-2020) it was also shipped with Open Bootloader. To use TinyGo with the nRF52840 MDK USB Dongle, please [upgrade to the UF2 bootloader](https://wiki.makerdiary.com/nrf52840-mdk-usb-dongle/programming/#upgrade-to-uf2-bootloader-from-open-bootloader).

### SoftDevice

TinyGo uses the S140 v6 [SoftDevice](https://infocenter.nordicsemi.com/topic/ug_gsg_ses/UG/gsg/softdevices.html) (proprietary firmware) for Bluetooth support and requires it to be installed before you can run TinyGo programs.

Some C++ examples in the [nRF52840 MDK USB Dongle repository](https://github.com/makerdiary/nrf52840-mdk-usb-dongle), like the [C++ blinky](https://github.com/makerdiary/nrf52840-mdk-usb-dongle/tree/master/examples/nrf5-sdk/blinky) example, do not use a SoftDevice and when flashed will be placed at the location where the SoftDevice is usually located, overwriting it in the process.

To install the S140 v6 SoftDevice, please follow the SoftDevice part of the ["Running examples that use a SoftDevice"](https://wiki.makerdiary.com/nrf52840-mdk-usb-dongle/nrf5-sdk/#running-examples-that-use-a-softdevice) instructions on the nRF52840 MDK USB Dongle wiki.

### CLI Flashing

- Enter DFU mode by holding the dongle's RESET/USR button while plugging it into your computer's USB port. A flash drive with the name MDK-DONGLE will appear.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=nrf52840-mdk-usb-dongle [PATH TO YOUR PROGRAM]
    ```

## Notes

You can use the USB port of the nRF52840 MDK USB Dongle as a serial port. `UART0` refers to this connection.

The button (pin 18) on the nRF52840 MDK USB Dongle can be to be configured as a reset-button or GPIO pin through the `PSELRESET[0]/[1]` UICR registers. When it is configured as reset instead of as GPIO, you won't be able to use the button in your code as pressing it will reset the nrf52840. To set the `PSELRESET[0]/[1]` UICR registers back to their default value (disabled), please flash the [pselreset erase example](https://github.com/makerdiary/nrf52840-mdk-usb-dongle/tree/master/examples/nrf5-sdk/pselreset_erase) using [these instructions](https://wiki.makerdiary.com/nrf52840-mdk-usb-dongle/programming/#dfu-via-uf2-bootloader) on the nRF52840 MDK USB Dongle wiki.

Bluetooth support is now available for nRF52840 boards. See https://github.com/tinygo-org/bluetooth for more information.
