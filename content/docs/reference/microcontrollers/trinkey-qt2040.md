---
title: "Adafruit Trinkey QT2040"
weight: 3
---

The [Adafruit Trinkey QT2040](https://www.adafruit.com/product/5056) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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

## Pins

| Pin               | Hardware pin | Alternative names | PWM                  |
| ----------------- | ------------ | ----------------- | -------------------- |
| `NEOPIXEL`        | `GPIO27`     | `WS2812`, `ADC1`  | `PWM5` (channel B)   |
| `I2C0_SDA_PIN`    | `GPIO16`     |                   | `PWM0` (channel A)   |
| `I2C0_SCL_PIN`    | `GPIO17`     |                   | `PWM0` (channel B)   |
| `ADC0`            | `GPIO26`     |                   | `PWM5` (channel A)   |
| `ADC2`            | `GPIO28`     |                   | `PWM6` (channel A)   |
| `ADC3`            | `GPIO29`     |                   | `PWM6` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Trinkey QT2040](../machine/trinkey-qt2040)

## Flashing

### UF2

The Trinkey QT2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=trinkey-qt2040 [PATH TO YOUR PROGRAM]
    ```

- The Trinkey QT2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Trinkey QT2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
