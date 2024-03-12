---
title: "Adafruit QT Py"
weight: 3
---

The [QT Py](https://www.adafruit.com/product/4600) is a tiny ARM development board based on the Atmel [ATSAMD21E18](https://www.microchip.com/wwwproducts/en/ATSAMD21E18) family of SoC.

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
| `D1`              | `PA03`       | `A1`              |                      |
| `D2`              | `PA04`       | `A2`              | `TCC0` (channel 0)   |
| `D3`              | `PA05`       | `A3`              | `TCC0` (channel 1)   |
| `D4`              | `PA16`       | `A4`, `SDA_PIN`   | `TCC2` (channel 0), `TCC0` (channel 2) |
| `D5`              | `PA17`       | `SCL_PIN`         | `TCC2` (channel 1), `TCC0` (channel 3) |
| `D6`              | `PA06`       | `UART_TX_PIN`     | `TCC1` (channel 0)   |
| `D7`              | `PA07`       | `UART_RX_PIN`     | `TCC1` (channel 1)   |
| `D8`              | `PA11`       | `SPI0_SCK_PIN`    | `TCC1` (channel 1), `TCC0` (channel 3) |
| `D9`              | `PA09`       | `SPI0_SDI_PIN`    | `TCC0` (channel 1), `TCC1` (channel 3) |
| `D10`             | `PA10`       | `SPI0_SDO_PIN`, `I2S_SCK_PIN` | `TCC1` (channel 0), `TCC0` (channel 2) |
| `D11`             | `PA18`       | `NEOPIXELS`, `WS2812` | `TCC0` (channel 2)   |
| `D12`             | `PA15`       | `NEOPIXELS_POWER` | `TCC0` (channel 1)   |
| `D13`             | `PA27`       |                   |                      |
| `D14`             | `PA23`       |                   | `TCC0` (channel 1)   |
| `D15`             | `PA19`       |                   | `TCC0` (channel 3)   |
| `D16`             | `PA22`       |                   | `TCC0` (channel 0)   |
| `D17`             | `PA08`       | `I2S_SD_PIN`      | `TCC0` (channel 0), `TCC1` (channel 2) |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC1` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC1` (channel 3)   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit QtPy](../machine/qtpy)

## Flashing

### UF2

The QtPy comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your QtPy into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=qtpy [PATH TO YOUR PROGRAM]
    ```

- The QtPy board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your QtPy board to receive code, try this:

- Press the "RESET" button on the board two times to get the QtPy board ready to receive code.
- The QtPy board will appear to your computer like a USB drive.
- Now try running the `tinygo flash` command as above:

    ```shell
    tinygo flash -target=qtpy [PATH TO YOUR PROGRAM]
    ```

Once you have updated your QtPy board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the QtPy as a serial port. `UART0` refers to this connection.
