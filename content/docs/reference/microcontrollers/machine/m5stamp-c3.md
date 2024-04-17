
---
title: m5stamp-c3
---


## Constants

```go
const (
	IO0	= GPIO0
	IO1	= GPIO1
	IO2	= GPIO2
	IO3	= GPIO3
	IO4	= GPIO4
	IO5	= GPIO5
	IO6	= GPIO6
	IO7	= GPIO7
	IO8	= GPIO8
	IO9	= GPIO9
	IO10	= GPIO10
	IO11	= GPIO11
	IO12	= GPIO12
	IO13	= GPIO13
	IO14	= GPIO14
	IO15	= GPIO15
	IO16	= GPIO16
	IO17	= GPIO17
	IO18	= GPIO18
	IO19	= GPIO19
	IO20	= GPIO20
	IO21	= GPIO21

	XTAL_32K_P	= IO0
	XTAL_32K_N	= IO1
	MTMS		= IO4
	MTDI		= IO5
	MTCK		= IO6
	MTDO		= IO7
	VDD_SPI		= IO11
	SPIHD		= IO12
	SPISP		= IO13
	SPICS0		= IO14
	SPICLK		= IO15
	SPID		= IO16
	SPIQ		= IO17
	U0RXD		= IO20
	U0TXD		= IO21

	UART_TX_PIN	= U0TXD
	UART_RX_PIN	= U0RXD
)
```



```go
const (
	WS2812 = IO2
)
```



```go
const Device = deviceName
```

Device is the running program's chip name, such as "ATSAMD51J19A" or
"nrf52840". It is not the same as the CPU name.

The constant is some hardcoded default value if the program does not target a
particular chip but instead runs in WebAssembly for example.


```go
const (
	KHz	= 1000
	MHz	= 1000_000
	GHz	= 1000_000_000
)
```

Generic constants.


```go
const NoPin = Pin(0xff)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	PinOutput	PinMode	= iota
	PinInput
	PinInputPullup
	PinInputPulldown
)
```



```go
const (
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
)
```



```go
const (
	PinRising	PinChange	= iota + 1
	PinFalling
	PinToggle
)
```

Pin change interrupt constants for SetInterrupt.


```go
const (
	SPI_MODE0	= uint8(0)
	SPI_MODE1	= uint8(1)
	SPI_MODE2	= uint8(2)
	SPI_MODE3	= uint8(3)

	FSPICLK_IN_IDX	= uint32(63)
	FSPICLK_OUT_IDX	= uint32(63)
	FSPIQ_IN_IDX	= uint32(64)
	FSPIQ_OUT_IDX	= uint32(64)
	FSPID_IN_IDX	= uint32(65)
	FSPID_OUT_IDX	= uint32(65)
	FSPIHD_IN_IDX	= uint32(66)
	FSPIHD_OUT_IDX	= uint32(66)
	FSPIWP_IN_IDX	= uint32(67)
	FSPIWP_OUT_IDX	= uint32(67)
	FSPICS0_IN_IDX	= uint32(68)
	FSPICS0_OUT_IDX	= uint32(68)
	FSPICS1_OUT_IDX	= uint32(69)
	FSPICS2_OUT_IDX	= uint32(70)
	FSPICS3_OUT_IDX	= uint32(71)
	FSPICS4_OUT_IDX	= uint32(72)
	FSPICS5_OUT_IDX	= uint32(73)
)
```



```go
const (
	// ParityNone means to not use any parity checking. This is
	// the most common setting.
	ParityNone	UARTParity	= iota

	// ParityEven means to expect that the total number of 1 bits sent
	// should be an even number.
	ParityEven

	// ParityOdd means to expect that the total number of 1 bits sent
	// should be an odd number.
	ParityOdd
)
```






## Variables

