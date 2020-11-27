
---
title: wioterminal
---


## Constants

```go
const RESET_MAGIC_VALUE = 0xf01669ef
```

used to reset into bootloader


```go
const (
	ADC0	= A0
	ADC1	= A1
	ADC2	= A2
	ADC3	= A3
	ADC4	= A4
	ADC5	= A5
	ADC6	= A6
	ADC7	= A7
	ADC8	= A8

	LED	= PIN_LED
	BUTTON	= BUTTON_1
)
```



```go
const (

	// LEDs
	PIN_LED_13	= PA15
	PIN_LED_RXL	= PA15
	PIN_LED_TXL	= PA15
	PIN_LED		= PIN_LED_13
	PIN_LED2	= PIN_LED_RXL
	PIN_LED3	= PIN_LED_TXL
	LED_BUILTIN	= PIN_LED_13
	PIN_NEOPIXEL	= PA15

	//Digital PINs
	D0	= PB08
	D1	= PB09
	D2	= PA07
	D3	= PB04
	D4	= PB05
	D5	= PB06
	D6	= PA04
	D7	= PB07
	D8	= PA06

	//Analog PINs
	A0	= PB08	// ADC/AIN[0]
	A1	= PB09	// ADC/AIN[2]
	A2	= PA07	// ADC/AIN[3]
	A3	= PB04	// ADC/AIN[4]
	A4	= PB05	// ADC/AIN[5]
	A5	= PB06	// ADC/AIN[10]
	A6	= PA04	// ADC/AIN[10]
	A7	= PB07	// ADC/AIN[10]
	A8	= PA06	// ADC/AIN[10]

	//PIN DEFINE FOR RPI
	BCM0	= PA13	// I2C Wire1
	BCM1	= PA12	// I2C Wire1
	BCM2	= PA17	// I2C Wire2
	BCM3	= PA16	// I2C Wire2
	BCM4	= PB14	// GCLK
	BCM5	= PB12	// GCLK
	BCM6	= PB13	// GCLK
	BCM7	= PA05	// DAC1
	BCM8	= PB01	// SPI SS
	BCM9	= PB00	// SPI SDI
	BCM10	= PB02	// SPI SDO
	BCM11	= PB03	// SPI SCK
	BCM12	= PB06
	BCM13	= PA07
	BCM14	= PB27	// UART Serial1
	BCM15	= PB26	// UART Serial1
	BCM16	= PB07
	BCM17	= PA02	// DAC0
	BCM18	= PB28	// FPC Digital & AD pins
	BCM19	= PA20	// WIO_IR
	BCM20	= PA21	// I2S SDO
	BCM21	= PA22	// I2S SDI
	BCM22	= PB09
	BCM23	= PA07
	BCM24	= PB04
	BCM25	= PB05
	BCM26	= PA06
	BCM27	= PB08

	// FPC NEW DEFINE
	FPC1	= PB28	// FPC Digital & AD pins
	FPC2	= PB17
	FPC3	= PB29
	FPC4	= PA14
	FPC5	= PC01
	FPC6	= PC02
	FPC7	= PC03
	FPC8	= PC04
	FPC9	= PC31
	FPC10	= PD00

	// RPI Analog RPIs
	RPI_A0	= PB08
	RPI_A1	= PB09
	RPI_A2	= PA07
	RPI_A3	= PB04
	RPI_A4	= PB05
	RPI_A5	= PB06
	RPI_A6	= PA04
	RPI_A7	= PB07
	RPI_A8	= PA06

	PIN_DAC0	= PA02
	PIN_DAC1	= PA05

	// USB
	PIN_USB_DM		= PA24
	PIN_USB_DP		= PA25
	PIN_USB_HOST_ENABLE	= PA27

	// BUTTON
	BUTTON_1	= PC26
	BUTTON_2	= PC27
	BUTTON_3	= PC28
	WIO_KEY_A	= PC26
	WIO_KEY_B	= PC27
	WIO_KEY_C	= PC28

	// SWITCH
	SWITCH_X	= PD20
	SWITCH_Y	= PD12
	SWITCH_Z	= PD09
	SWITCH_B	= PD08
	SWITCH_U	= PD10

	WIO_5S_UP	= PD20
	WIO_5S_LEFT	= PD12
	WIO_5S_RIGHT	= PD09
	WIO_5S_DOWN	= PD08
	WIO_5S_PRESS	= PD10

	// IRQ0 : RTL8720D
	IRQ0	= PC20

	// BUZZER_CTR
	BUZZER_CTR	= PD11
	WIO_BUZZER	= PD11

	// MIC_INPUT
	MIC_INPUT	= PC30
	WIO_MIC		= PC30

	// GCLK
	GCLK0	= PB14
	GCLK1	= PB12
	GCLK2	= PB13

	// Serial interfaces
	// Serial1
	PIN_SERIAL1_RX	= PB27
	PIN_SERIAL1_TX	= PB26

	// Serial2 : RTL8720D
	PIN_SERIAL2_RX	= PC23
	PIN_SERIAL2_TX	= PC22

	// Wire Interfaces
	// I2C Wire2
	// I2C1
	PIN_WIRE_SDA	= PA17
	PIN_WIRE_SCL	= PA16
	SDA		= PIN_WIRE_SDA
	SCL		= PIN_WIRE_SCL

	// I2C Wire1
	// I2C0 : LIS3DHTR and ATECC608
	PIN_WIRE1_SDA	= PA13
	PIN_WIRE1_SCL	= PA12

	SDA1	= PIN_WIRE1_SDA
	SCL1	= PIN_WIRE1_SCL

	PIN_GYROSCOPE_WIRE_SDA	= PIN_WIRE1_SDA
	PIN_GYROSCOPE_WIRE_SCL	= PIN_WIRE1_SCL
	GYROSCOPE_INT1		= PC21

	WIO_LIS3DH_SDA	= PIN_WIRE1_SDA
	WIO_LIS3DH_SCL	= PIN_WIRE1_SCL
	WIO_LIS3DH_INT	= PC21

	// SPI
	PIN_SPI_SDI	= PB00
	PIN_SPI_SDO	= PB02
	PIN_SPI_SCK	= PB03
	PIN_SPI_SS	= PB01

	SS	= PIN_SPI_SS
	SDO	= PIN_SPI_SDO
	SDI	= PIN_SPI_SDI
	SCK	= PIN_SPI_SCK

	// SPI1 RTL8720D_SPI
	PIN_SPI1_SDI	= PC24
	PIN_SPI1_SDO	= PB24
	PIN_SPI1_SCK	= PB25
	PIN_SPI1_SS	= PC25

	SS1	= PIN_SPI1_SS
	SDO1	= PIN_SPI1_SDO
	SDI1	= PIN_SPI1_SDI
	SCK1	= PIN_SPI1_SCK

	// SPI2 SD_SPI
	PIN_SPI2_SDI	= PC18
	PIN_SPI2_SDO	= PC16
	PIN_SPI2_SCK	= PC17
	PIN_SPI2_SS	= PC19

	SS2	= PIN_SPI2_SS
	SDO2	= PIN_SPI2_SDO
	SDI2	= PIN_SPI2_SDI
	SCK2	= PIN_SPI2_SCK

	// SPI3 LCD_SPI
	PIN_SPI3_SDI	= PB18
	PIN_SPI3_SDO	= PB19
	PIN_SPI3_SCK	= PB20
	PIN_SPI3_SS	= PB21

	SS3	= PIN_SPI3_SS
	SDO3	= PIN_SPI3_SDO
	SDI3	= PIN_SPI3_SDI
	SCK3	= PIN_SPI3_SCK

	// Needed for SD library
	SDCARD_SDI_PIN	= PIN_SPI2_SDI
	SDCARD_SDO_PIN	= PIN_SPI2_SDO
	SDCARD_SCK_PIN	= PIN_SPI2_SCK
	SDCARD_SS_PIN	= PIN_SPI2_SS
	SDCARD_DET_PIN	= PD21

	LCD_SDI_PIN	= PIN_SPI3_SDI
	LCD_SDO_PIN	= PIN_SPI3_SDO
	LCD_SCK_PIN	= PIN_SPI3_SCK
	LCD_SS_PIN	= PIN_SPI3_SS
	LCD_DC		= PC06
	LCD_RESET	= PC07
	LCD_BACKLIGHT	= PC05

	// 4 WIRE LCD TOUCH
	LCD_XL	= PC10
	LCD_YU	= PC11
	LCD_XR	= PC12
	LCD_YD	= PC13

	// Needed for RTL8720D
	RTL8720D_SDI_PIN	= PIN_SPI1_SDI
	RTL8720D_SDO_PIN	= PIN_SPI1_SDO
	RTL8720D_SCK_PIN	= PIN_SPI1_SCK
	RTL8720D_SS_PIN		= PIN_SPI1_SS

	//QSPI Pins
	PIN_QSPI_IO0	= PA08
	PIN_QSPI_IO1	= PA09
	PIN_QSPI_IO2	= PA10
	PIN_QSPI_IO3	= PA11
	PIN_QSPI_SCK	= PB10
	PIN_QSPI_CS	= PB11

	// I2S Interfaces
	PIN_I2S_FS	= PA20
	PIN_I2S_SCK	= PB16
	PIN_I2S_SDO	= PA22
	PIN_I2S_SDI	= PA21

	I2S_LRCLK	= PA20
	I2S_BLCK	= PB16
	I2S_SDOUT	= PA22
	I2S_SDIN	= PA21

	// RTL8720D Interfaces
	RTL8720D_CHIP_PU	= PA18
	RTL8720D_GPIO0		= PA19	// SYNC

	// SWD
	SWDCLK	= PA30
	SWDIO	= PA31
	SWO	= PB30

	// light sensor
	WIO_LIGHT	= PD01

	// ir sensor
	WIO_IR	= PB31

	// OUTPUT_CTR
	OUTPUT_CTR_5V	= PC14
	OUTPUT_CTR_3V3	= PC15
)
```



