
---
title: atsame54-xpro
---


## Constants

```go
const (
	LED	= PC18
	BUTTON	= PB31
)
```



```go
const (

	// Extension Header EXT1
	EXT1_PIN3_ADC_P		= PB04
	EXT1_PIN4_ADC_N		= PB05
	EXT1_PIN5_GPIO1		= PA06
	EXT1_PIN6_GPIO2		= PA07
	EXT1_PIN7_PWM_P		= PB08
	EXT1_PIN8_PWM_N		= PB09
	EXT1_PIN9_IRQ		= PB07
	EXT1_PIN9_GPIO		= PB07
	EXT1_PIN10_SPI_SS_B	= PA27
	EXT1_PIN10_GPIO		= PA27
	EXT1_PIN11_TWI_SDA	= PA22
	EXT1_PIN12_TWI_SCL	= PA23
	EXT1_PIN13_UART_RX	= PA05
	EXT1_PIN14_UART_TX	= PA04
	EXT1_PIN15_SPI_SS_A	= PB28
	EXT1_PIN16_SPI_SDO	= PB27
	EXT1_PIN17_SPI_SDI	= PB29
	EXT1_PIN18_SPI_SCK	= PB26

	// Extension Header EXT2
	EXT2_PIN3_ADC_P		= PB00
	EXT2_PIN4_ADC_N		= PA03
	EXT2_PIN5_GPIO1		= PB01
	EXT2_PIN6_GPIO2		= PB06
	EXT2_PIN7_PWM_P		= PB14
	EXT2_PIN8_PWM_N		= PB15
	EXT2_PIN9_IRQ		= PD00
	EXT2_PIN9_GPIO		= PD00
	EXT2_PIN10_SPI_SS_B	= PB02
	EXT2_PIN10_GPIO		= PB02
	EXT2_PIN11_TWI_SDA	= PD08
	EXT2_PIN12_TWI_SCL	= PD09
	EXT2_PIN13_UART_RX	= PB17
	EXT2_PIN14_UART_TX	= PB16
	EXT2_PIN15_SPI_SS_A	= PC06
	EXT2_PIN16_SPI_SDO	= PC04
	EXT2_PIN17_SPI_SDI	= PC07
	EXT2_PIN18_SPI_SCK	= PC05

	// Extension Header EXT3
	EXT3_PIN3_ADC_P		= PC02
	EXT3_PIN4_ADC_N		= PC03
	EXT3_PIN5_GPIO1		= PC01
	EXT3_PIN6_GPIO2		= PC10
	EXT3_PIN7_PWM_P		= PD10
	EXT3_PIN8_PWM_N		= PD11
	EXT3_PIN9_IRQ		= PC30
	EXT3_PIN9_GPIO		= PC30
	EXT3_PIN10_SPI_SS_B	= PC31
	EXT3_PIN10_GPIO		= PC31
	EXT3_PIN11_TWI_SDA	= PD08
	EXT3_PIN12_TWI_SCL	= PD09
	EXT3_PIN13_UART_RX	= PC23
	EXT3_PIN14_UART_TX	= PC22
	EXT3_PIN15_SPI_SS_A	= PC14
	EXT3_PIN16_SPI_SDO	= PC04
	EXT3_PIN17_SPI_SDI	= PC07
	EXT3_PIN18_SPI_SCK	= PC05

	// SD_CARD
	SD_CARD_MCDA0	= PB18
	SD_CARD_MCDA1	= PB19
	SD_CARD_MCDA2	= PB20
	SD_CARD_MCDA3	= PB21
	SD_CARD_MCCK	= PA21
	SD_CARD_MCCDA	= PA20
	SD_CARD_DETECT	= PD20
	SD_CARD_PROTECT	= PD21

	// I2C
	I2C_SDA	= PD08
	I2C_SCL	= PD09

	// CAN
	CAN0_TX	= PA22
	CAN0_RX	= PA23

	CAN1_STANDBY	= PC13
	CAN1_TX		= PB12
	CAN1_RX		= PB13

	CAN_STANDBY	= CAN1_STANDBY
	CAN_TX		= CAN1_TX
	CAN_RX		= CAN1_RX

	// PDEC
	PDEC_PHASE_A	= PC16
	PDEC_PHASE_B	= PC17
	PDEC_INDEX	= PC18

	// PCC
	PCC_I2C_SDA	= PD08
	PCC_I2C_SCL	= PD09
	PCC_VSYNC_DEN1	= PA12
	PCC_HSYNC_DEN2	= PA13
	PCC_CLK		= PA14
	PCC_XCLK	= PA15
	PCC_DATA00	= PA16
	PCC_DATA01	= PA17
	PCC_DATA02	= PA18
	PCC_DATA03	= PA19
	PCC_DATA04	= PA20
	PCC_DATA05	= PA21
	PCC_DATA06	= PA22
	PCC_DATA07	= PA23
	PCC_DATA08	= PB14
	PCC_DATA09	= PB15
	PCC_RESET	= PC12
	PCC_PWDN	= PC11

	// Ethernet
	ETHERNET_TXCK	= PA14
	ETHERNET_TXEN	= PA17
	ETHERNET_TX0	= PA18
	ETHERNET_TX1	= PA19
	ETHERNET_RXER	= PA15
	ETHERNET_RX0	= PA13
	ETHERNET_RX1	= PA12
	ETHERNET_RXDV	= PC20
	ETHERNET_MDIO	= PC12
	ETHERNET_MDC	= PC11
	ETHERNET_INT	= PD12
	ETHERNET_RESET	= PC21

	PIN_QT_BUTTON	= PA16
	PIN_BTN0	= PB31
	PIN_ETH_LED	= PC15
	PIN_LED0	= PC18
	PIN_ADC_DAC	= PA02
	PIN_VBUS_DETECT	= PC00
	PIN_USB_ID	= PC19
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
	// Extension Header EXT1
	UART_TX_PIN	= PA04	// TX : SERCOM0/PAD[0]
	UART_RX_PIN	= PA05	// RX : SERCOM0/PAD[1]

	// Extension Header EXT2
	UART2_TX_PIN	= PB16	// TX : SERCOM5/PAD[0]
	UART2_RX_PIN	= PB17	// RX : SERCOM5/PAD[1]

	// Extension Header EXT3
	UART3_TX_PIN	= PC22	// TX : SERCOM1/PAD[0]
	UART3_RX_PIN	= PC23	// RX : SERCOM1/PAD[1]

	// Virtual COM Port
	UART4_TX_PIN	= PB25	// TX : SERCOM2/PAD[0]
	UART4_RX_PIN	= PB24	// RX : SERCOM2/PAD[1]
)
```

