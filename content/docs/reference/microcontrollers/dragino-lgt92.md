---
title: "Dragino LGT-92"
weight: 3
---

The [Dragino LGT-92](https://www.dragino.com/products/lora-lorawan-end-node/item/142-lgt-92.html) includes a low power GPS module L76-L and 9-axis accelerometer for motion and attitude detection. It is based on the ST Micro [STM32L072CZT6](https://www.st.com/en/microcontrollers-microprocessors/stm32l072cz.html) SoC.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | YES | Not yet |

## Machine Package Docs

[Documentation for the machine package for the Dragino LGT-92](../machine/lgt-92)

## Flashing

### OpenOCD

Programs are loaded onto the Dragino LGT-92 with a separate hardware programmer such as the [STLink v2](https://www.st.com/en/development-tools/st-link-v2.html) board, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Dragino LGT-92 board with your TinyGo code.

- Plug your STLink v2 programmer into your computer's USB port.
- Plug your Dragino LGT-92 into the STLink v2 programmer using the Dragino LGT-92 `SWIO`, `SWCLK`, `3V3` and `GND` pins.
- Build and flash your TinyGo program using `tinygo flash -target=lgt92`
