
---
title: stm32f4disco
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

	PD0	= portD + 0
	PD1	= portD + 1
	PD2	= portD + 2
	PD3	= portD + 3
	PD4	= portD + 4
	PD5	= portD + 5
	PD6	= portD + 6
	PD7	= portD + 7
	PD8	= portD + 8
	PD9	= portD + 9
	PD10	= portD + 10
	PD11	= portD + 11
	PD12	= portD + 12
	PD13	= portD + 13
	PD14	= portD + 14
	PD15	= portD + 15

	PE0	= portE + 0
	PE1	= portE + 1
	PE2	= portE + 2
	PE3	= portE + 3
	PE4	= portE + 4
	PE5	= portE + 5
	PE6	= portE + 6
	PE7	= portE + 7
	PE8	= portE + 8
	PE9	= portE + 9
	PE10	= portE + 10
	PE11	= portE + 11
	PE12	= portE + 12
	PE13	= portE + 13
	PE14	= portE + 14
	PE15	= portE + 15

	PH0	= portH + 0
	PH1	= portH + 1
)
```



```go
const (
	LED		= LED_BUILTIN
	LED1		= LED_GREEN
	LED2		= LED_ORANGE
	LED3		= LED_RED
	LED4		= LED_BLUE
	LED_BUILTIN	= LED_GREEN
	LED_GREEN	= PD12
	LED_ORANGE	= PD13
	LED_RED		= PD14
	LED_BLUE	= PD15
)
```



```go
const (
	UART_TX_PIN	= PA2
	UART_RX_PIN	= PA3
)
```

UART pins


```go
const NoPin = Pin(-1)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const CPU_FREQUENCY = 168000000
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
	PinModeUartTX	PinMode	= 4
	PinModeUartRX	PinMode	= 5

	//GPIOx_MODER
	GPIO_MODE_INPUT			= 0
	GPIO_MODE_GENERAL_OUTPUT	= 1
	GPIO_MODE_ALTERNABTIVE		= 2
	GPIO_MODE_ANALOG		= 3

	//GPIOx_OTYPER
	GPIO_OUTPUT_MODE_PUSH_PULL	= 0
	GPIO_OUTPUT_MODE_OPEN_DRAIN	= 1

	// GPIOx_OSPEEDR
	GPIO_SPEED_LOW		= 0
	GPIO_SPEED_MID		= 1
	GPIO_SPEED_HI		= 2
	GPIO_SPEED_VERY_HI	= 3

	// GPIOx_PUPDR
	GPIO_FLOATING	= 0
	GPIO_PULL_UP	= 1
	GPIO_PULL_DOWN	= 2
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
	// Both UART0 and UART1 refer to USART2.
	UART0	= UART{Buffer: NewRingBuffer()}
	UART1	= &UART0
)
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

It has some limitations currently due to how "volatile" variables that are
members of a struct are not compiled correctly by TinyGo.
See https://github.com/tinygo-org/tinygo/issues/151 for details.



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

UART



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






