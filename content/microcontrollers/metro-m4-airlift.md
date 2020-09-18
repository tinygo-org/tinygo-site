---
title: "Adafruit Metro M4 Express AirLift"
weight: 3
---

The [Adafruit Metro M4 Express AirLift](https://www.adafruit.com/product/4000) is an ARM development board based on the Atmel [ATSAMD51J19](https://www.microchip.com/wwwproducts/en/ATSAMD51J19) family of SoC that has Arduino shield compatible form factor, and a built-in EspressIf ESP32 Wi-Fi Co processor.

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

[Documentation for the machine package for the Adafruit Metro M4 Airlift](../machine/metro-m4-airlift)

## Flashing

### UF2

The Metro M4 Express comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Metro M4 Express into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=metro-m4-airlift [PATH TO YOUR PROGRAM]
    ```

- The Metro M4 Express board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Metro M4 Express board to receive code, try this:

- Press the "RESET" button on the board two times to get the Metro M4 Express board ready to receive code.
- The Metro M4 Express board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=metro-m4-airlift [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Metro M4 Express board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Metro M4 Express as a serial port. `UART0` refers to this connection.
