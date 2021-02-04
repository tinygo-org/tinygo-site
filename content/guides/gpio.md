---
title: "GPIO"
weight: 1
---

GPIO is the most basic form of interacting with the outside world on a microcontroller. It makes it possible to set a pin to low (connected to ground) or high (connected to VCC, usually 3.3V or 5V). GPIO also allows reading whether a pin is low or high.

## Controlling GPIO output

You might have seen the classic blinking LED example before, but here it is again:

```go
package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        time.Sleep(time.Millisecond * 1000)

        led.High()
        time.Sleep(time.Millisecond * 1000)
    }
}
```

This example continuously sets an output (connected to an LED) to on or off.

Let's go through it line by line.

```go
led := machine.LED
led.Configure(machine.PinConfig{Mode: machine.PinOutput})
```

`machine.LED` is a constant of type `machine.Pin`. It is simply a named numeric type (`type Pin uint8`) to identify a given pin on the board. It has a few methods to interact with these hardware pins. It can also be used in other peripherals, but more on that in other guides. Most developer boards contain an on-board LED which you can access through `machine.LED`, but you can also us any other pin defined on the board.

The first method that's important is the `Configure` method. All hardware in the TinyGo machine package needs some initialization first before it can be used. In this case, only a single setting needs to be configured: the `Mode` field which indicates the pin mode. You could compare the `Configure` method to the Arduino [`pinMode`](https://www.arduino.cc/reference/en/language/functions/digital-io/pinmode/) function.

After configuration, you can set it to either low or high:

```go
led.Low()
led.High()
```

Or you can set it to a binary (`bool`) value, where low is false and high is true:

```go
led.Set(true) // same as led.High()
```

Again, low means that it is connected to ground while high means that it is connected to VCC, which is usually 3.3V or 5V.

**Warning**: make sure to limit the amount of current you draw from these pins. It is usually fine to drive a LED directly as long as you use a resistor (330Ω is safe in most cases) but you usually cannot drive anything more powerful directly, especially something as powerful as electromotors. Instead, you should look into transistors and MOSFETs (hint: MOSFETs are often easier to work with as they have a digital input).

You might notice that setting an onboard LED to low often turns it on while setting it to high turns it off. This is because LEDs on developer boards are often connected in a somewhat counter intuitive way: with the anode connected directly to VCC and the cathode connected to the output pin (with a resistor in between somewhere, of course). I believe this configuration is historical: chips might be able to "source" (provide) more current on VCC than they can "sink" current towards GND.

## Reading GPIO input

You can also read binary input from the outside world. This needs a slightly different configuration:

```go
pin := machine.BUTTON
pin.Configure(machine.PinConfig{Mode: machine.PinInput})
```

This configures the pin as a GPIO input. There are two other modes (`machine.PinInputPullup` and `machine.PinInputPulldown`) that will be covered later in this guide, but for now this will be good enough.

After configuring, you can read the current state of the pin using the `Get` method:

```go
value := pin.Get()
```

Like before with `Set`, `true` means high while `false` means low. So if the pin is connected to VCC (directly or through a resistor) you will read `true` while if the pin is connected to GND you'll read `false`.

This might raise the question: what if the pin is not connected to anything? In that case it is in an intermediate state called "floating". What you will read is entirely unpredictable, it might be either high or low and this state might randomly change (or not) even by subtle changes such as moving your hand near the pin. This state is generally undetectable. Therefore, you should always avoid floating inputs if possible.

The way to avoid floating inputs is through pull-up or pull-down resistors. You can read more about them [here](https://learn.sparkfun.com/tutorials/pull-up-resistors/all). In short, it connects a resistor with relatively high resistance to VCC or GND, usually somewhere around 10kΩ. You can think of it as pulling the input signal in a certain direction: a pull-up resistor gently pulls the input signal towards VCC while a pull-down resistor gently pulls down the signal to GND - all to avoid letting the input signal float in the middle and causing all kinds of issues. Of course, this pull is quite gentle so if you also connect the input directly to GND or VCC (for example with a button) the signal will still change, while avoiding the floating middle state.

You can set the pull mode as part of the configure step:

```go
pin := machine.BUTTON
pin.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
```

This simply adds an extra resistor from the pin to VCC - all in software! There is also `machine.PinInputPulldown` but it is slightly less commonly supported in microcontrollers so if it doesn't matter which way you pull, it's generally a good idea to use a pull up rather than a pull down.

Of course, instead of using a pull mode in software you can also connect a physical resistor to the pin as a pull up or pull down resistor. This has the advantage that you precisely control the resistor value.

```Warning```: make sure the input voltage is within the allowed range of the chip. Many chips will tolerate an input slightly out of the VCC..GND range (for example, allowing VCC+0.3V) but it's better to stay entirely within VCC..GND to be safe. Going outside this range may stress or destroy the chip, although the damage might not be immediately noticeable.

It's also worth noting that digital inputs expect exactly VCC or GND as input, but have some tolerance: the cutoff isn't exactly in the middle (VCC\*0.5V). Instead, most chips guarantee that low will be read when the voltage is below VCC\*0.3 and high will be read when the voltage is above VCC\*0.7. The area in between could be read either way.

## `machine` API reference

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
