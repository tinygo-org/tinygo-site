
---
title: nano-rp2040
---


## Constants

```go
const (
	D2	Pin	= GPIO25
	D3	Pin	= GPIO15
	D4	Pin	= GPIO16
	D5	Pin	= GPIO17
	D6	Pin	= GPIO18
	D7	Pin	= GPIO19
	D8	Pin	= GPIO20
	D9	Pin	= GPIO21
	D10	Pin	= GPIO5
	D11	Pin	= GPIO7
	D12	Pin	= GPIO4
	D13	Pin	= GPIO6
	D14	Pin	= GPIO26
	D15	Pin	= GPIO27
	D16	Pin	= GPIO28
	D17	Pin	= GPIO29
	D18	Pin	= GPIO12
	D19	Pin	= GPIO13
)
```

Digital Pins


```go
const (
	A0	Pin	= ADC0
	A1	Pin	= ADC1
	A2	Pin	= ADC2
	A3	Pin	= ADC3
)
```

Analog pins


```go
const (
	LED = GPIO6
)
```

Onboard LED


```go
const (
	I2C0_SDA_PIN	Pin	= GPIO12
	I2C0_SCL_PIN	Pin	= GPIO13

	I2C1_SDA_PIN	Pin	= GPIO18
	I2C1_SCL_PIN	Pin	= GPIO19
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	Pin	= GPIO6
	SPI0_SDO_PIN	Pin	= GPIO7
	SPI0_SDI_PIN	Pin	= GPIO4

	// GPIO22 does not have SPI functionality so we set it to avoid interfering with NINA.
	SPI1_SCK_PIN	Pin	= GPIO22
	SPI1_SDO_PIN	Pin	= GPIO22
	SPI1_SDI_PIN	Pin	= GPIO22
)
```

SPI pins. SPI1 not available on Nano RP2040 Connect.


```go
const (
	NINA_SCK	Pin	= GPIO14
	NINA_SDO	Pin	= GPIO11
	NINA_SDI	Pin	= GPIO8

	NINA_CS		Pin	= GPIO9
	NINA_ACK	Pin	= GPIO10
	NINA_GPIO0	Pin	= GPIO2
	NINA_RESETN	Pin	= GPIO3

	NINA_TX	Pin	= GPIO9
	NINA_RX	Pin	= GPIO8
)
```

NINA-W102 Pins


