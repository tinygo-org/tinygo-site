
---
title: nucleo-f722ze
---


## Constants

```go
const (
	LED		= LED_BUILTIN
	LED_BUILTIN	= LED_GREEN
	LED_GREEN	= PB0
	LED_BLUE	= PB7
	LED_RED		= PB14
)
```



```go
const (
	BUTTON		= BUTTON_USER
	BUTTON_USER	= PC13
)
```



```go
const (
	// PD8 and PD9 are connected to the ST-Link Virtual Com Port (VCP)
	UART_TX_PIN	= PD8
	UART_RX_PIN	= PD9
	UART_ALT_FN	= 7	// GPIO_AF7_UART3
)
```

UART pins


```go
const (
	SPI0_SCK_PIN	= PA5
	SPI0_SDI_PIN	= PA6
	SPI0_SDO_PIN	= PA7
)
```

SPI pins


```go
const (
	I2C0_SCL_PIN	= PB8
	I2C0_SDA_PIN	= PB9
)
```

I2C pins


```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.

Deprecated: use 100 * machine.KHz or 400 * machine.KHz instead.


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
	PinRising	PinChange	= 1 << iota
	PinFalling
	PinToggle	= PinRising | PinFalling
)
```



```go
const (
	MAX_NBYTE_SIZE	= 255

	// 100ms delay = 100e6ns / 16ns
	// In runtime_stm32_timers.go, tick is fixed at 16ns per tick
	TIMEOUT_TICKS	= 100e6 / 16

	I2C_NO_STARTSTOP		= 0x0
	I2C_GENERATE_START_WRITE	= 0x80000000 | stm32.I2C_CR2_START
	I2C_GENERATE_START_READ		= 0x80000000 | stm32.I2C_CR2_START | stm32.I2C_CR2_RD_WRN
	I2C_GENERATE_STOP		= 0x80000000 | stm32.I2C_CR2_STOP
)
```



```go
const (
	// Mode Flag
	PinOutput		PinMode	= 0
	PinInput		PinMode	= PinInputFloating
	PinInputFloating	PinMode	= 1
	PinInputPulldown	PinMode	= 2
	PinInputPullup		PinMode	= 3

	// for UART
	PinModeUARTTX	PinMode	= 4
	PinModeUARTRX	PinMode	= 5

	// for I2C
	PinModeI2CSCL	PinMode	= 6
	PinModeI2CSDA	PinMode	= 7

	// for SPI
	PinModeSPICLK	PinMode	= 8
	PinModeSPISDO	PinMode	= 9
	PinModeSPISDI	PinMode	= 10

	// for analog/ADC
	PinInputAnalog	PinMode	= 11

	// for PWM
	PinModePWMOutput	PinMode	= 12
)
```



```go
const RNG_MAX_READ_RETRIES = 1000
```



```go
const PWM_MODE1 = 0x6
```



```go
const (
	AF0_SYSTEM					= 0
	AF1_TIM1_2					= 1
	AF2_TIM3_4_5					= 2
	AF3_TIM8_9_10_11_LPTIM1				= 3
	AF4_I2C1_2_3_USART1				= 4
	AF5_SPI1_2_3_4_5_I2S1_2_3			= 5
	AF6_SPI2_3_I2S2_3_SAI1_UART4			= 6
	AF7_SPI2_3_I2S2_3_USART1_2_3_UART5		= 7
	AF8_SAI2_USART6_UART4_5_7_8_OTG1_FS		= 8
	AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS	= 9
	AF10_SAI2_QUADSPI_SDMMC2_OTG2_HS_OTG1_FS	= 10
	AF11_SDMMC2					= 11
	AF12_UART7_FMC_SDMMC1_OTG2_FS			= 12
	AF15_EVENTOUT					= 15
)
```

Alternative peripheral pin functions


