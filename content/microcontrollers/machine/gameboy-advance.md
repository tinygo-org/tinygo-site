
---
title: gameboy-advance
---


## Constants

```go
const NoPin = Pin(-1)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	IRQ_VBLANK	= 0
	IRQ_HBLANK	= 1
	IRQ_VCOUNT	= 2
	IRQ_TIMER0	= 3
	IRQ_TIMER1	= 4
	IRQ_TIMER2	= 5
	IRQ_TIMER3	= 6
	IRQ_COM		= 7
	IRQ_DMA0	= 8
	IRQ_DMA1	= 9
	IRQ_DMA2	= 10
	IRQ_DMA3	= 11
	IRQ_KEYPAD	= 12
	IRQ_GAMEPAK	= 13
)
```

Interrupt numbers as used on the GameBoy Advance. Register them with
runtime/interrupt.New.





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
var Display = FramebufDisplay{(*[160][240]volatile.Register16)(unsafe.Pointer(uintptr(0x06000000)))}
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






## type FramebufDisplay

```go
type FramebufDisplay struct {
	port *[160][240]volatile.Register16
}
```




### func (FramebufDisplay) Configure

```go
func (d FramebufDisplay) Configure()
```



### func (FramebufDisplay) Display

```go
func (d FramebufDisplay) Display() error
```



### func (FramebufDisplay) SetPixel

```go
func (d FramebufDisplay) SetPixel(x, y int16, c color.RGBA)
```



### func (FramebufDisplay) Size

```go
func (d FramebufDisplay) Size() (x, y int16)
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

Set has not been implemented.




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




