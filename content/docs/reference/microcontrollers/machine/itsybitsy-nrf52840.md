
---
title: itsybitsy-nrf52840
---


## Constants

```go
const HasLowFrequencyCrystal = true
```



```go
const (
	D0	= P0_25	// UART TX
	D1	= P0_24	// UART RX
	D2	= P1_02
	D3	= P0_06	// LED1
	D4	= P0_29	// Button
	D5	= P0_27
	D6	= P1_09	// DotStar Clock
	D7	= P1_08
	D8	= P0_08	// DotStar Data
	D9	= P0_07
	D10	= P0_05
	D11	= P0_26
	D12	= P0_11
	D13	= P0_12
	D14	= P0_04	// A0
	D15	= P0_30	// A1
	D16	= P0_28	// A2
	D17	= P0_31	// A3
	D18	= P0_02	// A4
	D19	= P0_03	// A5
	D20	= P0_05	// A6
	D21	= P0_16	// I2C SDA
	D22	= P0_14	// I2C SCL
	D23	= P0_20	// SPI SDI
	D24	= P0_15	// SPI SDO
	D25	= P0_13	// SPI SCK
	D26	= P0_19	// QSPI SCK
	D27	= P0_23	// QSPI CS
	D28	= P0_21	// QSPI Data 0
	D29	= P0_22	// QSPI Data 1
	D30	= P1_00	// QSPI Data 2
	D31	= P0_17	// QSPI Data 3
)
```

GPIO Pins


```go
const (
	A0	= D14
	A1	= D15
	A2	= D16
	A3	= D17
	A4	= D18
	A5	= D19
	A6	= D20
)
```

Analog Pins


```go
const (
	LED	= D3
	LED1	= LED
	BUTTON	= D4

	QSPI_SCK	= D26
	QSPI_CS		= D27
	QSPI_DATA0	= D28
	QSPI_DATA1	= D29
	QSPI_DATA2	= D30
	QSPI_DATA3	= D31
)
```



```go
const (
	UART_RX_PIN	= D0
	UART_TX_PIN	= D1
)
```

UART0 pins (logical UART1)


```go
const (
	SDA_PIN	= D21	// I2C0 external
	SCL_PIN	= D22	// I2C0 external
)
```

I2C pins


```go
const (
	SPI0_SCK_PIN	= D25
	SPI0_SDO_PIN	= D24
	SPI0_SDI_PIN	= D23
)
```

SPI pins


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
	PinInput		PinMode	= (nrf.GPIO_PIN_CNF_DIR_Input << nrf.GPIO_PIN_CNF_DIR_Pos) | (nrf.GPIO_PIN_CNF_INPUT_Connect << nrf.GPIO_PIN_CNF_INPUT_Pos)
	PinInputPullup		PinMode	= PinInput | (nrf.GPIO_PIN_CNF_PULL_Pullup << nrf.GPIO_PIN_CNF_PULL_Pos)
	PinInputPulldown	PinMode	= PinInput | (nrf.GPIO_PIN_CNF_PULL_Pulldown << nrf.GPIO_PIN_CNF_PULL_Pos)
	PinOutput		PinMode	= (nrf.GPIO_PIN_CNF_DIR_Output << nrf.GPIO_PIN_CNF_DIR_Pos) | (nrf.GPIO_PIN_CNF_INPUT_Connect << nrf.GPIO_PIN_CNF_INPUT_Pos)
)
```



```go
const (
	PinRising	PinChange	= nrf.GPIOTE_CONFIG_POLARITY_LoToHi
	PinFalling	PinChange	= nrf.GPIOTE_CONFIG_POLARITY_HiToLo
	PinToggle	PinChange	= nrf.GPIOTE_CONFIG_POLARITY_Toggle
)
```

Pin change interrupt constants for SetInterrupt.


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
	DefaultUART = UART0
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
	// UART0 is the hardware UART on the NRF SoC.
	_UART0	= UART{Buffer: NewRingBuffer()}
	UART0	= &_UART0
)
```

UART


```go
var (
	I2C0	= (*I2C)(unsafe.Pointer(nrf.TWI0))
	I2C1	= (*I2C)(unsafe.Pointer(nrf.TWI1))
)
```

There are 2 I2C interfaces on the NRF.


```go
var (
	PWM0	= &PWM{PWM: nrf.PWM0}
	PWM1	= &PWM{PWM: nrf.PWM1}
	PWM2	= &PWM{PWM: nrf.PWM2}
	PWM3	= &PWM{PWM: nrf.PWM3}
)
```

PWM


```go
var (
	SPI0	= SPI{Bus: nrf.SPIM0, buf: new([1]byte)}
	SPI1	= SPI{Bus: nrf.SPIM1, buf: new([1]byte)}
	SPI2	= SPI{Bus: nrf.SPIM2, buf: new([1]byte)}
)
```

There are 3 SPI interfaces on the NRF528xx.


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

