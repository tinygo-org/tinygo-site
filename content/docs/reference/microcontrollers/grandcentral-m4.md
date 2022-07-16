---
title: "Adafruit Grand Central M4"
weight: 3
---

The [Adafruit Grand Central M4](https://www.adafruit.com/product/4064) is a tiny ARM development board based on the Atmel [ATSAMD51P20](https://www.microchip.com/wwwproducts/en/ATSAMD51P20A) family of SoC.

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

[Documentation for the machine package for the Adafruit Grand Central M4](../machine/grandcentral-m4)

## Flashing

### UF2

The Grand Central M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Grand Central M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=grandcentral-m4 [PATH TO YOUR PROGRAM]
    ```

- The Grand Central M4 board should restart and then begin running your program.


### Troubleshooting

If you have troubles getting your Grand Central M4 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Grand Central M4 board ready to receive code.
- The Grand Central M4 board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=grandcentral-m4 [PATH TO YOUR PROGRAM]
```

Once you have updated your Grand Central M4 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Grand Central M4 as a serial port. `UART0` refers to this connection.

The Neopixel LED on the Grand Central M4 can be accessed using the [WS2812](https://pkg.go.dev/tinygo.org/x/drivers/ws2812) driver via the `PC24` pin
