---
title: "Basic command examples"
weight: 1
---

Here will go some basic examples of using TinyGo in the most common scenarios.

The examples can be obtained from the TinyGo source code on Github.

```shell
go get -d -u github.com/tinygo-org/tinygo
cd $GOPATH/src/github.com/tinygo-org/tinygo/src/examples
```

### Building "Hello, World" program for WebAssembly

To build the WASM example, run the following command:

```shell
tinygo build -o wasm.wasm -target=wasm examples/wasm/export
```

See the [WebAssembly page](../../webassembly) for more information on executing the compiled
WebAssembly.

### Building/flashing a "blink" program for micro:bit

To build and then flash a basic blink program for a micro:bit board:

- Plug your micro:bit into your computer's USB port.
- The micro:bit board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Run the following command, substituing the correct name for the board from the previous step:

```shell
tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.hex -target=microbit examples/microbit-blink
```

The top-left pixel of the built-in LED matrix on the micro:bit board should start to blink.

### Building/flashing a "blink" program for Arduino

To build and then flash a basic blink program for an Arduino Uno,

- Plug in the board to your computer's USB port.
- Run the following command, specifiying the serial port if needed. (It defaults to `/dev/ttyACM0` as that is the default port on Linux.)

```shell
tinygo flash -target=arduino -port=[PORT] examples/blinky1
```

The built-in LED on the board should start to blink.

### Building/flashing a "blink" program for Circuit Playground Express

To build and then flash a basic blink program for an Circuit Playground Express,

- Plug your Circuit Playground Express into your computer's USB port.
- Press the "RESET" button on the board two times to get the Circuit Playground Express board ready to receive code. If you just received the board, it will be running MakeCode, so press the "RESET" button just once. Either way, the NeoPixel LEDs should turn all green.
- The Circuit Playground Express board will appear to your computer like a USB drive. Determine the path to the board, for example on Linux it will be something like `/media/[USERNAME]/[NAME OF THE BOARD]`.
- Run the following command, substituing the correct name for the board from the previous step:

```shell
tinygo build -o=/media/[USERNAME]/[NAME OF THE BOARD]/flash.uf2 -target=circuitplay-express examples/blinky1
```

The built-in LED on the board should start to blink.