```go
var (
	ErrTimeoutRNG		= errors.New("machine: RNG Timeout")
	ErrClockRNG		= errors.New("machine: RNG Clock Error")
	ErrSeedRNG		= errors.New("machine: RNG Seed Error")
	ErrInvalidInputPin	= errors.New("machine: invalid input pin")
	ErrInvalidOutputPin	= errors.New("machine: invalid output pin")
	ErrInvalidClockPin	= errors.New("machine: invalid clock pin")
	ErrInvalidDataPin	= errors.New("machine: invalid data pin")
	ErrNoPinChangeChannel	= errors.New("machine: no channel available for pin interrupt")
)
```



```go
var (
	DefaultUART	= UART0

	UART0	= &_UART0
	_UART0	= UART{Bus: esp.UART0, Buffer: NewRingBuffer()}
	UART1	= &_UART1
	_UART1	= UART{Bus: esp.UART1, Buffer: NewRingBuffer()}

	onceUart		= sync.Once{}
	errSamePins		= errors.New("UART: invalid pin combination")
	errWrongUART		= errors.New("UART: unsupported UARTn")
	errWrongBitSize		= errors.New("UART: invalid data size")
	errWrongStopBitSize	= errors.New("UART: invalid bit size")
)
```



```go
var (
	_USBCDC	= &USB_DEVICE{
		Bus: esp.USB_DEVICE,
	}

	USBCDC	Serialer	= _USBCDC
)
```



```go
var (
	ErrInvalidSPIBus	= errors.New("machine: SPI bus is invalid")
	ErrInvalidSPIMode	= errors.New("machine: SPI mode is invalid")
)
```



```go
var (
	// SPI0 and SPI1 are reserved for use by the caching system etc.
	SPI2 = SPI{esp.SPI2}
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

CPUFrequency returns the current CPU frequency of the chip.
Currently it is a fixed frequency but it may allow changing in the future.


### func GetRNG

```go
func GetRNG() (ret uint32, err error)
```

GetRNG returns 32-bit random numbers using the ESP32-C3 true random number generator,
Random numbers are generated based on the thermal noise in the system and the
asynchronous clock mismatch.
For maximum entropy also make sure that the SAR_ADC is enabled.
See esp32-c3_technical_reference_manual_en.pdf p.524


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.




## type ADC

```go
type ADC struct {
	Pin Pin
}
```






## type ADCConfig

```go
type ADCConfig struct {
	Reference	uint32	// analog reference voltage (AREF) in millivolts
	Resolution	uint32	// number of bits for a single conversion (e.g., 8, 10, 12)
	Samples		uint32	// number of samples for a single conversion (e.g., 4, 8, 16, 32)
	SampleTime	uint32	// sample time, in microseconds (µs)
}
```

ADCConfig holds ADC configuration parameters. If left unspecified, the zero
value of each parameter will use the peripheral's default settings.





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




## type PDMConfig

```go
type PDMConfig struct {
	Stereo	bool
	DIN	Pin
	CLK	Pin
}
```






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

Configure this pin with the given configuration.


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get returns the current value of a GPIO pin when configured as an input or as
an output.


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


### func (Pin) PortMaskClear

```go
func (p Pin) PortMaskClear() (*uint32, uint32)
```

Return the register and mask to disable a given GPIO pin. This can be used to
implement bit-banged drivers.

Warning: only use this on an output pin!


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```

Return the register and mask to enable a given GPIO pin. This can be used to
implement bit-banged drivers.

Warning: only use this on an output pin!


### func (Pin) Set

```go
func (p Pin) Set(value bool)
```

Set the pin to high or low.
Warning: only use this on an output pin!


### func (Pin) SetInterrupt

```go
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) (err error)
```

SetInterrupt sets an interrupt to be executed when a particular pin changes
state. The pin should already be configured as an input, including a pull up
or down if no external pull is provided.

You can pass a nil func to unset the pin change interrupt. If you do so,
the change parameter is ignored and can be set to any value (such as 0).
If the pin is already configured with a callback, you must first unset
this pins interrupt before you can set a new callback.




## type PinChange

