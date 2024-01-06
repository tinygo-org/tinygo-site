---
title: "Adafruit Feather M4 CAN"
weight: 3
---

The [Feather M4 CAN](https://www.adafruit.com/product/4759) is a tiny ARM development board based on the Atmel [ATSAME51J19](https://www.microchip.com/wwwproducts/en/ATSAME51J19A) family of SoC.

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
| `D0`              | `PB17`       | `UART_RX_PIN`     | `TCC3` (channel 1), `TCC0` (channel 5) |
| `D1`              | `PB16`       | `UART_TX_PIN`     | `TCC3` (channel 0), `TCC0` (channel 4) |
| `D4`              | `PA14`       |                   | `TCC2` (channel 0), `TCC1` (channel 2) |
| `D5`              | `PA16`       |                   | `TCC1` (channel 0), `TCC0` (channel 4) |
| `D6`              | `PA18`       |                   | `TCC1` (channel 2), `TCC0` (channel 6) |
| `D7`              | `PB03`       |                   | `TCC2` (channel 3)   |
| `D8`              | `PB02`       | `NEOPIXELS`, `WS2812` | `TCC2` (channel 2)   |
| `D9`              | `PA19`       |                   | `TCC1` (channel 3), `TCC0` (channel 7) |
| `D10`             | `PA20`       |                   | `TCC1` (channel 4), `TCC0` (channel 0) |
| `D11`             | `PA21`       |                   | `TCC1` (channel 5), `TCC0` (channel 1) |
| `D12`             | `PA22`       | `CAN0_TX`         | `TCC1` (channel 6), `TCC0` (channel 2) |
| `D13`             | `PA23`       | `LED`, `CAN0_RX`  | `TCC1` (channel 7), `TCC0` (channel 3) |
| `D21`             | `PA13`       | `SCL_PIN`         | `TCC0` (channel 7), `TCC1` (channel 3) |
| `D22`             | `PA12`       | `SDA_PIN`         | `TCC0` (channel 6), `TCC1` (channel 2) |
| `D23`             | `PB22`       | `SPI0_SDI_PIN`    |                      |
| `D24`             | `PB23`       | `SPI0_SDO_PIN`    |                      |
| `D25`             | `PA17`       | `SPI0_SCK_PIN`    | `TCC1` (channel 1), `TCC0` (channel 5) |
| `A0`              | `PA02`       |                   |                      |
| `A1`              | `PA05`       |                   |                      |
| `A2`              | `PB08`       |                   |                      |
| `A3`              | `PB09`       |                   |                      |
| `A4`              | `PA04`       | `UART2_TX_PIN`    |                      |
| `A5`              | `PA06`       | `UART2_RX_PIN`    |                      |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC2` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC2` (channel 3)   |
| `CAN1_STANDBY`    | `PB12`       | `CAN_STANDBY`, `CAN_S` | `TCC3` (channel 0), `TCC0` (channel 0) |
| `CAN1_TX`         | `PB14`       | `CAN_TX`          | `TCC4` (channel 0), `TCC0` (channel 2) |
| `CAN1_RX`         | `PB15`       | `CAN_RX`          | `TCC4` (channel 1), `TCC0` (channel 3) |
| `BOOST_EN`        | `PB13`       |                   | `TCC3` (channel 1), `TCC0` (channel 1) |
| `QSPI_SCK`        | `PB10`       |                   | `TCC0` (channel 4), `TCC1` (channel 0) |
| `QSPI_CS`         | `PB11`       |                   | `TCC0` (channel 5), `TCC1` (channel 1) |
| `QSPI_DATA0`      | `PA08`       |                   | `TCC0` (channel 0), `TCC1` (channel 4) |
| `QSPI_DATA1`      | `PA09`       |                   | `TCC0` (channel 1), `TCC1` (channel 5) |
| `QSPI_DATA2`      | `PA10`       |                   | `TCC0` (channel 2), `TCC1` (channel 6) |
| `QSPI_DATA3`      | `PA11`       |                   | `TCC0` (channel 3), `TCC1` (channel 7) |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Feather M4 CAN](../machine/feather-m4-can)

## Flashing

### UF2

The Feather M4 CAN comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Feather M4 CAN into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-m4-can [PATH TO YOUR PROGRAM]
    ```

- The Feather M4 CAN board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Feather M4 CAN board to receive code, try this:

- Press the "RESET" button on the board two times to get the Feather M4 CAN board ready to receive code.
- The Feather M4 CAN board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=feather-m4-can [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Feather M4 CAN board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Feather M4 CAN as a serial port. `UART0` refers to this connection.
