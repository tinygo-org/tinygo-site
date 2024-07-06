import { Simulator } from './playground/simulator.js';
import { Editor } from './playground/resources/editor.bundle.min.js';

// Note: to test this locally, run the playground server!
// Run the following in a terminal:
//
//    cd static/playground
//    go run .
//
const PLAYGROUND_API = location.hostname == 'localhost' ? 'http://localhost:8080/api' : 'https://playground-bttoqog3vq-uc.a.run.app/api';

// Initialize the playground embedded in the homepage.
addEventListener('DOMContentLoaded', async () => {
	let playground = document.querySelector('#playground');
	let exampleSelect = playground.querySelector('.example_select');
	let root = playground.querySelector('.simulator');
	let firmwareButton = playground.querySelector('.playground-btn-flash');

	// Load editor.
	let editor = new Editor(playground.querySelector('.playground-editor'));

	// Load simulator.
	let state = structuredClone(examples[exampleSelect.value]);
	editor.setText(state.code);
	let simulator = new Simulator({
		root: root,
		editor: editor,
		firmwareButton: firmwareButton,
		baseURL: new URL('/playground/', document.baseURI),
		apiURL: PLAYGROUND_API,
		features: [],
	});
	await simulator.setState(state, state.target);

	// Respond to changes in the example select box.
	exampleSelect.disabled = false;
	exampleSelect.addEventListener('change', async () => {
		exampleSelect.disabled = true;
		state = structuredClone(examples[exampleSelect.value]);
		editor.setText(state.code);
		await simulator.setState(state, state.target);
		exampleSelect.disabled = false;
	})
});


const exampleHello = `// You can edit this code!
package main

func main() {
	println("hello world!")
}`;

const exampleBlink = `// You can edit this code!
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

const exampleST7789 = `package main

import (
	"machine"
	"image/color"
	"tinygo.org/x/drivers/st7789"
)

func main() {
	// configure the display
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 62_500_000, // 62.5MHz
		Mode:      3,
		SCK:       machine.SPI0_SCK_PIN,
		SDI:       machine.SPI0_SDI_PIN,
		SDO:       machine.SPI0_SDO_PIN,
	})
	display := st7789.New(machine.SPI0,
		machine.TFT_RST,       // reset
		machine.TFT_WRX,       // data/command
		machine.TFT_CS,        // chip select
		machine.TFT_BACKLIGHT) // backlight
	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

	width, height := display.Size()

	white := color.RGBA{255, 255, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	green := color.RGBA{0, 255, 0, 255}
	black := color.RGBA{0, 0, 0, 255}

	display.FillRectangle(0, 0, width/2, height/2, white)
	display.FillRectangle(width/2, 0, width/2, height/2, red)
	display.FillRectangle(0, height/2, width/2, height/2, green)
	display.FillRectangle(width/2, height/2, width/2, height/2, blue)
	display.FillRectangle(width/4, height/4, width/2, height/2, black)
}`;

const exampleLEDs = `package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

func main() {
	// Set up the ring of RGB LEDs.
	pin := machine.WS2812
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws := ws2812.New(pin)
	leds := make([]color.RGBA, 10)

	for cycle := uint8(0); ; cycle++ {
		// Calculate a new color for each LED.
		for i := range leds {
			leds[i] = colorWheel(cycle*2 + uint8(i*8))
		}

		// Write colors to the LED ring.
		ws.WriteColors(leds)

		// Wait a bit each loop (running at around 30fps).
		time.Sleep(time.Second / 30)
	}
}

// Pick a color from the color wheel (red, then, green, then blue, and back to
// red).
func colorWheel(index uint8) color.RGBA {
	index = 255-index
	if index < 85 {
		return color.RGBA{R: 255 - index*3, B: index * 3}
	} else if index < 170 {
		return color.RGBA{G: (index - 85) * 3, B: 255 - (index-85)*3}
	} else {
		return color.RGBA{R: (index - 170) * 3, G: 255 - (index-170)*3}
	}
}`;

var examples = {
	// Simple "hello world" example.
	hello: {
		target: 'console',
		code: exampleHello,
		parts: {
			main: {
				location: 'parts/console.json',
				x: 0,
				y: 0,
			}
		},
		wires: [],
	},

	// Simple "blinky light" example.
	// The Arduino Uno, despite its limitations, is very well known and many
	// people have one. So it seems like a good testcase.
	arduino: {
		target: 'arduino',
		code: exampleBlink,
		parts: {
			main: {
				location: 'parts/arduino.json',
				x: 0,
				y: 0,
			}
		},
		wires: [],
	},

	// Commonly used board, and a nice demo with LEDs.
	circuitplay_express: {
		target: 'circuitplay_express',
		code: exampleLEDs,
		parts: {
			main: {
				location: 'parts/circuitplay-express.json',
				x: 0,
				y: 0,
			},
		},
		wires: [],
	},

	// Our own Gopher Badge, which importantly has a display!
	gopher_badge: {
		target: 'gopher_badge',
		code: exampleST7789,
		parts: {
			main: {
				location: 'parts/gopher-badge.json',
				x: 0,
				y: 0,
			},
		},
		wires: [],
	},
};
