
---
title: particle-argon
---


## Constants

```go
const HasLowFrequencyCrystal = true
```



```go
const (
	LED		Pin	= 44
	LED_GREEN	Pin	= 14
	LED_RED		Pin	= 13
	LED_BLUE	Pin	= 15
)
```

LEDs


```go
const (
	A0	Pin	= 3
	A1	Pin	= 4
	A2	Pin	= 28
	A3	Pin	= 29
	A4	Pin	= 30
	A5	Pin	= 31
	D0	Pin	= 26	// Also SDA
	D1	Pin	= 27	// Also SCL
	D2	Pin	= 33
	D3	Pin	= 34
	D4	Pin	= 40
	D5	Pin	= 42
	D6	Pin	= 43
	D7	Pin	= 44	// Also LED
	D8	Pin	= 35
	D9	Pin	= 6	// Also TX
	D10	Pin	= 8	// Also RX
	D11	Pin	= 46	// Also MISO
	D12	Pin	= 45	// Also MOSI
	D13	Pin	= 47	// Also SCK
)
```

GPIOs


```go
const (
	UART_TX_PIN	Pin	= 6
	UART_RX_PIN	Pin	= 8
)
```



```go
const (
	SDA_PIN	Pin	= 26
	SCL_PIN	Pin	= 27
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	Pin	= 47
	SPI0_MOSI_PIN	Pin	= 45
	SPI0_MISO_PIN	Pin	= 46
)
```

SPI pins


```go
const (
	SPI1_SCK_PIN	Pin	= 19
	SPI1_MOSI_PIN	Pin	= 20
	SPI1_MISO_PIN	Pin	= 21
	SPI1_CS_PIN	Pin	= 17
	SPI1_WP_PIN	Pin	= 22
	SPI1_HOLD_PIN	Pin	= 23
)
```

Internal 4MB SPI Flash


```go
const (
	ESP32_TXD_PIN		Pin	= 36
	ESP32_RXD_PIN		Pin	= 37
	ESP32_CTS_PIN		Pin	= 39
	ESP32_RTS_PIN		Pin	= 38
	ESP32_BOOT_MODE_PIN	Pin	= 16
	ESP32_WIFI_EN_PIN	Pin	= 24
	ESP32_HOST_WK_PIN	Pin	= 7
)
```

ESP32 coprocessor


```go
const (
	MODE_BUTTON_PIN		Pin	= 11
	CHARGE_STATUS_PIN	Pin	= 41
	LIPO_VOLTAGE_PIN	Pin	= 5
	PCB_ANTENNA_PIN		Pin	= 2
	EXTERNAL_UFL_PIN	Pin	= 25
	NFC1_PIN		Pin	= 9
	NFC2_PIN		Pin	= 10
)
```

