
---
title: arduino-uno-r4-minima
---


## Constants

```go
const (
	D0	Pin	= P3_01
	D1	Pin	= P3_02
	D2	Pin	= P1_05
	D3	Pin	= P1_04
	D4	Pin	= P1_03
	D5	Pin	= P1_02
	D6	Pin	= P1_06
	D7	Pin	= P1_07
	D8	Pin	= P3_04
	D9	Pin	= P3_03
	D10	Pin	= P1_12
	D11	Pin	= P1_09
	D12	Pin	= P1_10
	D13	Pin	= P1_11

	D14	Pin	= A0
	D15	Pin	= A1
)
```



```go
const (
	A0	Pin	= P0_14
	A1	Pin	= P0_00
	A2	Pin	= P0_01
	A3	Pin	= P0_02
	A4	Pin	= P1_01
	A5	Pin	= P1_00
)
```



```go
const (
	PWM_D3_PIN	Pin	= D3
	PWM_D5_PIN	Pin	= D5
	PWM_D6_PIN	Pin	= D6
	PWM_D9_PIN	Pin	= D9
	PWM_D10_PIN	Pin	= D10
	PWM_D11_PIN	Pin	= D11
)
```



```go
const (
	SWDIO_PIN	Pin	= P1_08
	SWCLK_PIN	Pin	= P3_00
)
```



```go
const (
	LED	= D13

	LED_TX	Pin	= P0_12
	LED_RX	Pin	= P0_13
)
```



```go
const (
	UART_TX_PIN	Pin	= UART1_TX_PIN
	UART_RX_PIN	Pin	= UART1_RX_PIN

	UART1_TX_PIN	Pin	= D1
	UART1_RX_PIN	Pin	= D0
	UART_SWD_TX_PIN	Pin	= P5_01
	UART_SWD_RX_PIN	Pin	= P5_02
)
```



```go
const (
	SDA_PIN	Pin	= A4
	SCL_PIN	Pin	= A5
)
```



```go
const (
	SPI_SDO_PIN	Pin	= D11
	SPI_SDI_PIN	Pin	= D12
	SPI_SCK_PIN	Pin	= D13
	SPI_CS_PIN	Pin	= D10
)
```



```go
const (
	CAN_TX_PIN	Pin	= D4
	CAN_RX_PIN	Pin	= D5
)
```



```go
const DAC0_PIN Pin = A0
```



```go
const AVCC_MEASURE_PIN Pin = P5_00
```



