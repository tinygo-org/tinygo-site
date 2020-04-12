
---
title: bluepill
---


## Constants

```go
const (
	PA0	= portA + 0
	PA1	= portA + 1
	PA2	= portA + 2
	PA3	= portA + 3
	PA4	= portA + 4
	PA5	= portA + 5
	PA6	= portA + 6
	PA7	= portA + 7
	PA8	= portA + 8
	PA9	= portA + 9
	PA10	= portA + 10
	PA11	= portA + 11
	PA12	= portA + 12
	PA13	= portA + 13
	PA14	= portA + 14
	PA15	= portA + 15
	PB0	= portB + 0
	PB1	= portB + 1
	PB2	= portB + 2
	PB3	= portB + 3
	PB4	= portB + 4
	PB5	= portB + 5
	PB6	= portB + 6
	PB7	= portB + 7
	PB8	= portB + 8
	PB9	= portB + 9
	PB10	= portB + 10
	PB11	= portB + 11
	PB12	= portB + 12
	PB13	= portB + 13
	PB14	= portB + 14
	PB15	= portB + 15
	PC13	= portC + 13
	PC14	= portC + 14
	PC15	= portC + 15
)
```

https://wiki.stm32duino.com/index.php?title=File:Bluepillpinout.gif


```go
const (
	LED = PC13
)
```



```go
const (
	UART_TX_PIN	= PA9
	UART_RX_PIN	= PA10
	UART_ALT_TX_PIN	= PB6
	UART_ALT_RX_PIN	= PB7
)
```

UART pins


```go
const (
	SPI0_SCK_PIN	= PA5
	SPI0_MOSI_PIN	= PA7
	SPI0_MISO_PIN	= PA6
)
```

SPI pins


```go
const (
	SDA_PIN	= PB7
	SCL_PIN	= PB6
)
```

I2C pins


```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.


```go
const NoPin = Pin(-1)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	PinInput	PinMode	= 0	// Input mode
	PinOutput10MHz	PinMode	= 1	// Output mode, max speed 10MHz
	PinOutput2MHz	PinMode	= 2	// Output mode, max speed 2MHz
	PinOutput50MHz	PinMode	= 3	// Output mode, max speed 50MHz
	PinOutput	PinMode	= PinOutput2MHz

	PinInputModeAnalog	PinMode	= 0	// Input analog mode
	PinInputModeFloating	PinMode	= 4	// Input floating mode
	PinInputModePullUpDown	PinMode	= 8	// Input pull up/down mode
	PinInputModeReserved	PinMode	= 12	// Input mode (reserved)

	PinOutputModeGPPushPull		PinMode	= 0	// Output mode general purpose push/pull
	PinOutputModeGPOpenDrain	PinMode	= 4	// Output mode general purpose open drain
	PinOutputModeAltPushPull	PinMode	= 8	// Output mode alt. purpose push/pull
	PinOutputModeAltOpenDrain	PinMode	= 12	// Output mode alt. purpose open drain
)
```



```go
const (
	Mode0	= 0
	Mode1	= 1
	Mode2	= 2
	Mode3	= 3
)
```

SPI phase and polarity configs CPOL and CPHA





## Variables

```go
var (
	// USART1 is the first hardware serial port on the STM32.
	// Both UART0 and UART1 refer to USART1.
	UART0	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	stm32.USART1,
	}
	UART1	= &UART0
)
```



```go
var (
	ErrInvalidInputPin	= errors.New("machine: invalid input pin")
	ErrInvalidOutputPin	= errors.New("machine: invalid output pin")
	ErrInvalidClockPin	= errors.New("machine: invalid clock pin")
	ErrInvalidDataPin	= errors.New("machine: invalid data pin")
)
```



```go
var (
	SPI1	= SPI{Bus: stm32.SPI1}
	SPI0	= SPI1
)
```

There are 3 SPI interfaces on the STM32F103xx.
Since the first interface is named SPI1, both SPI0 and SPI1 refer to SPI1.
TODO: implement SPI2 and SPI3.


```go
var (
	I2C1	= I2C{Bus: stm32.I2C1}
	I2C0	= I2C1
)
```

There are 2 I2C interfaces on the STM32F103xx.
Since the first interface is named I2C1, both I2C0 and I2C1 refer to I2C1.
TODO: implement I2C2.


```go
var (
	ErrTxInvalidSliceSize = errors.New("SPI write and read slices must be same size")
)
```






### func CPUFrequency

```go
func CPUFrequency() uint32
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






## type I2C

```go
type I2C struct {
	Bus *stm32.I2C_Type
}
```




### func (I2C) Configure

```go
func (i2c I2C) Configure(config I2CConfig)
```

Configure is intended to setup the I2C interface.


### func (I2C) ReadRegister

```go
func (i2c I2C) ReadRegister(address uint8, register uint8, data []byte) error
```

ReadRegister transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (I2C) Tx

```go
func (i2c I2C) Tx(addr uint16, w, r []byte) error
```

Tx does a single I2C transaction at the specified address.
It clocks out the given address, writes the bytes in w, reads back len(r)
bytes and stores them in r, and generates a stop condition on the bus.


### func (I2C) WriteByte

```go
func (i2c I2C) WriteByte(data byte) error
```

WriteByte writes a single byte to the I2C bus.


### func (I2C) WriteRegister

```go
func (i2c I2C) WriteRegister(address uint8, register uint8, data []byte) error
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
	SCL		Pin
	SDA		Pin
}
```

I2CConfig is used to store config info for I2C.





## type PWM

```go
type PWM struct {
	Pin Pin
}
```






## type Pin

```go
type Pin int8
```

Pin is a single pin on a chip, which may be connected to other hardware
devices. It can either be used directly as GPIO pin or it can be used in
other peripherals like ADC, I2C, etc.



### func (Pin) Configure

```go
func (p Pin) Configure(config PinConfig)
```

Configure this pin with the given I/O settings.
stm32f1xx uses different technique for setting the GPIO pins than the stm32f407


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get returns the current value of a GPIO pin.


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
func (p Pin) Set(high bool)
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
	Bus *stm32.SPI_Type
}
```




### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig)
```

Configure is intended to setup the STM32 SPI1 interface.
Features still TODO:
- support SPI2 and SPI3
- allow setting data size to 16 bits?
- allow setting direction in HW for additional optimization?
- hardware SS pin?


### func (SPI) Transfer

```go
func (spi SPI) Transfer(w byte) (byte, error)
```

Transfer writes/reads a single byte using the SPI interface.


### func (SPI) Tx

```go
func (spi SPI) Tx(w, r []byte) error
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




## type SPIConfig

```go
type SPIConfig struct {
	Frequency	uint32
	SCK		Pin
	MOSI		Pin
	MISO		Pin
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig is used to store config info for SPI.





## type UART

```go
type UART struct {
	Buffer		*RingBuffer
	Bus		*stm32.USART_Type
	Interrupt	interrupt.Interrupt
}
```

UART representation



### func (UART) Buffered

```go
func (uart UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (UART) Configure

```go
func (uart UART) Configure(config UARTConfig)
```

Configure the UART.


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


### func (UART) SetBaudRate

```go
func (uart UART) SetBaudRate(br uint32)
```

SetBaudRate sets the communication speed for the UART. Defer to chip-specific
routines for calculation


### func (UART) Write

```go
func (uart UART) Write(data []byte) (n int, err error)
```

Write data to the UART.


### func (UART) WriteByte

```go
func (uart UART) WriteByte(c byte) error
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






