---
title: "Microchip SAM E54 Xplained Pro"
weight: 3
---

The [Microchip SAM E54 Xplained Pro](https://www.microchip.com/developmenttools/productdetails/atsame54-xpro) is a tiny ARM development board based on the Atmel [ATSAME54P20](https://www.microchip.com/wwwproducts/en/ATSAME54P20A) family of SoC.

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

[Documentation for the machine package for the Adafruit SAM E54 Xplained Pro](../machine/atsame54-xpro)

## Flashing

### OpenOCD

Programs are loaded onto the SAM E54 Xplained Pro with the onboard [Atmel EDBG](http://ww1.microchip.com/downloads/en/devicedoc/atmel-42096-microcontrollers-embedded-debugger_user-guide.pdf) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the SAM E54 Xplained Pro board with your TinyGo code.

- Plug your SAM E54 Xplained Pro into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=atsame54-xpro [PATH TO YOUR PROGRAM]`

## Notes

You can use the USB port to the SAM E54 Xplained Pro as a serial port. `UART0` refers to this connection.
