---
title: "machine package"
weight: 5
description: >
  How to interact with peripherals exposed through the `machine` package.
---

## GPIO


```go
type Pin uint8

const NoPin = Pin(0xff)
```

The pin type indicates a physical pin on a microcontroller. It can be used directly for GPIO or as part of a different peripheral (SPI, ADC, etc).

The `NoPin` value can be used in some places where you want to explicitly not use a pin. For example, depending on the chip you could use only SDO or SDI of an SPI peripheral, leaving the other pin unused.

Boards have various pins defined as constants. For example, ATsamd21 based chips (often marketed as "M0") have pin number PB01 defined as `PB01` in the machine package and the Arduino Uno has pin 13 defined as `D13`. However, please note that Arduino pin numbers do not necessarily match the pin numbers as used in TinyGo, therefore you should always use the named constants to refer to a pin if possible.

```go
type PinMode uint8

const (
    PinOutput        PinMode = ...
    PinInput         PinMode = ...
    PinInputPullup   PinMode = ...
    PinInputPulldown PinMode = ... // not always available
)
```

These `PinMode` are used to set a pin to various states: as input or output and optionally with a pull resistor (see above in the guide). The values these constants have vary by chip and should be considered implementation details.

Pull-down mode is not always available on the hardware. If it is not available, the `PinInputPulldown` constant is not available resulting in a compile error.

```go
type PinConfig struct {
    Mode PinMode
}
```

This struct contains options to configure a hardware pin. More options might be added in the future (with appropriate behavior on the zero value), but currently only the `Mode` field is used.

```go
func (p Pin) Configure(config machine.PinConfig)
```

Configure a pin, see `PinConfig` and `PinMode` above. `Configure` must be called before any other method may be called.

```go
func (p Pin) Low()
func (p Pin) High()
func (p Pin) Set(state bool)
```

Set an output pin to low or high. It must have been configured before as an output (`PinOutput`) or the behavior will be undefined. The `Set` method sets the state to low or high depending on the boolean value: `true` means high, `false` means low.

```go
func (p Pin) Get() bool
```

Get the input state of a pin. The returned value indicates whether the pin is low or high: `true` means high and `false` means low. `Get` may only be called when the pin has been configured as an input.

Note that if the pin is left floating (not connected to anything) the returned value is unpredictable and may appear random.
