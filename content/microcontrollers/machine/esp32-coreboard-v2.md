
---
title: esp32-coreboard-v2
---


## Constants

```go
const LED = IO2
```

Built-in LED on some ESP32 boards.


```go
const (
	CLK	Pin	= 6
	CMD	Pin	= 11
	IO0	Pin	= 0
	IO1	Pin	= 1
	IO10	Pin	= 10
	IO16	Pin	= 16
	IO17	Pin	= 17
	IO18	Pin	= 18
	IO19	Pin	= 19
	IO2	Pin	= 2
	IO21	Pin	= 21
	IO22	Pin	= 22
	IO23	Pin	= 23
	IO25	Pin	= 25
	IO26	Pin	= 26
	IO27	Pin	= 27
	IO3	Pin	= 3
	IO32	Pin	= 32
	IO33	Pin	= 33
	IO34	Pin	= 34
	IO35	Pin	= 35
	IO36	Pin	= 36
	IO39	Pin	= 39
	IO4	Pin	= 4
	IO5	Pin	= 5
	IO9	Pin	= 9
	RXD	Pin	= 3
	SD0	Pin	= 7
	SD1	Pin	= 8
	SD2	Pin	= 9
	SD3	Pin	= 10
	SVN	Pin	= 39
	SVP	Pin	= 36
	TCK	Pin	= 13
	TD0	Pin	= 15
	TDI	Pin	= 12
	TMS	Pin	= 14
	TXD	Pin	= 1
)
```



```go
const (
	SPI0_SCK_PIN	= IO18
	SPI0_SDO_PIN	= IO23
	SPI0_SDI_PIN	= IO19
	SPI0_CS0_PIN	= IO5
)
```

SPI pins


```go
const (
	SDA_PIN	= IO21
	SCL_PIN	= IO22
)
```

I2C pins


```go
const (
	ADC0	Pin	= IO34
	ADC1	Pin	= IO35
	ADC2	Pin	= IO36
	ADC3	Pin	= IO39
)
```

ADC pins


```go
const (
	UART_TX_PIN	= IO1
	UART_RX_PIN	= IO3
)
```

UART0 pins


```go
const (
	UART1_TX_PIN	= IO9
	UART1_RX_PIN	= IO10
)
```

UART1 pins


```go
const (
	PWM0_PIN	Pin	= IO2
	PWM1_PIN	Pin	= IO0
	PWM2_PIN	Pin	= IO4
)
```

PWM pins


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
	ErrInvalidSPIBus = errors.New("machine: invalid SPI bus")
)
```



```go
var (
	UART0	= UART{Bus: esp.UART0, Buffer: NewRingBuffer()}
	UART1	= UART{Bus: esp.UART1, Buffer: NewRingBuffer()}
	UART2	= UART{Bus: esp.UART2, Buffer: NewRingBuffer()}
)
```



```go
var (
	// SPI0 and SPI1 are reserved for use by the caching system etc.
	SPI2	= SPI{esp.SPI2}
	SPI3	= SPI{esp.SPI3}
)
```






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






## type PWM

```go
type PWM struct {
	Pin Pin
}
```






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
input.


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




### func (UART) Buffered

```go
func (uart UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (UART) Configure

```go
func (uart UART) Configure(config UARTConfig)
```



### func (UART) Read

```go
func (uart UART) Read(data []byte) (n int, err error)
```

Read from the RX buffer.


### func (UART) ReadByte

```go
func (uart UART) ReadByte() (byte, error)
```

ReadByte reads a single byte from the RX buffer.
If there is no data in the buffer, returns an error.


### func (UART) Receive

```go
func (uart UART) Receive(data byte)
```

Receive handles adding data to the UART's data buffer.
Usually called by the IRQ handler for a machine.


### func (UART) Write

```go
func (uart UART) Write(data []byte) (n int, err error)
```

Write data to the UART.


### func (UART) WriteByte

```go
func (uart UART) WriteByte(b byte) error
```





## type UARTConfig

```go
type UARTConfig struct {
	BaudRate	uint32
	TX		Pin
	RX		Pin
}
```






