---
title: "Seeed XIAO BLE"
weight: 3
---

The [Seeed XIAO BLE](https://www.seeedstudio.com/Seeed-XIAO-BLE-nRF52840-p-5201.html) is a tiny ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

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

## Machine Package Docs

[Documentation for the machine package for the Seeed XIAO BLE](../machine/xiao-ble)

## Flashing

### UF2

The XIAO BLE comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your XIAO BLE into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=xiao-ble [PATH TO YOUR PROGRAM]
    ```

- The XIAO BLE board should restart and then begin running your program.

### Troubleshooting

Add troubleshooting tips here.

## Notes

You can use the USB port to the XIAO BLE as a serial port. `UART0` refers to this connection.

