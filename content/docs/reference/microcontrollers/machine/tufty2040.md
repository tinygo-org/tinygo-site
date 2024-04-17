
---
title: tufty2040
---


## Constants

```go
const (
	LED	Pin	= GPIO25

	BUTTON_A	Pin	= GPIO7
	BUTTON_B	Pin	= GPIO8
	BUTTON_C	Pin	= GPIO9
	BUTTON_UP	Pin	= GPIO22
	BUTTON_DOWN	Pin	= GPIO6
	BUTTON_USER	Pin	= GPIO23

	LCD_BACKLIGHT	Pin	= GPIO2
	LCD_CS		Pin	= GPIO10
	LCD_DC		Pin	= GPIO11
	LCD_WR		Pin	= GPIO12
	LCD_RD		Pin	= GPIO13
	LCD_DB0		Pin	= GPIO14
	LCD_DB1		Pin	= GPIO15
	LCD_DB2		Pin	= GPIO16
	LCD_DB3		Pin	= GPIO17
	LCD_DB4		Pin	= GPIO18
	LCD_DB5		Pin	= GPIO19
	LCD_DB6		Pin	= GPIO20
	LCD_DB7		Pin	= GPIO21

	VBUS_DETECT	Pin	= GPIO24
	BATTERY		Pin	= GPIO29
	USER_LED	Pin	= GPIO25
	LIGHT_SENSE	Pin	= GPIO26
	SENSOR_POWER	Pin	= GPIO27
)
```



```go
const (
	I2C0_SDA_PIN	Pin	= GPIO4
	I2C0_SCL_PIN	Pin	= GPIO5

	I2C1_SDA_PIN	Pin	= NoPin
	I2C1_SCL_PIN	Pin	= NoPin
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	Pin	= NoPin
	SPI0_SDO_PIN	Pin	= NoPin
	SPI0_SDI_PIN	Pin	= NoPin

	SPI1_SCK_PIN	Pin	= NoPin
	SPI1_SDO_PIN	Pin	= NoPin
	SPI1_SDI_PIN	Pin	= NoPin
)
```

SPI pins.


```go
const (
	UART0_TX_PIN	= GPIO0
	UART0_RX_PIN	= GPIO1
	UART_TX_PIN	= UART0_TX_PIN
	UART_RX_PIN	= UART0_RX_PIN
)
```

