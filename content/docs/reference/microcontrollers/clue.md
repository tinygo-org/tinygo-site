---
title: "Adafruit CLUE"
weight: 3
---

The [Adafruit CLUE](https://www.adafruit.com/product/4500) is small ARM development board based on the Nordic Semiconductor [nrf52840](https://www.nordicsemi.com/eng/Products/nRF52840) processor. It has several built-in devices such as WS2812 "NeoPixel" LEDs, buttons, an accelerometer, and some other sensors.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |
| Bluetooth      | YES | YES |

## Machine Package Docs

[Documentation for the machine package for the Adafruit CLUE](../machine/adafruit-clue)

## Flashing

### UF2

The CLUE comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

**PLEASE NOTE** that for a good experience using TinyGo on your board you must be running version 0.4.1 or above of the UF2 bootloader on the board. For more information, [see below](#updating-the-uf2-bootloader)

### CLI Flashing

- Plug your CLUE into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=clue [PATH TO YOUR PROGRAM]
    ```

- The CLUE board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your CLUE board to receive code, try this:

- Press the "RESET" button on the board two times to get the CLUE board ready to receive code.
- The CLUE board will appear to your computer like a USB drive.
- Now try running the command:

    ```shell
    tinygo flash -target=clue [PATH TO YOUR PROGRAM]
    ```

Once you have updated your CLUE board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the CLUE as a serial port. `UART0` refers to this connection.

For an example that uses the built-in Neopixel LEDs, take a look at the TinyGo drivers repository located at [https://github.com/tinygo-org/drivers/tree/release/examples](https://github.com/tinygo-org/drivers)

Bluetooth support is now available for the Adafruit CLUE board. See https://github.com/tinygo-org/bluetooth for more information.

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
https://github.com/adafruit/Adafruit_nRF52_Bootloader/releases/download/0.4.1/clue_nrf52840_bootloader-0.4.1.zip

Unzip the files in this zip file and save them to a convenient location. 

Plug in your board to your computer's USB port.

Now we are ready to update the firmware. Run a command something like the following, adjusting for any difference based on where you have saved the files, and what your serial port is named:

```shell
 adafruit-nrfutil --verbose dfu serial --package clue_nrf52840_bootloader-0.4.1_s140_6.1.1.zip -p /dev/ttyACM0 -b 115200 --singlebank --touch 1200
```

Note that you should be flashing the board using the zip file that was contained within the zip file that you downloaded, NOT the file that you downloaded.

Once you have flashed the board with the `adafruit-nrfutil` it should restart the board with the new bootloader. You only need to do this update once, and then from that point on the new bootloader will be active.
