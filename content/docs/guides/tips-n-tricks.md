---
title: "Tips, Tricks and Gotchas"
weight: 4
description: |
  Tips and tricks for small places.
  How to write efficient embedded code and avoid common mistakes.
---

## Ensure Concurrency

TinyGo code runs on a single core, in a single thread (think `GOMAXPROCS=1`).
Since scheduling in TinyGo is [cooperative](https://en.wikipedia.org/wiki/Cooperative_multitasking), a goroutine that never does IO or other blocking calls (f.ex. `time.Sleep()`) will lock the single available thread only for itself and never allow other goroutines to execute. In such cases, you can use `runtime.Gosched()` as a workaround.

```
package main

import (
    "fmt"
    "time"
    "sync/atomic"
    "runtime"
)

func main() {

    var ops uint64 = 0
    for i := 0; i < 50; i++ {
        go func() {
            for {
                atomic.AddUint64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
```

## Save Memory

Use the fact [slices are descriptors of array segments](https://go.dev/blog/slices-intro#slice-internals) to save memory and avoid unnecessary allocations.
Allocate once an array of maximum size you ever need and slice it in your methods to the required sizes as you go.

```
type MyType struct {
    // ...other fields
    buf [6]byte
}

func (t *MyType) SomeMethod() {
    tmpBuf := t.buf[:2] // this method needs a buffer of size 2
    // ... use the buffer
}
```

You may even find yourself having two or more slices pointing at diferent regions of the same array simultaneously, if you are careful.

```
buf1 := t.buf[:2]
buf2 := t.buf[2:6]
```

For more recipes on how to avoid or minimize allocations while working with slices, please see [additional tricks](https://github.com/golang/go/wiki/SliceTricks#additional-tricks) section of [slice tricks](https://github.com/golang/go/wiki/SliceTricks) page.

Familiarize yourself with [heap allocation concept](../../concepts/compiler-internals/heap-allocation/) and use it to your advantage.