```go
const (
	USBCDC_DM_PIN	= PIN_USB_DM
	USBCDC_DP_PIN	= PIN_USB_DP
)
```

UART0 aka USBCDC pins


```go
const (
	UART_TX_PIN	= PIN_SERIAL1_TX
	UART_RX_PIN	= PIN_SERIAL1_RX
)
```

UART1 pins


```go
const (
	UART2_TX_PIN	= PIN_SERIAL2_TX
	UART2_RX_PIN	= PIN_SERIAL2_RX
)
```

UART2 pins RTL8720D


```go
const (
	SDA0_PIN	= PIN_WIRE_SDA	// SDA: SERCOM3/PAD[0]
	SCL0_PIN	= PIN_WIRE_SCL	// SCL: SERCOM3/PAD[1]

	SDA1_PIN	= PIN_WIRE1_SDA	// SDA: SERCOM4/PAD[0]
	SCL1_PIN	= PIN_WIRE1_SCL	// SCL: SERCOM4/PAD[1]

	SDA_PIN	= SDA0_PIN
	SCL_PIN	= SCL0_PIN
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	= SCK	// SCK:  SERCOM5/PAD[1]
	SPI0_SDO_PIN	= SDO	// SDO: SERCOM5/PAD[0]
	SPI0_SDI_PIN	= SDI	// SDI: SERCOM5/PAD[2]

	// RTL8720D
	SPI1_SCK_PIN	= SCK1	// SCK:  SERCOM0/PAD[1]
	SPI1_SDO_PIN	= SDO1	// SDO: SERCOM0/PAD[0]
	SPI1_SDI_PIN	= SDI1	// SDI: SERCOM0/PAD[2]

	// SD
	SPI2_SCK_PIN	= SCK2	// SCK:  SERCOM6/PAD[1]
	SPI2_SDO_PIN	= SDO2	// SDO: SERCOM6/PAD[0]
	SPI2_SDI_PIN	= SDI2	// SDI: SERCOM6/PAD[2]

	// LCD
	SPI3_SCK_PIN	= SCK3	// SCK:  SERCOM7/PAD[1]
	SPI3_SDO_PIN	= SDO3	// SDO: SERCOM7/PAD[3]
	SPI3_SDI_PIN	= SDI3	// SDI: SERCOM7/PAD[2]
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
	PinPWME			PinMode	= PinTimer
	PinPWMF			PinMode	= PinTimerAlt
	PinPWMG			PinMode	= PinTCCPDEC
	PinInputPulldown	PinMode	= 18
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
	UART1	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	sam.SERCOM2_USART_INT,
		SERCOM:	2,
	}

	// RTL8720D
	UART2	= UART{
		Buffer:	NewRingBuffer(),
		Bus:	sam.SERCOM1_USART_INT,
		SERCOM:	1,
	}
)
```



