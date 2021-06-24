---
title: "Arduino Nano33 BLE"
weight: 3
---

The [Arduino Nano33 BLE](h) is a very small ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Machine Package Docs

[Documentation for the machine package for the Arduino Nano33 IoT](../machine/arduino-nano33-ble)

## Installing BOSSA

In order to flash your TinyGo programs onto the Arduino Nano33 BLE board, you will need to install the "arduino_bossac" command line utility which is a special build of the [BOSSA command line utilities](https://github.com/shumatech/BOSSA).

### macOS

Instructions needed here.

### Linux

Instructions needed here.

### Windows

Instructions needed here.

## Flashing

Once you have installed the needed BOSSA command line utility, as in the previous section, you are ready to build and flash code to your Arduino Nano33 BLE board.

### CLI Flashing

- Plug your Arduino Nano33 BLE board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the Arduino Nano33 BLE with the blinky1 example:

    ```shell
    tinygo flash -target=arduino-nano33-ble examples/blinky1
    ```

- The Arduino Nano33 BLE board should restart and then begin running your program.

### Troubleshooting

Instructions needed here.

## Notes

You can use the USB port to the Arduino Nano33 BLE as a serial port. `UART0` refers to this connection.