UART pins


```go
const (
	// Extension Header EXT1
	SDA0_PIN	= PA22	// SDA: SERCOM3/PAD[0]
	SCL0_PIN	= PA23	// SCL: SERCOM3/PAD[1]

	// Extension Header EXT2
	SDA1_PIN	= PD08	// SDA: SERCOM7/PAD[0]
	SCL1_PIN	= PD09	// SCL: SERCOM7/PAD[1]

	// Extension Header EXT3
	SDA2_PIN	= PD08	// SDA: SERCOM7/PAD[0]
	SCL2_PIN	= PD09	// SCL: SERCOM7/PAD[1]

	// Data Gateway Interface
	SDA_DGI_PIN	= PD08	// SDA: SERCOM7/PAD[0]
	SCL_DGI_PIN	= PD09	// SCL: SERCOM7/PAD[1]

	SDA_PIN	= SDA0_PIN
	SCL_PIN	= SCL0_PIN
)
```

I2C pins


```go
const (
	// Extension Header EXT1
	SPI0_SCK_PIN	= PB26	// SCK: SERCOM4/PAD[1]
	SPI0_SDO_PIN	= PB27	// SDO: SERCOM4/PAD[0]
	SPI0_SDI_PIN	= PB29	// SDI: SERCOM4/PAD[3]
	SPI0_SS_PIN	= PB28	// SS : SERCOM4/PAD[2]

	// Extension Header EXT2
	SPI1_SCK_PIN	= PC05	// SCK: SERCOM6/PAD[1]
	SPI1_SDO_PIN	= PC04	// SDO: SERCOM6/PAD[0]
	SPI1_SDI_PIN	= PC07	// SDI: SERCOM6/PAD[3]
	SPI1_SS_PIN	= PC06	// SS : SERCOM6/PAD[2]

	// Extension Header EXT3
	SPI2_SCK_PIN	= PC05	// SCK: SERCOM6/PAD[1]
	SPI2_SDO_PIN	= PC04	// SDO: SERCOM6/PAD[0]
	SPI2_SDI_PIN	= PC07	// SDI: SERCOM6/PAD[3]
	SPI2_SS_PIN	= PC14	// SS : GPIO

	// Data Gateway Interface
	SPI_DGI_SCK_PIN	= PC05	// SCK: SERCOM6/PAD[1]
	SPI_DGI_SDO_PIN	= PC04	// SDO: SERCOM6/PAD[0]
	SPI_DGI_SDI_PIN	= PC07	// SDI: SERCOM6/PAD[3]
	SPI_DGI_SS_PIN	= PD01	// SS : GPIO
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
	PA08	Pin	= 8	// peripherals: TCC0 channel 0, TCC1 channel 4
	PA09	Pin	= 9	// peripherals: TCC0 channel 1, TCC1 channel 5
	PA10	Pin	= 10	// peripherals: TCC0 channel 2, TCC1 channel 6
	PA11	Pin	= 11	// peripherals: TCC0 channel 3, TCC1 channel 7
	PA12	Pin	= 12	// peripherals: TCC0 channel 6, TCC1 channel 2
	PA13	Pin	= 13	// peripherals: TCC0 channel 7, TCC1 channel 3
	PA14	Pin	= 14	// peripherals: TCC2 channel 0, TCC1 channel 2
	PA15	Pin	= 15	// peripherals: TCC2 channel 1, TCC1 channel 3
	PA16	Pin	= 16	// peripherals: TCC1 channel 0, TCC0 channel 4
	PA17	Pin	= 17	// peripherals: TCC1 channel 1, TCC0 channel 5
	PA18	Pin	= 18	// peripherals: TCC1 channel 2, TCC0 channel 6
	PA19	Pin	= 19	// peripherals: TCC1 channel 3, TCC0 channel 7
	PA20	Pin	= 20	// peripherals: TCC1 channel 4, TCC0 channel 0
	PA21	Pin	= 21	// peripherals: TCC1 channel 5, TCC0 channel 1
	PA22	Pin	= 22	// peripherals: TCC1 channel 6, TCC0 channel 2
	PA23	Pin	= 23	// peripherals: TCC1 channel 7, TCC0 channel 3
	PA24	Pin	= 24	// peripherals: TCC2 channel 2
	PA25	Pin	= 25	// peripherals: TCC2 channel 3
	PA26	Pin	= 26
	PA27	Pin	= 27
	PA28	Pin	= 28
	PA29	Pin	= 29
	PA30	Pin	= 30	// peripherals: TCC2 channel 0
	PA31	Pin	= 31	// peripherals: TCC2 channel 1
	PB00	Pin	= 32
	PB01	Pin	= 33
	PB02	Pin	= 34	// peripherals: TCC2 channel 2
	PB03	Pin	= 35	// peripherals: TCC2 channel 3
	PB04	Pin	= 36
	PB05	Pin	= 37
	PB06	Pin	= 38
	PB07	Pin	= 39
	PB08	Pin	= 40
	PB09	Pin	= 41
	PB10	Pin	= 42	// peripherals: TCC0 channel 4, TCC1 channel 0
	PB11	Pin	= 43	// peripherals: TCC0 channel 5, TCC1 channel 1
	PB12	Pin	= 44	// peripherals: TCC3 channel 0, TCC0 channel 0
	PB13	Pin	= 45	// peripherals: TCC3 channel 1, TCC0 channel 1
	PB14	Pin	= 46	// peripherals: TCC4 channel 0, TCC0 channel 2
	PB15	Pin	= 47	// peripherals: TCC4 channel 1, TCC0 channel 3
	PB16	Pin	= 48	// peripherals: TCC3 channel 0, TCC0 channel 4
	PB17	Pin	= 49	// peripherals: TCC3 channel 1, TCC0 channel 5
	PB18	Pin	= 50	// peripherals: TCC1 channel 0
	PB19	Pin	= 51	// peripherals: TCC1 channel 1
	PB20	Pin	= 52	// peripherals: TCC1 channel 2
	PB21	Pin	= 53	// peripherals: TCC1 channel 3
	PB22	Pin	= 54
	PB23	Pin	= 55
	PB24	Pin	= 56
	PB25	Pin	= 57
	PB26	Pin	= 58	// peripherals: TCC1 channel 2
	PB27	Pin	= 59	// peripherals: TCC1 channel 3
	PB28	Pin	= 60	// peripherals: TCC1 channel 4
	PB29	Pin	= 61	// peripherals: TCC1 channel 5
	PB30	Pin	= 62	// peripherals: TCC4 channel 0, TCC0 channel 6
	PB31	Pin	= 63	// peripherals: TCC4 channel 1, TCC0 channel 7
	PC00	Pin	= 64
	PC01	Pin	= 65
	PC02	Pin	= 66
	PC03	Pin	= 67
	PC04	Pin	= 68	// peripherals: TCC0 channel 0
	PC05	Pin	= 69	// peripherals: TCC0 channel 1
	PC06	Pin	= 70
	PC07	Pin	= 71
	PC08	Pin	= 72
	PC09	Pin	= 73
	PC10	Pin	= 74	// peripherals: TCC0 channel 0, TCC1 channel 4
	PC11	Pin	= 75	// peripherals: TCC0 channel 1, TCC1 channel 5
	PC12	Pin	= 76	// peripherals: TCC0 channel 2, TCC1 channel 6
	PC13	Pin	= 77	// peripherals: TCC0 channel 3, TCC1 channel 7
	PC14	Pin	= 78	// peripherals: TCC0 channel 4, TCC1 channel 0
	PC15	Pin	= 79	// peripherals: TCC0 channel 5, TCC1 channel 1
	PC16	Pin	= 80	// peripherals: TCC0 channel 0
	PC17	Pin	= 81	// peripherals: TCC0 channel 1
	PC18	Pin	= 82	// peripherals: TCC0 channel 2
	PC19	Pin	= 83	// peripherals: TCC0 channel 3
	PC20	Pin	= 84	// peripherals: TCC0 channel 4
	PC21	Pin	= 85	// peripherals: TCC0 channel 5
	PC22	Pin	= 86	// peripherals: TCC0 channel 6
	PC23	Pin	= 87	// peripherals: TCC0 channel 7
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
	PD08	Pin	= 104	// peripherals: TCC0 channel 1
	PD09	Pin	= 105	// peripherals: TCC0 channel 2
	PD10	Pin	= 106	// peripherals: TCC0 channel 3
	PD11	Pin	= 107	// peripherals: TCC0 channel 4
	PD12	Pin	= 108	// peripherals: TCC0 channel 5
	PD13	Pin	= 109	// peripherals: TCC0 channel 6
	PD14	Pin	= 110
	PD15	Pin	= 111
	PD16	Pin	= 112
	PD17	Pin	= 113
	PD18	Pin	= 114
	PD19	Pin	= 115
	PD20	Pin	= 116	// peripherals: TCC1 channel 0
	PD21	Pin	= 117	// peripherals: TCC1 channel 1
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
const (
	// WatchdogMaxTimeout in milliseconds (16s)
	WatchdogMaxTimeout = (16384 * 1000) / 1024	// CYC16384/1024kHz
)
```



