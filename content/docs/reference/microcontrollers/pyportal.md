---
title: "Adafruit PyPortal"
weight: 3
---

The [Adafruit PyPortal](https://www.adafruit.com/product/4116) is a ARM board based on the Microchip [ATSAMD51J20A](https://www.microchip.com/wwwproducts/en/ATSAMD51J20A) family of SoC.

The PyPortal also has an Espressif ESP32 Wi-Fi coprocessor with TLS/SSL support built-in. PyPortal has a 3.2″ 320 x 240 color TFT with resistive touch screen. PyPortal includes: speaker, light sensor, temperature sensor, NeoPixel, microSD card slot, 8MB flash, plug-in ports for I2C and 2 analog/digital pins,

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
| `D0`              | `PB13`       | `NINA_RX`, `UART_RX_PIN` | `TCC3` (channel 1), `TCC0` (channel 1) |
| `D1`              | `PB12`       | `NINA_TX`, `UART_TX_PIN` | `TCC3` (channel 0), `TCC0` (channel 0) |
| `D2`              | `PB22`       | `NEOPIXEL`, `WS2812` |                      |
| `D3`              | `PA04`       | `A1`              |                      |
| `D4`              | `PA05`       | `A3`              |                      |
| `D5`              | `PB16`       | `NINA_ACK`, `NINA_CTS` | `TCC3` (channel 0), `TCC0` (channel 4) |
| `D6`              | `PB15`       | `NINA_GPIO0`, `NINA_RTS` | `TCC4` (channel 1), `TCC0` (channel 3) |
| `D7`              | `PB17`       | `NINA_RESETN`     | `TCC3` (channel 1), `TCC0` (channel 5) |
| `D8`              | `PB14`       | `NINA_CS`         | `TCC4` (channel 0), `TCC0` (channel 2) |
| `D9`              | `PB04`       | `TFT_RD`          |                      |
| `D10`             | `PB05`       | `TFT_DC`          |                      |
| `D11`             | `PB06`       | `TFT_CS`          |                      |
| `D12`             | `PB07`       | `TFT_TE`          |                      |
| `D13`             | `PB23`       | `LED`             |                      |
| `D24`             | `PA00`       | `TFT_RESET`       |                      |
| `D25`             | `PB31`       | `TFT_BACKLIGHT`   | `TCC4` (channel 1), `TCC0` (channel 7) |
| `D26`             | `PB09`       | `TFT_WR`          |                      |
| `D27`             | `PB02`       | `SDA_PIN`         | `TCC2` (channel 2)   |
| `D28`             | `PB03`       | `SCL_PIN`         | `TCC2` (channel 3)   |
| `D29`             | `PA12`       | `SPI0_SDO_PIN`, `NINA_SDO` | `TCC0` (channel 6), `TCC1` (channel 2) |
| `D30`             | `PA13`       | `SPI0_SCK_PIN`, `NINA_SCK` | `TCC0` (channel 7), `TCC1` (channel 3) |
| `D31`             | `PA14`       | `SPI0_SDI_PIN`, `NINA_SDI` | `TCC2` (channel 0), `TCC1` (channel 2) |
| `D32`             | `PB30`       |                   | `TCC4` (channel 0), `TCC0` (channel 6) |
| `D33`             | `PA01`       |                   |                      |
| `D34`             | `PA16`       | `LCD_DATA0`       | `TCC1` (channel 0), `TCC0` (channel 4) |
| `D35`             | `PA17`       |                   | `TCC1` (channel 1), `TCC0` (channel 5) |
| `D36`             | `PA18`       |                   | `TCC1` (channel 2), `TCC0` (channel 6) |
| `D37`             | `PA19`       |                   | `TCC1` (channel 3), `TCC0` (channel 7) |
| `D38`             | `PA20`       |                   | `TCC1` (channel 4), `TCC0` (channel 0) |
| `D39`             | `PA21`       |                   | `TCC1` (channel 5), `TCC0` (channel 1) |
| `D40`             | `PA22`       |                   | `TCC1` (channel 6), `TCC0` (channel 2) |
| `D41`             | `PA23`       |                   | `TCC1` (channel 7), `TCC0` (channel 3) |
| `D42`             | `PB10`       | `QSPI_SCK`        | `TCC0` (channel 4), `TCC1` (channel 0) |
| `D43`             | `PB11`       | `QSPI_CS`         | `TCC0` (channel 5), `TCC1` (channel 1) |
| `D44`             | `PA08`       | `QSPI_DATA0`      | `TCC0` (channel 0), `TCC1` (channel 4) |
| `D45`             | `PA09`       | `QSPI_DATA1`      | `TCC0` (channel 1), `TCC1` (channel 5) |
| `D46`             | `PA10`       | `QSPI_DATA2`      | `TCC0` (channel 2), `TCC1` (channel 6) |
| `D47`             | `PA11`       | `QSPI_DATA3`      | `TCC0` (channel 3), `TCC1` (channel 7) |
| `D50`             | `PA02`       | `SPK_SD`, `A0`, `AUDIO_OUT` |                      |
| `D51`             | `PA15`       |                   | `TCC2` (channel 1), `TCC1` (channel 3) |
| `A2`              | `PA07`       | `LIGHT`           |                      |
| `A4`              | `PB00`       | `TOUCH_YD`        |                      |
| `A5`              | `PB01`       | `TOUCH_XL`        |                      |
| `A6`              | `PA06`       | `TOUCH_YU`        |                      |
| `A7`              | `PB08`       | `TOUCH_XR`        |                      |
| `USBCDC_DM_PIN`   | `PA24`       |                   | `TCC2` (channel 2)   |
| `USBCDC_DP_PIN`   | `PA25`       |                   | `TCC2` (channel 3)   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit PyPortal](../machine/pyportal)

## Flashing

### UF2

The PyPortal comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your PyPortal into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=pyportal [PATH TO YOUR PROGRAM]
    ```

- The PyPortal board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your PyPortal board to receive code, try this:

- Press the "RESET" button on the board two times to get the PyPortal board ready to receive code.
- The PyPortal board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=pyportal [PATH TO YOUR PROGRAM]
```

Once you have updated your PyPortal board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the PyPortal as a serial port. `UART0` refers to this connection.
