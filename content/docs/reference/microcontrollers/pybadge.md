---
title: "Adafruit PyBadge"
weight: 3
---

The [Adafruit PyBadge](https://www.adafruit.com/product/4200) is a ARM development board based on the Atmel [ATSAMD51J19A](https://www.microchip.com/wwwproducts/en/ATSAMD51J19A) family of SoC.

It has many built-in devices, such as a 1.8" 160x128 Color TFT Display, 8 x buttons, 5 x NeoPixels, a triple-axis accelerometer, a light sensor, and a speaker. The PyBadge uses the ST7735 display, so you may use the tinygo-org st7735 driver. The accelerometer is the LIS3DH so you may use the tinygo lis3dh driver

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

[Documentation for the machine package for the Adafruit PyBadge](../machine/pybadge)

## Flashing

### UF2

The PyBadge comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your PyBadge into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pybadge [PATH TO YOUR PROGRAM]
    ```

- The PyBadge board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your PyBadge board to receive code, try this:

- Press the "RESET" button on the board two times to get the PyBadge board ready to receive code.
- The PyBadge board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=pybadge [PATH TO YOUR PROGRAM]
```

Once you have updated your PyBadge board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the PyBadge as a serial port. `UART0` refers to this connection.
