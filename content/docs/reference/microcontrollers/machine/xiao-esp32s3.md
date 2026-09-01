
---
title: xiao-esp32s3
---


## Constants

```go
const (
	D0	= GPIO1
	D1	= GPIO2
	D2	= GPIO3
	D3	= GPIO4
	D4	= GPIO5
	D5	= GPIO6
	D6	= GPIO43
	D7	= GPIO44
	D8	= GPIO7
	D9	= GPIO8
	D10	= GPIO9
)
```

Digital Pins


```go
const (
	A0	= GPIO1
	A1	= GPIO2
	A2	= GPIO3
	A3	= GPIO4
)
```

Analog pins


```go
const (
	UART_RX_PIN	= GPIO44
	UART_TX_PIN	= GPIO43
)
```

UART pins


```go
const (
	SDA_PIN	= GPIO5
	SCL_PIN	= GPIO6
)
```

I2C pins


```go
const (
	SPI1_SCK_PIN	= GPIO7	// D8
	SPI1_MISO_PIN	= GPIO8	// D9
	SPI1_MOSI_PIN	= GPIO9	// D10
	SPI1_CS_PIN	= NoPin

	SPI2_SCK_PIN	= NoPin
	SPI2_MOSI_PIN	= NoPin
	SPI2_MISO_PIN	= NoPin
	SPI2_CS_PIN	= NoPin
)
```

SPI pins


```go
const (
	LED = GPIO21
)
```

Onboard LEDs


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
	GPIO26	Pin	= 26
	GPIO27	Pin	= 27
	GPIO28	Pin	= 28
	GPIO29	Pin	= 29
	GPIO30	Pin	= 30
	GPIO31	Pin	= 31
	GPIO32	Pin	= 32
	GPIO33	Pin	= 33
	GPIO34	Pin	= 34
	GPIO35	Pin	= 35
	GPIO36	Pin	= 36
	GPIO37	Pin	= 37
	GPIO38	Pin	= 38
	GPIO39	Pin	= 39
	GPIO40	Pin	= 40
	GPIO41	Pin	= 41
	GPIO42	Pin	= 42
	GPIO43	Pin	= 43
	GPIO44	Pin	= 44
	GPIO45	Pin	= 45
	GPIO46	Pin	= 46
	GPIO47	Pin	= 47
	GPIO48	Pin	= 48
)
```

Hardware pin numbers


```go
const (
	ADC0	Pin	= GPIO1
	ADC2	Pin	= GPIO2
	ADC3	Pin	= GPIO3
	ADC4	Pin	= GPIO4
	ADC5	Pin	= GPIO5
	ADC6	Pin	= GPIO6
	ADC7	Pin	= GPIO7
	ADC8	Pin	= GPIO8
	ADC9	Pin	= GPIO9
	ADC10	Pin	= GPIO10
	ADC11	Pin	= GPIO11
	ADC12	Pin	= GPIO12
	ADC13	Pin	= GPIO13
	ADC14	Pin	= GPIO14
	ADC15	Pin	= GPIO15
	ADC16	Pin	= GPIO16
	ADC17	Pin	= GPIO17
	ADC18	Pin	= GPIO18
	ADC19	Pin	= GPIO19
	ADC20	Pin	= GPIO20
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
	I2CEXT0_SCL_OUT_IDX	= 89
	I2CEXT0_SDA_OUT_IDX	= 90
	I2CEXT1_SCL_OUT_IDX	= 91
	I2CEXT1_SDA_OUT_IDX	= 92
)
```



```go
const (
	LEDC_LS_SIG_OUT0_IDX = 73
)
```

GPIO matrix output signal indices for LEDC (soc/gpio_sig_map.h)


```go
const (
	// ESP32-S3 PLL clock frequency (same as ESP32-C3)
	pplClockFreq	= 80e6

	// Default SPI frequency - maximum safe speed
	SPI_DEFAULT_FREQUENCY	= 80e6	// 80MHz
)
```



```go
const (
	// IO MUX function number for SPI direct connection
	SPI_IOMUX_FUNC = 4
)
```



```go
const (
	// SPI2 (FSPI) signals - Hardware SPI2 - CORRECT VALUES from ESP-IDF
	SPI2_CLK_OUT_IDX	= uint32(101)	// FSPICLK_OUT_IDX
	SPI2_CLK_IN_IDX		= uint32(101)	// FSPICLK_IN_IDX
	SPI2_Q_OUT_IDX		= uint32(102)	// FSPIQ_OUT_IDX (MISO)
	SPI2_Q_IN_IDX		= uint32(102)	// FSPIQ_IN_IDX
	SPI2_D_OUT_IDX		= uint32(103)	// FSPID_OUT_IDX (MOSI)
	SPI2_D_IN_IDX		= uint32(103)	// FSPID_IN_IDX
	SPI2_CS0_OUT_IDX	= uint32(110)	// FSPICS0_OUT_IDX

	// SPI3 (HSPI) signals - Hardware SPI3 - CORRECTED from ESP-IDF gpio_sig_map.h
	// Source: /esp-idf/components/soc/esp32s3/include/soc/gpio_sig_map.h
	SPI3_CLK_OUT_IDX	= uint32(66)	// Line 136: SPI3_CLK_OUT_IDX
	SPI3_CLK_IN_IDX		= uint32(66)	// Line 135: SPI3_CLK_IN_IDX
	SPI3_Q_OUT_IDX		= uint32(67)	// Line 138: SPI3_Q_OUT_IDX (MISO)
	SPI3_Q_IN_IDX		= uint32(67)	// Line 137: SPI3_Q_IN_IDX
	SPI3_D_OUT_IDX		= uint32(68)	// Line 140: SPI3_D_OUT_IDX (MOSI)
	SPI3_D_IN_IDX		= uint32(68)	// Line 139: SPI3_D_IN_IDX
	SPI3_CS0_OUT_IDX	= uint32(71)	// Line 146: SPI3_CS0_OUT_IDX
)
```

ESP32-S3 GPIO Matrix signal indices for SPI - CORRECTED from ESP-IDF gpio_sig_map.h


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
	ErrInvalidSPIBus = errors.New("machine: invalid SPI bus")
)
```



