---
title: "nice!nano"
weight: 3
---

The [nice!nano](https://nicekeyboards.com/products/nice-nano-v1-0) is a wireless, BLE enabled replacement for the Pro Micro powered by the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor.

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
| `D006`            | `P0_06`      | `UART_RX_PIN`     |
| `D008`            | `P0_08`      | `UART_TX_PIN`     |
| `D017`            | `P0_17`      | `SDA_PIN`         |
| `D020`            | `P0_20`      | `SCL_PIN`         |
| `D022`            | `P0_22`      | `SPI0_SCK_PIN`    |
| `D024`            | `P0_24`      | `SPI0_SDO_PIN`    |
| `D100`            | `P1_00`      | `SPI0_SDI_PIN`    |
| `D011`            | `P0_11`      |                   |
| `D104`            | `P1_04`      |                   |
| `D106`            | `P1_06`      |                   |
| `D004`            | `P0_04`      | `AIN2`            |
| `D013`            | `P0_13`      |                   |
| `D115`            | `P1_15`      |                   |
| `D113`            | `P1_13`      |                   |
| `D031`            | `P0_31`      | `AIN7`            |
| `D029`            | `P0_29`      | `AIN5`            |
| `D002`            | `P0_02`      | `AIN0`            |
| `D111`            | `P1_11`      |                   |
| `D010`            | `P0_10`      |                   |
| `D009`            | `P0_09`      |                   |
| `D026`            | `P0_26`      |                   |
| `D012`            | `P0_12`      |                   |
| `D101`            | `P1_01`      |                   |
| `D102`            | `P1_02`      |                   |
| `D107`            | `P1_07`      |                   |
| `LED`             | `P0_15`      |                   |

## Machine Package Docs

[Documentation for the machine package for the nice!nano](../machine/nicenano)

## Flashing

### UF2

The nice!nano comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

**PLEASE NOTE** that for a good experience using TinyGo on your board you must be running version 0.4.1 or above of the UF2 bootloader on the board. For more information, [see below](#updating-the-uf2-bootloader)

### CLI Flashing

- Plug your nice!nano into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=nicenano [PATH TO YOUR PROGRAM]
    ```

- The nice!nano board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your nice!nano board to receive code, try this:

- Press the "RESET" button on the board two times to get the nice!nano board ready to receive code.
- The nicenano board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=nicenano [PATH TO YOUR PROGRAM]
    ```

Once you have updated your nice!nano board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the nice!nano as a serial port. `UART0` refers to this connection.

Bluetooth support is available for the nice!nano board. See https://github.com/tinygo-org/bluetooth for more information.

## Updating the UF2 bootloader

This board uses a UF2 bootloader created by Adafruit: https://github.com/adafruit/Adafruit_nRF52_Bootloader

We recommend bootloader version 0.4.1 or above. You can check what version is installed on your board by double-clicking the button on the board to launch the bootloader. When you do this, a USB volume that should automatically be mounted on your computer. Check the file named "INFO_UF2.TXT" on that drive. The bootloader firmware version should be listed in that file, for example:

```
UF2 Bootloader 0.4.1 lib/nrfx (v2.0.0) lib/tinyusb (0.6.0-272-g4e6aa0d8) lib/uf2 (remotes/origin/configupdate-9-gadbb8c7)
Model: Adafruit ItsyBitsy nRF52840 Express
Board-ID: nRF52840-ItsyBitsy-revA
SoftDevice: S140 version 6.1.1
Date: Jan 19 2021
```

To update the bootloader, you will need to install the `adafruit-nrfutil` program. 

You can install it by running:

```shell
pip3 install --user adafruit-nrfutil
```

Once you have installed the `adafruit-nrfutil` program, download the firmware here: 
https://github.com/adafruit/Adafruit_nRF52_Bootloader/releases/download/0.4.1/nice_nano_bootloader-0.4.1.zip

Unzip the files in this zip file and save them to a convenient location. 

Plug in your board to your computer's USB port.

Now we are ready to update the firmware. Run a command something like the following, adjusting for any difference based on where you have saved the files, and what your serial port is named:

```shell
 adafruit-nrfutil --verbose dfu serial --package nice_nano_bootloader-0.4.1_s140_6.1.1.zip -p /dev/ttyACM0 -b 115200 --singlebank --touch 1200
```

Note that you should be flashing the board using the zip file that was contained within the zip file that you downloaded, NOT the file that you downloaded.

Once you have flashed the board with the `adafruit-nrfutil` it should restart the board with the new bootloader. You only need to do this update once, and then from that point on the new bootloader will be active.
