
---
title: esp32h2-devkitm-1
---


## Constants

```go
const (
	WS2812 = GPIO8
)
```

Data pin of the onboard addressable RGB LED. The board has no plain LED.


```go
const (
	BUTTON = GPIO9
)
```

BOOT button


```go
const (
	UART_TX_PIN	= GPIO24
	UART_RX_PIN	= GPIO23
)
```

UART pins


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
	PinOutput	PinMode	= iota
	PinInput
	PinInputPullup
	PinInputPulldown
	PinAnalog
)
```



```go
const (
	GPIO0	Pin	= 0
	GPIO1	Pin	= 1
	GPIO2	Pin	= 2
	GPIO3	Pin	= 3
	GPIO4	Pin	= 4
	GPIO5	Pin	= 5
	GPIO6	Pin	= 6
	GPIO7	Pin	= 7
	GPIO8	Pin	= 8
	GPIO9	Pin	= 9
	GPIO10	Pin	= 10
	GPIO11	Pin	= 11
	GPIO12	Pin	= 12
	GPIO13	Pin	= 13
	GPIO14	Pin	= 14
	GPIO15	Pin	= 15
	GPIO16	Pin	= 16
	GPIO17	Pin	= 17
	GPIO18	Pin	= 18
	GPIO19	Pin	= 19
	GPIO20	Pin	= 20
	GPIO21	Pin	= 21
	GPIO22	Pin	= 22
	GPIO23	Pin	= 23
	GPIO24	Pin	= 24
	GPIO25	Pin	= 25
	GPIO26	Pin	= 26
	GPIO27	Pin	= 27
)
```



```go
const (
	PinRising	PinChange	= iota + 1
	PinFalling
	PinToggle
)
```

Pin change interrupt constants for SetInterrupt.


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
var (
	_USBCDC	= &USB_DEVICE{
		Bus:	esp.USB_DEVICE,
		Buffer:	NewRingBuffer(),
	}

	USBCDC	Serialer	= _USBCDC
)
```



```go
var USBDev = &USBDevice{}
```



```go
var (
	DefaultUART	= UART0

	UART0	= &_UART0
	_UART0	= UART{Bus: esp.UART0, Buffer: NewRingBuffer()}
	UART1	= &_UART1
	_UART1	= UART{Bus: esp.UART1, Buffer: NewRingBuffer()}

	onceUart		= sync.Once{}
	errSamePins		= errors.New("UART: invalid pin combination")
	errWrongUART		= errors.New("UART: unsupported UARTn")
	errWrongBitSize		= errors.New("UART: invalid data size")
	errWrongStopBitSize	= errors.New("UART: invalid bit size")
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

CPUFrequency returns the current CPU frequency of the chip.
Currently it is a fixed frequency but it may allow changing in the future.


### func ConfigureUSBEndpoint

```go
func ConfigureUSBEndpoint(desc descriptor.Descriptor, epSettings []usb.EndpointConfig, setup []usb.SetupConfig)
```

ConfigureUSBEndpoint is a no-op on ESP32-C6 and ESP32-H2. The hardware does
not support programmable USB endpoints.


### func FlushSerial

```go
func FlushSerial()
```

FlushSerial flushes any pending USB serial TX data.


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


### func SendUSBInPacket

```go
func SendUSBInPacket(ep uint32, data []byte) bool
```

SendUSBInPacket is a no-op on ESP32-C6 and ESP32-H2.


### func SendZlp

```go
func SendZlp()
```

SendZlp is a no-op on ESP32-C6 and ESP32-H2.




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

Configure this pin with the given configuration.


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get returns the current value of a GPIO pin when configured as an input or as
an output.


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

Return the register and mask to disable a given GPIO pin. This can be used to
implement bit-banged drivers.

Warning: only use this on an output pin!


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```

Return the register and mask to enable a given GPIO pin. This can be used to
implement bit-banged drivers.

Warning: only use this on an output pin!


### func (Pin) Set

```go
func (p Pin) Set(value bool)
```

Set the pin to high or low.
Warning: only use this on an output pin!


