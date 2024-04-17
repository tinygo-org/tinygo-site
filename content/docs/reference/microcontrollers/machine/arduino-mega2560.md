
---
title: arduino-mega2560
---


## Constants

```go
const (
	A0	Pin	= PF0
	A1	Pin	= PF1
	A2	Pin	= PF2
	A3	Pin	= PF3
	A4	Pin	= PF4
	A5	Pin	= PF5
	A6	Pin	= PF6
	A7	Pin	= PF7
	A8	Pin	= PK0
	A9	Pin	= PK1
	A10	Pin	= PK2
	A11	Pin	= PK3
	A12	Pin	= PK4
	A13	Pin	= PK5
	A14	Pin	= PK6
	A15	Pin	= PK7

	// Analog Input
	ADC0	Pin	= PF0
	ADC1	Pin	= PF1
	ADC2	Pin	= PF2
	ADC3	Pin	= PF3
	ADC4	Pin	= PF4
	ADC5	Pin	= PF5
	ADC6	Pin	= PF6
	ADC7	Pin	= PF7
	ADC8	Pin	= PK0
	ADC9	Pin	= PK1
	ADC10	Pin	= PK2
	ADC11	Pin	= PK3
	ADC12	Pin	= PK4
	ADC13	Pin	= PK5
	ADC14	Pin	= PK6
	ADC15	Pin	= PK7

	// Digital pins
	D0	Pin	= PE0
	D1	Pin	= PE1
	D2	Pin	= PE4
	D3	Pin	= PE5
	D4	Pin	= PG5
	D5	Pin	= PE3
	D6	Pin	= PH3
	D7	Pin	= PH4
	D8	Pin	= PH5
	D9	Pin	= PH6
	D10	Pin	= PB4
	D11	Pin	= PB5
	D12	Pin	= PB6
	D13	Pin	= PB7
	D14	Pin	= PJ1	// TX3
	D15	Pin	= PJ0	// RX3
	D16	Pin	= PH1	// TX2
	D17	Pin	= PH0	// RX2
	D18	Pin	= PD3	// TX1
	D19	Pin	= PD2	// RX1
	D20	Pin	= PD1
	D21	Pin	= PD0
	D22	Pin	= PA0
	D23	Pin	= PA1
	D24	Pin	= PA2
	D25	Pin	= PA3
	D26	Pin	= PA4
	D27	Pin	= PA5
	D28	Pin	= PA6
	D29	Pin	= PA7
	D30	Pin	= PC7
	D31	Pin	= PC6
	D32	Pin	= PC5
	D33	Pin	= PC4
	D34	Pin	= PC3
	D35	Pin	= PC2
	D36	Pin	= PC1
	D37	Pin	= PC0
	D38	Pin	= PD7
	D39	Pin	= PG2
	D40	Pin	= PG1
	D41	Pin	= PG0
	D42	Pin	= PL7
	D43	Pin	= PL6
	D44	Pin	= PL5
	D45	Pin	= PL4
	D46	Pin	= PL3
	D47	Pin	= PL2
	D48	Pin	= PL1
	D49	Pin	= PL0
	D50	Pin	= PB3
	D51	Pin	= PB2
	D52	Pin	= PB1
	D53	Pin	= PB0

	AREF	Pin	= NoPin
	LED	Pin	= PB7
)
```



```go
const (
	UART_TX_PIN	Pin	= UART0_TX_PIN
	UART_RX_PIN	Pin	= UART0_RX_PIN
	UART0_TX_PIN	Pin	= D1
	UART0_RX_PIN	Pin	= D0
	UART1_TX_PIN	Pin	= D18
	UART1_RX_PIN	Pin	= D19
	UART2_TX_PIN	Pin	= D16
	UART2_RX_PIN	Pin	= D17
	UART3_TX_PIN	Pin	= D14
	UART3_RX_PIN	Pin	= D15
)
```

UART pins


