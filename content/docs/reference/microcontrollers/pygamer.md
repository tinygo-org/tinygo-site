---
title: "Adafruit PyGamer"
weight: 3
---

The [Adafruit PyGamer](https://www.adafruit.com/product/4242) is a ARM development board based on the Atmel [ATSAMD51J19A](https://www.microchip.com/wwwproducts/en/ATSAMD51J19A) family of SoC.

It has many built-in devices, such as a 1.8" 160x128 Color TFT Display, a dual-potentiometer analog stick, 4 x square-top buttons, 5 x NeoPixel LED, a triple-axis accelerometer, a light sensor, and a speaker.  The PyGamer uses the ST7735 display, so you may use the tinygo-org st7735 driver. The accelerometer is the LIS3DH so you may use the tinygo lis3dh driver.

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
| `D2`              | `PB03`       | `A8`              | `TCC2` (channel 3)   |
| `D3`              | `PB02`       | `A9`              | `TCC2` (channel 2)   |
| `D4`              | `PA14`       |                   | `TCC2` (channel 0), `TCC1` (channel 2) |
| `D5`              | `PA16`       |                   | `TCC1` (channel 0), `TCC0` (channel 4) |
| `D6`              | `PA18`       |                   | `TCC1` (channel 2), `TCC0` (channel 6) |
| `D7`              | `PB14`       | `SD_CS`           | `TCC4` (channel 0), `TCC0` (channel 2) |
| `D8`              | `PA15`       | `NEOPIXELS`, `WS2812` | `TCC2` (channel 1), `TCC1` (channel 3) |
| `D9`              | `PA19`       |                   | `TCC1` (channel 3), `TCC0` (channel 7) |
| `D10`             | `PA20`       |                   | `TCC1` (channel 4), `TCC0` (channel 0) |
| `D11`             | `PA21`       |                   | `TCC1` (channel 5), `TCC0` (channel 1) |
| `D12`             | `PA22`       |                   | `TCC1` (channel 6), `TCC0` (channel 2) |
| `D13`             | `PA23`       | `LED`             | `TCC1` (channel 7), `TCC0` (channel 3) |
| `A0`              | `PA02`       |                   |                      |
| `A1`              | `PA05`       |                   |                      |
| `A2`              | `PB08`       |                   |                      |
| `A3`              | `PB09`       |                   |                      |
| `A4`              | `PA04`       | `UART2_TX_PIN`    |                      |
| `A5`              | `PA06`       | `UART2_RX_PIN`    |                      |
| `A6`              | `PB01`       |                   |                      |
| `A7`              | `PB04`       | `LIGHTSENSOR`     |                      |
| `BUTTON_LATCH`    | `PB00`       |                   |                      |
| `BUTTON_OUT`      | `PB30`       |                   | `TCC4` (channel 0), `TCC0` (channel 6) |
| `BUTTON_CLK`      | `PB31`       |                   | `TCC4` (channel 1), `TCC0` (channel 7) |
| `JOYY`            | `PB06`       |                   |                      |
| `JOYX`            | `PB07`       |                   |                      |
| `TFT_DC`          | `PB05`       |                   |                      |
| `TFT_CS`          | `PB12`       |                   | `TCC3` (channel 0), `TCC0` (channel 0) |
| `TFT_RST`         | `PA00`       |                   |                      |
| `TFT_LITE`        | `PA01`       |                   |                      |
| `SPEAKER_ENABLE`  | `PA27`       |                   |                      |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC2` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC2` (channel 3)   |
| `SDA_PIN`         | `PA12`       |                   | `TCC0` (channel 6), `TCC1` (channel 2) |
| `SCL_PIN`         | `PA13`       |                   | `TCC0` (channel 7), `TCC1` (channel 3) |
| `SPI0_SCK_PIN`    | `PA17`       |                   | `TCC1` (channel 1), `TCC0` (channel 5) |
| `SPI0_SDO_PIN`    | `PB23`       |                   |                      |
| `SPI0_SDI_PIN`    | `PB22`       |                   |                      |
| `SPI1_SCK_PIN`    | `PB13`       |                   | `TCC3` (channel 1), `TCC0` (channel 1) |
| `SPI1_SDO_PIN`    | `PB15`       |                   | `TCC4` (channel 1), `TCC0` (channel 3) |
| `QSPI_SCK`        | `PB10`       |                   | `TCC0` (channel 4), `TCC1` (channel 0) |
| `QSPI_CS`         | `PB11`       |                   | `TCC0` (channel 5), `TCC1` (channel 1) |
| `QSPI_DATA0`      | `PA08`       |                   | `TCC0` (channel 0), `TCC1` (channel 4) |
| `QSPI_DATA1`      | `PA09`       |                   | `TCC0` (channel 1), `TCC1` (channel 5) |
| `QSPI_DATA2`      | `PA10`       |                   | `TCC0` (channel 2), `TCC1` (channel 6) |
| `QSPI_DATA3`      | `PA11`       |                   | `TCC0` (channel 3), `TCC1` (channel 7) |

## Machine Package Docs

[Documentation for the machine package for the Adafruit PyGamer](../machine/pygamer)

## Flashing

### UF2

The PyGamer comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your PyGamer into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pygamer [PATH TO YOUR PROGRAM]
    ```

- The PyGamer board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your PyGamer board to receive code, try this:

- Press the "RESET" button on the board two times to get the PyGamer board ready to receive code.
- The PyGamer board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=pygamer [PATH TO YOUR PROGRAM]
```

Once you have updated your PyGamer board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the PyGamer as a serial port. `UART0` refers to this connection.
