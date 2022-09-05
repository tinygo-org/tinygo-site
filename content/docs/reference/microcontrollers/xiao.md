---
title: "Seeed Seeeduino XIAO"
weight: 3
---

The [Seeed Seeeduino XIAO](https://www.seeedstudio.com/Seeeduino-XIAO-Arduino-Microcontroller-SAMD21-Cortex-M0+-p-4426.html) is a tiny ARM development board based on the Atmel [ATSAMD21G18](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of SoC.

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
| `D0`              | `PA02`       | `A0`              |                      |
| `D1`              | `PA04`       | `A1`              | `TCC0` (channel 0)   |
| `D2`              | `PA10`       | `A2`, `I2S_SCK_PIN` | `TCC1` (channel 0), `TCC0` (channel 2) |
| `D3`              | `PA11`       | `A3`              | `TCC1` (channel 1), `TCC0` (channel 3) |
| `D4`              | `PA08`       | `A4`, `SDA_PIN`, `I2S_SD_PIN` | `TCC0` (channel 0), `TCC1` (channel 2) |
| `D5`              | `PA09`       | `A5`, `SCL_PIN`   | `TCC0` (channel 1), `TCC1` (channel 3) |
| `D6`              | `PB08`       | `A6`, `UART_TX_PIN` |                      |
| `D7`              | `PB09`       | `A7`, `UART_RX_PIN` |                      |
| `D8`              | `PA07`       | `A8`, `SPI0_SCK_PIN` | `TCC1` (channel 1)   |
| `D9`              | `PA05`       | `A9`, `SPI0_SDI_PIN` | `TCC0` (channel 1)   |
| `D10`             | `PA06`       | `A10`, `SPI0_SDO_PIN` | `TCC1` (channel 0)   |
| `LED`             | `PA17`       |                   | `TCC2` (channel 1), `TCC0` (channel 3) |
| `LED_RXL`         | `PA18`       | `LED2`            | `TCC0` (channel 2)   |
| `LED_TXL`         | `PA19`       | `LED3`            | `TCC0` (channel 3)   |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC1` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC1` (channel 3)   |

## Machine Package Docs

[Documentation for the machine package for the Seeed Seeeduino XIAO](../machine/xiao)

## Flashing

### UF2

The XIAO comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your XIAO into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=xiao [PATH TO YOUR PROGRAM]
    ```

- The XIAO board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your XIAO board to receive code, try this:

- [Short the "RST pins" two times](https://wiki.seeedstudio.com/Seeeduino-XIAO/#reset) to get the XIAO board ready to receive code.
- The XIAO board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=xiao [PATH TO YOUR PROGRAM]
```

Once you have updated your XIAO board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the XIAO as a serial port. `UART0` refers to this connection.

