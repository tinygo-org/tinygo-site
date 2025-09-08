---
title: "Configuring the accelerometer"
description: "Turn on the accelerometer and start reading values"
weight: 2
---

The accelerometer that we are using, LIS3DH, starts in a low power mode. That's not the case for all such sensors, many start up in an already enabled state that uses more power. Starting up to a low power mode means that we can ignore the accelerometer when we don't use it and it won't use more power than necessary, but it does mean we need to configure it before we can use it.

The beginning of the code on the right is the same as before, but now we do some magic:

```go
	err = i2c.Tx(0x18, []byte{0x20, 0b0100_0111}, nil)
	if err != nil {
		println("got error while configuring the LIS3DH:", err.Error())
		return
	}
````

That's a lot of magic numbers! Let's go through them:

  * 0x18 is the I2C address, like before. It varies a bit depending on the electronics, but it will always be 0x18 or 0x19 for this particular sensor.
  * 0x20 is the register address.
  * 0b0100_0111 (8 bits of binary data) is the value that we're going to write to register 0x20.

Internally, most I2C devices are organized as a number of registers. You can think of it as a slice of bytes, one that's typically just 128 or so bytes long. This slice contains both configuration data (which can typically be read and written), and sensor output data (which is typically read-only). In one transaction, you can either read or write data to the I2C device:

  * For reading, you send a single byte (the register address) to the device and then read the response values into the `r` byte slice. This is what we did in the previous step of this tour, to read the `WHO_AM_I` register.
  * For writing, you send a slice of bytes to the device and read nothing back. The first byte in the slice is the address, the subsequent bytes are the values to write to the registers starting at the given address. This is what we do above.

To know what register 0x20 is, and what the value 0b0100_0111 means, we need to take another look at the [datasheet](https://cdn-shop.adafruit.com/datasheets/LIS3DH.pdf#page=29):

![Section of the LIS3DH datasheet describing the "CTRL_REG1" register](/images/i2c-lis3dh-config.png)

This section describes register CTRL_REG1 at address 20h. "20h" means "hexadecimal 20", which is the same as 0x20 in Go.

The register itself contains 8 bits: 4 bits for "ODR", one for "LPen", and 3 bits for the three axes (Z, Y, X).

Decoding our value 0b0100_0111 (starting from the lowest bits on the right), we see the following:

  * The lowest 3 bits are the three axes: X, Y, and Z. Of course we want to enable all of them but it is possible to disable some of them if you need. Who knows, it might reduce power usage a little (even though this chip already uses very little power).
  * The next bit is LPen, or "low power mode enable". This bit is zero, so that means low power mode is disabled (if we enable it, we get less precision from the sensor).
  * The upper bits, 0100, are the four ODR bits. These four bits are described in table 26 above. Looking at the table, "0 1 0 0" means that we will enable the accelerometer to run at 50Hz (measuring acceleration 50 times per second). This seems like a reasonable speed for our purposes.

Now if you look at the "Properties" tab on the right, you can see that the LIS3DH is indeed configured as "Normal mode 50Hz, enabled axes X, Y, Z", exactly what we wanted! If you go back one step (where we read the `WHO_AM_I` register) you can see it is set to "Power down mode" instead since that is the power-on default. Also, if you go to the "Power" tab, you can see that the LIS3DH now uses 11µA instead of 0.5µA as it did in power down mode.

Try changing the value 0b0100_0111 to something else, and see what happens in the simulator!

<script type="module">
let code = `
package main

import (
	"machine"
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

	println("accelerometer configured!")
}
`;
import {setupLIS3DH} from "/tour-lis3dh.js";
setupLIS3DH(code);
</script>
