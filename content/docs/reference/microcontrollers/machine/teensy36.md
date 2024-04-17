
---
title: teensy36
---


## Constants

```go
const (
	D00	= PB16
	D01	= PB17
	D02	= PD00
	D03	= PA12
	D04	= PA13
	D05	= PD07
	D06	= PD04
	D07	= PD02
	D08	= PD03
	D09	= PC03
	D10	= PC04
	D11	= PC06
	D12	= PC07
	D13	= PC05
	D14	= PD01
	D15	= PC00
	D16	= PB00
	D17	= PB01
	D18	= PB03
	D19	= PB02
	D20	= PD05
	D21	= PD06
	D22	= PC01
	D23	= PC02
	D24	= PE26
	D25	= PA05
	D26	= PA14
	D27	= PA15
	D28	= PA16
	D29	= PB18
	D30	= PB19
	D31	= PB10
	D32	= PB11
	D33	= PE24
	D34	= PE25
	D35	= PC08
	D36	= PC09
	D37	= PC10
	D38	= PC11
	D39	= PA17
	D40	= PA28
	D41	= PA29
	D42	= PA26
	D43	= PB20
	D44	= PB22
	D45	= PB23
	D46	= PB21
	D47	= PD08
	D48	= PD09
	D49	= PB04
	D50	= PB05
	D51	= PD14
	D52	= PD13
	D53	= PD12
	D54	= PD15
	D55	= PD11
	D56	= PE10
	D57	= PE11
	D58	= PE00
	D59	= PE01
	D60	= PE02
	D61	= PE03
	D62	= PE04
	D63	= PE05
)
```

digital IO


```go
const LED = PC05
```

LED on the Teensy


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
	PinInput	PinMode	= iota
	PinInputPullup
	PinInputPulldown
	PinOutput
	PinOutputOpenDrain
	PinDisable
)
```



```go
const (
	PinInputPullUp		= PinInputPullup
	PinInputPullDown	= PinInputPulldown
)
```

Deprecated: use PinInputPullup and PinInputPulldown instead.


```go
const (
	PA00	Pin	= iota
	PA01
	PA02
	PA03
	PA04
	PA05
	PA06
	PA07
	PA08
	PA09
	PA10
	PA11
	PA12
	PA13
	PA14
	PA15
	PA16
	PA17
	PA18
	PA19
	PA20
	PA21
	PA22
	PA23
	PA24
	PA25
	PA26
	PA27
	PA28
	PA29
)
```



```go
const (
	PB00	Pin	= iota + 32
	PB01
	PB02
	PB03
	PB04
	PB05
	PB06
	PB07
	PB08
	PB09
	PB10
	PB11
	_
	_
	_
	_
	PB16
	PB17
	PB18
	PB19
	PB20
	PB21
	PB22
	PB23
)
```



```go
const (
	PC00	Pin	= iota + 64
	PC01
	PC02
	PC03
	PC04
	PC05
	PC06
	PC07
	PC08
	PC09
	PC10
	PC11
	PC12
	PC13
	PC14
	PC15
	PC16
	PC17
	PC18
	PC19
)
```



```go
const (
	PD00	Pin	= iota + 96
	PD01
	PD02
	PD03
	PD04
	PD05
	PD06
	PD07
	PD08
	PD09
	PD10
	PD11
	PD12
	PD13
	PD14
	PD15
)
```



```go
const (
	PE00	Pin	= iota + 128
	PE01
	PE02
	PE03
	PE04
	PE05
	PE06
	PE07
	PE08
	PE09
	PE10
	PE11
	PE12
	PE13
	PE14
	PE15
	PE16
	PE17
	PE18
	PE19
	PE20
	PE21
	PE22
	PE23
	PE24
	PE25
	PE26
	PE27
	PE28
)
```



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
	TeensyUART1	= UART0
	TeensyUART2	= UART1
	TeensyUART3	= UART2
	TeensyUART4	= UART3
	TeensyUART5	= UART4
)
```