```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.


```go
const Device = deviceName
```

Device is the running program's chip name, such as "ATSAMD51J19A" or
"nrf52840". It is not the same as the CPU name.

The constant is some hardcoded default value if the program does not target a
particular chip but instead runs in WebAssembly for example.


```go
const NoPin = Pin(0xff)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	// GPIO pins
	GPIO0	Pin	= 0
	GPIO1	Pin	= 1
	GPIO2	Pin	= 2
	GPIO3	Pin	= 3
	GPIO4	Pin	= 4
	GPIO5	Pin	= 5
	GPIO6	Pin	= 6
	GPIO7	Pin	= 7
	GPIO8	Pin	= 8
	GPIO9	Pin	= 9
	GPIO10	Pin	= 10
	GPIO11	Pin	= 11
	GPIO12	Pin	= 12
	GPIO13	Pin	= 13
	GPIO14	Pin	= 14
	GPIO15	Pin	= 15
	GPIO16	Pin	= 16
	GPIO17	Pin	= 17
	GPIO18	Pin	= 18
	GPIO19	Pin	= 19
	GPIO20	Pin	= 20
	GPIO21	Pin	= 21
	GPIO22	Pin	= 22
	GPIO23	Pin	= 23
	GPIO24	Pin	= 24
	GPIO25	Pin	= 25
	GPIO26	Pin	= 26
	GPIO27	Pin	= 27
	GPIO28	Pin	= 28
	GPIO29	Pin	= 29

	// Analog pins
	ADC0	Pin	= GPIO26
	ADC1	Pin	= GPIO27
	ADC2	Pin	= GPIO28
	ADC3	Pin	= GPIO29
)
```



```go
const (
	UART_TX_PIN	= UART0_TX_PIN
	UART_RX_PIN	= UART0_RX_PIN
	UART0_TX_PIN	= GPIO0
	UART0_RX_PIN	= GPIO1
	UART1_TX_PIN	= GPIO8
	UART1_RX_PIN	= GPIO9
)
```

UART pins


```go
const (
	KHz	= 1000
	MHz	= 1000000
)
```



```go
const (
	PinOutput	PinMode	= iota
	PinInput
	PinInputPulldown
	PinInputPullup
	PinAnalog
	PinUART
	PinPWM
	PinI2C
	PinSPI
)
```



```go
const (
	// PinLevelLow triggers whenever pin is at a low (around 0V) logic level.
	PinLevelLow	PinChange	= 1 << iota
	// PinLevelLow triggers whenever pin is at a high (around 3V) logic level.
	PinLevelHigh
	// Edge falling
	PinFalling
	// Edge rising
	PinRising
)
```

Pin change interrupt constants for SetInterrupt.


```go
const RESETS_RESET_Msk = 0x01ffffff
```

RESETS_RESET_Msk is bitmask to reset all peripherals

TODO: This field is not available in the device file.


```go
const (
	// ParityNone means to not use any parity checking. This is
	// the most common setting.
	ParityNone	UARTParity	= 0

	// ParityEven means to expect that the total number of 1 bits sent
	// should be an even number.
	ParityEven	UARTParity	= 1

	// ParityOdd means to expect that the total number of 1 bits sent
	// should be an odd number.
	ParityOdd	UARTParity	= 2
)
```






## Variables

```go
var (
	NINA_SPI = SPI1
)
```



```go
var (
	ErrTimeoutRNG		= errors.New("machine: RNG Timeout")
	ErrInvalidInputPin	= errors.New("machine: invalid input pin")
	ErrInvalidOutputPin	= errors.New("machine: invalid output pin")
	ErrInvalidClockPin	= errors.New("machine: invalid clock pin")
	ErrInvalidDataPin	= errors.New("machine: invalid data pin")
	ErrNoPinChangeChannel	= errors.New("machine: no channel available for pin interrupt")
)
```



```go
var (
	UART0	= &_UART0
	_UART0	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	rp.UART0,
	}

	UART1	= &_UART1
	_UART1	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	rp.UART1,
	}
)
```

UART on the RP2040


```go
var DefaultUART = UART0
```



```go
var (
	I2C0	= &_I2C0
	_I2C0	= I2C{
		Bus: rp.I2C0,
	}
	I2C1	= &_I2C1
	_I2C1	= I2C{
		Bus: rp.I2C1,
	}
)
```

I2C on the RP2040.


```go
var (
	ErrInvalidI2CBaudrate	= errors.New("invalid i2c baudrate")
	ErrInvalidTgtAddr	= errors.New("invalid target i2c address not in 0..0x80 or is reserved")
	ErrI2CGeneric		= errors.New("i2c error")
	ErrRP2040I2CDisable	= errors.New("i2c rp2040 peripheral timeout in disable")
)
```



```go
var (
	ErrBadPeriod = errors.New("period outside valid range 8ns..268ms")
)
```



```go
var (
	PWM0	= getPWMGroup(0)
	PWM1	= getPWMGroup(1)
	PWM2	= getPWMGroup(2)
	PWM3	= getPWMGroup(3)
	PWM4	= getPWMGroup(4)
	PWM5	= getPWMGroup(5)
	PWM6	= getPWMGroup(6)
	PWM7	= getPWMGroup(7)
)
```

Hardware Pulse Width Modulation (PWM) API
PWM peripherals available on RP2040. Each peripheral has 2 pins available for
a total of 16 available PWM outputs. Some pins may not be available on some boards.

The RP2040 PWM block has 8 identical slices. Each slice can drive two PWM output signals, or
measure the frequency or duty cycle of an input signal. This gives a total of up to 16 controllable
PWM outputs. All 30 GPIOs can be driven by the PWM block

The PWM hardware functions by continuously comparing the input value to a free-running counter. This produces a
toggling output where the amount of time spent at the high output level is proportional to the input value. The fraction of
time spent at the high signal level is known as the duty cycle of the signal.

The default behaviour of a PWM slice is to count upward until the wrap value (\ref pwm_config_set_wrap) is reached, and then
immediately wrap to 0. PWM slices also offer a phase-correct mode, where the counter starts to count downward after
reaching TOP, until it reaches 0 again.


```go
var (
	SPI0	= &_SPI0
	_SPI0	= SPI{
		Bus: rp.SPI0,
	}
	SPI1	= &_SPI1
	_SPI1	= SPI{
		Bus: rp.SPI1,
	}
)
```

SPI on the RP2040


```go
var (
	ErrLSBNotSupported	= errors.New("SPI LSB unsupported on PL022")
	ErrTxInvalidSliceSize	= errors.New("SPI write and read slices must be same size")
	ErrSPITimeout		= errors.New("SPI timeout")
	ErrSPIBaud		= errors.New("SPI baud too low or above 66.5Mhz")
)
```



```go
var (
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = DefaultUART
```

Serial is implemented via the default (usually the first) UART on the chip.





### func CPUFrequency

```go
func CPUFrequency() uint32
```



### func CurrentCore

```go
func CurrentCore() int
```

CurrentCore returns the core number the call was made from.


### func InitADC

```go
func InitADC()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


### func NumCores

```go
func NumCores() int
```

NumCores returns number of cores available on the device.


### func PWMPeripheral

```go
func PWMPeripheral(pin Pin) (sliceNum uint8, err error)
```

Peripheral returns the RP2040 PWM peripheral which ranges from 0 to 7. Each
PWM peripheral has 2 channels, A and B which correspond to 0 and 1 in the program.
This number corresponds to the package's PWM0 throughout PWM7 handles




## type ADC

```go
type ADC struct {
	Pin Pin
}
```




### func (ADC) Configure

```go
func (a ADC) Configure(config ADCConfig)
```

Configure configures a ADC pin to be able to be used to read data.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```





## type ADCConfig

```go
type ADCConfig struct {
	Reference	uint32	// analog reference voltage (AREF) in millivolts
	Resolution	uint32	// number of bits for a single conversion (e.g., 8, 10, 12)
	Samples		uint32	// number of samples for a single conversion (e.g., 4, 8, 16, 32)
}
```

ADCConfig holds ADC configuration parameters. If left unspecified, the zero
value of each parameter will use the peripheral's default settings.





## type I2C

```go
type I2C struct {
	Bus		*rp.I2C0_Type
	restartOnNext	bool
}
```




### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```

Configure initializes i2c peripheral and configures I2C config's pins passed.
Here's a list of valid SDA and SCL GPIO pins on bus I2C0 of the rp2040:
 SDA: 0, 4, 8, 12, 16, 20
 SCL: 1, 5, 9, 13, 17, 21
Same as above for I2C1 bus:
 SDA: 2, 6, 10, 14, 18, 26
 SCL: 3, 7, 11, 15, 19, 27


### func (*I2C) ReadRegister

```go
func (i2c *I2C) ReadRegister(address uint8, register uint8, data []byte) error
```

ReadRegister transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (*I2C) SetBaudRate

```go
func (i2c *I2C) SetBaudRate(br uint32) error
```

SetBaudRate sets the I2C frequency. It has the side effect of also
enabling the I2C hardware if disabled beforehand.


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
```

Tx performs a write and then a read transfer placing the result in
in r.

Passing a nil value for w or r skips the transfer corresponding to write
or read, respectively.

 i2c.Tx(addr, nil, r)
Performs only a read transfer.

 i2c.Tx(addr, w, nil)
Performs only a write transfer.


### func (*I2C) WriteRegister

```go
func (i2c *I2C) WriteRegister(address uint8, register uint8, data []byte) error
```

WriteRegister transmits first the register and then the data to the
peripheral device.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily write to such registers. Also, it only works for
devices with 7-bit addresses, which is the vast majority.




## type I2CConfig

```go
type I2CConfig struct {
	Frequency	uint32
	// SDA/SCL Serial Data and clock pins. Refer to datasheet to see
	// which pins match the desired bus.
	SDA, SCL	Pin
}
```

I2CConfig is used to store config info for I2C.





## type NullSerial

```go
type NullSerial struct {
}
```

NullSerial is a serial version of /dev/null (or null router): it drops
everything that is written to it.



### func (NullSerial) Buffered

```go
func (ns NullSerial) Buffered() int
```

Buffered returns how many bytes are buffered in the UART. It always returns 0
as there are no bytes to read.


### func (NullSerial) Configure

```go
func (ns NullSerial) Configure(config UARTConfig) error
```

Configure does nothing: the null serial has no configuration.


### func (NullSerial) ReadByte

```go
func (ns NullSerial) ReadByte() (byte, error)
```

ReadByte always returns an error because there aren't any bytes to read.


### func (NullSerial) Write

```go
func (ns NullSerial) Write(p []byte) (n int, err error)
```

Write is a no-op: none of the data is being written and it will not return an
error.


### func (NullSerial) WriteByte

```go
func (ns NullSerial) WriteByte(b byte) error
```

WriteByte is a no-op: the null serial doesn't write bytes.




## type PWMConfig

```go
type PWMConfig struct {
	// PWM period in nanosecond. Leaving this zero will pick a reasonable period
	// value for use with LEDs.
	// If you want to configure a frequency instead of a period, you can use the
	// following formula to calculate a period from a frequency:
	//
	//     period = 1e9 / frequency
	//
	Period uint64
}
```

PWMConfig allows setting some configuration while configuring a PWM
peripheral. A zero PWMConfig is ready to use for simple applications such as
dimming LEDs.





## type Pin

```go
type Pin uint8
```

Pin is a single pin on a chip, which may be connected to other hardware
devices. It can either be used directly as GPIO pin or it can be used in
other peripherals like ADC, I2C, etc.



### func (Pin) Configure

```go
func (p Pin) Configure(config PinConfig)
```

Configure configures the gpio pin as per mode.


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get reads the pin value.


### func (Pin) High

```go
func (p Pin) High()
```

High sets this GPIO pin to high, assuming it has been configured as an output
pin. It is hardware dependent (and often undefined) what happens if you set a
pin to high that is not configured as an output pin.


### func (Pin) Low

```go
func (p Pin) Low()
```

Low sets this GPIO pin to low, assuming it has been configured as an output
pin. It is hardware dependent (and often undefined) what happens if you set a
pin to low that is not configured as an output pin.


### func (Pin) Set

```go
func (p Pin) Set(value bool)
```

Set drives the pin high if value is true else drives it low.


### func (Pin) SetInterrupt

```go
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) error
```

SetInterrupt sets an interrupt to be executed when a particular pin changes
state. The pin should already be configured as an input, including a pull up
or down if no external pull is provided.

This call will replace a previously set callback on this pin. You can pass a
nil func to unset the pin change interrupt. If you do so, the change
parameter is ignored and can be set to any value (such as 0).




## type PinChange

```go
type PinChange uint8
```

PinChange represents one or more trigger events that can happen on a given GPIO pin
on the RP2040. ORed PinChanges are valid input to most IRQ functions.





## type PinConfig

```go
type PinConfig struct {
	Mode PinMode
}
```






## type PinMode

```go
type PinMode uint8
```

PinMode sets the direction and pull mode of the pin. For example, PinOutput
sets the pin as an output and PinInputPullup sets the pin as an input with a
pull-up.





## type RingBuffer

```go
type RingBuffer struct {
	rxbuffer	[bufferSize]volatile.Register8
	head		volatile.Register8
	tail		volatile.Register8
}
```

RingBuffer is ring buffer implementation inspired by post at
https://www.embeddedrelated.com/showthread/comp.arch.embedded/77084-1.php



### func (*RingBuffer) Clear

```go
func (rb *RingBuffer) Clear()
```

Clear resets the head and tail pointer to zero.


### func (*RingBuffer) Get

```go
func (rb *RingBuffer) Get() (byte, bool)
```

Get returns a byte from the buffer. If the buffer is empty,
the method will return a false as the second value.


### func (*RingBuffer) Put

```go
func (rb *RingBuffer) Put(val byte) bool
```

Put stores a byte in the buffer. If the buffer is already
full, the method will return false.


### func (*RingBuffer) Used

```go
func (rb *RingBuffer) Used() uint8
```

Used returns how many bytes in buffer have been used.




## type SPI

```go
type SPI struct {
	Bus *rp.SPI0_Type
}
```




### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig) error
```

