---
title: "Nano RP2040"
weight: 3
---

The [Nano RP2040]() is a tiny development board based on the Raspberry Pi [RP2040]() microcontroller.

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

[Documentation for the machine package for the Nano RP2040](../machine/nano-rp2040)

## Flashing

### UF2

The Nano RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Nano RP2040 into your computer's USB port while holding down the RESET button on the board.
- One plugged in, release the RESET button.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=nano-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Nano RP2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Nano RP2040 as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.
