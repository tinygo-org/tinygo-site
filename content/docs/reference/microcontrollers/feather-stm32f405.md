---
title: "Adafruit Feather STM32F405 Express"
weight: 3
---

The [Feather STM32F405 Express](https://www.adafruit.com/product/4382) is a tiny ARM development board based on the ST Micro [STM32F405](https://www.st.com/resource/en/datasheet/dm00037051.pdf) family of microcontrollers.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | YES |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `PB11`       | `UART1_RX_PIN`, `UART0_RX_PIN`, `UART_RX_PIN`, `I2C2_SDA_PIN` |
| `D1`              | `PB10`       | `UART1_TX_PIN`, `UART0_TX_PIN`, `UART_TX_PIN`, `I2C2_SCL_PIN` |
| `D2`              | `PB3`        | `SPI2_SCK_PIN`    |
| `D3`              | `PB4`        | `SPI2_SDI_PIN`    |
| `D4`              | `PB5`        | `SPI2_SDO_PIN`    |
| `D5`              | `PC7`        | `UART2_RX_PIN`    |
| `D6`              | `PC6`        | `UART2_TX_PIN`    |
| `D7`              | `PA15`       |                   |
| `D8`              | `PC0`        | `LED_NEOPIXEL`, `WS2812` |
| `D9`              | `PB8`        | `I2C3_SDA_PIN`    |
| `D10`             | `PB9`        | `I2C3_SCL_PIN`    |
| `D11`             | `PC3`        |                   |
| `D12`             | `PC2`        |                   |
| `D13`             | `PC1`        | `LED_RED`, `LED_BUILTIN`, `LED` |
| `D14`             | `PB7`        | `UART3_RX_PIN`, `I2C1_SDA_PIN`, `I2C0_SDA_PIN`, `I2C_SDA_PIN`, `SDA_PIN` |
| `D15`             | `PB6`        | `UART3_TX_PIN`, `I2C1_SCL_PIN`, `I2C0_SCL_PIN`, `I2C_SCL_PIN`, `SCL_PIN` |
| `D16`             | `PA4`        | `A0`              |
| `D17`             | `PA5`        | `A1`, `SPI3_SCK_PIN` |
| `D18`             | `PA6`        | `A2`, `SPI3_SDI_PIN` |
| `D19`             | `PA7`        | `A3`, `SPI3_SDO_PIN` |
| `D20`             | `PC4`        | `A4`              |
| `D21`             | `PC5`        | `A5`              |
| `D22`             | `PA3`        | `A6`              |
| `D23`             | `PB13`       | `SPI1_SCK_PIN`, `SPI0_SCK_PIN`, `SPI_SCK_PIN` |
| `D24`             | `PB14`       | `SPI1_SDI_PIN`, `SPI0_SDI_PIN`, `SPI_SDI_PIN` |
| `D25`             | `PB15`       | `SPI1_SDO_PIN`, `SPI0_SDO_PIN`, `SPI_SDO_PIN` |
| `D26`             | `PC8`        |                   |
| `D27`             | `PC9`        |                   |
| `D28`             | `PC10`       |                   |
| `D29`             | `PC11`       |                   |
| `D30`             | `PC12`       |                   |
| `D31`             | `PD2`        |                   |
| `D32`             | `PB12`       |                   |
| `D33`             | `PC14`       |                   |
| `D34`             | `PC15`       |                   |
| `D35`             | `PA11`       |                   |
| `D36`             | `PA12`       |                   |
| `D37`             | `PA13`       |                   |
| `D38`             | `PA14`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Feather STM32F405](../machine/feather-stm32f405)

## Flashing

Flashing this board using its DFU bootloader can be a bit cumbersome, but it is possible - without requiring an external programmer - by pulling the B0 pin high (you can use the board's 3.3V output pin) at bootup when connected to host PC via USB. This puts the device in bootloader mode.
For more information see [Adafruit docs](https://learn.adafruit.com/adafruit-stm32f405-feather-express/dfu-bootloader-details).

Once in bootloader mode, the device can be programmed using the open-source tool `dfu-util`.

### CLI Flashing on Linux

You must first install the `dfu-util` program in order to flash the Adafruit Feather STM32F405 board.

    sudo apt update 
    sudo apt install dfu-util

### CLI Flashing on macOS

You must first install the `dfu-util` program in order to flash the Adafruit Feather STM32F405 board. You can do this using Homebrew on macOS:

    brew install dfu-util

### CLI Flashing on Windows

You must first install the `dfu-util` program in order to flash the Adafruit Feather STM32F405 board.

- Download dfu-util from the website here: https://dfu-util.sourceforge.net/releases/dfu-util-0.9-win64.zip
- Decompress the files to a directory such as `C:\dfu-util`
- Add `C:\dfu-util` to your `PATH`.

### Troubleshooting

If you run into trouble getting dfu-util installed and working on Windows, see the blog post at https://www.hanselman.com/blog/HowToFixDfuutilSTMWinUSBZadigBootloadersAndOtherFirmwareFlashingIssuesOnWindows.aspx
