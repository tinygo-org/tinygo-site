
---
title: m5stack-core2
---


## Constants

```go
const (
	IO0	Pin	= 0
	IO1	Pin	= 1	// U0TXD
	IO2	Pin	= 2
	IO3	Pin	= 3	// U0RXD
	IO4	Pin	= 4
	IO5	Pin	= 5
	IO6	Pin	= 6	// SD_CLK
	IO7	Pin	= 7	// SD_DATA0
	IO8	Pin	= 8	// SD_DATA1
	IO9	Pin	= 9	// SD_DATA2
	IO10	Pin	= 10	// SD_DATA3
	IO11	Pin	= 11	// SD_CMD
	IO12	Pin	= 12
	IO13	Pin	= 13	// U0RXD
	IO14	Pin	= 14	// U1TXD
	IO15	Pin	= 15
	IO16	Pin	= 16
	IO17	Pin	= 17
	IO18	Pin	= 18	// SPI0_SCK
	IO19	Pin	= 19
	IO21	Pin	= 21	// SDA0
	IO22	Pin	= 22	// SCL0
	IO23	Pin	= 23	// SPI0_SDO
	IO25	Pin	= 25
	IO26	Pin	= 26
	IO27	Pin	= 27
	IO32	Pin	= 32	// SDA1
	IO33	Pin	= 33	// SCL1
	IO34	Pin	= 34
	IO35	Pin	= 35	// ADC1
	IO36	Pin	= 36	// ADC2
	IO38	Pin	= 38	// SPI0_SDI
	IO39	Pin	= 39
)
```



```go
const (
	SPI0_SCK_PIN	= IO18
	SPI0_SDO_PIN	= IO23
	SPI0_SDI_PIN	= IO38
	SPI0_CS0_PIN	= IO5

	// LCD (ILI9342C)
	LCD_SCK_PIN	= SPI0_SCK_PIN
	LCD_SDO_PIN	= SPI0_SDO_PIN
	LCD_SDI_PIN	= SPI0_SDI_PIN
	LCD_SS_PIN	= SPI0_CS0_PIN
	LCD_DC_PIN	= IO15

	// SD CARD
	SDCARD_SCK_PIN	= SPI0_SCK_PIN
	SDCARD_SDO_PIN	= SPI0_SDO_PIN
	SDCARD_SDI_PIN	= SPI0_SDI_PIN
	SDCARD_SS_PIN	= IO4
)
```

SPI pins


```go
const (
	// Internal I2C (AXP192 / FT6336U / BM8563 / MPU6886)
	SDA0_PIN	= IO21
	SCL0_PIN	= IO22

	// External I2C (PORT A)
	SDA1_PIN	= IO32
	SCL1_PIN	= IO33

	SDA_PIN	= SDA1_PIN
	SCL_PIN	= SCL1_PIN
)
```

I2C pins


```go
const (
	ADC1	Pin	= IO35
	ADC2	Pin	= IO36
)
```

ADC pins


```go
const (
	DAC1	Pin	= IO25
	DAC2	Pin	= IO26
)
```

DAC pins


```go
const (
	// UART0 (CP2104)
	UART0_TX_PIN	= IO1
	UART0_RX_PIN	= IO3

	UART1_TX_PIN	= IO14
	UART1_RX_PIN	= IO13

	UART_TX_PIN	= UART0_TX_PIN
	UART_RX_PIN	= UART0_RX_PIN
)
```

UART pins


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
	PinOutput	PinMode	= iota
	PinInput
	PinInputPullup
	PinInputPulldown
)
```



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
	ErrInvalidSPIBus = errors.New("machine: invalid SPI bus")
)
```



```go
var DefaultUART = UART0
```



```go
var (
	UART0	= &_UART0
	_UART0	= UART{Bus: esp.UART0, Buffer: NewRingBuffer()}
	UART1	= &_UART1
	_UART1	= UART{Bus: esp.UART1, Buffer: NewRingBuffer()}
	UART2	= &_UART2
	_UART2	= UART{Bus: esp.UART2, Buffer: NewRingBuffer()}
)
```



```go
var (
	// SPI0 and SPI1 are reserved for use by the caching system etc.
	SPI2	= SPI{esp.SPI2}
	SPI3	= SPI{esp.SPI3}
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

Get returns the current value of a GPIO pin when the pin is configured as an
input or as an output.


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
	Bus *esp.SPI_Type
}
```

Serial Peripheral Interface on the ESP32.



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
	SCK		Pin
	SDO		Pin
	SDI		Pin
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig configures a SPI peripheral on the ESP32. Make sure to set at least
SCK, SDO and SDI (possibly to NoPin if not in use). The default for LSBFirst
(false) and Mode (0) are good for most applications. The frequency defaults
to 1MHz if not set but can be configured up to 40MHz. Possible values are
40MHz and integer divisions from 40MHz such as 20MHz, 13.3MHz, 10MHz, 8MHz,
etc.





## type UART

```go
type UART struct {
	Bus	*esp.UART_Type
	Buffer	*RingBuffer
}
```




### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
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


### func (*UART) Write

```go
func (uart *UART) Write(data []byte) (n int, err error)
```

Write data to the UART.


### func (*UART) WriteByte

```go
func (uart *UART) WriteByte(b byte) error
```





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





