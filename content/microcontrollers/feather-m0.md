---
title: "Adafruit Feather M0"
weight: 3
---

The [Adafruit Feather M0](https://www.adafruit.com/product/3403) is a tiny ARM development board based on the Atmel [ATSAMD21G18](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of SoC.

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

[Documentation for the machine package for the Adafruit Feather M0](../machine/feather-m0)

## Flashing

### UF2

The Feather M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing on Linux

- Plug your Feather M0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m0 [PATH TO YOUR PROGRAM]
    ```

- The Feather M0 board should restart and then begin running your program.

### CLI Flashing on macOS

- Plug your Feather M0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m0 [PATH TO YOUR PROGRAM]
    ```

- The Feather M0 board should restart and then begin running your program.

### CLI Flashing on Windows

- Plug your Feather M0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m0 [PATH TO YOUR PROGRAM]
    ```

- The Feather M0 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Feather M0 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Feather M0 board ready to receive code.
- The Feather M0 board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=feather-m0 [PATH TO YOUR PROGRAM]
```

Once you have updated your Feather M0 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Feather M0 as a serial port. `UART0` refers to this connection.
