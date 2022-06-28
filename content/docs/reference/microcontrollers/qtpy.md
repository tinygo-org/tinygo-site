---
title: "Adafruit Qt Py"
weight: 3
---

The [Adafruit QtPy](https://www.adafruit.com/product/4600) is a tiny ARM development board based on the Atmel [ATSAMD21E18](https://www.microchip.com/wwwproducts/en/ATSAMD21E18) family of SoC.

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

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `PA02`       |                   |
| `D1`              | `PA03`       | `A0`, `A1`        |
| `D2`              | `PA04`       | `A2`              |
| `D3`              | `PA05`       | `A3`              |
| `D4`              | `PA16`       | `A4`, `SDA_PIN`   |
| `D5`              | `PA17`       | `SCL_PIN`         |
| `D6`              | `PA06`       | `UART_TX_PIN`     |
| `D7`              | `PA07`       | `UART_RX_PIN`     |
| `D8`              | `PA11`       | `SPI0_SCK_PIN`    |
| `D9`              | `PA09`       | `SPI0_SDI_PIN`    |
| `D10`             | `PA10`       | `SPI0_SDO_PIN`, `I2S_SCK_PIN` |
| `D11`             | `PA18`       | `NEOPIXELS`, `WS2812` |
| `D12`             | `PA15`       | `NEOPIXELS_POWER` |
| `D13`             | `PA27`       |                   |
| `D14`             | `PA23`       |                   |
| `D15`             | `PA19`       |                   |
| `D16`             | `PA22`       |                   |
| `D17`             | `PA08`       | `I2S_SD_PIN`      |
| `USBCDC_DM_PIN`   | `PA24`       |                   |
| `USBCDC_DP_PIN`   | `PA25`       |                   |

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
