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

- Define a `Config` struct for peripherals with many configuration options. See example of a Config struct in action: https://github.com/soypat/lora/blob/main/lora.go
    - Define new types for configuration enums. Try to use values in the datasheet, if a register has 3 possible values for a configuration that a user may want to set in the `Config` define a type for that register value and provide the 3 exported values with that type as package level constants.
    - **Avoid**: [Functional options](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis) or interface options. These carry with them a higher overhead in both performance, memory and binary size. A peripheral API will never change after being manufactured so most benefits of this pattern are lost
    - Define a `DefaultConfig` at the package level so that the user may easily instantiate a configuration. Maybe add a parameter or two to `DefaultConfig` to allow the user to easily tweak the most important parameters that they may want to change.
    - Are some configuration parameters codependent and require calculating? Add a method on the `Config` type to set them consistently! Adding a `Validation() error` method is also useful during the `func (d *Device) Configure(cfg Config) error` call as well as letting users read the validation logic and understand what is a correct configuration. It's a great source for intuition on how the driver works!
- **Avoid**: Floating point operations. 
    - Peripherals rarely work with or return floating point representations of their measurements. More often than not you'll find they work with integers. There's a reason for this: integer operation hardware is widespread. Most small processors do not support floating point operations, and even if they do it is much slower than modern desktop PCs compared with integer operations
    - There is huge difficulty in working exclusively with integer operations instead of floats without losing precision. If possible define the API return values of sensors with integers and work internally with floats if that is easier for you. This will allow the API to remove the internal floating point bits in the future while not breaking existing users. The usage of floats vs integers has been discussed here: https://github.com/tinygo-org/drivers/pull/345

### Constrain memory use at compile time

TinyGo allows for heap allocations and that is a wonderful feature but can also subject embedded Go users to some pain
if the drivers they use make heavy use of heap allocations. 
Heap allocations over time can crash your program if memory becomes sufficiently fragmented. This is a huge problem for users who want their programs to run a long time.

Below is a short recap on heap allocations. There is also a separate section on this page, see the compiler internal page [Heap Allocations]({{<ref "heap-allocation.md">}}) for a more in depth dive.


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
Why is it important to know slices are referential? Well it is because the compiler is aware of this and will check
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
    // The assignment below is equivalent to a uint16 assignment to the compiler.
    // Static arrays in Go are not pointers so this will always be on the stack.
    d.bufWrite = [2]byte{addr, 0} 
    err= d.spi.Tx(d.bufWrite[:], d.bufRead[:])
    return d.bufRead[1], err
}


func (d *Device) writeRegister(addr, newValue uint8) error {
    d.bufWrite = [2]byte{addr, newValue}
    return d.spi.Tx(d.bufWrite[:], d.bufRead[:])
}
```
| Note on Go's array type in TinyGo |
|------|
| Note the use of static arrays types like `[2]byte`. "Arrays", as they are called in Go, if small enough will always be stack allocated. TinyGo currently has a **Max Stack Object Size** configuration which runs around the 256 byte size. That said, this fact is not relevant in this case since the array is attached to the `Device` struct which almost certainly is heap allocated. Most drivers, if not all, initialize and use a pointer type (`*Device`). If the `Device` is heap allocated then the arrays it contains will also be heap allocated- and this is OK! Remember, we are trying to constrain our allocations, not completely eliminate heap allocations. We win if during the lifetime of the program there are **ZERO** heap allocations. Allowing heap allocations at initialization is a good practice and actually encouraged in NASA's coding standards for extraplanetary missions. |

This can solve many such cases, but there are few cases where data to be read/written is also variable length.
The semantics for reading and writing data over the SPI bus are hypothetical, what is important is to take note on how to 
deal with variable length buffers. See [SliceTricks](https://go.dev/wiki/SliceTricks#additional-tricks) for more tips on slice manipulation.
```go
const maxCommSize = 32 // Never send more than 32 bytes in a single transaction.

