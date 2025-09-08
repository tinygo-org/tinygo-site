---
title: "Reading acceleration values"
description: "Read the X, Y, and Z values from the accelerometer."
weight: 3
---

Now that we have the accelerometer configured, we can start reading the values it reads.

Most of the code is the same as before, but the reading part is new:

```go
		// Read the acceleration data.
		w := []byte{0x28|0x80}
		r := make([]byte, 6)
		err := i2c.Tx(0x18, w, r)
```

Here we read 6 bytes starting with register 0x28. We'll get back to why we need to set the highest bit later (through `|0x80`). First let's take a look at the [datasheet](https://cdn-shop.adafruit.com/datasheets/LIS3DH.pdf#page=33) again:

![...](/images/i2c-lis3dh-acceldata.png)

So here we have 6 registers that contain the acceleration data in the 3 directions (X, Y, Z) where each direction takes up two registers. Since each register is 8 bits, and the data itself is 10 bits, two registers are needed for each direction.

Decoding it is a little bit tricky though, if you're not used to bitwise operations:

```go
			x := int16(r[0]) | int16(r[1])<<8
			y := int16(r[2]) | int16(r[3])<<8
			z := int16(r[4]) | int16(r[5])<<8
			println("x, y, z:", x, y, z)
````

This combines the two 8 bit values into a single 16-bit signed integer. If we take the X axis for example, the datasheet implies the first register (`OUT_X_L`) is low, since it has "L" in the name. (Yes, it's not the most helpful datasheet in this regard). The next register (`OUT_X_H`) contains the upper bits. So the way to convert them to a single 16-bit value is to convert them both individually to a 16-bit value (which keeps the lower bits but fills the missing bits with zero), then _shift_ the bits of `OUT_X_H` to the high position, and OR them together. This gives us a single 16-bit value, which is the signed X axis.

On the right you can see the values we read from the accelerometer. On a desktop computer you will see simulated values, as if the board is lying flat on a surface. However, if you are on a supported mobile device (phone, table), you will actually be able to see the real acceleration values of your device!

Some things to note:

  * **Gravity** is also measured. This is just how physics works: when the sensor is lying still, it shows values as if it was accelerating at *g* speed away from earth. (Unless you're doing this in a zero-gravity environment, of course). There are various algorithms to filter out the gravity component, which will usually be needed for practical applications.
  * There is **noise** in the output. Every sensor has noise, this one included. The simulator includes the noise to be more realistic, and if you use such a sensor in a project you need to be aware that such noise exists.
  * The sensor can be put on a board in various ways, for example it can either be on top or at the bottom. This will of course impact the measurements.

## What's up with that upper bit?

While for I2C devices you can read multiple sequential registers by writing the register address and then reading multiple values back, the LIS3DH is a bit different. By default it will return the same register on every read. To get the behavior we want, we need to set the highest bit of the register address to one, which gets the sensor to return values incrementally.

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

	err = i2c.Tx(0x18, []byte{0x20, 0b0100_0111}, nil)
	if err != nil {
		println("got error while configuring the LIS3DH:", err.Error())
		return
	}

	for {
		// Read the acceleration data.
		w := []byte{0x28|0x80}
		r := make([]byte, 6)
		err := i2c.Tx(0x18, w, r)
		if err != nil {
			println("got error while reading values:", err.Error())
		} else {
			x := int16(r[0]) | int16(r[1])<<8
			y := int16(r[2]) | int16(r[3])<<8
			z := int16(r[4]) | int16(r[5])<<8
			println("x, y, z:", x, y, z)
		}

		time.Sleep(time.Second)
	}
}
`;
import {setupLIS3DH} from "/tour-lis3dh.js";
setupLIS3DH(code);
</script>
