---
title: "Adafruit ItsyBitsy M4"
weight: 3
---

The [Adafruit ItsyBitsy M4](https://www.adafruit.com/product/3800) is very compact ARM development board based on the Atmel [SAMD51](https://www.microchip.com/wwwproducts/en/ATSAMD51G19A) family of SoC.

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

[Documentation for the machine package for the Adafruit ItsyBitsy M4](../machine/itsybitsy-m4)

## Flashing

### UF2

The ItsyBitsy M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your ItsyBitsy M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=itsybitsy-m4 [PATH TO YOUR PROGRAM]
    ```

- The ItsyBitsy M4 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your ItsyBitsy M4 board to receive code, try this:

- Press the "RESET" button on the board two times to get the ItsyBitsy M4 board ready to receive code.
- The ItsyBitsy M4 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=itsybitsy-m4 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your ItsyBitsy M4 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the ItsyBitsy M4 as a serial port. `UART0` refers to this connection.
