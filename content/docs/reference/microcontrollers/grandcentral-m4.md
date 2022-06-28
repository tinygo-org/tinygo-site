---
title: "Adafruit Grand Central M4"
weight: 3
---

The [Adafruit Grand Central M4](https://www.adafruit.com/product/4064) is a tiny ARM development board based on the Atmel [ATSAMD51P20](https://www.microchip.com/wwwproducts/en/ATSAMD51P20A) family of SoC.

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
| `D0`              | `PB25`       | `UART1_RX_PIN`, `UART_RX_PIN` |
| `D1`              | `PB24`       | `UART1_TX_PIN`, `UART_TX_PIN` |
| `D2`              | `PC18`       |                   |
| `D3`              | `PC19`       |                   |
| `D4`              | `PC20`       |                   |
| `D5`              | `PC21`       |                   |
| `D6`              | `PD20`       |                   |
| `D7`              | `PD21`       |                   |
| `D8`              | `PB18`       |                   |
| `D9`              | `PB02`       |                   |
| `D10`             | `PB22`       |                   |
| `D11`             | `PB23`       |                   |
| `D12`             | `PB00`       |                   |
| `D13`             | `PB01`       | `D87`, `LED_PIN`, `LED` |
| `D14`             | `PB16`       | `UART4_TX_PIN`, `I2S0_SCK_PIN`, `I2S_SCK_PIN` |
| `D15`             | `PB17`       | `UART4_RX_PIN`, `I2S0_MCK_PIN` |
| `D16`             | `PC22`       | `UART3_TX_PIN`    |
| `D17`             | `PC23`       | `UART3_RX_PIN`    |
| `D18`             | `PB12`       | `UART2_TX_PIN`    |
| `D19`             | `PB13`       | `UART2_RX_PIN`    |
| `D20`             | `PB20`       | `D62`, `I2C0_SDA_PIN`, `I2C_SDA_PIN`, `SDA_PIN` |
| `D21`             | `PB21`       | `D63`, `I2C0_SCL_PIN`, `I2C_SCL_PIN`, `SCL_PIN` |
| `D22`             | `PD12`       |                   |
| `D23`             | `PA15`       |                   |
| `D24`             | `PC17`       | `I2C1_SCL_PIN`    |
| `D25`             | `PC16`       | `I2C1_SDA_PIN`    |
| `D26`             | `PA12`       |                   |
| `D27`             | `PA13`       |                   |
| `D28`             | `PA14`       |                   |
| `D29`             | `PB19`       |                   |
| `D30`             | `PA23`       |                   |
| `D31`             | `PA22`       | `I2S0_SDI_PIN`    |
| `D32`             | `PA21`       | `I2S0_SDO_PIN`, `I2S_SD_PIN` |
| `D33`             | `PA20`       | `I2S0_FS_PIN`, `I2S_WS_PIN` |
| `D34`             | `PA19`       |                   |
| `D35`             | `PA18`       |                   |
| `D36`             | `PA17`       |                   |
| `D37`             | `PA16`       |                   |
| `D38`             | `PB15`       |                   |
| `D39`             | `PB14`       |                   |
| `D40`             | `PC13`       |                   |
| `D41`             | `PC12`       |                   |
| `D42`             | `PC15`       |                   |
| `D43`             | `PC14`       |                   |
| `D44`             | `PC11`       |                   |
| `D45`             | `PC10`       |                   |
| `D46`             | `PC06`       |                   |
| `D47`             | `PC07`       |                   |
| `D48`             | `PC04`       |                   |
| `D49`             | `PC05`       |                   |
| `D50`             | `PD11`       | `D64`, `SPI0_SDI_PIN`, `SPI_SDI_PIN` |
| `D51`             | `PD08`       | `D65`, `SPI0_SDO_PIN`, `SPI_SDO_PIN` |
| `D52`             | `PD09`       | `D66`, `SPI0_SCK_PIN`, `SPI_SCK_PIN` |
| `D53`             | `PD10`       | `SPI0_CS_PIN`, `SPI_CS_PIN` |
| `D54`             | `PB05`       | `A8`              |
| `D55`             | `PB06`       | `A9`              |
| `D56`             | `PB07`       | `A10`             |
| `D57`             | `PB08`       | `A11`             |
| `D58`             | `PB09`       | `A12`             |
| `D59`             | `PA04`       | `A13`             |
| `D60`             | `PA06`       | `A14`             |
| `D61`             | `PA07`       | `A15`             |
| `D67`             | `PA02`       | `D85`, `A0`       |
| `D68`             | `PA05`       | `D86`, `A1`       |
| `D69`             | `PB03`       | `A2`              |
| `D70`             | `PC00`       | `A3`              |
| `D71`             | `PC01`       | `A4`              |
| `D72`             | `PC02`       | `A5`              |
| `D73`             | `PC03`       | `A6`              |
| `D74`             | `PB04`       | `A7`              |
| `D75`             | `PC31`       | `UART_RX_LED_PIN`, `LED_RX` |
| `D76`             | `PC30`       | `UART_TX_LED_PIN`, `LED_TX` |
| `D77`             | `PA27`       | `USBCDC_HOSTEN_PIN` |
| `D78`             | `PA24`       | `USBCDC_DM_PIN`   |
| `D79`             | `PA25`       | `USBCDC_DP_PIN`   |
| `D80`             | `PB29`       | `SPI1_SDI_PIN`, `SD0_SDI_PIN`, `SDCARD_SDI_PIN` |
| `D81`             | `PB27`       | `SPI1_SCK_PIN`, `SD0_SCK_PIN`, `SDCARD_SCK_PIN` |
| `D82`             | `PB26`       | `SPI1_SDO_PIN`, `SD0_SDO_PIN`, `SDCARD_SDO_PIN` |
| `D83`             | `PB28`       | `SD0_CS_PIN`, `SDCARD_CS_PIN` |
| `D84`             | `PA03`       | `AREF`            |
| `D88`             | `PC24`       | `NEOPIXEL_PIN`, `NEOPIXEL`, `WS2812` |
| `D89`             | `PB10`       | `QSPI_SCK`        |
| `D90`             | `PB11`       | `QSPI_CS`         |
| `D91`             | `PA08`       | `QSPI_DATA0`      |
| `D92`             | `PA09`       | `QSPI_DATA1`      |
| `D93`             | `PA10`       | `QSPI_DATA2`      |
| `D94`             | `PA11`       | `QSPI_DATA3`      |
| `D95`             | `PB31`       | `SD0_DET_PIN`, `SDCARD_DET_PIN` |
| `D96`             | `PB30`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Grand Central M4](../machine/grandcentral-m4)

## Flashing

### UF2

The Grand Central M4 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Grand Central M4 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=grandcentral-m4 [PATH TO YOUR PROGRAM]
    ```

- The Grand Central M4 board should restart and then begin running your program.


### Troubleshooting

If you have troubles getting your Grand Central M4 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Grand Central M4 board ready to receive code.
- The Grand Central M4 board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=grandcentral-m4 [PATH TO YOUR PROGRAM]
```

Once you have updated your Grand Central M4 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Grand Central M4 as a serial port. `UART0` refers to this connection.

The Neopixel LED on the Grand Central M4 can be accessed using the [WS2812](https://pkg.go.dev/tinygo.org/x/drivers/ws2812) driver via the `PC24` pin