Other peripherals


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
	PinInput		PinMode	= (nrf.GPIO_PIN_CNF_DIR_Input << nrf.GPIO_PIN_CNF_DIR_Pos) | (nrf.GPIO_PIN_CNF_INPUT_Connect << nrf.GPIO_PIN_CNF_INPUT_Pos)
	PinInputPullup		PinMode	= PinInput | (nrf.GPIO_PIN_CNF_PULL_Pullup << nrf.GPIO_PIN_CNF_PULL_Pos)
	PinInputPulldown	PinMode	= PinInput | (nrf.GPIO_PIN_CNF_PULL_Pulldown << nrf.GPIO_PIN_CNF_PULL_Pos)
	PinOutput		PinMode	= (nrf.GPIO_PIN_CNF_DIR_Output << nrf.GPIO_PIN_CNF_DIR_Pos) | (nrf.GPIO_PIN_CNF_INPUT_Disconnect << nrf.GPIO_PIN_CNF_INPUT_Pos)
)
```



```go
const (
	P0_00	Pin	= 0
	P0_01	Pin	= 1
	P0_02	Pin	= 2
	P0_03	Pin	= 3
	P0_04	Pin	= 4
	P0_05	Pin	= 5
	P0_06	Pin	= 6
	P0_07	Pin	= 7
	P0_08	Pin	= 8
	P0_09	Pin	= 9
	P0_10	Pin	= 10
	P0_11	Pin	= 11
	P0_12	Pin	= 12
	P0_13	Pin	= 13
	P0_14	Pin	= 14
	P0_15	Pin	= 15
	P0_16	Pin	= 16
	P0_17	Pin	= 17
	P0_18	Pin	= 18
	P0_19	Pin	= 19
	P0_20	Pin	= 20
	P0_21	Pin	= 21
	P0_22	Pin	= 22
	P0_23	Pin	= 23
	P0_24	Pin	= 24
	P0_25	Pin	= 25
	P0_26	Pin	= 26
	P0_27	Pin	= 27
	P0_28	Pin	= 28
	P0_29	Pin	= 29
	P0_30	Pin	= 30
	P0_31	Pin	= 31
	P1_00	Pin	= 32
	P1_01	Pin	= 33
	P1_02	Pin	= 34
	P1_03	Pin	= 35
	P1_04	Pin	= 36
	P1_05	Pin	= 37
	P1_06	Pin	= 38
	P1_07	Pin	= 39
	P1_08	Pin	= 40
	P1_09	Pin	= 41
	P1_10	Pin	= 42
	P1_11	Pin	= 43
	P1_12	Pin	= 44
	P1_13	Pin	= 45
	P1_14	Pin	= 46
	P1_15	Pin	= 47
)
```

Hardware pins





## Variables

```go
var (
	Serial	= USB
	UART0	= NRF_UART0
)
```

UART


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
	ErrTxInvalidSliceSize = errors.New("SPI write and read slices must be same size")
)
```



```go
var (
	// NRF_UART0 is the hardware UART on the NRF SoC.
	NRF_UART0 = UART{Buffer: NewRingBuffer()}
)
```

UART


```go
var (
	I2C0	= I2C{Bus: nrf.TWI0}
	I2C1	= I2C{Bus: nrf.TWI1}
)
```

There are 2 I2C interfaces on the NRF.


```go
var (
	SPI0	= SPI{Bus: nrf.SPI0}
	SPI1	= SPI{Bus: nrf.SPI1}
)
```

There are 2 SPI interfaces on the NRF5x.


```go
var (
	USB	= USBCDC{Buffer: NewRingBuffer()}

	usbEndpointDescriptors	[8]usbDeviceDescriptor

	udd_ep_in_cache_buffer	[7][128]uint8
	udd_ep_out_cache_buffer	[7][128]uint8

	sendOnEP0DATADONE	struct {
		ptr	*byte
		count	int
	}
	isEndpointHalt		= false
	isRemoteWakeUpEnabled	= false
	endPoints		= []uint32{usb_ENDPOINT_TYPE_CONTROL,
		(usb_ENDPOINT_TYPE_INTERRUPT | usbEndpointIn),
		(usb_ENDPOINT_TYPE_BULK | usbEndpointOut),
		(usb_ENDPOINT_TYPE_BULK | usbEndpointIn)}

	usbConfiguration		uint8
	usbSetInterface			uint8
	usbLineInfo			= cdcLineInfo{115200, 0x00, 0x00, 0x08, 0x00}
	epinen				uint32
	epouten				uint32
	easyDMABusy			volatile.Register8
	epout0data_setlinecoding	bool
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

InitADC initializes the registers needed for ADC.


### func InitPWM

```go
func InitPWM()
```

InitPWM initializes the registers needed for PWM.


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

Configure configures an ADC pin to be able to read analog data.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get returns the current value of a ADC pin in the range 0..0xffff.




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
	Bus *nrf.TWI_Type
}
```

I2C on the NRF.



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
	Frequency	uint32
	SCL		Pin
	SDA		Pin
}
```

I2CConfig is used to store config info for I2C.





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
	Bus *nrf.SPI_Type
}
```

SPI on the NRF.



### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig)
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
	Buffer *RingBuffer
}
```

UART on the NRF.



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
	Buffer		*RingBuffer
	interrupt	interrupt.Interrupt
}
```

USBCDC is the USB CDC aka serial over USB interface on the nRF52840



### func (USBCDC) Buffered

```go
func (usbcdc USBCDC) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*USBCDC) Configure

```go
func (usbcdc *USBCDC) Configure(config UARTConfig)
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




