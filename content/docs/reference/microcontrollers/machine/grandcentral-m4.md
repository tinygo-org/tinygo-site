
---
title: grandcentral-m4
---


## Constants

```go
const (
	//  = Pin     Alt. Function        SERCOM   PWM Timer   Interrupt
	//   ------  -------------------- -------- ----------- -----------
	D0	= PB25	// UART1 RX              0[1]                 EXTI9
	D1	= PB24	// UART1 TX              0[0]                 EXTI8
	D2	= PC18	//                                TCC0[2]     EXTI2
	D3	= PC19	//                                TCC0[3]     EXTI3
	D4	= PC20	//                                TCC0[4]     EXTI4
	D5	= PC21	//                                TCC0[5]     EXTI5
	D6	= PD20	//                                TCC1[0]     EXTI10
	D7	= PD21	//                                TCC1[1]     EXTI11
	D8	= PB18	//                                TCC1[0]     EXTI2
	D9	= PB02	//                                TC6[0]      EXTI3
	D10	= PB22	//                                TC7[0]      EXTI6
	D11	= PB23	//                                            EXTI7
	D12	= PB00	//                                TC7[0]      EXTI0
	D13	= PB01	// On-board LED                   TC7[1]      EXTI1
	D14	= PB16	// UART4 TX, I2S0 SCK    5[0]     TC6[0]      EXTI0
	D15	= PB17	// UART4 RX, I2S0 MCK    5[1]                 EXTI1
	D16	= PC22	// UART3 TX              1[0]                 EXTI6
	D17	= PC23	// UART3 RX              1[1]                 EXTI6
	D18	= PB12	// UART2 TX              4[0]     TCC3[0]     EXTI12
	D19	= PB13	// UART2 RX              4[1]     TCC3[1]     EXTI13
	D20	= PB20	// I2C0 SDA              3[0]                 EXTI4
	D21	= PB21	// I2C0 SCL              3[1]                 EXTI5
	D22	= PD12	//                                            EXTI7
	D23	= PA15	//                                TCC2[1]     EXTI15
	D24	= PC17	// I2C1 SCL              6[1]     TCC0[1]     EXTI1
	D25	= PC16	// I2C1 SDA              6[0]     TCC0[0]     EXTI0
	D26	= PA12	// PCC DEN1                       TC2[0]      EXTI12
	D27	= PA13	// PCC DEN2                       TC2[1]      EXTI13
	D28	= PA14	// PCC CLK                        TCC2[0]     EXTI14
	D29	= PB19	// PCC XCLK                                   EXTI3
	D30	= PA23	// PCC D7                         TC4[1]      EXTI7
	D31	= PA22	// PCC D6, I2S0 SDI               TC4[0]      EXTI6
	D32	= PA21	// PCC D5, I2S0 SDO                           EXTI5
	D33	= PA20	// PCC D4, I2S0 FS                            EXTI4
	D34	= PA19	// PCC D3                         TC3[1]      EXTI3
	D35	= PA18	// PCC D2                         TC3[0]      EXTI2
	D36	= PA17	// PCC D1                                     EXTI1
	D37	= PA16	// PCC D0                                     EXTI0
	D38	= PB15	// PCC D9                         TCC4[1]     EXTI15
	D39	= PB14	// PCC D8                         TCC4[0]     EXTI14
	D40	= PC13	// PCC D11                                    EXTI13
	D41	= PC12	// PCC D10                                    EXTI12
	D42	= PC15	// PCC D13                                    EXTI15
	D43	= PC14	// PCC D12                                    EXTI14
	D44	= PC11	//                                            EXTI11
	D45	= PC10	//                                            EXTI10
	D46	= PC06	//                                            EXTI6
	D47	= PC07	//                                            EXTI5
	D48	= PC04	//                                            EXTI4
	D49	= PC05	//                                            EXTI5
	D50	= PD11	// SPI0 SDI              7[3]                 EXTI11
	D51	= PD08	// SPI0 SDO              7[0]                 EXTI8
	D52	= PD09	// SPI0 SCK              7[1]                 EXTI9
	D53	= PD10	// SPI0 CS                                    EXTI10
	D54	= PB05	// ADC1 (A8)                                  EXTI5
	D55	= PB06	// ADC1 (A9)                                  EXTI6
	D56	= PB07	// ADC1 (A10)                                 EXTI7
	D57	= PB08	// ADC1 (A11)                                 EXTI8
	D58	= PB09	// ADC1 (A12)                                 EXTI9
	D59	= PA04	// ADC0 (A13)                     TC0[0]      EXTI4
	D60	= PA06	// ADC0 (A14)                     TC1[0]      EXTI6
	D61	= PA07	// ADC0 (A15)                     TC1[1]      EXTI7
	D62	= PB20	// I2C0 SDA              3[0]     TCC1[2]     EXTI4
	D63	= PB21	// I2C0 SCL              3[1]     TCC1[3]     EXTI5
	D64	= PD11	// SPI0 SDI              7[3]                 EXTI6
	D65	= PD08	// SPI0 SDO              7[0]                 EXTI3
	D66	= PD09	// SPI0 SCK              7[1]                 EXTI4
	D67	= PA02	// ADC0 (A0), DAC0                            EXTI2
	D68	= PA05	// ADC0 (A1), DAC1                            EXTI5
	D69	= PB03	// ADC0 (A2)                      TC6[1]      EXTI3
	D70	= PC00	// ADC1 (A3)                                  EXTI0
	D71	= PC01	// ADC1 (A4)                                  EXTI1
	D72	= PC02	// ADC1 (A5)                                  EXTI2
	D73	= PC03	// ADC1 (A6)                                  EXTI3
	D74	= PB04	// ADC1 (A7)                                  EXTI4
	D75	= PC31	// UART RX LED
	D76	= PC30	// UART TX LED
	D77	= PA27	// USB HOST EN
	D78	= PA24	// USB DM                                     EXTI8
	D79	= PA25	// USB DP                                     EXTI9
	D80	= PB29	// SD/SPI1 SDI           2[3]
	D81	= PB27	// SD/SPI1 SCK           2[1]
	D82	= PB26	// SD/SPI1 SDO           2[0]
	D83	= PB28	// SD/SPI1 CS
	D84	= PA03	// AREF                                       EXTI3
	D85	= PA02	// DAC0                                       EXTI2
	D86	= PA05	// DAC1                                       EXTI5
	D87	= PB01	// On-board LED (D13)             TC7[1]      EXTI1
	D88	= PC24	// On-board NeoPixel
	D89	= PB10	// QSPI SCK                                   EXTI10
	D90	= PB11	// QSPI CS                                    EXTI11
	D91	= PA08	// QSPI ID0                                   EXTI(NMI)
	D92	= PA09	// QSPI ID1                                   EXTI9
	D93	= PA10	// QSPI ID2                                   EXTI10
	D94	= PA11	// QSPI ID3                                   EXTI11
	D95	= PB31	// SD Detect                                  EXTI15
	D96	= PB30	// SWO                                        EXTI14
)
```

