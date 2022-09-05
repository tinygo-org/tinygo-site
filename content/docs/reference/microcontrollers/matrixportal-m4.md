---
title: "Adafruit Matrix Portal M4"
weight: 3
---

The [Adafruit Matrix Portal M4](https://www.adafruit.com/product/4745) is an ARM development system based on the [ATSAMD51J19 Cortex M4 processor](https://www.microchip.com/wwwproducts/en/ATSAMD51J19). The Adafruit Matrix Portal M4 is designed to plugin easily to any HUB-75 LED display. In addition it has a built-in ESP32 Wi-Fi coprocessor, along with a LIS3DH accelerometer.

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
| `D0`              | `PA01`       | `UART1_RX_PIN`, `UART_RX_PIN` |                      |
| `D1`              | `PA00`       | `UART1_TX_PIN`, `UART_TX_PIN` |                      |
| `D2`              | `PB22`       | `BUTTON_UP`       |                      |
| `D3`              | `PB23`       | `BUTTON_DOWN`     |                      |
| `D4`              | `PA23`       | `NEOPIXEL`, `WS2812` | `TCC1` (channel 7), `TCC0` (channel 3) |
| `D5`              | `PB31`       | `I2C0_SDA_PIN`, `I2C_SDA_PIN`, `SDA_PIN` | `TCC4` (channel 1), `TCC0` (channel 7) |
| `D6`              | `PB30`       | `I2C0_SCL_PIN`, `I2C_SCL_PIN`, `SCL_PIN` | `TCC4` (channel 0), `TCC0` (channel 6) |
| `D7`              | `PB00`       | `HUB75_R1`        |                      |
| `D8`              | `PB01`       | `HUB75_G1`        |                      |
| `D9`              | `PB02`       | `HUB75_B1`        | `TCC2` (channel 2)   |
| `D10`             | `PB03`       | `HUB75_R2`        | `TCC2` (channel 3)   |
| `D11`             | `PB04`       | `HUB75_G2`        |                      |
| `D12`             | `PB05`       | `HUB75_B2`        |                      |
| `D13`             | `PA14`       | `LED`             | `TCC2` (channel 0), `TCC1` (channel 2) |
| `D14`             | `PB06`       | `HUB75_CLK`       |                      |
| `D15`             | `PB14`       | `HUB75_LAT`       | `TCC4` (channel 0), `TCC0` (channel 2) |
| `D16`             | `PB12`       | `HUB75_OE`        | `TCC3` (channel 0), `TCC0` (channel 0) |
| `D17`             | `PB07`       | `HUB75_ADDR_A`    |                      |
| `D18`             | `PB08`       | `HUB75_ADDR_B`    |                      |
| `D19`             | `PB09`       | `HUB75_ADDR_C`    |                      |
| `D20`             | `PB15`       | `HUB75_ADDR_D`    | `TCC4` (channel 1), `TCC0` (channel 3) |
| `D21`             | `PB13`       | `HUB75_ADDR_E`    | `TCC3` (channel 1), `TCC0` (channel 1) |
| `D22`             | `PA02`       | `A0`              |                      |
| `D23`             | `PA05`       | `D48`, `A1`, `SPI1_SCK_PIN` |                      |
| `D24`             | `PA04`       | `D49`, `A2`, `SPI1_SDO_PIN` |                      |
| `D25`             | `PA06`       | `A3`              |                      |
| `D26`             | `PA07`       | `D50`, `A4`, `SPI1_SDI_PIN` |                      |
| `D27`             | `PA12`       | `UART2_RX_PIN`, `NINA_RX` | `TCC0` (channel 6), `TCC1` (channel 2) |
| `D28`             | `PA13`       | `UART2_TX_PIN`, `NINA_TX` | `TCC0` (channel 7), `TCC1` (channel 3) |
| `D29`             | `PA20`       | `NINA_GPIO0`      | `TCC1` (channel 4), `TCC0` (channel 0) |
| `D30`             | `PA21`       | `NINA_RESETN`     | `TCC1` (channel 5), `TCC0` (channel 1) |
| `D31`             | `PA22`       | `NINA_ACK`        | `TCC1` (channel 6), `TCC0` (channel 2) |
| `D32`             | `PA18`       | `NINA_RTS`        | `TCC1` (channel 2), `TCC0` (channel 6) |
| `D33`             | `PB17`       | `NINA_CS`         | `TCC3` (channel 1), `TCC0` (channel 5) |
| `D34`             | `PA16`       | `SPI0_SCK_PIN`, `SPI_SCK_PIN`, `NINA_SCK` | `TCC1` (channel 0), `TCC0` (channel 4) |
| `D35`             | `PA17`       | `SPI0_SDI_PIN`, `SPI_SDI_PIN`, `NINA_SDI` | `TCC1` (channel 1), `TCC0` (channel 5) |
| `D36`             | `PA19`       | `SPI0_SDO_PIN`, `SPI_SDO_PIN`, `NINA_SDO` | `TCC1` (channel 3), `TCC0` (channel 7) |
| `D38`             | `PA24`       | `USBCDC_DM_PIN`, `UART0_RX_PIN` | `TCC2` (channel 2)   |
| `D39`             | `PA25`       | `USBCDC_DP_PIN`, `UART0_TX_PIN` | `TCC2` (channel 3)   |
| `D40`             | `PA03`       |                   |                      |
| `D41`             | `PB10`       | `QSPI_SCK`        | `TCC0` (channel 4), `TCC1` (channel 0) |
| `D42`             | `PB11`       | `QSPI_CS`         | `TCC0` (channel 5), `TCC1` (channel 1) |
| `D43`             | `PA08`       | `QSPI_DATA0`      | `TCC0` (channel 0), `TCC1` (channel 4) |
| `D44`             | `PA09`       | `QSPI_DATA1`      | `TCC0` (channel 1), `TCC1` (channel 5) |
| `D45`             | `PA10`       | `QSPI_DATA2`      | `TCC0` (channel 2), `TCC1` (channel 6) |
| `D46`             | `PA11`       | `QSPI_DATA3`      | `TCC0` (channel 3), `TCC1` (channel 7) |
| `D47`             | `PA27`       |                   |                      |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Matrix Portal M4](../machine/matrix-portal-m4)

## Flashing

### UF2

The Adafruit Matrix Portal M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Matrix Portal M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=matrixportal-m4 [PATH TO YOUR PROGRAM]
    ```

- The Matrix Portal M4 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Matrix Portal M4 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Matrix Portal M4 board ready to receive code.
- The Matrix Portal M4 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=matrixportal-m4 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Matrix Portal M4 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Matrix Portal M4 as a serial port. `UART0` refers to this connection.
