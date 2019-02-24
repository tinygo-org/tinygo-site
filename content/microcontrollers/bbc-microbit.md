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

### OpenOCD

Programs are loaded onto the BBC:Microbit using the `openocd` command line utility program. You must install [OpenOCD](http://openocd.org/) before you will be able to flash the BBC:Microbit board with your TinyGo code.

- Plug your Microbit into your computer's USB port.
- Build and flash your TinyGo program using `tinygo flash -target=microbit`

## Notes

The BBC:Microbit has two built-in I2C devices, a MMA8653 accelerometer and a MAG3110 magnetometer. You can use them via the I2C0 bus.

The BBC:Microbit I2C0 and SPI0 buses both share the same address space. This means you cannot use them both at the same time. However, you can still use both SPI and I2C at the same time, by using the SPI1 bus with the standard SPI pins at the same time as using the I2C0 bus to access the built-in devices.

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
