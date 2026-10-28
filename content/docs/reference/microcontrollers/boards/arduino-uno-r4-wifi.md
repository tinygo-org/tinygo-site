---
title: "Arduino Uno R4 WiFi"
weight: 3
---

The [Arduino Uno R4 WiFi](https://docs.arduino.cc/hardware/uno-r4-wifi/) is a development board based on the Renesas [RA4M1](https://www.renesas.com/en/products/microcontrollers-microprocessors/ra-cortex-m-mcus/ra4m1-32-bit-microcontrollers-48mhz-arm-cortex-m4-and-lcd-controller-and-cap-touch-hmi) 32-bit ARM Cortex-M4 microcontroller with 32KB of SRAM and 256KB of flash. It also has an Espressif [ESP32-S3](https://www.espressif.com/en/products/socs/esp32-s3) coprocessor that handles USB and wireless.

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
| `D0`              | `P3_01`      | `UART2_RX_PIN`    |
| `D1`              | `P3_02`      | `UART2_TX_PIN`    |
| `D2`              | `P1_04`      |                   |
| `D3`              | `P1_05`      | `PWM_D3_PIN`      |
| `D4`              | `P1_06`      |                   |
| `D5`              | `P1_07`      | `PWM_D5_PIN`      |
| `D6`              | `P1_11`      | `PWM_D6_PIN`      |
| `D7`              | `P1_12`      |                   |
| `D8`              | `P3_04`      |                   |
| `D9`              | `P3_03`      | `PWM_D9_PIN`      |
| `D10`             | `P1_03`      | `PWM_D10_PIN`, `SPI_CS_PIN`, `CAN_TX_PIN`, `ADC_CH19_PIN` |
| `D11`             | `P4_11`      | `PWM_D11_PIN`, `SPI_SDO_PIN` |
| `D12`             | `P4_10`      | `SPI_SDI_PIN`     |
| `D13`             | `P1_02`      | `LED`, `SPI_SCK_PIN`, `CAN_RX_PIN`, `ADC_CH20_PIN` |
| `D14`             | `P0_14`      | `A0`, `DAC0_PIN`, `ADC_CH9_PIN` |
| `D15`             | `P0_00`      | `A1`, `ADC_CH0_PIN` |
| `D22`             | `P1_09`      | `UART_TX_PIN`, `UART1_TX_PIN` |
| `D23`             | `P1_10`      | `UART_RX_PIN`, `UART1_RX_PIN` |
| `D24`             | `P5_01`      | `UART3_TX_PIN`, `ADC_CH17_PIN` |
| `D25`             | `P5_02`      | `UART3_RX_PIN`, `ADC_CH18_PIN` |
| `D26`             | `P4_00`      | `SCL1_PIN`        |
| `D27`             | `P4_01`      | `SDA1_PIN`        |
| `A2`              | `P0_01`      | `ADC_CH1_PIN`     |
| `A3`              | `P0_02`      | `ADC_CH2_PIN`     |
| `A4`              | `P1_01`      | `SDA_PIN`, `ADC_CH21_PIN` |
| `A5`              | `P1_00`      | `SCL_PIN`, `ADC_CH22_PIN` |
| `AVCC_MEASURE_PIN` | `P5_00`      | `ADC_CH16_PIN`    |

## Machine Package Docs

[Documentation for the machine package for the Arduino Uno R4 WiFi](../../machine/arduino-uno-r4-wifi)

## Flashing

### CLI Flashing

The Uno R4 WiFi is flashed using the Arduino fork of [bossac](https://github.com/arduino/BOSSA). It must be in your `PATH` as `bossac`.

- Plug your Uno R4 WiFi board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the board with the blinky example:

    ```shell
    tinygo flash -target=arduino-uno-r4-wifi examples/blinky1
    ```

- The Uno R4 WiFi board should restart and then begin running your program.

## Notes

The serial console uses `UART1`, which the ESP32-S3 bridges to the USB port. `UART2` is on the `D0` and `D1` header pins.
