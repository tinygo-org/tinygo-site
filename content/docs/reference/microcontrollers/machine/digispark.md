
---
title: digispark
---


## Constants

```go
const (
	P0	Pin	= PB0	// PWM available (Timer0 OC0A)
	P1	Pin	= PB1	// PWM available (Timer0 OC0B or Timer1 OC1A)
	P2	Pin	= PB2
	P3	Pin	= PB3
	P4	Pin	= PB4	// PWM available (Timer1 OC1B)
	P5	Pin	= PB5

	LED	= P1
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
	PB0	Pin	= iota
	PB1
	PB2
	PB3
	PB4
	PB5
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
	Timer0	= PWM{0}	// 8 bit timer for PB0 and PB1
	Timer1	= PWM{1}	// 8 bit high-speed timer for PB1 and PB4
)
```



```go
var SPI0 = SPI{}
```

SPI0 is the USI-based SPI interface on the ATTiny85


```go
var (
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = NullSerial{}
```

Serial is a null device: writes to it are ignored.


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

For Timer0, there is only a limited number of periods available, namely the
CPU frequency divided by 256 and again divided by 1, 8, 64, 256, or 1024.
For a MCU running at 8MHz, this would be a period of 32µs, 256µs, 2048µs,
8192µs, or 32768µs.

For Timer1, the period is more flexible as it uses OCR1C as the top value.
Timer1 also supports more prescaler values (1 to 16384).


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
	// Delay cycles for frequency control (0 = max speed)
	delayCycles	uint16

	// USICR value configured for the selected SPI mode
	usicrValue	uint8

	// LSB-first mode (requires software bit reversal)
	lsbFirst	bool
}
```

SPI is the USI-based SPI implementation for ATTiny85.
The ATTiny85 doesn't have dedicated SPI hardware, but uses the USI
(Universal Serial Interface) in three-wire mode.

Fixed pin mapping (directly controlled by USI hardware):
  - PB2: SCK (clock)
  - PB1: DO/MOSI (data out)
  - PB0: DI/MISO (data in)

Note: CS pin must be managed by the user.



### func (*SPI) Configure

```go
func (s *SPI) Configure(config SPIConfig) error
```

Configure sets up the USI for SPI communication.
Note: The user must configure and control the CS pin separately.


### func (*SPI) Transfer

```go
func (s *SPI) Transfer(b byte) (byte, error)
```

Transfer performs a single byte SPI transfer (send and receive simultaneously)
This implements the USI-based SPI transfer using the "clock strobing" technique


### func (*SPI) Tx

```go
func (spi *SPI) Tx(w, r []byte) error
```

Tx handles read/write operation for SPI interface. Since SPI is a synchronous write/read
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





