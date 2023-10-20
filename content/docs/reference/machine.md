---
title: "machine package"
weight: 5
description: >
  How to interact with peripherals exposed through the `machine` package.
---

## GPIO


```go
type Pin uint8

const NoPin = Pin(0xff)
```

The pin type indicates a physical pin on a microcontroller. It can be used directly for GPIO or as part of a different peripheral (SPI, ADC, etc).

The `NoPin` value can be used in some places where you want to explicitly not use a pin. For example, depending on the chip you could use only SDO or SDI of an SPI peripheral, leaving the other pin unused.

Boards have various pins defined as constants. For example, ATsamd21 based chips (often marketed as "M0") have pin number PB01 defined as `PB01` in the machine package and the Arduino Uno has pin 13 defined as `D13`. However, please note that Arduino pin numbers do not necessarily match the pin numbers as used in TinyGo, therefore you should always use the named constants to refer to a pin if possible.

```go
type PinMode uint8

const (
    PinOutput          PinMode = ...
    PinInput           PinMode = ...
    PinInputPullup     PinMode = ...
    PinInputPulldown   PinMode = ... // not always available
    PinOutputOpenDrain PinMode = ... // not always available
    PinDisable         PinMode = ... // not always available
)
```

These `PinMode` are used to set a pin to various states: as input or output and optionally with a pull resistor (see above in the guide). The values these constants have vary by chip and should be considered implementation details.

Pull-down mode is not always available on the hardware. If it is not available, the `PinInputPulldown` constant is not available resulting in a compile error.

```go
type PinConfig struct {
    Mode PinMode
}
```

This struct contains options to configure a hardware pin. More options might be added in the future (with appropriate behavior on the zero value), but currently only the `Mode` field is used.

```go
func (p Pin) Configure(config machine.PinConfig)
```

Configure a pin, see `PinConfig` and `PinMode` above. `Configure` must be called before any other method may be called.

```go
func (p Pin) Low()
func (p Pin) High()
func (p Pin) Set(state bool)
```

Set an output pin to low or high. It must have been configured before as an output (`PinOutput`) or the behavior will be undefined. The `Set` method sets the state to low or high depending on the boolean value: `true` means high, `false` means low.

```go
func (p Pin) Get() bool
```

Get the input state of a pin. The returned value indicates whether the pin is low or high: `true` means high and `false` means low. The pin must be configured as an input or as an output.

Note that if the pin is left floating (not connected to anything) the returned value is unpredictable and may appear random.

```go
type PinChange uint8

const (
    PinRising  PinChange = ... // call callback when the pin goes from low to high
    PinFalling PinChange = ... // call callback when the pin goes from high to low
    PinToggle  PinChange = ... // call callback when the pin changes either way
)

func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) error
```

SetInterrupt sets an interrupt to be executed when the pin changes state. The pin should already be configured as an input, including a pull up or down if no external pull is provided. The callback is called in an interrupt handler, which means the code is limited in what it can do: it cannot block and it may not allocate heap memory.

This call will replace a previously set callback on this pin. You can pass `nil` to disable the interrupt. If you do so, the change parameter is ignored and can be set to any value (such as 0).


## SPI

```go
type SPIConfig struct {
    Frequency uint32
    SCK       Pin
    SDO       Pin
    SDI       Pin
    LSBFirst  bool
    Mode      uint8
}
```

