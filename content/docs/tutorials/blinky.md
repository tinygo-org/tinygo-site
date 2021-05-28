---
title: "Blinking LED"
weight: 1
---

After you've [installed TinyGo]({{<ref "install">}}), you can now start using it! We're going to blink an LED. A blinking LED is like the [hello world](https://en.wikipedia.org/wiki/%22Hello,_World!%22_program) of hardware: it is the smallest piece of code that shows the hardware is working.

For this tutorial, previous experience with Go is not required but recommended.

You will need a board with an onboard LED. Most education and development boards will work, with the notable exception of the BBC micro:bit. Boards that are a good start are one of the following:

  - [Adafruit Circuit Playground Express](../../reference/microcontrollers/circuit-playground-express/)
  - [Arduino Nano 33 IoT](../../reference/microcontrollers/arduino-nano33-iot/)

Now that you have a board, let's get started. Start by creating a new directory (for example, named blink) and navigating into it. If you are using Go modules, initialize a new module with the following command:

    go mod init blinky

Next, create a new file named main.go with the following contents:

```go
package main

import (
    "machine"
    "time"
)

func main() {
    led := machine.LED
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    for {
        led.Low()
        time.Sleep(time.Millisecond * 500)

        led.High()
        time.Sleep(time.Millisecond * 500)
    }
}
```

As you can see, it has a main function (the entrypoint of the program) that configures a LED as an output and then in a loop sets the LED to low or high while sleeping for 500 milliseconds in between. We'll explain this more in-depth [later in this tutorial](#background), first let's run this on the microcontroller.

To transfer this program to the board, we'll use the `flash` subcommand. Flashing is the term used for writing a program to a microcontroller. You can do it by `cd`'ing to the directory and running this command:

    tinygo flash -target=circuitplay-express

A few seconds later, the program should be written to the board and an LED should start to blink. If it works, congratulations! You're running your first Go program on a microcontroller.

If you are using a different board, you need to replace the `-target` flag. For example, the [Arduino Uno](../../reference/microcontrollers/arduino-uno/) requires `-target=arduino` and an [Adafruit ItsyBitsy M4](../../reference/microcontrollers/itsybitsy-m4/) requires `-target=itsybitsy-m4`. You can get the full list by going to [Go on microcontrollers](../../reference/microcontrollers/), picking the board on the left and looking under the "Flashing" section.

## Background

So what's really happening here? We'll go through this line by line.

```go
package main
```

Every Go file starts with a `package` line. A Go package is a collection of Go files in a single directory. In this case we only have a single file but most packages will have multiple files.

This package is named main, which is a bit special. It contains the main function, which is the function that will be run when the program starts after initialization.

```go
import (
    "machine"
    "time"
)
```

For this small program we need two packages, which are a kind of libraries:

 1. The machine package, which provides direct access to the hardware in a somewhat portable way. For example, it allows controlling GPIO pins.
 2. The time package, which - you guessed - handles time. You can get the current time with `time.Now()` and sleep for a while with `time.Sleep(duration)`. It also provides some extra utilities such as parsing and formatting time and stuff like time zone handling.

Of these, only the time package is standard Go. The machine package was invented by the TinyGo project.

```go
func main() {
```

This main function is special: like the `main` function in C or `public static void main` in Java, this is the first function that will run. Returning from here will exit the program. It doesn't make sense to "exit" from a microcontroller, so you should avoid returning from it in TinyGo.

```go
led := machine.LED
```

This declares a new variable named `led`. If you are not familiar with Go, the `:=` operator is somewhat like `auto` in C++: it declares a new variable while inferencing the type from the right hand size. In this case, this is the `machine.LED` constant which has type `machine.Pin`.


```go
led.Configure(machine.PinConfig{Mode: machine.PinOutput})
```

This code configures the GPIO pin. Like most peripherals in the machine package, GPIO pins need initialization before they can be used. Maybe the following is clearer if you're not used to this syntax:

```go
config := machine.PinConfig{
    Mode: machine.PinOutput
}
led.Configure(config)
```

The first three lines define a new `machine.PinConfig` object and sets the `Mode` field to `machine.PinOutput`. This means that `config.Mode` will now equal `machine.PinOutput`.


```go
for {
    // ...
}
```

This defines an endless for loop. Go does not have a `while` keyword, instead it has a `for` keyword with zero, one, or three operands. Zero operands means that it is an infinite loop, similar to the Rust `loop { ... }` syntax.

```go
led.Low()
```

This sets the GPIO output to low, meaning ground (0V). Depending on the wiring, it might turn the LED either on or off.

```go
time.Sleep(time.Millisecond * 500)
```

As expected, this sleeps for 500 milliseconds. That doesn't mean nothing is happening: goroutines and interrupts mean that there can be plenty of activity in these 500 milliseconds. But you don't need to worry about that right now.

```go
led.High()
```

This changes the LED state. If it was on before, it will now be off. Otherwise, it will now be on. What is certain is that it will start blinking.

## Conclusion

That's it, your first TinyGo program! I hope this tutorial was helpful. If you have any questions, feel free to join the #tinygo channel on the [Gophers Slack](https://invite.slack.golangbridge.org/).