```go
var DefaultUART = UART0
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
var (
	ErrNotImplemented	= errors.New("device has not been implemented")
	ErrNotConfigured	= errors.New("device has not been configured")
)
```



```go
var (
	UART0	= &_UART0
	UART1	= &_UART1
	UART2	= &_UART2
	UART3	= &_UART3
	UART4	= &_UART4
	_UART0	= UART{UART_Type: nxp.UART0, SCGC: &nxp.SIM.SCGC4, SCGCMask: nxp.SIM_SCGC4_UART0, DefaultRX: defaultUART0RX, DefaultTX: defaultUART0TX}
	_UART1	= UART{UART_Type: nxp.UART1, SCGC: &nxp.SIM.SCGC4, SCGCMask: nxp.SIM_SCGC4_UART1, DefaultRX: defaultUART1RX, DefaultTX: defaultUART1TX}
	_UART2	= UART{UART_Type: nxp.UART2, SCGC: &nxp.SIM.SCGC4, SCGCMask: nxp.SIM_SCGC4_UART2, DefaultRX: defaultUART2RX, DefaultTX: defaultUART2TX}
	_UART3	= UART{UART_Type: nxp.UART3, SCGC: &nxp.SIM.SCGC4, SCGCMask: nxp.SIM_SCGC4_UART3, DefaultRX: defaultUART3RX, DefaultTX: defaultUART3TX}
	_UART4	= UART{UART_Type: nxp.UART4, SCGC: &nxp.SIM.SCGC1, SCGCMask: nxp.SIM_SCGC1_UART4, DefaultRX: defaultUART4RX, DefaultTX: defaultUART4TX}
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

CPUFrequency returns the frequency of the ARM core clock (180MHz)


### func CPUReset

```go
func CPUReset()
```

CPUReset performs a hard system reset.


### func ClockFrequency

```go
func ClockFrequency() uint32
```

ClockFrequency returns the frequency of the external oscillator (16MHz)


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


### func PollUART

```go
func PollUART(u *UART)
```

PollUART manually checks a UART status and calls the ISR. This should only be
called by runtime.abort.


### func PutcharUART

```go
func PutcharUART(u *UART, c byte)
```

PutcharUART writes a byte to the UART synchronously, without using interrupts
or calling the scheduler




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





## type FastPin

```go
type FastPin struct {
	PDOR	*volatile.BitRegister
	PSOR	*volatile.BitRegister
	PCOR	*volatile.BitRegister
	PTOR	*volatile.BitRegister
	PDIR	*volatile.BitRegister
	PDDR	*volatile.BitRegister
}
```




### func (FastPin) Clear

```go
func (p FastPin) Clear()
```



### func (FastPin) Read

```go
func (p FastPin) Read() bool
```



### func (FastPin) Set

```go
func (p FastPin) Set()
```



### func (FastPin) Toggle

```go
func (p FastPin) Toggle()
```



### func (FastPin) Write

```go
func (p FastPin) Write(v bool)
```





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

Configure this pin with the given configuration.


### func (Pin) Control

```go
func (p Pin) Control() *volatile.Register32
```



### func (Pin) Fast

```go
func (p Pin) Fast() FastPin
```



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




## type UART

```go
type UART struct {
	*nxp.UART_Type
	SCGC		*volatile.Register32
	SCGCMask	uint32

	DefaultRX	Pin
	DefaultTX	Pin

	// state
	Buffer		RingBuffer	// RX Buffer
	TXBuffer	RingBuffer
	Configured	bool
	Transmitting	volatile.Register8
	Interrupt	interrupt.Interrupt
}
```




### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (u *UART) Configure(config UARTConfig)
```

Configure the UART.


### func (*UART) Disable

```go
func (u *UART) Disable()
```



### func (*UART) Flush

```go
func (u *UART) Flush()
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





