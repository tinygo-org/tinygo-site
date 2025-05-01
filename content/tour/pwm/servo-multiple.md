---
title: "Servo: multiple servos"
description: "Control multiple servos at once"
weight: 12
---

Since many PWM peripherals have multiple channels (usually 2 or 4) we can in fact control multiple servos using a single PWM peripheral (freeing up other PWM peripherals for other tasks). This is why it's called `servo.NewArray`!

The code is very similar to using a single servo, but this time we control more than one from different goroutines. If you need more servos, you can of course use more than one servo array with different PWM peripherals.


<script type="module">
import { setupTour } from '/tour.js';
let code = `
package main

import "machine"
import "time"
import "tinygo.org/x/drivers/servo"

func main() {
	// Board specific configuration.
	servoPin1 := machine.GP16
	servoPin2 := machine.GP17
	pwm := machine.PWM0

	// Configure the servo array.
	array, err := servo.NewArray(pwm)
	if err != nil {
		println("could not configure:", err.Error())
		return
	}

	// Add both servos to the servo array.
	servo1, err := array.Add(servoPin1)
	if err != nil {
		println("could not add servo 1:", err.Error())
		return
	}
	servo2, err := array.Add(servoPin2)
	if err != nil {
		println("could not add servo 2:", err.Error())
		return
	}

	go moveServo(servo2, 1300*time.Millisecond)
	moveServo(servo1, 1000*time.Millisecond)
}

func moveServo(s servo.Servo, delay time.Duration) {
	for {
		// Move down.
		s.SetMicroseconds(1000)
		time.Sleep(delay)

		// Move up.
		s.SetMicroseconds(1500)
		time.Sleep(delay)
	}
}
`;
let servoConfig1 = {
	type: 'servo',
	humanName: 'Servo 1',
	svg: 'servo-sg90.svg',
};
let servoConfig2 = {
	type: 'servo',
	humanName: 'Servo 2',
	svg: 'servo-sg90.svg',
};
setupTour({
	boards: {
		'arduino-nano33': {
			code: code.replace('machine.PWM0', 'machine.TCC0').replace('machine.GP16', 'machine.D2').replace('machine.GP17', 'machine.D3'),
			parts: {
				main: {
					x: 30,
				},
				servo1: {
					config: servoConfig1,
					x: -30,
					y: -12,
				},
				servo2: {
					config: servoConfig2,
					x: -30,
					y: 12,
				},
			},
			wires: [
				{from: 'main.D2', to: 'servo1.control'},
				{from: 'main.D3', to: 'servo2.control'},
			],
		},
		'circuitplay-bluefruit': {
			code: code.replace('machine.GP16', 'machine.A2').replace('machine.GP17', 'machine.A1'),
			parts: {
				main: {
					x: 30,
				},
				servo1: {
					config: servoConfig1,
					x: -25,
					y: -12,
				},
				servo2: {
					config: servoConfig2,
					x: -25,
					y: 12,
				},
			},
			wires: [
				{from: 'main.A2', to: 'servo1.control'},
				{from: 'main.A1', to: 'servo2.control'},
			],
		},
		'circuitplay-express': {
			code: code.replace('machine.PWM0', 'machine.TCC1').replace('machine.GP16', 'machine.A3').replace('machine.GP17', 'machine.A2'),
			parts: {
				main: {
					x: 25,
				},
				servo1: {
					config: servoConfig1,
					x: -25,
					y: -12,
				},
				servo2: {
					config: servoConfig2,
					x: -25,
					y: 12,
				},
			},
			wires: [
				{from: 'main.A3', to: 'servo1.control'},
				{from: 'main.A2', to: 'servo2.control'},
			],
		},
		'pico': {
			code: code,
			parts: {
				main: {
					x: 30,
				},
				servo1: {
					config: servoConfig1,
					x: -30,
					y: -12,
				},
				servo2: {
					config: servoConfig2,
					x: -30,
					y: 12,
				},
			},
			wires: [
				{from: 'main.GP16', to: 'servo1.control'},
				{from: 'main.GP17', to: 'servo2.control'},
			],
		},
	}});
</script>
