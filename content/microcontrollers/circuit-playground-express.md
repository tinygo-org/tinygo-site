---
title: "Circuit Playground Express"
weight: 3
---

The [Adafruit Circuit Playground Express](https://www.adafruit.com/product/3333) is small ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of processors. It has several built-in devices such as WS2812 "NeoPixel" LEDs, buttons, an accelerometer, and some other sensors.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | Not yet |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | YES |

## Flashing

### UF2

The Circuit Playground Express comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

- Plug your Circuit Playground Express into your computer's USB port.
- Build your TinyGo program to `.bin` format using the `tinygo build -target=circuitplay-express` command.
- Press the "RESET" button on the board two times.
- Convert/flash the resulting `.bin` using the [`uf2conv.py`](https://github.com/Microsoft/uf2/blob/master/utils/uf2conv.py) program.

## Notes

You can use the USB port to the Circuit Playground Express as a serial port. `UART0` refers to this connection.

For an example that uses the built-in Neopixel LEDs, take a look at the TinyGo drivers repository located at [https://github.com/tinygo-org/drivers/tree/master/examples](https://github.com/tinygo-org/drivers)
