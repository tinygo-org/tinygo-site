---
title: "Seeed Sipeed MAix Bit"
weight: 3
---

The [Sipeed MAix Bit](https://www.seeedstudio.com/Sipeed-MAix-BiT-for-RISC-V-AI-IoT-p-2872.html) is a low-cost development board featuring the [Kendryte K210](https://canaan.io/product/kendryteai) [RISC-V](https://riscv.org/) (RV64GC) chip.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | NO | NO |
| PWM      | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `P00`        |                   |
| `D1`              | `P01`        |                   |
| `D2`              | `P02`        |                   |
| `D3`              | `P03`        |                   |
| `D4`              | `P04`        | `UART_RX_PIN`     |
| `D5`              | `P05`        | `UART_TX_PIN`     |
| `D6`              | `P06`        |                   |
| `D7`              | `P07`        |                   |
| `D8`              | `P08`        |                   |
| `D9`              | `P09`        |                   |
| `D10`             | `P10`        |                   |
| `D11`             | `P11`        |                   |
| `D12`             | `P12`        | `LED2`, `LED_GREEN` |
| `D13`             | `P13`        | `LED`, `LED1`, `LED_RED` |
| `D14`             | `P14`        | `LED3`, `LED_BLUE` |
| `D15`             | `P15`        |                   |
| `D16`             | `P16`        |                   |
| `D17`             | `P17`        |                   |
| `D18`             | `P18`        |                   |
| `D19`             | `P19`        |                   |
| `D20`             | `P20`        |                   |
| `D21`             | `P21`        |                   |
| `D22`             | `P22`        |                   |
| `D23`             | `P23`        |                   |
| `D24`             | `P24`        |                   |
| `D25`             | `P25`        |                   |
| `D26`             | `P26`        | `SPI0_SDI_PIN`    |
| `D27`             | `P27`        | `SPI0_SCK_PIN`    |
| `D28`             | `P28`        | `SPI0_SDO_PIN`    |
| `D29`             | `P29`        |                   |
| `D30`             | `P30`        |                   |
| `D31`             | `P31`        |                   |
| `D32`             | `P32`        |                   |
| `D33`             | `P33`        |                   |
| `D34`             | `P34`        | `I2C0_SDA_PIN`    |
| `D35`             | `P35`        | `I2C0_SCL_PIN`    |

## Machine Package Docs

[Documentation for the machine package for the Sipeed MAix Bit](../machine/maixbit)

## Flashing

### Kflash.py

Programs are loaded onto the MAix Bit using the `kflash.py` command line utility program. You must install this program before you will be able to flash the MAix Bit board with your TinyGo code.

The latest version of the `kflash.py` can be installed using `pip3 install kflash`.

- Plug your MAix Bit into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=maixbit [PATH TO YOUR PROGRAM]`

## Troubleshooting

You may get the following error when flashing the MAix Bit:

```shell
error: unable to locate a serial port
```
To resolve this, just specify the MAix Bit's serial port when flashing using `tinygo flash -target=maixbit -port=[MAIXBIT PORT] [PATH TO YOUR PROGRAM]`.
