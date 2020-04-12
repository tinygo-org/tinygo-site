
---
title: arduino-mega2560
---


## Constants

```go
const (
	AREF	Pin	= NoPin
	LED	Pin	= PB7

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
	D2	Pin	= PE6
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
	D14	Pin	= PJ1
	D15	Pin	= PJ0
	D16	Pin	= PH1
	D17	Pin	= PH0
	D18	Pin	= PD3
	D19	Pin	= PD2
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
)
```



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
	PE3	= portE + 2
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
	PK4	= portH + 4
	PK5	= portH + 5
	PK6	= portH + 6
	PK7	= portH + 7
	PL0	= portL + 0
	PL1	= portL + 1
	PL2	= portL + 2
	PL3	= portL + 3
	PL4	= portL + 4
	PL5	= portL + 5
	PL6	= portL + 6
	PL7	= portE + 7
)
```



```go
const (
	PinInput	PinMode	= iota
	PinOutput
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
var I2C0 = I2C{}
```

I2C0 is the only I2C interface on most AVRs.


```go
var (
	// UART0 is the hardware serial port on the AVR.
	UART0 = UART{Buffer: NewRingBuffer()}
)
```

UART





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


### func InitPWM

```go
func InitPWM()
```

InitPWM initializes the registers needed for PWM.


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
func (a ADC) Configure()
```

Configure configures a ADCPin to be able to be used to read data.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get returns the current value of a ADC pin, in the range 0..0xffff. The AVR
has an ADC of 10 bits precision so the lower 6 bits will be zero.




## type I2C

```go
type I2C struct {
}
```

I2C on AVR.



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
	Frequency uint32
}
```

I2CConfig is used to store config info for I2C.





## type PWM

```go
type PWM struct {
	Pin Pin
}
```




### func (PWM) Configure

```go
func (pwm PWM) Configure()
```

Configure configures a PWM pin for output.


### func (PWM) Set

```go
func (pwm PWM) Set(value uint16)
```

Set turns on the duty cycle for a PWM pin using the provided value. On the AVR this is normally a
8-bit value ranging from 0 to 255.




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

Configure sets the pin to input or output.


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




## type UART

```go
type UART struct {
	Buffer *RingBuffer
}
```

UART on the AVR.



### func (UART) Buffered

```go
func (uart UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (UART) Configure

```go
func (uart UART) Configure(config UARTConfig)
```

Configure the UART on the AVR. Defaults to 9600 baud on Arduino.


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