Configure is intended to setup/initialize the SPI interface.
Default baudrate of 115200 is used if Frequency == 0. Default
word length (data bits) is 8.
Below is a list of GPIO pins corresponding to SPI0 bus on the rp2040:
 SI : 0, 4, 17  a.k.a RX and MISO (if rp2040 is master)
 SO : 3, 7, 19  a.k.a TX and MOSI (if rp2040 is master)
 SCK: 2, 6, 18
SPI1 bus GPIO pins:
 SI : 8, 12
 SO : 11, 15
 SCK: 10, 14
No pin configuration is needed of SCK, SDO and SDI needed after calling Configure.


### func (SPI) GetBaudRate

```go
func (spi SPI) GetBaudRate() uint32
```



### func (SPI) PrintRegs

```go
func (spi SPI) PrintRegs()
```

PrintRegs prints SPI's peripheral common registries current values


### func (SPI) SetBaudRate

```go
func (spi SPI) SetBaudRate(br uint32) error
```



### func (SPI) Transfer

```go
func (spi SPI) Transfer(w byte) (byte, error)
```

Write a single byte and read a single byte from TX/RX FIFO.


### func (SPI) Tx

```go
func (spi SPI) Tx(w, r []byte) (err error)
```

Tx handles read/write operation for SPI interface. Since SPI is a syncronous write/read
interface, there must always be the same number of bytes written as bytes read.
The Tx method knows about this, and offers a few different ways of calling it.

