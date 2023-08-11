---
title: "Watchdog"
weight: 2
description: |
  How to interact with internal watchdog devices.
---

For API documentation, see the [machine API](../../../reference/machine#watchdog)

Many microcontrollers include some form of watchdog peripheral.  Watchdog peripherals are used to restart the microprocessor in the event of a fault condition (typically a software bug).  Watchdogs are typically based around a fixed-frequency hardware timer that counts down from a configured value to zero.  On reaching zero a reset is triggered.  To prevent a reset the software must periodically update the watchdog, causing the counter to return to it's initial value.

There are multiple variants of watchdog functionality. Some do not permit the watchdog to be disabled once enabled, others provide for multiple 'sources' enabling applications to have hardware protection over multiple subsystems.  The maximum time for the watchdog countdown can vary from hours to seconds on different hardware.

Tinygo provides an abstraction supported by most watchdog hardware:

1. No ability to disable a watchdog after it is enabled.

2. Where supported, the watchdog is automatically disabled if a debugger is attached.

3. The watchdog is not disabled during sleep (`time.Sleep`) or blocking I/O, so critical loops must take care to not sleep for too long or block on I/O indefinitely.

4. A single 'source' is supported on all platforms with Watchdog support.

5. The CPU is not guaranteed to reset exactly at the requested time.  Some hardware does not allow precise timing - the next largest supported timeout is used instead.

A typical use of Watchdog support in a simple app would look something like this:

```go
void main() {
    
    // ...

    config := machine.WatchdogConfig{
		TimeoutMillis: 1000,
	}
    machine.Watchdog.Configure(config)

    machine.Watchdog.Start()

    for {
        machine.Watchdog.Update()

        //
        // Main logic - poll GPIO, write to display, etc
        //
	}
}
```

In this example, the watchdog is configured for updates at least every 1s.  One second is a safe default value supported on all hardware.  Consult `machine.WatchdogMaxTimeout` for the hardware specific maximum value.