```go
const HSRAM_SIZE = 0x00040000
```



```go
const (
	CANRxFifoSize	= 16
	CANTxFifoSize	= 16
	CANEvFifoSize	= 16
)
```



```go
const (
	CANTransferRate125kbps	CANTransferRate	= 125000
	CANTransferRate250kbps	CANTransferRate	= 250000
	CANTransferRate500kbps	CANTransferRate	= 500000
	CANTransferRate1000kbps	CANTransferRate	= 1000000
	CANTransferRate2000kbps	CANTransferRate	= 2000000
	CANTransferRate4000kbps	CANTransferRate	= 4000000
)
```

CAN transfer rates for CANConfig


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
	// Extension Header EXT1
	UART1	= &sercomUSART0

	// Extension Header EXT2
	UART2	= &sercomUSART5

	// Extension Header EXT3
	UART3	= &sercomUSART1

	// EDBG Virtual COM Port
	UART4	= &sercomUSART2
)
```

UART on the SAM E54 Xplained Pro


```go
var (
	// Extension Header EXT1
	I2C0	= sercomI2CM3

	// Extension Header EXT2
	I2C1	= sercomI2CM7

	// Extension Header EXT3
	I2C2	= sercomI2CM7

	// Data Gateway Interface
	I2C3	= sercomI2CM7
)
```

I2C on the SAM E54 Xplained Pro


```go
var (
	// Extension Header EXT1
	SPI0	= sercomSPIM4

	// Extension Header EXT2
	SPI1	= sercomSPIM6

	// Extension Header EXT3
	SPI2	= sercomSPIM6

	// Data Gateway Interface
	SPI3	= sercomSPIM6
)
```

SPI on the SAM E54 Xplained Pro


```go
var (
	CAN0	= CAN{
		Bus: sam.CAN0,
	}

	CAN1	= CAN{
		Bus: sam.CAN1,
	}
)
```

CAN on the SAM E54 Xplained Pro


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
	DAC0	= DAC{Channel: 0}
	DAC1	= DAC{Channel: 1}
)
```



