---
title: "Arduino Nano33 IoT"
weight: 3
---

The [Arduino Nano33 IoT](https://store.arduino.cc/nano-33-iot) is a very small ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of processors. It also has a NINA-W102 chip onboard which provides an wireless communication abilities based on the popular ESP32 family of wireless chips from Espressif.

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

[Documentation for the machine package for the Arduino Nano33 IoT](../machine/arduino-nano33)

## Pin Mapping

Since the pin labels are located on the other side of the board, here is a picture showing the pin numbering from the front side perspective:

![Arduino Nano33 IoT](../../images/nano33pinmap.jpg)

## Installing BOSSA

In order to flash your TinyGo programs onto the Arduino Nano33 IoT board, you will need to install the "bossac" command line utility which is part of the [BOSSA command line utilities](https://github.com/shumatech/BOSSA).

### macOS

You can use Homebrew to install the BOSSA command line interface by using the following command:

```shell
brew cask install bossa
```

Or if you  prefer, you can also download the installer from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-1.9.1.dmg

Once you have downloaded it, double click on the .dmg file to perform the installation.

### Linux

On Linux, install from source:

```shell
sudo apt install libreadline-dev libwxgtk3.0-* 
git clone https://github.com/shumatech/BOSSA.git
cd BOSSA
make
sudo cp bin/bossac /usr/local/bin
```

### Windows

You can download BOSSA from https://github.com/shumatech/BOSSA/releases/download/1.9.1/bossa-x64-1.9.1.msi

*VERY IMPORTANT*: During the installation, you much choose to install into `c:\Program Files`. The installer might have the wrong path, so edit it to match  `c:\Program Files`.

After the installation, you must add BOSSA to your PATH:

```shell
set PATH=%PATH%;"c:\Program Files\BOSSA";
```

Test that you have installed "BOSSA" correctly by running this command:

```shell
bossac --help
```

## Flashing

Once you have installed the needed BOSSA command line utility, as in the previous section, you are ready to build and flash code to your Arduino Nano33 IoT board.

### CLI Flashing on Linux

- Plug your Arduino Nano33 IoT board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the Arduino Nano33 IoT with the blinky1 example:

    ```
    tinygo flash -target=arduino-nano33 examples/blinky1
    ```

- The Arduino Nano33 IoT board should restart and then begin running your program.

### CLI Flashing on macOS

- Plug your Arduino Nano33 IoT board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the Arduino Nano33 IoT with the blinky1 example:

    ```
    tinygo flash -target=arduino-nano33 examples/blinky1
    ```

- The Arduino Nano33 IoT board should restart and then begin running your program.

### CLI Flashing on Windows

- Plug your Arduino Nano33 IoT board into your computer's USB port.
- Double tap the "RESET" button on the board.
- Wait until the Arduino Nano33 IoT board appears as a serial drive.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the Arduino Nano33 IoT with the blinky1 example:

    ```
    tinygo flash -target=arduino-nano33 examples/blinky1
    ```

- The Arduino Nano33 IoT board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Arduino Nano33 IoTboard to receive code, try this:

- Press the "RESET" button on the board two times to get the Arduino Nano33 IoT board ready to receive code.
- Now try running the `tinygo flash` command as above:

    ```shell
    tinygo flash -target=arduino-nano33 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Arduino Nano33 IoT board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Arduino Nano33 IoT as a serial port. `UART0` refers to this connection.

For information on how to use the built-in NINA-W102 wireless chip, please see the "espat" driver in the TinyGo drivers repository located at [https://github.com/tinygo-org/drivers/tree/master/espat](https://github.com/tinygo-org/drivers/tree/master/espat).
