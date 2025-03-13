---
title: "Animating LEDs"
description: How to animate LEDs using a main loop.
weight: 2
---

Manually setting the color of each LED is nice, but of course it's much more interesting to animate them! Of course that's possible with a bit more code, see the code and demo on the right.

There are many ways to code an animation, but one that works well in many cases is by using a main loop that updates the in-memory color values, sends them out to the LED strip, and sleeps for a bit. Let's go through the code:

```go
	// Configure LEDs.
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws := ws2812.New(ledPin)
	leds := make([]color.RGBA, 10)
```

This configures the LEDs as before. One difference is that we create a slice of color values in advance, without setting their initial values. But since all color components get initialized to zero, this effectively creates an all-black color slice.

```go
	red := false
	for {
		// ...
	}
```

This is the main loop of the program. It's an infinite loop (much like the `loop` function on the Arduino platform) that controls the LEDs infinitely - or, more likely, until you pull the power.

Note also that it initializes the `red` variable to false, which is later used to toggle the color each cycle.

```go
		// Update the in-memory LED array.
		red = !red
		for i := range leds {
			red = !red
			if red {
				leds[i] = color.RGBA{R: 0xff}
			} else {
				leds[i] = color.RGBA{G: 0xff}
			}
		}
```

This is the first step of the loop. All values in the LED array get set to a specific color (red or green), toggling the color for each LED.

```go
		// Send the new colors to the LED strip.
		ws.WriteColors(leds)
```

Like before, this is where the LED strip gets updated.

```go
		// Wait a bit until the next update.
		time.Sleep(time.Second / 2)
```

This sleeps for a bit, so that the LED animation looks correct. In fact, this sleep is necessary for the correct functioning of the LED strip! Because the WS2812 protocol only knows when it has received all data when there's a short delay (up to 50µs) in the received data, we need to sleep for at least a short while.


<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

const ledPin = machine.WS2812

func main() {
	// Configure LEDs.
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws := ws2812.New(ledPin)
	leds := make([]color.RGBA, 10)

	red := false
	for {
		// Update the in-memory LED array.
		red = !red
		for i := range leds {
			red = !red
			if red {
				leds[i] = color.RGBA{R: 0xff}
			} else {
				leds[i] = color.RGBA{G: 0xff}
			}
		}

		// Send the new colors to the LED strip.
		ws.WriteColors(leds)

		// Wait a bit until the next update.
		time.Sleep(time.Second / 2)
	}
}
`;
let stickConfig = {
	type: 'ws2812',
	humanName: 'NeoPixel Stick 8',
	svg: 'ws2812-stick-8.svg',
	length: 8,
};
setupTour({
	boards: {
		'circuitplay-bluefruit': {},
		'circuitplay-express': {},
		'gopher-badge': {},
		'arduino': {
			code: code.replace('machine.WS2812', 'machine.D7'),
			parts: {
				main: {
					x: -30,
				},
				stick: {
					config: stickConfig,
					x: 40,
				}
			},
			wires: [
				{from: 'main.D7', to: 'stick.din'},
			],
		},
		'arduino-nano33': {
			code: code.replace('machine.WS2812', 'machine.A7'),
			parts: {
				main: {
					x: -30,
				},
				stick: {
					config: stickConfig,
					x: 40,
				}
			},
			wires: [
				{from: 'main.A7', to: 'stick.din'},
			],
		},
		'microbit': {
			code: code.replace('machine.WS2812', 'machine.P2'),
			parts: {
				main: {
					x: -30,
				},
				stick: {
					config: stickConfig,
					x: 40,
				}
			},
			wires: [
				{from: 'main.P2', to: 'stick.din'},
			],
		},
		'pico': {
			code: code.replace('machine.WS2812', 'machine.GP16'),
			parts: {
				main: {
					x: -30,
				},
				stick: {
					config: stickConfig,
					x: 40,
				}
			},
			wires: [
				{from: 'main.GP16', to: 'stick.din'},
			],
		},
	},
	code: code});
</script>
