
---
title: arduino-nano33
---


## Constants

```go
const RESET_MAGIC_VALUE = 0x07738135
```

used to reset into bootloader


```go
const (
	RX0	Pin	= PB23	// UART2 RX
	TX1	Pin	= PB22	// UART2 TX

	D2	Pin	= PB10	// PWM available
	D3	Pin	= PB11	// PWM available
	D4	Pin	= PA07
	D5	Pin	= PA05	// PWM available
	D6	Pin	= PA04	// PWM available
	D7	Pin	= PA06

	D8	Pin	= PA18
	D9	Pin	= PA20	// PWM available
	D10	Pin	= PA21	// PWM available
	D11	Pin	= PA16	// PWM available
	D12	Pin	= PA19	// PWM available

	D13	Pin	= PA17
)
```

GPIO Pins


```go
const (
	A0	Pin	= PA02	// ADC/AIN[0]
	A1	Pin	= PB02	// ADC/AIN[10]
	A2	Pin	= PA11	// ADC/AIN[19]
	A3	Pin	= PA10	// ADC/AIN[18],
	A4	Pin	= PB08	// ADC/AIN[2], SCL:  SERCOM2/PAD[1]
	A5	Pin	= PB09	// ADC/AIN[3], SDA:  SERCOM2/PAD[1]
	A6	Pin	= PA09	// ADC/AIN[17]
	A7	Pin	= PB03	// ADC/AIN[11]
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
	USBCDC_DM_PIN	Pin	= PA24
	USBCDC_DP_PIN	Pin	= PA25
)
```

USBCDC pins


```go
const (
	UART_TX_PIN	Pin	= PA22
	UART_RX_PIN	Pin	= PA23
)
```

UART1 pins


```go
const (
	SDA_PIN	Pin	= A4	// SDA: SERCOM4/PAD[1]
	SCL_PIN	Pin	= A5	// SCL: SERCOM4/PAD[1]
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	Pin	= D13	// SCK: SERCOM1/PAD[1]
	SPI0_SDO_PIN	Pin	= D11	// SDO: SERCOM1/PAD[0]
	SPI0_SDI_PIN	Pin	= D12	// SDI: SERCOM1/PAD[3]
)
```

SPI pins


```go
const (
	NINA_SDO	Pin	= PA12
	NINA_SDI	Pin	= PA13
	NINA_CS		Pin	= PA14
	NINA_SCK	Pin	= PA15
	NINA_GPIO0	Pin	= PA27
	NINA_RESETN	Pin	= PA08
	NINA_ACK	Pin	= PA28
	NINA_TX		Pin	= PA22
	NINA_RX		Pin	= PA23
)
```

NINA-W102 Pins


```go
const (
	I2S_SCK_PIN	Pin	= PA10
	I2S_SD_PIN	Pin	= PA08
	I2S_WS_PIN		= NoPin	// TODO: figure out what this is on Arduino Nano 33.
)
```

I2S pins


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
	PinAnalog	PinMode	= 1
	PinSERCOM	PinMode	= 2
	PinSERCOMAlt	PinMode	= 3
	PinTimer	PinMode	= 4
	PinTimerAlt	PinMode	= 5
	PinCom		PinMode	= 6
	//PinAC_CLK        PinMode = 7
	PinDigital		PinMode	= 8
	PinInput		PinMode	= 9
	PinInputPullup		PinMode	= 10
	PinOutput		PinMode	= 11
	PinTCC			PinMode	= PinTimer
	PinTCCAlt		PinMode	= PinTimerAlt
	PinInputPulldown	PinMode	= 12
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
var UART1 = &sercomUSART3
```

UART1 on the Arduino Nano 33 connects to the onboard NINA-W102 WiFi chip.


```go
var UART2 = &sercomUSART5
```

UART2 on the Arduino Nano 33 connects to the normal TX/RX pins.


```go
var (
	I2C0 = sercomI2CM4
)
```

I2C on the Arduino Nano 33.


```go
var SPI0 = sercomSPIM1
```

SPI on the Arduino Nano 33.


```go
var (
	SPI1		= sercomSPIM2
	NINA_SPI	= SPI1
)
```

SPI1 is connected to the NINA-W102 chip on the Arduino Nano 33.


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
var I2S0 = I2S{Bus: sam.I2S}
```



```go
var (
	ErrTxInvalidSliceSize = errors.New("SPI write and read slices must be same size")
)
```



```go
var (
	TCC0	= (*TCC)(sam.TCC0)
	TCC1	= (*TCC)(sam.TCC1)
	TCC2	= (*TCC)(sam.TCC2)
)
```

The SAM D21 has three TCC peripherals, which have PWM as one feature.


```go
var (
	USB = &USBCDC{Buffer: NewRingBuffer()}
)
```



```go
var (
	DAC0 = DAC{}
)
```



```go
var (
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = USB
```

Serial is implemented via USB (USB-CDC).





### func CPUFrequency

```go
func CPUFrequency() uint32
```

Return the current CPU frequency in hertz.


### func InitADC

```go
func InitADC()
```

InitADC initializes the ADC.


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

