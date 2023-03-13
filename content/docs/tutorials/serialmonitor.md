---
title: "Serial Monitor"
weight: 3
description: |
  How to write to and read from the serial port.
---

When you run a Go program on a desktop computer, you can use the `fmt.Print()`,
`fmt.Println()`, and `fmt.Printf()` functions from the Go standard [fmt
package](https://pkg.go.dev/fmt) to print strings and numbers to the terminal
program on the desktop computer. The Go language also supports the low-level
[print() and println()](https://go.dev/ref/spec#Bootstrapping) built-in
functions to print to the terminal.

A TinyGo program running on a microcontroller can use those same functions to
print strings and numbers to its serial monitor port and have them appear on the
terminal program on the host computer. By default, the `fmt` functions and the
`print()/println()` functions are configured to send to the `machine.Serial`
object of the microcontroller.

On some microcontrollers, the `machine.Serial` object is configured to send to
the
[UART](https://en.wikipedia.org/wiki/Universal_asynchronous_receiver-transmitter)
chip, which is often wired to a [USB-to-serial
adapter](https://en.wikipedia.org/wiki/USB-to-serial_adapter) chip on the dev
board. The adapter chip converts the serial bits into USB packets to the host
computer. On other microcontrollers, the USB controller is built directly into
the [microcontroller SoC](https://en.wikipedia.org/wiki/System_on_a_chip). The
`machine.Serial` object on these microcontrollers is configured to send to the
USB bus directly, instead of going through a USB-to-serial adapter.

In the context of this tutorial, it does not matter whether the microcontroller
uses a UART controller or a USB controller. In both cases, the microcontroller
will appear as a serial device on the host computer which can communicate via
applications that read from and write to the serial port on the host computer.

## Serial Output

### Using `fmt.Print()` and `fmt.Println()`

Here is a sample program that writes a line every second to the `machine.Serial`
port:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    count := 0
    for {
        fmt.Println(count, ": Hello, World")
        time.Sleep(time.Millisecond * 1000)
        count++
    }
}
```

This can be flashed to the microcontroller using the `tinygo flash` command
described in the [Blinky tutorial]({{<ref "blinky">}}).

### Using `print()` and `println()`

One problem with the above program is that `fmt` is a large package that
consumes substantial amount of flash memory on a microcontroller. The built-in
functions `print()` and `println()` consume far less resources. The above
program can be written like this:

```go
package main

import (
    "time"
)

func main() {
    count := 0
    for {
        println(count, ": Hello, World")
        time.Sleep(time.Millisecond * 1000)
        count++
    }
}
```

An estimate of the flash memory consumption can be printed by the TinyGo
compiler using the [-size]({{<ref "../reference/usage/misc-options">}}) flag.
Here is a table that shows the 2 versions of the program above for some
microcontrollers that I have readily available:

```
+-----------------+---------------+-----------+
| Board Type      | fmt.Println() | println() |
+-----------------+---------------+-----------+
| Arduino Zero    | 43532         | 7328      |
| Seeeduino Xiao  | 43532         | 7388      |
| STM32 BluePill  | 41756         | 6296      |
| ESP8266 D1 Mini | 44961         | 3588      |
| ESP32           | 42410         | 3335      |
+-----------------+---------------+-----------+
```

The `fmt` package increases flash memory consumption by 35 kB to 40 kB. On some
microcontrollers with limited amount of flash memory, (e.g. the STM32 Blue Pill
with 64 kB of flash, the Arduino Zero or Seeeduino Xiao both with 256 kB of
flash), it may be worth avoiding the overhead of the `fmt` package by using the
built-in `print()` and `println()` instead.

## Serial Monitor on Host Computer

To see the output of the serial port from the microcontroller, we need to run a
serial monitor application on the host computer. There are many ways to do this,
but the easiest is probably the `tinygo monitor` subcommand which is built
directly into the `tinygo` program itself.

### `monitor` subcommand

After flashing the program above, run the `tinygo monitor` program to see the
output every second from the microcontroller:

```
$ tinygo monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
4 : Hello, World
5 : Hello, World
[...]
```

In this example, the serial monitor missed the first 4 lines of "Hello, World"
(0 to 3) because the program started to print those lines immediately after
flashing, but before the serial monitor was connected.

### `-monitor` flag

It is often useful to automatically start the monitor immediately after flashing
your program to the microcontroller. The `tinygo flash` command takes an
optional `-monitor` flag to accomplish this:

```
$ tinygo flash -target=xiao -monitor
```

On some microcontrollers, the `-monitor` flag fails with the following error
message because the monitor starts too quickly:

```
$ tinygo flash -target=arduino-zero -monitor
[...]
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
error: read error: Port has been closed
```

If this happens, you can chain the `flash` and `monitor` subcommands manually,
with a 1 or 2-second delay between the two commands. On Linux or MacOS, the
command invocation looks like this:

```
$ tinygo flash -target=arduino-zero && sleep 1 && tinygo monitor
```

(The `&&` separator runs the next command only if the previous command completed
without errors. This is safer than using the semicolon `;` separator because the
semicolon continues to execute commands even if the previous command returned an
error code.)

### Baud Rate

The default [baud rate](https://en.wikipedia.org/wiki/Serial_port#Speed) of the
serial port for almost all microcontrollers supported by TinyGo is 115200. The
exceptions are boards using the AVR processors ([Arduino Nano]({{<ref
"../reference/microcontrollers/arduino-nano">}}), [Arduino Mega 1280]({{<ref
"../reference/microcontrollers/arduino-mega1280">}}), [Arduino Mega 2560]({{<ref
"../reference/microcontrollers/arduino-mega2560">}})). On these, the serial port
is set to 9600, so you need to override the baud rate of `tinygo monitor` like
this:

```
$ tinygo monitor -baudrate=9600
```

You can combine the `flash` subcommand, the `-monitor` flag, and the `-baudrate`
flag into a single invocation like this:

```
$ tinygo flash -target arduino-nano -monitor -baudrate 9600
```

(Notice that the `=` after each flag has been replaced with a space. It's an
alternative syntax that some people prefer because a space is easier to type
than an equal sign `=`.)

### Serial Port on Host

The microcontroller will be assigned a serial port on the host computer. If you
have only a single microcontroller attached, you will normally not need to worry
about what these serial ports are called. The `tinygo monitor` will
automatically figure out which serial port to use.

On Linux machines, the serial port will have a `USB` prefix or an `ACM` prefix
like this:

* `/dev/ttyUSB0`
* `/dev/ttyACM0`

On MacOS machines, the serial port will look like this:

* `/dev/cu.usbserial-1420`
* `/dev/cu.usbmodem6D8733AC53571`

On Windows machines, the serial port looks something like:

* `COM1`
* `COM31`

### Multiple Microcontrollers

If you have more than one microcontroller attached to the host computer, the
`tinygo flash` and `tinygo monitor` subcommands can *sometimes* figure out which
port it is using, but they will sometimes print out an error message, like this:

```
$ tinygo flash -target arduino-nano
error: multiple serial ports available - use -port flag,
available ports are /dev/ttyACM0, /dev/ttyUSB0
```

You then need to supply the `-port` flag to identify the microcontroller that
you want to flash and monitor:

```
$ tinygo flash -target=arduino-nano -port=/dev/ttyUSB0

$ tinygo monitor -port=/dev/ttyUSB0 -baudrate=9600
```

Sometimes it is possible to combine the two commands into a single command even
in the presence of multiple microcontrollers:

```
$ tinygo flash -target xiao -monitor
```

But sometimes, combining `flash` and `monitor` into a single command does not
work. In that case, you can issue the `flash` and `monitor` commands separately.
But it is often easier to just pull out the extra microcontroller(s) so that
only a single board is connected to the host computer.

```
$ tinygo flash -target=arduino-nano -monitor
error: multiple serial ports available - use -port flag,
available ports are /dev/ttyACM0, /dev/ttyUSB0

$ tinygo flash -target=arduino-nano -monitor -port=/dev/ttyUSB0 -baudrate=9600
[...]
avrdude: 4238 bytes of flash verified
avrdude done.  Thank you.
[...]
error: multiple serial ports available - use -port flag,
available ports are /dev/ttyACM0, /dev/ttyUSB0
```

## Serial Input

Occasionally it is useful to send characters from the host computer to the
microcontroller. The following program reads a single byte from the
`machine.Serial` object and prints the character back to the host computer.

The caveat is that the `Serial.ReadByte()` feature is *not* currently
implemented on every microcontroller supported by TinyGo. For example, the
following program does not work on the ESP32 or the ESP8266.

```go
package main

import (
    "machine"
    "time"
)

func main() {
    time.Sleep(time.Millisecond * 2000)
    println("Reading from the serial port...")

    for {
        c, err := machine.Serial.ReadByte()
        if err == nil {
            if c < 32 {
                // Convert nonprintable control characters to
                // ^A, ^B, etc.
                machine.Serial.WriteByte('^')
                machine.Serial.WriteByte(c + '@')
            } else if c >= 127 {
                // Anything equal or above ASCII 127, print ^?.
                machine.Serial.WriteByte('^')
                machine.Serial.WriteByte('?')
            } else {
                // Echo the printable character back to the
                // host computer.
                machine.Serial.WriteByte(c)
            }
        }

        // This assumes that the input is coming from a keyboard
        // so checking 120 times per second is sufficient. But if
        // the data comes from another processor, the port can
        // theoretically receive as much as 11000 bytes/second
        // (115200 baud). This delay can be removed and the
        // Serial.Read() method can be used to retrieve
        // multiple bytes from the receive buffer for each
        // iteration.
        time.Sleep(time.Millisecond * 8)
    }
}
```

You can flash this program to the microcontroller (in this example, a SAMD21 M0+
clone that emulates an Arduino Zero), and fire up the monitor like this:

```
$ tinygo flash -target=arduino-zero
$ tinygo monitor
Connected to /dev/ttyACM0. Press Ctrl-C to exit.
Reading from the serial port...
abcdef^A^B^D^E^F^G^H^I^J^K^L^M^N^O^P^R^T^U^V^W^X^Y^^^[^]^_^?
```

Type a few characters in the `tinygo monitor`, for example "abcdef". You should
see the characters echoed back by the microcontroller, as shown above. If you
type a [nonprintable control
characters](https://en.wikipedia.org/wiki/C0_and_C1_control_codes), these are
echoed back as 2 characters: the caret character `^` and a letter representing
the control character. For example, typing Control-P prints `^P`.

Of the 32 possible control characters, some of them are intercepted by the
`tinygo monitor` itself instead of being sent to the microcontroller:

* Control-C: terminates the `tinygo monitor`
* Control-Z: suspends the `tinygo monitor` and drops back into shell
* Control-\\: terminates the `tinygo monitor` with a stack trace
* Control-S: flow control, suspends output to the console
* Control-Q: flow control, resumes output to the console
* Control-@: thrown away by `tinygo monitor`

## Alternative Serial Monitors

There are many alternative serial monitor programs that can be used instead of
`tinygo monitor`. The setup is slightly more complicated because you will need
to supply the serial port and baud rate of the microcontroller as described in
the [Serial Port on Host](#serial-port-on-host) and [Baud Rate](#baud-rate)
subsections above.

### Arduino IDE

The [Arduino IDE](https://www.arduino.cc/en/software) contains its own serial
monitor. You may choose to use that instead. You need to set the serial port
(something like `/dev/ttyUSB0` on Linux, or `/dev/cu.usbserial-1420` on MacOS),
and set the baud rate to 115200 (or 9600 on AVR processors).

### pyserial

The [pyserial](https://pyserial.readthedocs.io/en/latest/pyserial.html) is a
Python library that comes with its own serial monitor. Setting up a python3
environment is a complex topic that is beyond the scope of this document. But if
you are able to install `python3` and `pip3`, you can install `pyserial` and use
its built-in `miniterm` tool roughly like this:

```
$ python3 -m pip install --user pyserial
$ python3 -m serial.tools.miniterm /dev/ttyUSB0 115200
```

Another useful feature of `pyserial` is the `list_ports` command:

```
$ python3 -m serial.tools.list_ports
/dev/ttyACM0
/dev/ttyS0
/dev/ttyUSB0
3 ports found
```

This is useful when you plug in a random microcontroller to the USB port, and
you cannot remember which serial port it is mapped to.

### picocom

The [picocom](https://github.com/npat-efault/picocom) terminal emulator can be
installed on both Linux and MacOS. If you are using an Ubuntu flavored Linux,
the installation is something like:

```
$ sudo apt install picocom
```

On MacOS, most people use [Homebrew](https://brew.sh/), and it can be installed
like this:

```
$ brew install picocom
```

It can be invoked like this:

```
$ picocom -b 115200 /dev/ttyACM0
port is        : /dev/ttyACM0
picocom v3.1
[...]
Type [C-a] [C-h] to see available commands
Terminal ready
```