The `SPIConfig` struct contains the configuration for the SPI peripheral.

  * `Frequency` is the maximum frequency that will be used: for example `1 * MHz` for 1MHz. Depending on chip capabilities, this or a lower frequency will be selected. When not set (or set to 0), the default of 4MHz will be used.
  * `SCK`, `SDO` and `SDI` are the clock, data out, and data in pins respectively, however support for setting pins other than the default pins may not be supported by a given SPI peripheral. Some chips are flexible and allow the use of any pin, while other boards only allow a limited range of pins or use fixed SCK/SDO/SDI pins. When these pins are left at the zero value, the default for the particular board is used.
  * `LSBFirst` configures the SPI peripheral to clock out the least significant bit (LSB) first. The default and most commonly used configuration is the most significant bit first (`LSBFirst=false`).
  * `Mode` is the [SPI mode (CPOL/CPHA) to be used](https://en.wikipedia.org/wiki/Serial_Peripheral_Interface#Clock_polarity_and_phase). Mode 0 is appropriate for most peripheral chips, but other modes may be needed in some cases. Check your peripheral documentation for details.
      | Mode    | CPOL | CPHA |
      | ------- | ---- | ---- |
      | `Mode0` | 0    | 0    |
      | `Mode1` | 0    | 1    |
      | `Mode2` | 1    | 0    |
      | `Mode3` | 1    | 1    |

```go
type SPI struct {
    // values are unexported or vary by chip
}

var (
    SPI0 = SPI{...}
    SPI1 = SPI{...}
)
```

The `SPI` object refers to a single (hardware) SPI instance. Depending on chip capabilities, various objects such as `SPI0` and perhaps others are defined.

```go
func (spi SPI) Configure(config SPIConfig) error
```

The `Configure` call enables and configures the hardware SPI for use, setting the configuration as described in `SPIConfig`. It will return an error when an incorrect configuration is provided (for example, using pins not usable with this SPI instance). See `SPIConfig` for details.

```go
func (spi SPI) Transfer(b byte) (byte, error)
```

Transmit and receive a single byte. Due to the nature of the SPI protocol, they happen both at the same time.

Use `Tx` instead of `Transfer` for high performance bulk transfers.

```go
func (spi SPI) Tx(w, r []byte) error
```

The `Tx` performs the actual SPI transaction, and return an error if there was an error during the transaction. Because SPI is a synchronous protocol, reading and writing happens at the same time. Therefore, `w` and `r` usually have to be the same length. There are some exceptions:

  * When `r` is nil, the SPI peripheral will only transmit the bytes in `w` and ignore what it receives.
  * When `w` is nil, the SPI peripheral will only read the given number of bytes and store it in `r`. It will send out a zero byte for each byte that it reads.

Some chips may also support mismatched lengths of `w` and `r`, in which case they will behave like above for the remaining bytes in the byte slice that's the longest of the two.


## I2C

```go
type I2CConfig struct {
    Frequency uint32
    SCL       Pin
    SDA       Pin
    Mode      I2CMode
}
```

The `I2CConfig` struct contains the configuration for the I2C peripheral.

  * `Frequency` can be set to either 100kHz (`100e3`), 400kHz (`400e3`), and sometimes to other values depending on the chip. The zero value defaults to 100kHz.
  * `SCL` and `SDA` can be set as desired, however support for different pins than the default is limited. Some chips are flexible and allow the use of any pin, while other boards only allow a limited range of pins or use fixed SCL/SDA pins. When both pins are left at the zero value, the default for the particular board is used.
  * `Mode` is present on peripherals that support I2C target mode.  The default
  is I2C controller mode, setting `Mode` to `I2CModeTarget` will configure the
  peripheral as an I2C target.
```go
type I2C struct {
    // values are unexported or vary by chip
}

var (
    I2C0 = I2C{...}
    I2C1 = I2C{...}
)
```

The `I2C` object refers to a single (hardware) I2C instance. Depending on chip capabilities, various objects such as `I2C0` and perhaps others are defined.

```go
func (i2c I2C) Configure(config I2CConfig) error
```

The `Configure` call enables and configures the hardware I2C for use, setting the pins and frequency. It will return an error when an incorrect configuration is provided (for example, using pins not usable with this I2C instance). See `I2CConfig` for details.

```go
func (i2c I2C) Tx(addr uint16, w, r []byte) error
```

_I2C Controller Mode Only_: The `Tx` call performs the actual I2C transaction. It first writes the bytes in `w` to the peripheral device indicated in `addr` and then reads `r` bytes from the peripheral and stores the read bytes in the `r` slice. It returns an error if the transaction failed. Both `w` and `r` can be `nil`.

```go
func (i2c I2C) Listen(addr uint16) error
```

_I2C Target Mode Only_: The `Listen` call starts the I2C peripheral listening for I2C transactions sent
to `addr` by the controller.  The peripheral must have been configured in target
mode (see `I2CConfig` struct) before `Listen` is called.

```go
func (i2c *I2C) WaitForEvent(buf []byte) (evt I2CTargetEvent, count int, err error)
```

_I2C Target Mode Only_: The `WaitForEvent` call blocks the current goroutine waiting for an I2C event.  For `I2CReceive` events, the message will be placed in `buf` and return the `count` of bytes received.  Oversize messages (those larger than `buf`) will be truncated.

The underlying peripheral will perform clock stretching, if necessary, in two cases:

  1. A correctly addressed message is received and the application is not blocked on a call to `WaitForEvent`,

  2. The application does not call `Reply` with a single I2C clock cycle for `I2CRequest` events.

Although the I2C target may perform clock stretching, controllers may implement arbitrary timeouts for pending devices.  To avoid timeouts from the perspective a controller, the application should:

  1. Handle the returned event in a timely manner, calling `Reply` if appropriate.

  2. Not have any go routines that may block indefinitely as `WaitForEvent` may yield the CPU to another go routine while waiting for an event.

```go
func (i2c I2C) Reply(buf []byte) error
```

_I2C Target Mode Only_: The `Reply` call sends a response to the controller when an `I2CRequest` event is received by the target.

## UART

```go
type UARTConfig struct {
    BaudRate uint32
    TX       Pin
    RX       Pin
}
```

The `UARTConfig` struct contains configuration for the UART peripheral.

  * `BaudRate` is the baud rate of the UART. Common values are 9600 and 115200. Other than that, most chips support multiples of 1200 (2400, 4800, 9600, etc). Support for other baud rates varies by chip, some chips even support high baudrates like 1MHz.
  * `TX` and `RX` are the transmit and receive pins of the UART. Many chips impose restrictions on which pins can be used, some only support a particular TX and RX pin.

```go
type UARTParity uint8

const (
    ParityNone UARTParity = iota
    ParityEven
    ParityOdd
)
```

Parity is a rarely used checksum feature of UART.

  * `ParityNone` is the default, meaning no parity. It is the most commonly used setting.
  * `ParityEven` means to expect that the total number of 1 bits sent should be an even number.
  * `ParityOdd` means to expect that the total number of 1 bits sent should be an odd number.

```go
type UART struct {
    // values are unexported or vary by chip
}

var (
    UART0 = &UART{...}
    UART1 = &UART{...}
)
```

The `UART` object refers to a single (hardware) UART instance. Depending on chip capabilities, various objects such as `UART0` and perhaps others are defined.

```go
func (uart UART) Configure(config UARTConfig) error
```

The `Configure` call enables and configures the hardware UART for use, setting the pins and baud rate. It will return an error when an incorrect configuration is provided (for example, using pins not usable with this UART instance). See `UARTConfig` for details.

Depending on the target configuration, a UART may already be configured if it is the stdout output for the given target. In that case, it is normally configured with a baud rate of 115200.

```go
func (uart *UART) SetBaudRate(br uint32)
```

Set the baud rate for the UART. This method is not available on all chips. See `UARTConfig` above for permissible baud rate values.

```go
func (uart *UART) SetFormat(dataBits, stopBits int, parity UARTParity) error
```

Set the UART format. The default format (8N1 meaning 8 bits, no parity, and 1 stop bit) is used for almost all external devices, but if you need it this method can be used to override the default.

This method is only available on a limited number of chips.

```go
func (uart *UART) Buffered() int
```

Return the number of bytes stored in the receive buffer.

```go
func (uart *UART) Read(data []byte) (n int, err error)
```

Read from the receive buffer. This method implements the `io.Reader` interface. It is non-blocking: it will return immediately with `n` set to 0 and `err` set to nil if no data is available.

```go
func (uart *UART) ReadByte() (byte, error)
```

ReadByte reads a single byte from the UART receive buffer. If there is no data available in the buffer, it returns an error. You can use `Buffered` to know whether there is data available.

```go
func (uart *UART) Write(data []byte) (n int, err error)
```

Write data to the UART. This method implements the `io.Writer` interface. It blocks until all data is written.

```go
func (uart *UART) WriteByte(c byte) error
```

Write a single byte to the UART output.

## Watchdog

```go
type WatchdogConfig struct {
	TimeoutMillis uint32
}
```

The `WatchdogConfig` struct contains configuration for the watchdog peripheral.

  * `TimeoutMillis` is the requested maximum delay between calls to `Update`.

```go
func (wd *Watchdog) Configure(config WatchdogConfig) error
```
Configures the watchdog.

This method should not be called after the watchdog is started and on some platforms attempting to reconfigure after starting the watchdog is explicitly forbidden / will not work.

```go
func (wd *Watchdog) Start() error
```
Starts the watchdog.  After calling this method, `Update` must be called periodically.

```go
func (wd *Watchdog) Update()
```
Updates the watchdog, indicating that the app is healthy.  This method must be called periodically to prevent the CPU from resetting.

## Other

```go
func CPUFrequency() uint32
```

Return the current CPU frequency in hertz (for example, 16MHz equals 16_000_000). It is often a fixed value.

```go
func CPUReset()
```

CPUReset performs a hard system reset.

Not all chips support CPUReset.

```go
func DeviceID() []byte
```

DeviceID returns a byte array containing a unique id (aka Serial Number) specific to this chip.  In some architectures (notably RP2040) the device ID is actually the ID of the flash chip.  The device ID can be useful for identifying specific devices within a family.  There is no guarantee the ID is globally unique.  The size of the ID is chip-family specific with 8 bytes (64 bits) and 16 bytes (128 bits) being common.

Not all chips have a hardware ID.

```go
func GetRNG() uint32
```

Return a 32-bit random number from a hardware random number generator. It is often (but not always) a cryptographic random number generator. Check the documentation of the chip to be sure.

Not all chips have a random number generator.

```go
func ReadTemperature() int32
```

Read the current die temperature of the chip. The return value is in milli-celsius: to convert to Celsius, divide the returned value by 1000.

Not all chips have a built-in temperature sensor.
