
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
	SDA_PIN	Pin	= GPIO12
	SCL_PIN	Pin	= GPIO13
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	Pin	= GPIO6
	SPI0_SDO_PIN	Pin	= GPIO7
	SPI0_SDI_PIN	Pin	= GPIO4
)
```

SPI pins


```go
const (
	NINA_SCK	Pin	= GPIO14
	NINA_SDO	Pin	= GPIO11
	NINA_SDI	Pin	= GPIO8

	NINA_CS		Pin	= GPIO9
	NINA_ACK	Pin	= GPIO10
	NINA_GPIO0	Pin	= GPIO0
	NINA_RESETN	Pin	= GPIO3

	NINA_TX	Pin	= GPIO9
	NINA_RX	Pin	= GPIO8
)
```

NINA-W102 Pins


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
)
```



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
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = DefaultUART
```

Serial is implemented via the default (usually the first) UART on the chip.





### func InitADC

```go
func InitADC()
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





