---
title: "LED: fade multiple"
description: "Control multiple channels from a single PWM peripheral by making use of channels."
weight: 3
---

One PWM peripheral can be used to control multiple LEDs as well, by using different channels. In that case, the frequency (or period) will be the same for all LEDs, but the duty cycle (percent on) will vary per channel and therefore per LED.

On the right, we have a single PWM instance but we use two goroutines to fade both LEDs in and out. Both LEDs are using a different PWM channel.

## Finding the correct PWM peripheral and pins

You may be wondering how we picked the right PWM peripheral and pins? It's actually all in the documentation if you know where to look:

  * [Arduino Nano 33 IoT](/docs/reference/microcontrollers/arduino-nano33/#pins)
  * [Adafruit Circuit Playground Bluefruit](/docs/reference/microcontrollers/circuitplay-bluefruit/#pins) (special case, see below)
  * [Adafruit Circuit Playground Express](/docs/reference/microcontrollers/circuitplay-express/#pins)
  * [Raspberry Pi Pico](/docs/reference/microcontrollers/pico/#pins)

Let's look at the first, the Arduino Nano 33 IoT. In the Pins table, there are these two rows:

> | Pin               | Hardware pin | Alternative names | PWM                  |
> | ----------------- | ------------ | ----------------- | -------------------- |
> | `A2`              | `PA11`       |                   | `TCC1` (channel 1), `TCC0` (channel 3) |
> | `A3`              | `PA10`       | `I2S_SCK_PIN`     | `TCC1` (channel 0), `TCC0` (channel 2) |

In this table you can see that we actually could have used two different PWM peripherals! Either TCC1 or TCC0. Both are supported for both pins, and importantly both use different channels for each pin. For TCC1 (as used on the right) channel 1 and 0 are used, while for TCC0 channel 2 and 3 would have been used.

Question: can you verify that the peripheral used for the other 3 boards corresponds to the PWM peripheral and channel used in the simulator on the right?

You might have noticed that three out of the four boards list PWMs, but one does not: the Circuit Playground Bluefruit. This is because the PWM peripherals on this chip can be connected to any pin. So what happens instead is that the `machine.PWM` type will not look up the pin in a table or something, but will instead use the first available unused channel. This brings a lot of flexibility, especially when you are working with special device types like servos (which will be addressed later in this tour).


<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import "machine"
import "time"

// Board specific configuration.
var pwm = machine.TCC1

func main() {
	// Configure the timer/PWM.
	// Use the default period, which works well enough for LEDs.
	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		println("could not configure:", err.Error())
		return
	}

	// Now fade both LEDs in and out.
	go fadeLED(machine.A2, time.Millisecond*1000)
	fadeLED(machine.A3, time.Millisecond*1300)
}

func fadeLED(led machine.Pin, duration time.Duration) {
	// Get the channel for this PWM peripheral.
	ch, err := pwm.Channel(led)
	if err != nil {
		println("could not obtain channel:", err.Error())
		return
	}

	println("fading a LED on channel", ch, "every", duration.String())

	for {
		// Fade the LED in.
		for percentOn := 0; percentOn <= 100; percentOn++ {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(duration / 202)
		}

		// Fade the LED out.
		for percentOn := 100; percentOn >= 0; percentOn-- {
			pwm.Set(ch, pwm.Top()*uint32(percentOn)/100)
			time.Sleep(duration / 202)
		}
	}
}
`;
let ledConfig1 = {
	type: 'led',
	humanName: 'External LED 1',
	color: [255, 0, 0],
	svg: 'led-tht-5mm.svg',
};
let ledConfig2 = {
	type: 'led',
	humanName: 'External LED 2',
	color: [255, 0, 0],
	svg: 'led-tht-5mm.svg',
};
setupTour({
	code: code,
	boards: {
		'arduino-nano33': {
			parts: {
				main: {
					y: -8,
				},
				led1: {
					config: ledConfig1,
					x: 0,
					y: 12,
				},
				led2: {
					config: ledConfig2,
					x: 14,
					y: 12,
				}
			},
			wires: [
				{from: 'main.A2', to: 'led1.anode'},
				{from: 'main.A3', to: 'led2.anode'},
				{from: 'main.GND#2', to: 'led1.cathode'},
				{from: 'main.GND#2', to: 'led2.cathode'},
			],
		},
		'circuitplay-bluefruit': {
			code: code.replace('machine.TCC1', 'machine.PWM0'),
			parts: {
				led1: {
					config: ledConfig1,
					x: 35,
					y: -8,
					rotation: 90,
				},
				led2: {
					config: ledConfig2,
					x: 45,
					y: -8,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.A2', to: 'led1.anode'},
				{from: 'main.A3', to: 'led2.anode'},
				{from: 'main.GND#1', to: 'led1.cathode'},
				{from: 'main.GND#1', to: 'led2.cathode'},
			],
		},
		'circuitplay-express': {
			parts: {
				led1: {
					config: ledConfig1,
					x: 35,
					y: -8,
					rotation: 90,
				},
				led2: {
					config: ledConfig2,
					x: 45,
					y: -8,
					rotation: 90,
				}
			},
			wires: [
				{from: 'main.A2', to: 'led1.anode'},
				{from: 'main.A3', to: 'led2.anode'},
				{from: 'main.GND#1', to: 'led1.cathode'},
				{from: 'main.GND#1', to: 'led2.cathode'},
			],
		},
		'pico': {
			code: code.replace('machine.TCC1', 'machine.PWM1').replace('machine.A2', 'machine.GP2').replace('machine.A3', 'machine.GP3'),
			parts: {
				main: {
					y: -8,
				},
				led1: {
					config: ledConfig1,
					x: -16,
					y: 12,
				},
				led2: {
					config: ledConfig2,
					x: -3,
					y: 12,
				}
			},
			wires: [
				{from: 'main.GP2', to: 'led1.anode'},
				{from: 'main.GP3', to: 'led2.anode'},
				{from: 'main.GND#6', to: 'led1.cathode'},
				{from: 'main.GND#6', to: 'led2.cathode'},
			],
		},
	}});
</script>
