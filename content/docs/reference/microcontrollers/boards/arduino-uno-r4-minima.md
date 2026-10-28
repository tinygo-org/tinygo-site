---
title: "Arduino Uno R4 Minima"
weight: 3
---

The [Arduino Uno R4 Minima](https://docs.arduino.cc/hardware/uno-r4-minima/) is a development board based on the Renesas [RA4M1](https://www.renesas.com/en/products/microcontrollers-microprocessors/ra-cortex-m-mcus/ra4m1-32-bit-microcontrollers-48mhz-arm-cortex-m4-and-lcd-controller-and-cap-touch-hmi) 32-bit ARM Cortex-M4 microcontroller with 32KB of SRAM and 256KB of flash.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | Not yet |
| I2C       | YES | Not yet |
| ADC       | YES | YES |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `P3_01`      | `UART_RX_PIN`, `UART1_RX_PIN` |
| `D1`              | `P3_02`      | `UART_TX_PIN`, `UART1_TX_PIN` |
| `D2`              | `P1_05`      |                   |
| `D3`              | `P1_04`      | `PWM_D3_PIN`      |
| `D4`              | `P1_03`      | `CAN_TX_PIN`, `ADC_CH19_PIN` |
| `D5`              | `P1_02`      | `PWM_D5_PIN`, `CAN_RX_PIN`, `ADC_CH20_PIN` |
| `D6`              | `P1_06`      | `PWM_D6_PIN`      |
| `D7`              | `P1_07`      |                   |
| `D8`              | `P3_04`      |                   |
| `D9`              | `P3_03`      | `PWM_D9_PIN`      |
| `D10`             | `P1_12`      | `PWM_D10_PIN`, `SPI_CS_PIN` |
| `D11`             | `P1_09`      | `PWM_D11_PIN`, `SPI_SDO_PIN` |
| `D12`             | `P1_10`      | `SPI_SDI_PIN`     |
| `D13`             | `P1_11`      | `LED`, `SPI_SCK_PIN` |
| `D14`             | `P0_14`      | `A0`, `DAC0_PIN`, `ADC_CH9_PIN` |
| `D15`             | `P0_00`      | `A1`, `ADC_CH0_PIN` |
| `A2`              | `P0_01`      | `ADC_CH1_PIN`     |
| `A3`              | `P0_02`      | `ADC_CH2_PIN`     |
| `A4`              | `P1_01`      | `SDA_PIN`, `ADC_CH21_PIN` |
| `A5`              | `P1_00`      | `SCL_PIN`, `ADC_CH22_PIN` |
| `SWDIO_PIN`       | `P1_08`      |                   |
| `SWCLK_PIN`       | `P3_00`      |                   |
| `LED_TX`          | `P0_12`      |                   |
| `LED_RX`          | `P0_13`      |                   |
| `UART_SWD_TX_PIN` | `P5_01`      | `ADC_CH17_PIN`    |
| `UART_SWD_RX_PIN` | `P5_02`      | `ADC_CH18_PIN`    |
| `AVCC_MEASURE_PIN` | `P5_00`      | `ADC_CH16_PIN`    |

## Machine Package Docs

[Documentation for the machine package for the Arduino Uno R4 Minima](../../machine/arduino-uno-r4-minima)

## Flashing

### CLI Flashing

The Uno R4 Minima is flashed using [dfu-util](https://dfu-util.sourceforge.net/).

- Plug your Uno R4 Minima board into your computer's USB port.
- Double tap the reset button to enter the bootloader.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the board with the blinky example:

    ```shell
    tinygo flash -target=arduino-uno-r4-minima examples/blinky1
    ```

- The Uno R4 Minima board should restart and then begin running your program.

## Notes

The serial console uses `UART1` on the `D0` and `D1` header pins, not the USB port.
