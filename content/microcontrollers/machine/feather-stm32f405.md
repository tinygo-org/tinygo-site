
---
title: feather-stm32f405
---


## Constants

```go
const (
	NUM_DIGITAL_IO_PINS	= 39
	NUM_ANALOG_IO_PINS	= 7
)
```



```go
const (
	// Arduino pin = MCU port pin // primary functions (alternate functions)
	D0	= PB11	// USART3 RX, PWM TIM2_CH4 (I2C2 SDA)
	D1	= PB10	// USART3 TX, PWM TIM2_CH3 (I2C2 SCL, I2S2 BCK)
	D2	= PB3	// GPIO, SPI3 FLASH SCK
	D3	= PB4	// GPIO, SPI3 FLASH MISO
	D4	= PB5	// GPIO, SPI3 FLASH MOSI
	D5	= PC7	// GPIO, PWM TIM3_CH2 (USART6 RX, I2S3 MCK)
	D6	= PC6	// GPIO, PWM TIM3_CH1 (USART6 TX, I2S2 MCK)
	D7	= PA15	// GPIO, SPI3 FLASH CS
	D8	= PC0	// GPIO, Neopixel
	D9	= PB8	// GPIO, PWM TIM4_CH3 (CAN1 RX, I2C1 SCL)
	D10	= PB9	// GPIO, PWM TIM4_CH4 (CAN1 TX, I2C1 SDA, I2S2 WSL)
	D11	= PC3	// GPIO (I2S2 SD, SPI2 MOSI)
	D12	= PC2	// GPIO (I2S2ext SD, SPI2 MISO)
	D13	= PC1	// GPIO, Builtin LED
	D14	= PB7	// I2C1 SDA, PWM TIM4_CH2 (USART1 RX)
	D15	= PB6	// I2C1 SCL, PWM TIM4_CH1 (USART1 TX, CAN2 TX)
	D16	= PA4	// A0 (DAC OUT1)
	D17	= PA5	// A1 (DAC OUT2, SPI1 SCK)
	D18	= PA6	// A2, PWM TIM3_CH1 (SPI1 MISO)
	D19	= PA7	// A3, PWM TIM3_CH2 (SPI1 MOSI)
	D20	= PC4	// A4
	D21	= PC5	// A5
	D22	= PA3	// A6
	D23	= PB13	// SPI2 SCK, PWM TIM1_CH1N (I2S2 BCK, CAN2 TX)
	D24	= PB14	// SPI2 MISO, PWM TIM1_CH2N (I2S2ext SD)
	D25	= PB15	// SPI2 MOSI, PWM TIM1_CH3N (I2S2 SD)
	D26	= PC8	// SDIO
	D27	= PC9	// SDIO
	D28	= PC10	// SDIO
	D29	= PC11	// SDIO
	D30	= PC12	// SDIO
	D31	= PD2	// SDIO
	D32	= PB12	// SD Detect
	D33	= PC14	// OSC32
	D34	= PC15	// OSC32
	D35	= PA11	// USB D+
	D36	= PA12	// USB D-
	D37	= PA13	// SWDIO
	D38	= PA14	// SWCLK
)
```

Digital pins


```go
const (
	A0	= D16	// ADC12 IN4
	A1	= D17	// ADC12 IN5
	A2	= D18	// ADC12 IN6
	A3	= D19	// ADC12 IN7
	A4	= D20	// ADC12 IN14
	A5	= D21	// ADC12 IN15
	A6	= D22	// VBAT
)
```

Analog pins


```go
const (
	NUM_BOARD_LED		= 1
	NUM_BOARD_NEOPIXEL	= 1

	LED_RED		= D13
	LED_NEOPIXEL	= D8
	LED_BUILTIN	= LED_RED
	LED		= LED_BUILTIN
)
```



```go
const (
	// #===========#==========#==============#============#=======#=======#
	// | Interface | Hardware |  Bus(Freq)   | RX/TX Pins | AltFn | Alias |
	// #===========#==========#==============#============#=======#=======#
	// |   UART1   |  USART3  | APB1(42 MHz) |   D0/D1    |   7   |   ~   |
	// |   UART2   |  USART6  | APB2(84 MHz) |   D5/D6    |   8   |   ~   |
	// |   UART3   |  USART1  | APB2(84 MHz) |  D14/D15   |   7   |   ~   |
	// | --------- | -------- | ------------ | ---------- | ----- | ----- |
	// |   UART0   |  USART3  | APB1(42 MHz) |   D0/D1    |   7   | UART1 |
	// #===========#==========#==============#============#=======#=======#
	NUM_UART_INTERFACES	= 3

	UART1_RX_PIN	= D0	// UART1 = hardware: USART3
	UART1_TX_PIN	= D1	//

	UART2_RX_PIN	= D5	// UART2 = hardware: USART6
	UART2_TX_PIN	= D6	//

	UART3_RX_PIN	= D14	// UART3 = hardware: USART1
	UART3_TX_PIN	= D15	//

	UART0_RX_PIN	= UART1_RX_PIN	// UART0 = alias: UART1
	UART0_TX_PIN	= UART1_TX_PIN	//

	UART_RX_PIN	= UART0_RX_PIN	// default/primary UART pins
	UART_TX_PIN	= UART0_TX_PIN	//
)
```



