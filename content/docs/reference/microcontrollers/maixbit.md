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
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | NO  | NO  |
| PWM       | YES | Not yet |
| USBDevice | ?   | ?   |

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
