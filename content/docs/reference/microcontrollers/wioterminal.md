---
title: "Seeed Wio Terminal"
weight: 3
---

The [Seeed Wio Terminal](https://www.seeedstudio.com/Wio-Terminal-p-4509.html) is a tiny ARM development board based on the Atmel [ATSAMD51P20](https://www.microchip.com/wwwproducts/en/ATSAMD51P20A) family of SoC.

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
| `ADC0`            | `PB08`       | `D0`, `A0`, `BCM27`, `RPI_A0` |
| `ADC1`            | `PB09`       | `D1`, `A1`, `BCM22`, `RPI_A1` |
| `ADC2`            | `PA07`       | `D2`, `A2`, `BCM13`, `BCM23`, `RPI_A2` |
| `ADC3`            | `PB04`       | `D3`, `A3`, `BCM24`, `RPI_A3` |
| `ADC4`            | `PB05`       | `D4`, `A4`, `BCM25`, `RPI_A4` |
| `ADC5`            | `PB06`       | `D5`, `A5`, `BCM12`, `RPI_A5` |
| `ADC6`            | `PA04`       | `D6`, `A6`, `RPI_A6` |
| `ADC7`            | `PB07`       | `D7`, `A7`, `BCM16`, `RPI_A7` |
| `ADC8`            | `PA06`       | `D8`, `A8`, `BCM26`, `RPI_A8` |
| `LED`             | `PA15`       | `PIN_LED_13`, `PIN_LED_RXL`, `PIN_LED_TXL`, `PIN_LED`, `PIN_LED2`, `PIN_LED3`, `LED_BUILTIN`, `PIN_NEOPIXEL` |
| `BUTTON`          | `PC26`       | `BUTTON_1`, `WIO_KEY_A` |
| `BCM0`            | `PA13`       | `PIN_WIRE1_SDA`, `SDA1`, `PIN_GYROSCOPE_WIRE_SDA`, `WIO_LIS3DH_SDA`, `SDA0_PIN` |
| `BCM1`            | `PA12`       | `PIN_WIRE1_SCL`, `SCL1`, `PIN_GYROSCOPE_WIRE_SCL`, `WIO_LIS3DH_SCL`, `SCL0_PIN` |
| `BCM2`            | `PA17`       | `PIN_WIRE_SDA`, `SDA`, `SDA1_PIN`, `SDA_PIN` |
| `BCM3`            | `PA16`       | `PIN_WIRE_SCL`, `SCL`, `SCL1_PIN`, `SCL_PIN` |
| `BCM4`            | `PB14`       | `GCLK0`           |
| `BCM5`            | `PB12`       | `GCLK1`           |
| `BCM6`            | `PB13`       | `GCLK2`           |
| `BCM7`            | `PA05`       | `PIN_DAC1`        |
| `BCM8`            | `PB01`       | `PIN_SPI_SS`, `SS` |
| `BCM9`            | `PB00`       | `PIN_SPI_SDI`, `SDI`, `SPI0_SDI_PIN` |
| `BCM10`           | `PB02`       | `PIN_SPI_SDO`, `SDO`, `SPI0_SDO_PIN` |
| `BCM11`           | `PB03`       | `PIN_SPI_SCK`, `SCK`, `SPI0_SCK_PIN` |
| `BCM14`           | `PB27`       | `PIN_SERIAL1_RX`, `UART_RX_PIN` |
| `BCM15`           | `PB26`       | `PIN_SERIAL1_TX`, `UART_TX_PIN` |
| `BCM17`           | `PA02`       | `PIN_DAC0`        |
| `BCM18`           | `PB28`       | `FPC1`            |
| `BCM19`           | `PA20`       | `PIN_I2S_FS`, `I2S_LRCLK` |
| `BCM20`           | `PA21`       | `PIN_I2S_SDI`, `I2S_SDIN` |
| `BCM21`           | `PA22`       | `PIN_I2S_SDO`, `I2S_SDOUT` |
| `FPC2`            | `PB17`       |                   |
| `FPC3`            | `PB29`       |                   |
| `FPC4`            | `PA14`       |                   |
| `FPC5`            | `PC01`       |                   |
| `FPC6`            | `PC02`       |                   |
| `FPC7`            | `PC03`       |                   |
| `FPC8`            | `PC04`       |                   |
| `FPC9`            | `PC31`       |                   |
| `FPC10`           | `PD00`       |                   |
| `PIN_USB_DM`      | `PA24`       | `USBCDC_DM_PIN`   |
| `PIN_USB_DP`      | `PA25`       | `USBCDC_DP_PIN`   |
| `PIN_USB_HOST_ENABLE` | `PA27`       |                   |
| `BUTTON_2`        | `PC27`       | `WIO_KEY_B`       |
| `BUTTON_3`        | `PC28`       | `WIO_KEY_C`       |
| `SWITCH_X`        | `PD20`       | `WIO_5S_UP`       |
| `SWITCH_Y`        | `PD12`       | `WIO_5S_LEFT`     |
| `SWITCH_Z`        | `PD09`       | `WIO_5S_RIGHT`    |
| `SWITCH_B`        | `PD08`       | `WIO_5S_DOWN`     |
| `SWITCH_U`        | `PD10`       | `WIO_5S_PRESS`    |
| `IRQ0`            | `PC20`       |                   |
| `BUZZER_CTR`      | `PD11`       | `WIO_BUZZER`      |
| `MIC_INPUT`       | `PC30`       | `WIO_MIC`         |
| `PIN_SERIAL2_RX`  | `PC23`       | `UART2_RX_PIN`    |
| `PIN_SERIAL2_TX`  | `PC22`       | `UART2_TX_PIN`    |
| `GYROSCOPE_INT1`  | `PC21`       | `WIO_LIS3DH_INT`  |
| `PIN_SPI1_SDI`    | `PC24`       | `SDI1`, `RTL8720D_SDI_PIN`, `SPI1_SDI_PIN` |
| `PIN_SPI1_SDO`    | `PB24`       | `SDO1`, `RTL8720D_SDO_PIN`, `SPI1_SDO_PIN` |
| `PIN_SPI1_SCK`    | `PB25`       | `SCK1`, `RTL8720D_SCK_PIN`, `SPI1_SCK_PIN` |
| `PIN_SPI1_SS`     | `PC25`       | `SS1`, `RTL8720D_SS_PIN` |
| `PIN_SPI2_SDI`    | `PC18`       | `SDI2`, `SDCARD_SDI_PIN`, `SPI2_SDI_PIN` |
| `PIN_SPI2_SDO`    | `PC16`       | `SDO2`, `SDCARD_SDO_PIN`, `SPI2_SDO_PIN` |
| `PIN_SPI2_SCK`    | `PC17`       | `SCK2`, `SDCARD_SCK_PIN`, `SPI2_SCK_PIN` |
| `PIN_SPI2_SS`     | `PC19`       | `SS2`, `SDCARD_SS_PIN` |
| `PIN_SPI3_SDI`    | `PB18`       | `SDI3`, `LCD_SDI_PIN`, `SPI3_SDI_PIN` |
| `PIN_SPI3_SDO`    | `PB19`       | `SDO3`, `LCD_SDO_PIN`, `SPI3_SDO_PIN` |
| `PIN_SPI3_SCK`    | `PB20`       | `SCK3`, `LCD_SCK_PIN`, `SPI3_SCK_PIN` |
| `PIN_SPI3_SS`     | `PB21`       | `SS3`, `LCD_SS_PIN` |
| `SDCARD_DET_PIN`  | `PD21`       |                   |
| `LCD_DC`          | `PC06`       |                   |
| `LCD_RESET`       | `PC07`       |                   |
| `LCD_BACKLIGHT`   | `PC05`       |                   |
| `LCD_XL`          | `PC10`       |                   |
| `LCD_YU`          | `PC11`       |                   |
| `LCD_XR`          | `PC12`       |                   |
| `LCD_YD`          | `PC13`       |                   |
| `PIN_QSPI_IO0`    | `PA08`       | `QSPI_DATA0`      |
| `PIN_QSPI_IO1`    | `PA09`       | `QSPI_DATA1`      |
| `PIN_QSPI_IO2`    | `PA10`       | `QSPI_DATA2`      |
| `PIN_QSPI_IO3`    | `PA11`       | `QSPI_DATA3`      |
| `PIN_QSPI_SCK`    | `PB10`       | `QSPI_SCK`        |
| `PIN_QSPI_CS`     | `PB11`       | `QSPI_CS`         |
| `PIN_I2S_SCK`     | `PB16`       | `I2S_BLCK`        |
| `RTL8720D_CHIP_PU` | `PA18`       |                   |
| `RTL8720D_GPIO0`  | `PA19`       |                   |
| `SWDCLK`          | `PA30`       |                   |
| `SWDIO`           | `PA31`       |                   |
| `SWO`             | `PB30`       |                   |
| `WIO_LIGHT`       | `PD01`       |                   |
| `WIO_IR`          | `PB31`       |                   |
| `OUTPUT_CTR_5V`   | `PC14`       |                   |
| `OUTPUT_CTR_3V3`  | `PC15`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Seeed Wio Terminal](../machine/wioterminal)

## Flashing

### UF2

The Wio Terminal comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Wio Terminal into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=wioterminal [PATH TO YOUR PROGRAM]
    ```

- The Wio Terminal board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Wio Terminal board to receive code, try this:

- [Slide the "RESET" switch on the board two times](https://wiki.seeedstudio.com/Wio-Terminal-Getting-Started/#faq) to get the Wio Terminal board ready to receive code.
- The Wio Terminal board will appear to your computer like a USB drive.
- Now try running the command as above:


```shell
tinygo flash -target=wioterminal [PATH TO YOUR PROGRAM]
```

Once you have updated your Wio Terminal board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Wio Terminal as a serial port. `UART0` refers to this connection.
