---
title: "Windows support"
weight: 4
description: |
  How to use TinyGo to create standard Windows executables.
---

TinyGo also lets you compile programs for Windows systems.

For cross compiling, you can use `GOOS` and `GOARCH` as usual. For example, you can cross compile the `examples/serial` example from a Linux host targeting Windows with the following command:

    GOOS=windows tinygo build -o serial.exe examples/serial
