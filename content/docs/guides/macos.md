---
title: "macOS support"
weight: 4
description: |
  How to use TinyGo to create standard macOS executables.
---

Using TinyGo you can compile programs for macOS systems.

For cross compiling, you can use `GOOS`. For example, you can cross compile the `examples/serial` example from a Linux host targeting macOS with the following command:

    GOOS=darwin tinygo build -o serial examples/serial