```go
var (
	I2C0	= I2C{
		Bus:	sam.SERCOM4_I2CM,
		SERCOM:	4,
	}

	I2C1	= I2C{
		Bus:	sam.SERCOM4_I2CM,
		SERCOM:	4,
	}
)
```

I2C on the Wio Terminal


```go
var (
	SPI0	= SPI{
		Bus:	sam.SERCOM5_SPIM,
		SERCOM:	5,
	}

	// RTL8720D
	SPI1	= SPI{
		Bus:	sam.SERCOM0_SPIM,
		SERCOM:	0,
	}

	// SD
	SPI2	= SPI{
		Bus:	sam.SERCOM6_SPIM,
		SERCOM:	6,
	}

	// LCD
	SPI3	= SPI{
		Bus:	sam.SERCOM7_SPIM,
		SERCOM:	7,
	}
)
```

SPI on the Wio Terminal


```go
var (
	ErrInvalidInputPin	= errors.New("machine: invalid input pin")
	ErrInvalidOutputPin	= errors.New("machine: invalid output pin")
	ErrInvalidClockPin	= errors.New("machine: invalid clock pin")
	ErrInvalidDataPin	= errors.New("machine: invalid data pin")
	ErrNoPinChangeChannel	= errors.New("machine: no channel available for pin interrupt")
)
```



```go
var (
	// UART0 is actually a USB CDC interface.
	UART0 = USBCDC{Buffer: NewRingBuffer()}
)
```



```go
var (
	DAC0 = DAC{}
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




## type DAC

```go
type DAC struct {
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
func (pwm PWM) Configure() error
```

Configure configures a PWM pin for output.


### func (PWM) Set

```go
func (pwm PWM) Set(value uint16)
```

Set turns on the duty cycle for a PWM pin using the provided value.




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




