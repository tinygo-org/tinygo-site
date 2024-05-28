---
title: "Blinking LED"
description: "Hello world of hardware. Blink an on-board LED using your own development board."
weight: 1
---

This is the classic "blinking LED" demo you will find in many places. You can find it for the [Arduino IDE](https://docs.arduino.cc/built-in-examples/basics/Blink/) for example. Here is the equivalent in Go.

If you have one of the supported boards (see the dropdown below the code), you can run this code on your own hardware! If you don't, you can just use the simulator that's already running.

Going through the code we can see the following:

```go
package main
```

This is just standard Go, so we begin with package main.

```go
import (
	"machine"
	"time"
)
```

We import two packages. The time package is the well-known time package which we need to sleep while blinking. However the machine package is new: it is a special standard library package used by TinyGo.

You can think of the machine package like the os or syscall package: it provides an abstraction layer over the underlying hardware and makes code more portable across boards.

```go
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
```

This configures the LED pin. The constant `LED` is defined in the machine package and represents the pin that is connected to the on-board LED. It is of type `machine.Pin`.

We configure the pin here as an output pin, that is, it outputs either a high or low signal. _High_ or _low_ in this case mean that the output voltage is either somewhere close to the supply voltage (usually 5V or 3.3V) or close to ground (0V).

```go
	for {
		led.High()
		time.Sleep(time.Second/2)

		led.Low()
		time.Sleep(time.Second/2)
	}
```

Here we just blink the LED: we set the output to high for ½ second and then set it to low for ½ second.

Try changing these values and see what happens! For example you can set both to 1 second, or only one of them.

When you play around with the sleep duration, you might discover that "high" and "low" don't actually match the "on" and "off" state of the LED. "high" usually means that the LED is off, while "low" means that it is on. This is because of the way the LED is wired: the anode is connected to the power source (probably via a resistor), and the cathode is connected to the microcontroller pin. This means that when the pin is high (close to the power source) both sides are positive and no current will flow. And when the pin is negative, current will flow from the power source through the LED into the microcontroller. Some boards are wired differently however, and "high" will mean the LED is on while "low" is off.

<script type="module">
import { setupTour } from '/tour.js';
setupTour({
	boards: {
		'arduino-nano33': {},
		'arduino': {},
		'circuitplay-bluefruit': {},
		'circuitplay-express': {},
		'pico': {},
	},
	code: `
// You can edit this code!
package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.High()
		time.Sleep(time.Second/2)

		led.Low()
		time.Sleep(time.Second/2)
	}
}`});
</script>
