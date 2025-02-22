---
title: "Pimoroni Tufty2040"
weight: 3
---

The [Pimoroni Tufty2040](https://shop.pimoroni.com/products/tufty-2040) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `LED`             | `GPIO25`     | `USER_LED`        | `PWM4` (channel B)   |
| `BUTTON_A`        | `GPIO7`      |                   | `PWM3` (channel B)   |
| `BUTTON_B`        | `GPIO8`      |                   | `PWM4` (channel A)   |
| `BUTTON_C`        | `GPIO9`      |                   | `PWM4` (channel B)   |
| `BUTTON_UP`       | `GPIO22`     |                   | `PWM3` (channel A)   |
| `BUTTON_DOWN`     | `GPIO6`      |                   | `PWM3` (channel A)   |
| `BUTTON_USER`     | `GPIO23`     |                   | `PWM3` (channel B)   |
| `LCD_BACKLIGHT`   | `GPIO2`      |                   | `PWM1` (channel A)   |
| `LCD_CS`          | `GPIO10`     |                   | `PWM5` (channel A)   |
| `LCD_DC`          | `GPIO11`     |                   | `PWM5` (channel B)   |
| `LCD_WR`          | `GPIO12`     |                   | `PWM6` (channel A)   |
| `LCD_RD`          | `GPIO13`     |                   | `PWM6` (channel B)   |
| `LCD_DB0`         | `GPIO14`     |                   | `PWM7` (channel A)   |
| `LCD_DB1`         | `GPIO15`     |                   | `PWM7` (channel B)   |
| `LCD_DB2`         | `GPIO16`     |                   | `PWM0` (channel A)   |
| `LCD_DB3`         | `GPIO17`     |                   | `PWM0` (channel B)   |
| `LCD_DB4`         | `GPIO18`     |                   | `PWM1` (channel A)   |
| `LCD_DB5`         | `GPIO19`     |                   | `PWM1` (channel B)   |
| `LCD_DB6`         | `GPIO20`     |                   | `PWM2` (channel A)   |
| `LCD_DB7`         | `GPIO21`     |                   | `PWM2` (channel B)   |
| `VBUS_DETECT`     | `GPIO24`     |                   | `PWM4` (channel A)   |
| `BATTERY`         | `GPIO29`     | `ADC3`            | `PWM6` (channel B)   |
| `LIGHT_SENSE`     | `GPIO26`     | `ADC0`            | `PWM5` (channel A)   |
| `SENSOR_POWER`    | `GPIO27`     | `ADC1`            | `PWM5` (channel B)   |
| `I2C0_SDA_PIN`    | `GPIO4`      |                   | `PWM2` (channel A)   |
| `I2C0_SCL_PIN`    | `GPIO5`      |                   | `PWM2` (channel B)   |
| `UART0_TX_PIN`    | `GPIO0`      | `UART_TX_PIN`     | `PWM0` (channel A)   |
| `UART0_RX_PIN`    | `GPIO1`      | `UART_RX_PIN`     | `PWM0` (channel B)   |
| `ADC2`            | `GPIO28`     |                   | `PWM6` (channel A)   |

## Machine Package Docs

[Documentation for the machine package for the Pimoroni Tufty2040](../machine/tufty2040)

## Flashing

### UF2

The Tufty2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=tufty2040 [PATH TO YOUR PROGRAM]
    ```

- The Tufty2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Tufty2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)

For support for the Tufty's display please see [https://github.com/tinygo-org/pio/tree/main/rp2-pio/examples/tufty](https://github.com/tinygo-org/pio/tree/main/rp2-pio/examples/tufty)
