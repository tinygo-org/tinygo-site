---
title: "Using PWM"
weight: 2
---

If you've gone through the [blinky](../blinky) example and want to take it to the next level with [PWM](https://en.wikipedia.org/wiki/Pulse-width_modulation), or **Pulse-Width Modulation**, this is the tutorial for you. PWM allows for control of most [servos](https://en.wikipedia.org/wiki/Servomotor) and LED intensity. We'll be trying out the latter here.

**Requirements**
* Go tool. Get it at [golang.org](https://golang.org/)
* TinyGo tool. Here's the [quick install guide](/getting-started/install/)
* A microcontroller with an on-board LED and PWM peripheral such as the [Raspberry Pi Pico](https://www.raspberrypi.org/products/raspberry-pi-pico/) (MSRP $4 USD)

Let's create a directory anywhere on your computer and navigate to it with the terminal. Run the following command in the directory:

```shell
go mod init pwm-blinky
```

Next, create a new file named main.go with the following contents:

```go
package main

import (
	"machine"
	"time"
)

var period uint64 = 1e9 / 500

func main() {
    // This program is specific to the Raspberry Pi Pico.
    pin := machine.LED
	pwm := machine.PWM4 // Pin 25 (LED on pico) corresponds to PWM4.

	// Configure the PWM with the given period.
	pwm.Configure(machine.PWMConfig{
		Period: period,
	})

	ch, err := pwm.Channel(pin)
	if err != nil {
		println(err.Error())
		return
	}

	for { 
		for i := 1; i < 255; i++ {
            // This performs a stylish fade-out blink
			pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Millisecond * 5)
		}
	}
}
```

There's quite a bit to unpack in the program above and there are some notable differences between the blinky example. The main take-away here is the two parts to the program: first we setup the PWM peripheral which corresponds to the LED pin, lastly we enter an infinite for loop which blinks the LED with variable intensity given by `pwm.Set` method.

To transfer this program to the board, we'll use the `flash` subcommand. Flashing is the term used for writing a program to a microcontroller. You can do it by `cd`'ing to the directory and running this command:

    tinygo flash -target=pico

Wait a few seconds and the LED should start to blink with a fade-out effect. This effect can be changed by changing the second argument to `pwm.Set`.

## Explanation

So let's take a deeper look, how does PWM work behind the curtains? We'll go through the program line by line.

The first ~10 or so lines are the package declaration and imports, those are explained in finer detail in the [blinky](../blinky) example.

Our logic for the program lies entirely within the `main` function. This is where the magic happens.

```go
led := machine.LED
```
This declares a new variable named `led`. If you are not familiar with Go, the `:=` operator is somewhat like `auto` in C++: it declares a new variable while inferencing the type from the right hand size. In this case, this is the `machine.LED` constant which has type `machine.Pin`. Boards which have an on-board LED will have this constant available for use.

```go
pwm := machine.PWM4
```
Now things start to look quite different to the blinky example. This code creates a variable which references a PWM peripheral. 

A _peripheral_ refers to any on-board integrated circuit which interfaces directly with the microprocessor and has I/O capabilities. Some other tasks which different peripherals may manage are communications ([I2C](https://en.wikipedia.org/wiki/I%C2%B2C), [SPI](https://en.wikipedia.org/wiki/Serial_Peripheral_Interface), UART, etc.) and [ADC](https://en.wikipedia.org/wiki/Analog-to-digital_converter)/[DAC](https://en.wikipedia.org/wiki/Digital-to-analog_converter).

In the case of a PWM peripheral, there is the additional idea of a channel which can be used to control more than one pin at the same time with a single PWM peripheral.

```go
var period uint64 = 1e9 / 500
```
We previously declared a new variable to calculate and store the PWM period of 2 milliseconds. By convention periods are unsigned integers which represent a duration in _nanoseconds_, also written as _ns_. The formula used above is `period = 1 / frequency`, where `500` is the desired frequency in hertz (periods per second) and `1e9` is a conversion constant to convert a period in seconds to nanoseconds. 

The period of a PWM signal is the time between rising edges.

```go
pwm.Configure(machine.PWMConfig{Period: period})
```
This code configures the PWM peripheral with the desired period. Keep in mind the output period of the PWM is usually not exact and may differ between microcontrollers.

```go
ch, err := pwm.Channel(pin)
```
This code obtains the channel of the peripheral corresponding to the `pin` we want output on.

```go
for { 
    // ... 
}
```
This code defines an endless `for` loop. All of it's contents are run on repeat forever. This can usually be found as "`while true`" in other languages.

```go
for i := 1; i < 255; i++ {
    // ...
}
```
This defines a loop which:
* declares and initializes the variable `i` to `1` (`i := 1`)
* Loops until `i` is equal to 255 or greater (`i < 255`)
* Adds `1` to `i` after each loop (`i++`)


```go
pwm.Set(ch, pwm.Top()/uint32(i))
```
The `Set` method sets the [duty cycle](https://en.wikipedia.org/wiki/Duty_cycle) of the PWM. We must pass in the channel `ch` we obtained earlier which sets the PWM on the desired pin.

The second argument to `Set` is the **threshold**, which is expressed as a part of the PWM's `top` value. There are various details to how PWM works at the lowest detail, to put it simply:

_PWM works by having a counter (like the variable `i` mentioned above) which decides if the pin is "off" or "on". If this counter is above the **threshold** then the pin is "off", if it is below the threshold it is "on". The counter is incremented automatically continuously until it reaches the **top** value, when the counter reaches top, it is reset to 0. The fraction of time the pin is "on" vs. "off" is called the **duty cycle**._ 


```go
pwm.Top()/uint32(i)
```
By using the above expression we are setting the threshold at a value between `top` and `top/255`, where `threshold=top` is a 100% dutycycle (always on) and `threshold=top/255` is 1/255% dutycycle, or 0.4% (practically off).


```go
time.Sleep(time.Millisecond * 5)
```
This code stops the program for 5 milliseconds and does nothing for that duration. It is used to slow down the rate at which the LED blinks. If it were not present then we would not notice the fade out effect because it would happen too fast to be perceived!

## Final thoughts

Although PWM may have its ups and downs which make it a complex topic to explain thoroughly, the program shown has everything one needs to begin harnessing the power of PWM control.

I hope this tutorial was helpful. If you have any questions, feel free to join the #tinygo channel on the [Gophers Slack](https://invite.slack.golangbridge.org/).