### func (Pin) SetInterrupt

```go
func (p Pin) SetInterrupt(change PinChange, callback func(Pin)) (err error)
```

SetInterrupt sets an interrupt to be executed when a particular pin changes
state. The pin should already be configured as an input, including a pull up
or down if no external pull is provided.

You can pass a nil func to unset the pin change interrupt. If you do so,
the change parameter is ignored and can be set to any value (such as 0).
If the pin is already configured with a callback, you must first unset
this pins interrupt before you can set a new callback.




## type PinChange

```go
type PinChange uint8
```






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




## type Serialer

```go
type Serialer interface {
	WriteByte(c byte) error
	Write(data []byte) (n int, err error)
	Configure(config UARTConfig) error
	Buffered() int
	ReadByte() (byte, error)
	DTR() bool
	RTS() bool
}
```






## type UART

```go
type UART struct {
	Bus			*esp.UART_Type
	Buffer			*RingBuffer
	ParityErrorDetected	bool	// set when parity error detected
	DataErrorDetected	bool	// set when data corruption detected
	DataOverflowDetected	bool	// set when data overflow detected in UART FIFO buffer or RingBuffer
}
```




### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig) error
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


### func (*UART) SetBaudRate

```go
func (uart *UART) SetBaudRate(baudRate uint32)
```



### func (*UART) SetFormat

```go
func (uart *UART) SetFormat(dataBits, stopBits int, parity UARTParity) error
```



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





## type USBDevice

```go
type USBDevice struct {
	initcomplete		bool
	InitEndpointComplete	bool
}
```

USBDevice provides a stub USB device for the ESP32-C6 and ESP32-H2. The
hardware only supports a fixed-function CDC-ACM serial port, so the
programmable USB device features are no-ops.



### func (*USBDevice) Attach

```go
func (dev *USBDevice) Attach()
```

Attach and Detach are no-ops: the USB Serial/JTAG controller has no
software-controlled soft-connect, it is always attached to the bus.


### func (*USBDevice) ClearStallEPIn

```go
func (dev *USBDevice) ClearStallEPIn(ep uint32)
```



### func (*USBDevice) ClearStallEPOut

```go
func (dev *USBDevice) ClearStallEPOut(ep uint32)
```



### func (*USBDevice) Detach

```go
func (dev *USBDevice) Detach()
```



### func (*USBDevice) SetStallEPIn

```go
func (dev *USBDevice) SetStallEPIn(ep uint32)
```



### func (*USBDevice) SetStallEPOut

```go
func (dev *USBDevice) SetStallEPOut(ep uint32)
```





## type USB_DEVICE

```go
type USB_DEVICE struct {
	Bus		*esp.USB_DEVICE_Type
	Buffer		*RingBuffer
	txPending	bool	// unflushed data in the EP1 TX FIFO
	txStalled	bool	// set when flushAndWait fails (no host reading); cleared when FIFO becomes writable
}
```




### func (*USB_DEVICE) Buffered

```go
func (usbdev *USB_DEVICE) Buffered() int
```

Buffered returns the number of bytes waiting in the receive ring buffer.


### func (*USB_DEVICE) Configure

```go
func (usbdev *USB_DEVICE) Configure(config UARTConfig) error
```

Configure initialises the USB Serial/JTAG controller clock, pads, and
interrupt so that received data is buffered automatically.


### func (*USB_DEVICE) DTR

```go
func (usbdev *USB_DEVICE) DTR() bool
```



### func (*USB_DEVICE) RTS

```go
func (usbdev *USB_DEVICE) RTS() bool
```



### func (*USB_DEVICE) ReadByte

```go
func (usbdev *USB_DEVICE) ReadByte() (byte, error)
```

ReadByte returns a byte from the receive ring buffer.


### func (*USB_DEVICE) Write

```go
func (usbdev *USB_DEVICE) Write(data []byte) (n int, err error)
```



### func (*USB_DEVICE) WriteByte

```go
func (usbdev *USB_DEVICE) WriteByte(c byte) error
```





