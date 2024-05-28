---
title: "External LED"
description: "Same blinking LED as in the previous example, but now with an external LED."
weight: 2
---

Instead of using an on-board LED, you can also connect a LED to an external pin. Just like you see in the example on the right. It connects an external LED (a red one in this case) to the pin A1 on the board.

The code is almost identical to the previous example. The only difference is this line:

```go
	led := machine.A1
```

We've replaced the `led` variable: instead of the on-board LED we're now using the external pin A1.

If you want to try this on a development board, make sure you use the correct wiring:

  * Connect the LED cathode (short lead with flat side) to a resistor of around 220Ω-330Ω. Use at least 150Ω to protect the LED, and use at most 1000Ω to still be able to see the LED (higher values will make the LED very dim).
  * Connect the other end of the resistor to ground.
  * Connect the LED anode to pin A1 on the board. See the silkscreen on the board or the graphic in the simulator to know which pin this is.

**Warning**: the resistor is not optional, not using the resistor can destroy the LED and/or your board. The LED might even explode!

The simulation doesn't need this resistor because the virtual LED can't burn out. This simplifies the demo, but remember that the real world is messy and LEDs need a resistor. (There are technical reasons for this, read [this answer](https://electronics.stackexchange.com/questions/28393/why-do-we-need-resistors-in-led) for more details).

<script type="module">
import { setupTour } from '/tour.js';
let ledConfig = {
	type: 'led',
	humanName: 'External LED',
	color: [255, 0, 0],
	svg: 'led-tht-5mm.svg',
};
let code = `
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
}`;
setupTour({
	boards: {
		'arduino': {
			code: code.replace('machine.LED', 'machine.D2 // Arduino Uno'),
			parts: {
				led: {
					config: ledConfig,
					x: 45,
					y: -10,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.D2', to: 'led.anode'},
				{from: 'main.GND#2', to: 'led.cathode'},
			],
		},
		'arduino-nano33': {
			code: code.replace('machine.LED', 'machine.D2 // Arduino Nano 33 IoT'),
			parts: {
				led: {
					config: ledConfig,
					x: 12,
					y: -18,
				}
			},
			wires: [
				{from: 'main.D2', to: 'led.anode'},
				{from: 'main.GND#1', to: 'led.cathode'},
			],
		},
		'circuitplay-bluefruit': {
			code: code.replace('machine.LED', 'machine.A2 // Circuit Playground'),
			parts: {
				led: {
					config: ledConfig,
					x: 35,
					y: -4,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.A2', to: 'led.anode'},
				{from: 'main.GND#1', to: 'led.cathode'},
			],
		},
		'circuitplay-express': {
			code: code.replace('machine.LED', 'machine.A2 // Circuit Playground'),
			parts: {
				led: {
					config: ledConfig,
					x: 35,
					y: -4,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.A2', to: 'led.anode'},
				{from: 'main.GND#1', to: 'led.cathode'},
			],
		},
		'pico': {
			code: code.replace('machine.LED', 'machine.GP20 // Raspberry Pi Pico'),
			parts: {
				led: {
					config: ledConfig,
					x: 16,
					y: -18,
				}
			},
			wires: [
				{from: 'main.GP20', to: 'led.anode'},
				{from: 'main.GND#4', to: 'led.cathode'},
			],
		},
	}});
</script>