UART pins


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
	PinInputPulldown
	PinInputPullup
	PinAnalog
	PinUART
	PinPWM
	PinI2C
	PinSPI
	PinPIO0
	PinPIO1
)
```



```go
const (
	// Edge falling
	PinFalling	PinChange	= 4 << iota
	// Edge rising
	PinRising

	PinToggle	= PinFalling | PinRising
)
```

Pin change interrupt constants for SetInterrupt.


```go
const (
	// GPIO pins
	GPIO0	Pin	= 0	// peripherals: PWM0 channel A
	GPIO1	Pin	= 1	// peripherals: PWM0 channel B
	GPIO2	Pin	= 2	// peripherals: PWM1 channel A
	GPIO3	Pin	= 3	// peripherals: PWM1 channel B
	GPIO4	Pin	= 4	// peripherals: PWM2 channel A
	GPIO5	Pin	= 5	// peripherals: PWM2 channel B
	GPIO6	Pin	= 6	// peripherals: PWM3 channel A
	GPIO7	Pin	= 7	// peripherals: PWM3 channel B
	GPIO8	Pin	= 8	// peripherals: PWM4 channel A
	GPIO9	Pin	= 9	// peripherals: PWM4 channel B
	GPIO10	Pin	= 10	// peripherals: PWM5 channel A
	GPIO11	Pin	= 11	// peripherals: PWM5 channel B
	GPIO12	Pin	= 12	// peripherals: PWM6 channel A
	GPIO13	Pin	= 13	// peripherals: PWM6 channel B
	GPIO14	Pin	= 14	// peripherals: PWM7 channel A
	GPIO15	Pin	= 15	// peripherals: PWM7 channel B
	GPIO16	Pin	= 16	// peripherals: PWM0 channel A
	GPIO17	Pin	= 17	// peripherals: PWM0 channel B
	GPIO18	Pin	= 18	// peripherals: PWM1 channel A
	GPIO19	Pin	= 19	// peripherals: PWM1 channel B
	GPIO20	Pin	= 20	// peripherals: PWM2 channel A
	GPIO21	Pin	= 21	// peripherals: PWM2 channel B
	GPIO22	Pin	= 22	// peripherals: PWM3 channel A
	GPIO23	Pin	= 23	// peripherals: PWM3 channel B
	GPIO24	Pin	= 24	// peripherals: PWM4 channel A
	GPIO25	Pin	= 25	// peripherals: PWM4 channel B
	GPIO26	Pin	= 26	// peripherals: PWM5 channel A
	GPIO27	Pin	= 27	// peripherals: PWM5 channel B
	GPIO28	Pin	= 28	// peripherals: PWM6 channel A
	GPIO29	Pin	= 29	// peripherals: PWM6 channel B

	// Analog pins
	ADC0	Pin	= GPIO26
	ADC1	Pin	= GPIO27
	ADC2	Pin	= GPIO28
	ADC3	Pin	= GPIO29
)
```



```go
const RESETS_RESET_Msk = 0x01ffffff
```

RESETS_RESET_Msk is bitmask to reset all peripherals

TODO: This field is not available in the device file.


```go
const (
	// DPRAM : Endpoint control register
	usbEpControlEnable			= 0x80000000
	usbEpControlDoubleBuffered		= 0x40000000
	usbEpControlInterruptPerBuff		= 0x20000000
	usbEpControlInterruptPerDoubleBuff	= 0x10000000
	usbEpControlEndpointType		= 0x0c000000
	usbEpControlInterruptOnStall		= 0x00020000
	usbEpControlInterruptOnNak		= 0x00010000
	usbEpControlBufferAddress		= 0x0000ffff

	usbEpControlEndpointTypeControl		= 0x00000000
	usbEpControlEndpointTypeISO		= 0x04000000
	usbEpControlEndpointTypeBulk		= 0x08000000
	usbEpControlEndpointTypeInterrupt	= 0x0c000000

	// Endpoint buffer control bits
	usbBuf1CtrlFull		= 0x80000000
	usbBuf1CtrlLast		= 0x40000000
	usbBuf1CtrlData0Pid	= 0x20000000
	usbBuf1CtrlData1Pid	= 0x00000000
	usbBuf1CtrlSel		= 0x10000000
	usbBuf1CtrlStall	= 0x08000000
	usbBuf1CtrlAvail	= 0x04000000
	usbBuf1CtrlLenMask	= 0x03FF0000
	usbBuf0CtrlFull		= 0x00008000
	usbBuf0CtrlLast		= 0x00004000
	usbBuf0CtrlData0Pid	= 0x00000000
	usbBuf0CtrlData1Pid	= 0x00002000
	usbBuf0CtrlSel		= 0x00001000
	usbBuf0CtrlStall	= 0x00000800
	usbBuf0CtrlAvail	= 0x00000400
	usbBuf0CtrlLenMask	= 0x000003FF

	USBBufferLen	= 64
)
```



```go
const (
	LS_SE0	= 0b00
	LS_J	= 0b01
	LS_K	= 0b10
	LS_SE1	= 0b11
)
```



```go
const (
	// WatchdogMaxTimeout in milliseconds (approx 8.3s).
	//
	// Nominal 1us per watchdog tick, 24-bit counter,
	// but due to errata two ticks consumed per 1us.
	// See: Errata RP2040-E1
	WatchdogMaxTimeout = (rp.WATCHDOG_LOAD_LOAD_Msk / 1000) / 2
)
```



```go
const XOSC_STARTUP_DELAY_MULTIPLIER = 64
```

On some boards, the XOSC can take longer than usual to stabilize. On such
boards, this is needed to avoid a hard fault on boot/reset. Refer to
PICO_XOSC_STARTUP_DELAY_MULTIPLIER in the Pico SDK for additional details.


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
var Flash flashBlockDevice
```



```go
var (
	UART0	= &_UART0
	_UART0	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	rp.UART0,
	}

	UART1	= &_UART1
	_UART1	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	rp.UART1,
	}
)
```

UART on the RP2040


```go
var (
	I2C0	= &_I2C0
	_I2C0	= I2C{
		Bus: rp.I2C0,
	}
	I2C1	= &_I2C1
	_I2C1	= I2C{
		Bus: rp.I2C1,
	}
)
```

I2C on the RP2040.


