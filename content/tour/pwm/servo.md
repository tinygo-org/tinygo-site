---
title: "Servo: intro"
description: "Control a servo motor directly using PWM"
weight: 10
---

Servos are rotary motors that allow for precise control of their angular position. The position is set using a form of pulse-width modification (PWM). Every 20ms or so (the precise interval doesn't matter too much) a pulse of 1ms to 2ms is sent that indicates the target position for the motor. The pulse width is what matters here: 1.5ms indicates the center position, 1ms all the way to the right, and 2ms all the way to the left.

There is a decent amount of variation among servo motors:

  * Servos generally accept a wider range of pulse intervals, 20ms is standard but it doesn't matter too much.
  * The usual range is 1ms to 2ms, but many servos accept a somewhat larger range from about 0.8ms to 2.2ms. Be careful with this: putting a servo out of range will let it keep trying to reach that position (which it can't) which will quickly wear out the motor!
  * Different servos have different ranges. Typical ranges are 90° and 180°, the simulator has a 180° servo.

So we need to send very precisely timed pulses every 20ms. That's exactly the sort of thing the built-in PWM is made for! Here is how we can do that:

  * Configure the PWM peripheral to use a period of 20ms.
  * Configure a channel/pin to be used for the servo.
  * Set the channel to a duty cycle that matches the 1-2ms range.

This is exactly what the code on the right is doing. Let's look at setting the servo in a specific position:

```go
	// Left position at 2000µs (2ms) "on" time.
	println("left:  -90°")
	pwm.Set(ch, pwm.Top()*20/200)
	time.Sleep(time.Second * 3)
```

We see it uses the following formula for calculating the duty cycle:

```math
\frac{T*20}{200}
```

Simplifying, we get:

```math
\frac{T}{10}
```

In other words, the 20ms period is divided by 10 to get an "on" time of 2ms: exactly what we wanted! The fractions for 1ms (all the way to the right) and 1.5ms (center position) are the same.

Because servos aren't instantaneous, we do need to wait a bit for the change to take effect:

```go
	time.Sleep(time.Second * 3)
```

3 seconds is much longer than it needs (it will arrive much faster in the target position) but it lets us see what is happening. If we'd set the PWM channel to 2ms (to the left), and then immediately to 1ms (to the right), the servo wouldn't have time to go to the left and immediately go to the right!

## Connecting a servo

Servos only need three wires to work:

  * Black or brown goes to ground.
  * Red goes to the power supply (usually 5V but check your servo datasheet).
  * Yellow/white/orange is the control wire, and is the one you will connect to the microcontroller.

Be careful about the power supply! Servos need a lot of current when they are rotating, and the 5V output of a board may not be enough. So an external power supply might be a better idea. If you use a separate power supply, do make sure to connect the ground wires together otherwise the signal won't be received properly.

## More information

SparkFun has made an [excellent guide](https://www.sparkfun.com/servos) on servos which is well worth a read!

<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import "machine"
import "time"

func main() {
	// Board specific configuration.
	servoPin := machine.GP16
	pwm := machine.PWM0

	// Configure the timer/PWM.
	err := pwm.Configure(machine.PWMConfig{
		Period: 20e6, // 20ms (for standard servos)
	})
	if err != nil {
		println("could not configure:", err.Error())
		return
	}

	// Get the channel for this PWM peripheral.
	ch, err := pwm.Channel(servoPin)
	if err != nil {
		println("could not obtain channel:", err.Error())
		return
	}

	// Left position at 2000µs (2ms) "on" time.
	println("left:  -90°")
	pwm.Set(ch, pwm.Top()*20/200)
	time.Sleep(time.Second * 3)

	// Right position, at 1000µs (1ms) "on" time.
	println("right: +90°")
	pwm.Set(ch, pwm.Top()*10/200)
	time.Sleep(time.Second * 3)

	// Center position, at 1500µs (1.5ms) "on" time.
	println("center: 0°")
	pwm.Set(ch, pwm.Top()*15/200)
	time.Sleep(time.Second * 3)
}
`;
let servoConfig = {
	type: 'servo',
	humanName: 'Servo',
	svg: 'servo-sg90.svg',
};
setupTour({
	boards: {
		'arduino-nano33': {
			code: code.replace('machine.PWM0', 'machine.TCC0').replace('machine.GP16', 'machine.D2'),
			parts: {
				main: {
					y: 15,
				},
				servo: {
					config: servoConfig,
					x: 0,
					y: -12,
				},
			},
			wires: [
				{from: 'main.D2', to: 'servo.control'},
			],
		},
		'circuitplay-bluefruit': {
			code: code.replace('machine.GP16', 'machine.A1'),
			parts: {
				main: {
					x: -25,
				},
				servo: {
					config: servoConfig,
					x: 25,
					y: -10,
				},
			},
			wires: [
				{from: 'main.A1', to: 'servo.control'},
			],
		},
		'circuitplay-express': {
			code: code.replace('machine.PWM0', 'machine.TCC0').replace('machine.GP16', 'machine.A1'),
			parts: {
				main: {
					x: -25,
				},
				servo: {
					config: servoConfig,
					x: 25,
					y: -10,
				},
			},
			wires: [
				{from: 'main.A1', to: 'servo.control'},
			],
		},
		'pico': {
			code: code,
			parts: {
				main: {
					y: 15,
				},
				servo: {
					config: servoConfig,
					x: 0,
					y: -12,
				},
			},
			wires: [
				{from: 'main.GP16', to: 'servo.control'},
			],
		},
	}});
</script>
