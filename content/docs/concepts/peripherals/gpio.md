---
title: "GPIO"
weight: 1
description: |
  How to control pins directly using GPIO.
---

For API documentation, see the [machine API](../../../reference/machine#gpio).

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

`machine.LED` is a constant of type `machine.Pin`. It is simply a named numeric type (`type Pin uint8`) to identify a given pin on the board. It has a few methods to interact with these hardware pins. It can also be used in other peripherals, but more on that in other guides. Most developer boards contain an on-board LED which you can access through `machine.LED`, but you can also use any other pin defined on the board.

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

**Warning**: make sure the input voltage is within the allowed range of the chip. Many chips will tolerate an input slightly out of the VCC..GND range (for example, allowing VCC+0.3V) but it's better to stay entirely within VCC..GND to be safe. Going outside this range may stress or destroy the chip, although the damage might not be immediately noticeable.

It's also worth noting that digital inputs expect exactly VCC or GND as input, but have some tolerance: the cutoff isn't exactly in the middle (VCC\*0.5V). Instead, most chips guarantee that low will be read when the voltage is below VCC\*0.3 and high will be read when the voltage is above VCC\*0.7. The area in between could be read either way.

## Interrupts on pin changes

Sometimes you may want to respond right away to a pin change, for example when it is connected to a button. You can do this by reading the state very frequently (every 10ms for example) but that introduces a lot of overhead and still doesn't respond right away. Instead, you can set an interrupt callback on the pin, which will be called when the pin changes state.

Here is an example which can be used on a Circuit Playground Express:

```go
led := machine.LED
led.Configure(machine.PinConfig{Mode: machine.PinOutput})
led.Low()

pin := machine.BUTTONA
pin.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
pin.SetInterrupt(machine.PinToggle, func(p machine.Pin) {
    led.Set(p.Get())
})
```

In this case, the button is connected between the input pin (here, `machine.BUTTONA`) and VCC. By configuring it with a pull down register, the button pin is normally logically low. However, when you press the button a connection is made between the button pin and VCC, which makes it logically high.

The most important part is this:

```go
pin.SetInterrupt(machine.PinToggle, func(p machine.Pin) {
    led.Set(p.Get())
})
```

It says that it should trigger an interrupt whenever the pin changes state (from low to high or from high to low, meaning when pressing or releasing the button) and that it sets the LED output to the given state. This means that the LED output will match the button input in software. That's not very useful on its own but demonstrates how pin interrupts can be used.