```go
const (
	PA0	= portA + 0
	PA1	= portA + 1
	PA2	= portA + 2
	PA3	= portA + 3
	PA4	= portA + 4
	PA5	= portA + 5
	PA6	= portA + 6
	PA7	= portA + 7
	PA8	= portA + 8
	PA9	= portA + 9
	PA10	= portA + 10
	PA11	= portA + 11
	PA12	= portA + 12
	PA13	= portA + 13
	PA14	= portA + 14
	PA15	= portA + 15

	PB0	= portB + 0
	PB1	= portB + 1
	PB2	= portB + 2
	PB3	= portB + 3
	PB4	= portB + 4
	PB5	= portB + 5
	PB6	= portB + 6
	PB7	= portB + 7
	PB8	= portB + 8
	PB9	= portB + 9
	PB10	= portB + 10
	PB11	= portB + 11
	PB12	= portB + 12
	PB13	= portB + 13
	PB14	= portB + 14
	PB15	= portB + 15

	PC0	= portC + 0
	PC1	= portC + 1
	PC2	= portC + 2
	PC3	= portC + 3
	PC4	= portC + 4
	PC5	= portC + 5
	PC6	= portC + 6
	PC7	= portC + 7
	PC8	= portC + 8
	PC9	= portC + 9
	PC10	= portC + 10
	PC11	= portC + 11
	PC12	= portC + 12
	PC13	= portC + 13
	PC14	= portC + 14
	PC15	= portC + 15

	PD0	= portD + 0
	PD1	= portD + 1
	PD2	= portD + 2
	PD3	= portD + 3
	PD4	= portD + 4
	PD5	= portD + 5
	PD6	= portD + 6
	PD7	= portD + 7
	PD8	= portD + 8
	PD9	= portD + 9
	PD10	= portD + 10
	PD11	= portD + 11
	PD12	= portD + 12
	PD13	= portD + 13
	PD14	= portD + 14
	PD15	= portD + 15

	PE0	= portE + 0
	PE1	= portE + 1
	PE2	= portE + 2
	PE3	= portE + 3
	PE4	= portE + 4
	PE5	= portE + 5
	PE6	= portE + 6
	PE7	= portE + 7
	PE8	= portE + 8
	PE9	= portE + 9
	PE10	= portE + 10
	PE11	= portE + 11
	PE12	= portE + 12
	PE13	= portE + 13
	PE14	= portE + 14
	PE15	= portE + 15

	PF0	= portF + 0
	PF1	= portF + 1
	PF2	= portF + 2
	PF3	= portF + 3
	PF4	= portF + 4
	PF5	= portF + 5
	PF6	= portF + 6
	PF7	= portF + 7
	PF8	= portF + 8
	PF9	= portF + 9
	PF10	= portF + 10
	PF11	= portF + 11
	PF12	= portF + 12
	PF13	= portF + 13
	PF14	= portF + 14
	PF15	= portF + 15

	PG0	= portG + 0
	PG1	= portG + 1
	PG2	= portG + 2
	PG3	= portG + 3
	PG4	= portG + 4
	PG5	= portG + 5
	PG6	= portG + 6
	PG7	= portG + 7
	PG8	= portG + 8
	PG9	= portG + 9
	PG10	= portG + 10
	PG11	= portG + 11
	PG12	= portG + 12
	PG13	= portG + 13
	PG14	= portG + 14
	PG15	= portG + 15

	PH0	= portH + 0
	PH1	= portH + 1
	PH2	= portH + 2
	PH3	= portH + 3
	PH4	= portH + 4
	PH5	= portH + 5
	PH6	= portH + 6
	PH7	= portH + 7
	PH8	= portH + 8
	PH9	= portH + 9
	PH10	= portH + 10
	PH11	= portH + 11
	PH12	= portH + 12
	PH13	= portH + 13
	PH14	= portH + 14
	PH15	= portH + 15

	PI0	= portI + 0
	PI1	= portI + 1
	PI2	= portI + 2
	PI3	= portI + 3
	PI4	= portI + 4
	PI5	= portI + 5
	PI6	= portI + 6
	PI7	= portI + 7
	PI8	= portI + 8
	PI9	= portI + 9
	PI10	= portI + 10
	PI11	= portI + 11
	PI12	= portI + 12
	PI13	= portI + 13
	PI14	= portI + 14
	PI15	= portI + 15
)
```



