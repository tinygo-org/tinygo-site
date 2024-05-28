---
title: "Reading button input"
description: Learn how to respond to button presses
weight: 1
---

Having some output using a LED is nice, but what's even better is if the electronics respond to you!

In this example, we'll use one of the supported boards that have a button on them. If you don't have one of these boards, you can just read the code here and try the next step in the tour where you can connect an external button (or just use a wire).

A lot of code will be familiar from the blinking LED in the previous step, but there are some important differences:

```go
	button := machine.BUTTONA
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})
```

Instead of configuring a LED, we'll configure the pin that's connected to the button labeled 'A' (usually the leftmost button). The mode is `PinInputPullup` or `PinInputPulldown`, the specific mode depends on the board (we'll come back to this later).

```go
		if button.Get() {
			println("button input is high")
		} else {
			println("button input is low")
		}
```

Here we read out the current button state and print whether it was high (`true`) or low (`false`).

You may find that the button pin is normally high but changes to low when pressed, and other boards are the reverse. This depends on how the button is wired, and will be explained later.

## Serial connection

To see the output of the program, you can use a serial monitor. There is a [tutorial for that]({{< ref "serialmonitor.md" >}}). But in short:

  * If you have TinyGo installed, you can open a serial monitor using `tinygo monitor` command. You may need to specify the board, like `tinygo monitor -target=circuitplay-express`.
  * If you have the Arduino IDE installed, you can use the built-in serial monitor. Make sure to select the right port and to set the baud rate to 115200.
  * On Linux and MacOS you can use a command-line tool like picocom (which you probably need to install manually).
  * If you use Windows, Adafruit has a [list of tools you can use](https://learn.adafruit.com/windows-tools-for-the-electrical-engineer/serial-terminal).

## Button wiring

There are many different types of buttons for loads of different applications. The one we're using here is a plain old pushbutton, or a _mometary_, _normally open_ _SPST_ button:

  * _Momentary_ means that you can press it to change state, but when you release it it will revert to its previous state. This is like a keyboard key, and unlike a light switch (which keeps its state).
  * _Normally open_ means that the connection between the two poles is open (unconnected) when it is not being pressed.
  * _SPST_ is short for _single-pole, single-throw_ which basically means it's a simple switch with two terminals and two states.

For more information about buttons, please [read this SparkFun article](https://learn.sparkfun.com/tutorials/button-and-switch-basics/all).

For the boards in this tour, one terminal of the button is connected to the pin of the microcontroller and another is connected to either VCC (the power source) or GND (ground, or 0V). This depends on the board.

With this, you can understand that as long as you're pressing the button, the input will be what the button is connected to. For example, if it is connected to GND, you will read `false` as long as you're pressing the button.

But what if you're not pressing the button? It's not connected to ground, but also not to VCC, so what is it? The answer is that it is in an intermediary state called _floating_. Sometimes it reads one way, sometimes another, and it's kind of random which it will do. Even moving your hand around the pin might change the state! Electronic designers will try to avoid this state, by gently pulling the voltage either to ground or to VCC, using a resistor. See the next section.

## Pull mode

To avoid the floating state, microcontrollers have a way to gently pull the voltage one way or another using a built-in resistor. This is why we need either `PinInputPullup` or `PinInputPulldown`. This tells the pin to connect a weak resistor, either between the pin and VCC or between the pin and ground. You can think of it as pulling the voltage either way so it avoids wandering off. Once you press the button, the voltage will be pulled very hard in the other direction so will override the weak pull from the pullup or pulldown resistor.

Sometimes, electronic designers will also put these resistors as physical resistors on the board. This is the case for the Gopher Badge for example. In that case, it is not necessary to specify a pull mode: the hardware already does that. But it won't harm either.


<script type="module">
import { setupTour } from '/tour.js';
let codePulldown = `
package main

import (
	"machine"
	"time"
)

func main() {
	button := machine.BUTTONA
	button.Configure(machine.PinConfig{Mode: machine.PinInputPulldown})

	for {
		if button.Get() {
			println("button input is high")
		} else {
			println("button input is low")
		}
		time.Sleep(200 * time.Millisecond)
	}
}`;
let codePullup = codePulldown.replace('PinInputPulldown', 'PinInputPullup');
setupTour({
	boards: {
		'circuitplay-bluefruit': { code: codePulldown },
		'circuitplay-express': { code: codePulldown },
		'gopher-badge': { code: codePullup.replace('BUTTONA', 'BUTTON_A // Gopher Badge uses BUTTON_A') },
		'microbit': { code: codePullup },
	},
});
</script>
