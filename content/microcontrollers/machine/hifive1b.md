
---
title: hifive1b
---


## Constants

```go
const (
	P00	Pin	= 0
	P01	Pin	= 1
	P02	Pin	= 2
	P03	Pin	= 3
	P04	Pin	= 4
	P05	Pin	= 5
	P06	Pin	= 6
	P07	Pin	= 7
	P08	Pin	= 8
	P09	Pin	= 9
	P10	Pin	= 10
	P11	Pin	= 11
	P12	Pin	= 12
	P13	Pin	= 13
	P14	Pin	= 14
	P15	Pin	= 15
	P16	Pin	= 16
	P17	Pin	= 17
	P18	Pin	= 18
	P19	Pin	= 19
	P20	Pin	= 20
	P21	Pin	= 21
	P22	Pin	= 22
	P23	Pin	= 23
	P24	Pin	= 24
	P25	Pin	= 25
	P26	Pin	= 26
	P27	Pin	= 27
	P28	Pin	= 28
	P29	Pin	= 29
	P30	Pin	= 30
	P31	Pin	= 31
)
```



```go
const (
	D0	= P16
	D1	= P17
	D2	= P18
	D3	= P19	// Green LED/PWM (PWM1_PWM1)
	D4	= P20	// PWM (PWM1_PWM0)
	D5	= P21	// Blue LED/PWM (PWM1_PWM2)
	D6	= P22	// Red LED/PWM (PWM1_PWM3)
	D7	= P16
	D8	= NoPin	// PWM?
	D9	= P01
	D10	= P02	// SPI1_CS0
	D11	= P03	// SPI1_DQ0
	D12	= P04	// SPI1_DQ1
	D13	= P05	// SPI1_SCK
	D14	= NoPin	// not connected
	D15	= P09	// does not seem to work?
	D16	= P10	// PWM (PWM2_PWM0)
	D17	= P11	// PWM (PWM2_PWM1)
	D18	= P12	// SDA (I2C0_SDA)/PWM (PWM2_PWM2)
	D19	= P13	// SDL (I2C0_SCL)/PWM (PWM2_PWM3)
)
```



```go
const (
	LED		= LED1
	LED1		= LED_RED
	LED2		= LED_GREEN
	LED3		= LED_BLUE
	LED_RED		= P22
	LED_GREEN	= P19
	LED_BLUE	= P21
)
```



```go
const (
	// TODO: figure out the pin numbers for these.
	UART_TX_PIN	= NoPin
	UART_RX_PIN	= NoPin
)
```



```go
const (
	SPI0_SCK_PIN	= NoPin
	SPI0_MOSI_PIN	= NoPin
	SPI0_MISO_PIN	= NoPin

	SPI1_SCK_PIN	= D13
	SPI1_MOSI_PIN	= D11
	SPI1_MISO_PIN	= D12
)
```

SPI pins


```go
const (
	I2C0_SDA_PIN	= D18
	I2C0_SCL_PIN	= D19
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
	PinInput	PinMode	= iota
	PinOutput
	PinPWM
	PinSPI
	PinI2C	= PinSPI
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
	SPI1 = SPI{
		Bus: sifive.QSPI1,
	}
)
```

SPI on the HiFive1.


```go
var (
	I2C0 = I2C{
		Bus: sifive.I2C0,
	}
)
```

I2C on the HiFive1 rev B.


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
	UART0 = UART{Bus: sifive.UART0, Buffer: NewRingBuffer()}
)
```



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
	Bus *sifive.I2C_Type
}
```

I2C on the FE310-G002.



### func (I2C) Configure

```go
func (i2c I2C) Configure(config I2CConfig) error
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

Configure this pin with the given configuration.


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
	Bus *sifive.QSPI_Type
}
```

SPI on the FE310. The normal SPI0 is actually a quad-SPI meant for flash, so it is best
to use SPI1 or SPI2 port for most applications.



### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig) error
```

Configure is intended to setup the SPI interface.


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
	Bus	*sifive.UART_Type
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
func (uart UART) WriteByte(c byte)
```





## type UARTConfig

```go
type UARTConfig struct {
	BaudRate	uint32
	TX		Pin
	RX		Pin
}
```






