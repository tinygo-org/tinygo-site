---
title: "Nintendo Switch"
weight: 3
---

The [Nintendo Switch](https://en.wikipedia.org/wiki/Nintendo_Switch) is a handheld videogame platform based on the [Nvidia Tegra X1](https://en.wikipedia.org/wiki/Tegra#Tegra_X1) SoC.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | ? | ? |
| UART      | ? | ? |
| SPI       | ? | ? |
| I2C       | ? | ? |
| ADC       | ? | ? |
| PWM       | ? | ? |
| USBDevice | ? | ? |

## Machine Package Docs

[Documentation for the machine package for the Nintendo Switch](../machine/nintendoswitch)

## Installing dependencies

You will need the `linkle` (https://github.com/MegatonHammer/linkle) program to convert to the NRO format needed by the Switch:

You can use a Nintendo Switch software emulator such as [suyu https://suyu.dev/](https://suyu.dev/) to test your programs.

## Building code

Build your Nintendo Switch programs using `-target nintendoswitch` like this:

```shell
tinygo build -o main.elf -target nintendoswitch examples/serial
```

Once you have created the ELF file, convert it into the NRO format using linkle:

```shell
linkle nro main.elf main.nro
```

You can now use the NRO file with your emulator or flash it onto your physical hardware.

## Flashing

Information needed here...

## Notes

See the gonx package (https://github.com/racerxdl/gonx) for wrappers around Nintendo Switch APIs.

Examples using gonx can be found at https://github.com/racerxdl/go-switch-examples
