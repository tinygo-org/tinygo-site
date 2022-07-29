
---
title: pyportal
---


## Constants

```go
const RESET_MAGIC_VALUE = 0xf01669ef
```

used to reset into bootloader


```go
const (
	D0	= PB13	// NINA_RX
	D1	= PB12	// NINA_TX
	D2	= PB22	// built-in neopixel
	D3	= PA04	// PWM available
	D4	= PA05	// PWM available
	D5	= PB16	// NINA_ACK
	D6	= PB15	// NINA_GPIO0
	D7	= PB17	// NINA_RESETN
	D8	= PB14	// NINA_CS
	D9	= PB04	// TFT_RD
	D10	= PB05	// TFT_DC
	D11	= PB06	// TFT_CS
	D12	= PB07	// TFT_TE
	D13	= PB23	// built-in LED
	D24	= PA00	// TFT_RESET
	D25	= PB31	// TFT_BACKLIGHT
	D26	= PB09	// TFT_WR
	D27	= PB02	// SDA
	D28	= PB03	// SCL
	D29	= PA12	// SDO
	D30	= PA13	// SCK
	D31	= PA14	// SDI
	D32	= PB30	// SD_CS
	D33	= PA01	// SD_CARD_DETECT
	D34	= PA16	// LCD_DATA0
	D35	= PA17	// LCD_DATA1
	D36	= PA18	// LCD_DATA2
	D37	= PA19	// LCD_DATA3
	D38	= PA20	// LCD_DATA4
	D39	= PA21	// LCD_DATA5
	D40	= PA22	// LCD_DATA6
	D41	= PA23	// LCD_DATA7
	D42	= PB10	// QSPI
	D43	= PB11	// QSPI
	D44	= PA08	// QSPI
	D45	= PA09	// QSPI
	D46	= PA10	// QSPI
	D47	= PA11	// QSPI
	D50	= PA02	// speaker amplifier shutdown
	D51	= PA15	// NINA_RTS

	NINA_CS		= D8
	NINA_ACK	= D5
	NINA_GPIO0	= D6
	NINA_RESETN	= D7

	NINA_TX		= D1
	NINA_RX		= D0
	NINA_RTS	= D51

	LCD_DATA0	= D34

	TFT_RD		= D9
	TFT_DC		= D10
	TFT_CS		= D11
	TFT_TE		= D12
	TFT_RESET	= D24
	TFT_BACKLIGHT	= D25
	TFT_WR		= D26

	NEOPIXEL	= D2
	WS2812		= D2
	SPK_SD		= D50
)
```

GPIO Pins


```go
const (
	A0	= PA02	// ADC0/AIN[0]
	A1	= D3	// ADC0/AIN[4]
	A2	= PA07	// ADC0/AIN[7]
	A3	= D4	// ADC0/AIN[5]
	A4	= PB00	// ADC0/AIN[12]
	A5	= PB01	// ADC0/AIN[13]
	A6	= PA06	// ADC0/AIN[6]
	A7	= PB08	// ADC1/AIN[0]

	AUDIO_OUT	= A0
	LIGHT		= A2
	TOUCH_YD	= A4
	TOUCH_XL	= A5
	TOUCH_YU	= A6
	TOUCH_XR	= A7
)
```

Analog pins


```go
const (
	LED = D13
)
```



```go
const (
	USBCDC_DM_PIN	= PA24
	USBCDC_DP_PIN	= PA25
)
```

USBCDC pins


```go
const (
	UART_TX_PIN	= D1
	UART_RX_PIN	= D0
)
```

UART1 aka NINA_TX/NINA_RX