```go
const (
	ARR_MAX	= 0x10000
	PSC_MAX	= 0x10000
)
```



```go
const APB1_TIM_FREQ = 54e6	// 54MHz

```

Internal use: configured speed of the APB1 and APB2 timers, this should be kept
in sync with any changes to runtime package which configures the oscillators
and clock frequencies


```go
const APB2_TIM_FREQ = 216e6	// 216MHz

```



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
	// USART3 is the hardware serial port connected to the onboard ST-LINK
	// debugger to be exposed as virtual COM port over USB on Nucleo boards.
	UART1	= &_UART1
	_UART1	= UART{
		Buffer:			NewRingBuffer(),
		Bus:			stm32.USART3,
		TxAltFuncSelector:	UART_ALT_FN,
		RxAltFuncSelector:	UART_ALT_FN,
	}
	DefaultUART	= UART1
)
```



```go
var (
	// I2C1 is documented, alias to I2C0 as well
	I2C1	= &I2C{
		Bus:			stm32.I2C1,
		AltFuncSelector:	4,
	}
	I2C0	= I2C1
)
```



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
	TIM1	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM1EN,
		Device:		stm32.TIM1,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA8, AF1_TIM1_2},
				{PE9, AF1_TIM1_2},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA9, AF1_TIM1_2},
				{PE11, AF1_TIM1_2},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA10, AF1_TIM1_2},
				{PE13, AF1_TIM1_2},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA11, AF1_TIM1_2},
				{PE14, AF1_TIM1_2},
			}},
		},
		busFreq:	APB2_TIM_FREQ,
	}

	TIM2	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM2EN,
		Device:		stm32.TIM2,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA0, AF1_TIM1_2},
				{PA5, AF1_TIM1_2},
				{PA15, AF1_TIM1_2},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA1, AF1_TIM1_2},
				{PB3, AF1_TIM1_2},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA2, AF1_TIM1_2},
				{PB10, AF1_TIM1_2},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA3, AF1_TIM1_2},
				{PB11, AF1_TIM1_2},
			}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM3	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM3EN,
		Device:		stm32.TIM3,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA6, AF2_TIM3_4_5},
				{PB4, AF2_TIM3_4_5},
				{PC6, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA7, AF2_TIM3_4_5},
				{PB5, AF2_TIM3_4_5},
				{PC7, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PB0, AF2_TIM3_4_5},
				{PC8, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PB1, AF2_TIM3_4_5},
				{PC9, AF2_TIM3_4_5},
			}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM4	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM4EN,
		Device:		stm32.TIM4,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PB6, AF2_TIM3_4_5},
				{PD12, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PB7, AF2_TIM3_4_5},
				{PD13, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PB8, AF2_TIM3_4_5},
				{PD14, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PB9, AF2_TIM3_4_5},
				{PD15, AF2_TIM3_4_5},
			}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM5	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM5EN,
		Device:		stm32.TIM5,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA0, AF2_TIM3_4_5},
				{PH10, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA1, AF2_TIM3_4_5},
				{PH11, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA2, AF2_TIM3_4_5},
				{PH12, AF2_TIM3_4_5},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA3, AF2_TIM3_4_5},
				{PI0, AF2_TIM3_4_5},
			}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM6	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM6EN,
		Device:		stm32.TIM6,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM7	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM7EN,
		Device:		stm32.TIM7,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM8	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM8EN,
		Device:		stm32.TIM8,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PC6, AF3_TIM8_9_10_11_LPTIM1},
				{PI5, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{
				{PC7, AF3_TIM8_9_10_11_LPTIM1},
				{PI6, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{
				{PC8, AF3_TIM8_9_10_11_LPTIM1},
				{PI7, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{
				{PC9, AF3_TIM8_9_10_11_LPTIM1},
				{PI2, AF3_TIM8_9_10_11_LPTIM1},
			}},
		},
		busFreq:	APB2_TIM_FREQ,
	}

	TIM9	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM9EN,
		Device:		stm32.TIM9,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA2, AF3_TIM8_9_10_11_LPTIM1},
				{PE5, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{
				{PA3, AF3_TIM8_9_10_11_LPTIM1},
				{PE6, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB2_TIM_FREQ,
	}

	TIM10	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM10EN,
		Device:		stm32.TIM10,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PB8, AF3_TIM8_9_10_11_LPTIM1},
				{PF6, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB2_TIM_FREQ,
	}

	TIM11	= TIM{
		EnableRegister:	&stm32.RCC.APB2ENR,
		EnableFlag:	stm32.RCC_APB2ENR_TIM11EN,
		Device:		stm32.TIM11,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PB9, AF3_TIM8_9_10_11_LPTIM1},
				{PF7, AF3_TIM8_9_10_11_LPTIM1},
			}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB2_TIM_FREQ,
	}

	TIM12	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM12EN,
		Device:		stm32.TIM12,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PB14, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
				{PH6, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
			}},
			TimerChannel{Pins: []PinFunction{
				{PB15, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
				{PH9, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
			}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM13	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM13EN,
		Device:		stm32.TIM13,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA6, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
				{PF8, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
			}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB1_TIM_FREQ,
	}

	TIM14	= TIM{
		EnableRegister:	&stm32.RCC.APB1ENR,
		EnableFlag:	stm32.RCC_APB1ENR_TIM14EN,
		Device:		stm32.TIM14,
		Channels: [4]TimerChannel{
			TimerChannel{Pins: []PinFunction{
				{PA7, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
				{PF9, AF9_CAN1_TIM12_13_14_QUADSPI_FMC_OTG2_HS},
			}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
			TimerChannel{Pins: []PinFunction{}},
		},
		busFreq:	APB1_TIM_FREQ,
	}
)
```



