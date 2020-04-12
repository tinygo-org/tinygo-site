
---
title: pybadge
---


## Constants

```go
const RESET_MAGIC_VALUE = 0xf01669ef
```

used to reset into bootloader


```go
const (
	D0	= PB17	// UART0 RX/PWM available
	D1	= PB16	// UART0 TX/PWM available
	D2	= PB03
	D3	= PB02
	D4	= PA14	// PWM available
	D5	= PA16	// PWM available
	D6	= PA18	// PWM available
	D7	= PB14
	D8	= PA15	// built-in neopixel
	D9	= PA19	// PWM available
	D10	= PA20	// can be used for PWM or UART1 TX
	D11	= PA21	// can be used for PWM or UART1 RX
	D12	= PA22	// PWM available
	D13	= PA23	// PWM available
)
```

GPIO Pins


```go
const (
	A0	= PA02	// ADC/AIN[0]
	A1	= PA05	// ADC/AIN[2]
	A2	= PB08	// ADC/AIN[3]
	A3	= PB09	// ADC/AIN[4]
	A4	= PA04	// ADC/AIN[5]
	A5	= PA06	// ADC/AIN[6]
	A6	= PB01	// ADC/AIN[12]/VMEAS
	A7	= PB04	// ADC/AIN[6]/LIGHT
	A8	= D2	// ADC/AIN[14]
	A9	= D3	// ADC/AIN[15]
)
```

Analog pins


```go
const (
	LED		= D13
	NEOPIXELS	= D8

	LIGHTSENSOR	= A7

	BUTTON_LATCH	= PB00
	BUTTON_OUT	= PB30
	BUTTON_CLK	= PB31

	TFT_DC		= PB05
	TFT_CS		= PB07
	TFT_RST		= PA00
	TFT_LITE	= PA01

	SPEAKER_ENABLE	= PA27
)
```



```go
const (
	BUTTON_LEFT_MASK	= 1
	BUTTON_UP_MASK		= 2
	BUTTON_DOWN_MASK	= 4
	BUTTON_RIGHT_MASK	= 8
	BUTTON_SELECT_MASK	= 16
	BUTTON_START_MASK	= 32
	BUTTON_A_MASK		= 64
	BUTTON_B_MASK		= 128
)
```



```go
const (
	USBCDC_DM_PIN	= PA24
	USBCDC_DP_PIN	= PA25
)
```

UART0 aka USBCDC pins


```go
const (
	UART_TX_PIN	= D1
	UART_RX_PIN	= D0
)
```

UART1 pins


```go
const (
	UART2_TX_PIN	= A4
	UART2_RX_PIN	= A5
)
```

UART2 pins


```go
const (
	SDA_PIN	= PA12	// SDA: SERCOM2/PAD[0]
	SCL_PIN	= PA13	// SCL: SERCOM2/PAD[1]
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	= PA17	// SCK: SERCOM1/PAD[1]
	SPI0_MOSI_PIN	= PB23	// MOSI: SERCOM1/PAD[3]
	SPI0_MISO_PIN	= PB22	// MISO: SERCOM1/PAD[2]
)
```

SPI pins


```go
const (
	SPI1_SCK_PIN	= PB13	// SCK: SERCOM4/PAD[1]
	SPI1_MOSI_PIN	= PB15	// MOSI: SERCOM4/PAD[3]
	SPI1_MISO_PIN	= NoPin
)
```

TFT SPI pins


```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.


```go
const (
	I2SModeMaster	I2SMode	= iota
	I2SModeSlave
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
const NoPin = Pin(-1)
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
	PinPWME			PinMode	= PinTimer
	PinPWMF			PinMode	= PinTimerAlt
	PinPWMG			PinMode	= PinTCCPDEC
	PinInputPulldown	PinMode	= 18
)
```



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
)
```

Hardware pins


```go
const (
	// SERCOM_FREQ_REF is always reference frequency on SAMD51 regardless of CPU speed.
	SERCOM_FREQ_REF	= 48000000

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
const HSRAM_SIZE = 0x00030000
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





## Variables

```go
var (
	I2C0 = I2C{
		Bus:	sam.SERCOM2_I2CM,
		SERCOM:	2,
	}
)
```

I2C on the ItsyBitsy M4.


```go
var (
	SPI0 = SPI{
		Bus:	sam.SERCOM1_SPIM,
		SERCOM:	1,
	}
)
```

SPI on the PyBadge.


```go
var (
	SPI1 = SPI{
		Bus:	sam.SERCOM4_SPIM,
		SERCOM:	4,
	}
)
```

TFT SPI on the PyBadge.


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
	// UART0 is actually a USB CDC interface.
	UART0	= USBCDC{Buffer: NewRingBuffer()}

	// The first hardware serial port on the SAMD51. Uses the SERCOM3 interface.
	UART1	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	sam.SERCOM3_USART_INT,
		SERCOM:	3,
	}

	// The second hardware serial port on the SAMD51. Uses the SERCOM0 interface.
	UART2	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	sam.SERCOM0_USART_INT,
		SERCOM:	0,
	}
)
```