EnterBootloader resets the chip into the UF2 bootloader. After reset, it
can be flashed via nrfutil or by copying a UF2 file to the mass storage device


### func EnterOTABootloader

```go
func EnterOTABootloader()
```

EnterOTABootloader resets the chip into the bootloader so that it can be
flashed via an OTA update


### func EnterSerialBootloader

```go
func EnterSerialBootloader()
```

EnterSerialBootloader resets the chip into the serial bootloader. After
reset, it can be flashed using serial/nrfutil.


### func EnterUF2Bootloader

```go
func EnterUF2Bootloader()
```

EnterUF2Bootloader resets the chip into the UF2 bootloader. After reset, it
can be flashed via nrfutil or by copying a UF2 file to the mass storage device


### func GetRNG

```go
func GetRNG() (ret uint32, err error)
```

GetRNG returns 32 bits of non-deterministic random data based on internal thermal noise.
According to Nordic's documentation, the random output is suitable for cryptographic purposes.


### func InitADC

```go
func InitADC()
```

InitADC initializes the registers needed for ADC.


### func InitSerial

```go
func InitSerial()
```



### func NewRingBuffer

```go
func NewRingBuffer() *RingBuffer
```

NewRingBuffer returns a new ring buffer.


### func ReadTemperature

```go
func ReadTemperature() int32
```

ReadTemperature reads the silicon die temperature of the chip. The return
value is in milli-celsius.


### func ReceiveUSBControlPacket

```go
func ReceiveUSBControlPacket() ([cdcLineInfoSize]byte, error)
```



### func SendUSBInPacket

```go
func SendUSBInPacket(ep uint32, data []byte) bool
```

SendUSBInPacket sends a packet for USBHID (interrupt in / bulk in).


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
func (a ADC) Configure(ADCConfig)
```

Configure configures an ADC pin to be able to read analog data.


### func (ADC) Get

```go
func (a ADC) Get() uint16
```

Get returns the current value of a ADC pin in the range 0..0xffff.




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
	Bus nrf.TWI_Type
}
```

I2C on the NRF.



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


### func (*I2C) Tx

```go
func (i2c *I2C) Tx(addr uint16, w, r []byte) (err error)
```

Tx does a single I2C transaction at the specified address.
It clocks out the given address, writes the bytes in w, reads back len(r)
bytes and stores them in r, and generates a stop condition on the bus.


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




## type PWM

```go
type PWM struct {
	PWM	*nrf.PWM_Type

	channelValues	[4]volatile.Register16
}
```

PWM is one PWM peripheral, which consists of a counter and multiple output
channels (that can be connected to actual pins). You can set the frequency
using SetPeriod, but only for all the channels in this PWM peripheral at
once.



### func (*PWM) Channel

```go
func (pwm *PWM) Channel(pin Pin) (uint8, error)
```

Channel returns a PWM channel for the given pin.


### func (*PWM) Configure

```go
func (pwm *PWM) Configure(config PWMConfig) error
```

Configure enables and configures this PWM.
On the nRF52 series, the maximum period is around 0.26s.


### func (*PWM) Set

```go
func (pwm *PWM) Set(channel uint8, value uint32)
```

Set updates the channel value. This is used to control the channel duty
cycle. For example, to set it to a 25% duty cycle, use:

	ch.Set(ch.Top() / 4)

ch.Set(0) will set the output to low and ch.Set(ch.Top()) will set the output
to high, assuming the output isn't inverted.


### func (*PWM) SetInverting

```go
func (pwm *PWM) SetInverting(channel uint8, inverting bool)
```

SetInverting sets whether to invert the output of this channel.
Without inverting, a 25% duty cycle would mean the output is high for 25% of
the time and low for the rest. Inverting flips the output as if a NOT gate
was placed at the output, meaning that the output would be 25% low and 75%
high with a duty cycle of 25%.


### func (*PWM) SetPeriod

```go
func (pwm *PWM) SetPeriod(period uint64) error
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


### func (*PWM) Top

```go
func (pwm *PWM) Top() uint32
```

Top returns the current counter top, for use in duty cycle calculation. It
will only change with a call to Configure or SetPeriod, otherwise it is
constant.

The value returned here is hardware dependent. In general, it's best to treat
it as an opaque value that can be divided by some number and passed to
pwm.Set (see pwm.Set for more information).




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
	Bus	*nrf.SPIM_Type
	buf	*[1]byte	// 1-byte buffer for the Transfer method
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

Tx handles read/write operation for SPI interface. Since SPI is a syncronous
write/read interface, there must always be the same number of bytes written
as bytes read. Therefore, if the number of bytes don't match it will be
padded until they fit: if len(w) > len(r) the extra bytes received will be
dropped and if len(w) < len(r) extra 0 bytes will be sent.




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






## type UART

```go
type UART struct {
	Buffer *RingBuffer
}
```

UART on the NRF.



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
type UARTParity uint8
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




