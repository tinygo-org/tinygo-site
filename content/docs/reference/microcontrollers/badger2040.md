---
title: "Pimoroni Badger2040"
weight: 3
---

The [Pimoroni Badger2040](https://shop.pimoroni.com/products/badger-2040) is a badge with E Ink display based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller.

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
| `LED`             | `GPIO25`     |                   | `PWM4` (channel B)   |
| `BUTTON_A`        | `GPIO12`     |                   | `PWM6` (channel A)   |
| `BUTTON_B`        | `GPIO13`     |                   | `PWM6` (channel B)   |
| `BUTTON_C`        | `GPIO14`     |                   | `PWM7` (channel A)   |
| `BUTTON_UP`       | `GPIO15`     |                   | `PWM7` (channel B)   |
| `BUTTON_DOWN`     | `GPIO11`     |                   | `PWM5` (channel B)   |
| `BUTTON_USER`     | `GPIO23`     |                   | `PWM3` (channel B)   |
| `EPD_BUSY_PIN`    | `GPIO26`     | `ADC0`            | `PWM5` (channel A)   |
| `EPD_RESET_PIN`   | `GPIO21`     |                   | `PWM2` (channel B)   |
| `EPD_DC_PIN`      | `GPIO20`     |                   | `PWM2` (channel A)   |
| `EPD_CS_PIN`      | `GPIO17`     |                   | `PWM0` (channel B)   |
| `EPD_SCK_PIN`     | `GPIO18`     | `SPI0_SCK_PIN`    | `PWM1` (channel A)   |
| `EPD_SDO_PIN`     | `GPIO19`     | `SPI0_SDO_PIN`    | `PWM1` (channel B)   |
| `VBUS_DETECT`     | `GPIO24`     |                   | `PWM4` (channel A)   |
| `VREF_POWER`      | `GPIO27`     | `ADC1`            | `PWM5` (channel B)   |
| `VREF_1V24`       | `GPIO28`     | `ADC2`            | `PWM6` (channel A)   |
| `VBAT_SENSE`      | `GPIO29`     | `BATTERY`, `ADC3` | `PWM6` (channel B)   |
| `ENABLE_3V3`      | `GPIO10`     |                   | `PWM5` (channel A)   |
| `I2C0_SDA_PIN`    | `GPIO4`      |                   | `PWM2` (channel A)   |
| `I2C0_SCL_PIN`    | `GPIO5`      |                   | `PWM2` (channel B)   |
| `SPI0_SDI_PIN`    | `GPIO16`     |                   | `PWM0` (channel A)   |
| `UART0_TX_PIN`    | `GPIO0`      | `UART_TX_PIN`     | `PWM0` (channel A)   |
| `UART0_RX_PIN`    | `GPIO1`      | `UART_RX_PIN`     | `PWM0` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for the Pimoroni Badger2040](../machine/badger2040)

## Flashing

### UF2

The Badger2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=badger2040 [PATH TO YOUR PROGRAM]
    ```

- The Badger2040 board should restart and then begin running your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the Badger2040 as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
