---
title: "Adafruit Feather nRF52840 Sense"
weight: 3
---

The [Adafruit Feather nRF52840 Sense](https://www.adafruit.com/product/4516) is a small ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

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
| `D0`              | `P0_25`      | `UART_TX_PIN`     |
| `D1`              | `P0_24`      | `UART_RX_PIN`     |
| `D2`              | `P0_10`      |                   |
| `D3`              | `P1_11`      |                   |
| `D4`              | `P1_10`      | `LED2`            |
| `D5`              | `P1_08`      |                   |
| `D6`              | `P0_07`      |                   |
| `D7`              | `P1_02`      | `BUTTON`          |
| `D8`              | `P0_16`      | `NEOPIXEL`, `WS2812` |
| `D9`              | `P0_26`      |                   |
| `D10`             | `P0_27`      |                   |
| `D11`             | `P0_06`      |                   |
| `D12`             | `P0_08`      |                   |
| `D13`             | `P1_09`      | `LED`, `LED1`     |
| `D14`             | `P0_04`      | `A0`              |
| `D15`             | `P0_05`      | `A1`              |
| `D16`             | `P0_30`      | `A2`              |
| `D17`             | `P0_28`      | `A3`              |
| `D18`             | `P0_02`      | `A4`              |
| `D19`             | `P0_03`      | `A5`              |
| `D20`             | `P0_29`      | `A6`              |
| `D21`             | `P0_31`      | `A7`              |
| `D22`             | `P0_12`      | `SDA_PIN`         |
| `D23`             | `P0_11`      | `SCL_PIN`         |
| `D24`             | `P0_15`      | `SPI0_SDI_PIN`    |
| `D25`             | `P0_13`      | `SPI0_SDO_PIN`    |
| `D26`             | `P0_14`      | `SPI0_SCK_PIN`    |
| `D27`             | `P0_19`      | `QSPI_SCK`        |
| `D28`             | `P0_20`      | `QSPI_CS`         |
| `D29`             | `P0_17`      | `QSPI_DATA0`      |
| `D30`             | `P0_22`      | `QSPI_DATA1`      |
| `D31`             | `P0_23`      | `QSPI_DATA2`      |
| `D32`             | `P0_21`      | `QSPI_DATA3`      |
| `D33`             | `P0_09`      |                   |

## Machine Package Docs

[Documentation for the machine package for the Adafruit Feather nRF52840 Sense](../machine/feather-nrf52840-sense)

## Flashing

### UF2

The Adafruit Feather nRF52840 Sense comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

**PLEASE NOTE** that for a good experience using TinyGo on your board you must be running version 0.4.1 or above of the UF2 bootloader on the board. For more information, [see below](#updating-the-uf2-bootloader)

### CLI Flashing

- Plug your Adafruit Feather nRF52840 Sense into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=feather-nrf52840-sense [PATH TO YOUR PROGRAM]
    ```

- The Adafruit Feather nRF52840 Sense board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Adafruit Feather nRF52840 Sense board to receive code, try this:

- Press the "RESET" button on the board two times to get the Adafruit Feather nRF52840 Sense board ready to receive code.
- The Adafruit Feather nRF52840 Sense board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=feather-nrf52840-sense [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Adafruit Feather nRF52840 Sense board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Adafruit Feather nRF52840 Sense as a serial port. `UART0` refers to this connection.

Bluetooth support is now available for the Adafruit Feather nRF52840 Sense board. See https://github.com/tinygo-org/bluetooth for more information.

## Updating the UF2 bootloader

This board uses a UF2 bootloader created by Adafruit: https://github.com/adafruit/Adafruit_nRF52_Bootloader

We recommend bootloader version 0.4.1 or above. You can check what version is installed on your board by double-clicking the button on the board to launch the bootloader. When you do this, a USB volume that should automatically be mounted on your computer. Check the file named "INFO_UF2.TXT" on that drive. The bootloader firmware version should be listed in that file, for example:

```
UF2 Bootloader 0.4.1 lib/nrfx (v2.0.0) lib/tinyusb (0.6.0-272-g4e6aa0d8) lib/uf2 (remotes/origin/configupdate-9-gadbb8c7)
Model: Adafruit CLUE nRF52840
Board-ID: nRF52840-CLUE-revA
SoftDevice: S140 version 6.1.1
Date: Jan 19 2021
```

To update the bootloader, you will need to install the `adafruit-nrfutil` program. 

You can install it by running:

```shell
pip3 install --user adafruit-nrfutil
```

Once you have installed the `adafruit-nrfutil` program, download the firmware here: 
https://github.com/adafruit/Adafruit_nRF52_Bootloader/releases/download/0.4.1/feather_nrf52840_sense_bootloader-0.4.1.zip

Unzip the files in this zip file and save them to a convenient location. 

Plug in your board to your computer's USB port.

Now we are ready to update the firmware. Run a command something like the following, adjusting for any difference based on where you have saved the files, and what your serial port is named:

```shell
 adafruit-nrfutil --verbose dfu serial --package feather_nrf52840_bootloader-0.4.1_s140_6.1.1.zip -p /dev/ttyACM0 -b 115200 --singlebank --touch 1200
```

Note that you should be flashing the board using the zip file that was contained within the zip file that you downloaded, NOT the file that you downloaded.

Once you have flashed the board with the `adafruit-nrfutil` it should restart the board with the new bootloader. You only need to do this update once, and then from that point on the new bootloader will be active.
