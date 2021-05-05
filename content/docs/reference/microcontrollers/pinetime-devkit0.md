---
title: "PineTime"
weight: 3
---

The [PineTime](https://wiki.pine64.org/index.php/PineTime) is a smartwatch by [Pine64](https://www.pine64.org/) that is based on the Nordic Semiconductor [nRF52832](https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF52832) SoC. As of October 2019, a limited amount has been produced for developers.

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

[Documentation for the machine package for the PineTime](../machine/pinetime-devkit0)

## Flashing

The PineTime comes with some (proprietary?) default firmware and has the flash locked. This means that to flash this smartwatch, the chip must be unlocked (which erases all flash). There are two options for that, OpenOCD or nrfjprog.

Make sure to connect the SWD and power pins. For example, if you're using the [J-Link EDU Mini](https://www.segger.com/products/debug-probes/j-link/models/j-link-edu-mini/), connect the following pins:

| [PineTime pin](https://wiki.pine64.org/index.php/File:PineTime_SWD_location.jpg) | J-Link pin |
| ------------ | ---------- |
| SWDIO        | SWIO       |
| SWDCLK       | CLK        |
| 3.3V VCC     | Vref       |
| GND          | GND        |

If you get the following error:

    Error: Could not find MEM-AP to control the core

it means the chip is locked. See below for how to unlock and erase it.

### Unlock chip using nrfjprog

If you have `nrfjprog` installed, you can easily unlock the chip using the following command:

    nrfjprog -f nrf52 --recover

You can find how to install nrfjprog in the [documentation for the pca10040 board](../pca10040#flashing).

### Unlock chip using OpenOCD

Recent OpenOCD versions support unlocking/erasing the flash. Unfortunately, it isn't supported by OpenOCD 0.10, which is (as of October 2019) included in most Linux distributions. So you'll have to compile OpenOCD from source and use that. See [this StackOverflow answer](https://stackoverflow.com/questions/52308978/problem-flashing-nrf52-chip-using-openocd#54372481) for more details.

After unlocking, you can use a regular OpenOCD version (like 0.10) for flashing.

### Flash a new program

See above for how to connect the PineTime to your computer. When it is connected, you can flash programs to it. For example, to blink the LCD backlight:

```sh
tinygo flash -target=pinetime-devkit0 examples/blinky1
```
