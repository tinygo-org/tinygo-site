---
title: "Microchip SAM E54 Xplained Pro"
weight: 3
---

The [Microchip SAM E54 Xplained Pro](https://www.microchip.com/developmenttools/productdetails/atsame54-xpro) is a tiny ARM development board based on the Atmel [ATSAME54P20](https://www.microchip.com/wwwproducts/en/ATSAME54P20A) family of SoC.

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
| `LED`             | `PC18`       | `PDEC_INDEX`, `PIN_LED0` |
| `BUTTON`          | `PB31`       | `PIN_BTN0`        |
| `EXT1_PIN3_ADC_P` | `PB04`       |                   |
| `EXT1_PIN4_ADC_N` | `PB05`       |                   |
| `EXT1_PIN5_GPIO1` | `PA06`       |                   |
| `EXT1_PIN6_GPIO2` | `PA07`       |                   |
| `EXT1_PIN7_PWM_P` | `PB08`       |                   |
| `EXT1_PIN8_PWM_N` | `PB09`       |                   |
| `EXT1_PIN9_IRQ`   | `PB07`       | `EXT1_PIN9_GPIO`  |
| `EXT1_PIN10_SPI_SS_B` | `PA27`       | `EXT1_PIN10_GPIO` |
| `EXT1_PIN11_TWI_SDA` | `PA22`       | `CAN0_TX`, `PCC_DATA06`, `SDA0_PIN`, `SDA_PIN` |
| `EXT1_PIN12_TWI_SCL` | `PA23`       | `CAN0_RX`, `PCC_DATA07`, `SCL0_PIN`, `SCL_PIN` |
| `EXT1_PIN13_UART_RX` | `PA05`       | `UART_RX_PIN`     |
| `EXT1_PIN14_UART_TX` | `PA04`       | `UART_TX_PIN`     |
| `EXT1_PIN15_SPI_SS_A` | `PB28`       | `SPI0_SS_PIN`     |
| `EXT1_PIN16_SPI_SDO` | `PB27`       | `SPI0_SDO_PIN`    |
| `EXT1_PIN17_SPI_SDI` | `PB29`       | `SPI0_SDI_PIN`    |
| `EXT1_PIN18_SPI_SCK` | `PB26`       | `SPI0_SCK_PIN`    |
| `EXT2_PIN3_ADC_P` | `PB00`       |                   |
| `EXT2_PIN4_ADC_N` | `PA03`       |                   |
| `EXT2_PIN5_GPIO1` | `PB01`       |                   |
| `EXT2_PIN6_GPIO2` | `PB06`       |                   |
| `EXT2_PIN7_PWM_P` | `PB14`       | `PCC_DATA08`      |
| `EXT2_PIN8_PWM_N` | `PB15`       | `PCC_DATA09`      |
| `EXT2_PIN9_IRQ`   | `PD00`       | `EXT2_PIN9_GPIO`  |
| `EXT2_PIN10_SPI_SS_B` | `PB02`       | `EXT2_PIN10_GPIO` |
| `EXT2_PIN11_TWI_SDA` | `PD08`       | `EXT3_PIN11_TWI_SDA`, `I2C_SDA`, `PCC_I2C_SDA`, `SDA1_PIN`, `SDA2_PIN`, `SDA_DGI_PIN` |
| `EXT2_PIN12_TWI_SCL` | `PD09`       | `EXT3_PIN12_TWI_SCL`, `I2C_SCL`, `PCC_I2C_SCL`, `SCL1_PIN`, `SCL2_PIN`, `SCL_DGI_PIN` |
| `EXT2_PIN13_UART_RX` | `PB17`       | `UART2_RX_PIN`    |
| `EXT2_PIN14_UART_TX` | `PB16`       | `UART2_TX_PIN`    |
| `EXT2_PIN15_SPI_SS_A` | `PC06`       | `SPI1_SS_PIN`     |
| `EXT2_PIN16_SPI_SDO` | `PC04`       | `EXT3_PIN16_SPI_SDO`, `SPI1_SDO_PIN`, `SPI2_SDO_PIN`, `SPI_DGI_SDO_PIN` |
| `EXT2_PIN17_SPI_SDI` | `PC07`       | `EXT3_PIN17_SPI_SDI`, `SPI1_SDI_PIN`, `SPI2_SDI_PIN`, `SPI_DGI_SDI_PIN` |
| `EXT2_PIN18_SPI_SCK` | `PC05`       | `EXT3_PIN18_SPI_SCK`, `SPI1_SCK_PIN`, `SPI2_SCK_PIN`, `SPI_DGI_SCK_PIN` |
| `EXT3_PIN3_ADC_P` | `PC02`       |                   |
| `EXT3_PIN4_ADC_N` | `PC03`       |                   |
| `EXT3_PIN5_GPIO1` | `PC01`       |                   |
| `EXT3_PIN6_GPIO2` | `PC10`       |                   |
| `EXT3_PIN7_PWM_P` | `PD10`       |                   |
| `EXT3_PIN8_PWM_N` | `PD11`       |                   |
| `EXT3_PIN9_IRQ`   | `PC30`       | `EXT3_PIN9_GPIO`  |
| `EXT3_PIN10_SPI_SS_B` | `PC31`       | `EXT3_PIN10_GPIO` |
| `EXT3_PIN13_UART_RX` | `PC23`       | `UART3_RX_PIN`    |
| `EXT3_PIN14_UART_TX` | `PC22`       | `UART3_TX_PIN`    |
| `EXT3_PIN15_SPI_SS_A` | `PC14`       | `SPI2_SS_PIN`     |
| `SD_CARD_MCDA0`   | `PB18`       |                   |
| `SD_CARD_MCDA1`   | `PB19`       |                   |
| `SD_CARD_MCDA2`   | `PB20`       |                   |
| `SD_CARD_MCDA3`   | `PB21`       |                   |
| `SD_CARD_MCCK`    | `PA21`       | `PCC_DATA05`      |
| `SD_CARD_MCCDA`   | `PA20`       | `PCC_DATA04`      |
| `SD_CARD_DETECT`  | `PD20`       |                   |
| `SD_CARD_PROTECT` | `PD21`       |                   |
| `CAN1_STANDBY`    | `PC13`       | `CAN_STANDBY`     |
| `CAN1_TX`         | `PB12`       | `CAN_TX`          |
| `CAN1_RX`         | `PB13`       | `CAN_RX`          |
| `PDEC_PHASE_A`    | `PC16`       |                   |
| `PDEC_PHASE_B`    | `PC17`       |                   |
| `PCC_VSYNC_DEN1`  | `PA12`       | `ETHERNET_RX1`    |
| `PCC_HSYNC_DEN2`  | `PA13`       | `ETHERNET_RX0`    |
| `PCC_CLK`         | `PA14`       | `ETHERNET_TXCK`   |
| `PCC_XCLK`        | `PA15`       | `ETHERNET_RXER`   |
| `PCC_DATA00`      | `PA16`       | `PIN_QT_BUTTON`   |
| `PCC_DATA01`      | `PA17`       | `ETHERNET_TXEN`   |
| `PCC_DATA02`      | `PA18`       | `ETHERNET_TX0`    |
| `PCC_DATA03`      | `PA19`       | `ETHERNET_TX1`    |
| `PCC_RESET`       | `PC12`       | `ETHERNET_MDIO`   |
| `PCC_PWDN`        | `PC11`       | `ETHERNET_MDC`    |
| `ETHERNET_RXDV`   | `PC20`       |                   |
| `ETHERNET_INT`    | `PD12`       |                   |
| `ETHERNET_RESET`  | `PC21`       |                   |
| `PIN_ETH_LED`     | `PC15`       |                   |
| `PIN_ADC_DAC`     | `PA02`       |                   |
| `PIN_VBUS_DETECT` | `PC00`       |                   |
| `PIN_USB_ID`      | `PC19`       |                   |
| `USBCDC_DM_PIN`   | `PA24`       |                   |
| `USBCDC_DP_PIN`   | `PA25`       |                   |
| `UART4_TX_PIN`    | `PB25`       |                   |
| `UART4_RX_PIN`    | `PB24`       |                   |
| `SPI_DGI_SS_PIN`  | `PD01`       |                   |
| `QSPI_SCK`        | `PB10`       |                   |
| `QSPI_CS`         | `PB11`       |                   |
| `QSPI_DATA0`      | `PA08`       |                   |
| `QSPI_DATA1`      | `PA09`       |                   |
| `QSPI_DATA2`      | `PA10`       |                   |
| `QSPI_DATA3`      | `PA11`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit SAM E54 Xplained Pro](../machine/atsame54-xpro)

## Flashing

### OpenOCD

Programs are loaded onto the SAM E54 Xplained Pro with the onboard [Atmel EDBG](http://ww1.microchip.com/downloads/en/devicedoc/atmel-42096-microcontrollers-embedded-debugger_user-guide.pdf) hardware programmer, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the SAM E54 Xplained Pro board with your TinyGo code.

- Plug your SAM E54 Xplained Pro into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=atsame54-xpro [PATH TO YOUR PROGRAM]`

## Notes

You can use the USB port to the SAM E54 Xplained Pro as a serial port. `UART0` refers to this connection.
