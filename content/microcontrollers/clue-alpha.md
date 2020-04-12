---
title: "Adafruit CLUE"
weight: 3
---

The [Adafruit CLUE](https://www.adafruit.com/product/4500) is small ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor. It has several built-in devices such as WS2812 "NeoPixel" LEDs, buttons, an accelerometer, and some other sensors.

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

[Documentation for the machine package for the Adafruit CLUE](../machine/clue-alpha)

## Flashing

### UF2

The CLUE comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing on Linux

- Plug your CLUE into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=clue-alpha [PATH TO YOUR PROGRAM]
    ```

- The CLUE board should restart and then begin running your program.

### CLI Flashing on macOS

- Plug your CLUE into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=clue-alpha [PATH TO YOUR PROGRAM]
    ```

- The CLUE board should restart and then begin running your program.

### CLI Flashing on Windows

- Plug your CLUE into your computer's USB port.
- Double tap the "RESET" button on the board.
- Wait until the CLUE board appears as a flash drive.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=clue-alpha [PATH TO YOUR PROGRAM]
    ```

- The CLUE board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your CLUE board to receive code, try this:

- Press the "RESET" button on the board two times to get the CLUE board ready to receive code.
- The CLUE board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=clue-alpha [PATH TO YOUR PROGRAM]
    ```

Once you have updated your CLUE board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the CLUE as a serial port. `UART0` refers to this connection.

For an example that uses the built-in Neopixel LEDs, take a look at the TinyGo drivers repository located at [https://github.com/tinygo-org/drivers/tree/master/examples](https://github.com/tinygo-org/drivers)

Bluetooth support is in development but not yet completed.
