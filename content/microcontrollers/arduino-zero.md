---
title: "Arduino Zero"
weight: 3
---

The [Arduino Zero](https://store.arduino.cc/arduino-zero) is a very small ARM development board based on the Atmel [SAMD21](https://www.microchip.com/wwwproducts/en/ATSAMD21G18) family of processors.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | YES |
| PWM      | YES | YES |

## Pins

| Pin           | Hardware pin | PWM
| ------------- | ------------ | ----------------
| `D0`          | `PA11`       | `TCC1`, `TCC0`
| `D1`          | `PA10`       | `TCC1`, `TCC0`
| `D2`          | `PA14`       | `TCC0`
| `D3`          | `PA09`       | `TCC0`, `TCC1`
| `D4`          | `PA08`       | `TCC0`, `TCC1`
| `D5`          | `PA15`       | `TCC0`
| `D6`          | `PA20`       | `TCC0`
| `D7`          | `PA21`       | `TCC0`
| `D8`          | `PA06`       | `TCC1`
| `D9`          | `PA07`       | `TCC1`
| `D10`         | `PA18`       | `TCC0`
| `D11`         | `PA16`       | `TCC2`, `TCC0`
| `D12`         | `PA19`       | `TCC0`
| `D13`, `LED`  | `PA17`       | `TCC2`, `TCC0`
| `LED2`        | `PA27`       |
| `LED3`        | `PB03`       |
| `AREF`        | `PA03`       |
| `A0`          | `PA02`       |
| `A1`          | `PB08`       |
| `A2`          | `PB09`       |
| `A3`          | `PA04`       | `TCC0`
| `A4`          | `PA05`       | `TCC0`
| `A5`          | `PA02`       |

## Machine Package Docs

[Documentation for the machine package for the Arduino Zero](../machine/arduino-zero)

## Installing BOSSA

In order to flash your TinyGo programs onto the Arduino Zero board, you will need to install the "bossac" command line utility which is part of the [BOSSA command line utilities](https://github.com/shumatech/BOSSA).

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

Once you have installed the needed BOSSA command line utility, as in the previous section, you are ready to build and flash code to your Arduino Zero board.

### CLI Flashing

- Plug your Arduino Zero board into your computer's USB port.
- Build and flash your TinyGo code using the `tinygo flash` command. This command flashes the Arduino Zero with the blinky1 example:

    ```
    tinygo flash -target=arduino-zero examples/blinky1
    ```

- The Arduino Zero board should restart and then begin running your program.


### Troubleshooting

If you have troubles getting your Arduino Zero board to receive code, try this:

- Press the "RESET" button on the board two times to get the Arduino Zero board ready to receive code.
- Now try running the `tinygo flash` command as above:

    ```shell
    tinygo flash -target=arduino-zero [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Arduino Zero board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You can use the USB port to the Arduino Zero as a serial port. `UART0` refers to this connection.
