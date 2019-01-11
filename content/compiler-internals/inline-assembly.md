---
title: "Inline assembly"
weight: 4
---

The device-specific packages like `device/avr` and `device/arm` provide `Asm` functions which you can use to write inline assembly:

```go
arm.Asm("wfi")
```

You can also pass parameters to the inline assembly:

```go
var result int32
arm.AsmFull(`
    lsls  {value}, #1
    str   {value}, {result}
`, map[string]interface{}{
    "value":  42,
    "result": &result,
})
println("result:", result)
```

In general, types are autodetected. That is, integer types are passed as raw registers and pointer types are passed as memory locations. This means you can't easily do pointer arithmetic. To do that, convert a pointer value to a `uintptr`.

Inline assembly support is expected to change in the future and may change in a backwards-incompatible manner.
