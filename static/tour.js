import { Simulator } from './playground/simulator.js';
import { Editor } from '/playground/resources/editor.bundle.min.js';

export { setupTour };

// Note: to test this locally, run the playground server!
// Run the following in a terminal:
//
//    cd static/playground
//    go run .
//
const PLAYGROUND_API = location.hostname == 'localhost' ? 'http://localhost:8080/api' : 'https://playground-bttoqog3vq-uc.a.run.app/api';

const boardNames = {
	'arduino': 'Arduino Uno',
	'arduino-nano33': 'Arduino Nano 33 IoT',
	'circuitplay-bluefruit': 'Adafruit Circuit Playground Bluefruit',
	'circuitplay-express': 'Adafruit Circuit Playground Express',
	'gopher-badge': 'Gopher Badge',
	'microbit': 'BBC micro:bit v1',
	'pico': 'Raspberry Pi Pico',
}

// Configure the tour. This works on pages that use the "tour" layout.
// The config parameter has the following layout:
//
//   {
//     code: '...',                // default code
//     boards: {
//       'arduino': {              // board name
//          code: '...',           // code for this board (optional)
//          parts: {               // extra parts besides the main part (optional)
//            led: { ... },        // configuration for this part
//          },
//          wires: [               // wires (optional)
//            {from: 'main.D2', to: 'led.anode'},
//            {from: 'main.GND#2', to: 'led.cathode'},
//          ],
//       },
//       'microbit': {},           // an empty value also works (and uses defaults)
//     },
//   }
async function setupTour(config) {
	let boards = [];
	for (let boardName in config.boards) {
		boards.push(boardName);
	}
	function createState(board) {
		let state = {
			target: board,
			parts: {
				main: {
					location: `parts/${board}.json`,
					x: 0,
					y: 0,
				}
			},
			wires: [],
		}
		let boardConfig = config.boards[board];
		if (boardConfig.code) {
			state.code = boardConfig.code.trim();
		} else {
			state.code = config.code.trim();
		}
		for (let key in boardConfig.parts) {
			let value = structuredClone(boardConfig.parts[key]);
			value.x = value.x || 0;
			value.y = value.y || 0;
			state.parts[key] = value;
		}
		for (let wire of (boardConfig.wires) || []) {
			state.wires.push(wire);
		}

		return state;
	}

	let root = document.querySelector('.simulator');
	let editor = new Editor(document.querySelector('.playground-editor'));
	let firmwareButton = document.querySelector('.playground-btn-flash');
	let resetButton = document.querySelector('.playground-btn-reset');

	// Pick default board.
	let currentBoard = boards[0];
	if (localStorage.tinygo_tour_board in config.boards) {
		// The user previously picked a different board.
		currentBoard = localStorage.tinygo_tour_board;
	}

	let boardButton = document.querySelector('#board .dropdown-toggle');
	let boardMenu = document.querySelector('#board .dropdown-menu');
	for (let board of boards) {
		let boardName = boardNames[board] || board;
		let a = document.createElement('a');
		a.textContent = boardName;
		a.classList.add('dropdown-item');
		if (board === currentBoard) {
			a.classList.add('active');
			boardButton.textContent = boardName;
		}
		a.href = "";
		a.addEventListener('click', async e => {
			// Select a different board.
			localStorage.tinygo_tour_board = board;
			e.preventDefault();
			boardMenu.querySelector('a.dropdown-item.active').classList.remove('active');
			a.classList.add('active');
			boardButton.textContent = boardName;
			currentBoard = board;
			state = createState(currentBoard);
			editor.setText(state.code);
			await simulator.setState(state, currentBoard);
		})
		boardMenu.appendChild(a);
	}

	// Load simulator.
	let state = createState(currentBoard);
	editor.setText(state.code);
	let simulator = new Simulator({
		root: root,
		editor: editor,
		features: [],
		firmwareButton: firmwareButton,
		baseURL: new URL('/playground/', document.baseURI),
		apiURL: PLAYGROUND_API,
	});
	await simulator.setState(state, state.target);

	resetButton.disabled = false;
	resetButton.addEventListener('click', async e => {
		resetButton.disabled = true;
		state = createState(currentBoard);
		editor.setText(state.code);
		await simulator.setState(state, currentBoard);
		resetButton.disabled = false;
	});
}