```go
const (
	SDA_PIN	= PB02	// SDA: SERCOM2/PAD[0]
	SCL_PIN	= PB03	// SCL: SERCOM2/PAD[1]
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	= PA13	// SCK: SERCOM1/PAD[1]
	SPI0_SDO_PIN	= PA12	// SDO: SERCOM1/PAD[3]
	SPI0_SDI_PIN	= PA14	// SDI: SERCOM1/PAD[2]

	NINA_SDO	= SPI0_SDO_PIN
	NINA_SDI	= SPI0_SDI_PIN
	NINA_SCK	= SPI0_SCK_PIN
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


```go
const (
	I2SModeSource	I2SMode	= iota
	I2SModeReceiver
	I2SModePDM
)
```



```go
const (
	I2StandardPhilips	I2SStandard	= iota
	I2SStandardMSB
	I2SStandardLSB
)
```



```go
const (
	I2SClockSourceInternal	I2SClockSource	= iota
	I2SClockSourceExternal
)
```



```go
const (
	I2SDataFormatDefault	I2SDataFormat	= 0
	I2SDataFormat8bit			= 8
	I2SDataFormat16bit			= 16
	I2SDataFormat24bit			= 24
	I2SDataFormat32bit			= 32
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
const NoPin = Pin(0xff)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	PinAnalog		PinMode	= 1
	PinSERCOM		PinMode	= 2
	PinSERCOMAlt		PinMode	= 3
	PinTimer		PinMode	= 4
	PinTimerAlt		PinMode	= 5
	PinTCCPDEC		PinMode	= 6
	PinCom			PinMode	= 7
	PinSDHC			PinMode	= 8
	PinI2S			PinMode	= 9
	PinPCC			PinMode	= 10
	PinGMAC			PinMode	= 11
	PinACCLK		PinMode	= 12
	PinCCL			PinMode	= 13
	PinDigital		PinMode	= 14
	PinInput		PinMode	= 15
	PinInputPullup		PinMode	= 16
	PinOutput		PinMode	= 17
	PinTCCE			PinMode	= PinTimer
	PinTCCF			PinMode	= PinTimerAlt
	PinTCCG			PinMode	= PinTCCPDEC
	PinInputPulldown	PinMode	= 18
	PinCAN			PinMode	= 19
	PinCAN0			PinMode	= PinSDHC
	PinCAN1			PinMode	= PinCom
)
```



```go
const (
	PinRising	PinChange	= sam.EIC_CONFIG_SENSE0_RISE
	PinFalling	PinChange	= sam.EIC_CONFIG_SENSE0_FALL
	PinToggle	PinChange	= sam.EIC_CONFIG_SENSE0_BOTH
)
```

Pin change interrupt constants for SetInterrupt.


```go
const (
	PA00	Pin	= 0
	PA01	Pin	= 1
	PA02	Pin	= 2
	PA03	Pin	= 3
	PA04	Pin	= 4
	PA05	Pin	= 5
	PA06	Pin	= 6
	PA07	Pin	= 7
	PA08	Pin	= 8
	PA09	Pin	= 9
	PA10	Pin	= 10
	PA11	Pin	= 11
	PA12	Pin	= 12
	PA13	Pin	= 13
	PA14	Pin	= 14
	PA15	Pin	= 15
	PA16	Pin	= 16
	PA17	Pin	= 17
	PA18	Pin	= 18
	PA19	Pin	= 19
	PA20	Pin	= 20
	PA21	Pin	= 21
	PA22	Pin	= 22
	PA23	Pin	= 23
	PA24	Pin	= 24
	PA25	Pin	= 25
	PA26	Pin	= 26
	PA27	Pin	= 27
	PA28	Pin	= 28
	PA29	Pin	= 29
	PA30	Pin	= 30
	PA31	Pin	= 31
	PB00	Pin	= 32
	PB01	Pin	= 33
	PB02	Pin	= 34
	PB03	Pin	= 35
	PB04	Pin	= 36
	PB05	Pin	= 37
	PB06	Pin	= 38
	PB07	Pin	= 39
	PB08	Pin	= 40
	PB09	Pin	= 41
	PB10	Pin	= 42
	PB11	Pin	= 43
	PB12	Pin	= 44
	PB13	Pin	= 45
	PB14	Pin	= 46
	PB15	Pin	= 47
	PB16	Pin	= 48
	PB17	Pin	= 49
	PB18	Pin	= 50
	PB19	Pin	= 51
	PB20	Pin	= 52
	PB21	Pin	= 53
	PB22	Pin	= 54
	PB23	Pin	= 55
	PB24	Pin	= 56
	PB25	Pin	= 57
	PB26	Pin	= 58
	PB27	Pin	= 59
	PB28	Pin	= 60
	PB29	Pin	= 61
	PB30	Pin	= 62
	PB31	Pin	= 63
	PC00	Pin	= 64
	PC01	Pin	= 65
	PC02	Pin	= 66
	PC03	Pin	= 67
	PC04	Pin	= 68
	PC05	Pin	= 69
	PC06	Pin	= 70
	PC07	Pin	= 71
	PC08	Pin	= 72
	PC09	Pin	= 73
	PC10	Pin	= 74
	PC11	Pin	= 75
	PC12	Pin	= 76
	PC13	Pin	= 77
	PC14	Pin	= 78
	PC15	Pin	= 79
	PC16	Pin	= 80
	PC17	Pin	= 81
	PC18	Pin	= 82
	PC19	Pin	= 83
	PC20	Pin	= 84
	PC21	Pin	= 85
	PC22	Pin	= 86
	PC23	Pin	= 87
	PC24	Pin	= 88
	PC25	Pin	= 89
	PC26	Pin	= 90
	PC27	Pin	= 91
	PC28	Pin	= 92
	PC29	Pin	= 93
	PC30	Pin	= 94
	PC31	Pin	= 95
	PD00	Pin	= 96
	PD01	Pin	= 97
	PD02	Pin	= 98
	PD03	Pin	= 99
	PD04	Pin	= 100
	PD05	Pin	= 101
	PD06	Pin	= 102
	PD07	Pin	= 103
	PD08	Pin	= 104
	PD09	Pin	= 105
	PD10	Pin	= 106
	PD11	Pin	= 107
	PD12	Pin	= 108
	PD13	Pin	= 109
	PD14	Pin	= 110
	PD15	Pin	= 111
	PD16	Pin	= 112
	PD17	Pin	= 113
	PD18	Pin	= 114
	PD19	Pin	= 115
	PD20	Pin	= 116
	PD21	Pin	= 117
	PD22	Pin	= 118
	PD23	Pin	= 119
	PD24	Pin	= 120
	PD25	Pin	= 121
	PD26	Pin	= 122
	PD27	Pin	= 123
	PD28	Pin	= 124
	PD29	Pin	= 125
	PD30	Pin	= 126
	PD31	Pin	= 127
)
```

Hardware pins


```go
const (
	// SERCOM_FREQ_REF is always reference frequency on SAMD51 regardless of CPU speed.
	SERCOM_FREQ_REF		= 48000000
	SERCOM_FREQ_REF_GCLK0	= 120000000

	// Default rise time in nanoseconds, based on 4.7K ohm pull up resistors
	riseTimeNanoseconds	= 125

	// wire bus states
	wireUnknownState	= 0
	wireIdleState		= 1
	wireOwnerState		= 2
	wireBusyState		= 3

	// wire commands
	wireCmdNoAction		= 0
	wireCmdRepeatStart	= 1
	wireCmdRead		= 2
	wireCmdStop		= 3
)
```



```go
const (
	QSPI_SCK	= PB10
	QSPI_CS		= PB11
	QSPI_DATA0	= PA08
	QSPI_DATA1	= PA09
	QSPI_DATA2	= PA10
	QSPI_DATA3	= PA11
)
```

The QSPI peripheral on ATSAMD51 is only available on the following pins


```go
const HSRAM_SIZE = 0x00040000
```



```go
const (
	// ParityNone means to not use any parity checking. This is
	// the most common setting.
	ParityNone	UARTParity	= 0

	// ParityEven means to expect that the total number of 1 bits sent
	// should be an even number.
	ParityEven	UARTParity	= 1

	// ParityOdd means to expect that the total number of 1 bits sent
	// should be an odd number.
	ParityOdd	UARTParity	= 2
)
```






## Variables

```go
var (
	UART1	= &sercomUSART4

	DefaultUART	= UART1
)
```



```go
var (
	I2C0 = sercomI2CM5
)
```

I2C on the PyPortal.


```go
var (
	SPI0		= sercomSPIM2
	NINA_SPI	= SPI0
)
```

SPI on the PyPortal.


```go
var (
	ErrTimeoutRNG		= errors.New("machine: RNG Timeout")
	ErrInvalidInputPin	= errors.New("machine: invalid input pin")
	ErrInvalidOutputPin	= errors.New("machine: invalid output pin")
	ErrInvalidClockPin	= errors.New("machine: invalid clock pin")
	ErrInvalidDataPin	= errors.New("machine: invalid data pin")
	ErrNoPinChangeChannel	= errors.New("machine: no channel available for pin interrupt")
)
```



```go
var (
	ErrTxInvalidSliceSize = errors.New("SPI write and read slices must be same size")
)
```



```go
var (
	DAC0	= DAC{Channel: 0}
	DAC1	= DAC{Channel: 1}
)
```



```go
var (
	TCC0	= (*TCC)(sam.TCC0)
	TCC1	= (*TCC)(sam.TCC1)
	TCC2	= (*TCC)(sam.TCC2)
	TCC3	= (*TCC)(sam.TCC3)
	TCC4	= (*TCC)(sam.TCC4)
)
```

This chip has five TCC peripherals, which have PWM as one feature.


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
	USBDev	= &USBDevice{}
	USBCDC	Serialer
)
```