```go
const (
	ADC_CH0_PIN	Pin	= A1
	ADC_CH1_PIN	Pin	= A2
	ADC_CH2_PIN	Pin	= A3
	ADC_CH9_PIN	Pin	= A0
	ADC_CH16_PIN	Pin	= AVCC_MEASURE_PIN
	ADC_CH17_PIN	Pin	= UART_SWD_TX_PIN
	ADC_CH18_PIN	Pin	= UART_SWD_RX_PIN
	ADC_CH19_PIN	Pin	= D4
	ADC_CH20_PIN	Pin	= D5
	ADC_CH21_PIN	Pin	= A4
	ADC_CH22_PIN	Pin	= A5
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
	P0_00	Pin	= 0
	P0_01	Pin	= 1
	P0_02	Pin	= 2
	P0_03	Pin	= 3
	P0_04	Pin	= 4
	P0_05	Pin	= 5
	P0_06	Pin	= 6
	P0_07	Pin	= 7
	P0_08	Pin	= 8
	P0_09	Pin	= 9
	P0_10	Pin	= 10
	P0_11	Pin	= 11
	P0_12	Pin	= 12
	P0_13	Pin	= 13
	P0_14	Pin	= 14
	P0_15	Pin	= 15
	P1_00	Pin	= 16
	P1_01	Pin	= 17
	P1_02	Pin	= 18
	P1_03	Pin	= 19
	P1_04	Pin	= 20
	P1_05	Pin	= 21
	P1_06	Pin	= 22
	P1_07	Pin	= 23
	P1_08	Pin	= 24
	P1_09	Pin	= 25
	P1_10	Pin	= 26
	P1_11	Pin	= 27
	P1_12	Pin	= 28
	P1_13	Pin	= 29
	P1_14	Pin	= 30
	P1_15	Pin	= 31
	P2_00	Pin	= 32
	P2_01	Pin	= 33
	P2_02	Pin	= 34
	P2_03	Pin	= 35
	P2_04	Pin	= 36
	P2_05	Pin	= 37
	P2_06	Pin	= 38
	P2_07	Pin	= 39
	P2_08	Pin	= 40
	P2_09	Pin	= 41
	P2_10	Pin	= 42
	P2_11	Pin	= 43
	P2_12	Pin	= 44
	P2_13	Pin	= 45
	P2_14	Pin	= 46
	P2_15	Pin	= 47
	P3_00	Pin	= 48
	P3_01	Pin	= 49
	P3_02	Pin	= 50
	P3_03	Pin	= 51
	P3_04	Pin	= 52
	P3_05	Pin	= 53
	P3_06	Pin	= 54
	P3_07	Pin	= 55
	P3_08	Pin	= 56
	P3_09	Pin	= 57
	P3_10	Pin	= 58
	P3_11	Pin	= 59
	P3_12	Pin	= 60
	P3_13	Pin	= 61
	P3_14	Pin	= 62
	P3_15	Pin	= 63
	P4_00	Pin	= 64
	P4_01	Pin	= 65
	P4_02	Pin	= 66
	P4_03	Pin	= 67
	P4_04	Pin	= 68
	P4_05	Pin	= 69
	P4_06	Pin	= 70
	P4_07	Pin	= 71
	P4_08	Pin	= 72
	P4_09	Pin	= 73
	P4_10	Pin	= 74
	P4_11	Pin	= 75
	P4_12	Pin	= 76
	P4_13	Pin	= 77
	P4_14	Pin	= 78
	P4_15	Pin	= 79
	P5_00	Pin	= 80
	P5_01	Pin	= 81
	P5_02	Pin	= 82
	P5_03	Pin	= 83
	P5_04	Pin	= 84
	P5_05	Pin	= 85
	P5_06	Pin	= 86
	P5_07	Pin	= 87
	P5_08	Pin	= 88
	P5_09	Pin	= 89
	P5_10	Pin	= 90
	P5_11	Pin	= 91
	P5_12	Pin	= 92
	P5_13	Pin	= 93
	P5_14	Pin	= 94
	P5_15	Pin	= 95
)
```

P0_00 through P5_15 identify the RA4M1 port pins.


```go
const (
	PinOutput	PinMode	= iota
	PinInput
	PinInputPullup
	// PinInputPulldown acts as PinInput because the RA4M1 has no internal pull-down.
	// See ArduinoCore-renesas 1.6.0 cores/arduino/digital.cpp.
	PinInputPulldown
)
```

PinOutput through PinInputPulldown select the GPIO pin mode.


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
	UART1		= &sci2UART
	DefaultUART	= UART1
)
```

UART1 is on the D0/D1 header pins.


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
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = DefaultUART
```

Serial is implemented via the default (usually the first) UART on the chip.





### func AssignCore

```go
func AssignCore(core int)
```

AssignCore assigns the calling goroutine to core until it exits. Without the
cores scheduler only core 0 is valid.

Experimental: this API may change or be removed in a future release.


### func CPUReset

```go
func CPUReset()
```

CPUReset performs a hard system reset.


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

Configure sets the pin mode.


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get reports whether the pin input is high.


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

Set drives the pin high if value is true, or low if value is false.




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
	Bus		*renesas.SCI2_Type
	rxIRQ		interrupt.Interrupt
	errIRQ		interrupt.Interrupt
	rxIELSR		*volatile.Register32
	errIELSR	*volatile.Register32
	event		uint32	// ELC event number of RXI
	mstpMask	uint32
	psel		uint32
	txPin		Pin
	rxPin		Pin
}
```

UART on the RA4M1 is a SCI module in asynchronous mode.



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


### func (*UART) DTR

```go
func (uart *UART) DTR() bool
```

DTR is not supported by the SCI peripheral.


### func (*UART) RTS

```go
func (uart *UART) RTS() bool
```

RTS is not supported by the SCI peripheral.


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

SetBaudRate sets the communication speed for the UART.


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
	InvertTX	bool	// Invert TX line (active low becomes active high, etc.)
	InvertRX	bool	// Invert RX line (active low becomes active high, etc.)
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