```go
var DefaultUART = UART0
```



```go
var (
	UART0	= &_UART0
	_UART0	= UART{Bus: esp.UART0, Buffer: NewRingBuffer()}
	UART1	= &_UART1
	_UART1	= UART{Bus: esp.UART1, Buffer: NewRingBuffer()}
	UART2	= &_UART2
	_UART2	= UART{Bus: esp.UART2, Buffer: NewRingBuffer()}
)
```



```go
var (
	I2C0	= &I2C{
		Bus:		esp.I2C0,
		funcSCL:	I2CEXT0_SCL_OUT_IDX,
		funcSDA:	I2CEXT0_SDA_OUT_IDX,
		useExt1:	false,
	}
	I2C1	= &I2C{
		Bus:		esp.I2C1,
		funcSCL:	I2CEXT1_SCL_OUT_IDX,
		funcSDA:	I2CEXT1_SDA_OUT_IDX,
		useExt1:	true,
	}
)
```



```go
var (
	PWM0	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsS3, timerNum: 0}
	PWM1	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsS3, timerNum: 1}
	PWM2	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsS3, timerNum: 2}
	PWM3	= &LEDCPWM{SigOutBase: LEDC_LS_SIG_OUT0_IDX, NumChannels: ledcChannelsS3, timerNum: 3}
)
```



```go
var (
	SPI0	= &SPI{Bus: esp.SPI2, busID: 2}	// Primary SPI (FSPI)
	SPI1	= &SPI{Bus: esp.SPI3, busID: 3}	// Secondary SPI (HSPI)
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






### func ConfigureUSBEndpoint

```go
func ConfigureUSBEndpoint(desc descriptor.Descriptor, epSettings []usb.EndpointConfig, setup []usb.SetupConfig)
```

ConfigureUSBEndpoint is a no-op on ESP32-S3.


### func FlushSerial

```go
func FlushSerial()
```

FlushSerial flushes any pending USB serial TX data. Called from the
runtime (e.g. before sleeping) to ensure data from print() without
a trailing newline gets sent promptly.


### func GetCPUFrequency

```go
func GetCPUFrequency() (uint32, error)
```

GetCPUFrequency returns the current CPU frequency of the chip.


### func GetRNG

```go
func GetRNG() (ret uint32, err error)
```

GetRNG returns 32-bit random numbers using the ESP32-S3 true random number generator,
Random numbers are generated based on the thermal noise in the system and the
asynchronous clock mismatch.
For maximum entropy also make sure that the SAR_ADC is enabled.
See esp32-s3_technical_reference_manual_en.pdf p.920


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


### func ReadTemperature

```go
func ReadTemperature() int32
```

ReadTemperature reads the on-chip temperature sensor (TSENS) and returns
a value in millicelsius (°C × 1000). Uses the default measurement range
(offset = 0, approximately −10 °C to 80 °C, ±3 °C accuracy).

The conversion uses the same formula as ESP-IDF with no eFuse calibration:

	T = (0.4386 × raw − 20.52) °C


### func SendUSBInPacket

```go
func SendUSBInPacket(ep uint32, data []byte) bool
```

SendUSBInPacket is a no-op on ESP32-S3.


### func SendZlp

```go
func SendZlp()
```

SendZlp is a no-op on ESP32-S3.


### func SetCPUFrequency

```go
func SetCPUFrequency(frequency uint32) error
```

SetCPUFrequency sets the frequency of the CPU to one of several targets




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



### func (ADC) Get

```go
func (a ADC) Get() uint16
```



### func (ADC) GetVoltage

```go
func (a ADC) GetVoltage() (raw uint32, v float64)
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
	Bus	interface{}
	busID	uint8
}
```




### func (*SPI) Configure

```go
func (spi *SPI) Configure(config SPIConfig) error
```

Configure and make the SPI peripheral ready to use.
Implementation following ESP-IDF HAL with GPIO Matrix routing


### func (*SPI) Transfer

```go
func (spi *SPI) Transfer(w byte) (byte, error)
```

Transfer writes/reads a single byte using the SPI interface.
Implementation following ESP-IDF HAL spi_ll_user_start with proper USER register setup


### func (*SPI) Tx

```go
func (spi *SPI) Tx(w, r []byte) error
```

Tx handles read/write operation for SPI interface. Since SPI is a synchronous write/read
interface, there must always be the same number of bytes written as bytes read.
This is accomplished by sending zero bits if r is bigger than w or discarding
the incoming data if w is bigger than r.
Optimized implementation ported from ESP32-C3 for better performance.




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
	Bus	*esp.UART_Type
	Buffer	*RingBuffer
}
```




### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
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

USBDevice provides a stub USB device for the ESP32-S3. The hardware
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
It drains any data sitting in the hardware FIFO and re-enables the
peripheral-level USB interrupt (which the ISR disables to prevent a
level-triggered interrupt storm).


### func (*USB_DEVICE) Configure

```go
func (usbdev *USB_DEVICE) Configure(config UARTConfig) error
```

Configure initialises the USB Serial/JTAG controller.


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





