---
title: "PJRC Teensy 3.6"
weight: 3
---

The [PJRC Teensy 3.6](https://www.pjrc.com/store/teensy36.html) is a small ARM development board based on the NXP [MK66FX1M0VMD18](https://www.nxp.com/docs/en/data-sheet/K66P144M180SF5V2.pdf) 32-bit 180 MHz ARM Cortex-M4.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI       | YES | Not yet |
| I2C       | YES | Not yet |
| ADC       | YES | Not yet |
| PWM       | YES | Not yet |
| USBDevice | YES | Not yet |

## Pins

| Pin               | Hardware pin | Alternative names |
| ----------------- | ------------ | ----------------- |
| `D00`             | `PB16`       |                   |
| `D01`             | `PB17`       |                   |
| `D02`             | `PD00`       |                   |
| `D03`             | `PA12`       |                   |
| `D04`             | `PA13`       |                   |
| `D05`             | `PD07`       |                   |
| `D06`             | `PD04`       |                   |
| `D07`             | `PD02`       |                   |
| `D08`             | `PD03`       |                   |
| `D09`             | `PC03`       |                   |
| `D10`             | `PC04`       |                   |
| `D11`             | `PC06`       |                   |
| `D12`             | `PC07`       |                   |
| `D13`             | `PC05`       | `LED`             |
| `D14`             | `PD01`       |                   |
| `D15`             | `PC00`       |                   |
| `D16`             | `PB00`       |                   |
| `D17`             | `PB01`       |                   |
| `D18`             | `PB03`       |                   |
| `D19`             | `PB02`       |                   |
| `D20`             | `PD05`       |                   |
| `D21`             | `PD06`       |                   |
| `D22`             | `PC01`       |                   |
| `D23`             | `PC02`       |                   |
| `D24`             | `PE26`       |                   |
| `D25`             | `PA05`       |                   |
| `D26`             | `PA14`       |                   |
| `D27`             | `PA15`       |                   |
| `D28`             | `PA16`       |                   |
| `D29`             | `PB18`       |                   |
| `D30`             | `PB19`       |                   |
| `D31`             | `PB10`       |                   |
| `D32`             | `PB11`       |                   |
| `D33`             | `PE24`       |                   |
| `D34`             | `PE25`       |                   |
| `D35`             | `PC08`       |                   |
| `D36`             | `PC09`       |                   |
| `D37`             | `PC10`       |                   |
| `D38`             | `PC11`       |                   |
| `D39`             | `PA17`       |                   |
| `D40`             | `PA28`       |                   |
| `D41`             | `PA29`       |                   |
| `D42`             | `PA26`       |                   |
| `D43`             | `PB20`       |                   |
| `D44`             | `PB22`       |                   |
| `D45`             | `PB23`       |                   |
| `D46`             | `PB21`       |                   |
| `D47`             | `PD08`       |                   |
| `D48`             | `PD09`       |                   |
| `D49`             | `PB04`       |                   |
| `D50`             | `PB05`       |                   |
| `D51`             | `PD14`       |                   |
| `D52`             | `PD13`       |                   |
| `D53`             | `PD12`       |                   |
| `D54`             | `PD15`       |                   |
| `D55`             | `PD11`       |                   |
| `D56`             | `PE10`       |                   |
| `D57`             | `PE11`       |                   |
| `D58`             | `PE00`       |                   |
| `D59`             | `PE01`       |                   |
| `D60`             | `PE02`       |                   |
| `D61`             | `PE03`       |                   |
| `D62`             | `PE04`       |                   |
| `D63`             | `PE05`       |                   |

## Machine Package Docs

[Documentation for the machine package for the Teensy 3.6](../machine/teensy36)

## Flashing

### teensy_loader_cli

In order to flash your TinyGo programs onto the Teensy 3.6 board, you will need to install the `teensy_loader_cli` (https://github.com/PaulStoffregen/teensy_loader_cli) on your machine.

### CLI Flashing

- Plug your Teensy 3.6 into your computer's USB port.
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=teensy36 [PATH TO YOUR PROGRAM]
    ```

- The Teensy 3.6 board should restart and then begin running your program.

### Troubleshooting

If you have troubles getting your Teensy 3.6 board to receive code, try this:

- Press the "RESET" button on the board two times to get the Teensy 3.6 board ready to receive code.
- The Teensy 3.6 board will appear to your computer like a USB drive.
- Now try running the command as above:

    ```shell
    tinygo flash -target=teensy36 [PATH TO YOUR PROGRAM]
    ```

Once you have updated your Teensy 3.6 board the first time, after that you should be able to flash it entirely from the command line.

## Notes

You cannot yet use the USB port to the Teensy 3.6 as a USB CDC serial port.
