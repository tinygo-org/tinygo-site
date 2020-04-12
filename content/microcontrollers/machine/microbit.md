
---
title: microbit
---


## Constants

```go
const HasLowFrequencyCrystal = false
```

The micro:bit does not have a 32kHz crystal on board.


```go
const (
	BUTTON	Pin	= BUTTONA
	BUTTONA	Pin	= 17
	BUTTONB	Pin	= 26
)
```

Buttons on the micro:bit (A and B)


```go
const (
	UART_TX_PIN	Pin	= 24
	UART_RX_PIN	Pin	= 25
)
```

UART pins


```go
const (
	ADC0	Pin	= 3	// P0 on the board
	ADC1	Pin	= 2	// P1 on the board
	ADC2	Pin	= 1	// P2 on the board
)
```

ADC pins


```go
const (
	SDA_PIN	Pin	= 30	// P20 on the board
	SCL_PIN	Pin	= 0	// P19 on the board
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	Pin	= 23	// P13 on the board
	SPI0_MOSI_PIN	Pin	= 21	// P15 on the board
	SPI0_MISO_PIN	Pin	= 22	// P14 on the board
)
```

SPI pins


```go
const (
	P0	Pin	= 3
	P1	Pin	= 2
	P2	Pin	= 1
	P3	Pin	= 4
	P4	Pin	= 5
	P5	Pin	= 17
	P6	Pin	= 12
	P7	Pin	= 11
	P8	Pin	= 18
	P9	Pin	= 10
	P10	Pin	= 6
	P11	Pin	= 26
	P12	Pin	= 20
	P13	Pin	= 23
	P14	Pin	= 22
	P15	Pin	= 21
	P16	Pin	= 16
)
```

GPIO/Analog pins


```go
const (
	LED_COL_1	Pin	= 4
	LED_COL_2	Pin	= 5
	LED_COL_3	Pin	= 6
	LED_COL_4	Pin	= 7
	LED_COL_5	Pin	= 8
	LED_COL_6	Pin	= 9
	LED_COL_7	Pin	= 10
	LED_COL_8	Pin	= 11
	LED_COL_9	Pin	= 12
	LED_ROW_1	Pin	= 13
	LED_ROW_2	Pin	= 14
	LED_ROW_3	Pin	= 15
)
```

LED matrix pins


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
	PinInput		PinMode	= (nrf.GPIO_PIN_CNF_DIR_Input << nrf.GPIO_PIN_CNF_DIR_Pos) | (nrf.GPIO_PIN_CNF_INPUT_Connect << nrf.GPIO_PIN_CNF_INPUT_Pos)
	PinInputPullup		PinMode	= PinInput | (nrf.GPIO_PIN_CNF_PULL_Pullup << nrf.GPIO_PIN_CNF_PULL_Pos)
	PinInputPulldown	PinMode	= PinInput | (nrf.GPIO_PIN_CNF_PULL_Pulldown << nrf.GPIO_PIN_CNF_PULL_Pos)
	PinOutput		PinMode	= (nrf.GPIO_PIN_CNF_DIR_Output << nrf.GPIO_PIN_CNF_DIR_Pos) | (nrf.GPIO_PIN_CNF_INPUT_Disconnect << nrf.GPIO_PIN_CNF_INPUT_Pos)
)
```






## Variables

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
	ErrTxInvalidSliceSize = errors.New("SPI write and read slices must be same size")
)
```



```go
var (
	// NRF_UART0 is the hardware UART on the NRF SoC.
	NRF_UART0 = UART{Buffer: NewRingBuffer()}
)
```

UART


```go
var (
	I2C0	= I2C{Bus: nrf.TWI0}
	I2C1	= I2C{Bus: nrf.TWI1}
)
```

There are 2 I2C interfaces on the NRF.


```go
var (
	SPI0	= SPI{Bus: nrf.SPI0}
	SPI1	= SPI{Bus: nrf.SPI1}
)
```

There are 2 SPI interfaces on the NRF5x.


```go
var (
	UART0 = NRF_UART0
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
	Bus *nrf.TWI_Type
}
```

I2C on the NRF.



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


### func (Pin) PortMaskClear

```go
func (p Pin) PortMaskClear() (*uint32, uint32)
```

Return the register and mask to disable a given port. This can be used to
implement bit-banged drivers.


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```

Return the register and mask to enable a given GPIO pin. This can be used to
implement bit-banged drivers.


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
	Bus *nrf.SPI_Type
}
```

SPI on the NRF.



### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig)
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
	Buffer *RingBuffer
}
```

UART on the NRF.



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

SetBaudRate sets the communication speed for the UART.


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