```go
var (
	ErrUSBReadTimeout	= errors.New("USB read timeout")
	ErrUSBBytesRead		= errors.New("USB invalid number of bytes read")
)
```






### func CPUFrequency

```go
func CPUFrequency() uint32
```



### func EnableCDC

```go
func EnableCDC(txHandler func(), rxHandler func([]byte), setupHandler func(usb.Setup) bool)
```



### func EnableHID

```go
func EnableHID(txHandler func(), rxHandler func([]byte), setupHandler func(usb.Setup) bool)
```

EnableHID enables HID. This function must be executed from the init().


### func EnableMIDI

```go
func EnableMIDI(txHandler func(), rxHandler func([]byte), setupHandler func(usb.Setup) bool)
```

EnableMIDI enables MIDI. This function must be executed from the init().


### func EnterBootloader

```go
func EnterBootloader()
```

EnterBootloader should perform a system reset in preparation
to switch to the bootloader to flash new firmware.


### func GetRNG

```go
func GetRNG() (uint32, error)
```

GetRNG returns 32 bits of cryptographically secure random data


### func InitADC

```go
func InitADC()
```

InitADC initializes the ADC.


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


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
func (a ADC) Configure(config ADCConfig)
```

Configure configures a ADCPin to be able to be used to read data.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get returns the current value of a ADC pin, in the range 0..0xffff.




## type ADCConfig

```go
type ADCConfig struct {
	Reference	uint32	// analog reference voltage (AREF) in millivolts
	Resolution	uint32	// number of bits for a single conversion (e.g., 8, 10, 12)
	Samples		uint32	// number of samples for a single conversion (e.g., 4, 8, 16, 32)
}
```

ADCConfig holds ADC configuration parameters. If left unspecified, the zero
value of each parameter will use the peripheral's default settings.





## type DAC

```go
type DAC struct {
	Channel uint8
}
```

DAC on the SAMD51.



### func (DAC) Configure

```go
func (dac DAC) Configure(config DACConfig)
```

Configure the DAC.
output pin must already be configured.


### func (DAC) Set

```go
func (dac DAC) Set(value uint16) error
```

Set writes a single 16-bit value to the DAC.
Since the ATSAMD51 only has a 12-bit DAC, the passed-in value will be scaled down.




## type DACConfig

```go
type DACConfig struct {
}
```

DACConfig placeholder for future expansion.





## type I2C

```go
type I2C struct {
	Bus	*sam.SERCOM_I2CM_Type
	SERCOM	uint8
}
```

I2C on the SAMD51.



### func (*I2C) Configure

```go
func (i2c *I2C) Configure(config I2CConfig) error
```

Configure is intended to setup the I2C interface.


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
func (i2c *I2C) SetBaudRate(br uint32)
```

