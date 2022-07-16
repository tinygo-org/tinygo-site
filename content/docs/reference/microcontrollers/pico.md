---
title: "Raspberry Pi Pico"
weight: 3
---

The [Raspberry Pi Pico](https://www.raspberrypi.org/products/raspberry-pi-pico/) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | YES |
| USBDevice | YES | YES |

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

You can refer to [getting started with Raspberry Pi Pico](https://datasheets.raspberrypi.org/pico/getting-started-with-pico.pdf) documentation on how to connect two Picos together (see Appendix A: Using Picoprobe) to debug and convert `UART0` output on target pico to USB output on picoprobe. You will need the [Picoprobe UF2](https://www.raspberrypi.org/documentation/rp2040/getting-started/#board-specifications), available on the Pico's website under "About" tab.
