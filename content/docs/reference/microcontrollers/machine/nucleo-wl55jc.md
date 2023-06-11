
---
title: nucleo-wl55jc
---


## Constants

```go
const (
	LED_BLUE	= PB15
	LED_GREEN	= PB9
	LED_RED		= PB11
	LED		= LED_RED

	BTN1	= PA0
	BTN2	= PA1
	BTN3	= PC6
	BUTTON	= BTN1

	// SubGhz (SPI3)
	SPI0_NSS_PIN	= PA4
	SPI0_SCK_PIN	= PA5
	SPI0_SDO_PIN	= PA6
	SPI0_SDI_PIN	= PA7

	//MCU USART1
	UART1_TX_PIN	= PB6
	UART1_RX_PIN	= PB7

	//MCU USART2
	UART2_RX_PIN	= PA3
	UART2_TX_PIN	= PA2

	// DEFAULT USART
	UART_RX_PIN	= UART2_RX_PIN
	UART_TX_PIN	= UART2_TX_PIN

	// I2C1 pins
	I2C1_SCL_PIN	= PA9
	I2C1_SDA_PIN	= PA10
	I2C1_ALT_FUNC	= 4

	// I2C2 pins
	I2C2_SCL_PIN	= PA12
	I2C2_SDA_PIN	= PA11
	I2C2_ALT_FUNC	= 4

	// I2C0 alias for I2C1
	I2C0_SDA_PIN	= I2C1_SDA_PIN
	I2C0_SCL_PIN	= I2C1_SCL_PIN
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
	MAX_NBYTE_SIZE	= 255

	// 100ms delay = 100e6ns / 16ns
	// In runtime_stm32_timers.go, tick is fixed at 16ns per tick
	TIMEOUT_TICKS	= 100e6 / 16

	I2C_NO_STARTSTOP		= 0x0
	I2C_GENERATE_START_WRITE	= 0x80000000 | stm32.I2C_CR2_START
	I2C_GENERATE_START_READ		= 0x80000000 | stm32.I2C_CR2_START | stm32.I2C_CR2_RD_WRN
	I2C_GENERATE_STOP		= 0x80000000 | stm32.I2C_CR2_STOP
)
```



```go
const (
	// Mode Flag
	PinOutput		PinMode	= 0
	PinInput		PinMode	= PinInputFloating
	PinInputFloating	PinMode	= 1
	PinInputPulldown	PinMode	= 2
	PinInputPullup		PinMode	= 3

	// for UART
	PinModeUARTTX	PinMode	= 4
	PinModeUARTRX	PinMode	= 5

	// for I2C
	PinModeI2CSCL	PinMode	= 6
	PinModeI2CSDA	PinMode	= 7

	// for SPI
	PinModeSPICLK	PinMode	= 8
	PinModeSPISDO	PinMode	= 9
	PinModeSPISDI	PinMode	= 10

	// for analog/ADC
	PinInputAnalog	PinMode	= 11

	// for PWM
	PinModePWMOutput	PinMode	= 12
)
```



```go
const RNG_MAX_READ_RETRIES = 1000
```



```go
const PWM_MODE1 = 0x6
```



```go
const (
	AF0_SYSTEM		= 0
	AF1_TIM1_2_LPTIM1	= 1
	AF2_TIM1_2		= 2
	AF3_SPIS2_TIM1_LPTIM3	= 3
	AF4_I2C1_2_3		= 4
	AF5_SPI1_SPI2S2		= 5
	AF6_RF			= 6
	AF7_USART1_2		= 7
	AF8_LPUART1		= 8
	AF12_COMP1_2_TIM1	= 12
	AF13_DEBUG		= 13
	AF14_TIM2_16_17_LPTIM2	= 14
	AF15_EVENTOUT		= 15
)
```