```go
const (
	// #===========#==========#==============#==================#=======#=======#
	// | Interface | Hardware |  Bus(Freq)   | SCK/SDI/SDO Pins | AltFn | Alias |
	// #===========#==========#==============#==================#=======#=======#
	// |   SPI1    |   SPI2   | APB1(42 MHz) |    D23/D24/D25   |   5   |   ~   |
	// |   SPI2    |   SPI3   | APB1(42 MHz) |     D2/D3/D4     |   6   |   ~   |
	// |   SPI3    |   SPI1   | APB2(84 MHz) |    D17/D18/D19   |   5   |   ~   |
	// | --------- | -------- | ------------ | ---------------- | ----- | ----- |
	// |   SPI0    |   SPI2   | APB1(42 MHz) |    D23/D24/D25   |   5   | SPI1  |
	// #===========#==========#==============#==================#=======#=======#
	NUM_SPI_INTERFACES	= 3

	SPI1_SCK_PIN	= D23	//
	SPI1_SDI_PIN	= D24	// SPI1 = hardware: SPI2
	SPI1_SDO_PIN	= D25	//

	SPI2_SCK_PIN	= D2	//
	SPI2_SDI_PIN	= D3	// SPI2 = hardware: SPI3
	SPI2_SDO_PIN	= D4	//

	SPI3_SCK_PIN	= D17	//
	SPI3_SDI_PIN	= D18	// SPI3 = hardware: SPI1
	SPI3_SDO_PIN	= D19	//

	SPI0_SCK_PIN	= SPI1_SCK_PIN	//
	SPI0_SDI_PIN	= SPI1_SDI_PIN	// SPI0 = alias: SPI1
	SPI0_SDO_PIN	= SPI1_SDO_PIN	//

	SPI_SCK_PIN	= SPI0_SCK_PIN	//
	SPI_SDI_PIN	= SPI0_SDI_PIN	// default/primary SPI pins
	SPI_SDO_PIN	= SPI0_SDO_PIN	//
)
```



```go
const (
	// #===========#==========#==============#==============#=======#=======#
	// | Interface | Hardware |  Bus(Freq)   | SDA/SCL Pins | AltFn | Alias |
	// #===========#==========#==============#==============#=======#=======#
	// |   I2C1    |   I2C1   | APB1(42 MHz) |   D14/D15    |   4   |   ~   |
	// |   I2C2    |   I2C2   | APB1(42 MHz) |    D0/D1     |   4   |   ~   |
	// |   I2C3    |   I2C1   | APB1(42 MHz) |    D9/D10    |   4   |   ~   |
	// | --------- | -------- | ------------ | ------------ | ----- | ----- |
	// |   I2C0    |   I2C1   | APB1(42 MHz) |   D14/D15    |   4   | I2C1  |
	// #===========#==========#==============#==============#=======#=======#
	NUM_I2C_INTERFACES	= 3

	I2C1_SDA_PIN	= D14	// I2C1 = hardware: I2C1
	I2C1_SCL_PIN	= D15	//

	I2C2_SDA_PIN	= D0	// I2C2 = hardware: I2C2
	I2C2_SCL_PIN	= D1	//

	I2C3_SDA_PIN	= D9	// I2C3 = hardware: I2C1
	I2C3_SCL_PIN	= D10	//   (interface duplicated on second pair of pins)

	I2C0_SDA_PIN	= I2C1_SDA_PIN	// I2C0 = alias: I2C1
	I2C0_SCL_PIN	= I2C1_SCL_PIN	//

	I2C_SDA_PIN	= I2C0_SDA_PIN	// default/primary I2C pins
	I2C_SCL_PIN	= I2C0_SCL_PIN	//
)
```



```go
const (
	TWI_FREQ_100KHZ	= 100000
	TWI_FREQ_400KHZ	= 400000
)
```

