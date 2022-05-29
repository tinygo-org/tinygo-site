---
title: "Sparkfun Thing Plus RP2040"
weight: 3
---

The [Sparkfun Thing Plus RP2040](https://www.sparkfun.com/products/17745) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller. 

Peripherals: 
- ws2812 Neopixel
- sdcard

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

[Documentation for the machine package for the Sparkfun Thing Plus RP2040](../machine/thingplus-rp2040)

## Flashing

### UF2

The Sparkfun Thing Plus RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Thing Plus RP2040 into your computer's USB port while pressing the BOOT button
- Release BOOT Button
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=thingplus-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Thing Plus RP2040 board should restart and run your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Thing Plus RP2040 as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.

The Neopixel LED and the SD Card are supported by [tinygo drivers](https://github.com/tinygo-org/drivers)