```go
var (
	ErrPWMPeriodTooLong = errors.New("pwm: period too long")
)
```



```go
var Serial = DefaultUART
```

Serial is implemented via the default (usually the first) UART on the chip.





### func CPUFrequency

```go
func CPUFrequency() uint32
```



### func GetRNG

```go
func GetRNG() (uint32, error)
```

GetRNG returns 32 bits of cryptographically secure random data


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.




## type ADC

```go
type ADC struct {
	Pin Pin
}
```






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





## type ChannelCallback

```go
type ChannelCallback func(channel uint8)
```






## type I2C

```go
type I2C struct {
	Bus		*stm32.I2C_Type
	AltFuncSelector	uint8
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


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) error
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
	SCL	Pin
	SDA	Pin
}
```

I2CConfig is used to store config info for I2C.





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

Configure this pin with the given configuration


### func (Pin) ConfigureAltFunc

```go
func (p Pin) ConfigureAltFunc(config PinConfig, altFunc uint8)
```

Configure this pin with the given configuration including alternate

	function mapping if necessary.


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

PortMaskClear returns the register and mask to disable a given port. This can
be used to implement bit-banged drivers.


### func (Pin) PortMaskSet

```go
func (p Pin) PortMaskSet() (*uint32, uint32)
```

PortMaskSet returns the register and mask to enable a given GPIO pin. This
can be used to implement bit-banged drivers.


### func (Pin) Set

```go
func (p Pin) Set(high bool)
```

Set the pin to high or low.
Warning: only use this on an output pin!


### func (Pin) SetAltFunc

```go
func (p Pin) SetAltFunc(af uint8)
```

SetAltFunc maps the given alternative function to the I/O pin


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

---------- General pin operations ----------





## type PinConfig

```go
type PinConfig struct {
	Mode PinMode
}
```






## type PinFunction

