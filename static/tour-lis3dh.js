import { setupTour } from '/tour.js';

export function setupLIS3DH(code) {
  let imuConfig = {
  	humanName: 'Adafruit LIS3DH',
  	svg: 'adafruit-lis3dh.svg',
  };
  setupTour({
  	boards: {
  		'arduino-nano33': {
  			code: code.replace('SCL_PIN', 'A5').replace('SDA_PIN', 'A4'),
  			parts: {
  				main: {
  					y: -10,
  				},
  				imu: {
  					config: imuConfig,
  					location: 'parts/adafruit-lis3dh.json',
  					x: 5,
  					y: 10,
  					rotation: 180,
  				}
  			},
  			wires: [
  				{from: 'main.A5', to: 'imu.SCL'},
  				{from: 'main.A4', to: 'imu.SDA'},
  			],
  		},
  		'arduino': {
  			code: code.replace('SCL_PIN', 'ADC5').replace('SDA_PIN', 'ADC4'),
  			parts: {
  				main: {
  					y: 10,
  				},
  				imu: {
  					config: imuConfig,
  					location: 'parts/adafruit-lis3dh.json',
  					x: -13,
  					y: -30,
  				},
  			},
  			wires: [
  				{from: 'main.A5#2', to: 'imu.SCL'},
  				{from: 'main.A4#2', to: 'imu.SDA'},
  			],
  		},
  		'circuitplay-bluefruit': {
    		code: code.replace('SCL_PIN', 'SCL1_PIN').replace('SDA_PIN', 'SDA1_PIN').replaceAll('Tx(0x18', 'Tx(0x19'),
  		},
  		'circuitplay-express': {
  			code: code.replace('SCL_PIN', 'SCL1_PIN').replace('SDA_PIN', 'SDA1_PIN').replace('machine.I2C0', 'machine.I2C1').replaceAll('Tx(0x18', 'Tx(0x19'),
  		},
  		'microbit': {
  			code: code.replace('SCL_PIN', 'P0').replace('SDA_PIN', 'P1'),
  			parts: {
  			  main: {
  					x: 15,
  				},
  				imu: {
  					config: imuConfig,
  					location: 'parts/adafruit-lis3dh.json',
  					y: -10,
  					x: -30,
  				},
  			},
  			wires: [
  				{from: 'main.P0', to: 'imu.SCL'},
  				{from: 'main.P1', to: 'imu.SDA'},
  			],
  		},
  		'gopher-badge': {
    		code: code.replace('SCL_PIN', 'I2C0_SCL_PIN').replace('SDA_PIN', 'I2C0_SDA_PIN'),
  		},
  		'reelboard': {
  			code: code.replace('SCL_PIN', 'P0_27').replace('SDA_PIN', 'P0_26'),
  			parts: {
  				main: {
  					x: -10,
  				},
  				imu: {
  					config: imuConfig,
  					location: 'parts/adafruit-lis3dh.json',
  					x: 50,
  				},
  			},
  			wires: [
  				{from: 'main.P19', to: 'imu.SCL'},
  				{from: 'main.P20', to: 'imu.SDA'},
  			],
  		},
  		'pico': {
  			code: code.replace('SCL_PIN', 'GP17').replace('SDA_PIN', 'GP16'),
  			parts: {
  				main: {
  					y: 12,
  				},
  				imu: {
  					config: imuConfig,
  					location: 'parts/adafruit-lis3dh.json',
  					x: 21,
  					y: -13,
  				},
  			},
  			wires: [
  				{from: 'main.GP17', to: 'imu.SCL'},
  				{from: 'main.GP16', to: 'imu.SDA'},
  			],
  		},
  		'hifive1b': {
  			code: code.replace('SCL_PIN', 'D19').replace('SDA_PIN', 'D18'),
  			parts: {
  				main: {
  					x: -10,
  				},
  				imu: {
  					config: imuConfig,
  					location: 'parts/adafruit-lis3dh.json',
  					x: 45,
  					rotation: 180,
  				},
  			},
  			wires: [
  				{from: 'main.D19', to: 'imu.SCL'},
  				{from: 'main.D18', to: 'imu.SDA'},
  			],
  		},
  	}});
}
