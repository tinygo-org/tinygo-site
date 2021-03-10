---
title: "Linux support"
weight: 3
---

TinyGo also lets you compile programs for Linux systems, both 32-bit and 64-bit, on both x86 and ARM architectures.

For cross compiling, you can use `GOOS` and `GOARCH` as usual. For example, you can cross compile the `examples/serial` example from a Linux host to a Raspberry Pi with the following command:

    GOARCH=arm tinygo build -o serial examples/serial

This currently requires having the right cross compiler installed. On Debian you can do that with the following command:

    sudo apt-get install gcc-arm-linux-gnueabihf

This limitation should be lifted in a future version.
