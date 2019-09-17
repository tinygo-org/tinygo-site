---
title: "Nucleo F103RB"
weight: 3
---

The [Nucleo F103RB](https://www.st.com/en/evaluation-tools/nucleo-f103rb.html) is a rather affordable ARM development board based on the ST Micro [STM32F103RB](https://www.st.com/en/microcontrollers/stm32f103rb.html) SoC providing lots of GPIO pins and an onboard programmer/debugger with mini-USB connector.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | YES |
| I2C       | YES | YES |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |

## Flashing

### OpenOCD

Programs are loaded onto the Nucleo F103RB with the onboard STLink v2-1 programmer connected via USB to the host, using the `openocd` command line utility program to perform the board flashing. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the Nucleo F103RB board with your TinyGo code.

- Plug the onboard STLink v2-1 programmer into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=nucleo-f103rb` - if you encounter OpenOCD connectivity errors try holding the black reset button down while flashing.
