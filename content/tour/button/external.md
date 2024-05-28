---
title: "External button"
description: Use an external push button
weight: 2
---

This time we'll be blinking an LED, but only when the button is pressed.

If you try it out, you'll see that the LED is initially off (or on), but when you press the button it'll quickly blink. When you release the button, the blinking will stop and it will keep the last state (either on or off).

If you have one of these boards but do not have a pushbutton, you can also use a small wire to connect the given pin (D2/A2/GP2 depending on the board) and then touch a ground point, such as the outside of the USB connector.


```go
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
```

This configures the on-board LED, same as in the [blink example]({{< ref "../blink/onboard.md" >}}).

```go
	button := machine.D2 // Arduino Uno
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
```

This configures the button input. The button pin varies by board, but this time it's always a pullup.

```go
		// if the button is low (pressed)
		if !button.Get() {
			// ...
		}
```

Here we check whether the button input is _low_, in other words, whether it is pressed. This is a bit counter intuitive, but that's how it is.

```go
			// toggle the LED
			led.Set(!led.Get())
```

Here we use a trick to toggle the LED: we can actually read the current output value using `led.Get()` and then we can change the value to the opposite value using `led.Set`. This is equivalent to the following:

```go
			if led.Get() {
				led.Low()
			} else {
				led.High()
			}
```

<script type="module">
import { setupTour } from '/tour.js';
let buttonConfig = {
	type: 'pushbutton',
	humanName: 'Push button',
	svg: 'pushbutton.svg',
};
let code = `
package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := EXT_BUTTON
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	for {
		// if the button is low (pressed)
		if !button.Get() {
			// toggle the LED
			led.Set(!led.Get())
		}

		// wait a bit, for the blinking effect
		time.Sleep(100 * time.Millisecond)
	}
}`;
setupTour({
	boards: {
		'arduino': {
			code: code.replace('EXT_BUTTON', 'machine.D2 // Arduino Uno'),
			parts: {
				button: {
					config: buttonConfig,
					x: 45,
					y: -10,
					rotation: 90,
				},
			},
			wires: [
				{from: 'main.D2', to: 'button.A'},
				{from: 'main.GND#2', to: 'button.B'},
			],
		},
		'arduino-nano33': {
			code: code.replace('EXT_BUTTON', 'machine.D2 // Arduino Nano 33 IoT'),
			parts: {
				button: {
					config: buttonConfig,
					x: 12,
					y: -18,
				}
			},
			wires: [
				{from: 'main.D2', to: 'button.A'},
				{from: 'main.GND#1', to: 'button.B'},
			],
		},
		'circuitplay-bluefruit': {
			code: code.replace('EXT_BUTTON', 'machine.A2 // Circuit Playground'),
			parts: {
				button: {
					config: buttonConfig,
					x: 35,
					y: -4,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.A2', to: 'button.A'},
				{from: 'main.GND#1', to: 'button.B'},
			],
		},
		'circuitplay-express': {
			code: code.replace('EXT_BUTTON', 'machine.A2 // Circuit Playground'),
			parts: {
				button: {
					config: buttonConfig,
					x: 35,
					y: -4,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.A2', to: 'button.A'},
				{from: 'main.GND#1', to: 'button.B'},
			],
		},
		'pico': {
			code: code.replace('EXT_BUTTON', 'machine.GP2 // Raspberry Pi Pico'),
			parts: {
				button: {
					config: buttonConfig,
					x: -11,
					y: 18,
				}
			},
			wires: [
				{from: 'main.GP2', to: 'button.A'},
				{from: 'main.GND#6', to: 'button.B'},
			],
		},
	},
});
</script>