```go
type PinChange uint8
```






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
	Bus *esp.SPI2_Type
}
```

Serial Peripheral Interface on the ESP32-C3.



### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig) error
```

Configure and make the SPI peripheral ready to use.


### func (SPI) Transfer

```go
func (spi SPI) Transfer(w byte) (byte, error)
```

Transfer writes/reads a single byte using the SPI interface. If you need to
transfer larger amounts of data, Tx will be faster.


### func (SPI) Tx

```go
func (spi SPI) Tx(w, r []byte) error
```

Tx handles read/write operation for SPI interface. Since SPI is a syncronous write/read
interface, there must always be the same number of bytes written as bytes read.
This is accomplished by sending zero bits if r is bigger than w or discarding
the incoming data if w is bigger than r.




## type SPIConfig

```go
type SPIConfig struct {
	Frequency	uint32
	SCK		Pin	// Serial Clock
	SDO		Pin	// Serial Data Out (MOSI)
	SDI		Pin	// Serial Data In  (MISO)
	CS		Pin	// Chip Select (optional)
	LSBFirst	bool	// MSB is default
	Mode		uint8	// SPI_MODE0 is default
}
```

SPIConfig is used to store config info for SPI.





## type Serialer

```go
type Serialer interface {
	WriteByte(c byte) error
	Write(data []byte) (n int, err error)
	Configure(config UARTConfig) error
	Buffered() int
	ReadByte() (byte, error)
	DTR() bool
	RTS() bool
}
```






## type UART

```go
type UART struct {
	Bus			*esp.UART_Type
	Buffer			*RingBuffer
	ParityErrorDetected	bool	// set when parity error detected
	DataErrorDetected	bool	// set when data corruption detected
	DataOverflowDetected	bool	// set when data overflow detected in UART FIFO buffer or RingBuffer
}
```




### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig) error
```



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
func (uart *UART) SetBaudRate(baudRate uint32)
```



### func (*UART) SetFormat

```go
func (uart *UART) SetFormat(dataBits, stopBits int, parity UARTParity) error
```



### func (*UART) Write

```go
func (uart *UART) Write(data []byte) (n int, err error)
```

Write data over the UART's Tx.
This function blocks until the data is finished being sent.


### func (*UART) WriteByte

```go
func (uart *UART) WriteByte(c byte) error
```

WriteByte writes a byte of data over the UART's Tx.
This function blocks until the data is finished being sent.




## type UARTConfig

```go
type UARTConfig struct {
	BaudRate	uint32
	TX		Pin
	RX		Pin
	RTS		Pin
	CTS		Pin
}
```

UARTConfig is a struct with which a UART (or similar object) can be
configured. The baud rate is usually respected, but TX and RX may be ignored
depending on the chip and the type of object.





## type UARTParity

```go
type UARTParity uint8
```

UARTParity is the parity setting to be used for UART communication.





## type USB_DEVICE

```go
type USB_DEVICE struct {
	Bus *esp.USB_DEVICE_Type
}
```

USB Serial/JTAG Controller
See esp32-c3_technical_reference_manual_en.pdf
pg. 736



### func (*USB_DEVICE) Buffered

```go
func (usbdev *USB_DEVICE) Buffered() int
```



### func (*USB_DEVICE) Configure

```go
func (usbdev *USB_DEVICE) Configure(config UARTConfig) error
```



### func (*USB_DEVICE) DTR

```go
func (usbdev *USB_DEVICE) DTR() bool
```



### func (*USB_DEVICE) RTS

```go
func (usbdev *USB_DEVICE) RTS() bool
```



### func (*USB_DEVICE) ReadByte

```go
func (usbdev *USB_DEVICE) ReadByte() (byte, error)
```



### func (*USB_DEVICE) Write

```go
func (usbdev *USB_DEVICE) Write(data []byte) (n int, err error)
```



### func (*USB_DEVICE) WriteByte

```go
func (usbdev *USB_DEVICE) WriteByte(c byte) error
```





