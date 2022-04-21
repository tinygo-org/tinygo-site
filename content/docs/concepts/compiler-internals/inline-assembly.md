---
title: "Inline assembly"
weight: 4
---

**Warning**: inline assembly is an unstable feature and may change or be removed entirely in the future! If you want to use inline assembly, it's better to use CGo instead (and [use inline assembly in C directly]({{< relref "#cgo" >}})).

The device-specific packages like `device/avr` and `device/arm` provide `Asm` functions which you can use to write inline assembly:

```go
arm.Asm("wfi")
```

You can also pass parameters to the inline assembly:

```go
result := arm.AsmFull(`
    add {}, {value}, #3
`, map[string]interface{}{
    "value": 42,
})
println("result:", int(result))
```

At the moment, only integer types are supported as operands.

## Inline assembly using CGo {#cgo}

Inline assembly directly in TinyGo is unstable. A more stable way to use inline assembly is through CGo:

```go
package main

/*
__attribute__((always_inline))
int inlineasm(int value) {
        int result;
        __asm__("add %[result], %[value], #3"
                : [result]"=r"(result)
                : [value]"r"(value));
        return result;
};
*/
import "C"

func main() {
        result := C.inlineasm(42)
        println("result:", int(result))
}
```

Because the `inlineasm` C function is inlined into the Go main function, there is no additional CGo overhead.

Inline assembly in C also has much more features that allow it to be used more safely (for example, by correctly marking registers as clobbered). A good introduction to inline assembly is this [ARM GCC inline assembly cookbook](http://www.ethernut.de/en/documents/arm-inline-asm.html).
