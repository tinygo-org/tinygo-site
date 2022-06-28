---
title: "Arduino Nano 33 BLE"
weight: 3
---

The [Arduino Nano33 BLE](https://store.arduino.cc/arduino-nano-33-ble) is a very small ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

There is also the [Arduino Nano33 BLE Sense](nano-33-ble-sense) which is the exact same board but with additional onboard sensors.

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
| `D2`              | `P1_11`      |                   |
| `D3`              | `P1_12`      |                   |
| `D4`              | `P1_15`      |                   |
| `D5`              | `P1_13`      |                   |
| `D6`              | `P1_14`      |                   |
| `D7`              | `P0_23`      |                   |
| `D8`              | `P0_21`      |                   |
| `D9`              | `P0_27`      |                   |
| `D10`             | `P1_02`      |                   |
| `D11`             | `P1_01`      | `SPI0_SDO_PIN`    |
| `D12`             | `P1_08`      | `SPI0_SDI_PIN`    |
| `D13`             | `P0_13`      | `LED`, `LED_BUILTIN`, `SPI0_SCK_PIN` |
| `A0`              | `P0_04`      |                   |
| `A1`              | `P0_05`      |                   |
| `A2`              | `P0_30`      |                   |
| `A3`              | `P0_29`      |                   |
| `A4`              | `P0_31`      | `SDA0_PIN`        |
| `A5`              | `P0_02`      | `SCL0_PIN`        |
| `A6`              | `P0_28`      |                   |
| `A7`              | `P0_03`      |                   |
| `LED1`            | `P0_24`      | `LED_RED`         |
| `LED2`            | `P0_16`      | `LED_GREEN`       |
| `LED3`            | `P0_06`      | `LED_BLUE`        |
| `LED_PWR`         | `P1_09`      |                   |
| `UART_RX_PIN`     | `P1_10`      |                   |
| `UART_TX_PIN`     | `P1_03`      |                   |
| `SDA_PIN`         | `P0_14`      | `SDA1_PIN`        |
| `SCL_PIN`         | `P0_15`      | `SCL1_PIN`        |
| `I2C_PULLUP`      | `P1_00`      |                   |
| `APDS_INT`        | `P0_19`      |                   |
| `LSM_PWR`         | `P0_22`      | `LPS_PWR`, `HTS_PWR` |
| `MIC_PWR`         | `P0_17`      |                   |
| `MIC_CLK`         | `P0_26`      |                   |
| `MIC_DIN`         | `P0_25`      |                   |

## Onboard sensors

* 9-axis IMU: [LSM9DS1](https://github.com/tinygo-org/drivers/tree/release/lsm9ds1)

## Machine Package Docs

[Documentation for the machine package for the Arduino Nano33 BLE](../machine/nano-33-ble)

## Installing BOSSA

In order to flash your TinyGo programs onto the Arduino Nano33 BLE board, you will need to install the "bossac_arduino2" command line utility which is a special build of the [BOSSA command line utilities](https://github.com/shumatech/BOSSA).

### macOS

If you have a Mac computer with an Intel processor, download the `bossac_arduino2` program from http://downloads.arduino.cc/tools/bossac-1.9.1-arduino2-osx.tar.gz

Extract the downloaded file to a directory on your computer.

Make sure to add that directory into your PATH.

### Linux

Download the `bossac_arduino2` program from http://downloads.arduino.cc/tools/bossac-1.9.1-arduino2-linux64.tar.gz

Extract the downloaded file to a directory on your computer.

Make sure to add that directory into your PATH.

### Windows

Download the `bossac_arduino2` program from http://downloads.arduino.cc/tools/bossac-1.9.1-arduino2-windows.tar.gz

Extract the downloaded file to a directory on your computer.

Make sure to add that directory into your PATH.

## Flashing

Once you have installed the needed BOSSA command line utility, as in the previous section, you are ready to build and flash code to your Arduino Nano33 BLE board.

### CLI Flashing

- Plug your Arduino Nano33 BLE board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the Arduino Nano33 BLE with the blinky1 example:

    ```shell
    tinygo flash -target=nano-33-ble examples/blinky1
    ```

- The Arduino Nano33 BLE board should restart and then begin running your program.

### Troubleshooting

Instructions needed here.

## Bluetooth

Nordic Semiconductor's SoftDevice (s140v7) must be flashed first to enable use of [bluetooth](https://github.com/tinygo-org/bluetooth) on this board.

SoftDevice overwrites original bootloader and flashing method described above is not avalable anymore.
Instead, please use [debug](../../guides/debugging.md) probe and flash your code with `nano-33-ble-s140v7` target.

## Notes

You can use the USB port to the Arduino Nano33 BLE as a serial port. `UART0` refers to this connection.