SetBaudRate sets the communication speed for the I2C.


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
```

Tx does a single I2C transaction at the specified address.
It clocks out the given address, writes the bytes in w, reads back len(r)
bytes and stores them in r, and generates a stop condition on the bus.


### func (*I2C) WriteByte

```go
func (i2c *I2C) WriteByte(data byte) error
```

WriteByte writes a single byte to the I2C bus.


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
	SCL		Pin
	SDA		Pin
}
```

I2CConfig is used to store config info for I2C.





## type I2SClockSource

```go
type I2SClockSource uint8
```






## type I2SConfig

```go
type I2SConfig struct {
	SCK		Pin
	WS		Pin
	SD		Pin
	Mode		I2SMode
	Standard	I2SStandard
	ClockSource	I2SClockSource
	DataFormat	I2SDataFormat
	AudioFrequency	uint32
	MainClockOutput	bool
	Stereo		bool
}
```

All fields are optional and may not be required or used on a particular platform.





## type I2SDataFormat

```go
type I2SDataFormat uint8
```






## type I2SMode

```go
type I2SMode uint8
```






## type I2SStandard

```go
type I2SStandard uint8
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

Return the register and mask to disable a given port. This can be used to
implement bit-banged drivers.


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```

Return the register and mask to enable a given GPIO pin. This can be used to
implement bit-banged drivers.


