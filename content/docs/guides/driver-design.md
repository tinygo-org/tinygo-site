---
title: "Driver Design and Development"
weight: 10
description: |
  Developing drivers from scratch for embedded Go
---
# Driver Design with TinyGo

## Driver for a peripheral
When designing a peripheral driver with which one interacts with over a communications bus one can usually adhere to the following rules

- One package per peripheral
    - So if your peripheral is the MPU6050 accelerometer create a new package called `mpu6050`
    - Sometimes the logic can be shared between a family of peripherals such as the AS5601 and AS5600. In these cases one usually creates a package that contains this shared logic and x's the shared character position, so `as560x`

- Define a base type `Device` for the handle to the peripheral. All interaction with the peripheral will be via methods on this type.
    - Store HAL required for peripheral operation here such as communication buses (I2C, SPI drivers) and pins
    - **Avoid**: using TinyGo `machine` package constructs in your driver (read as *don't import `"machine"`*). Users of the Go language (not only TinyGo) also consume the drivers package and they can't compile a program that uses the `"machine"` package. To abstract `machine.Pin` one can define `type PinOutput func(level bool)` and `type PinInput func() bool` functional interfaces.

- Define a `Config` struct for peripherals with many configuration options
    - Define new types for configuration enums. Try to use values in the datasheet, if a register has 3 possible values for a configuration that a user may want to set in the `Config` define a type for that register value and provide the 3 exported values with that type as package level constants.
    - **Avoid**: [Functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) or interface options. These carry with them a higher overhead in both performance, memory and binary size. A peripheral API will never change after being manufactured so most benefits of this pattern are lost

### Constrain memory use at compile time

TinyGo allows for heap allocations and that is a wonderful feature but can also subject embedded Go users to some pain
if the drivers they use make heavy use of heap allocations. 
Heap allocations over time can crash your program if memory becomes sufficiently fragmented. This is a huge problem for users who want their programs to run a long time.


#### Heap allocations: an introduction
The most common source of heap allocations in TinyGo drivers are in bus communications. Take for example the following code:
```go
type Device struct {
    spi drivers.SPI
}

func (d *Device) readRegister(addr uint8) (value uint8, err error) {
    var data [2]byte
    err= d.spi.Tx([]byte{addr, 0}, data[:])
    return data[1], err
}


func (d *Device) writeRegister(addr, newValue uint8) error {
    var data [2]byte
    return d.spi.Tx([]byte{addr, newvalue}, data[:])
}
```

The code above takes two buffers and passes them to the `drivers.SPI` interface `Tx` method as slices.
Slices are referential data structures, which is to say, the **data structure contains a pointer**.
This is a pointer to the start of the data to be sent over the bus.

In go we are not aware of this since we never interact with the low level representation of a slice:
```go
type slice struct {
    data      *type // Points to the first element of the slice:  &s[0]
    length    int   // This is returned when calling builtin len: len(s)
    capacity  int   // This is returned when calling builtin cap: cap(s)
}
```
Why is it important to know slices are referential? Well it is because the compile is aware of this and will check
referential arguments to function and "escape" them if it can't prove their reference (pointer) is not held by the calling function.
When a reference is escaped it will be forced to be heap allocated by the compiler.

Returning to the `Device` example above we started with: any referential argument to `spi.Tx` will escape since the `Tx` method is unknown at compile time since
`spi` is an interface. So every time we call `writeRegister` or `readRegister`, at worst two heap allocations will be performed, one being the `data [2]byte` array
and the second being the inline composite literal slice declaration `[]byte{x, x}`. Usually heap allocators have limitations on size allocated so it is likely that more than 4 bytes are allocated for every call.

#### Heap allocations: Mitigations
Luckily it can be relatively easy to mitigate heap allocations and eliminate them altogether. To do so with a driver that uses slices for bus communications one can include the array memory within the device struct:

```go
type Device struct {
    spi      drivers.SPI
    bufRead  [2]byte
    bufWrite [2]byte
}

func (d *Device) readRegister(addr uint8) (value uint8, err error) {
    d.bufWrite = [2]byte{addr, 0}
    err= d.spi.Tx(d.bufWrite[:], d.bufRead[:])
    return d.bufRead[1], err
}


func (d *Device) writeRegister(addr, newValue uint8) error {
    d.bufWrite = [2]byte{addr, newValue}
    return d.spi.Tx(d.bufWrite[:], d.bufRead[:])
}
```

This can solve many such cases, but there are few cases where data to be read/written is also variable length.
The semantics for reading and writing data over the SPI bus are hypothetical, what is important is to take note on how to 
deal with variable length buffers.
```go
const maxCommSize = 32 // Never send more than 32 bytes in a single transaction.

type Device struct {
    spi drivers.SPI
    wdata [maxCommSize]byte
    rdata [maxCommSize]byte
}

func (d *Device) Configure() error {
    // Since the argument buffer is not passed into Tx call,
    // data does NOT escape and may be stack allocated.
    data := []byte{1, 2, 3, 4}
    return d.writeData(CONFIG_REG, data)
}

func (d *Device) writeData(addr uint8, data []byte) error {
    d.data[0] = addr
    n := copy(d.data[1:], data)
    return d.spi.Tx(d.wdata[:n+1], d.rdata[:n+1])
}

func (d *Device) readData(addr uint8, data []byte) error {
    for i := range d.wdata {
        d.wdata[i] = 0 // Clear write data buffer.
    }
    dlen := len(data)
    d.wdata[0] = addr
    err := d.spi.Tx(d.wdata[:dlen+1], d.rdata[:dlen+1])
    if err != nil {
        return err
    }
    copy(data, d.rdata[1:])
    return nil
}
```
