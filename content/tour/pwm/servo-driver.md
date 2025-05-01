---
title: "Servo: using a driver"
description: "Control a servo motor using the servo driver to simplify things."
weight: 11
---

In many cases, it would be nice to avoid having to care about all these specific calculations. Luckily we have the [servo driver](https://pkg.go.dev/tinygo.org/x/drivers/servo) as part of the TinyGo drivers module. It still requires looking up the right PWM and pin combinations, but it does simplify controlling the servos a bit.

To use the driver, we need to configure it first:

```go
	// Configure the servo array.
	array, err := servo.NewArray(pwm)
	if err != nil {
		println("could not configure:", err.Error())
		return
	}
```

We just pass in the PWM peripheral, and it will configure the PWM to the right period size. Note it is called an "array" since we can actually control multiple servos! More on that later.

Next, we can configure the servo array to use this servo:

```go
	// Add a servo to the servo array.
	servo, err := array.Add(servoPin)
	if err != nil {
		println("could not add servo pin:", err.Error())
		return
	}
```

This will configure the pin to be used as a servo pin, and returns a `servo.Servo` instance. This instance can then be used to set the servo angle:

```go
	// Left position at 2000µs (2ms) "on" time.
	println("left:  -90°")
	servo.SetMicroseconds(2000)
	time.Sleep(time.Second * 3)
```


<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import "machine"
import "time"
import "tinygo.org/x/drivers/servo"

func main() {
	// Board specific configuration.
	servoPin := machine.GP16
	pwm := machine.PWM0

	// Configure the servo array.
	array, err := servo.NewArray(pwm)
	if err != nil {
		println("could not configure:", err.Error())
		return
	}

	// Add a servo to the servo array.
	servo, err := array.Add(servoPin)
	if err != nil {
		println("could not add servo:", err.Error())
		return
	}

	// Left position at 2000µs (2ms) "on" time.
	println("left:  -90°")
	servo.SetMicroseconds(2000)
	time.Sleep(time.Second * 3)

	// Right position, at 1000µs (1ms) "on" time.
	println("right: +90°")
	servo.SetMicroseconds(1000)
	time.Sleep(time.Second * 3)

	// Center position, at 1500µs (1.5ms) "on" time.
	println("center: 0°")
	servo.SetMicroseconds(1500)
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