```go
type PinFunction struct {
	Pin	Pin
	AltFunc	uint8
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




## type TIM

```go
type TIM struct {
	EnableRegister	*volatile.Register32
	EnableFlag	uint32
	Device		*stm32.TIM_Type
	Channels	[4]TimerChannel
	UpInterrupt	interrupt.Interrupt
	OCInterrupt	interrupt.Interrupt

	wraparoundCallback	TimerCallback
	channelCallbacks	[4]ChannelCallback

	busFreq	uint64
}
```




### func (*TIM) Channel

```go
func (t *TIM) Channel(pin Pin) (uint8, error)
```

Channel returns a PWM channel for the given pin.


### func (*TIM) Configure

```go
func (t *TIM) Configure(config PWMConfig) error
```

Configure enables and configures this PWM.


### func (*TIM) Count

```go
func (t *TIM) Count() uint32
```



### func (*TIM) Set

```go
func (t *TIM) Set(channel uint8, value uint32)
```

Set updates the channel value. This is used to control the channel duty
cycle. For example, to set it to a 25% duty cycle, use:

	t.Set(ch, t.Top() / 4)

ch.Set(0) will set the output to low and ch.Set(ch.Top()) will set the output
to high, assuming the output isn't inverted.


### func (*TIM) SetInverting

```go
func (t *TIM) SetInverting(channel uint8, inverting bool)
```

SetInverting sets whether to invert the output of this channel.
Without inverting, a 25% duty cycle would mean the output is high for 25% of
the time and low for the rest. Inverting flips the output as if a NOT gate
was placed at the output, meaning that the output would be 25% low and 75%
high with a duty cycle of 25%.


### func (*TIM) SetMatchInterrupt

```go
func (t *TIM) SetMatchInterrupt(channel uint8, callback ChannelCallback) error
```

Sets a callback to be called when a channel reaches it's set-point.

For example, if `t.Set(ch, t.Top() / 4)` is used then the callback will
be called every quarter-period of the timer's base Period.


### func (*TIM) SetPeriod

```go
func (t *TIM) SetPeriod(period uint64) error
```

SetPeriod updates the period of this PWM peripheral.
To set a particular frequency, use the following formula:

	period = 1e9 / frequency

If you use a period of 0, a period that works well for LEDs will be picked.

SetPeriod will not change the prescaler, but also won't change the current
value in any of the channels. This means that you may need to update the
value for the particular channel.

Note that you cannot pick any arbitrary period after the PWM peripheral has
been configured. If you want to switch between frequencies, pick the lowest
frequency (longest period) once when calling Configure and adjust the
frequency here as needed.


### func (*TIM) SetWraparoundInterrupt

```go
func (t *TIM) SetWraparoundInterrupt(callback TimerCallback) error
```

SetWraparoundInterrupt configures a callback to be called each
time the timer 'wraps-around'.

For example, if `Configure(PWMConfig{Period:1000000})` is used,
to set the timer period to 1ms, this callback will be called every
1ms.


### func (*TIM) Top

```go
func (t *TIM) Top() uint32
```

Top returns the current counter top, for use in duty cycle calculation. It
will only change with a call to Configure or SetPeriod, otherwise it is
constant.

The value returned here is hardware dependent. In general, it's best to treat
it as an opaque value that can be divided by some number and passed to
pwm.Set (see pwm.Set for more information).


### func (*TIM) Unset

```go
func (t *TIM) Unset(channel uint8)
```

Unset disables a channel, including any configured interrupts.




## type TimerCallback

```go
type TimerCallback func()
```






## type TimerChannel

```go
type TimerChannel struct {
	Pins []PinFunction
}
```






## type UART

```go
type UART struct {
	Buffer			*RingBuffer
	Bus			*stm32.USART_Type
	Interrupt		interrupt.Interrupt
	TxAltFuncSelector	uint8
	RxAltFuncSelector	uint8

	// Registers specific to the chip
	rxReg		*volatile.Register32
	txReg		*volatile.Register32
	statusReg	*volatile.Register32
	txEmptyFlag	uint32
}
```

UART representation



### func (*UART) Buffered

```go
func (uart *UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
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

SetBaudRate sets the communication speed for the UART. Defer to chip-specific
routines for calculation


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
type UARTParity uint8
```

UARTParity is the parity setting to be used for UART communication.





