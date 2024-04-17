---
title: "Adafruit Metro M4 Express AirLift"
weight: 3
---

The [Adafruit Metro M4 Express AirLift](https://www.adafruit.com/product/4000) is an ARM development board based on the Atmel [ATSAMD51J19](https://www.microchip.com/wwwproducts/en/ATSAMD51J19) family of SoC that has Arduino shield compatible form factor, and a built-in EspressIf ESP32 Wi-Fi Co processor.

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
| `D0`              | `PA23`       | `UART_RX_PIN`     | `TCC1` (channel 7), `TCC0` (channel 3) |
| `D1`              | `PA22`       | `UART_TX_PIN`     | `TCC1` (channel 6), `TCC0` (channel 2) |
| `D2`              | `PB17`       |                   | `TCC3` (channel 1), `TCC0` (channel 5) |
| `D3`              | `PB16`       |                   | `TCC3` (channel 0), `TCC0` (channel 4) |
| `D4`              | `PB13`       |                   | `TCC3` (channel 1), `TCC0` (channel 1) |
| `D5`              | `PB14`       |                   | `TCC4` (channel 0), `TCC0` (channel 2) |
| `D6`              | `PB15`       |                   | `TCC4` (channel 1), `TCC0` (channel 3) |
| `D7`              | `PB12`       |                   | `TCC3` (channel 0), `TCC0` (channel 0) |
| `D8`              | `PA21`       |                   | `TCC1` (channel 5), `TCC0` (channel 1) |
| `D9`              | `PA20`       |                   | `TCC1` (channel 4), `TCC0` (channel 0) |
| `D10`             | `PA18`       |                   | `TCC1` (channel 2), `TCC0` (channel 6) |
| `D11`             | `PA19`       | `SPI1_SDO_PIN`    | `TCC1` (channel 3), `TCC0` (channel 7) |
| `D12`             | `PA17`       | `SPI1_SCK_PIN`    | `TCC1` (channel 1), `TCC0` (channel 5) |
| `D13`             | `PA16`       | `LED`, `SPI1_SDI_PIN` | `TCC1` (channel 0), `TCC0` (channel 4) |
| `D40`             | `PB22`       | `WS2812`          |                      |
| `A0`              | `PA02`       |                   |                      |
| `A1`              | `PA05`       |                   |                      |
| `A2`              | `PB06`       |                   |                      |
| `A3`              | `PB00`       |                   |                      |
| `A4`              | `PB08`       |                   |                      |
| `A5`              | `PB09`       |                   |                      |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC2` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC2` (channel 3)   |
| `UART2_TX_PIN`    | `PA04`       | `NINA_TX`         |                      |
| `UART2_RX_PIN`    | `PA07`       | `NINA_RX`         |                      |
| `NINA_CS`         | `PA15`       |                   | `TCC2` (channel 1), `TCC1` (channel 3) |
| `NINA_ACK`        | `PB04`       | `NINA_CTS`        |                      |
| `NINA_GPIO0`      | `PB01`       | `NINA_RTS`        |                      |
| `NINA_RESETN`     | `PB05`       |                   |                      |
| `SDA_PIN`         | `PB02`       |                   | `TCC2` (channel 2)   |
| `SCL_PIN`         | `PB03`       |                   | `TCC2` (channel 3)   |
| `SPI0_SCK_PIN`    | `PA13`       | `NINA_SCK`        | `TCC0` (channel 7), `TCC1` (channel 3) |
| `SPI0_SDO_PIN`    | `PA12`       | `NINA_SDO`        | `TCC0` (channel 6), `TCC1` (channel 2) |
| `SPI0_SDI_PIN`    | `PA14`       | `NINA_SDI`        | `TCC2` (channel 0), `TCC1` (channel 2) |
| `QSPI_SCK`        | `PB10`       |                   | `TCC0` (channel 4), `TCC1` (channel 0) |
| `QSPI_CS`         | `PB11`       |                   | `TCC0` (channel 5), `TCC1` (channel 1) |
| `QSPI_DATA0`      | `PA08`       |                   | `TCC0` (channel 0), `TCC1` (channel 4) |
| `QSPI_DATA1`      | `PA09`       |                   | `TCC0` (channel 1), `TCC1` (channel 5) |
| `QSPI_DATA2`      | `PA10`       |                   | `TCC0` (channel 2), `TCC1` (channel 6) |
| `QSPI_DATA3`      | `PA11`       |                   | `TCC0` (channel 3), `TCC1` (channel 7) |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Metro M4 Airlift](../machine/metro-m4-airlift)

## Flashing

### UF2

The Metro M4 Express comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Metro M4 Express into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=metro-m4-airlift [PATH TO YOUR PROGRAM]
    ```

- The Metro M4 Express board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Metro M4 Express board to receive code, try this:

- Press the "RESET" button on the board two times to get the Metro M4 Express board ready to receive code.
- The Metro M4 Express board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=metro-m4-airlift [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Metro M4 Express board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Metro M4 Express as a serial port. `UART0` refers to this connection.

The Neopixel LED on the Metro M4 Express can be accessed using the [WS2812](https://pkg.go.dev/tinygo.org/x/drivers/ws2812) driver via the `PB22` pin