```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.

Deprecated: use 100 * machine.KHz or 400 * machine.KHz instead.


```go
const (
	// I2CReceive indicates target has received a message from the controller.
	I2CReceive	I2CTargetEvent	= iota

	// I2CRequest indicates the controller is expecting a message from the target.
	I2CRequest

	// I2CFinish indicates the controller has ended the transaction.
	//
	// I2C controllers can chain multiple receive/request messages without
	// relinquishing the bus by doing 'restarts'.  I2CFinish indicates the
	// bus has been relinquished by an I2C 'stop'.
	I2CFinish
)
```



```go
const (
	// I2CModeController represents an I2C peripheral in controller mode.
	I2CModeController	I2CMode	= iota

	// I2CModeTarget represents an I2C peripheral in target mode.
	I2CModeTarget
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
	PA0	= portA + 0
	PA1	= portA + 1
	PA2	= portA + 2
	PA3	= portA + 3
	PA4	= portA + 4
	PA5	= portA + 5
	PA6	= portA + 6
	PA7	= portA + 7
	PB0	= portB + 0
	PB1	= portB + 1
	PB2	= portB + 2
	PB3	= portB + 3
	PB4	= portB + 4
	PB5	= portB + 5
	PB6	= portB + 6
	PB7	= portB + 7
	PC0	= portC + 0
	PC1	= portC + 1
	PC2	= portC + 2
	PC3	= portC + 3
	PC4	= portC + 4
	PC5	= portC + 5
	PC6	= portC + 6
	PC7	= portC + 7
	PD0	= portD + 0
	PD1	= portD + 1
	PD2	= portD + 2
	PD3	= portD + 3
	PD7	= portD + 7
	PE0	= portE + 0
	PE1	= portE + 1
	PE3	= portE + 3
	PE4	= portE + 4
	PE5	= portE + 5
	PE6	= portE + 6
	PF0	= portF + 0
	PF1	= portF + 1
	PF2	= portF + 2
	PF3	= portF + 3
	PF4	= portF + 4
	PF5	= portF + 5
	PF6	= portF + 6
	PF7	= portF + 7
	PG0	= portG + 0
	PG1	= portG + 1
	PG2	= portG + 2
	PG5	= portG + 5
	PH0	= portH + 0
	PH1	= portH + 1
	PH3	= portH + 3
	PH4	= portH + 4
	PH5	= portH + 5
	PH6	= portH + 6
	PJ0	= portJ + 0
	PJ1	= portJ + 1
	PK0	= portK + 0
	PK1	= portK + 1
	PK2	= portK + 2
	PK3	= portK + 3
	PK4	= portK + 4
	PK5	= portK + 5
	PK6	= portK + 6
	PK7	= portK + 7
	PL0	= portL + 0
	PL1	= portL + 1
	PL2	= portL + 2
	PL3	= portL + 3
	PL4	= portL + 4
	PL5	= portL + 5
	PL6	= portL + 6
	PL7	= portL + 7
)
```



```go
const (
	PinInput	PinMode	= iota
	PinInputPullup
	PinOutput
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
	UART1	= &_UART1
	_UART1	= UART{
		Buffer:	NewRingBuffer(),

		dataReg:	avr.UDR1,
		baudRegH:	avr.UBRR1H,
		baudRegL:	avr.UBRR1L,
		statusRegA:	avr.UCSR1A,
		statusRegB:	avr.UCSR1B,
		statusRegC:	avr.UCSR1C,
	}
	UART2	= &_UART2
	_UART2	= UART{
		Buffer:	NewRingBuffer(),

		dataReg:	avr.UDR2,
		baudRegH:	avr.UBRR2H,
		baudRegL:	avr.UBRR2L,
		statusRegA:	avr.UCSR2A,
		statusRegB:	avr.UCSR2B,
		statusRegC:	avr.UCSR2C,
	}
	UART3	= &_UART3
	_UART3	= UART{
		Buffer:	NewRingBuffer(),

		dataReg:	avr.UDR3,
		baudRegH:	avr.UBRR3H,
		baudRegL:	avr.UBRR3L,
		statusRegA:	avr.UCSR3A,
		statusRegB:	avr.UCSR3B,
		statusRegC:	avr.UCSR3C,
	}
)
```



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
var DefaultUART = UART0
```

Always use UART0 as the serial output.


```go
var (
	// UART0 is the hardware serial port on the AVR.
	UART0	= &_UART0
	_UART0	= UART{
		Buffer:	NewRingBuffer(),

		dataReg:	avr.UDR0,
		baudRegH:	avr.UBRR0H,
		baudRegL:	avr.UBRR0L,
		statusRegA:	avr.UCSR0A,
		statusRegB:	avr.UCSR0B,
		statusRegC:	avr.UCSR0C,
	}
)
```

UART


