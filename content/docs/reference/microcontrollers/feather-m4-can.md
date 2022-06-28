---
title: "Adafruit Feather M4 CAN"
weight: 3
---

The [Adafruit Feather M4 CAN](https://www.adafruit.com/product/4759) is a tiny ARM development board based on the Atmel [ATSAME51J19](https://www.microchip.com/wwwproducts/en/ATSAME51J19A) family of SoC.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `PB17`       | `UART_RX_PIN`     |
| `D1`              | `PB16`       | `UART_TX_PIN`     |
| `D4`              | `PA14`       |                   |
| `D5`              | `PA16`       |                   |
| `D6`              | `PA18`       |                   |
| `D7`              | `PB03`       |                   |
| `D8`              | `PB02`       | `NEOPIXELS`, `WS2812` |
| `D9`              | `PA19`       |                   |
| `D10`             | `PA20`       |                   |
| `D11`             | `PA21`       |                   |
| `D12`             | `PA22`       | `CAN0_TX`         |
| `D13`             | `PA23`       | `LED`, `CAN0_RX`  |
| `D21`             | `PA13`       | `SCL_PIN`         |
| `D22`             | `PA12`       | `SDA_PIN`         |
| `D23`             | `PB22`       | `SPI0_SDI_PIN`    |
| `D24`             | `PB23`       | `SPI0_SDO_PIN`    |
| `D25`             | `PA17`       | `SPI0_SCK_PIN`    |
| `A0`              | `PA02`       |                   |
| `A1`              | `PA05`       |                   |
| `A2`              | `PB08`       |                   |
| `A3`              | `PB09`       |                   |
| `A4`              | `PA04`       | `UART2_TX_PIN`    |
| `A5`              | `PA06`       | `UART2_RX_PIN`    |
| `USBCDC_DM_PIN`   | `PA24`       |                   |
| `USBCDC_DP_PIN`   | `PA25`       |                   |
| `CAN1_STANDBY`    | `PB12`       | `CAN_STANDBY`, `CAN_S` |
| `CAN1_TX`         | `PB14`       | `CAN_TX`          |
| `CAN1_RX`         | `PB15`       | `CAN_RX`          |
| `BOOST_EN`        | `PB13`       |                   |
| `QSPI_SCK`        | `PB10`       |                   |
| `QSPI_CS`         | `PB11`       |                   |
| `QSPI_DATA0`      | `PA08`       |                   |
| `QSPI_DATA1`      | `PA09`       |                   |
| `QSPI_DATA2`      | `PA10`       |                   |
| `QSPI_DATA3`      | `PA11`       |                   |

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