```go
var (
	ErrTxInvalidSliceSize = errors.New("SPI write and read slices must be same size")
)
```






### func CPUFrequency

```go
func CPUFrequency() uint32
```



### func InitADC

```go
func InitADC()
```

InitADC initializes the ADC.


### func InitPWM

```go
func InitPWM()
```

InitPWM initializes the PWM interface.


### func NewACMFunctionalDescriptor

```go
func NewACMFunctionalDescriptor(subtype, d0 uint8) ACMFunctionalDescriptor
```

NewACMFunctionalDescriptor returns a new USB ACMFunctionalDescriptor.


### func NewCDCCSInterfaceDescriptor

```go
func NewCDCCSInterfaceDescriptor(subtype, d0, d1 uint8) CDCCSInterfaceDescriptor
```

NewCDCCSInterfaceDescriptor returns a new USB CDCCSInterfaceDescriptor.


### func NewCDCDescriptor

```go
func NewCDCDescriptor(i IADDescriptor, c InterfaceDescriptor,
	h CDCCSInterfaceDescriptor,
	cm ACMFunctionalDescriptor,
	fd CDCCSInterfaceDescriptor,
	callm CMFunctionalDescriptor,
	ci EndpointDescriptor,
	di InterfaceDescriptor,
	outp EndpointDescriptor,
	inp EndpointDescriptor) CDCDescriptor
```



### func NewCMFunctionalDescriptor

```go
func NewCMFunctionalDescriptor(subtype, d0, d1 uint8) CMFunctionalDescriptor
```

NewCMFunctionalDescriptor returns a new USB CMFunctionalDescriptor.


### func NewConfigDescriptor

```go
func NewConfigDescriptor(totalLength uint16, interfaces uint8) ConfigDescriptor
```

NewConfigDescriptor returns a new USB ConfigDescriptor.


### func NewDeviceDescriptor

```go
func NewDeviceDescriptor(class, subClass, proto, packetSize0 uint8, vid, pid, version uint16, im, ip, is, configs uint8) DeviceDescriptor
```

NewDeviceDescriptor returns a USB DeviceDescriptor.


### func NewEndpointDescriptor

```go
func NewEndpointDescriptor(addr, attr uint8, packetSize uint16, interval uint8) EndpointDescriptor
```

NewEndpointDescriptor returns a new USB EndpointDescriptor.


### func NewIADDescriptor

```go
func NewIADDescriptor(firstInterface, count, class, subClass, protocol uint8) IADDescriptor
```

NewIADDescriptor returns a new USB IADDescriptor.


### func NewInterfaceDescriptor

```go
func NewInterfaceDescriptor(n, numEndpoints, class, subClass, protocol uint8) InterfaceDescriptor
```

NewInterfaceDescriptor returns a new USB InterfaceDescriptor.


### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


### func ResetProcessor

```go
func ResetProcessor()
```

ResetProcessor should perform a system reset in preparation
to switch to the bootloader to flash new firmware.




## type ACMFunctionalDescriptor

```go
type ACMFunctionalDescriptor struct {
	len		uint8
	dtype		uint8	// 0x24
	subtype		uint8	// 1
	bmCapabilities	uint8
}
```

ACMFunctionalDescriptor is a Abstract Control Model (ACM) USB descriptor.



### func (ACMFunctionalDescriptor) Bytes

```go
func (d ACMFunctionalDescriptor) Bytes() []byte
```

Bytes returns the ACMFunctionalDescriptor data.




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

Get returns the current value of a ADC pin, in the range 0..0xffff.




## type CDCCSInterfaceDescriptor

```go
type CDCCSInterfaceDescriptor struct {
	len	uint8	// 5
	dtype	uint8	// 0x24
	subtype	uint8
	d0	uint8
	d1	uint8
}
```

CDCCSInterfaceDescriptor is a CDC CS interface descriptor.



### func (CDCCSInterfaceDescriptor) Bytes

```go
func (d CDCCSInterfaceDescriptor) Bytes() []byte
```

Bytes returns CDCCSInterfaceDescriptor data.




## type CDCDescriptor

