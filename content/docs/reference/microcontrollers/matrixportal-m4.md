---
title: "Adafruit Matrix Portal M4"
weight: 3
---

The [Adafruit Matrix Portal M4](https://www.adafruit.com/product/4745) is an ARM development system based on the [ATSAMD51J19 Cortex M4 processor](https://www.microchip.com/wwwproducts/en/ATSAMD51J19). The Adafruit Matrix Portal M4 is designed to plugin easily to any HUB-75 LED display. In addition it has a built-in ESP32 Wi-Fi coprocessor, along with a LIS3DH accelerometer.

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

[Documentation for the machine package for the Adafruit Matrix Portal M4](../machine/matrix-portal-m4)

## Flashing

### UF2

The Adafruit Matrix Portal M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Matrix Portal M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=matrixportal-m4 [PATH TO YOUR PROGRAM]
    ```

- The Matrix Portal M4 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Matrix Portal M4 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Matrix Portal M4 board ready to receive code.
- The Matrix Portal M4 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=matrixportal-m4 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Matrix Portal M4 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Matrix Portal M4 as a serial port. `UART0` refers to this connection.
