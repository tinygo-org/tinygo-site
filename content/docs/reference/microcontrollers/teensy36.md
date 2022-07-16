---
title: "PJRC Teensy 3.6"
weight: 3
---

The [PJRC Teensy 3.6](https://www.pjrc.com/store/teensy36.html) is a small ARM development board based on the NXP [MK66FX1M0VMD18](https://www.nxp.com/docs/en/data-sheet/K66P144M180SF5V2.pdf) 32-bit 180 MHz ARM Cortex-M4.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | Not yet |
| I2C       | YES | Not yet |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the Teensy 3.6](../machine/teensy36)

## Flashing

### teensy_loader_cli

In order to flash your TinyGo programs onto the Teensy 3.6 board, you will need to install the `teensy_loader_cli` (https://github.com/PaulStoffregen/teensy_loader_cli) on your machine.

### CLI Flashing

- Plug your Teensy 3.6 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=teensy36 [PATH TO YOUR PROGRAM]
    ```

- The Teensy 3.6 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Teensy 3.6 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Teensy 3.6 board ready to receive code.
- The Teensy 3.6 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=teensy36 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Teensy 3.6 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You cannot yet use the USB port to the Teensy 3.6 as a USB CDC serial port.