```go
var (
	ErrInvalidI2CBaudrate	= errors.New("invalid i2c baudrate")
	ErrInvalidTgtAddr	= errors.New("invalid target i2c address not in 0..0x80 or is reserved")
	ErrI2CGeneric		= errors.New("i2c error")
	ErrRP2040I2CDisable	= errors.New("i2c rp2040 peripheral timeout in disable")
	errInvalidI2CSDA	= errors.New("invalid I2C SDA pin")
	errInvalidI2CSCL	= errors.New("invalid I2C SCL pin")
	ErrI2CAlreadyListening	= errors.New("i2c already listening")
	ErrI2CWrongMode		= errors.New("i2c wrong mode")
	ErrI2CUnderflow		= errors.New("i2c underflow")
)
```



```go
var (
	ErrBadPeriod = errors.New("period outside valid range 8ns..268ms")
)
```



```go
var (
	PWM0	= getPWMGroup(0)
	PWM1	= getPWMGroup(1)
	PWM2	= getPWMGroup(2)
	PWM3	= getPWMGroup(3)
	PWM4	= getPWMGroup(4)
	PWM5	= getPWMGroup(5)
	PWM6	= getPWMGroup(6)
	PWM7	= getPWMGroup(7)
)
```

Hardware Pulse Width Modulation (PWM) API
PWM peripherals available on RP2040. Each peripheral has 2 pins available for
a total of 16 available PWM outputs. Some pins may not be available on some boards.

The RP2040 PWM block has 8 identical slices. Each slice can drive two PWM output signals, or
measure the frequency or duty cycle of an input signal. This gives a total of up to 16 controllable
PWM outputs. All 30 GPIOs can be driven by the PWM block

The PWM hardware functions by continuously comparing the input value to a free-running counter. This produces a
toggling output where the amount of time spent at the high output level is proportional to the input value. The fraction of
time spent at the high signal level is known as the duty cycle of the signal.

The default behaviour of a PWM slice is to count upward until the wrap value (\ref pwm_config_set_wrap) is reached, and then
immediately wrap to 0. PWM slices also offer a phase-correct mode, where the counter starts to count downward after
reaching TOP, until it reaches 0 again.


```go
var RTC = (*rtcType)(unsafe.Pointer(rp.RTC))
```



```go
var (
	ErrRtcDelayTooSmall	= errors.New("RTC interrupt deplay is too small, shall be at least 1 second")
	ErrRtcDelayTooLarge	= errors.New("RTC interrupt deplay is too large, shall be no more than 1 day")
)
```



```go
var (
	SPI0	= &_SPI0
	_SPI0	= SPI{
		Bus: rp.SPI0,
	}
	SPI1	= &_SPI1
	_SPI1	= SPI{
		Bus: rp.SPI1,
	}
)
```

SPI on the RP2040


```go
var (
	ErrLSBNotSupported	= errors.New("SPI LSB unsupported on PL022")
	ErrSPITimeout		= errors.New("SPI timeout")
	ErrSPIBaud		= errors.New("SPI baud too low or above 66.5Mhz")
	errSPIInvalidSDI	= errors.New("invalid SPI SDI pin")
	errSPIInvalidSDO	= errors.New("invalid SPI SDO pin")
	errSPIInvalidSCK	= errors.New("invalid SPI SCK pin")
)
```



```go
var Watchdog = &watchdogImpl{}
```

Watchdog provides access to the hardware watchdog available
in the RP2040.


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



```go
var (
	USBDev	= &USBDevice{}
	USBCDC	Serialer
)
```



