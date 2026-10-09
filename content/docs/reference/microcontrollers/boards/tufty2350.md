---
title: "Pimoroni Tufty 2350"
weight: 3
---

The [Pimoroni Tufty 2350](https://shop.pimoroni.com/products/tufty-2350) is a development board based on the Raspberry Pi [RP2350B](https://datasheets.raspberrypi.org/rp2350/rp2350-datasheet.pdf) microcontroller. It also includes an onboard [Infineon CYW43439](https://www.infineon.com/cms/en/product/wireless-connectivity/airoc-wi-fi-plus-bluetooth-combos/wi-fi-4-802.11n/cyw43439/) wireless chip.

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
| WiFi      | YES | YES |
| Bluetooth | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names | I2C                  | PWM                  |
| ----------------- | ------------ | ----------------- | -------------------- | -------------------- |
| `CL0`             | `GPIO0`      | `LED`, `UART0_TX_PIN`, `UART_TX_PIN` | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `CL1`             | `GPIO1`      | `UART0_RX_PIN`, `UART_RX_PIN` | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `CL2`             | `GPIO2`      |                   | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `CL3`             | `GPIO3`      |                   | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `BUTTON_DOWN`     | `GPIO6`      |                   | `I2C1` (SDA)         | `PWM3` (channel A)   |
| `BUTTON_A`        | `GPIO7`      |                   | `I2C1` (SCL)         | `PWM3` (channel B)   |
| `BUTTON_B`        | `GPIO9`      |                   | `I2C0` (SCL)         | `PWM4` (channel B)   |
| `BUTTON_C`        | `GPIO10`     |                   | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `BUTTON_UP`       | `GPIO11`     |                   | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `BUTTON_HOME`     | `GPIO22`     |                   |                      | `PWM3` (channel A)   |
| `BUTTON_RESET`    | `GPIO14`     |                   | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `BUTTON_INT`      | `GPIO15`     |                   | `I2C1` (SCL)         | `PWM7` (channel B)   |
| `VBUS_DETECT`     | `GPIO12`     |                   | `I2C0` (SDA)         | `PWM6` (channel A)   |
| `RTC_ALARM`       | `GPIO13`     |                   | `I2C0` (SCL)         | `PWM6` (channel B)   |
| `LCD_BACKLIGHT`   | `GPIO26`     |                   | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `LCD_CS`          | `GPIO27`     |                   | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `LCD_DC`          | `GPIO28`     |                   |                      | `PWM6` (channel A)   |
| `LCD_WR`          | `GPIO30`     |                   | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `LCD_RD`          | `GPIO31`     |                   | `I2C1` (SCL)         | `PWM7` (channel B)   |
| `LCD_DB0`         | `GPIO32`     |                   | `I2C0` (SDA)         | `PWM8` (channel A)   |
| `LCD_DB1`         | `GPIO33`     |                   | `I2C0` (SCL)         | `PWM8` (channel B)   |
| `LCD_DB2`         | `GPIO34`     |                   | `I2C1` (SDA)         | `PWM9` (channel A)   |
| `LCD_DB3`         | `GPIO35`     |                   | `I2C1` (SCL)         | `PWM9` (channel B)   |
| `LCD_DB4`         | `GPIO36`     |                   | `I2C0` (SDA)         | `PWM10` (channel A)  |
| `LCD_DB5`         | `GPIO37`     |                   | `I2C0` (SCL)         | `PWM10` (channel B)  |
| `LCD_DB6`         | `GPIO38`     |                   | `I2C1` (SDA)         | `PWM11` (channel A)  |
| `LCD_DB7`         | `GPIO39`     |                   | `I2C1` (SCL)         | `PWM11` (channel B)  |
| `VBAT_SENSE`      | `GPIO40`     | `ADC0`            | `I2C0` (SDA)         | `PWM8` (channel A)   |
| `POWER_EN`        | `GPIO41`     | `ADC1`            | `I2C0` (SCL)         | `PWM8` (channel B)   |
| `SENSE_1V1`       | `GPIO42`     | `ADC2`            | `I2C1` (SDA)         | `PWM9` (channel A)   |
| `LIGHT_SENSE`     | `GPIO43`     | `ADC3`            | `I2C1` (SCL)         | `PWM9` (channel B)   |
| `WL_REG_ON`       | `GPIO23`     |                   |                      | `PWM3` (channel B)   |
| `WL_DATA`         | `GPIO24`     |                   |                      | `PWM4` (channel A)   |
| `WL_CLOCK`        | `GPIO29`     |                   |                      | `PWM6` (channel B)   |
| `WL_CS`           | `GPIO25`     |                   |                      | `PWM4` (channel B)   |
| `I2C0_SDA_PIN`    | `GPIO4`      |                   | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `I2C0_SCL_PIN`    | `GPIO5`      |                   | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `ADC4`            | `GPIO44`     |                   | `I2C0` (SDA)         | `PWM10` (channel A)  |
| `ADC5`            | `GPIO45`     |                   | `I2C0` (SCL)         | `PWM10` (channel B)  |
| `ADC6`            | `GPIO46`     |                   | `I2C1` (SDA)         | `PWM11` (channel A)  |
| `ADC7`            | `GPIO47`     |                   | `I2C1` (SCL)         | `PWM11` (channel B)  |

## Machine Package Docs

[Documentation for the machine package for the Pimoroni Tufty 2350](../../machine/tufty2350)

## Flashing

### UF2

The Pimoroni Tufty 2350 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=tufty2350 [PATH TO YOUR PROGRAM]
    ```

- The Pimoroni Tufty 2350 board should restart and then begin running your program.

## Notes

You can use the USB port to the Pimoroni Tufty 2350 as a serial port.

TinyGo has support for the RP2350's on-board Programmable Input/Output (PIO) block.

For more information, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)

### WiFi

You can use the onboard wireless chip for WiFi using the CYW43439 package.

For more information, see [https://github.com/soypat/cyw43439](https://github.com/soypat/cyw43439)

### Bluetooth

You can also use the onboard wireless chip for Bluetooth Low Energy using the TinyGo Bluetooth package.

For more information, see [https://github.com/tinygo-org/bluetooth](https://github.com/tinygo-org/bluetooth)
