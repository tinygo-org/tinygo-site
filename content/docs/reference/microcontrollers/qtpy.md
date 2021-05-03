---
title: "Adafruit QtPy"
weight: 3
---

The [Adafruit QtPy](https://www.adafruit.com/product/4600) is a tiny ARM development board based on the Atmel [ATSAMD21E18](https://www.microchip.com/wwwproducts/en/ATSAMD21E18) family of SoC.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Machine Package Docs

[Documentation for the machine package for the Adafruit QtPy](../machine/qtpy)

## Flashing

### UF2

The QtPy comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your QtPy into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=qtpy [PATH TO YOUR PROGRAM]
    ```

- The QtPy board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your QtPy board to receive code, try this:

- Press the "RESET" button on the board two times to get the QtPy board ready to receive code.
- The QtPy board will appear to your computer like a USB drive.
- Now try running the `tinygo flash` command as above:

    ```shell
    tinygo flash -target=qtpy [PATH TO YOUR PROGRAM]
    ```

Once you have updated your QtPy board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the QtPy as a serial port. `UART0` refers to this connection.