```go
var SPI0 = SPI{
	spcr:	avr.SPCR,
	spdr:	avr.SPDR,
	spsr:	avr.SPSR,
	sck:	PB1,
	sdo:	PB2,
	sdi:	PB3,
	cs:	PB0}
```

SPI configuration


```go
var (
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = DefaultUART
```

Serial is implemented via the default (usually the first) UART on the chip.


```go
var (
	ErrTxInvalidSliceSize		= errors.New("SPI write and read slices must be same size")
	errSPIInvalidMachineConfig	= errors.New("SPI port was not configured properly by the machine")
)
```






### func CPUFrequency

```go
func CPUFrequency() uint32
```

Return the current CPU frequency in hertz.


### func InitADC

```go
func InitADC()
```

InitADC initializes the registers needed for ADC.


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




### func (ADC) Configure

```go
func (a ADC) Configure(ADCConfig)
```

Configure configures a ADCPin to be able to be used to read data.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get returns the current value of a ADC pin, in the range 0..0xffff. The AVR
has an ADC of 10 bits precision so the lower 6 bits will be zero.




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





## type I2C

```go
type I2C struct {
	srReg	*volatile.Register8
	brReg	*volatile.Register8
	crReg	*volatile.Register8
	drReg	*volatile.Register8

	srPS0	byte
	srPS1	byte
	crEN	byte
	crINT	byte
	crSTO	byte
	crEA	byte
	crSTA	byte
}
```

I2C on AVR.



### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```

Configure is intended to setup the I2C interface.


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

SetBaudRate sets the communication speed for I2C.


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
```

Tx does a single I2C transaction at the specified address.
It clocks out the given address, writes the bytes in w, reads back len(r)
bytes and stores them in r, and generates a stop condition on the bus.


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
	Frequency uint32
}
```

I2CConfig is used to store config info for I2C.





## type I2CMode

```go
type I2CMode int
```

I2CMode determines if an I2C peripheral is in Controller or Target mode.





## type I2CTargetEvent

```go
type I2CTargetEvent uint8
```

I2CTargetEvent reflects events on the I2C bus





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

Configure sets the pin to input or output.


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
func (p Pin) PortMaskClear() (*volatile.Register8, uint8)
```

Return the register and mask to disable a given port. This can be used to
implement bit-banged drivers.

Warning: there are no separate pin set/clear registers on the AVR. The
returned mask is only valid as long as no other pin in the same port has been
changed.


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*volatile.Register8, uint8)
```

Return the register and mask to enable a given GPIO pin. This can be used to
implement bit-banged drivers.

Warning: there are no separate pin set/clear registers on the AVR. The
returned mask is only valid as long as no other pin in the same port has been
changed.


### func (Pin) Set

```go
func (p Pin) Set(value bool)
```

Set changes the value of the GPIO pin. The pin must be configured as output.




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
	// The registers for the SPIx port set by the chip
	spcr	*volatile.Register8
	spdr	*volatile.Register8
	spsr	*volatile.Register8

	spcrR0		byte
	spcrR1		byte
	spcrCPHA	byte
	spcrCPOL	byte
	spcrDORD	byte
	spcrSPE		byte
	spcrMSTR	byte

	spsrI2X		byte
	spsrSPIF	byte

	// The io pins for the SPIx port set by the chip
	sck	Pin
	sdi	Pin
	sdo	Pin
	cs	Pin
}
```

SPI is for the Serial Peripheral Interface
Data is taken from http://ww1.microchip.com/downloads/en/DeviceDoc/ATmega48A-PA-88A-PA-168A-PA-328-P-DS-DS40002061A.pdf page 169 and following



### func (SPI) Configure

```go
func (s SPI) Configure(config SPIConfig) error
```

Configure is intended to setup the SPI interface.


### func (SPI) Transfer

```go
func (s SPI) Transfer(b byte) (byte, error)
```

Transfer writes the byte into the register and returns the read content


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
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig is used to store config info for SPI.





## type UART

```go
type UART struct {
	Buffer	*RingBuffer

	dataReg		*volatile.Register8
	baudRegH	*volatile.Register8
	baudRegL	*volatile.Register8

	statusRegA	*volatile.Register8
	statusRegB	*volatile.Register8
	statusRegC	*volatile.Register8
}
```

UART on the AVR.



### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
```

Configure the UART on the AVR. Defaults to 9600 baud on Arduino.


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





