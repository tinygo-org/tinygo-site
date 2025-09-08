---
title: "Get started with I2C"
description: "Get started with the LIS3DH by verifying the connection."
weight: 1
---

The LIS3DH sensor can measure acceleration and changes in temperature. But to get this data from the sensor, we need a way to establish communication. We'll be using the I2C protocol, which is the most common way to communicate with low-speed external devices.

```go
	i2c := machine.I2C0
	err := i2c.Configure(machine.I2CConfig{
		SCL: machine.GP17,
		SDA: machine.GP16,
	})
```

To start using I2C, we need to configure I2C on the microcontroller. This code (for the Raspberry Pi Pico) configures GP17 and GP16 as the pins to be used. Note that you can't use all pins! Many microcontrollers only allow using specific pins for I2C. We'll get to that later.

```go
		w := []byte{0x0f}
		r := make([]byte, 1)
		err := i2c.Tx(0x18, w, r)
```

This does an actual transaction. This code will first write the I2C peripheral address (0x18 or 0x19 depending on the electronic design), then write the byte 0x0f in the `w` byte slice, and then read a single byte back to store in the `r` byte slice. The peripheral address is the one for the LIS3DH, and 0x0f is a register called `WHO_AM_I` which when read returns a fixed value to identify the chip.

![WHO_AM_I register description, with the bits in the register being "0 0 1 1 0 0 1 1"](/images/lis3dh-who-am-i.png)

Here is the description from the [datasheet](https://cdn-shop.adafruit.com/datasheets/LIS3DH.pdf#page=29). It says the returned bits are "0 0 1 1 0 0 1 1". Decoding this to decimal, we get the number 51. And that's exactly what we get back from the chip, which shows that we managed to communicate with the chip!


## Connecting I2C devices

I2C uses two wires to communicate: a clock wire (SCL) and a data wire (SDA). Connect the SCL pin on the microcontroller to the SCL pin on the sensor, and the SDA pin on the microcontroller to the SDA pin on the sensor. Also, make sure there is a ground connection between the two, otherwise communication won't work!

Apart from SCL and SDA (and a shared ground), the sensor also needs some power. This is often 3.3V, but check the datasheet to be sure. If the sensor needs 3.3V and you provide it 5V, it will likely burn out! In the simulator we skip this requirement to avoid clutter.

I2C may also need some external pullup registers, that make sure the wires stay high unless they are set to a low voltage by the controlling device. In practice, many sensors already include pullup resistors and microcontrollers often configure I2C wires with a pullup, so this may not be required in practice. It's something to be aware of when you design your own board though.

## I2C basics

I2C allows using multiple devices with just one pair of wires! This pair of wires (SCL and SDA) is called a bus, and can be connected to multiple sensors and actuators. It typically only allows one controlling device (the main microcontroller), but on this bus up to 127 or so devices can be used. These are usually sensors or other low-speed devices, but you can also configure another microcontroller as a peripheral device to allow communication with the controlling device.

To support multiple devices on a single bus, every I2C peripheral device has an address. In the case of the LIS3DH, that is 0x18 or 0x19 in hexadecimal. You can change the address by connecting the SA0/SDO pin to VCC or GND.

Communication is always initiated from the controller device (typically the microcontroller), and happens in a single transaction with a peripheral device. The controller sends the address, then sends a few bytes (which can be zero-length), and then receives a number of bytes. That all happens in a single operation called a "transaction".

In other documentation you may see the term "master" and "slave", but we chose to use the [less problematic](https://oshwa.org/resources/a-resolution-to-redefine-spi-signal-names/) "controller" and "peripheral" instead.

As usual, [SparkFun has an excellent article explaining all the ins and outs of I2C](https://learn.sparkfun.com/tutorials/i2c/all) if you want to learn more.


<script type="module">
let code = `
package main

import (
	"machine"
	"time"
)

func main() {
	i2c := machine.I2C0
	err := i2c.Configure(machine.I2CConfig{
		SCL: machine.SCL_PIN,
		SDA: machine.SDA_PIN,
	})
	if err != nil {
		println("could not configure I2C:", err.Error())
		return
	}

	for {
		// Read the "who am I" register.
		w := []byte{0x0f}
		r := make([]byte, 1)
		err := i2c.Tx(0x18, w, r)
		if err != nil {
			println("got error while reading 'who am I' register:", err.Error())
		} else {
			println("chip says:", r[0])
		}

		time.Sleep(time.Second)
	}
}
`;
import {setupLIS3DH} from "/tour-lis3dh.js";
setupLIS3DH(code);
</script>
