---
title: "Teensy 4.0"
weight: 3
---

The [PJRC Teensy 4.0](https://www.pjrc.com/store/teensy40.html) is a small ARM development board based on the NXP [iMXRT1062](https://www.nxp.com/docs/en/nxp/data-sheets/IMXRT1060CEC.pdf) 32-bit 600 MHz ARM Cortex-M7.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | Not yet |
| I2C      | YES | Not yet |
| ADC      | YES | Not yet |
| PWM      | YES | Not Yet |

## Machine Package Docs

[Documentation for the machine package for the Teensy 4.0](../machine/teensy40)

## Flashing

### teensy_loader_cli

In order to flash your TinyGo programs onto the Teensy 4.0 board, you will need to install the `teensy_loader_cli` (https://github.com/PaulStoffregen/teensy_loader_cli) on your machine.

### CLI Flashing

- Plug your Teensy 4.0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=teensy36 [PATH TO YOUR PROGRAM]
    ```

- The Teensy 4.0 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Teensy 4.0 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Teensy 4.0 board ready to receive code.
- The Teensy 4.0 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=teensy40 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Teensy 4.0 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You cannot yet use the USB port to the Teensy 4.0 as a USB CDC serial port.
