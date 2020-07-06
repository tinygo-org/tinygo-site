---
title: "Adafruit Feather M4"
weight: 3
---

The [Adafruit Feather M4](https://www.adafruit.com/product/3857) is a tiny ARM development board based on the Atmel [ATSAMD51J19](https://www.microchip.com/wwwproducts/en/ATSAMD51J19A) family of SoC.

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

[Documentation for the machine package for the Adafruit Feather M4](../machine/feather-m4)

## Flashing

### UF2

The Feather M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing on Linux

- Plug your Feather M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m4 [PATH TO YOUR PROGRAM]
    ```

- The Feather M4 board should restart and then begin running your program.

### CLI Flashing on macOS

- Plug your Feather M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m4 [PATH TO YOUR PROGRAM]
    ```

- The Feather M4 board should restart and then begin running your program.

### CLI Flashing on Windows

- Plug your Feather M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m4 [PATH TO YOUR PROGRAM]
    ```

- The Feather M4 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Feather M4 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Feather M4 board ready to receive code.
- The Feather M4 board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=feather-m4 [PATH TO YOUR PROGRAM]
```

Once you have updated your Feather M4 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Feather M4 as a serial port. `UART0` refers to this connection.