```go
var Flash flashBlockDevice
```



```go
var Watchdog = &watchdogImpl{}
```

Watchdog provides access to the hardware watchdog available
in the SAMD51.


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
var CANRxFifo [2][(8 + 64) * CANRxFifoSize]byte
```



```go
var CANTxFifo [2][(8 + 64) * CANTxFifoSize]byte
```



```go
var CANEvFifo [2][(8) * CANEvFifoSize]byte
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






### func CANDlcToLength

```go
func CANDlcToLength(dlc byte, isFD bool) byte
```

CANDlcToLength() converts a DLC value to its actual length.


### func CANLengthToDlc

```go
func CANLengthToDlc(length byte, isFD bool) byte
```

CANLengthToDlc() converts its actual length to a DLC value.


### func CPUFrequency

```go
func CPUFrequency() uint32
```



### func CPUReset

```go
func CPUReset()
```

CPUReset performs a hard system reset.


### func ConfigureUSBEndpoint

```go
func ConfigureUSBEndpoint(desc descriptor.Descriptor, epSettings []usb.EndpointConfig, setup []usb.SetupConfig)
```



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
8 bytes (64 bits) and 16 bytes (128 bits) are common.


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





## type CAN

```go
type CAN struct {
	Bus *sam.CAN_Type
}
```




