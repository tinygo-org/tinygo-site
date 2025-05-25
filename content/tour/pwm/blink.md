---
title: "LED: blink"
description: "Get a LED to blink using the PWM peripheral."
weight: 1
---

PWM, or pulse-width modulation, is a method of turning a GPIO pin high and low at a fixed frequency but changing the percent of time it is turned on. This has numerous applications, such as dimming LEDs. We'll see various applications of PWM throughout this tour.

To start off with, we're going to do something slightly unusual with PWM: we're going to blink an LED. Remember that PWM is a way of quickly turning an output high and low, if we do that slow enough we can see the LED blink. This should make it easier to see how PWM works.

At first we'll just define some variables. These may vary by board, so it's useful to define them in a single place.

```go
	// Board specific configuration.
	led := machine.LED
	pwm := machine.TCC0
```

Next we're going to configure the PWM peripheral itself:

```go
	// Configure the timer/PWM.
	err := pwm.Configure(machine.PWMConfig{
		Period: 2e9, // two seconds for a single cycle
	})
	if err != nil {
		println("could not configure:", err.Error())
		return
	}
```

This configures the PWM peripheral for a cycle time of two seconds. That's much longer than PWM is commonly used for, but it allows to see us what happens.

Next up, we connect the PWM peripheral to the given output pin:

```go
	// Get the channel for this PWM peripheral.
	ch, err := pwm.Channel(led)
	if err != nil {
		println("could not obtain channel:", err.Error())
		return
	}
```

We get a channel number back. One PWM peripheral typically has 2-4 channels, which can be connected to certain pins. Which pins can be used depends a lot on the hardware, which we'll [get to later](../multiple/#finding-the-correct-pwm-peripheral-and-pins). For now, all you need to know is that the PWM channel is now connected to a single GPIO pin.

To actually make the LED blink, we will need to set the duty cycle. The duty cycle refers to the part of the time when the PWM is turned on. In this case, we will set the duty cycle to half the "top" value of the LED:

```go
	// Blink the LED, setting it to "on" half the time.
	pwm.Set(ch, pwm.Top()/2)
```

The top value is equivalent to a duty cycle of 100% - or 100% of the time on. So using the top divided by 2 results in a duty cycle of 50%, or "high" for half the time. Dividing by 3 results in a duty cycle of 33%, etc. By varying the formula you can either use a particular percentage (like we do here) or if you know the period size you can calculate a specific "on" time within the cycle (such as "500ms on, 2000ms off").

To get a feel for how it works, you can try a few things:

  * You can change the period to something other than `2e9` (two seconds) and see what happens to the "top" value.
  * You can change the duty cycle by changing the `pwm.Set` call.
  * You can connect a different LED to the virtual board, and blink it.
  * You can try using a different PWM peripheral. These boards also have `TCC1` and `TCC2`.

## More information

SparkFun has written an [excellent tutorial on PWM](https://learn.sparkfun.com/tutorials/pulse-width-modulation) that is well worth a read!



<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import "machine"

func main() {
	// Board specific configuration.
	led := machine.LED
	pwm := machine.TCC0

	// Configure the timer/PWM.
	err := pwm.Configure(machine.PWMConfig{
		Period: 2e9, // two seconds for a single cycle
	})
	if err != nil {
		println("could not configure:", err.Error())
		return
	}

	// Get the channel for this PWM peripheral.
	ch, err := pwm.Channel(led)
	if err != nil {
		println("could not obtain channel:", err.Error())
		return
	}

	// Blink the LED, setting it to "on" half the time.
	pwm.Set(ch, pwm.Top()/2)

	println("top value:  ", pwm.Top())
	println("LED channel:", ch)

	// Main returns, so the program exits. Yet the LED continues blinking.
}
`;
setupTour({
	code: code,
	boards: {
		'arduino-nano33': {},
		'circuitplay-express': {},
	}});
</script>
