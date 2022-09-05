---
title: "Adafruit ItsyBitsy M0"
weight: 3
---

The [Adafruit ItsyBitsy M0](https://www.adafruit.com/product/3727) is very compact ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of SoC.

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
| `D0`              | `PA11`       |                   | `TCC1` (channel 1), `TCC0` (channel 3) |
| `D1`              | `PA10`       | `I2S_SCK_PIN`     | `TCC1` (channel 0), `TCC0` (channel 2) |
| `D2`              | `PA14`       |                   | `TCC0` (channel 0)   |
| `D3`              | `PA09`       |                   | `TCC0` (channel 1), `TCC1` (channel 3) |
| `D4`              | `PA08`       | `I2S_SD_PIN`      | `TCC0` (channel 0), `TCC1` (channel 2) |
| `D5`              | `PA15`       |                   | `TCC0` (channel 1)   |
| `D6`              | `PA20`       |                   | `TCC0` (channel 2)   |
| `D7`              | `PA21`       |                   | `TCC0` (channel 3)   |
| `D8`              | `PA06`       |                   | `TCC1` (channel 0)   |
| `D9`              | `PA07`       |                   | `TCC1` (channel 1)   |
| `D10`             | `PA18`       | `UART_TX_PIN`     | `TCC0` (channel 2)   |
| `D11`             | `PA16`       | `UART_RX_PIN`     | `TCC2` (channel 0), `TCC0` (channel 2) |
| `D12`             | `PA19`       |                   | `TCC0` (channel 3)   |
| `D13`             | `PA17`       | `LED`             | `TCC2` (channel 1), `TCC0` (channel 3) |
| `A0`              | `PA02`       |                   |                      |
| `A1`              | `PB08`       |                   |                      |
| `A2`              | `PB09`       |                   |                      |
| `A3`              | `PA04`       |                   | `TCC0` (channel 0)   |
| `A4`              | `PA05`       |                   | `TCC0` (channel 1)   |
| `A5`              | `PB02`       |                   |                      |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC1` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC1` (channel 3)   |
| `SDA_PIN`         | `PA22`       |                   | `TCC0` (channel 0)   |
| `SCL_PIN`         | `PA23`       |                   | `TCC0` (channel 1)   |
| `SPI0_SCK_PIN`    | `PB11`       |                   | `TCC0` (channel 1)   |
| `SPI0_SDO_PIN`    | `PB10`       |                   | `TCC0` (channel 0)   |
| `SPI0_SDI_PIN`    | `PA12`       |                   | `TCC2` (channel 0), `TCC0` (channel 2) |
| `SPI1_CS_PIN`     | `PA27`       |                   |                      |
| `SPI1_SCK_PIN`    | `PB23`       |                   |                      |
| `SPI1_SDO_PIN`    | `PB22`       |                   |                      |
| `SPI1_SDI_PIN`    | `PB03`       |                   |                      |

## Machine Package Docs

[Documentation for the machine package for the Adafruit ItsyBitsy M0](../machine/itsybitsy-m0)

## Flashing

### UF2

The ItsyBitsy M0 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your ItsyBitsy M0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=itsybitsy-m0 [PATH TO YOUR PROGRAM]
    ```

- The ItsyBitsy M0 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your ItsyBitsy M0 board to receive code, try this:

- Press the "RESET" button on the board two times to get the ItsyBitsy M0 board ready to receive code.
- The ItsyBitsy M0 board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=itsybitsy-m0 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your ItsyBitsy M0 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the ItsyBitsy M0 as a serial port. `UART0` refers to this connection.

The DotStar LED on the ItsyBitsy M0 can be accessed using the [APA102](https://pkg.go.dev/tinygo.org/x/drivers/apa102) driver via a software SPI on the `PA00` and `PA01` pins
