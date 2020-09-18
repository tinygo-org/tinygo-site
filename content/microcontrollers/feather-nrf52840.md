---
title: "Adafruit Feather nRF52840"
weight: 3
---

The [Adafruit Feather nRF52840](https://www.adafruit.com/product/4500) is a small ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

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

[Documentation for the machine package for the Adafruit Feather nRF52840](../machine/feather-nrf52840)

## Flashing

### UF2

The Adafruit Feather nRF52840 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Adafruit Feather nRF52840 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-nrf52840 [PATH TO YOUR PROGRAM]
    ```

- The Adafruit Feather nRF52840 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Adafruit Feather nRF52840 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Adafruit Feather nRF52840 board ready to receive code.
- The Adafruit Feather nRF52840 board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=feather-nrf52840 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Adafruit Feather nRF52840 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Adafruit Feather nRF52840 as a serial port. `UART0` refers to this connection.

Bluetooth support is now available for the Adafruit Feather nRF52840 board. See https://github.com/tinygo-org/bluetooth for more information.
