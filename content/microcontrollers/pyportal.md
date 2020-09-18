---
title: "Adafruit PyPortal"
weight: 3
---

The [Adafruit PyPortal](https://www.adafruit.com/product/4116) is a ARM board based on the Microchip [ATSAMD51J20A](https://www.microchip.com/wwwproducts/en/ATSAMD51J20A) family of SoC.

The PyPortal also has an Espressif ESP32 Wi-Fi coprocessor with TLS/SSL support built-in. PyPortal has a 3.2″ 320 x 240 color TFT with resistive touch screen. PyPortal includes: speaker, light sensor, temperature sensor, NeoPixel, microSD card slot, 8MB flash, plug-in ports for I2C and 2 analog/digital pins,

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

[Documentation for the machine package for the Adafruit PyPortal](../machine/pyportal)

## Flashing

### UF2

The PyPortal comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your PyPortal into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pyportal [PATH TO YOUR PROGRAM]
    ```

- The PyPortal board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your PyPortal board to receive code, try this:

- Press the "RESET" button on the board two times to get the PyPortal board ready to receive code.
- The PyPortal board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=pyportal [PATH TO YOUR PROGRAM]
```

Once you have updated your PyPortal board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the PyPortal as a serial port. `UART0` refers to this connection.