```go
var (
	ErrUSBReadTimeout	= errors.New("USB read timeout")
	ErrUSBBytesRead		= errors.New("USB invalid number of bytes read")
	ErrUSBBytesWritten	= errors.New("USB invalid number of bytes written")
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


### func ChipVersion

```go
func ChipVersion() uint8
```

ChipVersion returns the version of the chip. 1 is returned for B0 and B1
chip.


### func ConfigureUSBEndpoint

```go
func ConfigureUSBEndpoint(desc descriptor.Descriptor, epSettings []usb.EndpointConfig, setup []usb.SetupConfig)
```



### func CurrentCore

```go
func CurrentCore() int
```

CurrentCore returns the core number the call was made from.


### func DeviceID

```go
func DeviceID() []byte
```

DeviceID returns an identifier that is unique within
a particular chipset.

The identity is one burnt into the MCU itself, or the
flash chip at time of manufacture.

It's possible that two different vendors may allocate
the same DeviceID, so callers should take this into
account if needing to generate a globally unique id.

The length of the hardware ID is vendor-specific, but
8 bytes (64 bits) is common.


### func EnableCDC

```go
func EnableCDC(txHandler func(), rxHandler func([]byte), setupHandler func(usb.Setup) bool)
```



### func EnterBootloader

```go
func EnterBootloader()
```

EnterBootloader should perform a system reset in preparation
to switch to the bootloader to flash new firmware.


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

GetRNG returns 32 bits of semi-random data based on ring oscillator.

Unlike some other implementations of GetRNG, these random numbers are not
cryptographically secure and must not be used for cryptographic operations
(nonces, etc).


### func InitADC

```go
func InitADC()
```

InitADC resets the ADC peripheral.


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


### func NumCores

```go
func NumCores() int
```

NumCores returns number of cores available on the device.


### func PWMPeripheral

```go
func PWMPeripheral(pin Pin) (sliceNum uint8, err error)
```

Peripheral returns the RP2040 PWM peripheral which ranges from 0 to 7. Each
PWM peripheral has 2 channels, A and B which correspond to 0 and 1 in the program.
This number corresponds to the package's PWM0 throughout PWM7 handles


### func ReadTemperature

```go
func ReadTemperature() (millicelsius int32)
```

ReadTemperature does a one-shot sample of the internal temperature sensor and returns a milli-celsius reading.


### func ReceiveUSBControlPacket

```go
func ReceiveUSBControlPacket() ([cdcLineInfoSize]byte, error)
```



### func SendUSBInPacket

```go
func SendUSBInPacket(ep uint32, data []byte) bool
```

SendUSBInPacket sends a packet for USB (interrupt in / bulk in).


### func SendZlp

```go
func SendZlp()
```





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

Configure sets the ADC pin to analog input mode.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get returns a one-shot ADC sample reading.


### func (ADC) GetADCChannel

```go
func (a ADC) GetADCChannel() (c ADCChannel, err error)
```

GetADCChannel returns the channel associated with the ADC pin.




## type ADCChannel

```go
type ADCChannel uint8
```

ADCChannel is the ADC peripheral mux channel. 0-4.



### func (ADCChannel) Configure

```go
func (c ADCChannel) Configure(config ADCConfig) error
```

Configure sets the channel's associated pin to analog input mode.
The powered on temperature sensor increases ADC_AVDD current by approximately 40 μA.


### func (ADCChannel) Pin

```go
func (c ADCChannel) Pin() (p Pin, err error)
```

The Pin method returns the GPIO Pin associated with the ADC mux channel, if it has one.




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





## type I2C

```go
type I2C struct {
	Bus		*rp.I2C0_Type
	mode		I2CMode
	txInProgress	bool
}
```




### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```

Configure initializes i2c peripheral and configures I2C config's pins passed.
Here's a list of valid SDA and SCL GPIO pins on bus I2C0 of the rp2040:

	SDA: 0, 4, 8, 12, 16, 20
	SCL: 1, 5, 9, 13, 17, 21

Same as above for I2C1 bus:

	SDA: 2, 6, 10, 14, 18, 26
	SCL: 3, 7, 11, 15, 19, 27


### func (*I2C) Listen

```go
func (i2c *I2C) Listen(addr uint16) error
```

Listen starts listening for I2C requests sent to specified address

addr is the address to listen to


### func (*I2C) ReadRegister

```go
func (i2c *I2C) ReadRegister(address uint8, register uint8, data []byte) error
```

ReadRegister transmits the register, restarts the connection as a read
operation, and reads the response.

Many I2C-compatible devices are organized in terms of registers. This method
is a shortcut to easily read such registers. Also, it only works for devices
with 7-bit addresses, which is the vast majority.


### func (*I2C) Reply

```go
func (i2c *I2C) Reply(buf []byte) error
```



### func (*I2C) SetBaudRate

```go
func (i2c *I2C) SetBaudRate(br uint32) error
```

SetBaudRate sets the I2C frequency. It has the side effect of also
enabling the I2C hardware if disabled beforehand.


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
```

Tx performs a write and then a read transfer placing the result in
in r.

Passing a nil value for w or r skips the transfer corresponding to write
or read, respectively.

	i2c.Tx(addr, nil, r)

Performs only a read transfer.

	i2c.Tx(addr, w, nil)

Performs only a write transfer.


### func (*I2C) WaitForEvent

```go
func (i2c *I2C) WaitForEvent(buf []byte) (evt I2CTargetEvent, count int, err error)
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
	Frequency	uint32
	// SDA/SCL Serial Data and clock pins. Refer to datasheet to see
	// which pins match the desired bus.
	SDA, SCL	Pin
	Mode		I2CMode
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

