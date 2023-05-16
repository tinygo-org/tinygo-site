---
title: "EEPROM"
weight: 3
description: |
  How to use the internal EEPROM.
---

For API documentation, see the [machine API](../../../reference/machine#eeprom).

The EEPROM is a form of non-volatile memory that is uusually integrated along side an MCU. EEPROMs allows you to store a relatively small amount of data by allowing bytes to be erased and reprogrammed. For the most part, EEPROMs are not integrated directly onto a board, but in some cases they are provided as a standard peripheral.

Although it may seem useful to have some type of data storage for your applications, the EEPROM is often limited in the amount of write cycles that it contains. This means that after a certain amount of writes, the EEPROM will not writeable, but you may still read data from it.

## Controlling the EEPROM

Control of the EEPROM has been simplified through the `EEPROM` variable that can be accessed through the `machine` package.

Unlike most peripherals within TinyGo, you do not need to configure the EEPROM to use it.

Example usage of the EEPROM can be seen below.

```go
package main

import (
    "machine"
    "time"
)

// Address of the data to be stored.
// Addresses can be either in hex or dec.
const DataAddr = 0x10 


func main() {
    eeprom := machine.EEPROM0 // the onboard EEPROM

    _, err := eeprom.WriteAt([]byte("Hello, World!"), DataAddr)
    if err != nil{
        panic(err)
    }
    
    // Create a buffer to read the data into
    // We need to define a buffer with the correct size of the bytes we want to read.
    buf := make([]byte, len("Hello, World!"))
    _, err := eeprom.ReadAt(buf, DataAddr)
    if err != nil {
        panic(err)
    }

    fmt.Println(buf) // Hello, World!
}
```

This example will store a value, `Hello, World!`, into the address `0x10`.
