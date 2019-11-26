---
title: "Adafruit Circuit Playground Express"
weight: 3
---

The [Adafruit Circuit Playground Express](https://www.adafruit.com/product/3333) is small ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of processors. It has several built-in devices such as WS2812 "NeoPixel" LEDs, buttons, an accelerometer, and some other sensors.

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

[Documentation for the machine package for the Circuit Playground Express](../machine/circuitplay-express)

## Flashing

### UF2

The Circuit Playground Express comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing on Linux

- Plug your Circuit Playground Express into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=circuitplay-express [PATH TO YOUR PROGRAM]
    ```

- The Circuit Playground Express board should restart and then begin running your program.

### CLI Flashing on macOS

In order to talk to flash the board using macOS, you need to discover how macOS system has named the serial port.

- Plug your Circuit Playground Express into your computer's USB port.
- Run this command to display the connected USB devices:

    ```shell
    ls /dev | grep usb
    ```

    The above command should result in output like this:

    ```shell
    /dev/cu.usbmodem141201
    /dev/tty.usbmodem141201
    ```

- Using this information, you should now be able to flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=circuitplay-express -port=[PORT TO YOUR BOARD] [PATH TO YOUR PROGRAM]
    ```

    Replace `[PORT TO YOUR BOARD]` in the command above with the correct USB port name for your board.

- The Circuit Playground Express board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Circuit Playground Express board to receive code, try this:

- Press the "RESET" button on the board two times to get the Circuit Playground Express board ready to receive code.
- The Circuit Playground Express board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=circuitplay-express [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Circuit Playground Express board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Circuit Playground Express as a serial port. `UART0` refers to this connection.

For an example that uses the built-in Neopixel LEDs, take a look at the TinyGo drivers repository located at [https://github.com/tinygo-org/drivers/tree/master/examples](https://github.com/tinygo-org/drivers)
