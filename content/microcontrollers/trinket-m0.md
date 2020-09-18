---
title: "Adafruit Trinket M0"
weight: 3
---

The [Adafruit Trinket M0](https://www.adafruit.com/product/3500) is a tiny ARM development board based on the Atmel [ATSAMD21E18](https://www.microchip.com/wwwproducts/en/ATSAMD21E18) family of SoC.

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

[Documentation for the machine package for the Adafruit Trinket M0](../machine/trinket-m0)

## Flashing

### UF2

The Trinket M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Trinket M0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=trinket-m0 [PATH TO YOUR PROGRAM]
    ```

- The Trinket M0 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Trinket M0 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Trinket M0 board ready to receive code.
- The Trinket M0 board will appear to your computer like a USB drive.
- Now try running the `tinygo flash` command as above:

    ```shell
    tinygo flash -target=Trinket-m0 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Trinket M0 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Trinket M0 as a serial port. `UART0` refers to this connection.