```go
type CDCDescriptor struct {
	//	IAD
	iad	IADDescriptor	// Only needed on compound device

	//	Control
	cif	InterfaceDescriptor
	header	CDCCSInterfaceDescriptor

	// CDC control
	controlManagement	ACMFunctionalDescriptor		// ACM
	functionalDescriptor	CDCCSInterfaceDescriptor	// CDC_UNION
	callManagement		CMFunctionalDescriptor		// Call Management
	cifin			EndpointDescriptor

	//	CDC Data
	dif	InterfaceDescriptor
	in	EndpointDescriptor
	out	EndpointDescriptor
}
```

CDCDescriptor is the Communication Device Class (CDC) descriptor.



### func (CDCDescriptor) Bytes

```go
func (d CDCDescriptor) Bytes() []byte
```

Bytes returns CDCDescriptor data.




## type CMFunctionalDescriptor

```go
type CMFunctionalDescriptor struct {
	bFunctionLength		uint8
	bDescriptorType		uint8	// 0x24
	bDescriptorSubtype	uint8	// 1
	bmCapabilities		uint8
	bDataInterface		uint8
}
```

CMFunctionalDescriptor is the functional descriptor general format.



### func (CMFunctionalDescriptor) Bytes

```go
func (d CMFunctionalDescriptor) Bytes() []byte
```

Bytes returns the CMFunctionalDescriptor data.




## type ConfigDescriptor

```go
type ConfigDescriptor struct {
	bLength			uint8	// 9
	bDescriptorType		uint8	// 2
	wTotalLength		uint16	// total length
	bNumInterfaces		uint8
	bConfigurationValue	uint8
	iConfiguration		uint8
	bmAttributes		uint8
	bMaxPower		uint8
}
```

ConfigDescriptor implements the standard USB configuration descriptor.

Table 9-10. Standard Configuration Descriptor
bLength, bDescriptorType, wTotalLength, bNumInterfaces, bConfigurationValue, iConfiguration
bmAttributes, bMaxPower



### func (ConfigDescriptor) Bytes

```go
func (d ConfigDescriptor) Bytes() []byte
```

Bytes returns ConfigDescriptor data.




## type DeviceDescriptor

```go
type DeviceDescriptor struct {
	bLength			uint8	// 18
	bDescriptorType		uint8	// 1 USB_DEVICE_DESCRIPTOR_TYPE
	bcdUSB			uint16	// 0x200
	bDeviceClass		uint8
	bDeviceSubClass		uint8
	bDeviceProtocol		uint8
	bMaxPacketSize0		uint8	// Packet 0
	idVendor		uint16
	idProduct		uint16
	bcdDevice		uint16	// 0x100
	iManufacturer		uint8
	iProduct		uint8
	iSerialNumber		uint8
	bNumConfigurations	uint8
}
```

DeviceDescriptor implements the USB standard device descriptor.

Table 9-8. Standard Device Descriptor
bLength, bDescriptorType, bcdUSB, bDeviceClass, bDeviceSubClass, bDeviceProtocol, bMaxPacketSize0,
   idVendor, idProduct, bcdDevice, iManufacturer, iProduct, iSerialNumber, bNumConfigurations */



### func (DeviceDescriptor) Bytes

```go
func (d DeviceDescriptor) Bytes() []byte
```

Bytes returns DeviceDescriptor data




## type EndpointDescriptor

```go
type EndpointDescriptor struct {
	bLength			uint8	// 7
	bDescriptorType		uint8	// 5
	bEndpointAddress	uint8
	bmAttributes		uint8
	wMaxPacketSize		uint16
	bInterval		uint8
}
```

EndpointDescriptor implements the standard USB endpoint descriptor.

Table 9-13. Standard Endpoint Descriptor
bLength, bDescriptorType, bEndpointAddress, bmAttributes, wMaxPacketSize, bInterval



### func (EndpointDescriptor) Bytes

```go
func (d EndpointDescriptor) Bytes() []byte
```

Bytes returns EndpointDescriptor data.




## type I2C

```go
type I2C struct {
	Bus	*sam.SERCOM_I2CM_Type
	SERCOM	uint8
}
```

I2C on the SAMD51.



### func (I2C) Configure

