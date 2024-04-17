
---
title: arduino
---


## Constants

```go
const (
	D0	= PD0	// RX
	D1	= PD1	// TX
	D2	= PD2
	D3	= PD3
	D4	= PD4
	D5	= PD5
	D6	= PD6
	D7	= PD7
	D8	= PB0
	D9	= PB1
	D10	= PB2
	D11	= PB3
	D12	= PB4
	D13	= PB5
)
```

Digital pins, marked as plain numbers on the board.


```go
const LED Pin = D13
```

LED on the Arduino


```go
const (
	ADC0	Pin	= PC0
	ADC1	Pin	= PC1
	ADC2	Pin	= PC2
	ADC3	Pin	= PC3
	ADC4	Pin	= PC4	// Used by TWI for SDA
	ADC5	Pin	= PC5	// Used by TWI for SCL
)
```

ADC on the Arduino


```go
const (
	UART_TX_PIN	Pin	= PD1
	UART_RX_PIN	Pin	= PD0
)
```

UART pins


```go
const (
	PB0	= portB + 0
	PB1	= portB + 1	// peripherals: Timer1 channel A
	PB2	= portB + 2	// peripherals: Timer1 channel B
	PB3	= portB + 3	// peripherals: Timer2 channel A
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
	PD3	= portD + 3	// peripherals: Timer2 channel B
	PD4	= portD + 4
	PD5	= portD + 5	// peripherals: Timer0 channel B
	PD6	= portD + 6	// peripherals: Timer0 channel A
	PD7	= portD + 7
)
```



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
	PinRising	PinChange	= 1 << iota
	PinFalling
	PinToggle	= PinRising | PinFalling
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
var (
	Timer0	= PWM{0}	// 8 bit timer for PD5 and PD6
	Timer1	= PWM{1}	// 16 bit timer for PB1 and PB2
	Timer2	= PWM{2}	// 8 bit timer for PB3 and PD3
)
```



```go
var I2C0 = &I2C{
	srReg:	avr.TWSR,
	brReg:	avr.TWBR,
	crReg:	avr.TWCR,
	drReg:	avr.TWDR,
	srPS0:	avr.TWSR_TWPS0,
	srPS1:	avr.TWSR_TWPS1,
	crEN:	avr.TWCR_TWEN,
	crINT:	avr.TWCR_TWINT,
	crSTO:	avr.TWCR_TWSTO,
	crEA:	avr.TWCR_TWEA,
	crSTA:	avr.TWCR_TWSTA,
}
```

I2C0 is the only I2C interface on most AVRs.


```go
var SPI0 = SPI{
	spcr:	avr.SPCR,
	spdr:	avr.SPDR,
	spsr:	avr.SPSR,

	spcrR0:		avr.SPCR_SPR0,
	spcrR1:		avr.SPCR_SPR1,
	spcrCPHA:	avr.SPCR_CPHA,
	spcrCPOL:	avr.SPCR_CPOL,
	spcrDORD:	avr.SPCR_DORD,
	spcrSPE:	avr.SPCR_SPE,
	spcrMSTR:	avr.SPCR_MSTR,

	spsrI2X:	avr.SPSR_SPI2X,
	spsrSPIF:	avr.SPSR_SPIF,

	sck:	PB5,
	sdo:	PB3,
	sdi:	PB4,
	cs:	PB2,
}
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






## type PWM

```go
type PWM struct {
	num uint8
}
```

PWM is one PWM peripheral, which consists of a counter and two output
channels (that can be connected to two fixed pins). You can set the frequency
using SetPeriod, but only for all the channels in this PWM peripheral at
once.



### func (PWM) Channel

```go
func (pwm PWM) Channel(pin Pin) (uint8, error)
```

Channel returns a PWM channel for the given pin.


### func (PWM) Configure

```go
func (pwm PWM) Configure(config PWMConfig) error
```

Configure enables and configures this PWM.

For the two 8 bit timers, there is only a limited number of periods
available, namely the CPU frequency divided by 256 and again divided by 1, 8,
64, 256, or 1024. For a MCU running at 16MHz, this would be a period of 16µs,
128µs, 1024µs, 4096µs, or 16384µs.


### func (PWM) Counter

```go
func (pwm PWM) Counter() uint32
```

Counter returns the current counter value of the timer in this PWM
peripheral. It may be useful for debugging.


### func (PWM) Period

```go
func (pwm PWM) Period() uint64
```

Period returns the used PWM period in nanoseconds. It might deviate slightly
from the configured period due to rounding.


### func (PWM) Set

```go
func (pwm PWM) Set(channel uint8, value uint32)
```

Set updates the channel value. This is used to control the channel duty
cycle, in other words the fraction of time the channel output is high (or low
when inverted). For example, to set it to a 25% duty cycle, use:

	pwm.Set(channel, pwm.Top() / 4)

pwm.Set(channel, 0) will set the output to low and pwm.Set(channel,
pwm.Top()) will set the output to high, assuming the output isn't inverted.


### func (PWM) SetInverting

```go
func (pwm PWM) SetInverting(channel uint8, inverting bool)
```

SetInverting sets whether to invert the output of this channel.
Without inverting, a 25% duty cycle would mean the output is high for 25% of
the time and low for the rest. Inverting flips the output as if a NOT gate
was placed at the output, meaning that the output would be 25% low and 75%
high with a duty cycle of 25%.

Note: the invert state may not be applied on the AVR until the next call to
ch.Set().


### func (PWM) SetPeriod

```go
func (pwm PWM) SetPeriod(period uint64) error
```

SetPeriod updates the period of this PWM peripheral.
To set a particular frequency, use the following formula:

	period = 1e9 / frequency

If you use a period of 0, a period that works well for LEDs will be picked.

SetPeriod will not change the prescaler, but also won't change the current
value in any of the channels. This means that you may need to update the
value for the particular channel.

Note that you cannot pick any arbitrary period after the PWM peripheral has
been configured. If you want to switch between frequencies, pick the lowest
frequency (longest period) once when calling Configure and adjust the
frequency here as needed.


### func (PWM) Top

```go
func (pwm PWM) Top() uint32
```

Top returns the current counter top, for use in duty cycle calculation. It
will only change with a call to Configure or SetPeriod, otherwise it is
constant.

The value returned here is hardware dependent. In general, it's best to treat
it as an opaque value that can be divided by some number and passed to Set
(see Set documentation for more information).




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


### func (Pin) SetInterrupt

```go
func (pin Pin) SetInterrupt(pinChange PinChange, callback func(Pin)) (err error)
```





## type PinChange

```go
type PinChange uint8
```

Pin Change Interrupts





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





