---
title: "PJRC Teensy 4.0"
weight: 3
---

The [Teensy 4.0](https://www.pjrc.com/store/teensy40.html) is a small ARM development board based on the NXP [iMXRT1062](https://www.nxp.com/docs/en/nxp/data-sheets/IMXRT1060CEC.pdf) 32-bit 600 MHz ARM Cortex-M7.

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
| `D0`              | `PA3`        | `UART_RX_PIN`, `UART1_RX_PIN`, `SPI2_CS_PIN` |
| `D1`              | `PA2`        | `UART_TX_PIN`, `UART1_TX_PIN`, `SPI2_SDI_PIN` |
| `D2`              | `PD4`        |                   |
| `D3`              | `PD5`        |                   |
| `D4`              | `PD6`        |                   |
| `D5`              | `PD8`        |                   |
| `D6`              | `PB10`       |                   |
| `D7`              | `PB17`       | `UART2_RX_PIN`    |
| `D8`              | `PB16`       | `UART2_TX_PIN`    |
| `D9`              | `PB11`       |                   |
| `D10`             | `PB0`        | `SPI_CS_PIN`, `SPI1_CS_PIN` |
| `D11`             | `PB2`        | `SPI_SDO_PIN`, `SPI1_SDO_PIN` |
| `D12`             | `PB1`        | `SPI_SDI_PIN`, `SPI1_SDI_PIN` |
| `D13`             | `PB3`        | `LED`, `SPI_SCK_PIN`, `SPI1_SCK_PIN` |
| `D14`             | `PA18`       | `A0`, `UART3_TX_PIN` |
| `D15`             | `PA19`       | `A1`, `UART3_RX_PIN` |
| `D16`             | `PA23`       | `A2`, `UART4_RX_PIN`, `I2C2_SCL_PIN` |
| `D17`             | `PA22`       | `A3`, `UART4_TX_PIN`, `I2C2_SDA_PIN` |
| `D18`             | `PA17`       | `A4`, `I2C_SDA_PIN`, `I2C1_SDA_PIN` |
| `D19`             | `PA16`       | `A5`, `I2C_SCL_PIN`, `I2C1_SCL_PIN` |
| `D20`             | `PA26`       | `A6`, `UART5_TX_PIN` |
| `D21`             | `PA27`       | `A7`, `UART5_RX_PIN` |
| `D22`             | `PA24`       | `A8`              |
| `D23`             | `PA25`       | `A9`              |
| `D24`             | `PA12`       | `A10`, `UART6_TX_PIN`, `I2C3_SCL_PIN` |
| `D25`             | `PA13`       | `A11`, `UART6_RX_PIN`, `I2C3_SDA_PIN` |
| `D26`             | `PA30`       | `A12`, `SPI2_SDO_PIN` |
| `D27`             | `PA31`       | `A13`, `SPI2_SCK_PIN` |
| `D28`             | `PC18`       | `UART7_RX_PIN`    |
| `D29`             | `PD31`       | `UART7_TX_PIN`    |
| `D30`             | `PC23`       |                   |
| `D31`             | `PC22`       |                   |
| `D32`             | `PB12`       |                   |
| `D33`             | `PD7`        |                   |
| `D34`             | `PC15`       | `SPI3_SDI_PIN`    |
| `D35`             | `PC14`       | `SPI3_SDO_PIN`    |
| `D36`             | `PC13`       | `SPI3_CS_PIN`     |
| `D37`             | `PC12`       | `SPI3_SCK_PIN`    |
| `D38`             | `PC17`       |                   |
| `D39`             | `PC16`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Teensy 4.0](../machine/teensy40)

## Flashing

### teensy_loader_cli

In order to flash your TinyGo programs onto the Teensy 4.0 board, you will need to install the [teensy_loader_cli](https://github.com/PaulStoffregen/teensy_loader_cli) on your machine.

### CLI Flashing

- Plug your Teensy 4.0 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=teensy36 [PATH TO YOUR PROGRAM]
    ```

- The Teensy 4.0 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Teensy 4.0 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Teensy 4.0 board ready to receive code.
- The Teensy 4.0 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=teensy40 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Teensy 4.0 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You cannot yet use the USB port to the Teensy 4.0 as a USB CDC serial port.
