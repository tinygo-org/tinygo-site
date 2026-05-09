---
title: "Adding runtime support for new processors"
weight: 1
description: >
  How to add runtime support to a new processor.
---

In order for a processor to support `time.Sleep()` and other things you normally expect, there are some functions that need to be implemented in the runtime for that processor.

## runtime timers/sleep

Functions that need to be implemented:

```go
// timeUnit defines the type of the unit of time used for that processor per tick.
// can be measured in nanoseconds, microseconds, or milliseconds.
// the important part is that it be consistent for any particular processor.
type timeUnit int64

// ticksToNanoseconds converts ticks to nanoseconds
func ticksToNanoseconds(ticks timeUnit) int64

// nanosecondsToTicks converts nanoseconds  to ticks
func nanosecondsToTicks(ns int64) timeUnit

// sleepTicks should sleep for d number of ticks.
func sleepTicks(d timeUnit)

// ticks returns the elapsed time since reset as a number of timeUnit ticks
func ticks() timeUnit
```

Defining the `timeUnit` is crucial, since the unit of time for most processors is usually dependent on the system clock settings.

Based on that initial information, the functions `ticksToNanoseconds()` and `nanosecondsToTicks()` can be implemented in a fairly straightforward fashion.

It is with the function `sleepTicks()` that implementation details become very important. This function needs to be able to yield the processor to perform other tasks while it is waiting, in order for it to work the way that you expect with Go routines.

Lastly the `ticks()` function is needed, in order for many `time` package functions to work.
