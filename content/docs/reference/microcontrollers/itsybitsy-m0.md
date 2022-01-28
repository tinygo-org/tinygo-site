---
title: "Adafruit ItsyBitsy M0"
weight: 3
---

The [Adafruit ItsyBitsy M0](https://www.adafruit.com/product/3727) is very compact ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of SoC.

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

[Documentation for the machine package for the Adafruit ItsyBitsy M0](../machine/itsybitsy-m0)

## Flashing

### UF2

The ItsyBitsy M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your ItsyBitsy M0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=itsybitsy-m0 [PATH TO YOUR PROGRAM]
    ```

- The ItsyBitsy M0 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your ItsyBitsy M0 board to receive code, try this:

- Press the "RESET" button on the board two times to get the ItsyBitsy M0 board ready to receive code.
- The ItsyBitsy M0 board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=itsybitsy-m0 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your ItsyBitsy M0 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the ItsyBitsy M0 as a serial port. `UART0` refers to this connection.

The DotStar LED on the ItsyBitsy M0 can be accessed using the [APA102](https://pkg.go.dev/tinygo.org/x/drivers/apa102) driver via a software SPI on the `PA00` and `PA01` pins