Configure configures the gpio pin as per mode.


### func (Pin) Get

```go
func (p Pin) Get() bool
```

Get reads the pin value.


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



### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```



### func (Pin) Set

```go
func (p Pin) Set(value bool)
```

Set drives the pin high if value is true else drives it low.


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

PinChange represents one or more trigger events that can happen on a given GPIO pin
on the RP2040. ORed PinChanges are valid input to most IRQ functions.





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
	Bus *rp.SPI0_Type
}
```




### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig) error
```

Configure is intended to setup/initialize the SPI interface.
Default baudrate of 4MHz is used if Frequency == 0. Default
word length (data bits) is 8.
Below is a list of GPIO pins corresponding to SPI0 bus on the rp2040:

	SI : 0, 4, 17  a.k.a RX and MISO (if rp2040 is master)
	SO : 3, 7, 19  a.k.a TX and MOSI (if rp2040 is master)
	SCK: 2, 6, 18

SPI1 bus GPIO pins:

	SI : 8, 12
	SO : 11, 15
	SCK: 10, 14

No pin configuration is needed of SCK, SDO and SDI needed after calling Configure.


### func (SPI) GetBaudRate

```go
func (spi SPI) GetBaudRate() uint32
```



### func (SPI) PrintRegs

```go
func (spi SPI) PrintRegs()
```

PrintRegs prints SPI's peripheral common registries current values


### func (SPI) SetBaudRate

```go
func (spi SPI) SetBaudRate(br uint32) error
```



### func (SPI) Transfer

```go
func (spi SPI) Transfer(w byte) (byte, error)
```

Write a single byte and read a single byte from TX/RX FIFO.


### func (SPI) Tx

```go
func (spi SPI) Tx(w, r []byte) (err error)
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

Remark: This implementation (RP2040) allows reading into buffer with a custom repeated
value on tx.

	spi.Tx([]byte{0xff}, rx) // may cause unwanted heap allocations.

This form sends 0xff and puts the result into rx buffer. Useful for reading from SD cards
which require 0xff input on SI.




## type SPIConfig

```go
type SPIConfig struct {
	Frequency	uint32
	// LSB not supported on rp2040.
	LSBFirst	bool
	// Mode's two most LSB are CPOL and CPHA. i.e. Mode==2 (0b10) is CPOL=1, CPHA=0
	Mode	uint8
	// Serial clock pin
	SCK	Pin
	// TX or Serial Data Out (MOSI if rp2040 is master)
	SDO	Pin
	// RX or Serial Data In (MISO if rp2040 is master)
	SDI	Pin
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
	Buffer		*RingBuffer
	Bus		*rp.UART0_Type
	Interrupt	interrupt.Interrupt
}
```

UART on the RP2040.



### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig) error
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

SetBaudRate sets the baudrate to be used for the UART.


### func (*UART) SetFormat

```go
func (uart *UART) SetFormat(databits, stopbits uint8, parity UARTParity) error
```

SetFormat for number of data bits, stop bits, and parity for the UART.


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





## type USBBuffer

```go
type USBBuffer struct {
	Buffer0	[USBBufferLen]byte
	Buffer1	[USBBufferLen]byte
}
```






## type USBBufferControlRegister

```go
type USBBufferControlRegister struct {
	In	volatile.Register32
	Out	volatile.Register32
}
```






## type USBDPSRAM

```go
type USBDPSRAM struct {
	// Note that EPxControl[0] is not EP0Control but 8-byte setup data.
	EPxControl	[16]USBEndpointControlRegister

	EPxBufferControl	[16]USBBufferControlRegister

	EPxBuffer	[16]USBBuffer
}
```






## type USBDevice

```go
type USBDevice struct {
	initcomplete		bool
	InitEndpointComplete	bool
}
```




### func (*USBDevice) Configure

```go
func (dev *USBDevice) Configure(config UARTConfig)
```

Configure the USB peripheral. The config is here for compatibility with the UART interface.




## type USBEndpointControlRegister

```go
type USBEndpointControlRegister struct {
	In	volatile.Register32
	Out	volatile.Register32
}
```






## type WatchdogConfig

```go
type WatchdogConfig struct {
	// The timeout (in milliseconds) before the watchdog fires.
	//
	// If the requested timeout exceeds `MaxTimeout` it will be rounded
	// down.
	TimeoutMillis uint32
}
```

WatchdogConfig holds configuration for the watchdog timer.





