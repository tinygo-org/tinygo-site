---
title: "Vicharak Shrike-Lite"
weight: 3
---

The [Vicharak Shrike-Lite](https://github.com/vicharak-in/shrike) is a tiny development board based on the Raspberry Pi [RP2040](https://datasheets.raspberrypi.org/rp2040/rp2040-datasheet.pdf) microcontroller paired with a [Renesas ForgeFPGA SLG47910V](https://www.renesas.com/en/products/slg47910) FPGA.

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

## Pins

| Pin               | Hardware pin | Alternative names     | I2C                  | PWM                  |
| ----------------- | ------------ | -----------------     | -------------------- | -------------------- |
| `IO0`             | `GPIO0`      | `F6`                  | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `IO1`             | `GPIO1`      | `F4`                  | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `IO2`             | `GPIO2`      | `F3`                  | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `IO3`             | `GPIO3`      | `F5`                  | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `IO4`             | `GPIO4`      | `LED`                 | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `IO5`             | `GPIO5`      |                       | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `IO6`             | `GPIO6`      | `I2C1_SDA_PIN`        | `I2C1` (SDA)         | `PWM3` (channel A)   |
| `IO7`             | `GPIO7`      | `I2C1_SCL_PIN`        | `I2C1` (SCL)         | `PWM3` (channel B)   |
| `IO8`             | `GPIO8`      | `SPI1_SDI_PIN`        | `I2C0` (SDA)         | `PWM4` (channel A)   |
| `IO9`             | `GPIO9`      | `UART1_RX_PIN`        | `I2C0` (SCL)         | `PWM4` (channel B)   |
| `IO10`            | `GPIO10`     | `SPI1_SCK_PIN`        | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `IO11`            | `GPIO11`     | `SPI1_SDO_PIN`        | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `IO12`            | `GPIO12`     | `FPGA_PWR`            | `I2C0` (SDA)         | `PWM6` (channel A)   |
| `IO13`            | `GPIO13`     | `FPGA_EN`             | `I2C0` (SCL)         | `PWM6` (channel B)   |
| `IO14`            | `GPIO14`     | `F18`                 | `I2C1` (SDA)         | `PWM7` (channel A)   |
| `IO15`            | `GPIO15`     | `F17`                 | `I2C1` (SCL)         | `PWM7` (channel B)   |
| `IO16`            | `GPIO16`     | `SPI0_SDI_PIN`        | `I2C0` (SDA)         | `PWM0` (channel A)   |
| `IO17`            | `GPIO17`     |                       | `I2C0` (SCL)         | `PWM0` (channel B)   |
| `IO18`            | `GPIO18`     | `SPI0_SCK_PIN`        | `I2C1` (SDA)         | `PWM1` (channel A)   |
| `IO19`            | `GPIO19`     | `SPI0_SDO_PIN`        | `I2C1` (SCL)         | `PWM1` (channel B)   |
| `IO20`            | `GPIO20`     | `SPI0_SDI_PIN`        | `I2C0` (SDA)         | `PWM2` (channel A)   |
| `IO21`            | `GPIO21`     |                       | `I2C0` (SCL)         | `PWM2` (channel B)   |
| `IO22`            | `GPIO22`     |                       |                      | `PWM3` (channel A)   |
| `IO23`            | `GPIO23`     |                       |                      | `PWM3` (channel B)   |
| `IO24`            | `GPIO24`     | `UART1_TX_PIN`        |                      | `PWM4` (channel A)   |
| `IO25`            | `GPIO25`     | `UART1_RX_PIN`        |                      | `PWM4` (channel B)   |
| `IO26`            | `GPIO26`     | `ADC0`                | `I2C1` (SDA)         | `PWM5` (channel A)   |
| `IO27`            | `GPIO27`     | `ADC1`                | `I2C1` (SCL)         | `PWM5` (channel B)   |
| `IO28`            | `GPIO28`     | `ADC2`, `UART0_TX_PIN`|                      | `PWM6` (channel A)   |
| `IO29`            | `GPIO29`     | `ADC3`, `UART0_RX_PIN`|                      | `PWM6` (channel B)   |

## Machine Package Docs

[Documentation for the machine package for Shrike-Lite](../machine/vicharak-shrike-lite)

## Flashing

### UF2

The board comes with the [UF2 bootloader](https://github.com/Microsoft/uf2) already installed.

### CLI Flashing

- Put your board into Bootloader mode
- Flash your TinyGo program to the board using this command:

    ```shell
    tinygo flash -target=vicharak_shrike_lite [PATH TO YOUR PROGRAM]
    ```

- The board should restart and then begin running your program.

### Flashing the FPGA
An SPI interface should be configured at specific pins in order to flash the FPGA
```go
	machine.SPI0.Configure(machine.SPIConfig{
		SDO:       machine.IO3, // MOSI
		SDI:       machine.IO0, // MISO
		SCK:       machine.IO2,
		Frequency: 1600000,
		Mode:      0b00,
	})
	machine.FPGA_PWR.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.FPGA_EN.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.IO1.Configure(machine.PinConfig{Mode: machine.PinOutput}) // SS
```

To start flashing, set `FPGA_PWR` and `FPGA_EN` to `HIGH`, `IO1` to `LOW`, and write the bitstream. \
Once flashing is successful, set `IO1` to `HIGH`.
```go
	err := machine.SPI0.Tx(data, nil) // data[] is the bitstream

	if err != nil {
		panic(err)
	}
	machine.IO1.High() // Setting SS to HIGH
```

A simple tool for flashing the FPGA can be found at [https://github.com/YajTPG/shrike_go/blob/master/shrikeflash.go](https://github.com/YajTPG/shrike_go)


### Troubleshooting

Any troubleshooting tips go here.

## Notes

You can use the USB port to the board as a serial port.

TinyGo has support for the RP2040's on-board Programmable Input/Output (PIO) block.

For more informantion, see [https://github.com/tinygo-org/pio](https://github.com/tinygo-org/pio)
