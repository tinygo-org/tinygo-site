---
title: "Pico"
weight: 3
---

The [Raspberry Pi Pico](https://www.raspberrypi.org/products/raspberry-pi-pico/) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | Not yet |
| I2C      | YES | Not yet |
| ADC      | YES | YES |
| PWM      | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the Pico](../machine/pico)

## Flashing

### UF2

The Pico comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Pico into your computer's USB port while holding down the RESET button on the board.
- One plugged in, release the RESET button.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pico [PATH TO YOUR PROGRAM]
    ```

- The Pico board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Pico as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.
