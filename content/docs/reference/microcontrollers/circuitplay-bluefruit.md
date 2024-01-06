---
title: "Adafruit Circuit Playground Bluefruit"
weight: 3
---

The [Circuit Playground Bluefruit](https://www.adafruit.com/product/4333) is small ARM development board based on the Nordic Semiconductor [nRF52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

## Peripherals and Drivers

- [nRF52840](https://github.com/tinygo-org/bluetooth) Bluetooth
- [LIS3DH](https://pkg.go.dev/tinygo.org/x/drivers/lis3dh) IMU chip (acceleration, tap detection, free-fall detection)
- MEMS microphone
- [Thermistor](https://pkg.go.dev/tinygo.org/x/drivers/thermistor) temperature sensor
- Light sensor (phototransistor)
- [WS2812](https://pkg.go.dev/tinygo.org/x/drivers/ws2812) Neopixel via the `D8` pin, [example](https://github.com/tinygo-org/drivers/tree/release/examples/ws2812)
- Buttons

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
| Bluetooth | YES | YES |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D0`              | `P0_30`      | `A6`, `UART_RX_PIN` |
| `D1`              | `P0_14`      | `A7`, `UART_TX_PIN` |
| `D2`              | `P0_05`      | `A5`, `SDA_PIN`   |
| `D3`              | `P0_04`      | `A4`, `SCL_PIN`   |
| `D4`              | `P1_02`      | `BUTTONA`, `BUTTON` |
| `D5`              | `P1_15`      | `BUTTONB`, `BUTTON1` |
| `D6`              | `P0_02`      | `A1`              |
| `D7`              | `P1_06`      | `SLIDER`          |
| `D8`              | `P0_13`      | `NEOPIXELS`, `WS2812` |
| `D9`              | `P0_29`      | `A2`              |
| `D10`             | `P0_03`      | `A3`              |
| `D11`             | `P1_04`      |                   |
| `D12`             | `P0_26`      |                   |
| `D13`             | `P1_14`      | `LED`             |
| `A8`              | `P0_28`      | `LIGHTSENSOR`     |
| `A9`              | `P0_31`      | `TEMPSENSOR`      |
| `SDA1_PIN`        | `P1_10`      |                   |
| `SCL1_PIN`        | `P1_12`      |                   |
| `SPI0_SCK_PIN`    | `P0_19`      |                   |
| `SPI0_SDO_PIN`    | `P0_21`      |                   |
| `SPI0_SDI_PIN`    | `P0_23`      |                   |
| `PDM_CLK_PIN`     | `P0_17`      |                   |
| `PDM_DIN_PIN`     | `P0_16`      |                   |

## Machine Package Docs

[Documentation for the machine package for the Circuit Playground Bluefruit](../machine/circuitplay-bluefruit)

## Flashing

### UF2

The Circuit Playground Bluefruit comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

**PLEASE NOTE** that for a good experience using TinyGo on your board you must be running version 0.4.1 or above of the UF2 bootloader on the board. For more information, [see below](#updating-the-uf2-bootloader)

### CLI Flashing

- Plug your Circuit Playground Bluefruit into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=circuitplay-bluefruit [PATH TO YOUR PROGRAM]
    ```

- The Circuit Playground Bluefruit board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Circuit Playground Bluefruit board to receive code, try this:

- Press the "RESET" button on the board two times to get the Circuit Playground Bluefruit board ready to receive code.
- The Circuit Playground Bluefruit board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=circuitplay-bluefruit [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Circuit Playground Bluefruit board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Circuit Playground Bluefruit as a serial port. `UART0` refers to this connection.

## Updating the UF2 bootloader

This board uses a UF2 bootloader created by Adafruit: https://github.com/adafruit/Adafruit_nRF52_Bootloader

We recommend bootloader version 0.4.1 or above. You can check what version is installed on your board by double-clicking the button on the board to launch the bootloader. When you do this, a USB volume that should automatically be mounted on your computer. Check the file named "INFO_UF2.TXT" on that drive. The bootloader firmware version should be listed in that file, for example:

```shell
UF2 Bootloader 0.4.1 lib/nrfx (v2.0.0) lib/tinyusb (0.6.0-272-g4e6aa0d8) lib/uf2 (remotes/origin/configupdate-9-gadbb8c7)
Model: Adafruit Circuit Playground nRF52840
Board-ID: nRF52840-CircuitPlayground-revD
SoftDevice: S140 version 6.1.1
Date: Jan 19 2021
```

To update the bootloader, you will need to install the `adafruit-nrfutil` program.

You can install it by running:

```shell
pip3 install --user adafruit-nrfutil
```

Once you have installed the `adafruit-nrfutil` program, download the firmware here:
https://github.com/adafruit/Adafruit_nRF52_Bootloader/releases/download/0.4.1/circuitplayground_nrf52840_bootloader-0.4.1.zip

Unzip the files in this zip file and save them to a convenient location.

Plug in your Circuit Playground Bluefruit board to your computer's USB port.

Now we are ready to update the firmware. Run a command something like the following, adjusting for any difference based on where you have saved the files, and what your serial port is named:

```shell
 adafruit-nrfutil --verbose dfu serial --package circuitplayground_nrf52840_bootloader-0.4.1_s140_6.1.1.zip -p /dev/ttyACM0 -b 115200 --singlebank --touch 1200
```

Note that you should be flashing the board using the zip file that was contained within the zip file that you downloaded, NOT the file that you downloaded.

Once you have flashed the board with the `adafruit-nrfutil` it should restart the board with the new bootloader. You only need to do this update once, and then from that point on the new bootloader will be active.