TWI_FREQ is the I2C bus speed. Normally either 100 kHz, or 400 kHz for high-speed bus.


```go
const NoPin = Pin(0xff)
```

NoPin explicitly indicates "not a pin". Use this pin if you want to leave one
of the pins in a peripheral unconfigured (if supported by the hardware).


```go
const (
	DutyCycle2	= 0
	DutyCycle16x9	= 1
)
```

I2C fast mode (Fm) duty cycle


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
)
```



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

	PH0	= portH + 0
	PH1	= portH + 1
)
```



```go
const (
	AF0_SYSTEM			= 0
	AF1_TIM1_2			= 1
	AF2_TIM3_4_5			= 2
	AF3_TIM8_9_10_11		= 3
	AF4_I2C1_2_3			= 4
	AF5_SPI1_SPI2			= 5
	AF6_SPI3			= 6
	AF7_USART1_2_3			= 7
	AF8_USART4_5_6			= 8
	AF9_CAN1_CAN2_TIM12_13_14	= 9
	AF10_OTG_FS_OTG_HS		= 10
	AF11_ETH			= 11
	AF12_FSMC_SDIO_OTG_HS_1		= 12
	AF13_DCMI			= 13
	AF14				= 14
	AF15_EVENTOUT			= 15
)
```

Alternative peripheral pin functions


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
		Buffer:			NewRingBuffer(),
		Bus:			stm32.USART3,
		AltFuncSelector:	AF7_USART1_2_3,
	}
	UART2	= UART{
		Buffer:			NewRingBuffer(),
		Bus:			stm32.USART6,
		AltFuncSelector:	AF8_USART4_5_6,
	}
	UART3	= UART{
		Buffer:			NewRingBuffer(),
		Bus:			stm32.USART1,
		AltFuncSelector:	AF7_USART1_2_3,
	}
	UART0	= UART1
)
```



```go
var (
	SPI1	= SPI{
		Bus:			stm32.SPI2,
		AltFuncSelector:	AF5_SPI1_SPI2,
	}
	SPI2	= SPI{
		Bus:			stm32.SPI3,
		AltFuncSelector:	AF6_SPI3,
	}
	SPI3	= SPI{
		Bus:			stm32.SPI1,
		AltFuncSelector:	AF5_SPI1_SPI2,
	}
	SPI0	= SPI1
)
```



```go
var (
	I2C1	= I2C{
		Bus:			stm32.I2C1,
		AltFuncSelector:	AF4_I2C1_2_3,
	}
	I2C2	= I2C{
		Bus:			stm32.I2C2,
		AltFuncSelector:	AF4_I2C1_2_3,
	}
	I2C3	= I2C{
		Bus:			stm32.I2C1,
		AltFuncSelector:	AF4_I2C1_2_3,
	}
	I2C0	= I2C1
)
```



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
	ErrTxInvalidSliceSize		= errors.New("SPI write and read slices must be same size")
	errSPIInvalidMachineConfig	= errors.New("SPI port was not configured properly by the machine")
)
```






### func CPUFrequency

```go
func CPUFrequency() uint32
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





## type I2C

```go
type I2C struct {
	Bus		*stm32.I2C_Type
	AltFuncSelector	uint8
}
```




### func (I2C) Configure

```go
func (i2c I2C) Configure(config I2CConfig) error
```

Configure is intended to setup the STM32 I2C interface.


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
	DutyCycle	uint8
}
```

I2CConfig is used to store config info for I2C.





## type PWM

```go
type PWM struct {
	Pin Pin
}
```






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
	Bus		*stm32.SPI_Type
	AltFuncSelector	uint8
}
```




### func (SPI) Configure

```go
func (spi SPI) Configure(config SPIConfig)
```

Configure is intended to setup the STM32 SPI1 interface.


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
	Bus		*stm32.USART_Type
	Interrupt	interrupt.Interrupt
	AltFuncSelector	uint8

	// Registers specific to the chip
	rxReg		*volatile.Register32
	txReg		*volatile.Register32
	statusReg	*volatile.Register32
	txEmptyFlag	uint32
}
```

UART representation



### func (UART) Buffered

```go
func (uart UART) Buffered() int
```

Buffered returns the number of bytes currently stored in the RX buffer.


### func (*UART) Configure

```go
func (uart *UART) Configure(config UARTConfig)
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


### func (*UART) SetBaudRate

```go
func (uart *UART) SetBaudRate(br uint32)
```

SetBaudRate sets the communication speed for the UART. Defer to chip-specific
routines for calculation


### func (UART) Write

```go
func (uart UART) Write(data []byte) (n int, err error)
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