### func (*CAN) Configure

```go
func (can *CAN) Configure(config CANConfig) error
```

Configure this CAN peripheral with the given configuration.


### func (*CAN) Rx

```go
func (can *CAN) Rx() (id uint32, dlc byte, data []byte, isFd, isExtendedID bool)
```

Rx receives a CAN frame. It is easier to use than RxRaw, but not as
flexible.


### func (*CAN) RxFifoIsEmpty

```go
func (can *CAN) RxFifoIsEmpty() bool
```

RxFifoIsEmpty returns whether RxFifo is empty or not.


### func (*CAN) RxFifoIsFull

```go
func (can *CAN) RxFifoIsFull() bool
```

RxFifoIsFull returns whether RxFifo is full or not.


### func (*CAN) RxFifoSize

```go
func (can *CAN) RxFifoSize() int
```

RxFifoSize returns the number of CAN Frames currently stored in the RXFifo.


### func (*CAN) RxRaw

```go
func (can *CAN) RxRaw(e *CANRxBufferElement)
```

RxRaw copies the received CAN frame to CANRxBufferElement.


### func (*CAN) SetInterrupt

```go
func (can *CAN) SetInterrupt(ie uint32, callback func(*CAN)) error
```

SetInterrupt sets an interrupt to be executed when a particular CAN state.

This call will replace a previously set callback. You can pass a nil func
to unset the CAN interrupt. If you do so, the change parameter is ignored
and can be set to any value (such as 0).


### func (*CAN) Tx

```go
func (can *CAN) Tx(id uint32, data []byte, isFD, isExtendedID bool)
```

The Tx transmits CAN frames. It is easier to use than TxRaw, but not as
flexible.


### func (*CAN) TxFifoIsFull

```go
func (can *CAN) TxFifoIsFull() bool
```

TxFifoIsFull returns whether TxFifo is full or not.


### func (*CAN) TxRaw

```go
func (can *CAN) TxRaw(e *CANTxBufferElement)
```

TxRaw sends a CAN Frame according to CANTxBufferElement.




## type CANConfig

```go
type CANConfig struct {
	TransferRate	CANTransferRate
	TransferRateFD	CANTransferRate
	Tx		Pin
	Rx		Pin
	Standby		Pin
}
```

CANConfig holds CAN configuration parameters. Tx and Rx need to be
specified with some pins. When the Standby Pin is specified, configure it
as an output pin and output Low in Configure(). If this operation is not
necessary, specify NoPin.





## type CANRxBufferElement

```go
type CANRxBufferElement struct {
	ESI	bool
	XTD	bool
	RTR	bool
	ID	uint32
	ANMF	bool
	FIDX	uint8
	FDF	bool
	BRS	bool
	DLC	uint8
	RXTS	uint16
	DB	[64]uint8
}
```

CANRxBufferElement is a struct that corresponds to the same5x Rx Buffer and
FIFO Element.



### func (CANRxBufferElement) Data

```go
func (e CANRxBufferElement) Data() []byte
```

Data returns the received data as a slice of the size according to dlc.


### func (CANRxBufferElement) Length

```go
func (e CANRxBufferElement) Length() byte
```

Length returns its actual length.




## type CANTransferRate

```go
type CANTransferRate uint32
```






## type CANTxBufferElement

```go
type CANTxBufferElement struct {
	ESI	bool
	XTD	bool
	RTR	bool
	ID	uint32
	MM	uint8
	EFC	bool
	FDF	bool
	BRS	bool
	DLC	uint8
	DB	[64]uint8
}
```

CANTxBufferElement is a struct that corresponds to the same5x' Tx Buffer
Element.





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
func (i2c *I2C) SetBaudRate(br uint32) error
```

SetBaudRate sets the communication speed for I2C.


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