This form sends the bytes in tx buffer, putting the resulting bytes read into the rx buffer.
Note that the tx and rx buffers must be the same size:

		spi.Tx(tx, rx)

This form sends the tx buffer, ignoring the result. Useful for sending "commands" that return zeros
until all the bytes in the command packet have been received:

		spi.Tx(tx, nil)

This form sends zeros, putting the result into the rx buffer. Good for reading a "result packet":

		spi.Tx(nil, rx)

Remark: This implementation (RP2040) allows reading into buffer with a custom repeated
value on tx.

		spi.Tx([]byte{0xff}, rx) // may cause unwanted heap allocations.

This form sends 0xff and puts the result into rx buffer. Useful for reading from SD cards
which require 0xff input on SI.




## type SPIConfig

```go
type SPIConfig struct {
	Frequency	uint32
	// LSB not supported on rp2040.
	LSBFirst	bool
	// Mode's two most LSB are CPOL and CPHA. i.e. Mode==2 (0b10) is CPOL=1, CPHA=0
	Mode	uint8
	// Number of data bits per transfer. Valid values 4..16. Default and recommended is 8.
	DataBits	uint8
	// Serial clock pin
	SCK	Pin
	// TX or Serial Data Out (MOSI if rp2040 is master)
	SDO	Pin
	// RX or Serial Data In (MISO if rp2040 is master)
	SDI	Pin
}
```