Digital pins


```go
const (
	A0	= D67	// (PA02) ADC0 ch. 0,
	A1	= D68	// (PA05) ADC0 ch. 5,
	A2	= D69	// (PB03) ADC0 ch. 15
	A3	= D70	// (PC00) ADC1 ch. 10
	A4	= D71	// (PC01) ADC1 ch. 11
	A5	= D72	// (PC02) ADC1 ch. 4
	A6	= D73	// (PC03) ADC1 ch. 5
	A7	= D74	// (PB04) ADC1 ch. 6
	A8	= D54	// (PB05) ADC1 ch. 7
	A9	= D55	// (PB06) ADC1 ch. 8
	A10	= D56	// (PB07) ADC1 ch. 9
	A11	= D57	// (PB08) ADC1 ch. 0
	A12	= D58	// (PB09) ADC1 ch. 1
	A13	= D59	// (PA04) ADC0 ch. 4
	A14	= D60	// (PA06) ADC0 ch. 6
	A15	= D61	// (PA07) ADC0 ch. 7

	AREF	= D84	// (PA03)
)
```

Analog pins


```go
const (
	LED_PIN		= D13	// (PB01), also on D87
	UART_RX_LED_PIN	= D75	// (PC31)
	UART_TX_LED_PIN	= D76	// (PC30)
	NEOPIXEL_PIN	= D88	// (PC24)

	// aliases used by examples and drivers
	LED		= LED_PIN
	LED_RX		= UART_RX_LED_PIN
	LED_TX		= UART_TX_LED_PIN
	NEOPIXEL	= NEOPIXEL_PIN
	WS2812		= NEOPIXEL_PIN
)
```

LED pins