```go
const (
	SYSCLK		= 48e6
	APB1_TIM_FREQ	= SYSCLK
	APB2_TIM_FREQ	= SYSCLK
)
```



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

	PC0	= portC + 0
	PC1	= portC + 1
	PC2	= portC + 2
	PC3	= portC + 3
	PC4	= portC + 4
	PC5	= portC + 5
	PC6	= portC + 6
	PC7	= portC + 7
	PC8	= portC + 8
	PC9	= portC + 9
	PC10	= portC + 10
	PC11	= portC + 11
	PC12	= portC + 12
	PC13	= portC + 13
	PC14	= portC + 14
	PC15	= portC + 15

	PH3	= portH + 3
)
```



```go
const (
	ARR_MAX	= 0x10000
	PSC_MAX	= 0x10000
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
	// STM32 UART2 is connected to the embedded STLINKV3 Virtual Com Port
	UART0	= &_UART0
	_UART0	= UART{
		Buffer:			NewRingBuffer(),
		Bus:			stm32.USART2,
		TxAltFuncSelector:	7,
		RxAltFuncSelector:	7,
	}

	// UART1 is free
	UART1	= &_UART1
	_UART1	= UART{
		Buffer:			NewRingBuffer(),
		Bus:			stm32.USART1,
		TxAltFuncSelector:	7,
		RxAltFuncSelector:	7,
	}

	DefaultUART	= UART0

	// I2C Busses
	I2C1	= &I2C{
		Bus:			stm32.I2C1,
		AltFuncSelector:	I2C1_ALT_FUNC,
	}
	I2C2	= &I2C{
		Bus:			stm32.I2C2,
		AltFuncSelector:	I2C2_ALT_FUNC,
	}
	I2C0	= I2C1

	// SPI
	SPI3	= SPI{
		Bus: stm32.SPI3,
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
var Flash flashBlockDevice
```



```go
var (
	TIM1	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM1EN,
		Device:		stm32.TIM1,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{{PA8, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{{PA9, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{{PA10, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{{PA11, AF1_TIM1_2_LPTIM1}}},
		},
		busFreq:	APB2_TIM_FREQ,
	}
	TIM2	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR1,
		EnableFlag:	stm32.RCC_APB1ENR1_TIM2EN,
		Device:		stm32.TIM2,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{{PA0, AF1_TIM1_2_LPTIM1}, {PA5, AF1_TIM1_2_LPTIM1}, {PA15, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{{PA1, AF1_TIM1_2_LPTIM1}, {PB3, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{{PA2, AF1_TIM1_2_LPTIM1}, {PB10, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{{PA3, AF1_TIM1_2_LPTIM1}, {PB11, AF1_TIM1_2_LPTIM1}}},
		},
		busFreq:	APB1_TIM_FREQ,
	}
	TIM16	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM16EN,
		Device:		stm32.TIM16,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{{PA6, AF14_TIM2_16_17_LPTIM2}}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB2_TIM_FREQ,
	}
	TIM17	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM17EN,
		Device:		stm32.TIM17,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{{PA7, AF1_TIM1_2_LPTIM1}, {PB9, AF1_TIM1_2_LPTIM1}}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB2_TIM_FREQ,
	}
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



### func CPUReset

```go
func CPUReset()
```

CPUReset performs a hard system reset.


### func FlashDataEnd

```go
func FlashDataEnd() uintptr
```

Return the end of the writable flash area. Usually this is the address one
past the end of the on-chip flash.


### func FlashDataStart

```go
func FlashDataStart() uintptr
```

Return the start of the writable flash area, aligned on a page boundary. This
is usually just after the program and static data.


### func GetRNG

```go
func GetRNG() (uint32, error)
```

GetRNG returns 32 bits of cryptographically secure random data


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





## type BlockDevice

```go
type BlockDevice interface {
	// ReadAt reads the given number of bytes from the block device.
	io.ReaderAt

	// WriteAt writes the given number of bytes to the block device.
	io.WriterAt

	// Size returns the number of bytes in this block device.
	Size() int64

	// WriteBlockSize returns the block size in which data can be written to
	// memory. It can be used by a client to optimize writes, non-aligned writes
	// should always work correctly.
	WriteBlockSize() int64

	// EraseBlockSize returns the smallest erasable area on this particular chip
	// in bytes. This is used for the block size in EraseBlocks.
	// It must be a power of two, and may be as small as 1. A typical size is 4096.
	EraseBlockSize() int64

	// EraseBlocks erases the given number of blocks. An implementation may
	// transparently coalesce ranges of blocks into larger bundles if the chip
	// supports this. The start and len parameters are in block numbers, use
	// EraseBlockSize to map addresses to blocks.
	EraseBlocks(start, len int64) error
}
```

BlockDevice is the raw device that is meant to store flash data.





## type ChannelCallback

```go
type ChannelCallback func(channel uint8)
```






## type I2C

```go
type I2C struct {
	Bus		*stm32.I2C_Type
	AltFuncSelector	uint8
}
```




### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```



### func (*I2C) ReadRegister

```go
func (i2c *I2C) ReadRegister(address uint8, register uint8, data []byte) error
```

ReadRegister transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
```



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
	SCL	Pin
	SDA	Pin
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

Configure this pin with the given configuration


### func (Pin) ConfigureAltFunc

```go
func (p Pin) ConfigureAltFunc(config PinConfig, altFunc uint8)
```

Configure this pin with the given configuration including alternate

	function mapping if necessary.


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

PortMaskClear returns the register and mask to disable a given port. This can
be used to implement bit-banged drivers.


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```

PortMaskSet returns the register and mask to enable a given GPIO pin. This
can be used to implement bit-banged drivers.


### func (Pin) Set

```go
func (p Pin) Set(high bool)
```

Set the pin to high or low.
Warning: only use this on an output pin!


### func (Pin) SetAltFunc

```go
func (p Pin) SetAltFunc(af uint8)
```

SetAltFunc maps the given alternative function to the I/O pin


### func (Pin) SetInterrupt

```go
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) error
```

SetInterrupt sets an interrupt to be executed when a particular pin changes
state. The pin should already be configured as an input, including a pull up
or down if no external pull is provided.

This call will replace a previously set callback on this pin. You can pass a
nil func to unset the pin change interrupt. If you do so, the change
parameter is ignored and can be set to any value (such as 0).




## type PinChange

```go
type PinChange uint8
```

---------- General pin operations ----------





## type PinConfig

```go
type PinConfig struct {
	Mode PinMode
}
```






## type PinFunction

```go
type PinFunction struct {
	Pin	Pin
	AltFunc	uint8
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
	Bus		*stm32.SPI_Type
	AltFuncSelector	uint8
}
```




### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig)
```

Configure is intended to setup the STM32 SPI1 interface.


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
	SDO		Pin
	SDI		Pin
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig is used to store config info for SPI.





## type TIM

```go
type TIM struct {
	EnableRegister	*volatile.Register32
	EnableFlag	uint32
	Device		*stm32.TIM_Type
	Channels	[4]TimerChannel
	UpInterrupt	interrupt.Interrupt
	OCInterrupt	interrupt.Interrupt

	wraparoundCallback	TimerCallback
	channelCallbacks	[4]ChannelCallback

	busFreq	uint64
}
```




### func (*TIM) Channel

```go
func (t *TIM) Channel(pin Pin) (uint8, error)
```

Channel returns a PWM channel for the given pin.


### func (*TIM) Configure

```go
func (t *TIM) Configure(config PWMConfig) error
```

Configure enables and configures this PWM.


### func (*TIM) Count

```go
func (t *TIM) Count() uint32
```



### func (*TIM) Set

```go
func (t *TIM) Set(channel uint8, value uint32)
```

Set updates the channel value. This is used to control the channel duty
cycle. For example, to set it to a 25% duty cycle, use:

	t.Set(ch, t.Top() / 4)

ch.Set(0) will set the output to low and ch.Set(ch.Top()) will set the output
to high, assuming the output isn't inverted.


### func (*TIM) SetInverting

```go
func (t *TIM) SetInverting(channel uint8, inverting bool)
```

SetInverting sets whether to invert the output of this channel.
Without inverting, a 25% duty cycle would mean the output is high for 25% of
the time and low for the rest. Inverting flips the output as if a NOT gate
was placed at the output, meaning that the output would be 25% low and 75%
high with a duty cycle of 25%.


### func (*TIM) SetMatchInterrupt

```go
func (t *TIM) SetMatchInterrupt(channel uint8, callback ChannelCallback) error
```

Sets a callback to be called when a channel reaches it's set-point.

For example, if `t.Set(ch, t.Top() / 4)` is used then the callback will
be called every quarter-period of the timer's base Period.


### func (*TIM) SetPeriod

```go
func (t *TIM) SetPeriod(period uint64) error
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


### func (*TIM) SetWraparoundInterrupt

```go
func (t *TIM) SetWraparoundInterrupt(callback TimerCallback) error
```

SetWraparoundInterrupt configures a callback to be called each
time the timer 'wraps-around'.

For example, if `Configure(PWMConfig{Period:1000000})` is used,
to set the timer period to 1ms, this callback will be called every
1ms.


### func (*TIM) Top

```go
func (t *TIM) Top() uint32
```

Top returns the current counter top, for use in duty cycle calculation. It
will only change with a call to Configure or SetPeriod, otherwise it is
constant.

The value returned here is hardware dependent. In general, it's best to treat
it as an opaque value that can be divided by some number and passed to
pwm.Set (see pwm.Set for more information).


### func (*TIM) Unset

```go
func (t *TIM) Unset(channel uint8)
```

Unset disables a channel, including any configured interrupts.




## type TimerCallback

```go
type TimerCallback func()
```






## type TimerChannel

```go
type TimerChannel struct {
	Pins []PinFunction
}
```






## type UART

```go
type UART struct {
	Buffer			*RingBuffer
	Bus			*stm32.USART_Type
	Interrupt		interrupt.Interrupt
	TxAltFuncSelector	uint8
	RxAltFuncSelector	uint8

	// Registers specific to the chip
	rxReg		*volatile.Register32
	txReg		*volatile.Register32
	statusReg	*volatile.Register32
	txEmptyFlag	uint32
}
```

UART representation



### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
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

SetBaudRate sets the communication speed for the UART. Defer to chip-specific
routines for calculation


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
type UARTParity uint8
```

UARTParity is the parity setting to be used for UART communication.