SPIConfig is used to store config info for SPI.





## type UART

```go
type UART struct {
	Buffer		*RingBuffer
	Bus		*rp.UART0_Type
	Interrupt	interrupt.Interrupt
}
```

UART on the RP2040.



### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig) error
```

Configure the UART.


### func (*UART) Read

```go
func (uart *UART) Read(data []byte) (n int, err error)
```

Read from the RX buffer.


### func (*UART) ReadByte

```go
func (uart *UART) ReadByte() (byte, error)
```

ReadByte reads a single byte from the RX buffer.
If there is no data in the buffer, returns an error.


### func (*UART) Receive

```go
func (uart *UART) Receive(data byte)
```

Receive handles adding data to the UART's data buffer.
Usually called by the IRQ handler for a machine.


### func (*UART) SetBaudRate

```go
func (uart *UART) SetBaudRate(br uint32)
```

SetBaudRate sets the baudrate to be used for the UART.


### func (*UART) SetFormat

```go
func (uart *UART) SetFormat(databits, stopbits uint8, parity UARTParity) error
```

SetFormat for number of data bits, stop bits, and parity for the UART.


### func (*UART) Write

```go
func (uart *UART) Write(data []byte) (n int, err error)
```

Write data to the UART.


### func (*UART) WriteByte

```go
func (uart *UART) WriteByte(c byte) error
```

WriteByte writes a byte of data to the UART.




## type UARTConfig

```go
type UARTConfig struct {
	BaudRate	uint32
	TX		Pin
	RX		Pin
}
```

UARTConfig is a struct with which a UART (or similar object) can be
configured. The baud rate is usually respected, but TX and RX may be ignored
depending on the chip and the type of object.





## type UARTParity

```go
type UARTParity int
```

UARTParity is the parity setting to be used for UART communication.