```go
const (
	UART1_RX_PIN	= D0	// (PB25)
	UART1_TX_PIN	= D1	// (PB24)

	UART2_RX_PIN	= D19	// (PB13)
	UART2_TX_PIN	= D18	// (PB12)

	UART3_RX_PIN	= D17	// (PC23)
	UART3_TX_PIN	= D16	// (PC22)

	UART4_RX_PIN	= D15	// (PB17)
	UART4_TX_PIN	= D14	// (PB16)

	UART_RX_PIN	= UART1_RX_PIN	// default pins
	UART_TX_PIN	= UART1_TX_PIN	//
)
```

UART pins


```go
const (
	SPI0_SCK_PIN	= D66	// (PD09), also on D52
	SPI0_SDO_PIN	= D65	// (PD08), also on D51
	SPI0_SDI_PIN	= D64	// (PD11), also on D50
	SPI0_CS_PIN	= D53	// (PD10)

	SPI1_SCK_PIN	= D81	// (PB27)
	SPI1_SDO_PIN	= D82	// (PB26)
	SPI1_SDI_PIN	= D80	// (PB29)

	SPI_SCK_PIN	= SPI0_SCK_PIN	// default pins
	SPI_SDO_PIN	= SPI0_SDO_PIN	//
	SPI_SDI_PIN	= SPI0_SDI_PIN	//
	SPI_CS_PIN	= SPI0_CS_PIN	//
)
```

SPI pins


```go
const (
	I2C0_SDA_PIN	= D62	// (PB20), also on D20
	I2C0_SCL_PIN	= D63	// (PB21), also on D21

	I2C1_SDA_PIN	= D25	// (PC16)
	I2C1_SCL_PIN	= D24	// (PC17)

	I2C_SDA_PIN	= I2C0_SDA_PIN	// default pins
	I2C_SCL_PIN	= I2C0_SCL_PIN	//

	SDA_PIN	= I2C_SDA_PIN	// unconventional pin names
	SCL_PIN	= I2C_SCL_PIN	//  (required by machine_atsamd51.go)
)
```

I2C pins


```go
const (
	I2S0_SCK_PIN	= D14	// (PB16)
	I2S0_MCK_PIN	= D15	// (PB17)
	I2S0_FS_PIN	= D33	// (PA20)
	I2S0_SDO_PIN	= D32	// (PA21)
	I2S0_SDI_PIN	= D31	// (PA22)

	I2S_SCK_PIN	= I2S0_SCK_PIN	// default pins
	I2S_WS_PIN	= I2S0_FS_PIN	//
	I2S_SD_PIN	= I2S0_SDO_PIN	//
)
```

I2S pins


```go
const (
	SD0_SCK_PIN	= D81	// (PB27)
	SD0_SDO_PIN	= D82	// (PB26)
	SD0_SDI_PIN	= D80	// (PB29)
	SD0_CS_PIN	= D83	// (PB28)
	SD0_DET_PIN	= D95	// (PB31)

	SDCARD_SCK_PIN	= SD0_SCK_PIN	// default pins
	SDCARD_SDO_PIN	= SD0_SDO_PIN	//
	SDCARD_SDI_PIN	= SD0_SDI_PIN	//
	SDCARD_CS_PIN	= SD0_CS_PIN	//
	SDCARD_DET_PIN	= SD0_DET_PIN	//
)
```

SD card pins


```go
const (
	RESET_MAGIC_VALUE = 0xF01669EF	// Used to reset into bootloader
)
```

Other peripheral constants


```go
const (
	USBCDC_HOSTEN_PIN	= D77	// (PA27) host enable
	USBCDC_DM_PIN		= D78	// (PA24) D-
	USBCDC_DP_PIN		= D79	// (PA25) D+
)
```

USB CDC pins


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
	UART1	= &sercomUSART0
	UART2	= &sercomUSART4
	UART3	= &sercomUSART1
	UART4	= &sercomUSART5

	DefaultUART	= UART1
)
```

UART on the Grand Central M4


```go
var (
	SPI0	= sercomSPIM7
	SPI1	= sercomSPIM2	// SD card
)
```

SPI on the Grand Central M4


```go
var (
	I2C0	= sercomI2CM3
	I2C1	= sercomI2CM6
)
```

I2C on the Grand Central M4


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
	// USB is a USB CDC interface.
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
var Serial = USB
```

Serial is implemented via USB (USB-CDC).





### func CPUFrequency

```go
func CPUFrequency() uint32
```



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





## type USBCDC

```go
type USBCDC struct {
	Buffer			*RingBuffer
	TxIdx			volatile.Register8
	waitTxc			bool
	waitTxcRetryCount	uint8
	sent			bool
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