```go
func (i2c I2C) Configure(config I2CConfig) error
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


### func (I2C) SetBaudRate

```go
func (i2c I2C) SetBaudRate(br uint32)
```

SetBaudRate sets the communication speed for the I2C.


### func (I2C) Tx

```go
func (i2c I2C) Tx(addr uint16, w, r []byte) error
```

Tx does a single I2C transaction at the specified address.
It clocks out the given address, writes the bytes in w, reads back len(r)
bytes and stores them in r, and generates a stop condition on the bus.


### func (I2C) WriteByte

```go
func (i2c I2C) WriteByte(data byte) error
```

WriteByte writes a single byte to the I2C bus.


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
	SCK			Pin
	WS			Pin
	SD			Pin
	Mode			I2SMode
	Standard		I2SStandard
	ClockSource		I2SClockSource
	DataFormat		I2SDataFormat
	AudioFrequency		uint32
	MasterClockOutput	bool
	Stereo			bool
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






## type IADDescriptor

```go
type IADDescriptor struct {
	bLength			uint8	// 8
	bDescriptorType		uint8	// 11
	bFirstInterface		uint8
	bInterfaceCount		uint8
	bFunctionClass		uint8
	bFunctionSubClass	uint8
	bFunctionProtocol	uint8
	iFunction		uint8
}
```

IADDescriptor is an Interface Association Descriptor, which is used
to bind 2 interfaces together in CDC composite device.

Standard Interface Association Descriptor:
bLength, bDescriptorType, bFirstInterface, bInterfaceCount, bFunctionClass, bFunctionSubClass,
bFunctionProtocol, iFunction



### func (IADDescriptor) Bytes

```go
func (d IADDescriptor) Bytes() []byte
```

Bytes returns IADDescriptor data.




## type InterfaceDescriptor

```go
type InterfaceDescriptor struct {
	bLength			uint8	// 9
	bDescriptorType		uint8	// 4
	bInterfaceNumber	uint8
	bAlternateSetting	uint8
	bNumEndpoints		uint8
	bInterfaceClass		uint8
	bInterfaceSubClass	uint8
	bInterfaceProtocol	uint8
	iInterface		uint8
}
```

InterfaceDescriptor implements the standard USB interface descriptor.

Table 9-12. Standard Interface Descriptor
bLength, bDescriptorType, bInterfaceNumber, bAlternateSetting, bNumEndpoints, bInterfaceClass,
bInterfaceSubClass, bInterfaceProtocol, iInterface



### func (InterfaceDescriptor) Bytes

```go
func (d InterfaceDescriptor) Bytes() []byte
```

Bytes returns InterfaceDescriptor data.




## type MSCDescriptor

```go
type MSCDescriptor struct {
	msc	InterfaceDescriptor
	in	EndpointDescriptor
	out	EndpointDescriptor
}
```

MSCDescriptor is not used yet.





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

Set turns on the duty cycle for a PWM pin using the provided value.




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


### func (Pin) Toggle

```go
func (p Pin) Toggle()
```

Toggle switches an output pin from low to high or from high to low.
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
	MOSI		Pin
	MISO		Pin
	LSBFirst	bool
	Mode		uint8
}
```

SPIConfig is used to store config info for SPI.





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



### func (UART) Buffered

```go
func (uart UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (UART) Configure

```go
func (uart UART) Configure(config UARTConfig) error
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


### func (UART) SetBaudRate

```go
func (uart UART) SetBaudRate(br uint32)
```

SetBaudRate sets the communication speed for the UART.


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






## type USBCDC

```go
type USBCDC struct {
	Buffer *RingBuffer
}
```

USBCDC is the USB CDC aka serial over USB interface on the SAMD21.



### func (USBCDC) Buffered

```go
func (usbcdc USBCDC) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (USBCDC) Configure

```go
func (usbcdc USBCDC) Configure(config UARTConfig)
```

Configure the USB CDC interface. The config is here for compatibility with the UART interface.


### func (USBCDC) DTR

```go
func (usbcdc USBCDC) DTR() bool
```



### func (USBCDC) RTS

```go
func (usbcdc USBCDC) RTS() bool
```



### func (USBCDC) Read

```go
func (usbcdc USBCDC) Read(data []byte) (n int, err error)
```

Read from the RX buffer.


### func (USBCDC) ReadByte

```go
func (usbcdc USBCDC) ReadByte() (byte, error)
```

ReadByte reads a single byte from the RX buffer.
If there is no data in the buffer, returns an error.


### func (USBCDC) Receive

```go
func (usbcdc USBCDC) Receive(data byte)
```

Receive handles adding data to the UART's data buffer.
Usually called by the IRQ handler for a machine.


### func (USBCDC) Write

```go
func (usbcdc USBCDC) Write(data []byte) (n int, err error)
```

Write data to the USBCDC.


### func (USBCDC) WriteByte

```go
func (usbcdc USBCDC) WriteByte(c byte) error
```

WriteByte writes a byte of data to the USB CDC interface.




