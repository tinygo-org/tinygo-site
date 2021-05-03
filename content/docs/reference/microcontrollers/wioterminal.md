---
title: "Seeed Wio Terminal"
weight: 3
---

The [Seeed Wio Terminal](https://www.seeedstudio.com/Wio-Terminal-p-4509.html) is a tiny ARM development board based on the Atmel [ATSAMD51P20](https://www.microchip.com/wwwproducts/en/ATSAMD51P20A) family of SoC.

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

[Documentation for the machine package for the Seeed Wio Terminal](../machine/wioterminal)

## Flashing

### UF2

The Wio Terminal comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Wio Terminal into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=wioterminal [PATH TO YOUR PROGRAM]
    ```

- The Wio Terminal board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Wio Terminal board to receive code, try this:

- [Slide the "RESET" switch on the board two times](https://wiki.seeedstudio.com/Wio-Terminal-Getting-Started/#faq) to get the Wio Terminal board ready to receive code.
- The Wio Terminal board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=wioterminal [PATH TO YOUR PROGRAM]
```

Once you have updated your Wio Terminal board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Wio Terminal as a serial port. `UART0` refers to this connection.
