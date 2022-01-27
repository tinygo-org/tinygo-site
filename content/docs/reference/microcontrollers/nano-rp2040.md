---
title: "Arduino Nano RP2040 Connect"
weight: 3
---

The [Nano RP2040 Connect](https://store.arduino.cc/nano-rp2040-connect) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller. 

Peripherals: 
- NINA-W102 chip with [wifinina](https://github.com/tinygo-org/drivers/tree/release/wifinina) firmware (wifi and bluetooth)
- [lsm6dox](https://github.com/tinygo-org/drivers/tree/release/lsm6dox) IMU chip (acceleration, rotation and temperature)
- microphone

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

[Documentation for the machine package for the Nano RP2040](../machine/nano-rp2040)

## Flashing

### UF2

The Nano RP2040 comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Plug your Nano RP2040 into your computer's USB port while shorting the pins REC and GND with a jumper wire.
- Once plugged in, remove the jumper pin.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=nano-rp2040 [PATH TO YOUR PROGRAM]
    ```

- The Nano RP2040 board should restart and run your program.

### Troubleshooting

Any troubleshooting tips go here.

## Notes

You cannot yet use the USB port to the Nano RP2040 as a serial port. Instead `UART0` refers to the TX/RX pins on the board itself.