ResetProcessor should perform a system reset in preperation
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
func (d ACMFunctionalDescriptor) Bytes() [acmFunctionalDescriptorSize]byte
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
func (a ADC) Configure(config ADCConfig)
```

Configure configures a ADC pin to be able to be used to read data.


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
func (d CDCCSInterfaceDescriptor) Bytes() [cdcCSInterfaceDescriptorSize]byte
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
func (d CDCDescriptor) Bytes() [cdcSize]byte
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
func (d CMFunctionalDescriptor) Bytes() [cmFunctionalDescriptorSize]byte
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
func (d ConfigDescriptor) Bytes() [configDescriptorSize]byte
```

Bytes returns ConfigDescriptor data.




## type DAC

```go
type DAC struct {
}
```

DAC on the SAMD21.



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
Since the ATSAMD21 only has a 10-bit DAC, the passed-in value will be scaled down.




## type DACConfig

```go
type DACConfig struct {
}
```

DACConfig placeholder for future expansion.





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
func (d DeviceDescriptor) Bytes() [deviceDescriptorSize]byte
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
func (d EndpointDescriptor) Bytes() [endpointDescriptorSize]byte
```

Bytes returns EndpointDescriptor data.




## type I2C

```go
type I2C struct {
	Bus	*sam.SERCOM_I2CM_Type
	SERCOM	uint8
}
```

I2C on the SAMD21.



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





## type I2S

```go
type I2S struct {
	Bus *sam.I2S_Type
}
```

I2S



### func (I2S) Close

```go
func (i2s I2S) Close() error
```

Close the I2S bus.


### func (I2S) Configure

```go
func (i2s I2S) Configure(config I2SConfig)
```

Configure is used to configure the I2S interface. You must call this
before you can use the I2S bus.


### func (I2S) Read

```go
func (i2s I2S) Read(p []uint32) (n int, err error)
```

Read data from the I2S bus into the provided slice.
The I2S bus must already have been configured correctly.


### func (I2S) Write

```go
func (i2s I2S) Write(p []uint32) (n int, err error)
```

Write data to the I2S bus from the provided slice.
The I2S bus must already have been configured correctly.




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
func (d IADDescriptor) Bytes() [iadDescriptorSize]byte
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
func (d InterfaceDescriptor) Bytes() [interfaceDescriptorSize]byte
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
	Bus	*sam.SERCOM_SPI_Type
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





## type TCC

```go
type TCC sam.TCC_Type
```

TCC is one timer/counter peripheral, which consists of a counter and multiple
output channels (that can be connected to actual pins). You can set the
frequency using SetPeriod, but only for all the channels in this TCC
peripheral at once.



### func (*TCC) Channel

```go
func (tcc *TCC) Channel(pin Pin) (uint8, error)
```

Channel returns a PWM channel for the given pin. Note that one channel may be
shared between multiple pins, and so will have the same duty cycle. If this
is not desirable, look for a different TCC peripheral or consider using a
different pin.


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
it as an opaque value that can be divided by some number and passed to Set
(see Set documentation for more information).




## type UART

```go
type UART struct {
	Buffer		*RingBuffer
	Bus		*sam.SERCOM_USART_Type
	SERCOM		uint8
	Interrupt	interrupt.Interrupt
}
```

UART on the SAMD21.



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





## type USBCDC

```go
type USBCDC struct {
	Buffer			*RingBuffer
	TxIdx			volatile.Register8
	waitTxc			bool
	waitTxcRetryCount	uint8
	sent			bool
	configured		bool
}
```

USBCDC is the USB CDC aka serial over USB interface on the SAMD21.



### func (*USBCDC) Buffered

```go
func (usbcdc *USBCDC) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*USBCDC) Configure

```go
func (usbcdc *USBCDC) Configure(config UARTConfig)
```

Configure the USB CDC interface. The config is here for compatibility with the UART interface.


### func (*USBCDC) Configured

```go
func (usbcdc *USBCDC) Configured() bool
```

Configured returns whether usbcdc is configured or not.


### func (*USBCDC) DTR

```go
func (usbcdc *USBCDC) DTR() bool
```



### func (*USBCDC) Flush

```go
func (usbcdc *USBCDC) Flush() error
```

Flush flushes buffered data.


### func (*USBCDC) RTS

```go
func (usbcdc *USBCDC) RTS() bool
```



### func (*USBCDC) Read

```go
func (usbcdc *USBCDC) Read(data []byte) (n int, err error)
```

Read from the RX buffer.


### func (*USBCDC) ReadByte

```go
func (usbcdc *USBCDC) ReadByte() (byte, error)
```

ReadByte reads a single byte from the RX buffer.
If there is no data in the buffer, returns an error.


### func (*USBCDC) Receive

```go
func (usbcdc *USBCDC) Receive(data byte)
```

Receive handles adding data to the UART's data buffer.
Usually called by the IRQ handler for a machine.


### func (*USBCDC) Write

```go
func (usbcdc *USBCDC) Write(data []byte) (n int, err error)
```

Write data to the USBCDC.


### func (*USBCDC) WriteByte

```go
func (usbcdc *USBCDC) WriteByte(c byte) error
```

WriteByte writes a byte of data to the USB CDC interface.




