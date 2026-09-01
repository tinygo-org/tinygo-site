
---
title: xiao-esp32c3
---


## Constants

```go
const (
	D0	= GPIO2
	D1	= GPIO3
	D2	= GPIO4
	D3	= GPIO5
	D4	= GPIO6
	D5	= GPIO7
	D6	= GPIO21
	D7	= GPIO20
	D8	= GPIO8
	D9	= GPIO9
	D10	= GPIO10
)
```

Digital Pins


```go
const (
	A0	= GPIO2
	A1	= GPIO3
	A2	= GPIO4
	A3	= GPIO5
)
```

Analog pins


```go
const (
	UART_RX_PIN	= GPIO20
	UART_TX_PIN	= GPIO21
)
```

UART pins


```go
const (
	SDA_PIN	= GPIO6
	SCL_PIN	= GPIO7
)
```

I2C pins


```go
const (
	SPI_SCK_PIN	= GPIO8
	SPI_SDI_PIN	= GPIO9
	SPI_SDO_PIN	= GPIO10
)
```

SPI pins


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
)
```



```go
const (
	ADC0	Pin	= GPIO0
	ADC1	Pin	= GPIO1
	ADC2	Pin	= GPIO2
	ADC3	Pin	= GPIO3
	ADC4	Pin	= GPIO4
	ADC5	Pin	= GPIO5	// avoid when WiFi is used.
)
```



```go
const (
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
	I2CEXT0_SCL_OUT_IDX	= 53
	I2CEXT0_SDA_OUT_IDX	= 54
)
```



```go
const (
	LEDC_LS_SIG_OUT0_IDX = 45
)
```

GPIO matrix output signal indices for LEDC (soc/gpio_sig_map.h)


```go
const (
	FSPICLK_IN_IDX	= uint32(63)
	FSPICLK_OUT_IDX	= uint32(63)
	FSPIQ_IN_IDX	= uint32(64)
	FSPIQ_OUT_IDX	= uint32(64)
	FSPID_IN_IDX	= uint32(65)
	FSPID_OUT_IDX	= uint32(65)
	FSPIHD_IN_IDX	= uint32(66)
	FSPIHD_OUT_IDX	= uint32(66)
	FSPIWP_IN_IDX	= uint32(67)
	FSPIWP_OUT_IDX	= uint32(67)
	FSPICS0_IN_IDX	= uint32(68)
	FSPICS0_OUT_IDX	= uint32(68)
	FSPICS1_OUT_IDX	= uint32(69)
	FSPICS2_OUT_IDX	= uint32(70)
	FSPICS3_OUT_IDX	= uint32(71)
	FSPICS4_OUT_IDX	= uint32(72)
	FSPICS5_OUT_IDX	= uint32(73)
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
	I2C0 = &I2C{
		Bus:		esp.I2C0,
		funcSCL:	I2CEXT0_SCL_OUT_IDX,
		funcSDA:	I2CEXT0_SDA_OUT_IDX,
		useExt1:	false,
	}
)
```



```go
var (
	PWM0	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsC3, timerNum: 0}
	PWM1	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsC3, timerNum: 1}
	PWM2	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsC3, timerNum: 2}
	PWM3	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsC3, timerNum: 3}
)
```



```go
var (
	ErrInvalidSPIBus	= errors.New("machine: SPI bus is invalid")
	ErrInvalidSPIMode	= errors.New("machine: SPI mode is invalid")
)
```



```go
var (
	// SPI0 and SPI1 are reserved for use by the caching system etc.
	SPI2	= &SPI{esp.SPI2}
	SPI0	= SPI2
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
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial Serialer
```

Serial is implemented via USB (USB-CDC).


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

CPUFrequency returns the current CPU frequency of the chip.
Currently it is a fixed frequency but it may allow changing in the future.


### func ConfigureUSBEndpoint

```go
func ConfigureUSBEndpoint(desc descriptor.Descriptor, epSettings []usb.EndpointConfig, setup []usb.SetupConfig)
```

ConfigureUSBEndpoint is a no-op on ESP32-C3 — the hardware does not
support programmable USB endpoints.


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


### func FlushSerial

```go
func FlushSerial()
```

FlushSerial flushes any pending USB serial TX data. Called from the
runtime (e.g. before sleeping) to ensure data from print() without
a trailing newline gets sent promptly.


### func GetRNG

```go
func GetRNG() (ret uint32, err error)
```

GetRNG returns 32-bit random numbers using the ESP32-C3 true random number generator,
Random numbers are generated based on the thermal noise in the system and the
asynchronous clock mismatch.
For maximum entropy also make sure that the SAR_ADC is enabled.
See esp32-c3_technical_reference_manual_en.pdf p.524


### func InitADC

```go
func InitADC()
```



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

SendUSBInPacket is a no-op on ESP32-C3 — the hardware does not
support arbitrary IN endpoints.  Returns false to indicate the
packet was not sent.


### func SendZlp

```go
func SendZlp()
```

SendZlp is a no-op on ESP32-C3 — the hardware handles control
transfers internally.




## type ADC

```go
type ADC struct {
	Pin Pin
}
```




### func (ADC) Configure

```go
func (a ADC) Configure(config ADCConfig) error
```

ESP32-C3: ADC1 = GPIO0–GPIO4 (ch 0–4), ADC2 = GPIO5 (ch 0). ADC2 shares with Wi‑Fi;
readings may be noisy when Wi‑Fi is active.


### func (ADC) Get

```go
func (a ADC) Get() uint16
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
	//
	// This interface directly writes data to the underlying block device.
	// Different kinds of devices have different requirements: most can only
	// write data after the page has been erased, and many can only write data
	// with specific alignment (such as 4-byte alignment).
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





## type I2C

```go
type I2C struct {
	Bus			*esp.I2C_Type
	funcSCL, funcSDA	uint32
	useExt1			bool
	txCmdBuf		[8]i2cCommand
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


### func (*I2C) SetBaudRate

```go
func (i2c *I2C) SetBaudRate(br uint32) error
```



### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) (err error)
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
	Frequency	uint32	// in Hz
	SCL		Pin
	SDA		Pin
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





## type LEDCPWM

```go
type LEDCPWM struct {
	SigOutBase	uint32	// GPIO matrix signal index for channel 0 (e.g. 73 on S3, 45 on C3)
	NumChannels	uint8
	timerNum	uint8	// 0–3: which LEDC timer (frequency) this PWM uses
	dutyRes		uint8
	configured	bool
	channelPin	[8]Pin
}
```




### func (*LEDCPWM) Channel

```go
func (pwm *LEDCPWM) Channel(pin Pin) (uint8, error)
```



### func (*LEDCPWM) Configure

```go
func (pwm *LEDCPWM) Configure(config PWMConfig) error
```



### func (*LEDCPWM) Set

```go
func (pwm *LEDCPWM) Set(channel uint8, value uint32)
```



### func (*LEDCPWM) SetInverting

```go
func (pwm *LEDCPWM) SetInverting(channel uint8, inverting bool)
```



### func (*LEDCPWM) Top

```go
func (pwm *LEDCPWM) Top() uint32
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




## type SPI

```go
type SPI struct {
	Bus *esp.SPI2_Type
}
```

Serial Peripheral Interface on the ESP32-C3.



### func (*SPI) Configure

```go
func (spi *SPI) Configure(config SPIConfig) error
```

Configure and make the SPI peripheral ready to use.


### func (*SPI) Transfer

```go
func (spi *SPI) Transfer(w byte) (byte, error)
```

Transfer writes/reads a single byte using the SPI interface. If you need to
transfer larger amounts of data, Tx will be faster.


### func (*SPI) Tx

```go
func (spi *SPI) Tx(w, r []byte) error
```

Tx handles read/write operation for SPI interface. Since SPI is a synchronous write/read
interface, there must always be the same number of bytes written as bytes read.
This is accomplished by sending zero bits if r is bigger than w or discarding
the incoming data if w is bigger than r.




## type SPIConfig

```go
type SPIConfig struct {
	Frequency	uint32
	SCK		Pin	// Serial Clock
	SDO		Pin	// Serial Data Out (MOSI)
	SDI		Pin	// Serial Data In  (MISO)
	CS		Pin	// Chip Select (optional)
	LSBFirst	bool	// MSB is default
	Mode		uint8	// Mode0 is default
}
```

SPIConfig is used to store config info for SPI.





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

USBDevice provides a stub USB device for the ESP32-C3.  The hardware
only supports a fixed-function CDC-ACM serial port, so the programmable
USB device features are no-ops.



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