### func (Pin) Set

```go
func (p Pin) Set(high bool)
```

Set the pin to high or low.
Warning: only use this on an output pin!


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


### func (Pin) Toggle

```go
func (p Pin) Toggle()
```

Toggle switches an output pin from low to high or from high to low.
Warning: only use this on an output pin!




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
	Bus	*sam.SERCOM_SPIM_Type
	SERCOM	uint8
}
```

SPI



### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig) error
```

Configure is intended to setup the SPI interface.


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






## type TCC

```go
type TCC sam.TCC_Type
```

TCC is one timer peripheral, which consists of a counter and multiple output
channels (that can be connected to actual pins). You can set the frequency
using SetPeriod, but only for all the channels in this timer peripheral at
once.



### func (*TCC) Channel

```go
func (tcc *TCC) Channel(pin Pin) (uint8, error)
```

Channel returns a PWM channel for the given pin. Note that one channel may be
shared between multiple pins, and so will have the same duty cycle. If this
is not desirable, look for a different TCC or consider using a different pin.


### func (*TCC) Configure

```go
func (tcc *TCC) Configure(config PWMConfig) error
```

Configure enables and configures this TCC.


### func (*TCC) Counter

```go
func (tcc *TCC) Counter() uint32
```

Counter returns the current counter value of the timer in this TCC
peripheral. It may be useful for debugging.


### func (*TCC) Set

```go
func (tcc *TCC) Set(channel uint8, value uint32)
```

Set updates the channel value. This is used to control the channel duty
cycle, in other words the fraction of time the channel output is high (or low
when inverted). For example, to set it to a 25% duty cycle, use:

    tcc.Set(channel, tcc.Top() / 4)

tcc.Set(channel, 0) will set the output to low and tcc.Set(channel,
tcc.Top()) will set the output to high, assuming the output isn't inverted.


### func (*TCC) SetInverting

```go
func (tcc *TCC) SetInverting(channel uint8, inverting bool)
```

SetInverting sets whether to invert the output of this channel.
Without inverting, a 25% duty cycle would mean the output is high for 25% of
the time and low for the rest. Inverting flips the output as if a NOT gate
was placed at the output, meaning that the output would be 25% low and 75%
high with a duty cycle of 25%.


### func (*TCC) SetPeriod

```go
func (tcc *TCC) SetPeriod(period uint64) error
```

SetPeriod updates the period of this TCC peripheral.
To set a particular frequency, use the following formula:

    period = 1e9 / frequency

If you use a period of 0, a period that works well for LEDs will be picked.

SetPeriod will not change the prescaler, but also won't change the current
value in any of the channels. This means that you may need to update the
value for the particular channel.

Note that you cannot pick any arbitrary period after the TCC peripheral has
been configured. If you want to switch between frequencies, pick the lowest
frequency (longest period) once when calling Configure and adjust the
frequency here as needed.


### func (*TCC) Top

```go
func (tcc *TCC) Top() uint32
```

Top returns the current counter top, for use in duty cycle calculation. It
will only change with a call to Configure or SetPeriod, otherwise it is
constant.

The value returned here is hardware dependent. In general, it's best to treat
it as an opaque value that can be divided by some number and passed to
tcc.Set (see tcc.Set for more information).




## type UART

```go
type UART struct {
	Buffer		*RingBuffer
	Bus		*sam.SERCOM_USART_INT_Type
	SERCOM		uint8
	Interrupt	interrupt.Interrupt	// RXC interrupt
}
```

UART on the SAMD51.



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

SetBaudRate sets the communication speed for the UART.


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
type UARTParity int
```

UARTParity is the parity setting to be used for UART communication.





## type USBDevice

```go
type USBDevice struct {
	initcomplete bool
}
```




### func (*USBDevice) Configure

```go
func (dev *USBDevice) Configure(config UARTConfig)
```

Configure the USB peripheral. The config is here for compatibility with the UART interface.




