---
title: "BBC:Microbit"
weight: 3
---

The [BBC micro:bit](https://microbit.org) is a tiny programmable computer designed for learning. It is based on the Nordic Semiconductor [nRF51822](https://www.nordicsemi.com/eng/Products/Bluetooth-low-energy/nRF51822) ARM Cortex MO chip.

## Interfaces

| Interface | Hardware Supported | TinyGo Support |
| --------- | ------------- | ----- |
| GPIO      | YES | YES |
| UART      | YES | YES |
| SPI      | YES | YES |
| I2C      | YES | YES |
| ADC      | YES | Not yet |
| PWM      | Software support | Not yet |

## Flashing

### File Copy

The simplest way to load programs onto the micro:bit is to copy the hex file onto the device. You can do this either using drag-and-drop with your operating system's file explorer.

- Plug your micro:bit into one of your computer's USB ports
- Often this will cause the machine's graphical explorer utility to run listing a device "MICROBIT"
- You can drag-and-drop hex files onto the "MICROBIT" device to flash the device
- Alternatively (see below), to use the command-line, you will need to determine the path to the board then copy the file.
 
#### Linux
- Determine the path to the board.
- It will be something like `/media/$(whoami)/MICROBIT`.
- You can confirm the device by listing its contents (`ls /media/$(whoami)/MICROBIT`).
- The directory should contain two files: `DETAILS.TXT` and `MICROBIT.HTM`
- Copy your hex file to the device using `cp /path/to/your/code.hex /media/$(whoami)/MICROBIT`.
 
#### Windows
- Determine the path to the board.
- It will be mounted to one of the machine's available drive letters, perhaps `D:`.
- You can confirm the device is mounted to the drive by listing its contents (`dir D:`).
- The directory should contain two files: `DETAILS.TXT` and `MICROBIT.HTM`
- Copy your hex file to the device using `copy \path\to\your\code.hex D:`.

You can simply right-click on a hex file, select "Send To" and "MICROBIT (X:)".

#### MacOS
- Determine the path to the board.
- It will be something like `/Volumes/"MICROBIT"`.
- You can confirm the device by listings its contents (`ls /Volumes/"MICROBIT"`)
- The directory should contain two files: `DETAILS.TXT` and `MICROBIT.HTM`
- Copy your hex file to the device using `cp /path/to/your/code.hex /Volumes/"MICROBIT"`.

### OpenOCD

An alternative approach to load programs onto the micro:bit is by using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the micro:bit board with your TinyGo code.

- Plug your Microbit into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=microbit`

## Notes

The micro:bit has two built-in I2C devices, a MMA8653 accelerometer and a MAG3110 magnetometer. You can use them via the I2C0 bus.

The micro:bit I2C0 and SPI0 buses both share the same address space. This means you cannot use them both at the same time. However, you can still use both SPI and I2C at the same time, by using the SPI1 bus with the standard SPI pins at the same time as using the I2C0 bus to access the built-in devices.

For example:

```go
// use the same pins as for SPI0, but with SPI1
machine.SPI1.Configure(machine.SPIConfig{
    SCK:       machine.SPI0_SCK_PIN,
    MOSI:      machine.SPI0_MOSI_PIN,
    MISO:      machine.SPI0_MISO_PIN,
    Frequency: 500000,
    Mode:      0})

// use I2C0 as normally do
machine.I2C0.Configure(machine.I2CConfig{})
```
