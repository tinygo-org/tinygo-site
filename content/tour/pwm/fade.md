---
title: "LED: fade"
description: "Fade a LED using the PWM peripheral."
weight: 2
---

Next up, we're going to do what PWM was actually designed for! Namely, quickly controlling an output so that the average output can be controlled precisely.

Most of it is similar to blinking an LED using the PWM peripheral, but fading needs a little bit more work:

```go
		// Fade the LED in.
		for percentOn := 0; percentOn <= 100; percentOn++ {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(time.Second / 100)
		}
```

Here we fade the LED in. This loop loops 101 times (from 0 to 100 inclusive), setting the output to a percentage of the input. The code `pwm.Top()*uint32(percentOn)/100` is essentially the integer variant of the following formula:

```math
\frac{T}{100} * P
```

Where \\(T\\) is the top value (`pwm.Top()`) and \\(P\\) is the percentage (`percentOn`).

Fading out is very similar. It's almost identical, except it loops from 100 to 0 (inclusive):

```go
		// Fade the LED out.
		for percentOn := 100; percentOn >= 0; percentOn-- {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(time.Second / 100)
		}
```


<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import "machine"
import "time"

func main() {
	// Board specific configuration.
	led := machine.LED
	pwm := machine.TCC0

	// Configure the timer/PWM.
	// Use the default period, which works well enough for LEDs.
	err := pwm.Configure(machine.PWMConfig{})
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

	for {
		// Fade the LED in.
		for percentOn := 0; percentOn <= 100; percentOn++ {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(time.Second / 100)
		}

		// Fade the LED out.
		for percentOn := 100; percentOn >= 0; percentOn-- {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(time.Second / 100)
		}
	}
}
`;
setupTour({
	code: code,
	boards: {
		'arduino-nano33': {},
		'circuitplay-bluefruit': {
			code: code.replace('machine.TCC0', 'machine.PWM0'),
		},
		'circuitplay-express': {},
		'pico': {
			code: code.replace('machine.TCC0', 'machine.PWM4'),
		},
	}});
</script>