type Device struct {
    spi   drivers.SPI
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
    n := copy(d.wdata[1:], data) // First byte reserved for address.
    d.wdata[0] = addr
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

### Multi-protocol devices
Some devices support SPI and I2C interfaces (UART, RS485, etc.). What's the ideal approach to this? Should one generate a separate driver for each protocol?

The answer is, as usual in engineering "it depends". That said more often than not a peripheral which supports different protocols shares a whole lot of 
functioning. Take for example the Honeywell HSC TruStability pressure sensors. They are a VERY simple pressure sensor which come in SPI and I2C varieties.
One can write a driver that shares the underlying shared logic like so:

```go
package hsc

// device implements the shared logic between Honeywell HSC pressure sensor varieties.
type device struct {
    // calibration factors.
    k1, k2, k3 uint32
    lastRead   uint32
}

func (d *device) update(bridgeData []byte) error {
    if bridgeData[0] & (1<<7) != 0 {
        return errors.New("data unavailable")
    }
    d.lastRead = uint32(bridgeData[0])*d.k1 + uint32(bridgeData[1])*d.k2 + uint32(bridgeData[2])*d.k3
    return nil
}

// Pressure implements the drivers.Sensor API for pressure sensors.
func (d *device) Pressure() uint32 {
    return d.lastRead
}

type DeviceSPI struct {
    device
    spi        drivers.SPI
    wbuf, rbuf [4]byte
}

// Update implements the drivers.Sensor API.
type (d *DeviceSPI) Update(which drivers.Measurement) error {
    if which&drivers.Pressure == 0 {
        return nil
    }
    err := d.spi.Tx(d.wbuf[:], d.rbuf[:])
    if err != nil {
        return err
    }
    return d.update(d.rbuf[1:4])
}
```
| Note on `drivers.Sensor` API (`Update`/`Pressure`) |
|----|
| When designing a sensor API be sure to read [this PR](https://github.com/tinygo-org/drivers/pull/345). It is a good practice to separate the update from reading sensor data to have a way of allowing users of the sensor API to either do a calculation with the most recently read pressure or update the latest pressure value. This is because it is costly to update a pressure value and at times some sensors can also get multiple different readings like pressure and humidity in a single call. This API exposes this functionality in a composable and performant way. Users are responsible of calling `Update` regularly and having applications call the sensor value method (`Pressure`/`AngularVelocity`/etc) |

There are times where the logic is much more complex and having separate types for different protocols would make the driver implementation much harder to read. 
In these cases it is justified to have a single `Device` type which has a generic bus interface.

```go
func NewI2C(i2c drivers.I2C) *Device {
    return &Device{
        bus:bus{i2c:i2c, isSPI:false}
    }
}

func NewSPI(spi drivers.SPI) *Device {
    return &Device{
        bus:bus{spi:spi, isSPI:true}
    }
}

type Device struct {
    bus
    // ...
}

type bus struct {
    spi drivers.SPI
    i2c drivers.I2C
    wbuf, rbuf [4]byte
    isSPI bool
}

func (b *bus) ReadRegister(addr uint8) (uint8, error) {
    if b.isSPI {
        // SPI logic to read register.
    } else {
        // I2C logic to read register.
    }
}
```



## Hot loops
Hot loops are a natural part of writing drivers. Sometimes you need to check a driver state before proceeding like during initialization.
If hot loops run a couple of times the impact is negligible (but one still must mitigate risks of hot loops!). An issue does arise when the loop happens thousands  or even millions of times.
Below is an example of a hot loop in which the developer is debugging to see how many times the hot loop runs:

```go
func (d *Device) Configure(cfg Config) error {
    d.Reset()
    start := time.Now()
    hotloops := 0
    for !d.IsConnected() {
        hotloops++
    }
    // Make sure to not print in hot loop when measuring!
    println(time.Since(start).String(), "elapsed during hot loop in Configure with calls done=", hotloops)
    // ...
}
```

After measuring one can evaluate the impact of a hot loop. When working with peripheral drivers one usually looks at the time 
it took for the hot loop to finish. This is the amount of time the CPU thread was blocked and unable to do other possibly important work.
Unmitigated blocking hot loops raise the CPU usage and can bring the program to a standstill. In certain cases where the peripheral hardware
is inconsistent or unstable it can freeze your computer.

One mitigates hot loop effects by yielding the processor inside the hot loop and also setting a maximum number
of times to retry the hot loop before "giving up" and returning an error. A good number of retries to use is x2 the number of retries measured and at least 10.
Remember the number of retries performed will vary from CPU to CPU since some CPUs are faster. When measuring use the highest communication bitrate for the peripheral to get the best case scenario.
```go
// Make sure `retries` makes sense. Use a larger number than the maximum typical hotloops done.
retries := 1000 
for retries > 0 && !ready() {
    runtime.Gosched() // For short hot loops around <1ms
    retries--
}
if retries <= 0 {
    return errFailedReady
}

for retries > 0 && !longready() {
    time.Sleep(typicalLongDuration/5)
    retries--
}
if retries <= 0 {
    return errFailedReady
}

// Very friendly to CPU mitigation. We sleep for the typical duration
// and then start the hot loop.
// This avoids using the CPU during the time the peripheral is almost certainly not ready.
time.Sleep(reallyLongDuration)
for retries > 0 && !reallyLongReady() {
    time.Sleep(reallyLongDuration/10)
    retries--
}
if retries <= 0 {
    return errFailedReady
}
```
Note the retry check is in front of the hot loop check in the `for` statement condition: this is for a reason! If not ordered that way then maybe the hot loop check passes but retry fails resulting in a zero retry and a returned error even though the check passed.

## Miscellaneous tips
### Consolidate I/O into as few dynamic calls as possible

Consider the code below, it does some configuration over an SPI bus and a pin which apparently selects if spi bytes are data or commands.
```go
func (d *Device) Configure(cfg Config) {
    d.Command(CONFIG_BYTE)
    d.Data(0x01)
    d.Data(0x02)
    d.Data(0x03)
    d.Data(0x04)
    d.Data(0x05)
}

func (d *Device) Data(data uint8) {
    d.pinData(true)
    d.spi.Tx([]byte{data}, nil)
}

func (d *Device) Command(cmd uint8) {
    d.pinData(false)
    d.spi.Tx([]byte{cmd}, nil)
}
```

Note the string of `d.Data` calls, they all do two I/O operations: turn the data pin on and send a single byte over SPI.
Below are some problems that can be solved with this approach regarding I/O overhead:
- Successive `Data` calls turn a pin on that has remained on since last `Data` call
- Every call of Data does a dynamic call to the SPI `Tx` method

If we are aware of how a `SPI` bus works then we'll know we can concatenate the single byte `Tx` calls into a single multi-byte `Tx` call.

```go
func (d *Device) Configure(cfg Config) {
    d.Command(CONFIG_BYTE)
    d.Data(0x01, 0x02, 0x03, 0x04, 0x05)
}
func (d *Device) Data(data ...byte) {
    d.pinData(true)
    d.spi.Tx(data, nil)
}
```
The above code is equivalent in functionality. One may worry we are not calling `d.pinData(true)` between bytes sent over SPI until one realizes calling `d.pinData(true)` after the first call is completely inneffective. Once a pin turns on it stays that way until it is turned off with `d.pinData(false)`.

To get the gold star though we'd want to eliminate heap allocations. Since `Data` may receive thousands of bytes at a time in some cases we define a helper method:

```go
func (d *Device) Configure(cfg Config) {
    d.Command(CONFIG_BYTE)
    d.dataShort(0x01, 0x02, 0x03, 0x04, 0x05)
}
func (d *Device) dataShort(data ...byte) {
    n := copy(d.buf[:], data) // Buf is a static array that is the max size of calls performed by the library internally.
    d.Data(d.buf[:n])
}
func (d *Device) Data(data []byte) {
    d.pinData(true)
    d.spi.Tx(data, nil)
}
func (d *Device) Command(cmd uint8) {
    d.pinData(false)
    d.buf[0] = cmd
    d.spi.Tx(d.buf[:1], nil)
}
```