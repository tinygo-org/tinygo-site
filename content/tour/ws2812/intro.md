---
title: "Introduction"
description: How to use the extremely common WS2812/NeoPixel addressable LEDs.
weight: 1
---

You've probably seen or used these LEDs before. They're extremely common in the maker scene, for a good reason! They're cheap, widely available, and easy to use. You can control them from basically any microcontroller with just a single pin and a good power supply.

On the right you can see how to control them. Some boards include WS2812 boards, so you can use those without fiddling with wires.

```go
ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
```

Before being able to send any data, the pin needs to be set as an output. This is similar to how you'd control a single LED.

```go
ws := ws2812.New(ledPin)
```

The `ws` object is the driver for these LEDs and is later used to update the LEDs.


```go
leds := []color.RGBA{
	{R: 255, G: 0, B: 0},
	{R: 0, G: 255, B: 0},
	{R: 0, G: 0, B: 255},
}
```

This initializes a slice of RGB values to be sent to the LED strip. The first is red, the second green, and the third blue.

```go
ws.WriteColors(leds)
```

And here we send the LED colors to the LED strip. Easy, right?

Though if you look at the output of the simulator, you can see that it is actually reversed: the first LED is blue, the second green, and the third red. What happened here?

## WS2812 daisy chaining

To understand why the LEDs are reversed, we need to understand a bit of how the protocol works.

For the first LED, the microcontroller sends the first color in the array to the first LED. This means the first LED receives the red color. It doesn't apply the color immediately, instead it keeps the color in a buffer and waits to see whether more will be coming.

But the microcontroller doesn't wait, it sends the next color (green) to the first LED. The first LED already has the color red in the buffer, so it sends the color red to the next (second) LED while receiving the green color.

Now, the third color (blue) is sent. Like before, the first LED receives blue and sends the green it has in the buffer. The second LED receives green, and sends the red color to the third LED.

Now the microcontroller is finished, it doesn't send anything else. In fact, the main function returns so it locks up. After a while (50µs according to the datasheet, [around 9µs in practice](https://cpldcpu.wordpress.com/2014/01/14/light_ws2812-library-v2-0-part-i-understanding-the-ws2812/)) all the LEDs in the string realize no more data is coming so they apply the color they have in their buffer.

Knowing this we can do interesting things with these LEDs:

  * It's possible to daisy chain many of these LED strips! They pass color data from one to the other, and so connecting multiple in a chain means you get one big LED strip. Though watch out for power consumption, this'll quickly add up!
  * You can, if you want, control multiple LED strips in parallel. They will look the same, but it does save a pin on the microcontroller.
  * You can update multiple chained LED strips independently! Just send them all together without any delay in between (from last to first). This can be useful for controlling chained RGB and RGBW LED strips, if you happen to have such a setup. We'll get to RGBW LED strips later.

## Connecting WS2812 LEDs

The simulator simplifies a lot of things. Sadly, the real world is a bit more complicated.

There are extensive guides on the internet on how to connect these LEDs. [Here is one by Adafruit](https://learn.adafruit.com/adafruit-neopixel-uberguide/the-magic-of-neopixels), and [here is one by SparkFun](https://learn.sparkfun.com/tutorials/ws2812-breakout-hookup-guide/all). But in short, here are a few things you need to be aware of:

  * The power wire (usually red) needs to be connected to a decent power supply. For just a few LEDs you can usually use a pin on the microcontroller board itself, this one is usually labeled "5V", "VBUS", or "VOUT". Check the pinout to be sure! For longer LED strips you will need a separate 5V power supply.
  * The ground wire (usually white or black) needs to be connected to the microcontroller ground. If you use a separate power supply, it also needs to be connected to the ground of that power supply.
  * The data pin (usually green) needs to be connected to the GPIO pin you are using on the board.
  * Make sure the wires are not too long! One meter will usually work fine, but for longer wires you may need special handling.
  * These LEDs can be very bright, especially in the dark! Many people don't run them at full brightness for this reason.


## WS2812 vs NeoPixel vs SK6812 vs ...

These are all the same type of LEDs! There are multiple companies that make these, but apart from some details these LEDs are all interchangeable. In this tour, we'll just use the name "WS2812" for the whole family.

WS2812 was the original by WorldSemi (hence WS), followed by the WS2812B which made it slightly more robust. SK6812 is a very similar chip from a different company (known for making a RGBW version of the WS2812). There are many others that are usually just sold as "WS2812" even though they aren't from WorldSemi.

NeoPixel is the Adafruit brand for this type of LED. They're still the same LEDs (from WorldSemi and others), just with a different name.


<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

const ledPin = machine.WS2812

func main() {
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws := ws2812.New(ledPin)
	leds := []color.RGBA{
		{R: 255, G: 0, B: 0},
		{R: 0, G: 255, B: 0},
		{R: 0, G: 0, B: 255},
	}
	ws.WriteColors(leds)
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
